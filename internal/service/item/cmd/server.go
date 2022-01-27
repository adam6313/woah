package cmd

import (
	"context"
	"fmt"
	"log"
	"net/http"
	"woah/config"
	"woah/internal/service/item/domain/service"
	"woah/internal/service/item/infra/persistemce/mongo"
	itemRepo "woah/internal/service/item/infra/persistemce/mongo/item"
	"woah/internal/service/item/usecase/inventory"
	"woah/internal/service/item/usecase/shelf"
	"woah/internal/service/item/usecase/warehouse"
	"woah/pkg/broadcast"
	"woah/pkg/logger"

	itemConf "woah/internal/service/item/infra/config"
	b "woah/pkg/broadcast"

	//httpController "woah/internal/service/item/interface/controller/http"
	grpcController "woah/internal/service/item/interface/controller/grpc"
	itemGrpc "woah/internal/service/item/interface/controller/grpc/item"

	"github.com/kataras/iris/v12"
	"github.com/spf13/cobra"
	"github.com/tyr-tech-team/hawk/srv"
	"go.uber.org/fx"
	"go.uber.org/zap"
	"google.golang.org/grpc"
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

// serve -
func serve() {
	// di
	app := fx.New(
		fx.NopLogger,
		fx.Provide(
			context.Background,
			broadcast.New,
			config.New,
			itemConf.NewRegisterClient,
			itemConf.New,
			logger.NewLogger,
			mongo.NewDial,
			itemRepo.New,
			service.New,
			warehouse.New,
			shelf.New,
			inventory.New,
			itemGrpc.New,
			grpcController.NewGrpcServer,
			//httpController.NewHTTPServer,
		),
		fx.Invoke(grpcHandle),
	)

	if err := app.Err(); err != nil {
		log.Fatal(err)
	}

	app.Run()
}

func grpcHandle(lc fx.Lifecycle, f fx.Shutdowner, ic config.IConfig, broadcast b.Broadcast, server *grpc.Server, register srv.Register, log *zap.Logger) error {

	s := srv.New(
		srv.SetName(itemConf.C.Info.Name),
		srv.SetHost(itemConf.C.Info.Host),
		srv.SetPort(itemConf.C.Info.Port),
		srv.SetRegister(register),
		srv.SetGRPC(),
		srv.SetEnableTraefik(),
	)

	lc.Append(fx.Hook{
		OnStart: func(context.Context) error {

			// start server
			go server.Serve(s.GetListener())

			// register server
			s.Register()

			log.Sugar().Info("start service on ", s.GetHost())

			return nil
		},
		OnStop: func(ctx context.Context) error {
			// deregister service
			s.Deregister()

			// close service
			s.Close()

			return nil
		},
	})

	return nil

}

func httpHandle(lc fx.Lifecycle, f fx.Shutdowner, ic config.IConfig, broadcast b.Broadcast, log *zap.Logger, h http.Handler) error {
	lc.Append(fx.Hook{
		OnStart: func(ctx context.Context) error {
			// watch config
			ic.Watch(context.Background())

			go func() {
				broadcast.Sub("conf-item", func(d []byte) {
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
