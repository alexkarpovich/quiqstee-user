package root

import (
  "net/http"
  "encoding/json"
  "github.com/alexkarpovich/quiqstee-user/lib"
  "github.com/alexkarpovich/quiqstee-user/database"
  "github.com/alexkarpovich/quiqstee-user/database/models"
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

  user := models.User{
    Email: sus.Email,
    Password: sus.Password,
    FirstName: sus.FirstName,
    LastName: sus.LastName,
    Role: models.Member,
    Status: models.Inactive,
  }
  database.Db.Create(&user)
  go email.SendSignup(&user)

  lib.SendJson(w, "Success", http.StatusOK)
}

func (h *RootHandler) Login(w http.ResponseWriter, r *http.Request) {
  var lis structs.Login
  var user models.User
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

  database.Db.Where(&models.User{Email: lis.Email}).First(&user)
  lib.SendJson(w, map[string]string{"token": lib.NewToken(&user)}, http.StatusOK)
}
