package infra

import (
	"context"
	"fmt"
	"testing"

	"github.com/stretchr/testify/assert"
	"go.mongodb.org/mongo-driver/bson"
)

func TestMongoDriver(t *testing.T) {
	mongo := MongoDriver{}
	ctx := context.Background()
	mongo.Config(ctx, "mongodb://root:example@127.0.0.1:27018")
	t.Run("Test MongoDriver Connect", func(t *testing.T) {
		err := mongo.Connect(ctx)
		assert.NoError(t, err)

	})

	t.Run("Test MongoDriver FInd", func(t *testing.T) {
		results, err := mongo.Find(ctx, "products", bson.M{"name": "test"})
		fmt.Println(results)
		assert.NoError(t, err)
	})

	t.Run("Test MongoDriver Insert", func(t *testing.T) {
		result, err := mongo.Insert(ctx, "products", Test{Name: "test"})
		assert.NoError(t, err)
		assert.NotEmpty(t, result)
	})

	t.Run("Test MongoDriver Delete", func(t *testing.T) {
		count, err := mongo.Delete(ctx, "products", bson.M{})
		assert.NoError(t, err)
		assert.GreaterOrEqual(t, count, int64(1))
	})
}
