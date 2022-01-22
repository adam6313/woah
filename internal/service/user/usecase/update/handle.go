package update

import (
	"context"
	"fmt"
)

// updateUsecase  -
type updateUsecase struct {
	//AggregateID string
}

// NewUseCase -
func NewUseCase() UpdateUserUsecase {
	return &updateUsecase{}
}

// Update -
func (c *updateUsecase) Update(ctx context.Context, in *UpdateUserInfo) error {
	fmt.Println("AggregateID", c.AggregateID)
	fmt.Println("Update~~", in)
	return nil
}
