package root

import (
    "net/http"
    "encoding/json"
    "github.com/gorilla/mux"
    "github.com/alexkarpovich/quiqstee-user/lib"
    "github.com/alexkarpovich/quiqstee-user/database"
    "github.com/alexkarpovich/quiqstee-user/database/users"
    "github.com/alexkarpovich/quiqstee-user/database/regs"
    "github.com/alexkarpovich/quiqstee-user/requests/structs"
    "github.com/alexkarpovich/quiqstee-user/service/email"
)

func (h *RootHandler) Signup(w http.ResponseWriter, r *http.Request) {
    var s structs.Signup
    err := json.NewDecoder(r.Body).Decode(&s)
    defer r.Body.Close()
    if err != nil {
        lib.SendJsonError(w, "Invalid request data", http.StatusBadRequest)
        return
    }

    if !s.Validate() {
        lib.SendJsonError(w, s.Errors, http.StatusBadRequest)
        return
    }

    reg := regs.Registration{
        Email: s.Email,
    }
    database.Db.Create(&reg)
    go email.SendSignup(&reg)

    lib.SendJson(w, "Success", http.StatusOK)
}

func (h *RootHandler) ConfirmSignup(w http.ResponseWriter, r *http.Request) {
    var s structs.ConfirmSignup
    err := json.NewDecoder(r.Body).Decode(&s)
    defer r.Body.Close()
    if err != nil {
        lib.SendJsonError(w, "Invalid request data", http.StatusBadRequest)
        return
    }

    vars := mux.Vars(r)
    s.Token = vars["token"]

    if !s.Validate() {
        lib.SendJsonError(w, s.Errors, http.StatusBadRequest)
        return
    }

    user := users.User{
        Email: s.Email,
        FirstName: s.FirstName,
        LastName: s.LastName,
        Role: users.Member,
        Status: users.Active,
    }
    user.SetPassword(s.Password)
    database.Db.Create(&user)

    lib.SendJson(w, "Success", http.StatusOK)
}

func (h *RootHandler) Login(w http.ResponseWriter, r *http.Request) {
    var s structs.Login
    
    err := json.NewDecoder(r.Body).Decode(&s)
    defer r.Body.Close()
    if err != nil {
        lib.SendJsonError(w, "Invalid request data.", http.StatusBadRequest)
        return
    }

    if !s.Validate() {
        lib.SendJsonError(w, s.Errors, http.StatusBadRequest)
        return
    }

    lib.SendJson(w, map[string]string{"token": lib.NewToken(s.User)}, http.StatusOK)
}
