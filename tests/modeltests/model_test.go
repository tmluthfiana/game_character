package modeltests

import (
	"fmt"
	"log"
	"os"
	"testing"

	"game_character/api/controllers"
	"game_character/api/models"

	"github.com/jinzhu/gorm"
	"github.com/joho/godotenv"
)

var server = controllers.Server{}
var itemInstance = models.Item{}

func TestMain(m *testing.M) {
	var err error
	err = godotenv.Load(os.ExpandEnv("../../.env"))
	if err != nil {
		log.Fatalf("Error getting env %v\n", err)
	}
	Database()

	os.Exit(m.Run())
}

func Database() {

	var err error

	TestDbDriver := os.Getenv("TestDbDriver")

	if TestDbDriver == "mysql" {
		DBURL := fmt.Sprintf("%s:%s@tcp(%s:%s)/%s?charset=utf8&parseTime=True&loc=Local", os.Getenv("TestDbUser"), os.Getenv("TestDbPassword"), os.Getenv("TestDbHost"), os.Getenv("TestDbPort"), os.Getenv("TestDbName"))
		server.DB, err = gorm.Open(TestDbDriver, DBURL)
		if err != nil {
			fmt.Printf("Cannot connect to %s database\n", TestDbDriver)
			log.Fatal("This is the error:", err)
		} else {
			fmt.Printf("We are connected to the %s database\n", TestDbDriver)
		}
	}
	if TestDbDriver == "postgres" {
		DBURL := fmt.Sprintf("host=%s port=%s user=%s dbname=%s sslmode=disable password=%s", os.Getenv("TestDbHost"), os.Getenv("TestDbPort"), os.Getenv("TestDbUser"), os.Getenv("TestDbName"), os.Getenv("TestDbPassword"))
		server.DB, err = gorm.Open(TestDbDriver, DBURL)
		if err != nil {
			fmt.Printf("Cannot connect to %s database\n", TestDbDriver)
			log.Fatal("This is the error:", err)
		} else {
			fmt.Printf("We are connected to the %s database\n", TestDbDriver)
		}
	}
}

func refreshItemTable() error {

	err := server.DB.DropTableIfExists(&models.Item{}).Error
	if err != nil {
		return err
	}
	err = server.DB.AutoMigrate(&models.Item{}).Error
	if err != nil {
		return err
	}
	log.Printf("Successfully refreshed tables")
	return nil
}

func seedOneItem() (models.Item, error) {

	err := refreshUserAndItemTable()
	if err != nil {
		return models.Item{}, err
	}
	
	item := models.Item{
		Name:   "Bilbo",
		Character_code: 3,
		Power: 20,
	},

	err = server.DB.Model(&models.Item{}).Create(&item).Error
	if err != nil {
		return models.Item{}, err
	}
	return item, nil
}

func seedItems() ([]models.Item, error) {

	var err error

	if err != nil {
		return []models.User{}, []models.Item{}, err
	}
	var items = []models.Item{
		models.Item{
			Name:   "Sauron",
			Character_code: 1,
			Power: 110,
		},
		models.Item{
			Name:   "Turiel",
			Character_code: 2,
			Power: 70,
		},
	}

	for i, _ := range items {

		err = server.DB.Model(&models.Item{}).Create(&items[i]).Error
		if err != nil {
			log.Fatalf("cannot seed items table: %v", err)
		}
	}
	return items, nil
}
