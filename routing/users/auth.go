package users

import (
  "net/http"
  "encoding/json"
  "github.com/alexkarpovich/quiqstee-user/lib"
  "github.com/alexkarpovich/quiqstee-user/database"
  "github.com/alexkarpovich/quiqstee-user/database/models"
  "github.com/alexkarpovich/quiqstee-user/requests/structs"
)

func (h *UserHandler) Signup(w http.ResponseWriter, r *http.Request) {
  var sus structs.Signup
  decoder := json.NewDecoder(r.Body)
  err := decoder.Decode(&sus)
  if err != nil {
      json.NewEncoder(w).Encode("Invalid request data")
      return
  }

  if !sus.Validate() {
    json.NewEncoder(w).Encode("User already exists")
    return
  }

  user := models.User{
    Email: sus.Email,
    Password: sus.Password,
    FirstName: sus.FirstName,
    LastName: sus.LastName,
  }
  database.Db.Create(&user)

  json.NewEncoder(w).Encode(map[string]string{"token": lib.NewToken(&user)})
}

func (h *UserHandler) Login(w http.ResponseWriter, r *http.Request) {

}

func (h *UserHandler) Logout(w http.ResponseWriter, r *http.Request) {

}
