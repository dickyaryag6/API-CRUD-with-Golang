package driver

import(
  "context"
  "go.mongodb.org/mongo-driver/mongo"
  "go.mongodb.org/mongo-driver/mongo/options"
  "os"
  // . "fmt"
  "github.com/joho/godotenv"
  "log"
)

func Connect() (*mongo.Database, error) {

  err := godotenv.Load()
  if err != nil {
    log.Fatal("Error loading .env file")
  }
  dbURI:=os.Getenv("DB_URI")

  clientOptions := options.Client().ApplyURI(dbURI)
  client, err := mongo.NewClient(clientOptions)

  if err != nil {
    return nil, err
  }

  err = client.Connect(context.Background())
  if err != nil {
        return nil, err
    }
  dbName:=os.Getenv("DB_NAME")
  return client.Database(dbName), nil

}
