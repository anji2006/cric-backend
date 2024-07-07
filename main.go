package main

import (
	"context"
	"fmt"
	"log"

	"cric.com/backend/controllers"
	"cric.com/backend/services"
	"github.com/gin-gonic/gin"
	"go.mongodb.org/mongo-driver/bson"
	"go.mongodb.org/mongo-driver/mongo"
	"go.mongodb.org/mongo-driver/mongo/options"
	"go.mongodb.org/mongo-driver/mongo/readpref"
)

var (
	server         *gin.Engine
	teamservice    services.TeamService
	teamcontroller controllers.TeamController
	ctx            context.Context
	teamcollection *mongo.Collection
	mongoclient    *mongo.Client
	err            error
)

func init() {
	ctx = context.TODO()

	mongoconn := options.Client().ApplyURI("mongodb://localhost:27017")
	mongoclient, err = mongo.Connect(ctx, mongoconn)

	if err != nil {
		log.Fatal(err)
	}

	err = mongoclient.Ping(ctx, readpref.Primary())

	if err != nil {
		log.Fatal(err)
	}

	fmt.Println("Mongo connection is established.....")

	teamcollection = mongoclient.Database("cric").Collection("team")
	indexModel := mongo.IndexModel{
		Keys:    bson.M{"name": 1},
		Options: options.Index().SetUnique(true),
	}
	_, err := teamcollection.Indexes().CreateOne(context.TODO(), indexModel)
	if err != nil {
		log.Fatal(err)
	}
	teamservice = services.NewTeamService(teamcollection, ctx)
	teamcontroller = controllers.New(teamservice)

	server = gin.Default()

}

func main() {
	defer mongoclient.Disconnect(ctx)

	basepath := server.Group("/v1")

	teamcontroller.RegisterTeamRoutes(basepath)

	log.Fatal(server.Run(":9090"))

}
