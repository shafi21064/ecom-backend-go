package rest

import (
	"net/http"
	"strconv"

	"e-com/config"
	"e-com/rest/handlers/product"
	"e-com/rest/handlers/user"
	"e-com/rest/middleware"
)

type Server struct {
	cnf            *config.Config
	userHandler    *user.Handler
	productHandler *product.Handler
}

func NewServer(
	cnf *config.Config,
	userHandler *user.Handler,
	productHanler *product.Handler,
) *Server {
	return &Server{
		cnf:            cnf,
		userHandler:    userHandler,
		productHandler: productHanler,
	}
}

func (server *Server) Start() {
	mngr := middleware.NewManager()
	mngr.Use(middleware.Cors, middleware.Preflight, middleware.Logger)

	mux := http.NewServeMux() // router

	server.userHandler.RegisterRoutes(mux, mngr)
	server.productHandler.RegisterRoutes(mux, mngr)

	wrapedMux := mngr.WrapMux(mux)

	address := ":" + strconv.Itoa(server.cnf.HttpPort)

	println("Server running on port", address)
	err := http.ListenAndServe(address, wrapedMux)

	if err != nil {
		println("Error on server start", err)
		return
	}

}
