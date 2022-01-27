package cmd

import (
	"context"
	"fmt"
	"log"
	"net"
	"time"
	"woah/config"
	grpc_user "woah/internal/service/user/interface/controller/grpc"
	"woah/internal/service/user/interface/controller/grpc/user"
	"woah/internal/service/user/usecase/create"
	"woah/internal/service/user/usecase/update"
	"woah/pkg/broadcast"
	"woah/pkg/logger"

	"github.com/gogo/protobuf/types"
	"github.com/spf13/cobra"
	"go.uber.org/fx"
	"go.uber.org/zap"

	"google.golang.org/grpc"

	"woah/internal/common/conv/topic"
	pb "woah/internal/common/protobuf/user"
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

	// set consul address
	serverCmd.Flags().StringVarP(&config.Cmd.ConsulAddress, "consul", "c", "127.0.01:8500", "consul address")
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
			create.NewUseCase,
			update.NewUseCase,
			user.NewServer,
			grpc_user.NewGrpcServer,
			//controller_http.NewHTTPServer,
		),
		fx.Invoke(handle),
	)

	if err := app.Err(); err != nil {
		log.Fatal(err)
	}

	app.Run()
}

func handle(lc fx.Lifecycle, f fx.Shutdowner, ic config.IConfig, broadcast b.Broadcast, log *zap.Logger, server *grpc.Server) error {
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

			ln, _ := net.Listen("tcp", ":3333")
			go server.Serve(ln)

			log.Sugar().Info("start service on ", ln.Addr().String())

			time.Sleep(1 * time.Second)

			conn, err := grpc.Dial("localhost:3333", grpc.WithInsecure(), grpc.WithBlock())
			if err != nil {
				panic(err)
			}

			defer conn.Close()
			c := pb.NewUserServiceClient(conn)

			d := pb.CreateUserRequest{
				Name: "Adam",
			}

			dd, _ := d.Marshal()

			c.Execute(context.Background(), &types.Any{
				TypeUrl: topic.EVENT.String(),
				Value:   dd,
			})

			//go h.(*iris.Application).Run(iris.Addr(":3000"))

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
