package shelf

import "context"

// ShelfUsecase - 上下架管理
type ShelfUsecase interface {
	// 上架
	OnSold(ctx context.Context, command *OnSoldCmd) error

	// 下架
	OffSold(ctx context.Context, command *OffSoldCmd) error
}
