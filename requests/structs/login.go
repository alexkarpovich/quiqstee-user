package structs

import (
  "github.com/alexkarpovich/quiqstee-user/database"
  "github.com/alexkarpovich/quiqstee-user/database/models"
)

type Login struct {
  Email string `json:"email"`
  Password string `json:"password"`
}

func (l *Login) Validate() bool {
  var user models.User

  database.Db.Where("email=?", l.Email).First(&user)

  return user.ID != 0 && user.CheckPassword(l.Password)
}
