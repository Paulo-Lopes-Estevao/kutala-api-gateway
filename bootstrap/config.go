package bootstrap

import (
	"github.com/joho/godotenv"
	"log"
	"os"
	"path/filepath"
	"runtime"
)

func GoDotEnvVariable(key string) string {

	_, b, _, _ := runtime.Caller(0)
	basepath := filepath.Dir(b)

	err := godotenv.Load(basepath + "/../.env")

	if err != nil {
		log.Fatal("Error loading .env file", err)
	}

	return os.Getenv(key)

}
