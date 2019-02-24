package structs

import (
    "time"
    "github.com/alexkarpovich/quiqstee-user/database"
    "github.com/alexkarpovich/quiqstee-user/database/regs"
)

type ConfirmSignup struct {
    Base
    Email string
    Password string
    FirstName string
    LastName string
    Token string
}

func (s *ConfirmSignup) Validate() bool {
    var reg regs.Registration
    s.Errors = make(map[string]string)

    database.Db.Where("token=? and expires_at > ?", s.Token, time.Now()).First(&reg)

    if reg.ID == 0 {
        s.Errors["token"] = "Token is invalid or expired."
    }

    s.Email = reg.Email

    return len(s.Errors) == 0
}
