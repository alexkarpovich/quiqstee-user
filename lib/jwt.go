package lib

import (
  "log"
  "time"
  "github.com/dgrijalva/jwt-go"
  "github.com/alexkarpovich/quiqstee-user/database/models"
)

func NewToken(user *models.User) string {
  token := jwt.New(jwt.SigningMethodHS256)

  token.Claims = jwt.MapClaims{
    "exp": time.Now().Add(time.Hour * 24).Unix(),
    "iat": time.Now().Unix(),
    "name": user.FullName(),
  }

  tokenString, err := token.SignedString([]byte("mySigningKey"))
  if err != nil {
    log.Fatal(err)
  }

  return tokenString
}
