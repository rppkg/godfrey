package godfrey

import (
	"context"
	"errors"
	"log/slog"
	"net/http"
	"os"
	"os/signal"
	"syscall"
	"time"

	"github.com/gin-gonic/gin"
	"github.com/rppkg/godfrey/internal/pkg/middleware"
	"github.com/rppkg/godfrey/pkg/log"
	"github.com/spf13/cobra"
	"github.com/spf13/viper"
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
			if viper.GetString("GF_SERVE_GIN_MODE") == "prod" {
				gin.SetMode(gin.ReleaseMode)
			}

			if err := initDal(); err != nil {
				return err
			}

			g := gin.New()
			g.Use(gin.Recovery(), middleware.SlogInPrint())

			if err := initRouters(g); err != nil {
				return err
			}

			sv := &http.Server{Addr: viper.GetString("GF_SERVE_GIN_ADDR"), Handler: g}

			go func() {
				if err := sv.ListenAndServe(); err != nil {
					if !errors.Is(err, http.ErrServerClosed) {
						log.Fatal("Server listen", slog.Any("error", err))
					}
				}
			}()

			sig := make(chan os.Signal, 1)
			signal.Notify(sig, os.Interrupt, syscall.SIGINT, syscall.SIGTERM)
			<-sig

			log.Info("Shutting down server ...")

			ctx, cancel := context.WithTimeout(context.Background(), 10*time.Second)
			defer cancel()

			if err := sv.Shutdown(ctx); err != nil {
				log.Error("Server forced to shutdown", slog.Any("error", err))
				return err
			}

			return nil
		},
	}

	cobra.OnInitialize(initConfig)
	cmd.Flags().BoolP("toggle", "t", false, "Help message for toggle")

	return cmd
}
