package rest

import (
	"net/http"
	"strconv"

	"github.com/shafi21064/ecom-go/config"
	"github.com/shafi21064/ecom-go/rest/handlers/product"
	"github.com/shafi21064/ecom-go/rest/handlers/user"
	"github.com/shafi21064/ecom-go/rest/middleware"
)

type Server struct {
	userHandler    *user.Handler
	productHandler *product.Handler
}

func NewServer(
	userHandler *user.Handler,
	productHanler *product.Handler,
) *Server {
	return &Server{
		userHandler:    userHandler,
		productHandler: productHanler,
	}
}

func (server *Server) Start(cnf config.Config) {
	mngr := middleware.NewManager()
	mngr.Use(middleware.Cors, middleware.Preflight, middleware.Logger)

	mux := http.NewServeMux() // router

	server.userHandler.RegisterRoutes(mux, mngr)
	server.productHandler.RegisterRoutes(mux, mngr)
	

	wrapedMux := mngr.WrapMux(mux)

	address := ":" + strconv.Itoa(cnf.HttpPort)

	println("Server running on port", address)
	err := http.ListenAndServe(address, wrapedMux)

	if err != nil {
		println("Error on server start", err)
		return
	}

}
