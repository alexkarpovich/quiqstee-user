package structs

import (
  "log"
  "github.com/alexkarpovich/quiqstee-user/database"
  "github.com/alexkarpovich/quiqstee-user/database/models"
)

type Signup struct {
  Email string `json:"email"`
  Password string `json:"password"`
}

func (s *Signup) Validate() bool {
  var user models.User

  database.Db.Where("email=?", s.Email).First(&user)

  log.Printf("%s", user)

  return user.ID == 0
}
