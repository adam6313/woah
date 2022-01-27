package item

import (
	"woah/internal/service/item/usecase/inventory"

	pb "woah/internal/common/protobuf/item"
)

func convertPurchaseItemListToDto(in []*pb.PurchaseItem) []*inventory.PurchaseItem {
	list := make([]*inventory.PurchaseItem, len(in))

	for i, v := range in {
		list[i] = &inventory.PurchaseItem{
			ItemID:   v.GetItemId(),
			Quantity: v.GetQuantity(),
		}
	}

	return list
}
