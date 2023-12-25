package logic

import (
	"context"
	"encoding/json"
	"github.com/zeromicro/go-zero/core/threading"
	"miniZhihu/app/like/rpc/internal/types"

	"miniZhihu/app/like/rpc/internal/svc"
	"miniZhihu/app/like/rpc/like"

	"github.com/zeromicro/go-zero/core/logx"
)

type ThumbupLogic struct {
	ctx    context.Context
	svcCtx *svc.ServiceContext
	logx.Logger
}

func NewThumbupLogic(ctx context.Context, svcCtx *svc.ServiceContext) *ThumbupLogic {
	return &ThumbupLogic{
		ctx:    ctx,
		svcCtx: svcCtx,
		Logger: logx.WithContext(ctx),
	}
}

func (l *ThumbupLogic) Thumbup(in *like.ThumbupRequest) (*like.ThumbupResponse, error) {
	// todo: add your logic here and delete this line
	// 1. 查询是否点过赞
	// 2. 计算当前内容的总点赞数和点踩数
	msg := &types.ThumbupMsg{
		BizId:    in.BizId,
		ObjId:    in.ObjId,
		UserId:   in.UserId,
		LikeType: in.LikeType,
	}
	// 发送kafka消息，异步
	threading.GoSafe(func() {
		data,err := json.Marshal(msg)
		if err != nil {
			l.Logger.Errorf("[Thumbup] marshal msg: %+v error: %v", msg, err)
			return
		}
		err = l.svcCtx.KqPusherClient.Push(string(data))
		if err != nil {
			l.Logger.Errorf("[Thumbup] kq push data: %s error: %v", data, err)
		}
	})
	return &like.ThumbupResponse{}, nil
}
