package structs

import (
    "github.com/alexkarpovich/quiqstee-user/database"
    "github.com/alexkarpovich/quiqstee-user/database/users"
)

type Login struct {
    Base
    Email string `json:"email"`
    Password string `json:"password"`
    User *users.User
}

func (s *Login) Validate() bool {
    var user users.User
    s.Errors = make(map[string]string)

    database.Db.Where("email=? and status=?", s.Email, users.Active).First(&user)

    if user.ID == 0 {
        s.Errors["email"] = "There is no active user with such email." 
    } else if !user.CheckPassword(s.Password) {
        s.Errors["password"] = "Password is incorrect."   
    }

    s.User = &user

    return len(s.Errors) == 0
}
