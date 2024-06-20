package repository

import (
	"encoding/json"
	"fmt"
	"github.com/ko44d/hrmos-time-aggregator/pkg/config"
	"net/http"
	"net/url"
)

type AuthenticationTokenRepository interface {
	Get() (*AuthenticationToken, error)
}

type authenticationTokenRepository struct {
	client *http.Client
	config config.Config
}

type AuthenticationToken struct {
	ExpiredAt string `json:"expired_at"`
	Token     string `json:"token"`
}

func NewAuthenticationTokenRepository(client *http.Client, config config.Config) AuthenticationTokenRepository {
	return &authenticationTokenRepository{client: client, config: config}
}

func (atr *authenticationTokenRepository) Get() (*AuthenticationToken, error) {
	u, err := url.Parse(fmt.Sprintf("%s://%s/api/%s/v1/authentication/token", "https", "ieyasu.co", atr.config.CompanyURL))
	if err != nil {
		fmt.Errorf(err.Error())
	}

	req, err := http.NewRequest(http.MethodGet, u.String(), nil)
	if err != nil {
		return nil, err
	}

	req.Header.Set("Authorization", fmt.Sprintf("Basic %s", atr.config.APIKey))
	req.Header.Set("Content-Type", "application/json")

	res, err := atr.client.Do(req)
	if err != nil {
		return nil, err
	}
	defer res.Body.Close()

	var resj *AuthenticationToken
	if err = json.NewDecoder(res.Body).Decode(&resj); err != nil {
		return nil, err
	}

	return resj, nil
}
