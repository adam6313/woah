package repository

import (
	"context"
	"woah/internal/service/item/domain/model/aggregate"
)

// ItemRepo -
type ItemRepo interface {
	// Create - 建立
	Create(ctx context.Context, in *aggregate.Item) error

	// Update - 更新
	Update(ctx context.Context, in *aggregate.Item) error

	// GetByID - 透過商品編號取得
	GetByID(ctx context.Context, id string) (*aggregate.Item, error)

	// SearchByIDList - 透過商品編號列表搜尋
	SearchByIDList(ctx context.Context, idList []string) ([]*aggregate.Item, error)

	// Bulk - 批次處理
	Bulk(ctx context.Context, create, update, remove []*aggregate.Item) error
}
