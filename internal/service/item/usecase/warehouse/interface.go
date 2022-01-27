package warehouse

import "context"

// WarehouseUsecase - 入庫管理
type WarehouseUsecase interface {
	// 入庫
	Warehouse(ctx context.Context, command *WarehouseCmd) error
}
