package regs

import (
	"log"
	"time"
	"github.com/google/uuid"
	"github.com/jinzhu/gorm"
)

type Registration struct {
	ID uint `gorm:"primary_key"`
	Email string 
	Token string
	ExpiresAt time.Time `gorm:"type:time"`
}

func (reg *Registration) BeforeCreate(scope *gorm.Scope) error {
 
	log.Printf("%s", time.Now().Add(time.Hour * 1).Unix())
	scope.SetColumn("token", uuid.New().String())
	scope.SetColumn("expires_at", time.Now().Add(time.Hour * 1).UTC())
  
	return nil
  }