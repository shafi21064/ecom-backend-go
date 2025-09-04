package cmd

import (
	"fmt"
	"net/http"

	"github.com/shafi21064/ecom-go/middleware"
)

func Serve() {
	mngr := middleware.NewManager()
	mngr.Use(middleware.Cors, middleware.Preflight, middleware.Logger)

	mux := http.NewServeMux() // router

	initRoutes(mux, mngr)

	wrapedMux := mngr.WrapMux(mux)

	fmt.Println("Server running on 8080")

	err := http.ListenAndServe(":8080", wrapedMux)

	if err != nil {
		fmt.Println("Error on server start", err)
	}

}
