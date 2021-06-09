package models

import (
	"errors"
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

func (d *DBStruct) CreateDeduction(data Deduction) (Deduction, error) {
	err = d.db.Model(&Deduction{}).Create(&data).Error
	if err != nil {
		return Deduction{}, err
	}
	return data, nil
}

func (d *DBStruct) FindAllDeduction() ([]Deduction, error) {
	datas := []Deduction{}
	err = d.db.Model(&Deduction{}).Order("id desc").Find(&datas).Error
	if err != nil {
		return []Deduction{}, err
	}
	return datas, nil
}

func (d *DBStruct) FindByIdDeduction(pid uint) (Deduction, error) {
	data := Deduction{}
	err = d.db.Model(&Deduction{}).Where("id = ?", pid).Take(&data).Error
	if err != nil {
		return Deduction{}, err
	}
	return data, nil
}

func (d *DBStruct) FindByCountryDeduction(country string) (Deduction, error) {
	data := Deduction{}
	err = d.db.Model(&Deduction{}).Where("country= ?", country).Take(&data).Error
	if err != nil {
		return Deduction{}, err
	}
	return data, nil
}

func (d *DBStruct) UpdateDeduction(data Deduction) (Deduction, error) {

	err = d.db.Model(&Deduction{}).Update(&data).Error
	if err != nil {
		return Deduction{}, err
	}
	return data, nil
}

func (d *DBStruct) DeleteDeduction(pid uint) (int64, error) {
	result := d.db.Model(&Deduction{}).Where("id = ?", pid).Take(&Deduction{}).Delete(&Deduction{})
	if result.Error != nil {
		if gorm.IsRecordNotFoundError(result.Error) {
			return 0, errors.New("Deduction  not found")
		}
		return 0, result.Error
	}
	return result.RowsAffected, nil
}
