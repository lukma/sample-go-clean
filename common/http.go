package common

import (
	"io"
	"net/http"
	"os"

	"github.com/gin-gonic/gin"
	mgo "gopkg.in/mgo.v2"
)

// FileHandler - Handle file upload to server.
// It returns destination file that already store and any write error encountered.
func FileHandler(c *gin.Context, directory string, query string) (string, error) {
	var destination string

	formFile, err := c.FormFile(query)
	if err == nil {
		file, _ := formFile.Open()
		defer file.Close()

		var randomText string
		randomText, err := GenerateRandomText()
		if err == nil {
			destination = os.Getenv("GOPATH") + "/upload/sample-go-clean/post/" +
				directory + "/" +
				randomText + "-" + formFile.Filename
		}

		dst, _ := os.Create(destination)
		defer dst.Close()

		io.Copy(dst, file)
	}

	return destination, err
}

// JsonHandler - Write json response to client.
func JsonHandler(c *gin.Context, err error, data map[string]interface{}) {
	var status int
	var json gin.H

	if err != nil {
		switch err {
		case mgo.ErrNotFound:
			status = http.StatusNotFound
			json = gin.H{"error_message": err.Error()}
			break
		default:
			status = http.StatusBadRequest
			json = gin.H{"error_message": err.Error()}
			break
		}
	} else {
		status = http.StatusOK
		json = data
	}

	c.Writer.Header().Set("Content-Type", "application/json")
	c.JSON(status, json)
}
