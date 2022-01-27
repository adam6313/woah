package service

import (
	"context"
	"woah/internal/service/item/domain/model/aggregate"
)

// Service -
type Service interface {
	// BeforeSave - 儲存前檢查
	BeforeSave(ctx context.Context, in *aggregate.Item) error
}
