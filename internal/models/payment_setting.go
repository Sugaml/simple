package models

import (
	"errors"

	"github.com/jinzhu/gorm"
)

type PaymentSetting struct {
	gorm.Model
	User        *User  `gorm:"foreignkey:UserID" json:"user"`
	UserID      uint   `gorm:"not null" json:"user_id"`
	Country     string `gorm:"not null" json:"country"`
	State       string `gorm:"not null" json:"state"`
	City        string `gorm:"not null" json:"city"`
	Street      string `gorm:"not null" json:"street"`
	Postal_Code string `gorm:"not null" json:"postal_code"`
	Promocode   string `gorm:"not null" json:"promocode"`
}

func (d *DBStruct) CreatePaymentSetting(data PaymentSetting) (PaymentSetting, error) {
	dataExist, _ := d.FindByUserIDPaymentSetting(data.UserID)
	if (PaymentSetting{}) != dataExist {
		data.ID = dataExist.ID
		dataupdate, _ := d.UpdatePaymentSetting(data)
		return dataupdate, nil
	}
	err := d.db.Model(&PaymentSetting{}).Create(&data).Error
	if err != nil {
		return PaymentSetting{}, err
	}
	return data, nil
}
func (d *DBStruct) FindAllPaymentSetting() ([]PaymentSetting, error) {
	datas := []PaymentSetting{}
	err = d.db.Model(&PaymentSetting{}).Find(&datas).Error
	if err != nil {
		return []PaymentSetting{}, err
	}
	return datas, nil
}

func (d *DBStruct) FindByIdPaymentSetting(uid uint) (PaymentSetting, error) {
	data := PaymentSetting{}
	err = d.db.Model(&PaymentSetting{}).Where("id = ?", uid).Take(&data).Error
	if err != nil {
		return PaymentSetting{}, err
	}
	return data, nil
}
func (d *DBStruct) FindByUserIDPaymentSetting(uid uint) (PaymentSetting, error) {
	data := PaymentSetting{}
	err = d.db.Model(&PaymentSetting{}).Where("user_id = ?", uid).Take(&data).Error
	if err != nil {
		return PaymentSetting{}, err
	}
	return data, nil
}
func (d *DBStruct) UpdatePaymentSetting(paymentsetting PaymentSetting) (PaymentSetting, error) {
	err = d.db.Model(&PaymentSetting{}).Update(&paymentsetting).Error
	if err != nil {
		return PaymentSetting{}, err
	}
	return paymentsetting, nil
}
func (d *DBStruct) DeletePaymentSetting(pid uint) (int64, error) {
	result := d.db.Model(&PaymentSetting{}).Where("id = ?", pid).Take(&PaymentSetting{}).Delete(&PaymentSetting{})
	if result.Error != nil {
		if gorm.IsRecordNotFoundError(result.Error) {
			return 0, errors.New("PaymentSetting not found")
		}
		return 0, result.Error
	}
	return result.RowsAffected, nil
}
func (data *PaymentSetting) FindPromocode(db *gorm.DB) (*PromoCode, error) {
	datas := PromoCode{}
	err = db.Model(&PromoCode{}).Where("code = ? and expiry_date", data.Promocode).Find(&datas).Error
	if err != nil {
		return &PromoCode{}, err
	}
	return &datas, nil
}
