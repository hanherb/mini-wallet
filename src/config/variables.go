package config

import (
	"os"
)

// Mysql environment
var Mysql struct {
	Host     string
	User     string
	Pass     string
	Database string
}

// Token key for encryption
var JWTSecretKey string

func VariablesInitialization() {
	Mysql.Host = os.Getenv("MYSQL_HOST")
	Mysql.User = os.Getenv("MYSQL_USER")
	Mysql.Pass = os.Getenv("MYSQL_PASSWORD")
	Mysql.Database = os.Getenv("MYSQL_DATABASE")

	JWTSecretKey = os.Getenv("JWT_SECRET_KEY")
}
