package check

import (
	"GithubTools/pkg/client"
	"context"
	"fmt"
	"github.com/Gebes/there/v2"
	"github.com/google/go-github/v44/github"
	"golang.org/x/sync/errgroup"
	"sort"
	"strings"
)

func NoFollowBacks(user string) ([]*github.User, error) {

	toTest, _, err := client.Following(user)
	if err != nil {
		return nil, err
	}

	var noFollowBacks []*github.User
	errs, _ := errgroup.WithContext(context.Background())

	for _, currentUser := range toTest {
		currentUser := currentUser
		errs.Go(func() error {
			currentUserFollowings, _, err := client.Following(*currentUser.Login)
			if err != nil {
				return err
			}

			contains := false
			for _, following := range currentUserFollowings {
				if strings.ToLower(*following.Login) == strings.ToLower(user) {
					contains = true
					break
				}
			}
			if !contains {
				noFollowBacks = append(noFollowBacks, currentUser)
			}
			return nil
		})
	}

	err = errs.Wait()
	if err != nil {
		return nil, err
	}

	sort.Slice(noFollowBacks, func(i, j int) bool {
		return strings.Compare(safeUnwrap(noFollowBacks[i].Login), safeUnwrap(noFollowBacks[j].Login)) == 0
	})

	return noFollowBacks, nil
}

func ContainsFiles(owner, repo string, files []string) ([]string, error) {
	out := make([]string, 0)
	for _, file := range files {
		commits, resp, err := client.CommitsWithFile(owner, repo, file)
		if err != nil {
			return nil, err
		}
		if resp.StatusCode != there.StatusOK {
			return nil, fmt.Errorf("invalid response: %d", resp.StatusCode)
		}
		if len(commits) != 0 {
			out = append(out, file)
		}
	}
	return out, nil
}

func safeUnwrap(p *string) string {
	if p == nil {
		return ""
	}
	return *p
}
