package server

import (
	"fmt"

	"github.com/clykk-user-atif/app/clykk-v1.0/server/internal/handlers"
	"github.com/clykk-user-atif/config"
	"github.com/fasthttp/router"
	"github.com/valyala/fasthttp"
)

// Index is the index handler
func Index(ctx *fasthttp.RequestCtx) {
	fmt.Fprint(ctx, "Not protected!\n")
}

func initRouter() (*router.Router, error) {

	handler, err := handlers.NewHandler(config.DBConnectionString())
	if err != nil {
		return nil, err
	}

	r := router.New()

	r.POST("/login", handler.Login)

	r.GET("/user/preferences", handlers.AuthUser(handler.UserPreferencesGET))
	r.POST("/user/preferences", handlers.AuthUser(handler.UserPreferencesPOST))
	r.PUT("/user/preferences", handlers.AuthUser(handler.UserPreferencesPUT))
	return r, nil

}
