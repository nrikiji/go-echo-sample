package store

import (
	"context"

	"firebase.google.com/go/auth"
)

type firebaseClient interface {
	VerifyIDToken(context.Context, string) (*auth.Token, error)
}

type AuthStore struct {
	client firebaseClient
}

func NewAuthStore(client firebaseClient) *AuthStore {
	return &AuthStore{client: client}
}

func (as *AuthStore) VerifyIdToken(idToken string) (string, error) {
	token, err := as.client.VerifyIDToken(context.Background(), idToken)
	if err != nil {
		return "", err
	}
	return token.UID, nil
}
