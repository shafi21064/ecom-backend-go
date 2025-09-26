package cmd

import (
	"os"

	"e-com/config"
	"e-com/infra/db"
	"e-com/repo"
	"e-com/rest"
	productHandler "e-com/rest/handlers/product"
	userHandler "e-com/rest/handlers/user"
	"e-com/rest/middleware"
	"e-com/user"
)

func Serve() {
	cnf := config.GetConfig()

	dbCon, err := db.NewDBConnection(cnf.DBConfig)
	if err != nil {
		println(err.Error())
		os.Exit(1)
	}

	err = db.MigrateDB(dbCon, "./migrations")
	if err != nil {
		println(err.Error())
		os.Exit(1)
	}

	// Repo
	productRepo := repo.NewProductRepo(dbCon)
	userRepo := repo.NewUserRepo(dbCon)

	// Domain
	userSvc := user.NewService(userRepo)

	// Middleware
	middleware := middleware.NewMiddlewares(cnf)

	// Handler
	userHandler := userHandler.NewHandler(cnf, userSvc)
	productHandler := productHandler.NewHandler(middleware, productRepo)

	server := rest.NewServer(
		cnf,
		userHandler,
		productHandler,
	)

	server.Start()
}
