package common

import (
	"context"

	firebase "firebase.google.com/go"
	"google.golang.org/api/option"
)

func GetFirebaseApp(path string) (*firebase.App, error) {
	opt := option.WithCredentialsFile(path)
	return firebase.NewApp(context.Background(), nil, opt)
}
