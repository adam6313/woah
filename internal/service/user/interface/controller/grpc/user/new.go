package user

import (
	"context"
	"fmt"
	"woah/internal/common/command"
	"woah/internal/common/conv/topic"
	pb "woah/internal/common/protobuf/user"
	"woah/internal/service/user/usecase/create"
	"woah/internal/service/user/usecase/update"

	"github.com/gogo/protobuf/types"
	"github.com/golang/protobuf/proto"
)

type server struct {
	CreateUserUsecase create.CreateUserUsecase
	UpdateUserUsecase update.UpdateUserUsecase
	Dispatch          command.Dispatch
}

// NewServer -
func NewServer(createUserUsecase create.CreateUserUsecase, updateUserUsecase update.UpdateUserUsecase) pb.UserServiceServer {
	return &server{
		Dispatch: command.NewDispatch(
			createUserUsecase,
			updateUserUsecase,
		),
	}
}

// Execute -
func (s *server) Execute(ctx context.Context, in *types.Any) (*types.Any, error) {
	var (
		message     interface{}
		aggregateID string
	)

	switch in.GetTypeUrl() {
	case topic.EVENT.String():
		imple := new(pb.CreateUserRequest)
		proto.Unmarshal(in.GetValue(), imple)

		message = &create.CreateUser{
			Name: imple.GetName(),
		}
	}

	// new command
	cmd := command.New(aggregateID, message)

	result, err := s.Dispatch.Handle(ctx, cmd)
	if err != nil {
		return &types.Any{}, err
	}

	fmt.Println(result)

	return &types.Any{}, nil
}
