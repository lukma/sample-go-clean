package domain

import (
	"time"

	"github.com/gin-gonic/gin"
	"gopkg.in/mgo.v2/bson"
)

// AuthEntity - Auth entity data schema.
type AuthEntity struct {
	ID            bson.ObjectId `bson:"_id,omitempty" json:"id"`
	FaID          string        `bson:"fa_id,omitempty" json:"fa_id"`
	FcmID         string        `bson:"fcm_id,omitempty" json:"fcm_id"`
	Email         string        `bson:"email,omitempty" json:"email"`
	Phone         string        `bson:"phone,omitempty" json:"phone"`
	FacebookToken string        `bson:"facebook_token,omitempty" json:"facebook_token"`
	GoogleToken   string        `bson:"google_token,omitempty" json:"google_token"`
	CreatedDate   time.Time     `bson:"created_date,omitempty" json:"created_date"`
}

// AuthRepository - Auth repository.
type AuthRepository interface {

	// GetAuth - Get selected auth by fcm id from database.
	// It returns selected auth and any write error encountered.
	GetAuth(faID string) (AuthEntity, error)

	// CreateAuth - Insert new auth to database.
	// It returns one auth and any write error encountered.
	CreateAuth(obj AuthEntity) (string, error)

	// UpdateAuth - Update selected auth to database.
	// It any write error encountered.
	UpdateAuth(id string, obj AuthEntity) error
}

type AuthService interface {

	// LoginHandler - Handle http request login.
	LoginHandler(context *gin.Context)

	// RegisterHandler - Handle http request register.
	RegisterHandler(context *gin.Context)

	// RefreshTokenHandler - Handle http request refresh token.
	RefreshTokenHandler(context *gin.Context)
}

// LoginForm - Login form schema.
type LoginForm struct {
	FaID  string `form:"fa_id" binding:"required"`
	FcmID string `form:"fcm_id" binding:"required"`
}

// RegisterForm - Register form schema.
type RegisterForm struct {
	FaID          string `form:"fa_id" binding:"required"`
	FcmID         string `form:"fcm_id" binding:"required"`
	FacebookToken string `form:"facebook_token"`
	GoogleToken   string `form:"google_token"`
}

// RefreshTokenForm - Refresh token form schema.
type RefreshTokenForm struct {
	RefreshToken string `form:"refresh_token" binding:"required"`
}
