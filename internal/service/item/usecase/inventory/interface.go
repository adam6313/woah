package inventory

import "context"

// InventoryUsecase - 庫存管理
type InventoryUsecase interface {
	// Purchase - 進貨
	Purchase(ctx context.Context, command *PurchaseCmd) error
}
