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
  User *models.User `json:"user"`
}

func NewToken(user *models.User) string {
  token := jwt.NewWithClaims(jwt.SigningMethodHS256, TokenClaims {
    jwt.StandardClaims{
      ExpiresAt: time.Now().Add(time.Hour * 24).Unix(),
      Issuer: "quiqstee-user",
    },
    user,
  })

  tokenString, err := token.SignedString([]byte("mySigningKey"))
  if err != nil {
    log.Fatal(err)
  }

  return tokenString
}

func GetTokenClaims(tokenString string) *models.User {
  token, _ := jwt.ParseWithClaims(tokenString, &TokenClaims{}, func(token *jwt.Token) (interface{}, error) {
    if _, ok := token.Method.(*jwt.SigningMethodHMAC); !ok {
        return nil, fmt.Errorf("Unexpected signing method: %v", token.Header["alg"])
    }

    // hmacSampleSecret is a []byte containing your secret, e.g. []byte("my_secret_key")
    return []byte("mySigningKey"), nil
  })

  if claims, ok := token.Claims.(*TokenClaims); ok && token.Valid {
      return claims.User
  }

  return nil
}
