package config

import (
	"os"
	"strconv"

	"github.com/joho/godotenv"
)

type DBConfig struct {
	DbUser      string
	DbPassword  string
	DbHost      string
	DbPort      int
	DbName      string
	DbEnableSSL bool
}

type Config struct {
	Version       string
	ServiceName   string
	HttpPort      int
	JwtSecrateKey string
	DBConfig      *DBConfig
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

	dBUser := os.Getenv("DB_USER")
	if dBUser == "" {
		println("DB user is not defiend")
		os.Exit(1)
	}

	dBPassword := os.Getenv("DB_PASSSWORD")
	if dBPassword == "" {
		println("DB password is not defiend")
		os.Exit(1)
	}

	dbHost := os.Getenv("DB_HOST")
	if dbHost == "" {
		println("DB host is not defiend")
		os.Exit(1)
	}

	dBPortString := os.Getenv("DB_PORT")
	if dBPortString == "" {
		println("DB port is not defiend")
		os.Exit(1)
	}

	dBName := os.Getenv("DB_NAME")
	if dBName == "" {
		println("DB name is not defiend")
		os.Exit(1)
	}

	dBEnableSSh := os.Getenv("DB_ENABLE_SSL")
	if dBEnableSSh == "" {
		dBEnableSSh = "false"
	}

	dBPort, err := strconv.Atoi(dBPortString)
	if err != nil {
		println("Can't convert the port")
		os.Exit(1)
	}
	isSslEnable, err := strconv.ParseBool(dBEnableSSh)
	if err != nil {
		println("Can't convert the port")
		os.Exit(1)
	}
	configuration = &Config{
		Version:       version,
		ServiceName:   serviceName,
		HttpPort:      port,
		JwtSecrateKey: jwtSecrateKey,
		DBConfig: &DBConfig{
			DbUser:      dBUser,
			DbPassword:  dBPassword,
			DbHost:      dbHost,
			DbPort:      dBPort,
			DbName:      dBName,
			DbEnableSSL: isSslEnable,
		},
	}

}

func GetConfig() *Config {
	if configuration == nil {
		loadConfig()
	}
	return configuration
}
