package config

import (
	"os"

	"github.com/joho/godotenv"
)

var Stage string
var Jwt_token string

func DotEnvInit() error {
	err := godotenv.Load(".env")
	if err != nil {
		return err
	}
	return nil
}

func InitEnvVar() error {
	Stage = os.Getenv("STAGE")
	Jwt_token = os.Getenv("JWT_TOKEN")
	return nil
}
