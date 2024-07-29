package main

import (
	"os"

	_ "github.com/joho/godotenv/autoload"
	_ "go.uber.org/automaxprocs"

	"github.com/rppkg/godfrey/internal/apiserver"
)

func main() {
	app := apiserver.App()
	if err := app.Execute(); err != nil {
		os.Exit(2)
	}
}
