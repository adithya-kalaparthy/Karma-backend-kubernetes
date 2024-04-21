package handlers

import (
	"context"
	"encoding/json"
	"fmt"
	"log"
	"net/http"
	"os"
	"strconv"

	"github.com/gin-gonic/gin"
	"github.com/karma/karma-backend/pkg/api/types"
	"go.mongodb.org/mongo-driver/bson"
	"go.mongodb.org/mongo-driver/mongo"
)

func GetTasks(c *gin.Context) {

	mongoSession, err := c.Get("mongoSession")
	if !err || mongoSession == nil {
		log.Fatalf("Middleware did not provide mongoSession")
		return
	}

	dbSession, ok := mongoSession.(mongo.Session)

	if !ok {
		log.Fatalf("MongoDB session is not valid")
		return
	}

	filter := bson.D{
		{Key: "project_id", Value: 1},
	}

	collection := dbSession.Client().Database(os.Getenv("MONGODB_DATABASE")).Collection("karma")

	cursor, errFilteringCollection := collection.Find(context.TODO(), filter)

	if errFilteringCollection != nil {
		panic(errFilteringCollection)
	}

	defer cursor.Close(context.TODO())

	var results []types.Task

	if err := cursor.All(context.TODO(), &results); err != nil {
		panic(err)
	}

	c.JSON(http.StatusOK, gin.H{
		"task_list": results,
	})
}

func CreateTask(c *gin.Context) {

	body := c.Request.Body

	mongoSession, err := c.Get("mongoSession")
	if !err || mongoSession == nil {
		log.Fatalf("Middleware did not provide mongoSession")
		return
	}

	dbSession, ok := mongoSession.(mongo.Session)

	if !ok {
		log.Fatalf("MongoDB session is not valid")
		return
	}

	// Close the body after create task is correctly run.
	defer body.Close()

	var jsonData types.Task
	errConvertingJson := json.NewDecoder(body).Decode(&jsonData)

	if errConvertingJson != nil {
		log.Fatalf("error decoding json %v", errConvertingJson)
		body.Close()
		return
	}

	collection := dbSession.Client().Database(os.Getenv("MONGODB_DATABASE")).Collection("karma")

	result, errInsertingDoc := collection.InsertOne(context.TODO(), jsonData)

	if errInsertingDoc != nil {
		panic(errInsertingDoc)
	}

	c.JSON(http.StatusCreated, gin.H{
		"Insert result": result,
	})
}

func UpdateTask(c *gin.Context) {
	body := c.Request.Body

	// Close the body after create task is correctly run.
	defer body.Close()

	var jsonData types.Task
	errConvertingJson := json.NewDecoder(body).Decode(&jsonData)

	if errConvertingJson != nil {
		log.Fatalf("error decoding json %v", errConvertingJson)
		body.Close()
		return
	}

	mongoSession, err := c.Get("mongoSession")
	if !err || mongoSession == nil {
		log.Fatalf("Middleware did not provide mongoSession")
		return
	}

	dbSession, ok := mongoSession.(mongo.Session)

	if !ok {
		log.Fatalf("MongoDB session is not valid")
		return
	}

	collection := dbSession.Client().Database(os.Getenv("MONGODB_DATABASE")).Collection("karma")

	filter := bson.D{
		{Key: "task_id", Value: jsonData.TaskID},
	}

	updateDoc := bson.D{
		{Key: "$set", Value: jsonData},
	}

	// Update the collection with new data.
	result, errUpdatingDb := collection.UpdateOne(context.TODO(), filter, updateDoc)

	if errUpdatingDb != nil {
		fmt.Print(jsonData)
		panic(errUpdatingDb)
	}

	c.JSON(http.StatusOK, gin.H{
		"updated tasks": result,
	})
}

func DeleteTask(c *gin.Context) {
	deletedTaskId, errConvertingInt := strconv.ParseInt(c.Query("task_id"), 10, 64)

	if errConvertingInt != nil {
		log.Fatalf("Error converting to int %v", errConvertingInt)
	}

	mongoSession, err := c.Get("mongoSession")
	if !err || mongoSession == nil {
		log.Fatalf("Middleware did not provide mongoSession")
		return
	}

	dbSession, ok := mongoSession.(mongo.Session)

	if !ok {
		log.Fatalf("MongoDB session is not valid")
		return
	}

	collection := dbSession.Client().Database(os.Getenv("MONGODB_DATABASE")).Collection("karma")

	filter := bson.D{
		{Key: "task_id", Value: deletedTaskId},
	}

	result, errDeletingTask := collection.DeleteOne(context.TODO(), filter)

	if errDeletingTask != nil {
		panic(errDeletingTask)
	}

	c.JSON(http.StatusCreated, gin.H{
		"deleted list": result,
	})
}

func CreateMultipleTasks(c *gin.Context) {
	body := c.Request.Body

	mongoSession, err := c.Get("mongoSession")
	if !err || mongoSession == nil {
		log.Fatalf("Middleware did not provide mongoSession")
		return
	}

	dbSession, ok := mongoSession.(mongo.Session)

	if !ok {
		log.Fatalf("MongoDB session is not valid")
		return
	}

	// Close the body after create task is correctly run.
	defer body.Close()

	var jsonData []interface{}
	errConvertingJson := json.NewDecoder(body).Decode(&jsonData)

	if errConvertingJson != nil {
		log.Fatalf("error decoding json %v", errConvertingJson)
		body.Close()
		return
	}

	collection := dbSession.Client().Database(os.Getenv("MONGODB_DATABASE")).Collection("karma")

	result, errInsertingDoc := collection.InsertMany(context.TODO(), jsonData)

	if errInsertingDoc != nil {
		panic(errInsertingDoc)
	}

	c.JSON(http.StatusCreated, gin.H{
		"Insert result": result,
	})
}
