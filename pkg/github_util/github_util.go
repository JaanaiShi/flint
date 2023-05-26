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
