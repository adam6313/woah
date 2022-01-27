package aggregate

import (
	"woah/internal/common/conv"
	"woah/internal/service/item/domain/model/value_object"
)

// Item - 商品
type Item struct {
	// 編號
	ID string

	// 名稱
	Name string

	// 庫存量
	Quantity int64

	// 價格資訊
	PriceInfo *value_object.PriceInfo

	// 狀態
	Status conv.ItemStatusEnum
}
