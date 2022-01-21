package create

import (
	"context"
	"fmt"
)

// CreateUsecase -
type createUsecase struct {
}

// NewUseCase -
func NewUseCase() CreateUserUsecase {
	return &createUsecase{}
}

// Create -
func (c *createUsecase) Create(ctx context.Context, in *CreateUser) error {
	fmt.Println("create~~")
	return nil
}

//Handle -
//func (c createUsecase) Handle(ctx context.Context, message command.Command) error {
//switch cmd := message.Command().(type) {
//case CreateUser:
//c.Create(ctx, aggregate.User{
//Name: cmd.Name,
//})
//}

//return nil
//}
