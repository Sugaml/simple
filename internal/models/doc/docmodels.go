package doc

import (
	"time"
)

type PaymentHistory struct {
	User          *User        `json:"user"`
	UserID        uint         `json:"user_id"`
	Debit         uint         `json:"debit"`
	Credit        uint         `json:"credit"`
	Balance       uint         `json:"balance"`
	Invoice       *Invoice     `json:"invoice"`
	InvoiceID     uint         `json:"invoice_id"`
	Transaction   *Transaction `json:"transaction"`
	TransactionID uint         `json:"transaction_id"`
}

type Transaction struct {
	Title     string   `json:"title"`
	User      *User    `json:"user"`
	UserID    uint     `json:"user_id"`
	Amount    uint     `json:"amount"`
	Credit    uint     `json:"promo_code_id"`
	Balance   uint     `json:"balance"`
	Invoice   *Invoice `son:"invoice"`
	InvoiceID uint     `json:"invoice_id"`
	Gateway   *Gateway `json:"gateway"`
	GatewayID uint     `json:"gateway_id"`
	Status    string   `json:"status"`
}

type Gateway struct {
	Name      string `json:"name"`
	AccessKey string `json:"accesskey"`
	SecretKey string `json:"secretkey"`
	Token     string `json:"token"`
	Others    string `json:"others"`
	Url       string `json:"url"`
	Active    bool   `json:"active"`
}

type Threshold struct {
	User           uint   `gorm:"foreignkey:UserID" json:"user"`
	UserID         uint   `gorm:"not null" json:"user_id"`
	ThresholdLimit uint   `gorm:"not null" json:"threshold_limit"`
	Email          string `gorm:"not null" json:"email"`
	Active         bool   `gorm:"not null" json:"active"`
}

type Deduction struct {
	Name      string `json:"name"`
	Value     uint   `json:"value"`
	IsPercent bool   `json:"is_percent"`
	Country   string `json:"country"`
}

type PromoCode struct {
	Title      string    `json:"title"`
	Code       uint      `json:"code"`
	IsPercent  bool      `json:"is_percent"`
	Discount   uint      `json:"discount"`
	ExpiryDate time.Time `json:"expiry_date"`
	Limit      uint      `json:"limit"`
	Count      uint      `json:"count"`
	Active     bool      `json:"active"`
}

type InvoiceItems struct {
	Invoice    uint   `json:"invoice"`
	InvoiceID  uint   `json:"invoice_id"`
	User       *User  `json:"user"`
	UserID     uint   `json:"user_id"`
	Particular string `json:"particular"`
	Rate       uint   `json:"rate"`
	Days       uint   `json:"days"`
	Total      uint   `json:"total"`
}

type PaymentSetting struct {
	User        *User  `json:"user"`
	UserID      uint   `json:"user_id"`
	Country     string `json:"country"`
	State       string `json:"state"`
	City        string `json:"city"`
	Street      string `json:"street"`
	Postal_Code string `json:"postal_code"`
	Promocode   string `json:"promocode"`
}

type User struct {
	FirstName     string `json:"first_name,omitempty"`
	LastName      string `json:"last_name,omitempty"`
	Email         string `json:"email,omitempty"`
	Image         string `json:"image,omitempty"`
	Company       string `json:"company,omitempty"`
	Designation   string `json:"designation,omitempty"`
	Password      string `json:"password,omitempty"`
	EmailVerified bool   `json:"email_verified"`
	Active        bool   `json:"active"`
	IsAdmin       bool   `json:"is_admin,omitempty"`
}

type Invoice struct {
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
