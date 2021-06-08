package models

import (
	"errors"

	"github.com/jinzhu/gorm"
)

type Transaction struct {
	gorm.Model
	Title     string   `gorm:"not null" json:"title"`
	User      *User    `gorm:"foreignkey:UserID" json:"user"`
	UserID    uint     `gorm:"not null" json:"user_id"`
	Amount    uint     `gorm:"not null" json:"amount"`
	Credit    uint     `gorm:"not null" json:"promo_code_id"`
	Balance   uint     `gorm:"not null" json:"balance"`
	Invoice   *Invoice `gorm:"foreignkey:InvoiceID" json:"invoice"`
	InvoiceID uint     `gorm:"not null" json:"invoice_id"`
	Gateway   *Gateway `gorm:"foreignkey:GatewayID" json:"gateway"`
	GatewayID uint     `gorm:"not null" json:"gateway_id"`
	Status    string   `gorm:"not null" json:"status"`
}

func (d *DBStruct) CreateTransaction(data Transaction) (Transaction, error) {
	err = d.db.Model(&Transaction{}).Create(&data).Error
	if err != nil {
		return Transaction{}, err
	}
	return data, nil
}
func (d *DBStruct) FindAllTransaction() ([]Transaction, error) {
	datas := []Transaction{}
	err = d.db.Model(&Transaction{}).Order("id desc").Find(&datas).Error
	if err != nil {
		return []Transaction{}, err
	}
	return datas, nil
}

func (d *DBStruct) FindByIdTransaction(pid uint) (Transaction, error) {
	data := Transaction{}
	err = d.db.Model(&Transaction{}).Where("id = ?", pid).Take(&data).Error
	if err != nil {
		return Transaction{}, err
	}
	return data, nil
}

func (d *DBStruct) UpdateTransaction(data Transaction) (Transaction, error) {
	err = d.db.Model(&Transaction{}).Update(&data).Error
	if err != nil {
		return Transaction{}, err
	}
	return data, nil
}

func (d *DBStruct) DeleteTransaction(pid uint) (int64, error) {
	result := d.db.Model(&Transaction{}).Where("id = ?", pid).Take(&Transaction{}).Delete(&Transaction{})
	if result.Error != nil {
		if gorm.IsRecordNotFoundError(result.Error) {
			return 0, errors.New("promocode not found")
		}
		return 0, result.Error
	}
	return result.RowsAffected, nil
}
