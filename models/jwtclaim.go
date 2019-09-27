package models

import (
  jwt "github.com/dgrijalva/jwt-go"
)

type TheClaims struct {
    jwt.StandardClaims
    Username string `json:"Username"`
    Email    string `json:"Email"`
    Group    string `json:"Group"`
}
