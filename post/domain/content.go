package domain

import (
	"time"

	"github.com/gin-gonic/gin"
	"gopkg.in/mgo.v2/bson"
)

// ContentEntity - Content entity data schema.
type ContentEntity struct {
	ID          bson.ObjectId `bson:"_id,omitempty" json:"id"`
	Title       string        `bson:"title,omitempty" json:"title"`
	Thumbnail   string        `bson:"thumbnail,omitempty" json:"thumbnail"`
	Content     string        `bson:"content,omitempty" json:"content"`
	CreatedDate time.Time     `bson:"created_date,omitempty" json:"created_date"`
}

// ContentRepository - Content repository.
type ContentRepository interface {

	// CountContent - Count contents from database.
	// It returns count contents and any write error encountered.
	CountContent(query string) (int, error)

	// GetContents - Get many content from database.
	// It returns many content and any write error encountered.
	GetContents(limit int, offset int, query string, sort string) ([]ContentEntity, error)

	// GetContent - Get selected content by id from database.
	// It returns selected content and any write error encountered.
	GetContent(id string) (ContentEntity, error)

	// CreateContent - Insert new content to database.
	// It returns one content and any write error encountered.
	CreateContent(obj ContentEntity) (string, error)

	// UpdateContent - Update selected content to database.
	// It returns any write error encountered.
	UpdateContent(id string, obj ContentEntity) error

	// DeleteContent - Delete selected content to database.
	// It returns any write error encountered.
	DeleteContent(id string) error
}

type ContentService interface {

	// GetContentsHandler - Handle http request get many content.
	GetContentsHandler(context *gin.Context)

	// GetContentHandler - Handle http request get selected content.
	GetContentHandler(context *gin.Context)

	// CreateContentHandler - Handle http request create selected content.
	CreateContentHandler(context *gin.Context)

	// UpdateContentHandler - Handle http request update selected content.
	UpdateContentHandler(context *gin.Context)

	// DeleteContentHandler - Handle http request delete selected content.
	DeleteContentHandler(context *gin.Context)
}

// CreateContentForm - Create post form schema.
type CreateContentForm struct {
	Title   string `form:"title" binding:"required"`
	Content string `form:"content" binding:"required"`
}

// UpdateContentForm - Update post form schema.
type UpdateContentForm struct {
	Title   string `form:"title"`
	Content string `form:"content"`
}
