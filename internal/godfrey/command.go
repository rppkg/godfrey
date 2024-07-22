package godfrey

import (
	"net/http"
	"os"

	"github.com/gin-gonic/gin"
	"github.com/spf13/cobra"
)

func Command() *cobra.Command {
	cmd := &cobra.Command{
		Use: "serve",

		Short: "GF serve command CLI.",

		Long: `A longer description that spans multiple lines and likely contains examples
and usage of using your command. For example:

Cobra is a CLI library for Go that empowers applications.
This application is a tool to generate the needed files
to quickly create a Cobra application.`,

		RunE: func(cmd *cobra.Command, args []string) error {
			var err error

			if os.Getenv("mode") == "prod" {
				gin.SetMode(gin.ReleaseMode)
			}

			r := gin.New()
			r.RedirectTrailingSlash = false
			r.Use(gin.Recovery())

			r.GET("/ping", func(c *gin.Context) {
				c.JSON(http.StatusOK, gin.H{
					"message": "ok",
				})
			})
			if err != nil {
				return err
			}

			r.Run(":8080")

			return nil
		},
	}

	return cmd
}
