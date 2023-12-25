// @Author zhangjiaozhu 2023/12/19 19:53:00
package service

import (
	"car/cache"
	"car/conf"
	"car/model"
	"car/pkg/e"
	"car/pkg/util"
	"car/serializer"
	"encoding/json"
	"fmt"
	logging "github.com/sirupsen/logrus"
	"math/rand"
	"strconv"
	"time"

	"github.com/tencentcloud/tencentcloud-sdk-go/tencentcloud/common"
	"github.com/tencentcloud/tencentcloud-sdk-go/tencentcloud/common/errors"
	"github.com/tencentcloud/tencentcloud-sdk-go/tencentcloud/common/profile"
	sms "github.com/tencentcloud/tencentcloud-sdk-go/tencentcloud/sms/v20210111" // 引入sms
)

type ValidPhoneService struct {
	OperationType int    `form:"operation_type" json:"operation_type"` // 操作类型 1:绑定手机 2:解绑手机
	Phone         string `form:"phone" json:"phone"`                   // 手机号
	Code          string `form:"code" json:"code"`                     // 验证码
}

type GetCodeService struct {
	Phone string `form:"phone" json:"phone"`
}

func (service *ValidPhoneService) Valid(authorization string) serializer.Response {
	var phone string
	var openid string
	code := e.Success
	phone = service.Phone
	claims, _ := util.ParseToken(authorization)
	openid = claims.OpenID
	// 根据操作类型执行绑定和解绑手机的操作
	if service.OperationType == 1 {
		// 1 绑定手机
		if err := cache.RedisClient.Get("code").Err(); err != nil {
			fmt.Println(err)
		}
		RedisCode := fmt.Sprintf("%s", cache.RedisClient.Get("code"))[10:]
		if RedisCode != service.Code {
			fmt.Println(RedisCode, service.Code)
			code = e.ErrorMsgCode
			return serializer.Response{
				Status: code,
				Msg:    e.GetMsg(code),
			}
		}
		if err := model.DB.Model(model.User{}).Where("open_id=?", openid).Update("phone", phone).Error; err != nil {
			logging.Info(err)
			code = e.ErrorDatabase
			return serializer.Response{
				Status: code,
				Msg:    e.GetMsg(code),
				Error:  err.Error(),
			}
		} else if service.OperationType == 2 {
			//2 解绑手机
			if err := cache.RedisClient.Get("code").Err(); err != nil {
				fmt.Println(err)
			}
			RedisCode := fmt.Sprintf("%s", cache.RedisClient.Get("code"))
			if RedisCode != service.Code {
				code = e.ErrorMsgCode
				return serializer.Response{
					Status: code,
					Msg:    e.GetMsg(code),
				}
			}
			if err := model.DB.Model(model.User{}).Where("open_id=?", openid).Update("phone", "").Error; err != nil {
				logging.Info(err)
				code = e.ErrorDatabase
				return serializer.Response{
					Status: code,
					Msg:    e.GetMsg(code),
					Error:  err.Error(),
				}
			}
		}
	}
	//获取该用户信息
	var user model.User
	if err := model.DB.First(&user).Where("open_id = ?", openid).Error; err != nil {
		logging.Info(err)
		code = e.ErrorDatabase
		return serializer.Response{
			Status: code,
			Msg:    e.GetMsg(code),
		}
	}
	return serializer.Response{
		Status: code,
		Data:   serializer.BuildUser(user),
		Msg:    e.GetMsg(code),
	}
}

// 发送短信
func (service *GetCodeService) SendMsg() serializer.Response {
	// 生成随机验证码并存入redis
	code := e.Success
	rand.Seed(time.Now().UnixNano())
	codeInt := rand.Intn(10000)
	codeString := strconv.Itoa(codeInt)
	if err := cache.RedisClient.Set("code", codeString, 0).Err(); err != nil {
		logging.Info(err) //将code存入redis中
	}
	if err := cache.RedisClient.Get("code").Err(); err != nil {
		logging.Info(err) //将code从redis拿出来
	}
	/* 必要步骤：
	 * 实例化一个认证对象，入参需要传入腾讯云账户密钥对secretId，secretKey。
	 * 这里采用的是从环境变量读取的方式，需要在环境变量中先设置这两个值。
	 * 您也可以直接在代码中写死密钥对，但是小心不要将代码复制、上传或者分享给他人，
	 * 以免泄露密钥对危及您的财产安全。
	 * SecretId、SecretKey 查询: https://console.cloud.tencent.com/cam/capi */
	credential := common.NewCredential( //创建第一个实例对象，登陆用
		conf.TxSecretId,
		conf.TxSecretKey,
	)
	/* 非必要步骤:
	 * 实例化一个客户端配置对象，可以指定超时时间等配置 */
	cpf := profile.NewClientProfile()
	/* SDK默认使用POST方法。
	 * 如果您一定要使用GET方法，可以在这里设置。GET方法无法处理一些较大的请求 */

	cpf.HttpProfile.ReqMethod = "POST"
	/* SDK有默认的超时时间，非必要请不要进行调整
	 * 如有需要请在代码中查阅以获取最新的默认值 */
	// cpf.HttpProfile.ReqTimeout = 5

	/* 指定接入地域域名，默认就近地域接入域名为 sms.tencentcloudapi.com ，也支持指定地域域名访问，例如广州地域的域名为 sms.ap-guangzhou.tencentcloudapi.com */
	cpf.HttpProfile.Endpoint = "sms.tencentcloudapi.com"
	cpf.SignMethod = "HmacSHA1" //SDK 默认用 TC3-HMAC-SHA256 进行签名，非必要请不要修改该字段
	/* 实例化要请求产品(以sms为例)的client对象
	 * 第二个参数是地域信息，可以直接填写字符串ap-guangzhou，支持的地域列表参考 https://cloud.tencent.com/document/api/382/52071#.E5.9C.B0.E5.9F.9F.E5.88.97.E8.A1.A8 */
	client, _ := sms.NewClient(credential, "ap-guangzhou", cpf)

	request := sms.NewSendSmsRequest()                                 //创建第二个实例对象，发送信息用
	request.SmsSdkAppId = common.StringPtr(conf.TxSmsSdkAppid)         //短信应用 ID: 在 [短信控制台] 添加应用后生成的实际 SDKAppID
	request.SignName = common.StringPtr(conf.TxSmsSign)                //短信签名内容: 使用 UTF-8 编码，必须填写已审核通过的签名，可登录 [短信控制台] 查看签名信息 */
	request.TemplateParamSet = common.StringPtrs([]string{codeString}) //放{1}参数，验证码
	request.TemplateId = common.StringPtr(conf.TxTemplateID)
	request.PhoneNumberSet = common.StringPtrs([]string{"+86" + service.Phone}) //发送的号码
	response, err := client.SendSms(request)                                    // 通过 client 对象调用想要访问的接口，需要传入请求对象
	// 处理异常
	if _, ok := err.(*errors.TencentCloudSDKError); ok {
		code = e.ErrorSendMsg
		return serializer.Response{
			Status: code,
			Data:   fmt.Sprintf("%s", err),
			Msg:    e.GetMsg(code),
		}
	}
	Info, err := json.Marshal(response.Response)
	InfoStr := fmt.Sprintf("%s", Info)
	if err != nil {
		code = e.ErrorSendMsg
		return serializer.Response{
			Status: code,
			Data:   err,
			Msg:    e.GetMsg(code),
		}
	}
	return serializer.Response{
		Status: code,
		Data:   InfoStr,
		Msg:    e.GetMsg(code),
	}
}
