package models

import (
	"github.com/jinzhu/gorm"
)

type Store interface {
	Create(data PaymentSetting) (PaymentSetting, error)
	CreateInvoice(invoice Invoice) (Invoice, error)
	FindAll() ([]PaymentSetting, error)
	FindById(uid uint) (PaymentSetting, error)
	Update(data PaymentSetting) (PaymentSetting, error)
	Delete(pid uint) (int64, error)
	MigrateDB()
}
type DBStruct struct {
	db *gorm.DB
}

func NewDB(db *gorm.DB) Store {
	return &DBStruct{db}
}

func (d *DBStruct) MigrateDB() {
	d.db.AutoMigrate(
		User{},
		Project{},
		Subscription{},
		Invoice{},
		InvoiceItems{},
		PaymentHistory{},
		Deduction{},
		PromoCode{},
		PaymentSetting{},
	)
}
