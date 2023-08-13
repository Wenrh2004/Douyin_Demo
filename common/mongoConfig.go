package common

import (
	"context"
	"fmt"
	"github.com/spf13/viper"
	"go.mongodb.org/mongo-driver/bson"
	"go.mongodb.org/mongo-driver/mongo"
	"go.mongodb.org/mongo-driver/mongo/options"
)

type MongoDB struct {
	passWord string
	userName string
}

func getMongoDBConfig() *MongoDB {
	viper.SetConfigType("yaml")
	viper.SetConfigFile("application.yaml")
	if err := viper.ReadInConfig(); err != nil {
		if _, ok := err.(viper.ConfigFileNotFoundError); ok {
			panic("config file can not be founded")
		}
		panic("config file read error+" + err.Error())
	}
	return &MongoDB{
		passWord: viper.GetString("mongodb.password"),
		userName: viper.GetString("mongodb.username"),
	}
}

// InitMongoDB ==> Initialize mongodb database connection.
// 1. Use the SetServerAPIOptions() method to set the Stable API version to 1
// 2. Create a new client and connect to the server
// 3. Send a ping to confirm a successful connection
func InitMongoDB() {
	var config = getMongoDBConfig()
	dsn := fmt.Sprintf("mongodb+srv://%s:%s@douyin.js4evzp.mongodb.net/?retryWrites=true&w=majority", config.userName, config.passWord)
	// Use the SetServerAPIOptions() method to set the Stable API version to 1
	serverAPI := options.ServerAPI(options.ServerAPIVersion1)
	opts := options.Client().ApplyURI(dsn).SetServerAPIOptions(serverAPI)

	// Create a new client and connect to the server
	client, err := mongo.Connect(context.TODO(), opts)
	if err != nil {
		panic(err)
	}

	defer func() {
		if err = client.Disconnect(context.TODO()); err != nil {
			panic(err)
		}
	}()

	// Send a ping to confirm a successful connection
	if err := client.Database("admin").RunCommand(context.TODO(), bson.D{{"ping", 1}}).Err(); err != nil {
		panic(err)
	}
	fmt.Println("Pinged your deployment. You successfully connected to MongoDB!")
}
