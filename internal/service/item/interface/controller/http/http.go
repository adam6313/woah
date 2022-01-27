package http

import (
	"net/http"
	"woah/internal/common/command"
	"woah/internal/service/item/interface/controller/http/item"
	"woah/internal/service/item/usecase/inventory"
	"woah/internal/service/item/usecase/shelf"
	"woah/internal/service/item/usecase/warehouse"

	"github.com/kataras/iris/v12"
)

// NewHTTPServer -
func NewHTTPServer(warehouseUsecase warehouse.WarehouseUsecase, shelfUsecase shelf.ShelfUsecase, inventoryUsecase inventory.InventoryUsecase) http.Handler {
	h := item.Server{
		App: iris.New(),
		Dispatch: command.NewDispatch(
			warehouseUsecase,
			shelfUsecase,
			inventoryUsecase,
		),
	}

	h.SetRouter()

	return h.App
}
