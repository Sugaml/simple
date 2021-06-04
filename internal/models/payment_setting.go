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

func (d *DBStruct) Create(data PaymentSetting) (PaymentSetting, error) {
	err := d.db.Model(&PaymentSetting{}).Create(&data).Error
	if err != nil {
		return PaymentSetting{}, err
	}
	return data, nil
}

func (d *DBStruct) FindById(uid uint) (PaymentSetting, error) {
	data := PaymentSetting{}
	err = d.db.Model(&PaymentSetting{}).Where("id = ?", uid).Take(&data).Error
	if err != nil {
		return PaymentSetting{}, err
	}
	return data, nil
}

func (d *DBStruct) FindAll() ([]PaymentSetting, error) {
	data := []PaymentSetting{}
	err = d.db.Model(&PaymentSetting{}).Order("id desc").Find(&data).Error
	if err != nil {
		return []PaymentSetting{}, err
	}
	return data, nil
}
func (d *DBStruct) Update(paymentsetting PaymentSetting) (PaymentSetting, error) {
	err = d.db.Model(&PaymentSetting{}).Update(&paymentsetting).Error
	if err != nil {
		return PaymentSetting{}, err
	}
	return paymentsetting, nil
}
func (d *DBStruct) Delete(pid uint) (int64, error) {
	result := d.db.Model(&PaymentSetting{}).Where("id = ?", pid).Take(&PaymentSetting{}).Delete(&PaymentSetting{})
	if result.Error != nil {
		if gorm.IsRecordNotFoundError(result.Error) {
			return 0, errors.New("PaymentSetting not found")
		}
		return 0, result.Error
	}
	return result.RowsAffected, nil
}
func (paymentsetting *PaymentSetting) CreatePayment(db *gorm.DB) *gorm.DB {
	result := db.Create(&paymentsetting)
	if err != nil {
		panic(err)
	}
	return result
}
func (data *PaymentSetting) FindAll(db *gorm.DB) (*[]PaymentSetting, error) {
	datas := []PaymentSetting{}
	err = db.Model(&PaymentSetting{}).Order("id desc").Find(&datas).Error
	if err != nil {
		return &[]PaymentSetting{}, err
	}
	return &datas, nil
}

func (data PaymentSetting) Find(db *gorm.DB, pid uint) (PaymentSetting, error) {
	err = db.Model(&PaymentSetting{}).Where("id = ?", pid).Take(&data).Error
	if err != nil {
		return PaymentSetting{}, err
	}
	return data, nil
}

func (data *PaymentSetting) FindByUserID(db *gorm.DB, uid uint64) (*PaymentSetting, error) {
	err = db.Model(&PaymentSetting{}).Where("user_id = ?", uid).Take(&data).Error
	if err != nil {
		return &PaymentSetting{}, err
	}
	return data, nil
}
func (data *PaymentSetting) Update(db *gorm.DB) (*PaymentSetting, error) {
	err = db.Model(&PaymentSetting{}).Update(&data).Error
	if err != nil {
		return &PaymentSetting{}, err
	}
	return data, nil
}
func (data *PaymentSetting) Delete(db *gorm.DB) (*PaymentSetting, error) {
	err = db.Model(&PaymentSetting{}).Delete(&data).Error
	if err != nil {
		return &PaymentSetting{}, err
	}
	return data, nil
}
