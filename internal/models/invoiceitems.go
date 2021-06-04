package models

import (
	"github.com/jinzhu/gorm"
)

type InvoiceItems struct {
	gorm.Model
	Invoice    uint   `gorm:"foreignkey:InvoiceID" json:"invoice"`
	InvoiceID  uint   `gorm:"not null" json:"invoice_id"`
	User       *User  `gorm:"foreignkey:UserID" json:"user"`
	UserID     uint   `gorm:"not null" json:"user_id"`
	Particular string `gorm:"not null" json:"particular"`
	Rate       uint   `gorm:"not null" json:"rate"`
	Days       uint   `gorm:"not null" json:"days"`
	Total      uint   `gorm:"not null" json:"total"`
}

func (data *InvoiceItems) Save(db *gorm.DB) (*InvoiceItems, error) {
	err = db.Model(&InvoiceItems{}).Create(&data).Error
	if err != nil {
		return &InvoiceItems{}, err
	}
	return data, nil
}

func (data *InvoiceItems) FindAll(db *gorm.DB) (*[]InvoiceItems, error) {
	datas := []InvoiceItems{}
	err = db.Model(&InvoiceItems{}).Order("id desc").Find(&datas).Error
	if err != nil {
		return &[]InvoiceItems{}, err
	}
	return &datas, nil
}

func (data *InvoiceItems) Find(db *gorm.DB, pid uint64) (*InvoiceItems, error) {
	err = db.Model(&InvoiceItems{}).Where("id = ?", pid).Take(&data).Error
	if err != nil {
		return &InvoiceItems{}, err
	}
	return data, nil
}

func (data *InvoiceItems) Update(db *gorm.DB) (*InvoiceItems, error) {
	var invoiceItems = InvoiceItems{}
	if data.InvoiceID != 0 {
		invoiceItems.InvoiceID = data.InvoiceID
	}
	if data.UserID != 0 {
		invoiceItems.UserID = data.UserID
	}
	if data.Particular != "" {
		invoiceItems.Particular = data.Particular
	}
	if data.Rate != 0 {
		invoiceItems.Rate = data.Rate
	}
	if data.Days != 0 {
		invoiceItems.Days = data.Days
	}
	if data.Total != 0 {
		invoiceItems.Total = data.Total
	}
	err = db.Model(&InvoiceItems{}).Where("id=?", data.ID).Updates(invoiceItems).Error
	if err != nil {
		return &InvoiceItems{}, err
	}
	return data, nil
}

func (data *InvoiceItems) Delete(db *gorm.DB) (*InvoiceItems, error) {
	err = db.Model(&InvoiceItems{}).Delete(&data).Error
	if err != nil {
		return &InvoiceItems{}, err
	}
	return data, nil
}
