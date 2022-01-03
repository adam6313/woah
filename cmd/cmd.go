package cmd

import (
	"context"
	"fmt"
	"log"

	"woah/config"

	"github.com/spf13/cobra"
	"go.uber.org/fx"
)

// serve cmd
var serveCmd = &cobra.Command{
	Use:   "serve",
	Short: "s",
	Run: func(cmd *cobra.Command, args []string) {
		serve()
	},
}

func init() {
	// add serve cmd
	rootCmd.AddCommand(serveCmd)
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

			fmt.Println("@@", string(ic.Mod()))
			fmt.Println("@@", string(ic.Log()))
			fmt.Println("@@", string(ic.Log()))

			fmt.Println(string(ic.Service("user")))

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
