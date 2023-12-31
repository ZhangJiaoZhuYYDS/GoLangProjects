package logic

import (
	"GopherTok/common/errorx"
	"context"
	"fmt"
	"github.com/pkg/errors"
	"gorm.io/gorm"
	"strconv"

	"GopherTok/server/relation/rpc/internal/svc"
	"GopherTok/server/relation/rpc/pb"

	"github.com/zeromicro/go-zero/core/logx"
)

type GetFriendCountLogic struct {
	ctx    context.Context
	svcCtx *svc.ServiceContext
	logx.Logger
}

func NewGetFriendCountLogic(ctx context.Context, svcCtx *svc.ServiceContext) *GetFriendCountLogic {
	return &GetFriendCountLogic{
		ctx:    ctx,
		svcCtx: svcCtx,
		Logger: logx.WithContext(ctx),
	}
}

func (l *GetFriendCountLogic) GetFriendCount(in *pb.GetFriendCountReq) (*pb.GetFriendCountResp, error) {
	count := l.svcCtx.Rdb.HGet(l.ctx, "cache:gopherTok:follow:friendCount", fmt.Sprintf("%d:friendCount", in.Userid))
	if count.Err() != nil {
		if count.Err().Error() != "redis: nil" {
			return nil,
				errors.Wrapf(errorx.NewDefaultError("redis get err:"+count.Err().Error()), "redis get err ：%v", count.Err())
		}

		if count.Err().Error() == "redis: nil" {
			//如果redis中没有则从mysql中拉取并更新至redis中
			friend := []pb.FollowSubject{}
			err := l.svcCtx.MysqlDb.WithContext(l.ctx).Table("follow_subject").
				Where("user_id = ?", in.Userid).Find(&friend).Error
			if err != nil {
				return nil,
					errors.Wrapf(errorx.NewDefaultError("mysql get err:"+err.Error()), "mysql get err ：%v", err)
			}
			countMysql := 0
			for _, v := range friend {

				err := l.svcCtx.MysqlDb.WithContext(l.ctx).Table("follow_subject").
					Where("user_id = ? AND follower_id = ?", v.FollowerId, in.Userid).First(&pb.FollowSubject{}).Error
				if err != nil {
					if err != gorm.ErrRecordNotFound {
						return nil,
							errors.Wrapf(errorx.NewDefaultError("mysql get err:"+err.Error()), "mysql get err ：%v", err)
					}
				}
				countMysql++
			}
			l.svcCtx.Rdb.HSet(l.ctx, "friendCount", fmt.Sprintf("%d:friendCount", in.Userid), strconv.Itoa(countMysql))

			return &pb.GetFriendCountResp{StatusCode: 0,
				StatusMsg: "get friendCount successfully",
				Count:     int64(countMysql)}, nil
		}

	}
	countInt, _ := strconv.ParseInt(count.Val(), 10, 64)

	return &pb.GetFriendCountResp{StatusCode: 0,
		StatusMsg: "get friendCount successfully",
		Count:     countInt}, nil
}
