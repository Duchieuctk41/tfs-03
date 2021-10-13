package driver

import (
	"context"
	"fmt"
	"time"

	"go.mongodb.org/mongo-driver/mongo"
	"go.mongodb.org/mongo-driver/mongo/options"
	"go.mongodb.org/mongo-driver/mongo/readpref"
)

type Mongodb struct {
	Client *mongo.Client
}

var Mongo = &Mongodb{}

func ConnectMongoDB(user, password string) *Mongodb {
	connStr := fmt.Sprintf("mongodb+srv://%s:%s@cluster0.duth3.mongodb.net/myFirstDatabase?retryWrites=true&w=majority",
		user, password)
	client, err := mongo.NewClient(options.Client().ApplyURI(connStr))
	if err != nil {
		panic(err)
	}

	ctx, _ := context.WithTimeout(context.Background(), 10*time.Second)
	err = client.Connect(ctx)
	if err != nil {
		panic(err)
	}

	err = client.Ping(ctx, readpref.Primary()) // ping vào bản chính
	if err != nil {
		panic(err)
	}

	fmt.Println("connection ok")
	Mongo.Client = client
	return Mongo

}
