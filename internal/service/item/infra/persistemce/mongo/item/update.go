package item

import (
	"context"
	"woah/internal/service/item/domain/model/aggregate"

	"github.com/tyr-tech-team/hawk/status"
	"go.mongodb.org/mongo-driver/bson"
)

// Update-
func (r *repo) Update(ctx context.Context, in *aggregate.Item) error {
	coll := r.db.Database(r.databaseName).Collection(C)

	filter := bson.M{
		"id": in.ID,
	}

	update := bson.M{
		"$set": setUpdate(in),
	}

	if _, err := coll.UpdateOne(ctx, filter, update); err != nil {
		return status.UpdatedFailed.Err()
	}

	return nil
}

func setUpdate(in *aggregate.Item) bson.M {
	update := bson.M{}

	update["id"] = in.ID
	update["name"] = in.Name
	update["quantity"] = in.Quantity
	update["status"] = in.Status

	return update
}
