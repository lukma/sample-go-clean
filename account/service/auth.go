package service

import (
	"github.com/gin-gonic/gin"
	"github.com/lukma/sample-go-clean/account/data"
	"github.com/lukma/sample-go-clean/account/domain"
	"github.com/lukma/sample-go-clean/common"
)

type service struct {
	repository      domain.AuthRepository
	responseHandler func(c *gin.Context, err error, data map[string]interface{})
}

func NewAuthService() domain.AuthService {
	repository := data.NewAuthRepository()
	responseHandler := common.JsonHandler

	return &service{
		repository:      repository,
		responseHandler: responseHandler,
	}
}

func (service *service) LoginHandler(c *gin.Context) {
	var form domain.LoginForm
	err := c.Bind(&form)

	var auth domain.AuthEntity
	if err == nil {
		auth, err = service.repository.GetAuth(form.FaID)
	}

	if err == nil {
		err = service.repository.UpdateAuth(auth.ID.Hex(), domain.AuthEntity{
			FcmID: form.FcmID,
		})
	}

	var token map[string]interface{}
	if err == nil {
		token, err = common.GenerateToken(auth.ID.Hex())
	}

	service.responseHandler(c, err, gin.H{"token": token})
}

func (service *service) RegisterHandler(c *gin.Context) {
	var form domain.RegisterForm
	err := c.Bind(&form)

	id, err := service.repository.CreateAuth(domain.AuthEntity{
		FaID:          form.FaID,
		FcmID:         form.FcmID,
		FacebookToken: form.FacebookToken,
		GoogleToken:   form.GoogleToken,
	})

	service.responseHandler(c, err, gin.H{"data": id})
}

func (service *service) RefreshTokenHandler(c *gin.Context) {
	form := domain.RefreshTokenForm{}
	err := c.Bind(&form)

	claims := map[string]interface{}{}
	if err == nil {
		claims, err = common.GetClaimsFromToken(form.RefreshToken)
	}

	auth := domain.AuthEntity{}
	if err == nil {
		service.repository.GetAuth(claims["client"].(string))
	}

	token := map[string]interface{}{}
	if err == nil {
		token, err = common.GenerateToken(auth.ID.Hex())
	}

	service.responseHandler(c, err, gin.H{"token": token})
}
