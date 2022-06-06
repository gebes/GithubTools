package router

import (
	"GithubTools/pkg/check"
	"fmt"
	"github.com/Gebes/there/v2"
	"strings"
)

func UserNoFollowBackGet(request there.HttpRequest) there.HttpResponse {
	user := request.RouteParams.GetDefault("user", "")
	users, err := check.NoFollowBacks(user)
	if err != nil {
		return there.Error(there.StatusInternalServerError, fmt.Errorf("could not run follow back check: %v", err))
	}

	return there.Json(there.StatusOK, users)
}

func ContainsFileGet(request there.HttpRequest) there.HttpResponse {
	owner := request.RouteParams.GetDefault("owner", "")
	repo := request.RouteParams.GetDefault("repo", "")
	filesParam, ok := request.Params.Get("files")
	if !ok {
		return there.Error(there.StatusBadRequest, "no files parameter provided")
	}

	files, err := check.ContainsFiles(owner, repo, strings.Split(filesParam, ","))
	if err != nil {
		return there.Error(there.StatusInternalServerError, fmt.Errorf("could not run contains files check: %v", err))
	}

	return there.Json(there.StatusOK, files)
}
