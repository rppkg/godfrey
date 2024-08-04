package main

import (
	"os"

	"github.com/rppkg/godfrey/internal/gencode"
	// This line is necessary for go-swagger to find your docs!
	_ "github.com/rppkg/godfrey/docs/swagger"
)

//go:generate swagger generate spec -o ../../docs/swagger.yaml --scan-models

func main() {
	app := gencode.App()
	if err := app.Execute(); err != nil {
		os.Exit(2)
	}
}
