package inventory

import (
	"context"
	"woah/internal/service/item/domain/repository"

	"github.com/tyr-tech-team/hawk/status"
)

// inventoryUsecase -
type inventoryUsecase struct {
	itemRepo repository.ItemRepo
}

// New -
func New(itemRepo repository.ItemRepo) InventoryUsecase {
	return &inventoryUsecase{
		itemRepo: itemRepo,
	}
}

// Purchase -
func (u *inventoryUsecase) Purchase(ctx context.Context, command *PurchaseCmd) error {
	var (
		itemQuantityMap = make(map[string]int64)
		itemIDList      = make([]string, len(command.List))
	)

	for i, v := range command.List {
		itemQuantityMap[v.ItemID] = v.Quantity
		itemIDList[i] = v.ItemID
	}

	// search by item idList
	itemList, err := u.itemRepo.SearchByIDList(ctx, itemIDList)
	if err != nil {
		return err
	}

	// 處理數量
	for _, v := range itemList {
		quantity := itemQuantityMap[v.ID]
		v.Quantity += quantity
	}

	// 批次處理
	if err := u.itemRepo.Bulk(ctx, nil, itemList, nil); err != nil {
		return status.UpdatedFailed.Err()
	}

	return nil
}
