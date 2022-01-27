package item

import (
	"context"
	"woah/internal/service/item/domain/model/aggregate"
	"woah/internal/service/item/infra/persistemce/mongo/po"

	"go.mongodb.org/mongo-driver/bson"
	"go.mongodb.org/mongo-driver/mongo"
	"go.mongodb.org/mongo-driver/mongo/options"
)

func (r *repo) Bulk(ctx context.Context, create, update, remove []*aggregate.Item) error {
	var (
		index      = 0
		operations = make([]mongo.WriteModel, len(create)+len(update)+len(remove))
	)

	createListPo := po.ConvertItemListToPo(create)

	coll := r.db.Database(r.databaseName).Collection(C)

	// add create data
	if len(create) > 0 {
		for _, v := range createListPo {
			operation := mongo.NewInsertOneModel()
			operation.SetDocument(v)

			operations[index] = operation
			index++
		}
	}

	// add update data
	if len(update) > 0 {
		for _, v := range update {
			operation := mongo.NewUpdateOneModel()

			filter := bson.M{"id": v.ID}
			operation.SetFilter(filter)

			update := bson.M{"$set": setUpdate(v)}
			operation.SetUpdate(update)

			operations[index] = operation
			index++
		}
	}

	// add remove date
	if len(remove) > 0 {
		for _, v := range remove {
			operation := mongo.NewDeleteOneModel()

			filter := bson.M{"id": v.ID}
			operation.SetFilter(filter)

			operations[index] = operation
			index++
		}
	}

	opts := options.BulkWriteOptions{}
	opts.SetOrdered(false)

	_, err := coll.BulkWrite(ctx, operations, &opts)
	if err != nil {
		return err
	}

	return nil
}
