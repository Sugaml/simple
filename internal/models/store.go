package models

import (
	"github.com/jinzhu/gorm"
)

type Store interface {
	CreatePaymentSetting(data *PaymentSetting) (*PaymentSetting, error)
	FindAllPaymentSetting(datas *[]PaymentSetting) (*[]PaymentSetting, error)
	FindByIdPaymentSetting(uid uint) (*PaymentSetting, error)
	UpdatePaymentSetting(datas *PaymentSetting) (*PaymentSetting, error)
	DeletePaymentSetting(pid uint) (int64, error)

	CreateInvoice(invoice Invoice) (Invoice, error)
	FindAllInvoice() ([]Invoice, error)
	FindByIdInvoice(uid uint) (Invoice, error)
	UpdateInvoice(data Invoice) (Invoice, error)
	DeleteInvoice(pid uint) (int64, error)

	CreatePaymentHistory(data *PaymentHistory) (*PaymentHistory, error)
	FindAllPaymentHistory(datas *[]PaymentHistory) (*[]PaymentHistory, error)
	FindByIdPaymentHistory(uid uint) (*PaymentHistory, error)
	UpdatePaymentHistory(data *PaymentHistory) (*PaymentHistory, error)
	DeletePaymentHistory(pid uint) (int64, error)

	CreatePromocode(data *PromoCode) (*PromoCode, error)
	FindAllPromocode(datas *[]PromoCode) (*[]PromoCode, error)
	FindByIdPromocode(uid uint) (*PromoCode, error)
	UpdatePromocode(data *PromoCode) (*PromoCode, error)
	DeletePromocode(pid uint) (int64, error)

	CreateDeduction(data Deduction) (Deduction, error)
	FindAllDeduction() ([]Deduction, error)
	FindByIdDeduction(uid uint) (Deduction, error)
	UpdateDeduction(data Deduction) (Deduction, error)
	DeleteDeduction(pid uint) (int64, error)

	CreateThreshold(data Threshold) (Threshold, error)
	FindAllThreshold() ([]Threshold, error)
	FindByIdThreshold(uid uint) (Threshold, error)
	UpdateThreshold(data Threshold) (Threshold, error)
	DeleteThreshold(pid uint) (int64, error)

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
