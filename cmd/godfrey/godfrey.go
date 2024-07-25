package main

import (
	"os"

	_ "github.com/joho/godotenv/autoload"
	_ "go.uber.org/automaxprocs"

	"github.com/rppkg/godfrey/internal/godfrey"
)

func main() {
	if err := godfrey.App().Execute(); err != nil {
		os.Exit(2)
	}
}
