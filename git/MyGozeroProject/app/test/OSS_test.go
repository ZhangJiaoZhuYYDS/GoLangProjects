package main

import (
	"fmt"
	"github.com/aliyun/aliyun-oss-go-sdk/oss"
	"os"
	"testing"
)

func TestOSS(t *testing.T) {
	// 创建OSSClient实例。
	// yourEndpoint填写Bucket对应的Endpoint，以华东1（杭州）为例，填写为https://oss-cn-hangzhou.aliyuncs.com。其它Region请按实际情况填写。
	// 阿里云账号AccessKey拥有所有API的访问权限，风险很高。强烈建议您创建并使用RAM用户进行API访问或日常运维，请登录RAM控制台创建RAM用户。
	fmt.Println("OSS Go SDK Version: ", oss.Version)
	client, err := oss.New("https://oss-cn-beijing.aliyuncs.com", "LTAI5tSk3oEZqinKhFWHx9ym", "4Pznq6IvLtcHzP3X6WtZdYArIqDpAa")
	if err != nil {
		fmt.Println("Error:", err)
		os.Exit(-1)
	}
	//列举所有的存储空间
	marker := ""
	for {
		lsRes, err := client.ListBuckets(oss.Marker(marker))
		if err != nil {
			fmt.Println("Error:", err)
			os.Exit(-1)
		}

		// 默认情况下一次返回100条记录。
		for _, bucket := range lsRes.Buckets {
			fmt.Println("Bucket: ", bucket.Name)
		}

		if lsRes.IsTruncated {
			marker = lsRes.NextMarker
		} else {
			break
		}
	}
	//判断存储空间是否存在
	//yourBucketName存储空间名称
	BucketName := "go-zero-project"
	isExist, err := client.IsBucketExist(BucketName)
	if err != nil {
		fmt.Println("Error:", err)
		os.Exit(-1)
	}
	fmt.Println("IsBucketExist result : ", isExist)
	//文件上传，文件上传有简单上传，追加上传，断点续传上传，分片上传
	if !isExist {
		os.Exit(-1)
	}
	bucket, err := client.Bucket(BucketName) //注意此处不要写错，写错的话，err让然是nil，我们应该需要先判断一下是否存在
	if err != nil {
		fmt.Println("Error:", err)
		os.Exit(-1)
	}
	err = bucket.PutObjectFromFile("exampledir/产品经理必懂的技术那点事儿 (唐韧) (z-lib.org).pdf", "E:\\demo.txt")
	if err != nil {
		fmt.Println("Error:", err)
		os.Exit(-1)
	}
	file, _ := os.Open("E:\\demo.txt")
	err = bucket.PutObject("demo", file)
	if err != nil {
		fmt.Println(err)
	}
}