package update

import "context"

// UpdateUserUsecase -
type UpdateUserUsecase interface {
	// Update - 更新使用者資料
	Update(ctx context.Context, in *UpdateUserInfo) error
}
