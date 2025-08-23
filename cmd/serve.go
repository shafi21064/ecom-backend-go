package cmd

import (
	"fmt"
	"net/http"

	"github.com/shafi21064/ecom-go/global_router"
	"github.com/shafi21064/ecom-go/middleware"
)

func Serve() {
	mngr := middleware.NewManager()
	mngr.Use(middleware.Hudai, middleware.Logger)

	mux := http.NewServeMux() // router

	initRoutes(mux, mngr)

	allRouter := global_router.GlobalRouter(mux)

	fmt.Println("Server running on 8080")

	err := http.ListenAndServe(":8080", allRouter)

	if err != nil {
		fmt.Println("Error on server start", err)
	}

}
