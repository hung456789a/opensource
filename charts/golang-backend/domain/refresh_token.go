package domain

import "context"

type RefreshToken struct {
	AccessToken  string `json:"accessToken"`
	RefreshToken string `json:"refreshToken"`
}

type RefresTokenUseCase interface {
	GetUserByID(c context.Context, id string) (User, error)
	CreateAccessToken(user *User, secret string, expiry int) (accessToken string, err error)
	CreateRefreshToken(user *User, secret string, expiry int) (refreshToken string, err error)
	ExtractIDFromToken(requestToken string, secret string) (string, error)
}
