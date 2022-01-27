package grpc

import (
	pb "woah/internal/common/protobuf/item"

	grpc_middleware "github.com/grpc-ecosystem/go-grpc-middleware"
	grpc_auth "github.com/grpc-ecosystem/go-grpc-middleware/auth"
	grpc_zap "github.com/grpc-ecosystem/go-grpc-middleware/logging/zap"
	grpc_recovery "github.com/grpc-ecosystem/go-grpc-middleware/recovery"
	grpc_ctxtags "github.com/grpc-ecosystem/go-grpc-middleware/tags"

	hawk_grpc "github.com/tyr-tech-team/hawk/middleware/grpc"
	"go.uber.org/zap"
	"google.golang.org/grpc"
)

//NewGrpcServer -
func NewGrpcServer(log *zap.Logger, item pb.ItemServiceServer) *grpc.Server {
	unaryInterceptor := []grpc.UnaryServerInterceptor{
		grpc_ctxtags.UnaryServerInterceptor(
			grpc_ctxtags.WithFieldExtractor(grpc_ctxtags.CodeGenRequestFieldExtractor),
		),
		grpc_zap.UnaryServerInterceptor(log),
		grpc_auth.UnaryServerInterceptor(hawk_grpc.TraceID),
		grpc_auth.UnaryServerInterceptor(hawk_grpc.GetOperator),
		grpc_recovery.UnaryServerInterceptor(),
	}

	streamInterceptor := []grpc.StreamServerInterceptor{
		grpc_ctxtags.StreamServerInterceptor(
			grpc_ctxtags.WithFieldExtractor(grpc_ctxtags.CodeGenRequestFieldExtractor),
		),
		grpc_zap.StreamServerInterceptor(log),
		grpc_auth.StreamServerInterceptor(hawk_grpc.TraceID),
		grpc_auth.StreamServerInterceptor(hawk_grpc.GetOperator),
		grpc_recovery.StreamServerInterceptor(),
	}

	// tracing
	//if c.Trace.Environment != "" && c.Trace.URL != "" {
	//unaryInterceptor = append(unaryInterceptor, otelgrpc.UnaryServerInterceptor())
	//streamInterceptor = append(streamInterceptor, otelgrpc.StreamServerInterceptor())
	//}

	server := grpc.NewServer(
		grpc.UnaryInterceptor(
			grpc_middleware.ChainUnaryServer(unaryInterceptor...),
		),
		grpc.StreamInterceptor(
			grpc_middleware.ChainStreamServer(streamInterceptor...),
		),
	)

	pb.RegisterItemServiceServer(server, item)

	return server
}
