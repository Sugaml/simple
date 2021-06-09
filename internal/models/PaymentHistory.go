package models

import (
	"time"

	"github.com/jinzhu/gorm"
)

type PaymentHistory struct {
	gorm.Model
	User          *User        `gorm:"foreignkey:UserID" json:"user"`
	UserID        uint         `gorm:"not null" json:"user_id"`
	Debit         float64      `gorm:"not null" json:"debit"`
	Credit        float64      `gorm:"not null" json:"credit"`
	Balance       float64      `gorm:"not null" json:"balance"`
	Invoice       *Invoice     `gorm:"foreignkey:InvoiceID" json:"invoice"`
	InvoiceID     uint         `gorm:"not null" json:"invoice_id"`
	Transaction   *Transaction `gorm:"foreignkey:TransactionID" json:"transaction"`
	TransactionID uint         `gorm:"not null" json:"transaction_id"`
	Date          time.Time    `gorm:"not null" json:"date"`
}

func (d *DBStruct) CreatePaymentHistory(data PaymentHistory) (PaymentHistory, error) {
	err = d.db.Model(&PaymentHistory{}).Create(&data).Error
	if err != nil {
		return PaymentHistory{}, err
	}
	return data, nil
}

func (d *DBStruct) FindByUserIDPaymentHistory(uid uint) ([]PaymentHistory, error) {
	datas := []PaymentHistory{}
	err = d.db.Model(&PaymentHistory{}).Where("user_id = ?", uid).Find(&datas).Error
	if err != nil {
		return []PaymentHistory{}, err
	}
	return datas, nil
}
