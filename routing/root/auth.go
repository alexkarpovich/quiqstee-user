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
    var sus structs.Signup
    err := json.NewDecoder(r.Body).Decode(&sus)
    defer r.Body.Close()
    if err != nil {
        lib.SendJsonError(w, "Invalid request data", http.StatusBadRequest)
        return
    }

    if !sus.Validate() {
        lib.SendJsonError(w, "User already exists", http.StatusBadRequest)
        return
    }

    reg := regs.Registration{
        Email: sus.Email,
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
        lib.SendJsonError(w, "Invalid data", http.StatusBadRequest)
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
    var lis structs.Login
    var user users.User
    err := json.NewDecoder(r.Body).Decode(&lis)
    defer r.Body.Close()
    if err != nil {
        lib.SendJsonError(w, "Invalid request data.", http.StatusBadRequest)
        return
    }

    if !lis.Validate() {
        lib.SendJsonError(w, "Email or password is wrong.", http.StatusBadRequest)
        return
    }

    database.Db.Where(&users.User{Email: lis.Email}).First(&user)
    lib.SendJson(w, map[string]string{"token": lib.NewToken(&user)}, http.StatusOK)
}
