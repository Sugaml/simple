package models

import (
	"errors"

	"github.com/jinzhu/gorm"
)

type InvoiceItems struct {
	gorm.Model
	Invoice    uint    `gorm:"foreignkey:InvoiceID" json:"invoice"`
	InvoiceID  uint    `gorm:"not null" json:"invoice_id"`
	User       *User   `gorm:"foreignkey:UserID" json:"user"`
	UserID     uint    `gorm:"not null" json:"user_id"`
	Particular string  `gorm:"not null" json:"particular"`
	Rate       uint    `gorm:"not null" json:"rate"`
	Days       uint    `gorm:"not null" json:"days"`
	Total      float64 `gorm:"not null" json:"total"`
}

func (d *DBStruct) CreateInvoiceItems(data InvoiceItems) (InvoiceItems, error) {
	err = d.db.Model(&InvoiceItems{}).Create(&data).Error
	if err != nil {
		return InvoiceItems{}, err
	}
	return data, nil
}

func (d *DBStruct) FindAllInvoiceItems() ([]InvoiceItems, error) {
	datas := []InvoiceItems{}
	err = d.db.Model(&InvoiceItems{}).Order("id desc").Find(&datas).Error
	if err != nil {
		return []InvoiceItems{}, err
	}
	return datas, nil
}

func (d *DBStruct) FindByIdInvoiceItems(pid uint) (InvoiceItems, error) {
	data := InvoiceItems{}
	err = d.db.Model(&InvoiceItems{}).Where("id = ?", pid).Take(&data).Error
	if err != nil {
		return InvoiceItems{}, err
	}
	return data, nil
}

func (d *DBStruct) UpdateInvoiceItems(data InvoiceItems) (InvoiceItems, error) {
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
	err = d.db.Model(&InvoiceItems{}).Where("id=?", data.ID).Updates(invoiceItems).Error
	if err != nil {
		return InvoiceItems{}, err
	}
	return data, nil
}

func (d *DBStruct) DeleteInvoiceItems(pid uint) (int64, error) {
	result := d.db.Model(&InvoiceItems{}).Where("id = ?", pid).Take(&InvoiceItems{}).Delete(&InvoiceItems{})
	if result.Error != nil {
		if gorm.IsRecordNotFoundError(result.Error) {
			return 0, errors.New("InvoiceItems not found")
		}
		return 0, result.Error
	}
	return result.RowsAffected, nil
}
