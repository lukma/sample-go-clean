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
		auth, err = service.repository.GetAuthByUsernameOrEmail(form.Username, form.Password)
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

	var id string
	if err == nil {
		id, err = service.repository.CreateAuth(domain.AuthEntity{
			Username: form.Username,
			Password: form.Password,
			FullName: form.FullName,
			Email:    form.Email,
		})
	}

	service.responseHandler(c, err, gin.H{"data": id})
}

func (service *service) ConnectWithThirdPartyHandler(c *gin.Context) {
	var form domain.ConnectWithThirdPartyForm
	err := c.Bind(&form)

	var auth domain.AuthEntity
	if err == nil {
		auth, err = service.repository.GetAuthByThirdParty(form.ThirdParty, form.Token)
	}

	var token map[string]interface{}
	if err == nil {
		token, err = common.GenerateToken(auth.ID.Hex())
	}

	service.responseHandler(c, err, gin.H{"token": token})
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
		service.repository.GetAuthByID(claims["client"].(string))
	}

	token := map[string]interface{}{}
	if err == nil {
		token, err = common.GenerateToken(auth.ID.Hex())
	}

	service.responseHandler(c, err, gin.H{"token": token})
}
