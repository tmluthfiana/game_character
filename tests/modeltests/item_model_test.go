package modeltests

import (
	"log"
	"testing"

	"game_character/api/models"

	_ "github.com/jinzhu/gorm/dialects/mysql"
	"gopkg.in/go-playground/assert.v1"
)

func TestFindAllItems(t *testing.T) {

	err := refreshItemTable()
	if err != nil {
		log.Fatalf("Error refreshing item table %v\n", err)
	}
	_, _, err = seedItems()
	if err != nil {
		log.Fatalf("Error seeding item table %v\n", err)
	}
	items, err := itemInstance.FindAllItems(server.DB)
	if err != nil {
		t.Errorf("this is the error getting the items: %v\n", err)
		return
	}
	assert.Equal(t, len(*items), 2)
}

func TestSaveItem(t *testing.T) {

	err := refreshItemTable()
	if err != nil {
		log.Fatalf("Error item refreshing table %v\n", err)
	}

	newItem := models.Item{
		Name:   "Gollum",
		Character_code: 3,
		Power: 30,
	}
	savedItem, err := newItem.SaveItem(server.DB)
	if err != nil {
		t.Errorf("this is the error getting the item: %v\n", err)
		return
	}
	assert.Equal(t, newItem.ID, savedItem.ID)
	assert.Equal(t, newItem.Name, savedItem.Name)
	assert.Equal(t, newItem.Character_code, savedItem.Character_code)
	assert.Equal(t, newItem.Power, savedItem.Power)

}

func TestGetItemByID(t *testing.T) {

	err := refreshItemTable()
	if err != nil {
		log.Fatalf("Error refreshing item table: %v\n", err)
	}
	item, err := seedOneItem()
	if err != nil {
		log.Fatalf("Error Seeding table")
	}
	foundItem, err := itemInstance.FindItemByID(server.DB, item.ID)
	if err != nil {
		t.Errorf("this is the error getting one user: %v\n", err)
		return
	}
	assert.Equal(t, foundItem.ID, item.ID)
	assert.Equal(t, foundItem.Name, item.Name)
	assert.Equal(t, foundItem.Character_code, item.Character_code)
}

func TestUpdateItem(t *testing.T) {

	err := refreshUserAndItemTable()
	if err != nil {
		log.Fatalf("Error refreshing user and item table: %v\n", err)
	}
	item, err := seedOneItem()
	if err != nil {
		log.Fatalf("Error Seeding table")
	}
	itemUpdate := models.Item{
		Name:   "Bilbooo",
		Character_code: 3,
		Power: 15,
	}
	updatedItem, err := itemUpdate.UpdateAItem(server.DB)
	if err != nil {
		t.Errorf("this is the error updating the user: %v\n", err)
		return
	}
	assert.Equal(t, updatedItem.ID, itemUpdate.ID)
	assert.Equal(t, updatedItem.Name, itemUpdate.Name)
	assert.Equal(t, updatedItem.Character_code, itemUpdate.Character_code)
	assert.Equal(t, updatedItem.Power, itemUpdate.Power)
}

