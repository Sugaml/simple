package models

import (
	"time"

	"github.com/jinzhu/gorm"
)

type Store interface {
	CreatePaymentSetting(data PaymentSetting) (PaymentSetting, error)
	FindAllPaymentSetting() ([]PaymentSetting, error)
	FindByUserIDPaymentSetting(uid uint) (PaymentSetting, error)
	FindByIdPaymentSetting(uid uint) (PaymentSetting, error)
	UpdatePaymentSetting(data PaymentSetting) (PaymentSetting, error)
	DeletePaymentSetting(pid uint) (int64, error)

	CreateInvoice(data Invoice) (Invoice, error)
	FindByIdInvoice(uid uint) (Invoice, error)
	UpdateInvoice(data Invoice) (Invoice, error)
	DeleteInvoice(pid uint) (int64, error)

	CreatePromocode(data PromoCode) (PromoCode, error)
	FindAllPromocode() ([]PromoCode, error)
	FindByIdPromocode(uid uint) (PromoCode, error)
	FindByPromoCode(promocode string) (PromoCode, error)
	UpdatePromocode(data PromoCode) (PromoCode, error)
	DeletePromocode(pid uint) (int64, error)

	CreateDeduction(data Deduction) (Deduction, error)
	FindAllDeduction() ([]Deduction, error)
	FindByIdDeduction(uid uint) (Deduction, error)
	FindByCountryDeduction(country string) (Deduction, error)
	UpdateDeduction(data Deduction) (Deduction, error)
	DeleteDeduction(pid uint) (int64, error)

	CreateInvoiceItems(data InvoiceItems) (InvoiceItems, error)
	FindAllInvoiceItems() ([]InvoiceItems, error)
	FindByIdInvoiceItems(uid uint) (InvoiceItems, error)
	UpdateInvoiceItems(data InvoiceItems) (InvoiceItems, error)
	DeleteInvoiceItems(pid uint) (int64, error)

	CreateThreshold(data PaymentThreshold) (PaymentThreshold, error)
	FindAllThreshold() ([]PaymentThreshold, error)
	FindByIdThreshold(uid uint) (PaymentThreshold, error)
	FindByUserIDThreshold(uid uint) (PaymentThreshold, error)
	UpdateThreshold(data PaymentThreshold) (PaymentThreshold, error)
	DeleteThreshold(pid uint) (int64, error)

	CreatePaymentHistory(data PaymentHistory) (PaymentHistory, error)
	FindByUserIDPaymentHistory(uid uint) ([]PaymentHistory, error)

	FindAllByUser(userID uint, startDate, endDate time.Time) ([]Project, error)
	FindSubscription(pid uint) (Subscription, error)

	CreateTransaction(data Transaction) (Transaction, error)
	FindAllTransaction() ([]Transaction, error)
	FindByIdTransaction(uid uint) (Transaction, error)
	UpdateTransaction(data Transaction) (Transaction, error)
	DeleteTransaction(pid uint) (int64, error)

	CreateGateway(data Gateway) (Gateway, error)
	FindAllGateway() ([]Gateway, error)
	FindByIdGateway(uid uint) (Gateway, error)
	UpdateGateway(data Gateway) (Gateway, error)
	DeleteGateway(pid uint) (int64, error)

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
