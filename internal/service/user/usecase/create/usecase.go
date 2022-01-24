package create

import (
	"context"
	"fmt"
)

// CreateUsecase -
type createUsecase struct {
	AggregateID string
}

// NewUseCase -
func NewUseCase() CreateUserUsecase {
	return &createUsecase{}
}

// Create -
func (c *createUsecase) Create(ctx context.Context, command *CreateUser) (*UserCreated, error) {
	//aggregateID := ctx.Value("aggregateID").(string)

	fmt.Println("AggregateID", c.AggregateID)
	return &UserCreated{ID: c.AggregateID}, nil
}
