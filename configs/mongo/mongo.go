package mongo

import (
    "context"
    "fmt"
    "log"
    . "we.care/they-chat/configs"
    // "go.mongodb.org/mongo-driver/bson"
    "go.mongodb.org/mongo-driver/mongo"
    "go.mongodb.org/mongo-driver/mongo/options"
)

type Conn struct {
	Server   string
	Database string
}

var config = Config{}
var singleton *conn
var once sync.Once
var db *mgo.Database

func init(){
    config.Read()
    this.Server = config.Server
	this.Database = config.Database
	this.Connect()

}


func GetDB() *conn {
    once.Do(func() {
        singleton = Connect()
    })
    return singleton
}

func Connect () (Conn) {

    // Set client options
clientOptions := options.Client().ApplyURI("mongodb://"+config.Server+":27017")

// Connect to MongoDB
client, err := mongo.Connect(context.TODO(), clientOptions)

if err != nil {
    log.Fatal(err)
}

// Check the connection
err = client.Ping(context.TODO(), nil)

if err != nil {
    log.Fatal(err)
}

fmt.Println("Connected to MongoDB!")
// Connected
db:=client.Database(config.Database)
return Conn{}
}
// collection := client.Database(config.Database).Collection("trainers")

// fmt.Println(collection)
// // collection

// err = client.Disconnect(context.TODO())

// if err != nil {
//     log.Fatal(err)
// }
// fmt.Println("Connection to MongoDB closed.")
// close
// }
