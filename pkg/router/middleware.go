package router

import (
	"GithubTools/pkg/env"
	"github.com/Gebes/there/v2"
)

func KeyMiddleware(request there.HttpRequest, next there.HttpResponse) there.HttpResponse {
	authorization := request.Headers.GetDefault(there.RequestHeaderAuthorization, "")
	if authorization != env.Key {
		return there.Error(there.StatusUnauthorized, "no authorization header provided")
	}
	return next
}
