package usecase

import "github.com/ko44d/hrmos-time-aggregator/pkg/repository"

type AuthenticationTokenUsecase interface {
	Get() (*repository.AuthenticationToken, error)
}

type authenticationTokenUsecase struct {
	atr repository.AuthenticationTokenRepository
}

func NewAuthenticationTokenUsecase(atr repository.AuthenticationTokenRepository) AuthenticationTokenUsecase {
	return &authenticationTokenUsecase{atr: atr}
}

func (atu *authenticationTokenUsecase) Get() (*repository.AuthenticationToken, error) {
	return atu.atr.Get()
}
