package db

import (
	"context"
	firebase "firebase.google.com/go/v4"
	"firebase.google.com/go/v4/auth"
	"google.golang.org/api/option"
	"os"
)

var FS_PATH_PREFIX = os.Getenv("FS_PATH_PREFIX")

func NewAuthClient() *auth.Client {
	opt := option.WithCredentialsFile(FS_PATH_PREFIX + "firebase/service-account.json")
	app, err := firebase.NewApp(context.Background(), nil, opt)
	if err != nil {
		panic(err)
	}
	authClient, err := app.Auth(context.Background())
	if err != nil {
		panic(err)
	}
	return authClient
}

func VerifyIDToken(authClient *auth.Client, token string) (*auth.Token, error) {
	return authClient.VerifyIDToken(context.Background(), token)
}

func GetUser(authClient *auth.Client, token string) (*auth.UserRecord, error) {
	return authClient.GetUser(context.Background(), token)
}
