package middlewares

import (
  "strings"
  "context"
  "net/http"
  "github.com/alexkarpovich/quiqstee-user/lib"
)

func User(next http.Handler) http.Handler {
  return http.HandlerFunc(func(w http.ResponseWriter, r *http.Request) {

    reqToken := r.Header.Get("Authorization")

    if len(reqToken) > 0 {
      splitToken := strings.Split(reqToken, "Bearer ")
      reqToken = splitToken[1]

      user := lib.GetTokenClaims(reqToken)

      ctx := context.WithValue(r.Context(), "user", user)

      r = r.WithContext(ctx)
    }

    next.ServeHTTP(w, r)
  })
}
