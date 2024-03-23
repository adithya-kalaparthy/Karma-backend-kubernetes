package config

import (
	"fmt"
	"path"
	"runtime"

	"github.com/joho/godotenv"
)

func LoadEnv() error {
	_, filename, _, _ := runtime.Caller(0)
	envFilePath := path.Join(path.Dir(filename), "../../.env")

	err := godotenv.Load(envFilePath)

	if err != nil {
		return fmt.Errorf("load env data is not working properly %v", err)
	}

	return nil
}
