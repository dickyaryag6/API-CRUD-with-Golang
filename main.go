package main

import (
  . "fmt"
  "net/http"
  "log"
  "github.com/gorilla/mux"
  handler "api/handler/http"
  mid "api/middleware"
)



// TODO: tambah .env

// func homepage(w http.ResponseWriter, r *http.Request) {
//   Fprintf(w, "Welcome to Homepage")
//   Println("Endpoint Hit: homepage")
// }


func handleRequest() {

  router := mux.NewRouter().StrictSlash(true) //strictslash true supaya kalo
  //user make slash di akhir atau ga, tetep bisa jalan
  // router.Use(mid.MiddlewareJWTAuthorization)

  // router.HandleFunc("/", homepage)
  router.HandleFunc("/login", handler.Loginhandler).Methods("POST")

  art := router.PathPrefix("/articles").Subrouter()
  art.Use(mid.MiddlewareJWTAuthorization)
  art.HandleFunc("", handler.GetAllArticles).Methods("GET")
  art.HandleFunc("/{id}", handler.GetArticle).Methods("GET")
  art.HandleFunc("", handler.CreateArticle).Methods("POST")
  art.HandleFunc("/{id}", handler.DeleteArticle).Methods("DELETE")
  art.HandleFunc("/{id}", handler.UpdateArticle).Methods("PUT")

  Println("Server listen at :8000")

  log.Fatal(http.ListenAndServe(":8000", router))
}

func main() {
  handleRequest()
}
