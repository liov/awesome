package api

import (
	"github.com/gin-gonic/gin"
	"test/example/cherrywithinit/proto"
	"test/example/cherrywithinit/service"
	"net/http"
)

func GinRegister(app *gin.Engine) {
	_ = user.RegisterUserServiceHandlerServer(app, service.GetUserService())
	//oauth.RegisterOauthServiceHandlerServer(app, service.GetOauthService())
	app.StaticFS("/oauth/login", http.Dir("./static/login.html"))
	app.GET("/api/test/:id", service.Test)
}
