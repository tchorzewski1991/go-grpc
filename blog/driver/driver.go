package driver

import (
	"github.com/mongodb/mongo-go-driver/mongo"
	"context"
	"fmt"
	h "github.com/tchorzewski1991/go-grpc/blog/helpers"
)

type DB struct {
	Mongo *mongo.Database
}

var dbConn = &DB{}

func ConnectMongoDB(ctx context.Context) *DB {
	uri := fmt.Sprintf(`mongodb://%s:%s`,
		ctx.Value("MONGO_DB_HOST"),
		ctx.Value("MONGO_DB_PORT"),
	)

	client, err := mongo.NewClient(uri)
	h.LogIfError(err, "Failed to initialize mongo db client: %v")

	err = client.Connect(ctx)
	h.LogIfError(err, "Failed to connect mongo db client: %v")

	dbConn.Mongo = client.Database("root-db")

	return dbConn
}
