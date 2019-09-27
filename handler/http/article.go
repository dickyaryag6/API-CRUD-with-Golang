package handler

import(
  "encoding/json"
  . "fmt"
  "io/ioutil"
  "net/http"
  "github.com/gorilla/mux"
  models "api/models"
  // repository "api/repository/article"
  method "api/repository/article"
  // mid "api/middleware"

)

func GetAllArticles(w http.ResponseWriter, r *http.Request) {

  Println("Endpoint Hit: getAllArticles")
  articles:=method.FindAll()
  //query data
  json.NewEncoder(w).Encode(articles) //encoding our articles array into a JSON string and then writing as part of our response.
}

func GetArticle(w http.ResponseWriter, r *http.Request) {
  Println("Endpoint Hit: getArticle")

  vars := mux.Vars(r) //ambil isi dari w,isinya parameter dari endpoint
  // Println("sebelum fungsi vars : ",vars)
  id := "art"+vars["id"] //ambil nilai parameter id dari endpoint

  // Println("sebelum fungsi : ",id)
  article:=method.Find(id)
  json.NewEncoder(w).Encode(article)
  // articles:=method.FindAll()
  // json.NewEncoder(w).Encode(articles)
}

func CreateArticle(w http.ResponseWriter, r *http.Request) {
  Println("Endpoint Hit: createArticle")
  reqBody, _ := ioutil.ReadAll(r.Body) //ngebaca json

  var article models.Article

  json.Unmarshal(reqBody, &article) //ubah json ke dalam bentuk objek
  Println(article)

  art:=method.Insert(article.ArticleID, article.Title, article.Desc, article.Content)

  json.NewEncoder(w).Encode(art)
}

func DeleteArticle(w http.ResponseWriter, r *http.Request) {
  Println("Endpoint Hit: deleteArticle")
  vars := mux.Vars(r) //ambil isi dari w,isinya parameter dari endpoint
  id := "art"+vars["id"] //ambil id dari endpoint
  // query data
  method.Remove(id)
}

func UpdateArticle(w http.ResponseWriter, r *http.Request) {
  Println("Endpoint Hit: updateArticle")
  vars := mux.Vars(r) //ambil isi dari w,isinya parameter dari endpoint
  id := "art"+vars["id"] //ambil id dari endpoint

  //ambil isi dari body
  reqBody, _ := ioutil.ReadAll(r.Body)
  var article models.Article
  json.Unmarshal(reqBody, &article)
  //update data
  art:=method.Update(id, article.Title, article.Desc, article.Content)

  json.NewEncoder(w).Encode(art)
}
