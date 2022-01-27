package service

import (
	"context"
	"woah/internal/service/item/domain/model/aggregate"

	"github.com/tyr-tech-team/hawk/status"
)

// BeforeSave -
func (s *service) BeforeSave(ctx context.Context, in *aggregate.Item) error {
	var errDescription []string

	if in.Name == "" {
		errDescription = append(errDescription, "商品名稱不得為空")
	}

	if !in.Status.Verify() {
		errDescription = append(errDescription, "商品狀態設定錯誤")
	}

	if len(errDescription) > 0 {
		return status.InvalidParameter.WithDetail(errDescription...).Err()
	}

	return nil
}
