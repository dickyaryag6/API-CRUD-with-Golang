package middleware

import (
  dbCon "api/driver"
  models "api/models"
  "go.mongodb.org/mongo-driver/bson"
  "log"
  "context"
)

func AuthenticateUser(username, password string) (bool, models.User) {
  //cari document dengan username dan password yg diberikan
  db, err := dbCon.Connect()
  if err != nil {
      log.Fatal(err.Error())
  }
  var user models.User

  collection := db.Collection("user")
  cur := collection.FindOne(context.Background(), bson.M{"username": username})
  cur.Decode(&user)

  if user.Username == "" {
    return false, user
  } else if password != user.Password {
    return false, user
  }

  return true, user
}
