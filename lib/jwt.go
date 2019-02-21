package lib

import (
  "log"
  "fmt"
  "time"
  "github.com/dgrijalva/jwt-go"
  "github.com/alexkarpovich/quiqstee-user/database/models"
)

type TokenClaims struct {
  jwt.StandardClaims
  Uid uint `json:"uid"`
  Name string `json:"name"`
  Img string `json:"img"`
}

func NewToken(user *models.User) string {
  token := jwt.NewWithClaims(jwt.SigningMethodHS256, TokenClaims {
    jwt.StandardClaims{
      ExpiresAt: time.Now().Add(time.Hour * 24).Unix(),
      Issuer: "quiqstee-user",
    },
    user.ID,
    user.FullName(),
    "http://icons.iconarchive.com/icons/paomedia/small-n-flat/256/sign-info-icon.png",
  })

  tokenString, err := token.SignedString([]byte("mySigningKey"))
  if err != nil {
    log.Fatal(err)
  }

  return tokenString
}

func GetTokenClaims(tokenString string) (*TokenClaims, bool) {
  token, _ := jwt.ParseWithClaims(tokenString, &TokenClaims{}, func(token *jwt.Token) (interface{}, error) {
    if _, ok := token.Method.(*jwt.SigningMethodHMAC); !ok {
        return nil, fmt.Errorf("Unexpected signing method: %v", token.Header["alg"])
    }

    return []byte("mySigningKey"), nil
  })

  if claims, ok := token.Claims.(*TokenClaims); ok && token.Valid {
      return claims, true
  }

  return nil, false
}
