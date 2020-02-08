package mongodb

import (
    "context"
    "fmt"
    "log"
    "time"

    "go.mongodb.org/mongo-driver/mongo"
    "go.mongodb.org/mongo-driver/mongo/options"
		"go.mongodb.org/mongo-driver/mongo/readpref"
		
		"go-gin-cli/pkg/setting"
)

var MongoClient *mongo.Client

func init() {
    var err error

    mongoURI := fmt.Sprintf(setting.MongoSetting.Host)

    // connect to mongo
    ctx, _ := context.WithTimeout(context.Background(), 10*time.Second)
    MongoClient, err = mongo.Connect(ctx, options.Client().ApplyURI(mongoURI))
    if err != nil {
        log.Fatal("connet err:", err.Error())
        panic(err)
    }

    // Check the connection
    ctx, _ = context.WithTimeout(context.Background(), 10*time.Second)
    err = MongoClient.Ping(ctx, readpref.Primary())
    if err != nil {
        fmt.Println("could not ping to mongo db service: %v\n", err)
        return
    } else {
        fmt.Println("Successful connected to MongoDB!")
    }    
}