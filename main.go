package main

import (
	"os"

	"github.com/joho/godotenv"
	"github.com/shafi21064/ecom-go/cmd"
	"github.com/shafi21064/ecom-go/config"
)

func main() {
	err := godotenv.Load()
	if err != nil {
		println("Faild to load env")
		os.Exit(1)
	}

	config := config.GetConfig()

	println(config.Version)
	println(config.ServiceName)
	println(config.HttpPort)
	cmd.Serve()
}
