package github_util

import (
	"context"
	"fmt"
	"log"
	"testing"
)

var (
	git = NewGithubUtil(context.Background(), "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")
)

func TestGithubUtil_ListPublicRepository(t *testing.T) {

	projectList, err := git.ListPublicRepository(0)
	if err != nil {
		log.Fatal(err)
	}

	for i := 0; i < len(projectList); i++ {
		fmt.Println(projectList[i])
	}

}

func TestGithubUtil_ListUsers(t *testing.T) {

	r, err := git.ListUsers(0, 1, 100)
	if err != nil {
		log.Fatal(err)
	}

	for i := 0; i < len(r); i++ {
		fmt.Println(r[i].Login, r[i].Id, r[i].AvatarUrl, r[i].FollowersUrl)
	}
}
