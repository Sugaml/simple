package models

import (
	"errors"
	"time"

	"github.com/jinzhu/gorm"
)

type PromoCode struct {
	gorm.Model
	Title      string    `gorm:"not null" json:"title"`
	Code       uint      `gorm:"not null" json:"code"`
	IsPercent  bool      `gorm:"not null" json:"is_percent"`
	Discount   uint      `gorm:"not null" json:"discount"`
	ExpiryDate time.Time `gorm:"not null" json:"expiry_date"`
	Limit      uint      `gorm:"not null" json:"limit"`
	Count      uint      `gorm:"not null" json:"count"`
	Active     bool      `gorm:"not null" json:"active"`
}

func (d *DBStruct) CreatePromocode(data *PromoCode) (*PromoCode, error) {
	err = d.db.Model(&PromoCode{}).Create(&data).Error
	if err != nil {
		return &PromoCode{}, err
	}
	return data, nil
}
func (d *DBStruct) FindAllPromocode(datas *[]PromoCode) (*[]PromoCode, error) {
	err = d.db.Model(&PromoCode{}).Order("id desc").Find(&datas).Error
	if err != nil {
		return &[]PromoCode{}, err
	}
	return datas, nil
}

func (d *DBStruct) FindByIdPromocode(pid uint) (*PromoCode, error) {
	data := PromoCode{}
	err = d.db.Model(&PromoCode{}).Where("id = ?", pid).Take(&data).Error
	if err != nil {
		return &PromoCode{}, err
	}
	return &data, nil
}

func (d *DBStruct) UpdatePromocode(data *PromoCode) (*PromoCode, error) {
	err = d.db.Model(&PromoCode{}).Update(&data).Error
	if err != nil {
		return &PromoCode{}, err
	}
	return data, nil
}

func (d *DBStruct) DeletePromocode(pid uint) (int64, error) {
	result := d.db.Model(&PromoCode{}).Where("id = ?", pid).Take(&PromoCode{}).Delete(&PromoCode{})
	if result.Error != nil {
		if gorm.IsRecordNotFoundError(result.Error) {
			return 0, errors.New("promocode not found")
		}
		return 0, result.Error
	}
	return result.RowsAffected, nil
}
