package main

import (
	"context"
	"crud_gin_gonic/controllers"
	"crud_gin_gonic/services"
	"fmt"
	"log"

	"github.com/gin-gonic/gin"
	"go.mongodb.org/mongo-driver/mongo"
	"go.mongodb.org/mongo-driver/mongo/options"
	"go.mongodb.org/mongo-driver/mongo/readpref"
)

var (
	server *gin.Engine
	userservice services.UserService
	usercontoller controllers.UserController
	ctx    context.Context
	usercollection *mongo.Collection
	mongoclient *mongo.Client
	err error
)

func init(){
	ctx = context.TODO()

	mongoconn := options.Client().ApplyURI("mongodb://127.0.0.1:27017/?compressors=disabled&gssapiServiceName=mongodb")
	mongoclient , err = mongo.Connect(ctx,mongoconn)
	if err!= nil {
		log.Fatal(err)
	}
	err = mongoclient.Ping(ctx , readpref.Primary())
	if err != nil {
		log.Fatal(err)
	}

	fmt.Println("mongo connection established")

	usercollection = mongoclient.Database("kgf").Collection("rocky")
	userservice = services.NewUserService(usercollection, ctx)
	usercontoller = controllers.New(userservice)
	server = gin.Default()


}
func main() {

	defer mongoclient.Disconnect(ctx)
	basepath := server.Group("/v1")
	usercontoller.RegisterUserRoutes(basepath)
	log.Fatal(server.Run(":9090"))
}