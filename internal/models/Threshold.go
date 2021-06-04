package models

import (
	"errors"

	"github.com/jinzhu/gorm"
)

type Threshold struct {
	gorm.Model
	User           uint   `gorm:"foreignkey:UserID" json:"user"`
	UserID         uint   `gorm:"not null" json:"user_id"`
	ThresholdLimit uint   `gorm:"not null" json:"threshold_limit"`
	Email          string `gorm:"not null" json:"email"`
	Active         bool   `gorm:"not null" json:"active"`
}

func (data *Threshold) Save(db *gorm.DB) (*Threshold, error) {
	err = db.Model(&Threshold{}).Create(&data).Error
	if err != nil {
		return &Threshold{}, err
	}
	return data, nil
}

func (data *Threshold) FindAll(db *gorm.DB) (*[]Threshold, error) {
	datas := []Threshold{}
	err = db.Model(&Threshold{}).Order("id desc").Find(&datas).Error
	if err != nil {
		return &[]Threshold{}, err
	}
	return &datas, nil
}

func (data *Threshold) Find(db *gorm.DB, pid uint64) (*Threshold, error) {
	err = db.Model(&Threshold{}).Where("id = ?", pid).Take(&data).Error
	if err != nil {
		return &Threshold{}, err
	}
	return data, nil
}

func (data *Threshold) FindByUserID(db *gorm.DB, uid uint64) (*Threshold, error) {
	err = db.Model(&Threshold{}).Order("id desc").Where("user_id = ?", uid).Take(&data).Error
	if err != nil {
		return &Threshold{}, err
	}
	return data, nil
}

func (data *Threshold) Update(db *gorm.DB) (*Threshold, error) {
	var threshold = Threshold{}
	if data.UserID != 0 {
		threshold.UserID = data.UserID
	}
	if data.ThresholdLimit != 0 {
		threshold.ThresholdLimit = data.ThresholdLimit
	}
	if data.Email != "" {
		threshold.Email = data.Email
	}
	err = db.Model(&Threshold{}).Where("id = ?", data.ID).Updates(threshold).Error
	if err != nil {
		return &Threshold{}, err
	}
	return data, nil
}

func (data *Threshold) Delete(db *gorm.DB, id uint64) (int64, error) {
	db = db.Model(&Threshold{}).Where("id = ?", id).Take(&Threshold{}).Delete(&Threshold{})
	if db.Error != nil {
		if gorm.IsRecordNotFoundError(db.Error) {
			return 0, errors.New("Threshold not found")
		}
		return 0, db.Error
	}
	return db.RowsAffected, nil
}
