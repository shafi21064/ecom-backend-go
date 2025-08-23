package cmd

import (
	"fmt"
	"net/http"

	"github.com/shafi21064/ecom-go/global_router"
	"github.com/shafi21064/ecom-go/handlers"
)

func Serve() {
	mux := http.NewServeMux() // router

	mux.Handle("GET /products", http.HandlerFunc(handlers.GetProducts))

	mux.Handle("POST /products", http.HandlerFunc(handlers.CreateProduct))

	mux.Handle("GET /products/{id}", http.HandlerFunc(handlers.GetProductsByID))

	allRouter := global_router.GlobalRouter(mux)

	fmt.Println("Server running on 8080")

	err := http.ListenAndServe(":8080", allRouter)

	if err != nil {
		fmt.Println("Error on server start", err)
	}

}
