package cmd

import (
	"context"
	"log"

	"woah/config"
	"woah/pkg/broadcast"
	"woah/pkg/broker"
	"woah/service"

	"github.com/davecgh/go-spew/spew"
	"github.com/spf13/cobra"
	mb "go-micro.dev/v4/broker"
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
			broker.NewMemoryBroker,
			broadcast.NewTest,
		),
		fx.Invoke(handle),
	)

	if err := app.Err(); err != nil {
		log.Fatal(err)
	}

	app.Run()
}

func handle(lc fx.Lifecycle, f fx.Shutdowner, ic config.IConfig, broadCast broadcast.BroadCast) error {
	lc.Append(fx.Hook{
		OnStart: func(context.Context) error {
			broadCast.Subscribe("test", func(e mb.Event) error {
				spew.Dump(e.Message().Header["target"])
				return nil
			})

			broadCast.OnSubscribe("test", &config.Values{}, func(header map[string]string, in interface{}) error {
				b, ok := in.(*config.Values)
				spew.Dump(b, ok)
				return nil
			})

			go ic.Watch(context.Background())

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

			broadCast.CloseAll()

			return nil
		},
	})

	return nil
}
