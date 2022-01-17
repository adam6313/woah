package cmd

import (
	"context"
	"log"

	"woah/config"
	"woah/service"

	"github.com/spf13/cobra"
	"go.uber.org/fx"
)

// serve cmd
var serveCmd = &cobra.Command{
	Use:   "srv",
	Short: "s",
	Run: func(cmd *cobra.Command, args []string) {
		serve()
	},
}

func init() {
	// add serve cmd
	rootCmd.AddCommand(serveCmd)

	// set Name default is ""
	serveCmd.Flags().StringVarP(&config.Cmd.Run, "run", "r", "", "service run")
}

func serve() {
	// di
	app := fx.New(
		fx.NopLogger,
		fx.Provide(
			context.Background,
			config.New,
		),
		fx.Invoke(handle),
	)

	if err := app.Err(); err != nil {
		log.Fatal(err)
	}

	app.Run()
}

func handle(lc fx.Lifecycle, f fx.Shutdowner, ic config.IConfig) error {
	lc.Append(fx.Hook{
		OnStart: func(context.Context) error {

			// new service
			service.New(
				service.WithIC(ic),
			).Run()

			return nil
		},
		OnStop: func(ctx context.Context) error {
			// Shutdown fx
			f.Shutdown()

			//close conf (connect and watch)
			ic.Close()

			return nil
		},
	})

	return nil
}
