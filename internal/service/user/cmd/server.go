package cmd

import (
	"context"
	"fmt"
	"log"
	"woah/config"
	service_config "woah/internal/service/user/infra/config"
	"woah/internal/service/user/infra/persistence/mongo"
	user_repo "woah/internal/service/user/infra/persistence/mongo/user"
	controller_http "woah/internal/service/user/interface/controller/http"
	"woah/internal/service/user/usecase/create"
	"woah/internal/service/user/usecase/update"
	"woah/pkg/broadcast"
	"woah/pkg/logger"

	"github.com/kataras/iris/v12"
	"github.com/spf13/cobra"
	"go.uber.org/fx"
	"go.uber.org/zap"

	b "woah/pkg/broadcast"

	"net/http"
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
			service_config.NewServiceConfig,
			logger.NewLogger,
			create.NewUseCase,
			update.NewUseCase,
			controller_http.NewHTTPServer,
			user_repo.NewRepository,
			mongo.NewDial,
		),
		fx.Invoke(handle),
	)

	if err := app.Err(); err != nil {
		log.Fatal(err)
	}

	app.Run()
}

func handle(lc fx.Lifecycle, f fx.Shutdowner, ic config.IConfig, broadcast b.Broadcast, log *zap.Logger, h http.Handler, c *service_config.Config) error {
	lc.Append(fx.Hook{
		OnStart: func(ctx context.Context) error {
			// watch config
			ic.Watch(context.Background())

			go func() {
				broadcast.Sub("conf-user", func(d []byte) {

					fmt.Println("123")
					fmt.Println(string(d))
				})
			}()

			go h.(*iris.Application).Run(iris.Addr(":3000"))

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
