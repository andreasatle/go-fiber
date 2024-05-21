package config

import (
	"os"
	"strconv"
)

type ServerConfig struct {
	Host string
	Port string
}

type DatabaseConfig struct {
	Host string
	Port int
	User string
	Password string
	Name string
}

type TlsConfig struct {
	CertFile string
	KeyFile string
}
type Config struct {
	Server ServerConfig
	Tls TlsConfig
	Database DatabaseConfig
}

var Cfg *Config

func init() {
	serverConfig := ServerConfig {
		Host: os.Getenv("WEBSERVER_HOST"),
		Port: os.Getenv("WEBSERVER_PORT"),
	}

	db_port, err := strconv.Atoi(os.Getenv("DATABASE_PORT"))
	if err != nil {
		panic("Database port is not an integer")
	}
	databaseConfig := DatabaseConfig {
		Host: os.Getenv("DATABASE_HOST"),
		Port: db_port,
		User: os.Getenv("DATABASE_USER"),
		Password: os.Getenv("DATABASE_PASSWORD"),
		Name: os.Getenv("DATABASE_NAME"),
	}

	tlsConfig := TlsConfig {
		CertFile: os.Getenv("TLS_CERTFILE"),
		KeyFile: os.Getenv("TLS_KEYFILE"),
	}

	Cfg = &Config{
		Server: serverConfig,
		Tls: tlsConfig,
		Database: databaseConfig,
	}
}