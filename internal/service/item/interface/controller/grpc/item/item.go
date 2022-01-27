package item

import (
	"context"
	"woah/internal/common/command"
	pb "woah/internal/common/protobuf/item"
	"woah/internal/service/item/usecase/inventory"
	"woah/internal/service/item/usecase/shelf"
	"woah/internal/service/item/usecase/warehouse"

	"github.com/google/uuid"
)

// Warehouse -
func (s *Server) Warehouse(ctx context.Context, in *pb.WarehouseRequest) (*pb.Empty, error) {
	aggregateID := uuid.New().String()

	// new command
	cmd := command.New(aggregateID, &warehouse.WarehouseCmd{
		Name:      in.GetName(),
		Quantity:  in.GetQuantity(),
		SalePrice: in.GetSalePrice(),
		CostPrice: in.GetCostPrice(),
	})

	// dispatch command handle
	if _, err := s.Dispatch.Handle(ctx, cmd); err != nil {
		return &pb.Empty{}, err
	}

	return &pb.Empty{}, nil
}

// Shelf -
func (s *Server) Shelf(ctx context.Context, in *pb.ShelfRequest) (*pb.Empty, error) {
	var cmd command.Command
	switch in.GetStatus() {
	case 1:
		cmd = command.New(in.GetId(), &shelf.OnSoldCmd{})
	case 2:
		cmd = command.New(in.GetId(), &shelf.OffSoldCmd{})
	}

	if _, err := s.Dispatch.Handle(ctx, cmd); err != nil {
		return &pb.Empty{}, err
	}

	return &pb.Empty{}, nil
}

// Purchase -
func (s *Server) Purchase(ctx context.Context, in *pb.PurchaseRequest) (*pb.Empty, error) {
	cmd := command.New("", &inventory.PurchaseCmd{
		List: convertPurchaseItemListToDto(in.GetList()),
	})

	if _, err := s.Dispatch.Handle(ctx, cmd); err != nil {
		return &pb.Empty{}, err
	}

	return &pb.Empty{}, nil
}
