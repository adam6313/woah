package po

import "time"

// ItemPo - 商品
type ItemPo struct {
	// 編號
	ID string `bson:"id"`

	// 名稱
	Name string `bson:"name"`

	// 庫存量
	Quantity int64 `bson:"quantity"`

	// 價格資訊
	PriceInfo *PriceInfoPo `bson:"priceInfo"`

	// 狀態
	Status uint32 `bson:"status"`

	// ----- Po -----
	// 建立時間
	CreatedAt time.Time `bson:"createdAt"`
}

// PriceInfoPo - 商品價格資訊
type PriceInfoPo struct {
	// 販售單價
	SalePrice int64 `bson:"salePrice"`

	// 單價成本
	CostPrice int64 `bson:"costPrice"`
}
