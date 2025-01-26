package middlewares

import (
	"context"
	"log"
	"net/url"
	"os"
	"time"

	"github.com/gin-gonic/gin"
	"go.mongodb.org/mongo-driver/mongo"
	"go.mongodb.org/mongo-driver/mongo/options"
	"go.mongodb.org/mongo-driver/mongo/readpref"
	"go.mongodb.org/mongo-driver/mongo/writeconcern"
)

const (
	mongoDBHostEnvKey      = "DB_HOST"
	mongoDBPortEnvKey      = "DB_PORT"
	mongoDBUsernameEnvKey  = "DB_USER"
	mongoDBPasswordEnvKey  = "DB_PASS"
	mongoDBDatabaseEnvKey  = "MONGODB_DATABASE"
	envKey                 = "ENV"
	mongoClusterNameEnvKey = "MONGO_CLUSTER_NAME"
)

// MongoDBMiddleware connects to MongoDB and provides a database session in the Gin API request context.
func MongoDBMiddleware() gin.HandlerFunc {
	return func(c *gin.Context) {
		// Get MongoDB connection details from environment variables
		mongoDBHost := os.Getenv(mongoDBHostEnvKey)
		mongoDBPort := os.Getenv(mongoDBPortEnvKey)
		mongoDBUsername := os.Getenv(mongoDBUsernameEnvKey)
		mongoDBPassword := os.Getenv(mongoDBPasswordEnvKey)
		env := os.Getenv(envKey)
		mongoClusterName := os.Getenv(mongoClusterNameEnvKey)
		serverApi := options.ServerAPI(options.ServerAPIVersion1)

		// Construct MongoDB connection URI
		mongoDBURI := "mongodb://" +
			url.QueryEscape(mongoDBUsername) +
			":" +
			url.QueryEscape(mongoDBPassword) +
			"@" +
			url.QueryEscape(mongoDBHost)

		if env == "local" {
			mongoDBURI += ":" + mongoDBPort + "/"
		} else if env == "prod" {
			mongoDBURI = "mongodb+srv://" + url.QueryEscape(mongoDBUsername) +
				":" +
				url.QueryEscape(mongoDBPassword) +
				"@" +
				url.QueryEscape(mongoDBHost) +
				"/?retryWrites=true&w=majority&appName=" +
				mongoClusterName
		}

		log.Printf("ENV recieved as %s", env)
		log.Printf("DB connection string %s", mongoDBURI)

		// Set up MongoDB client options
		clientOptions := options.Client().ApplyURI(mongoDBURI).SetServerAPIOptions(serverApi)
		clientOptions.SetReadPreference(readpref.Primary())    // Set read preference
		clientOptions.SetWriteConcern(writeconcern.Majority()) // Set write concern
		clientOptions.SetMaxConnIdleTime(10 * time.Second)     // Set max conn idle time

		ctx, cancel := context.WithTimeout(context.Background(), 10*time.Second)
		defer cancel()

		// Connect to MongoDB
		client, err := mongo.Connect(ctx, clientOptions)
		if err != nil {
			log.Fatalf("Error creating MongoDB client: %v", err)
		}

		// Start a new transaction
		session, err := client.StartSession()
		if err != nil {
			log.Fatalf("Error starting session: %v", err)
		}

		// Begin a transaction
		err = session.StartTransaction()
		if err != nil {
			log.Fatalf("Error starting transaction: %v", err)
		}

		// Add the MongoDB client and session to the Gin context
		c.Set("mongoClient", client)
		c.Set("mongoSession", session)

		// Pass the control to the next middleware or route handler
		c.Next()

		// Check if there was an error in the request processing
		if len(c.Errors) > 0 {
			// Roll back the transaction if an error occurred
			err := session.AbortTransaction(ctx)
			if err != nil {
				log.Printf("Error rolling back transaction: %v", err)
			}
		} else {
			// Commit the transaction if there were no errors
			err := session.CommitTransaction(ctx)
			if err != nil {
				log.Printf("Error committing transaction: %v", err)
			}
		}

		// Disconnect from MongoDB after the request is processed
		err = client.Disconnect(ctx)
		if err != nil {
			log.Printf("Error disconnecting from MongoDB: %v", err)
		}
	}
}
