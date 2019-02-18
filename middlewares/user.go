package middlewares

import (
  "strings"
  "context"
  "net/http"
  "github.com/alexkarpovich/quiqstee-user/lib"
  "github.com/alexkarpovich/quiqstee-user/database"
  "github.com/alexkarpovich/quiqstee-user/database/models"
)

func User(next http.Handler) http.Handler {
  return http.HandlerFunc(func(w http.ResponseWriter, r *http.Request) {

    var user models.User
    reqToken := r.Header.Get("Authorization")

    if len(reqToken) > 0 {
      splitToken := strings.Split(reqToken, "Bearer ")
      reqToken = splitToken[1]

      userId := lib.GetTokenClaims(reqToken)
      err := database.Db.Where(&models.User{ID: userId}).First(&user)

      if err == nil {
        ctx := context.WithValue(r.Context(), "user", user)

        r = r.WithContext(ctx)
      }      
    }

    next.ServeHTTP(w, r)
  })
}
