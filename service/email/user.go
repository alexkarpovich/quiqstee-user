package email

import (
    "github.com/alexkarpovich/quiqstee-user/database/models"
)

func SendSignup(user *models.User) {
    subject := "Подтверждение регистрации"
    from := "alexsure.k@gmail.com"

    SendWithView(
        subject,
        from,
        []string{user.Email},
        []string{
            "./email/templates/layout/base.html",
            "./email/templates/auth/signup.html",
        },
        "layout",
        user,
    )
}
