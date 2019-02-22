package structs

import (
  "github.com/alexkarpovich/quiqstee-user/database"
  "github.com/alexkarpovich/quiqstee-user/database/regs"
  "github.com/alexkarpovich/quiqstee-user/database/users"
)

type Signup struct {
  Email string `json:"email"`
}

func (s *Signup) Validate() bool {
  var user users.User
  var reg regs.Registration

  database.Db.Where("email=?", s.Email).First(&user)
  database.Db.Where("email=?", s.Email).First(&reg)

  return user.ID == 0 && reg.ID == 0
}
