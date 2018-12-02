package account

import (
	"github.com/gin-gonic/gin"
	"github.com/lukma/sample-go-clean/account/service"
)

func ApplyAccountRoute(router *gin.RouterGroup) {
	group := router.Group("/account/auth")
	{
		group.POST("/login", service.NewAuthService().LoginHandler)
		group.POST("/register", service.NewAuthService().RegisterHandler)
		group.POST("/connectWithThirdParty", service.NewAuthService().ConnectWithThirdPartyHandler)
		group.POST("/refreshToken", service.NewAuthService().RefreshTokenHandler)
	}
}
