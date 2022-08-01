package api

import (
	"context"
	"encoding/base64"
	"go.mongodb.org/mongo-driver/mongo"
	"go.mongodb.org/mongo-driver/mongo/options"
	"log"
	"os"
)

const (
	webPort  = "80"
	rpcPort  = "5001"
	mongoUrl = "mongodb://mongo:27017"
	gRpcPort = "50001"
)

var client *mongo.Client

type Config struct {
}

func main() {
	mongoClient, err := connectToMongoDB()
	if err != nil {
		log.Panic(err)
	}
	client = mongoClient
}

func connectToMongoDB() (*mongo.Client, error) {
	clientOptions := options.Client().ApplyURI(mongoUrl)
	clientOptions.SetAuth(options.Credential{
		Username: decode(os.Getenv("MONGO_USER")),
		Password: decode(os.Getenv("MONGO_PASSWORD")),
	})
	return mongo.Connect(context.TODO(), clientOptions)
}

func decode(value string) string {
	result, err := base64.StdEncoding.DecodeString(value)
	if err != nil {
		log.Panic(err)
	}
	return string(result)
}
