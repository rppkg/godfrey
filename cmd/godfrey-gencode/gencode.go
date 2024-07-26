package main

import (
	"os"

	"github.com/rppkg/godfrey/internal/gencode"
)

func main() {
	app := gencode.App()
	if err := app.Execute(); err != nil {
		os.Exit(2)
	}
}
