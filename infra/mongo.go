package infra

import (
	"context"

	config "github.com/GinanTrade/go-common-libs/config"

	"go.mongodb.org/mongo-driver/bson"
	"go.mongodb.org/mongo-driver/bson/primitive"
	"go.mongodb.org/mongo-driver/mongo"
	"go.mongodb.org/mongo-driver/mongo/options"
)

type Test struct {
	ID   primitive.ObjectID `bson:"_id,omitempty"`
	Name string             `bson:"name"`
}

// IMongoDriver is the interface for the abstraction of mongo driver
type IMongoDriver interface {
	Insert(ctx context.Context, collection string, data interface{}) (interface{}, error)
	Find(ctx context.Context, collection string, filter interface{}) ([]interface{}, error)
	Delete(ctx context.Context, collection string, filter interface{}) (int64, error)
	// FindOne(ctx context.Context, string, filter interface{}) (interface{}, error)
	// Update(ctx context.Context, string, filter interface{}, data T) error
	// Deletefctx context.Context, string, filter interface{}) error
	Config(ctx context.Context, connectionString string) error
	Connect(ctx context.Context) error

	GetClient() *mongo.Client
	Disconnect() error
}

type MongoDriver struct {
	uri    string
	client *mongo.Client
}

func (m *MongoDriver) Config(ctx context.Context, connectionString string) error {
	m.uri = connectionString
	return nil
}
func (m *MongoDriver) Connect(ctx context.Context) error {
	serverApi := options.ServerAPI(options.ServerAPIVersion1)
	opts := options.Client().ApplyURI(m.uri).SetServerAPIOptions(serverApi)
	client, err := mongo.Connect(ctx, opts)
	if err != nil {
		panic(err)
	}

	// defer func() {
	// 	if err = client.Disconnect(ctx); err != nil {
	// 		panic(err)
	// 	}
	// }()
	//var result bson.M
	m.client = client
	return nil
}

func (m *MongoDriver) Disconnect() error {
	return m.client.Disconnect(context.Background())
}

func (m *MongoDriver) Find(ctx context.Context, collection string, filter interface{}) ([]interface{}, error) {
	// colection := m.client.Database("productservice").Collection(collection)
	cur, err := m.client.Database("productservice").Collection(collection).Find(ctx, filter)
	results := []interface{}{}
	if err != nil {
		return nil, err
	}

	for cur.Next(ctx) {
		var result bson.M

		err := cur.Decode(&result)
		if err != nil {
			return nil, err
		}
		results = append(results, result)

	}
	return results, nil

	// if err != nil {
	// 	panic(err)
	// }
	// fmt.Println(result)

	// return nil, err
}

// Generate for insert
func (m *MongoDriver) Insert(ctx context.Context, collection string, data interface{}) (interface{}, error) {
	col := m.client.Database("productservice").Collection(collection)

	result, err := col.InsertOne(ctx, data)
	if err != nil {
		return nil, err
	}
	id := result.InsertedID.(primitive.ObjectID)

	return id, nil
}

func (m *MongoDriver) Delete(ctx context.Context, collection string, filter interface{}) (int64, error) {
	col := m.client.Database("productservice").Collection(collection)
	result, err := col.DeleteMany(ctx, filter)
	return result.DeletedCount, err
}

// Returns the mongodb client used by the repo
// @param ctx context
func (m *MongoDriver) GetClient() *mongo.Client {
	return m.client
}

// var Driver = &MongoDriver{
// 	uri: config.Config.Get("MONGODB_URL"),
// }

// func init() {

// 	fmt.Println("Initializing mongo driver")
// 	Driver.Connect(context.Background())
// }

func NewMongoDriver(c config.IConfiguration) IMongoDriver {
	var mongoUrl = c.Get("MONGODB_URL")
	driver := &MongoDriver{
		uri: mongoUrl,
	}
	driver.Connect(context.Background())
	return driver
}

var Module = func() IMongoDriver {
	c := config.Config
	c.Init()
	return NewMongoDriver(c)
}()
