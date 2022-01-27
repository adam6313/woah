package update

import (
	"context"
	"fmt"
	"woah/internal/common/command"
)

// updateUsecase  -
type updateUsecase struct {
}

// NewUseCase -
func NewUseCase() UpdateUserUsecase {
	return &updateUsecase{}
}

// Update -
func (c *updateUsecase) Update(ctx context.Context, in *UpdateUserInfo) error {
	aggregateID := command.AggregateID(ctx)

	fmt.Println("AggregateID", aggregateID)

	fmt.Println("Update~~", in)
	return nil
}
