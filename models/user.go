package models

import (
  "github.com/jinzhu/gorm"
)

type User struct {
  gorm.Model
  Email string
  Phone string
  FirstName string
  LastName string
}
