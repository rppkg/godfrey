package gencode

import (
	"github.com/spf13/cobra"
	"gorm.io/gen"

	"github.com/rppkg/godfrey/internal/pkg/models"
)

func App() *cobra.Command {
	cmd := &cobra.Command{
		Use: "gencode",

		Short: "The gencode is gencode service CLI for godfrey.",

		Long: `A longer description that spans multiple lines and likely contains examples
and usage of using your command. For example:
		
Cobra is a CLI library for Go that empowers applications.
This application is a tool to generate the needed files
to quickly create a Cobra application.`,

		RunE: func(cmd *cobra.Command, args []string) error {
			gormGen := gen.NewGenerator(gen.Config{
				OutPath: "internal/apiserver/dal/query",
				Mode:    gen.WithDefaultQuery,
			})

			gormGen.ApplyBasic(
				models.User{},
				models.Role{},
			)
			gormGen.Execute()
			return nil
		},
	}

	return cmd
}
