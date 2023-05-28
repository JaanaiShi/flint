package github_util

import (
	"context"
	"github.com/google/go-github/v52/github"
	"time"
)

type GithubUtil struct {
	ctx    context.Context
	client *github.Client
}

func NewGithubUtil(ctx context.Context, token string) *GithubUtil {
	return &GithubUtil{
		ctx:    ctx,
		client: github.NewTokenClient(ctx, token),
	}
}

type GithubProjectInfo struct {
	Name          string     `json:"name"`
	FullName      string     `json:"full_name"`
	Description   string     `json:"description"`
	DefaultBranch string     `json:"master_branch"`
	CreatedAt     *time.Time `json:"created_at"`
	PushedAt      *time.Time `json:"pushed_at"`
	UpdatedAt     *time.Time `json:"updated_at"`
}

func (g *GithubUtil) ListPublicRepository(since int64) ([]GithubProjectInfo, error) {
	var (
		option github.RepositoryListAllOptions
		result []GithubProjectInfo
	)

	if since != 0 {
		option.Since = since
	}

	repositories, _, err := g.client.Repositories.ListAll(g.ctx, &option)
	if err != nil {
		return result, err
	}

	for i := 0; i < len(repositories); i++ {
		createdAt := repositories[i].GetCreatedAt()
		pushedAt := repositories[i].GetPushedAt()
		updatedAt := repositories[i].GetUpdatedAt()

		result = append(result, GithubProjectInfo{
			Name:          repositories[i].GetName(),
			FullName:      repositories[i].GetFullName(),
			Description:   repositories[i].GetDescription(),
			DefaultBranch: repositories[i].GetDefaultBranch(),
			CreatedAt:     createdAt.GetTime(),
			PushedAt:      pushedAt.GetTime(),
			UpdatedAt:     updatedAt.GetTime(),
		})
	}

	return result, nil
}

type UserInfo struct {
	Login             string `json:"login"`
	Id                int64  `json:"id"`
	NodeId            string `json:"node_id"`
	AvatarUrl         string `json:"avatar_url"`
	GravatarId        string `json:"gravatar_id"`
	Url               string `json:"url"`
	HtmlUrl           string `json:"html_url"`
	FollowersUrl      string `json:"followers_url"`
	FollowingUrl      string `json:"following_url"`
	GistsUrl          string `json:"gists_url"`
	StarredUrl        string `json:"starred_url"`
	SubscriptionsUrl  string `json:"subscriptions_url"`
	OrganizationsUrl  string `json:"organizations_url"`
	ReposUrl          string `json:"repos_url"`
	EventsUrl         string `json:"events_url"`
	ReceivedEventsUrl string `json:"received_events_url"`
	Type              string `json:"type"`
	SiteAdmin         bool   `json:"site_admin"`
}

func (g GithubUtil) ListUsers(since int64, page, limit int) (r []UserInfo, err error) {

	opt := github.UserListOptions{
		ListOptions: github.ListOptions{
			Page:    page,
			PerPage: limit,
		},
	}

	if since != 0 {
		opt.Since = since
	}

	data, _, err := g.client.Users.ListAll(g.ctx, &opt)
	if err != nil {
		return
	}

	for i := 0; i < len(data); i++ {
		r = append(r, UserInfo{
			Login:        data[i].GetLogin(),
			Id:           data[i].GetID(),
			AvatarUrl:    data[i].GetAvatarURL(),
			FollowersUrl: data[i].GetFollowingURL(),
			ReposUrl:     data[i].GetReposURL(),
		})
	}

	return

}

func (g GithubUtil) ListFowerrs(username string, page, limit int) (r []UserInfo, err error) {

	opt := github.ListOptions{
		Page:    page,
		PerPage: limit,
	}

	data, _, err := g.client.Users.ListFollowers(g.ctx, username, &opt)
	if err != nil {
		return
	}

	for i := 0; i < len(data); i++ {
		r = append(r, UserInfo{
			Login:        data[i].GetLogin(),
			Id:           data[i].GetID(),
			AvatarUrl:    data[i].GetAvatarURL(),
			FollowersUrl: data[i].GetFollowingURL(),
			ReposUrl:     data[i].GetReposURL(),
		})
	}

	return
}
