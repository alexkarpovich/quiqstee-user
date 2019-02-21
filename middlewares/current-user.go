package middlewares

import (
    "strings"
    "context"
    "net/http"
    "github.com/alexkarpovich/quiqstee-user/lib"
    "github.com/alexkarpovich/quiqstee-user/database"
    "github.com/alexkarpovich/quiqstee-user/database/models"
)

func CurrentUser(next http.Handler) http.Handler {
    return http.HandlerFunc(func(w http.ResponseWriter, r *http.Request) {

        var user models.User
        reqToken := r.Header.Get("Authorization")

        if len(reqToken) > 0 {
            splitToken := strings.Split(reqToken, "Bearer ")
            reqToken = splitToken[1]

            claims, isValid := lib.GetTokenClaims(reqToken)

            if !isValid {
                lib.SendJsonError(w, "invalid token", http.StatusUnauthorized)
                return
            }

            if claims != nil {
                database.Db.Where("id=? and status=?", claims.Uid, models.Active).First(&user)

                if user.ID != 0 {
                    ctx := context.WithValue(r.Context(), "user", &user)

                    r = r.WithContext(ctx)
                }
            }
        }

        next.ServeHTTP(w, r)
    })
}
