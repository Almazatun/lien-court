package helper

import (
	"fmt"
	"os"

	"github.com/joho/godotenv"
)

func GetEnvVar(key string) string {
	err := godotenv.Load(".env")

	if err != nil {
		fmt.Println("🔴 Error loading .env variables")
	}

	return os.Getenv(key)
}
