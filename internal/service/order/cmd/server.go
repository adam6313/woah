package cmd

import (
	"context"
	"log"
	"woah/config"
	"woah/pkg/broadcast"
	"woah/pkg/logger"

	"github.com/spf13/cobra"
	"go.uber.org/fx"
	"go.uber.org/zap"

	b "woah/pkg/broadcast"
)

// serve cmd
var serverCmd = &cobra.Command{
	Use: "server",
	Run: func(cmd *cobra.Command, args []string) {
		serve()
	},
}

func init() {
	// add serve cmd
	rootCmd.AddCommand(serverCmd)

	// set Name default is ""
	serverCmd.Flags().StringVarP(&config.Cmd.Run, "run", "r", "", "service run")
}

func serve() {
	// di
	app := fx.New(
		fx.NopLogger,
		fx.Provide(
			context.Background,
			broadcast.New,
			config.New,
			logger.NewLogger,
		),
		fx.Invoke(handle),
	)

	if err := app.Err(); err != nil {
		log.Fatal(err)
	}

	app.Run()
}

func handle(lc fx.Lifecycle, f fx.Shutdowner, ic config.IConfig, broadcast b.Broadcast, log *zap.Logger) error {
	lc.Append(fx.Hook{
		OnStart: func(ctx context.Context) error {
			// watch config
			ic.Watch(context.Background())

			// new service
			//service.New(
			//service.WithIC(ic),
			//service.WithBroadcast(broadcast),
			//service.WithLogger(log),
			//).Run()

			return nil
		},
		OnStop: func(ctx context.Context) error {
			// Shutdown fx
			f.Shutdown()

			// close conf (connect and watch)
			ic.Close()

			// close broadcast
			broadcast.CloseAll()

			return nil
		},
	})

	return nil
}
