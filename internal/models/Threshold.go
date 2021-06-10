package models

import (
	"errors"

	"github.com/jinzhu/gorm"
)

type PaymentThreshold struct {
	gorm.Model
	User           uint   `gorm:"foreignkey:UserID" json:"user"`
	UserID         uint   `gorm:"not null" json:"user_id"`
	ThresholdLimit uint   `gorm:"not null" json:"threshold_limit"`
	Email          string `gorm:"not null" json:"email"`
	Active         bool   `gorm:"default:false" json:"active"`
}

func (d *DBStruct) CreateThreshold(data PaymentThreshold) (PaymentThreshold, error) {
	dataExist, _ := d.FindByUserIDThreshold(data.UserID)
	if (PaymentThreshold{}) != dataExist {
		data.ID = dataExist.ID
		dataupdate, _ := d.UpdateThreshold(data)
		return dataupdate, nil
	}
	err = d.db.Model(&PaymentThreshold{}).Create(&data).Error
	if err != nil {
		return PaymentThreshold{}, err
	}
	return data, nil
}

func (d *DBStruct) FindAllThreshold() ([]PaymentThreshold, error) {
	datas := []PaymentThreshold{}
	err = d.db.Model(&PaymentThreshold{}).Find(&datas).Error
	if err != nil {
		return []PaymentThreshold{}, err
	}
	return datas, nil
}

func (d *DBStruct) FindByIdThreshold(pid uint) (PaymentThreshold, error) {
	data := PaymentThreshold{}
	err = d.db.Model(&PaymentThreshold{}).Where("id = ?", pid).Take(&data).Error
	if err != nil {
		return PaymentThreshold{}, err
	}
	return data, nil
}

func (d *DBStruct) FindByUserIDThreshold(uid uint) (PaymentThreshold, error) {
	data := PaymentThreshold{}
	err = d.db.Model(&PaymentThreshold{}).Order("id desc").Where("user_id = ?", uid).Take(&data).Error
	if err != nil {
		return PaymentThreshold{}, err
	}
	return data, nil
}

func (d *DBStruct) UpdateThreshold(data PaymentThreshold) (PaymentThreshold, error) {
	threshold := map[string]interface{}{
		"active": data.Active,
	}
	if data.UserID != 0 {
		threshold["user_id"] = data.UserID
	}
	if data.ThresholdLimit != 0 {
		threshold["threshold_limit"] = data.ThresholdLimit
	}
	if data.Email != "" {
		threshold["email"] = data.Email
	}

	err = d.db.Model(&PaymentThreshold{}).Where("id = ?", data.ID).Updates(threshold).Error
	if err != nil {
		return PaymentThreshold{}, err
	}
	return data, nil
}

func (d *DBStruct) DeleteThreshold(id uint) (int64, error) {
	result := d.db.Model(&PaymentThreshold{}).Where("id = ?", id).Take(&PaymentThreshold{}).Delete(&PaymentThreshold{})
	if result.Error != nil {
		if gorm.IsRecordNotFoundError(result.Error) {
			return 0, errors.New("threshold not found")
		}
		return 0, result.Error
	}
	return result.RowsAffected, nil
}
