package router

import (
	"context"
	"github.com/Gebes/there/v2"
	"github.com/Gebes/there/v2/middlewares"
)

var router = there.NewRouter()

func Listen() error {

	router.Use(middlewares.Cors(middlewares.AllowAllConfiguration()))

	router.Get("/no-follow-back/:user", UserNoFollowBackGet)
	router.Get("/contains-file/:owner/:repo", ContainsFileGet)

	return router.Listen(8080)
}

func Shutdown() error {
	return router.Server.Shutdown(context.Background())
}
