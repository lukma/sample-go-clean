package domain

import (
	"time"

	"github.com/gin-gonic/gin"
	"gopkg.in/mgo.v2/bson"
)

// AuthEntity - Auth entity data schema.
type AuthEntity struct {
	ID            bson.ObjectId `bson:"_id,omitempty" json:"id"`
	Username      string        `bson:"username,omitempty" json:"username"`
	Password      string        `bson:"password,omitempty" json:"-"`
	FullName      string        `bson:"fullname,omitempty" json:"fullname"`
	Email         string        `bson:"email,omitempty" json:"email"`
	Phone         string        `bson:"phone,omitempty" json:"phone"`
	FacebookToken string        `bson:"facebook_token,omitempty" json:"facebook_token"`
	GoogleToken   string        `bson:"google_token,omitempty" json:"google_token"`
	FcmID         string        `bson:"fcm_id,omitempty" json:"fcm_id"`
	CreatedDate   time.Time     `bson:"created_date,omitempty" json:"created_date"`
}

// AuthRepository - Auth repository.
type AuthRepository interface {

	// GetAuthByID - Get selected auth by id from database.
	// It returns selected auth and any write error encountered.
	GetAuthByID(id string) (AuthEntity, error)

	// GetAuthByUsernameOrEmail - Get selected auth by username or email from database.
	// It returns selected auth and any write error encountered.
	GetAuthByUsernameOrEmail(usernameOrEmail string, password string) (AuthEntity, error)

	// GetAuthByThirdParty - Get selected auth by fcm id from database.
	// It returns selected auth and any write error encountered.
	GetAuthByThirdParty(thirdParty string, token string) (AuthEntity, error)

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

	// ConnectWithThirdPartyHandler - Handle http request connect third party.
	ConnectWithThirdPartyHandler(context *gin.Context)

	// RefreshTokenHandler - Handle http request refresh token.
	RefreshTokenHandler(context *gin.Context)
}

// LoginForm - Login form schema.
type LoginForm struct {
	Username string `form:"username" binding:"required"`
	Password string `form:"password" binding:"required"`
}

// RegisterForm - Register form schema.
type RegisterForm struct {
	Username string `form:"username" binding:"required"`
	Password string `form:"password" binding:"required"`
	FullName string `form:"fullname" binding:"required"`
	Email    string `form:"email" binding:"required"`
}

// RegisterForm - Register form schema.
type ConnectWithThirdPartyForm struct {
	ThirdParty string `form:"fcm_id" binding:"required"`
	Token      string `form:"token" binding:"required"`
}

// RefreshTokenForm - Refresh token form schema.
type RefreshTokenForm struct {
	RefreshToken string `form:"refresh_token" binding:"required"`
}
