package value_object

// PriceInfo - 商品價格資訊
type PriceInfo struct {
	// 販售單價
	SalePrice int64

	// 單價成本
	CostPrice int64
}
