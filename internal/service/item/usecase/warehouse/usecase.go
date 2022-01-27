package warehouse

import (
	"context"
	"woah/internal/common/conv"
	"woah/internal/service/item/domain/model/aggregate"
	"woah/internal/service/item/domain/model/value_object"
	"woah/internal/service/item/domain/repository"
	"woah/internal/service/item/domain/service"
)

// warehouseUsecase -
type warehouseUsecase struct {
	// 聚合編號
	AggregateID string
	itemRepo    repository.ItemRepo
	service     service.Service
}

// New -
func New(itemRepo repository.ItemRepo, service service.Service) WarehouseUsecase {
	return &warehouseUsecase{
		itemRepo: itemRepo,
		service:  service,
	}
}

// Warehouse -
func (u *warehouseUsecase) Warehouse(ctx context.Context, command *WarehouseCmd) error {
	// 轉換成 model
	in := u.convertItemToModel(command)

	// 儲存前檢查
	if err := u.service.BeforeSave(ctx, in); err != nil {
		return err
	}

	// 建立
	if err := u.itemRepo.Create(ctx, in); err != nil {
		return err
	}

	return nil
}

// convertItemToModel -
func (u *warehouseUsecase) convertItemToModel(command *WarehouseCmd) *aggregate.Item {
	return &aggregate.Item{
		ID:       u.AggregateID,
		Name:     command.Name,
		Quantity: command.Quantity,
		PriceInfo: &value_object.PriceInfo{
			SalePrice: command.SalePrice,
			CostPrice: command.CostPrice,
		},
		Status: conv.ItemStatusOffSold, // 預設下架
	}
}
