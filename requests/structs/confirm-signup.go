package structs

import (
    "time"
    "github.com/alexkarpovich/quiqstee-user/database"
    "github.com/alexkarpovich/quiqstee-user/database/regs"
)

type ConfirmSignup struct {
    Email string
    Password string
    FirstName string
    LastName string
    Token string
}

func (s *ConfirmSignup) Validate() bool {
    var reg regs.Registration

    database.Db.Where("token=? and expires_at > ?", s.Token, time.Now()).First(&reg)

    s.Email = reg.Email

    return reg.ID != 0
}
