package github_util

import (
	"context"
	"fmt"
	"log"
	"testing"
)

func TestGithubUtil_ListPublicRepository(t *testing.T) {
	git := NewGithubUtil(context.Background(), "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")

	projectList, err := git.ListPublicRepository(0)
	if err != nil {
		log.Fatal(err)
	}

	for i := 0; i < len(projectList); i++ {
		fmt.Println(projectList[i])
	}

}
