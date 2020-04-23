package models

import (
	"errors"
	"github.com/badoux/checkmail"
	"github.com/dgrijalva/jwt-go"
	"github.com/go-redis/redis"
	"github.com/jinzhu/gorm"
	"html"
	"strings"
	"time"
	"woden/src/database"
)

type User struct {
	ID       uint32 `gorm:"primary_key;auto_increment" json:"id"`
	Username string `gorm:"size:100;not null;unique" json:"username"`
	Email    string `gorm:"size:100;not null;unique" json:"email"`
	Password string `gorm:"size:100;not null;" json:"password"`
	Token    string `json:"token"`
}

type Token struct {
	Token  string
	UserId string
	jwt.StandardClaims
}

func (u *User) Prepare() {
	u.ID = 0
	u.Username = html.EscapeString(strings.TrimSpace(u.Username))
	u.Email = html.EscapeString(strings.TrimSpace(u.Email))
}

func (u *User) Validate(action string) error {
	switch strings.ToLower(action) {
	case "update":
		if u.Password == "" {
			return errors.New("Required Password")
		}
		return nil

	case "login":
		if u.Password == "" {
			return errors.New("Required Password")
		}
		if u.Email == "" && u.Username == "" {
			return errors.New("Required Email or Username")
		}
		if u.Username == "" {
			if err := checkmail.ValidateFormat(u.Email); err != nil {
				return errors.New("Invalid Email")
			}
		}
		return nil

	default:
		if u.Username == "" {
			return errors.New("Required Username")
		}
		if u.Password == "" {
			return errors.New("Required Password")
		}
		if u.Email == "" {
			return errors.New("Required Email")
		}
		if err := checkmail.ValidateFormat(u.Email); err != nil {
			return errors.New("Invalid Email")
		}
		return nil
	}
}

func (u *User) SaveUser(db *gorm.DB) (*User, error) {
	var err error
	err = db.Debug().Create(&u).Error
	if err != nil {
		return &User{}, err
	}
	return u, nil
}

func (u *User) UpdateAUser(db *gorm.DB, id uint32) (*User, error) {
	db = db.Debug().Model(&User{}).Where("id = ?", id).Take(&User{}).UpdateColumns(
		map[string]interface{}{
			"password": u.Password,
		},
	)
	if db.Error != nil {
		return &User{}, db.Error
	}
	// This is the display the updated user
	err := db.Debug().Model(&User{}).Where("id = ?", id).Take(&u).Error
	if err != nil {
		return &User{}, err
	}
	return u, nil
}

func Logout(id uint32, token string) (bool, error) {
	byToken, err := GetToken(token)
	if byToken == nil {
		if err := database.ClientForToken.Set(token, id, time.Hour).Err(); err != nil {
			return false, err
		}
		return true, nil
	}
	return false, err
}

func GetToken(token string) (*Token, error) {
	name, err := database.ClientForToken.Get(token).Result()
	if err == redis.Nil {
		return nil, nil
	}
	if err != nil {
		return nil, errors.New("While getting token")
	}
	tokenFromDB := &Token{}
	tokenFromDB.Token = token
	tokenFromDB.UserId = name
	return tokenFromDB, nil
}
