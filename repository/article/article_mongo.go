package article

import (
  "context"
  "log"
  . "fmt"
  models "api/models"
  dbCon "api/driver"
  // "go.mongodb.org/mongo-driver/mongo"
  // "go.mongodb.org/mongo-driver/mongo/options"
  "go.mongodb.org/mongo-driver/bson"
)


func FindAll() []models.Article {

  var Articles []models.Article

  db, err := dbCon.Connect()
  if err != nil {
      log.Fatal(err.Error())
  }

  collection := db.Collection("article")
  cur, err := collection.Find(context.Background(), bson.D{})

  for cur.Next(context.Background()) {
      var row models.Article
      err := cur.Decode(&row)
      if err != nil {
          log.Fatal(err.Error())
      }
      Articles = append(Articles, row)
  }

  // Println(Articles)

  return Articles
}

func Find(id string) models.Article {
  var article models.Article
  db, err := dbCon.Connect()
  if err != nil {
      log.Fatal(err.Error())
  }

  // if id == "art" {
  //   Println("id is not given")
  //   return
  // }

  collection := db.Collection("article")
  cur := collection.FindOne(context.Background(), bson.M{"ArticleID": id})
  // defer cur.Close(ctx)

  cur.Decode(&article)

  return article
}

func Insert(id string, title string, desc string, content string) models.Article {
  db, err := dbCon.Connect()

  var article models.Article
  if err != nil {
    log.Fatal(err.Error())
  }

  collection := db.Collection("article")
  //cari jumlah document yg udah ada
  // id := collection.Count()+1
  _, err = collection.InsertOne(context.Background(), models.Article{id, title, desc, content})
  if err != nil {
    log.Fatal(err.Error())
  }

  article.ArticleID = id
  article.Title = title
  article.Desc = desc
  article.Content = content

  Println("Insert success!")

  return article
}

func Remove(id string) {
  db, err := dbCon.Connect()
  if err != nil {
        log.Fatal(err.Error())
  }
  // if id == "art" {
  //   Println("id is not given")
  //   return
  // }

  collection := db.Collection("article")
  _, err = collection.DeleteOne(context.Background(), bson.M{"ArticleID":id})
  if err != nil {
    log.Fatal(err.Error())
  }

  Println("Delete success!")

  return
}

func Update(id string, title string, desc string, content string) models.Article {
  db, err := dbCon.Connect()
  var article models.Article

  if err != nil {
      log.Fatal(err.Error())
  }

  // if id == "art" {
  //   Println("id is not given")
  //   return
  // }

  collection := db.Collection("article")

  var selector = bson.M{"ArticleID": id}
  var changes = models.Article{id, title, desc, content}
  _, err = collection.UpdateOne(context.Background(), selector, bson.M{"$set": changes})
  if err != nil {
    log.Fatal(err.Error())
  }

  article.ArticleID = id
  article.Title = title
  article.Desc = desc
  article.Content = content

  Println("Update success!")

  return article
}
