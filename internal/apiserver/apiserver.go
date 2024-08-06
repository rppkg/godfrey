package apiserver

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
	"github.com/spf13/cobra"
	"github.com/spf13/viper"

	"github.com/rppkg/godfrey/internal/pkg/middleware"
	"github.com/rppkg/godfrey/pkg/log"
)

func App() *cobra.Command {
	cmd := &cobra.Command{
		Use: "apiserver",

		Short: "The apiserver is apiserver service CLI for godfrey.",

		Long: `A longer description that spans multiple lines and likely contains examples
and usage of using your command. For example:

Cobra is a CLI library for Go that empowers applications.
This application is a tool to generate the needed files
to quickly create a Cobra application.`,

		RunE: func(cmd *cobra.Command, args []string) error {
			if viper.GetString("APISERVER_GIN_MODE") == "prod" {
				gin.SetMode(gin.ReleaseMode)
			}

			if err := initToken(); err != nil {
				log.Error("Init token", slog.Any("error", err))
				return err
			}

			if err := initDal(); err != nil {
				log.Error("Init dal", slog.Any("error", err))
				return err
			}

			g := gin.New()
			g.Use(gin.Recovery(), middleware.SlogInPrint())
			if err := initRouters(g); err != nil {
				log.Error("Init routers", slog.Any("error", err))
				return err
			}

			sv := &http.Server{Addr: viper.GetString("APISERVER_GIN_ADDR"), Handler: g}
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
	cmd.PersistentFlags().StringVarP(&cfg, "config", "c", "",
		"The path to the godfrey configuration file. Empty string for no configuration file.")

	return cmd
}
