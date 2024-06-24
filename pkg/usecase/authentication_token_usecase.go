package usecase

import "github.com/ko44d/hrmos-time-aggregator/pkg/repository"

type AuthenticationTokenUsecase interface {
	GetToken(apiKey, companyUrl string) (*repository.AuthenticationToken, error)
}

type authenticationTokenUsecase struct {
	atr repository.AuthenticationTokenRepository
}

func NewAuthenticationTokenUsecase(atr repository.AuthenticationTokenRepository) AuthenticationTokenUsecase {
	return &authenticationTokenUsecase{atr: atr}
}

func (atu *authenticationTokenUsecase) GetToken(apiKey, companyURL string) (*repository.AuthenticationToken, error) {
	return atu.atr.Get(apiKey, companyURL)
}
