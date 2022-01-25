package update

import (
	"context"
	"fmt"
	"woah/internal/service/user/domain/model/aggregate"
	"woah/internal/service/user/domain/repository"
	"woah/internal/service/user/domain/service"
)

// updateUsecase  -
type updateUsecase struct {
	userRepo    repository.UserRepository
	userService service.UserService
	AggregateID string
}

// NewUseCase -
func NewUseCase(userRepo repository.UserRepository) UpdateUserUsecase {
	return &updateUsecase{}
}

// Update - 更新使用者資料
func (c *updateUsecase) Update(ctx context.Context, in *UpdateUserInfo) error {

	user := &aggregate.User{
		ID:   c.AggregateID,
		Name: in.Name,
	}

	// 檢查使用者資料
	if err := c.userService.CheckUserData(ctx, user); err != nil {
		return err
	}

	// 取得使用者資料
	origin, err := c.userRepo.GetUser(ctx, in.ID)
	if err != nil {
		return err
	}

	// 設置使用者更新資料
	if err := c.userService.SetUpdateUser(ctx, origin, user); err != nil {
		return err
	}

	// 更新資料
	if err := c.userRepo.UpdateUser(ctx, origin); err != nil {
		return err
	}

	fmt.Println("AggregateID", c.AggregateID)
	fmt.Println("Update~~", in)
	return nil
}
