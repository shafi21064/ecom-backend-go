package middleware

import "e-com/config"

type Middlewares struct {
	cofig *config.Config
}

func NewMiddlewares(config *config.Config) *Middlewares {
	return &Middlewares{
		cofig: config,
	}
}
