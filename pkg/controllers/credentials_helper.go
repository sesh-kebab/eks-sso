package controllers

import (
	"context"

	"github.com/quintilesims/eks-sso/pkg/models"
)

// advised against using basic types as a key for context values
type key int

const (
	credentialsCtxKey     key    = 0
	credentialsSessionKey string = "credentials-key"
)

func newContextWithCredentials(ctx context.Context, iamCreds models.IAMCredentials) context.Context {
	return context.WithValue(ctx, credentialsCtxKey, iamCreds)
}

func iamCredsFromContext(ctx context.Context) (models.IAMCredentials, bool) {
	creds := ctx.Value(credentialsCtxKey)
	iamCreds, ok := creds.(models.IAMCredentials)
	return iamCreds, ok
}
