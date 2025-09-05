package cmd

import (
	"github.com/shafi21064/ecom-go/config"
	"github.com/shafi21064/ecom-go/rest"
)

func Serve() {
	cnf := config.GetConfig()
	rest.Start(cnf)
}
