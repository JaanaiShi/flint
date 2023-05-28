package model

import "time"

type Repository struct {
	Id            int64      `json:"id"`
	GithubId      int64      `json:"github_id"`
	Name          string     `json:"name"`
	FullName      string     `json:"full_name"`
	Description   string     `json:"description"`
	DefaultBranch string     `json:"master_branch"`
	CreatedAt     *time.Time `json:"created_at"` // 时间可以设置为Null
	PushedAt      *time.Time `json:"pushed_at"`
	UpdatedAt     *time.Time `json:"updated_at"`
}
