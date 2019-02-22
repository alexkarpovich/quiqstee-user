package users

import (
    "fmt"
    "time"
    "golang.org/x/crypto/bcrypt"
)

type UserRole string
type UserStatus uint

const (
    Anonymous UserRole = "anonymous"
    Admin UserRole = "admin"
    Member UserRole = "member"
)

const (
    Inactive UserStatus = iota
    Active
    Deleted
)

type User struct {
    ID uint `json:"id";gorm:"primary_key"`
    CreatedAt time.Time `json:"createdAt"`
    UpdatedAt time.Time `json:"updatedAt"`
    DeletedAt *time.Time `json:"deletedAt"`
    Email string `json:"email"`
    Phone string `json:"phone"`
    PasswordHash string `json:"-"` //omit passwordhash field
    FirstName string `json:"firstName"`
    LastName string `json:"lastName"`
    Role UserRole `json:"role";gorm:"default:'member'"`
    Status UserStatus `json:"status"`
}

func (user *User) SetPassword(plainPassword string) {
    bytes, _ := bcrypt.GenerateFromPassword([]byte(plainPassword), 14) 
    user.PasswordHash = string(bytes)
}

func (user *User) CheckPassword(password string) bool {
    err := bcrypt.CompareHashAndPassword([]byte(user.PasswordHash), []byte(password))

    return err == nil
}

func (user *User) FullName() string {
    return fmt.Sprintf("%s %s", user.FirstName, user.LastName)
}
