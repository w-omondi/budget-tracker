package services

import (
	"context"
	"errors"
	"log"
	"path/filepath"
	"strings"

	"firebase.google.com/go/v4"
	"firebase.google.com/go/v4/auth"
	"github.com/w-omondi/budget-tracker.git/internal/configs"

	//"firebase.google.com/go/v4/messaging"
	"google.golang.org/api/option"
)

var FirebaseApp *firebase.App

type FirebaseAuthManager interface {
	VerifyToken(string) (*auth.Token, error)
	ExtractTokenMetadata(string) (string, error)
}

type firebaseAuthManager struct {
	app        *firebase.App
	authClient *auth.Client
}

func NewFirebaseAuthManager() FirebaseAuthManager {
	configs.CheckEnvs("SERVICE_ACCOUNT_JSON")

	if FirebaseApp != nil {
		return &firebaseAuthManager{app: FirebaseApp}
	}

	serviceJsonFile := configs.GetEnv("SERVICE_ACCOUNT_JSON")
	serviceJsonFilePath := filepath.Clean(serviceJsonFile)
	opt := option.WithCredentialsFile(serviceJsonFilePath)

	app, err := firebase.NewApp(context.Background(), nil, opt)
	if err != nil {
		log.Fatalf("error initializing app: %v", err)
	}

	authClient, err := app.Auth(context.Background())
	if err != nil {
		log.Fatalf("error initializing auth client: %v", err)
	}

	return &firebaseAuthManager{
		app:        app,
		authClient: authClient,
	}
}

func (fm *firebaseAuthManager) ExtractTokenMetadata(authString string) (string, error) {
	split := strings.Split(authString, "Bearer ")
	if len(split) != 2 {
		return "", errors.New("invalid auth string")
	}
	token := split[1]
	if token == "" {
		return "", errors.New("invalid auth string")
	}
	return token, nil
}

func (fm *firebaseAuthManager) VerifyToken(token string) (*auth.Token, error) {
	decodedToken, err := fm.authClient.VerifyIDToken(context.Background(), token)
	if err != nil {
		return nil, err
	}
	return decodedToken, nil
}
