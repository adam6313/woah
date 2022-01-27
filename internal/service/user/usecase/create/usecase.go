package create

import (
	"context"
	"fmt"

	"github.com/google/uuid"
)

// CreateUsecase -
type createUsecase struct {
}

// NewUseCase -
func NewUseCase() CreateUserUsecase {
	return &createUsecase{}
}

// Create -
func (c *createUsecase) Create(ctx context.Context, cmd *CreateUser) (*UserCreated, error) {
	fmt.Println("Create~~~~~~~")
	aggregateID := uuid.New().String()

	fmt.Println(aggregateID)

	return &UserCreated{ID: aggregateID}, nil
}
