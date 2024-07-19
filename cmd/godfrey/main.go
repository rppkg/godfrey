package main

import (
	"os"

	_ "github.com/joho/godotenv/autoload"
	"github.com/spf13/cobra"
	_ "go.uber.org/automaxprocs"

	"github.com/rppkg/godfrey/internal/godfrey"
)

var cmd = &cobra.Command{
	Use:   "gf",
	Short: "The gf is root CLI for godfrey.",
}

func main() {
	cmd.AddCommand(godfrey.Command())
	if err := cmd.Execute(); err != nil {
		os.Exit(2)
	}
}
