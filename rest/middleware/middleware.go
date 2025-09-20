package middleware

import "github.com/shafi21064/ecom-go/config"

type Middlewares struct {
	cofig *config.Config
}

func NewMiddlewares(config *config.Config) *Middlewares {
	return &Middlewares{
		cofig: config,
	}
}
