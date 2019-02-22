package email

import (
    "github.com/alexkarpovich/quiqstee-user/database/regs"
)

func SendSignup(reg *regs.Registration) {
    subject := "Подтверждение регистрации"
    from := "alexsure.k@gmail.com"

    SendWithView(
        subject,
        from,
        []string{reg.Email},
        []string{
            "./email/templates/layout/base.html",
            "./email/templates/auth/signup.html",
        },
        "layout",
        reg,
    )
}
