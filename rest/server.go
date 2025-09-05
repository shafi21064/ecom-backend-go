package rest

import (
	"net/http"
	"strconv"

	"github.com/shafi21064/ecom-go/config"
	"github.com/shafi21064/ecom-go/rest/middleware"
)

func Start(cnf config.Config) {
	mngr := middleware.NewManager()
	mngr.Use(middleware.Cors, middleware.Preflight, middleware.Logger)

	mux := http.NewServeMux() // router

	initRoutes(mux, mngr)

	wrapedMux := mngr.WrapMux(mux)

	address := ":" + strconv.Itoa(cnf.HttpPort)

	println("Server running on port", address)
	err := http.ListenAndServe(address, wrapedMux)

	if err != nil {
		println("Error on server start", err)
		return
	}

}
