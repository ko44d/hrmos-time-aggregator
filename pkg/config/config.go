package config

import "github.com/caarlos0/env/v11"

type Config struct {
	CompanyURL string `env:"COMPANY_URL"`
	APIKey     string `env:"API_KEY"`
}

func New() (Config, error) {

	var c Config
	if err := env.Parse(&c); err != nil {
		return c, err
	}
	return c, nil
}
