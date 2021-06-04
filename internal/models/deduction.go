package models

import (
	"time"

	"github.com/jinzhu/gorm"
)

type Deduction struct {
	gorm.Model
	Name          string    `gorm:"not null" json:"name"`
	Value         uint      `gorm:"not null" json:"value"`
	IsPercent     bool      `gorm:"not null" json:"is_percent"`
	Country       string    `gorm:"not null" json:"country"`
	Description   string    `gorm:"not null" json:"description"`
	Attributes    string    `gorm:"not null" json:"attributes"`
	EffectiveDate time.Time `gorm:"not null" json:"date"`
}

func (data *Deduction) Save(db *gorm.DB) (*Deduction, error) {
	err = db.Model(&Deduction{}).Create(&data).Error
	if err != nil {
		return &Deduction{}, err
	}
	return data, nil
}

func (data *Deduction) FindAll(db *gorm.DB) (*[]Deduction, error) {
	datas := []Deduction{}
	err = db.Model(&Deduction{}).Order("id desc").Find(&datas).Error
	if err != nil {
		return &[]Deduction{}, err
	}
	return &datas, nil
}

func (data *Deduction) Find(db *gorm.DB, pid uint64) (*Deduction, error) {
	err = db.Model(&Deduction{}).Where("id = ?", pid).Take(&data).Error
	if err != nil {
		return &Deduction{}, err
	}
	return data, nil
}

func (data *Deduction) FindByCountry(db *gorm.DB, country string) (*Deduction, error) {
	err = db.Model(&Deduction{}).Where("country= ?", country).Take(&data).Error
	if err != nil {
		return &Deduction{}, err
	}
	return data, nil
}

func (data *Deduction) Update(db *gorm.DB) (*Deduction, error) {
	err = db.Model(&Deduction{}).Update(&data).Error
	if err != nil {
		return &Deduction{}, err
	}
	return data, nil
}

func (data *Deduction) Delete(db *gorm.DB) (*Deduction, error) {
	err = db.Model(&Deduction{}).Delete(&data).Error
	if err != nil {
		return &Deduction{}, err
	}
	return data, nil
}
