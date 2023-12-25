// Code generated by goctl. DO NOT EDIT.
package types

type BaseResponse struct {
	Code    int64  `json:"status_code"`
	Message string `json:"status_msg,omitempty"`
}

type AuthorInfo struct {
	ID              int64  `json:"id"`
	Name            string `json:"name"`
	FollowCount     int64  `json:"follow_count"`
	FollowerCount   int64  `json:"follower_count"`
	IsFollow        bool   `json:"is_follow"`
	Avatar          string `json:"avatar"`
	BackgroundImage string `json:"background_image"`
	Signature       string `json:"signature"`
	TotalFavorited  string `json:"total_favorited"`
	WorkCount       int64  `json:"work_count"`
	FavoriteCount   int64  `json:"favorite_count"`
}

type VideoInfo struct {
	ID            int64      `json:"id"`
	Author        AuthorInfo `json:"author"`
	PlayURL       string     `json:"play_url"`
	CoverURL      string     `json:"cover_url"`
	FavoriteCount int64      `json:"favorite_count"`
	CommentCount  int64      `json:"comment_count"`
	IsFavorite    bool       `json:"is_favorite"`
	Title         string     `json:"title"`
}

type VideoList struct {
	List []*VideoInfo `json:"video_list"`
}

type PublishVideoReq struct {
	Token string `form:"token"`
	Title string `form:"title"`
}

type PublishVideoResp struct {
	BaseResponse
}

type UserVideoListReq struct {
	Token  string `form:"token"`
	UserId string `form:"user_id"`
}

type UserVideoListResp struct {
	BaseResponse
	VideoList
}

type VideoListReq struct {
	LatestTime string `form:"latest_time,optional"`
	Token      string `form:"token,optional"`
}

type VideoListResp struct {
	BaseResponse
	NextTime int64 `json:"next_time"`
	VideoList
}
