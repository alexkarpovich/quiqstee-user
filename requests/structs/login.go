package structs

import (
  "github.com/alexkarpovich/quiqstee-user/database"
  "github.com/alexkarpovich/quiqstee-user/database/users"
)

type Login struct {
  Email string `json:"email"`
  Password string `json:"password"`
}

func (l *Login) Validate() bool {
  var user users.User

  database.Db.Where("email=? and status=?", l.Email, users.Active).First(&user)

  return user.ID != 0 && user.CheckPassword(l.Password)
}
