package item

import (
	"context"
	"time"
	"woah/internal/service/item/domain/model/aggregate"
	"woah/internal/service/item/infra/persistemce/mongo/po"

	"github.com/tyr-tech-team/hawk/status"
)

// Create -
func (r *repo) Create(ctx context.Context, in *aggregate.Item) error {
	coll := r.db.Database(r.databaseName).Collection(C)

	po := po.ConvertItemToPo(in)
	po.CreatedAt = time.Now().In(time.Local)

	if _, err := coll.InsertOne(ctx, po); err != nil {
		return status.CreatedFailed.Err()
	}

	return nil
}
