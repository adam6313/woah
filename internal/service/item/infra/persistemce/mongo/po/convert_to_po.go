package po

import "woah/internal/service/item/domain/model/aggregate"

// ConvertItemListToPo -
func ConvertItemListToPo(in []*aggregate.Item) []*ItemPo {
	list := make([]*ItemPo, len(in))

	for i, v := range in {
		list[i] = ConvertItemToPo(v)
	}

	return list
}

// ConvertItemToPo -
func ConvertItemToPo(in *aggregate.Item) *ItemPo {
	if in == nil {
		return &ItemPo{
			PriceInfo: &PriceInfoPo{},
		}
	}

	return &ItemPo{
		ID:       in.ID,
		Name:     in.Name,
		Quantity: in.Quantity,
		PriceInfo: &PriceInfoPo{
			SalePrice: in.PriceInfo.SalePrice,
			CostPrice: in.PriceInfo.CostPrice,
		},
		Status: uint32(in.Status),
	}
}
