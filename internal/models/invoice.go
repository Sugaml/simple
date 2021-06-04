package models

import (
	"errors"
	"time"

	"github.com/jinzhu/gorm"
)

type Invoice struct {
	gorm.Model
	User         *User           `gorm:"foreignkey:UserID" json:"user"`
	UserID       uint            `gorm:"not null" json:"user_id"`
	Date         time.Time       `gorm:"not null" json:"date"`
	StartDate    time.Time       `gorm:"not null" json:"start_date"`
	EndDate      time.Time       `gorm:"not null" json:"end_date"`
	TotalCost    uint            `gorm:"not null" json:"total_cost"`
	PromoCode    *PromoCode      `gorm:"foreignkey:PromoCodeID" json:"promocode"`
	PromoCodeID  uint            `gorm:"null" json:"promo_code_id"`
	Deduction    *Deduction      `gorm:"foreignkey:DeductionID" json:"deduction"`
	DeductionID  uint            `gorm:"null" json:"deduction_id"`
	InvoiceItems *[]InvoiceItems `gorm:"null" json:"invoice_items"`
}

func (d *DBStruct) CreateInvoice(invoice Invoice) (Invoice, error) {
	err = d.db.Model(&Invoice{}).Create(&invoice).Error
	if err != nil {
		return Invoice{}, err
	}
	return invoice, nil
}

func (data *Invoice) Save(db *gorm.DB) (*Invoice, error) {
	err = db.Model(&Invoice{}).Create(&data).Error
	if err != nil {
		return &Invoice{}, err
	}
	return data, nil
}

func (data *Invoice) FindAll(db *gorm.DB) (*[]Invoice, error) {
	datas := []Invoice{}
	err = db.Model(&Invoice{}).Order("id desc").Find(&datas).Error
	if err != nil {
		return &[]Invoice{}, err
	}
	return &datas, nil
}

func (data *Invoice) Find(db *gorm.DB, pid uint64) (*Invoice, error) {
	err = db.Model(&Invoice{}).Where("id = ?", pid).Take(&data).Error
	if err != nil {
		return &Invoice{}, err
	}
	return data, nil
}
func (data *Invoice) Update(db *gorm.DB) (*Invoice, error) {
	var invoice = Invoice{}
	if data.UserID != 0 {
		invoice.UserID = data.UserID
	}
	if data.TotalCost != 0 {
		invoice.TotalCost = data.TotalCost
	}
	if data.PromoCodeID != 0 {
		invoice.PromoCodeID = data.PromoCodeID
	}
	if data.DeductionID != 0 {
		invoice.DeductionID = data.DeductionID
	}
	err = db.Model(&Invoice{}).Where("id = ?", data.ID).Updates(invoice).Error
	if err != nil {
		return &Invoice{}, err
	}
	return data, nil
}

func (data *Invoice) Delete(db *gorm.DB, id uint64) (int64, error) {
	db = db.Model(&Invoice{}).Where("id = ?", id).Take(&Invoice{}).Delete(&Invoice{})
	if db.Error != nil {
		if gorm.IsRecordNotFoundError(db.Error) {
			return 0, errors.New("Invoice not found")
		}
		return 0, db.Error
	}
	return db.RowsAffected, nil
}
