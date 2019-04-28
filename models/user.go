package models

import (
  "time"

  "github.com/youngjinpark20/gosu/modules/app"

  "golang.org/x/crypto/bcrypt"
  "github.com/lib/pq"
  "github.com/teris-io/shortid"
  "github.com/jinzhu/gorm"
)

// User stores various user data like authentication,
// name, maps, etc.
type User struct {
  gorm.Model
  ShortID string `gorm:"size:15;primary_key:true;unique;not_null"`
  Name string `gorm:"size:25"`
  Email string `gorm:"size:255;unique;not_null"`
  JoinDate time.Time
  PasswordHash []byte `gorm:"size:100;not_null"`

  FavoriteMaps pq.StringArray `gorm:"type:varchar(15)[]"`
}

// Init Initializes the User with a ShortID and returns user.
func (u *User) Init(ctx *app.Context) error {
  var err error
  u.ShortID, err = shortid.Generate()
  return err
}

// SetPassword sets password for a user given password `p` by taking in the
// string, generating a salt and a hash for authentication.
func (u *User) SetPassword(password string) error {
  hash, err := bcrypt.GenerateFromPassword([]byte(password), bcrypt.DefaultCost)
  if err != nil {
    return err
  }
  u.PasswordHash = hash
  return nil
}

// Authenticate Password returns a boolean value depending on if password
// string argument is the right password for the stored hash
func (u *User) Authenticate(password string) bool {
  err := bcrypt.CompareHashAndPassword(u.PasswordHash, []byte(password))
  if err == nil {
    return true
  }
  return false
}
