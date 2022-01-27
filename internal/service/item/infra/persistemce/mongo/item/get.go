package item

import (
	"context"
	"woah/internal/service/item/domain/model/aggregate"
	"woah/internal/service/item/infra/persistemce/mongo/po"

	"github.com/tyr-tech-team/hawk/status"
	"go.mongodb.org/mongo-driver/bson"
)

// GetByID -
func (r *repo) GetByID(ctx context.Context, id string) (*aggregate.Item, error) {

	coll := r.db.Database(r.databaseName).Collection(C)

	filter := bson.M{
		"id": id,
	}

	itemPo := new(po.ItemPo)
	if err := coll.FindOne(ctx, filter).Decode(&itemPo); err != nil {
		return nil, status.NotFound.Err()
	}

	return po.ConvertItemToModel(itemPo), nil
}

// SearchByIDList -
func (r *repo) SearchByIDList(ctx context.Context, idList []string) ([]*aggregate.Item, error) {

	coll := r.db.Database(r.databaseName).Collection(C)

	filter := bson.M{
		"id": bson.M{
			"$in": idList,
		},
	}

	itemPoList := make([]*po.ItemPo, 0)
	cur, err := coll.Find(ctx, filter)
	if err != nil {
		return nil, status.NotFound.Err()
	}

	if err := cur.All(ctx, &itemPoList); err != nil {
		return nil, status.NotFound.Err()
	}

	return po.ConvertItemListToModel(itemPoList), nil
}
