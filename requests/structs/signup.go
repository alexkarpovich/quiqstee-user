package structs

import (
  "github.com/alexkarpovich/quiqstee-user/database"
  "github.com/alexkarpovich/quiqstee-user/database/users"
)

type Signup struct {
  Email string `json:"email"`
}

func (s *Signup) Validate() bool {
  var user users.User

  database.Db.Where("email=?", s.Email).First(&user)

  return user.ID == 0
}
