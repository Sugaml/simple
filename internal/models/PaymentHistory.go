package models

import (
	"errors"

	"github.com/jinzhu/gorm"
)

type PaymentHistory struct {
	gorm.Model
	User          *User        `gorm:"foreignkey:UserID" json:"user"`
	UserID        uint         `gorm:"not null" json:"user_id"`
	Debit         uint         `gorm:"not null" json:"debit"`
	Credit        uint         `gorm:"not null" json:"credit"`
	Balance       uint         `gorm:"not null" json:"balance"`
	Invoice       *Invoice     `gorm:"foreignkey:InvoiceID" json:"invoice"`
	InvoiceID     uint         `gorm:"not null" json:"invoice_id"`
	Transaction   *Transaction `gorm:"foreignkey:TransactionID" json:"transaction"`
	TransactionID uint         `gorm:"not null" json:"transaction_id"`
}

func (d *DBStruct) CreatePaymentHistory(data *PaymentHistory) (*PaymentHistory, error) {
	err = d.db.Model(&PaymentHistory{}).Create(&data).Error
	if err != nil {
		return &PaymentHistory{}, err
	}
	return data, nil
}
func (d *DBStruct) FindAllPaymentHistory(datas *[]PaymentHistory) (*[]PaymentHistory, error) {
	err = d.db.Model(&PaymentHistory{}).Order("id desc").Find(&datas).Error
	if err != nil {
		return &[]PaymentHistory{}, err
	}
	return datas, nil
}

func (d *DBStruct) FindByIdPaymentHistory(pid uint) (*PaymentHistory, error) {
	data := PaymentHistory{}
	err = d.db.Model(&PaymentHistory{}).Where("id = ?", pid).Take(&data).Error
	if err != nil {
		return &PaymentHistory{}, err
	}
	return &data, nil
}

func (d *DBStruct) UpdatePaymentHistory(data *PaymentHistory) (*PaymentHistory, error) {
	err = d.db.Model(&PaymentHistory{}).Update(&data).Error
	if err != nil {
		return &PaymentHistory{}, err
	}
	return data, nil
}

func (d *DBStruct) DeletePaymentHistory(pid uint) (int64, error) {
	result := d.db.Model(&PaymentHistory{}).Where("id = ?", pid).Take(&PaymentHistory{}).Delete(&PaymentHistory{})
	if result.Error != nil {
		if gorm.IsRecordNotFoundError(result.Error) {
			return 0, errors.New("paymenthistory not found")
		}
		return 0, result.Error
	}
	return result.RowsAffected, nil
}
