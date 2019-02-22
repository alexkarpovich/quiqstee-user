package middlewares

import (
    "log"
    "net/http"
    "github.com/casbin/casbin"
    "github.com/alexkarpovich/quiqstee-user/lib"
    "github.com/alexkarpovich/quiqstee-user/database/users"
)

func Authorizer(e *casbin.Enforcer) func(next http.Handler) http.Handler {
    return func(next http.Handler) http.Handler {
        fn := func(w http.ResponseWriter, r *http.Request) {
            role := users.Anonymous
            user, ok := r.Context().Value("user").(*users.User)
            if ok {
                role = user.Role
            }

            log.Printf("%s, %s, %s", role, r.URL.Path, r.Method)

            res, err := e.EnforceSafe(string(role), r.URL.Path, r.Method)

            if err != nil {
                lib.SendJsonError(w, err, http.StatusInternalServerError)
                return
            }

            if res {
                next.ServeHTTP(w, r)
            } else {
                lib.SendJsonError(w, "Forbidden", http.StatusForbidden)
                return
            }
        }

        return http.HandlerFunc(fn)
    }
}
