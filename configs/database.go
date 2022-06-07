package configs

import (
	"context"
	"fmt"
	"github.com/MrMohebi/golang-gin-boilerplate.git/common"
	"go.mongodb.org/mongo-driver/mongo"
	"go.mongodb.org/mongo-driver/mongo/options"
	"time"
)

// DB Client instance
var DB *mongo.Client

func GetDBClint() *mongo.Client {
	if DB != nil {
		return DB
	}
	DB = ConnectDB()
	return DB
}

func ConnectDB() *mongo.Client {
	EnvSetup()
	client, err := mongo.NewClient(options.Client().ApplyURI(EnvMongoURI()))
	common.IsErr(err)

	ctx, _ := context.WithTimeout(context.Background(), 10*time.Second)
	err = client.Connect(ctx)
	common.IsErr(err)

	//ping the database
	err = client.Ping(ctx, nil)
	common.IsErr(err)

	fmt.Println("Connected to MongoDB")
	DB = client
	return client
}

// GetCollection getting database collections
func GetCollection(client *mongo.Client, collectionName string) *mongo.Collection {
	collection := client.Database(EvnMongoDB()).Collection(collectionName)
	return collection
}
