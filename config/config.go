package config

import (
	"os"
	"strconv"
)

type Config struct {
	Version     string
	ServiceName string
	HttpPort    int
}

var configuration Config

func loadConfig() {
	version := os.Getenv("VERSION")
	if version == "" {
		println("Version is not defiend")
		os.Exit(1)
	}
	serviceName := os.Getenv("SERVICE_NAME")
	if serviceName == "" {
		println("Version is not defiend")
		os.Exit(1)
	}
	httpPort := os.Getenv("HTTP_PORT")
	if httpPort == "" {
		println("Version is not defiend")
		os.Exit(1)
	}

	port, err := strconv.Atoi(httpPort)
	if err != nil {
		println(err)
		os.Exit(1)
	}

	configuration = Config{
		Version:     version,
		ServiceName: serviceName,
		HttpPort:    port,
	}

}

func GetConfig() Config {
	loadConfig()
	return configuration
}
