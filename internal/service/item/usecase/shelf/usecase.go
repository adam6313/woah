package shelf

import (
	"context"
	"woah/internal/common/conv"
	"woah/internal/service/item/domain/repository"
)

// shelfUsecase -
type shelfUsecase struct {
	// 聚合編號
	AggregateID string
	itemRepo    repository.ItemRepo
}

// New -
func New(itemRepo repository.ItemRepo) ShelfUsecase {
	return &shelfUsecase{
		itemRepo: itemRepo,
	}
}

// OnSold -
func (u *shelfUsecase) OnSold(ctx context.Context, command *OnSoldCmd) error {
	// 取得原資料
	originalItem, err := u.itemRepo.GetByID(ctx, u.AggregateID)
	if err != nil {
		return err
	}

	// 一樣不處理
	if originalItem.Status == conv.ItemStatusOnSold {
		return nil
	}

	// 上架邏輯
	originalItem.Status = conv.ItemStatusOnSold

	// 更新狀態
	if err := u.itemRepo.Update(ctx, originalItem); err != nil {
		return err
	}

	return nil
}

// OffSold -
func (u *shelfUsecase) OffSold(ctx context.Context, command *OffSoldCmd) error {
	// 取得原資料
	originalItem, err := u.itemRepo.GetByID(ctx, u.AggregateID)
	if err != nil {
		return err
	}

	// 一樣不處理
	if originalItem.Status == conv.ItemStatusOffSold {
		return nil
	}

	// 下架邏輯
	originalItem.Status = conv.ItemStatusOffSold

	// 更新狀態
	if err := u.itemRepo.Update(ctx, originalItem); err != nil {
		return err
	}

	return nil
}
