package structs

import (
  "github.com/alexkarpovich/quiqstee-user/database"
  "github.com/alexkarpovich/quiqstee-user/database/models"
)

type Signup struct {
  Email string `json:"email"`
  Password string `json:"password"`
  FirstName string `json:"firstName"`
  LastName string `json:"lastName"`
}

func (s *Signup) Validate() bool {
  var user models.User

  database.Db.Where("email=?", s.Email).First(&user)

  return user.ID == 0
}
