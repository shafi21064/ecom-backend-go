package config

import (
	"os"
	"strconv"

	"github.com/joho/godotenv"
)

type Config struct {
	Version       string
	ServiceName   string
	HttpPort      int
	JwtSecrateKey string
}

var configuration *Config

func loadConfig() {

	err := godotenv.Load()
	if err != nil {
		println("Faild to load env")
		os.Exit(1)
	}

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
	jwtSecrateKey := os.Getenv("JWT_SECRATE_KEY")
	if jwtSecrateKey == "" {
		println("JWT secret is not defiend")
		os.Exit(1)
	}

	port, err := strconv.Atoi(httpPort)
	if err != nil {
		println(err)
		os.Exit(1)
	}

	configuration = &Config{
		Version:       version,
		ServiceName:   serviceName,
		HttpPort:      port,
		JwtSecrateKey: jwtSecrateKey,
	}

}

func GetConfig() *Config {
	if configuration == nil{
	loadConfig()
	}
	return configuration
}
