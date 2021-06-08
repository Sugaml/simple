package models

import (
	"errors"

	"github.com/jinzhu/gorm"
)

type Gateway struct {
	gorm.Model
	Name      string `gorm:"not null" json:"name"`
	AccessKey string `gorm:"not null" json:"accesskey"`
	SecretKey string `gorm:"not null" json:"secretkey"`
	Token     string `gorm:"not null" json:"token"`
	Others    string `gorm:"not null" json:"others"`
	Url       string `gorm:"not null" json:"url"`
	Active    bool   `gorm:"not null" json:"active"`
}

var err error

func (d *DBStruct) CreateGateway(data Gateway) (Gateway, error) {
	err = d.db.Model(&Gateway{}).Create(&data).Error
	if err != nil {
		return Gateway{}, err
	}
	return data, nil
}
func (d *DBStruct) FindAllGateway() ([]Gateway, error) {
	datas := []Gateway{}
	err = d.db.Model(&Gateway{}).Order("id desc").Find(&datas).Error
	if err != nil {
		return []Gateway{}, err
	}
	return datas, nil
}

func (d *DBStruct) FindByIdGateway(pid uint) (Gateway, error) {
	data := Gateway{}
	err = d.db.Model(&Gateway{}).Where("id = ?", pid).Take(&data).Error
	if err != nil {
		return Gateway{}, err
	}
	return data, nil
}

func (d *DBStruct) UpdateGateway(data Gateway) (Gateway, error) {
	err = d.db.Model(&Gateway{}).Update(&data).Error
	if err != nil {
		return Gateway{}, err
	}
	return data, nil
}

func (d *DBStruct) DeleteGateway(pid uint) (int64, error) {
	result := d.db.Model(&Gateway{}).Where("id = ?", pid).Take(&Gateway{}).Delete(&Gateway{})
	if result.Error != nil {
		if gorm.IsRecordNotFoundError(result.Error) {
			return 0, errors.New("promocode not found")
		}
		return 0, result.Error
	}
	return result.RowsAffected, nil
}
