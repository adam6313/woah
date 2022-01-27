package item

import (
	"woah/internal/common/command"
	"woah/internal/service/item/usecase/inventory"
	"woah/internal/service/item/usecase/shelf"
	"woah/internal/service/item/usecase/warehouse"
	"woah/internal/service/user/interface/controller/http/middle"

	"github.com/google/uuid"
	"github.com/kataras/iris/v12"
)

// Server -
type Server struct {
	App              *iris.Application
	WarehouseUsecase warehouse.WarehouseUsecase
	ShelfUsecase     shelf.ShelfUsecase
	Dispatch         command.Dispatch
}

// Warehouse -
func (s *Server) Warehouse(c *middle.C) {
	aggregateID := uuid.New().String()

	data := new(warehouse.WarehouseCmd)
	if err := c.ReadJSON(data); err != nil {
		c.E(err)
		return
	}

	// new command
	cmd := command.New(aggregateID, data)

	// dispatch command handle
	if _, err := s.Dispatch.Handle(c.Request().Context(), cmd); err != nil {
		c.E(err)
		return
	}

	c.R(nil)
}

// Shelf -
func (s *Server) Shelf(c *middle.C) {
	aggregateID := c.Params().Get("id")
	itemStatus := c.Params().Get("status")

	// new command
	var cmd command.Command

	switch itemStatus {
	case "1":
		cmd = command.New(aggregateID, &shelf.OnSoldCmd{})
	case "2":
		cmd = command.New(aggregateID, &shelf.OffSoldCmd{})
	}

	// dispatch command handle
	if _, err := s.Dispatch.Handle(c.Request().Context(), cmd); err != nil {
		c.E(err)
		return
	}

	c.R(nil)
}

// Purchase -
func (s *Server) Purchase(c *middle.C) {
	data := new(inventory.PurchaseCmd)
	if err := c.ReadJSON(data); err != nil {
		c.E(err)
		return
	}

	// new command
	cmd := command.New("", data)

	// dispatch command handle
	if _, err := s.Dispatch.Handle(c.Request().Context(), cmd); err != nil {
		c.E(err)
		return
	}

	c.R(nil)
}
