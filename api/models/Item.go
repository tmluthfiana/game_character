package models

import (
	"errors"
	"html"
	"strings"
	"time"

	"github.com/jinzhu/gorm"
)

type Item struct {
	ID        uint64    `gorm:"primary_key;auto_increment" json:"id"`
	Name     string    `gorm:"size:255;not null;unique" json:"name"`
	Character_code  uint32    `gorm:"not null" json:"character_code"`
	Power  uint32    `gorm:"null" json:"power"`
	Value  uint32    `gorm:"null" json:"value"`
}

func (p *Item) Prepare() {
	p.ID = 0
	p.Name = html.EscapeString(strings.TrimSpace(p.Name))
	p.Character_code = 0
	p.Power = 0
	p.Value = 0
}

func (p *Item) Validate() error {

	if p.Name == "" {
		return errors.New("Required Name")
	}

	if p.Character_code >= 4 && p.Character_code <= 0 {
		return errors.New("Wrong Character code")
	}
	
	return nil
}

func (p *Item) SaveItem(db *gorm.DB) (*Item, error) {
	var err error

	GetValue(p)

	err = db.Debug().Model(&Item{}).Create(&p).Error
	if err != nil {
		return &Item{}, err
	}
	
	return p, nil
}

func (p *Item) FindAllItems(db *gorm.DB) (*[]Item, error) {
	var err error
	Items := []Item{}
	err = db.Debug().Model(&Item{}).Limit(100).Find(&Items).Error
	if err != nil {
		return &[]Item{}, err
	}
	
	return &Items, nil
}

func (p *Item) FindItemByID(db *gorm.DB, pid uint64) (*Item, error) {
	var err error
	err = db.Debug().Model(&Item{}).Where("id = ?", pid).Take(&p).Error
	if err != nil {
		return &Item{}, err
	}
	
	return p, nil
}

func (p *Item) UpdateAItem(db *gorm.DB) (*Item, error) {

	var err error
	GetValue(p)

	err = db.Debug().Model(&Item{}).Where("id = ?", p.ID).Updates(Item{Name: p.Name, Power: p.Power}).Error
	if err != nil {
		return &Item{}, err
	}
	
	return p, nil
}

func (p *Item) GetValue()  (*Item) {
	if p.Character_code == 1 {
		p.Value = 150 * p.Power / 100
	} else if  p.Character_code == 2 {
		p.Value = 2 + (110 * p.Power / 100)
	} else if  p.Character_code == 3 {
		if p.Power < 20 {
			p.Value = 200 * p.Power / 100
		} else {
			p.Value = 300 * p.Power / 100
		}
	}

	return p
}
