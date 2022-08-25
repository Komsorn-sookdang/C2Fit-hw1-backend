package databases

import (
	"context"
	"log"
	"os"
	"time"

	"github.com/joho/godotenv"
	"go.mongodb.org/mongo-driver/mongo"
	"go.mongodb.org/mongo-driver/mongo/options"
	"go.mongodb.org/mongo-driver/mongo/readpref"
)

type mongoDB struct {
	Client   *mongo.Client
	Database *mongo.Database
}

func ConnectToDatabase() {
	if err := godotenv.Load(); err != nil {
		log.Println("No .env file found")
	}

	uri := os.Getenv("MONGODB_URI")
	if uri == "" {
		log.Println("No MONGODB_URI")
	}
	clientOptions := options.Client().ApplyURI(uri)

	ctx, cancel := context.WithTimeout(context.Background(), 10*time.Second)
	defer cancel()
	client, err := mongo.Connect(ctx, clientOptions)
	if err != nil {
		log.Fatal("cant connect to MongoDB :", err)
	}
	if err := client.Ping(context.TODO(), readpref.Primary()); err != nil {
		// Can't connect to Mongo server
		log.Fatal("cant ping Mongo server :", err)
	}
	db := client.Database("c2fit-hw-db")

	MyDB = mongoDB{
		Client:   client,
		Database: db,
	}
	log.Println("Connected to MongoDB Successfully")
}

var MyDB mongoDB
