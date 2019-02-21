package email

import (
    "bytes"
    "html/template"
	"gopkg.in/gomail.v2"
)

func Send(subject string, from string, recipients []string, body string) {
    m := gomail.NewMessage()
    m.SetHeader("From", from)
    m.SetHeader("To", recipients...)
    m.SetHeader("Subject", subject)
    m.SetBody("text/html", body)

    d := gomail.NewDialer("smtp.gmail.com", 587, "alexsure.k", "Test123!")

    if err := d.DialAndSend(m); err != nil {
        panic(err)
    }
}

func SendWithView(subject string, from string, recipients []string, views []string, layout string, data interface{}) {
    t, err := template.ParseFiles(views...)
    if err != nil {
        panic(err)
    }

    var tpl bytes.Buffer
    if err := t.ExecuteTemplate(&tpl, layout, data); err != nil {
        panic(err)
    }

    Send(subject, from, recipients, tpl.String())
}
