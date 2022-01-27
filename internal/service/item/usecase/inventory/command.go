package inventory

// PurchaseCmd -
type PurchaseCmd struct {
	// 商品進貨列表
	List []*PurchaseItem
}

// PurchaseItem -
type PurchaseItem struct {
	// 商品編號
	ItemID string

	// 數量
	Quantity int64
}
