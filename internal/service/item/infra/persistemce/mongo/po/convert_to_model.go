package po

import (
	"woah/internal/common/conv"
	"woah/internal/service/item/domain/model/aggregate"
	"woah/internal/service/item/domain/model/value_object"
)

// ConvertItemListToModel -
func ConvertItemListToModel(in []*ItemPo) []*aggregate.Item {
	list := make([]*aggregate.Item, len(in))

	for i, v := range in {
		list[i] = ConvertItemToModel(v)
	}

	return list
}

// ConvertItemToModel -
func ConvertItemToModel(in *ItemPo) *aggregate.Item {
	if in == nil {
		return &aggregate.Item{
			PriceInfo: &value_object.PriceInfo{},
		}
	}

	return &aggregate.Item{
		ID:       in.ID,
		Name:     in.Name,
		Quantity: in.Quantity,
		PriceInfo: &value_object.PriceInfo{
			SalePrice: in.PriceInfo.SalePrice,
			CostPrice: in.PriceInfo.CostPrice,
		},
		Status: conv.ItemStatusEnum(in.Status),
	}
}
