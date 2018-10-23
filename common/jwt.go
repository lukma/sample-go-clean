package common

import (
	"fmt"
	"net/http"
	"os"
	"time"

	jwt "github.com/dgrijalva/jwt-go"
	"github.com/dgrijalva/jwt-go/request"
	"github.com/gin-gonic/gin"
)

func createToken(clientID string, exp int64) (string, error) {
	token := jwt.NewWithClaims(jwt.SigningMethodHS256, &jwt.MapClaims{
		"client": clientID,
		"exp":    exp,
	})
	tokenString, err := token.SignedString([]byte(os.Getenv("AUTH_SECRET_KEY")))

	return tokenString, err
}

func parseClaimsFromToken(token *jwt.Token) (map[string]interface{}, error) {
	if claims, ok := token.Claims.(jwt.MapClaims); ok && token.Valid {
		return map[string]interface{}{
			"client": claims["client"].(string),
			"exp":    int64(claims["exp"].(float64)),
		}, nil
	}

	return map[string]interface{}{}, fmt.Errorf("Token is invalid")
}

// ApplyJwt - Apply jwt authentication for access route.
func ApplyJwt(c *gin.Context) {
	token, err := request.ParseFromRequest(c.Request, request.OAuth2Extractor, func(token *jwt.Token) (interface{}, error) {
		return ([]byte(os.Getenv("AUTH_SECRET_KEY"))), nil
	})

	if err == nil {
		claimsToken := map[string]interface{}{}
		if claimsToken, err = parseClaimsFromToken(token); err == nil {
			c.Set("claims", claimsToken)
			c.Next()
		}
	}

	if err != nil {
		c.AbortWithStatusJSON(http.StatusUnauthorized, gin.H{"error_message": err.Error()})
	}
}

// GenerateToken - Generate token that assigned to client
// It returns token jwt and any write error encountered.
func GenerateToken(clientID string) (map[string]interface{}, error) {
	var result map[string]interface{}

	token, err := createToken(clientID, time.Now().Add(time.Hour*1).Unix())

	var refreshToken string
	if err == nil {
		refreshToken, err = createToken(clientID, time.Now().Add(time.Hour*168).Unix())
	}

	if err == nil {
		result = map[string]interface{}{
			"access_token":  token,
			"token_type":    "Bearer",
			"refresh_token": refreshToken,
		}
	}

	return result, err
}

// GetClaimsFromToken - Get claim from valid token
// It returns claim jwt and any write error encountered.
func GetClaimsFromToken(tokenString string) (map[string]interface{}, error) {
	token, err := jwt.Parse(tokenString, func(token *jwt.Token) (interface{}, error) {
		return ([]byte(os.Getenv("AUTH_SECRET_KEY"))), nil
	})

	if err == nil {
		return parseClaimsFromToken(token)
	}

	return map[string]interface{}{}, err
}
