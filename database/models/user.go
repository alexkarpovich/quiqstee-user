package models

import (
  "log"
  "github.com/jinzhu/gorm"
  "golang.org/x/crypto/bcrypt"
)

type User struct {
  gorm.Model
  Email string `json:"email"`
  Phone string `json:"phone"`
  Password string `gorm:"-"` // Ignore this field
  PasswordHash string `json:"-"` //omit passwordhash field
  FirstName string `json:"firstName"`
  LastName string `json:"lastName"`
}

func (user *User) BeforeCreate(scope *gorm.Scope) error {
  log.Printf("Hello")
  bytes, err := bcrypt.GenerateFromPassword([]byte(user.Password), 14)

  if err != nil {
    panic(err)
  }

  scope.SetColumn("password_hash", string(bytes))

  return nil
}

func (user *User) CheckPassword(password string) bool {
  err := bcrypt.CompareHashAndPassword([]byte(user.PasswordHash), []byte(password))

  return err == nil
}
