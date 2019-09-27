package handler

import (
  "net/http"
  mid "api/middleware"
  models "api/models"
  jwt "github.com/dgrijalva/jwt-go"
  "encoding/json"
  "time"
  "os"
  . "fmt"
  "log"
  "github.com/joho/godotenv"
)



// var APPLICATION_NAME = "My Simple JWT App"
// var APPLICATION_NAME = os.Getenv("APPLICATION_NAME")
var LOGIN_EXP_DURATION = time.Duration(1) * time.Hour
var JWT_SIGNING_METHOD = jwt.SigningMethodHS256
// var JWT_SIGNATURE_KEY = []byte("not so secret")
// var JWT_SIGNATURE_KEY = []byte(os.Getenv("SIGNATURE_KEY"))



type  M map[string]interface{}

func Loginhandler(w http.ResponseWriter, r *http.Request) {

  err := godotenv.Load()
  if err != nil {
    log.Fatal("Error loading .env file")
  }
  Println(os.Getenv("APPLICATION_NAME"))
  // Println(APPLICATION_NAME)
  username, password, ok := r.BasicAuth()
  if !ok {
    http.Error(w, "Invalid username or password", http.StatusBadRequest)
    return
  }
  //fungsi  authenticateUser dengan return semua info user
  // ok, userInfo := authenticateUser(username, password)
  ok, userInfo := mid.AuthenticateUser(username, password)

  if !ok {
    http.Error(w, "Invalid username or password", http.StatusBadRequest)
    return
  }
  //
  //initiate objek claim
  claims := models.TheClaims{
          StandardClaims: jwt.StandardClaims{
          Issuer:    os.Getenv("APPLICATION_NAME"),
          ExpiresAt: time.Now().Add(LOGIN_EXP_DURATION).Unix(),
      },
      Username: userInfo.Username,
      Email:    userInfo.Email,
      Group:    userInfo.Group,
  }
  //buat token baru
  token := jwt.NewWithClaims(
      JWT_SIGNING_METHOD,
      claims,
  )
  //menandatangani token yg akan dikembalikan ke client
  signedToken, err := token.SignedString([]byte(os.Getenv("SIGNATURE_KEY")))
  if err != nil {
      http.Error(w, err.Error(), http.StatusBadRequest)
      return
  }
  //tokennya dijadiin json
  tokenString, _ := json.Marshal(M{ "token": signedToken })
  w.Write([]byte(tokenString))
}
