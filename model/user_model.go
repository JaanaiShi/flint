package model

type UserInfo struct {
	Id        int64  `json:"id" gorm:"column:id"`
	Username  string `json:"username" gorm:"column:username"`
	AvatarUrl string `json:"avatar_url" gorm:"column:avatar_url"`
	GithubId  int64  `json:"github_id" gorm:"column:github_id"` // 设置唯一索引

}
