package warehouse

// TODO: command 是否要完全扁平

// WarehouseCmd -
type WarehouseCmd struct {
	// 商品名稱
	Name string

	// 數量
	Quantity int64

	// 販售單價
	SalePrice int64

	// 單價成本
	CostPrice int64
}
