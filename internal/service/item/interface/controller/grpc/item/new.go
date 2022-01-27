package item

import (
	"woah/internal/common/command"
	pb "woah/internal/common/protobuf/item"
	"woah/internal/service/item/usecase/inventory"
	"woah/internal/service/item/usecase/shelf"
	"woah/internal/service/item/usecase/warehouse"
)

// Server -
type Server struct {
	//WarehouseUsecase warehouse.WarehouseUsecase
	//ShelfUsecase     shelf.ShelfUsecase
	Dispatch command.Dispatch
}

// New -
func New(warehouseUsecase warehouse.WarehouseUsecase, shelfUsecase shelf.ShelfUsecase, inventoryUsecase inventory.InventoryUsecase) pb.ItemServiceServer {
	return &Server{
		Dispatch: command.NewDispatch(
			warehouseUsecase,
			shelfUsecase,
			inventoryUsecase,
		),
	}
}
