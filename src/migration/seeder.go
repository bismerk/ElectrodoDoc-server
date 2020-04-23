package seed

import (
	"log"
	"woden/src/models"

	"github.com/jinzhu/gorm"
)

var users = []models.User{
	{
		Username: "user1",
		Email:    "user1@gmail.com",
		Password: "password",
	},
	{
		Username: "user2",
		Email:    "user2@gmail.com",
		Password: "password",
	},
}

func Load(db *gorm.DB) {

	err := db.Debug().DropTableIfExists(&models.User{}).Error
	if err != nil {
		log.Fatalf("cannot drop table: %v", err)
	}
	err = db.Debug().AutoMigrate(&models.User{}).Error
	if err != nil {
		log.Fatalf("cannot migrate table: %v", err)
	}

	for i, _ := range users {
		err = db.Debug().Model(&models.User{}).Create(&users[i]).Error
		if err != nil {
			log.Fatalf("cannot seed users table: %v", err)
		}
	}
}
