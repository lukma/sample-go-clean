package post

import (
	"github.com/gin-gonic/gin"
	"github.com/lukma/sample-go-clean/post/service"
)

func ApplyContentRoute(router *gin.RouterGroup) {
	group := router.Group("/post/content")
	{
		group.GET("/", service.NewContentService().GetContentsHandler)
		group.GET("/:id", service.NewContentService().GetContentHandler)
		group.POST("/", service.NewContentService().CreateContentHandler)
		group.PUT("/:id", service.NewContentService().UpdateContentHandler)
		group.DELETE("/:id", service.NewContentService().DeleteContentHandler)
	}
}
