package seed

import (
	"log"

	"game_character/api/models"

	"github.com/jinzhu/gorm"
)

var items = []models.Item{
	models.Item{
		Name:   "Gandalf",
		Character_code: 1,
		Power: 100,
	},
	models.Item{
		Name:   "Legolas",
		Character_code: 2,
		Power: 60,
	},
	models.Item{
		Name:   "Frodo",
		Character_code: 3,
		Power: 10,
	},
}

func Load(db *gorm.DB) {

	err := db.Debug().DropTableIfExists(&models.Item{}).Error
	if err != nil {
		log.Fatalf("cannot drop table: %v", err)
	}
	err = db.Debug().AutoMigrate(&models.Item{}).Error
	if err != nil {
		log.Fatalf("cannot migrate table: %v", err)
	}

}
