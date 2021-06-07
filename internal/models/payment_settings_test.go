package models

import (
	"regexp"
	"testing"
	"time"

	"github.com/DATA-DOG/go-sqlmock"
	"github.com/stretchr/testify/assert"
)

func TestGetPaymentSettingByID(t *testing.T) {
	paymentsetting := PaymentSetting{
		UserID:      1,
		Country:     "Nepal",
		State:       "Bagmati",
		City:        "Kathmandu",
		Street:      "Lainchour",
		Postal_Code: "44600",
		Promocode:   "New year offer",
	}

	server.Mock.ExpectQuery(regexp.QuoteMeta(`SELECT`)).WillReturnRows(sqlmock.NewRows([]string{"id", "user_id", "country", "state", "city", "street", "postal_code", "promocode", "CreatedAt", "UpdatedAt"}).AddRow(1, 1, "Nepal", "Bagmati", "Kathmandu", "Lainchour", "44600", "New year offer", time.Now(), time.Now()))
	server.Mock.ExpectCommit()
	server.Mock.MatchExpectationsInOrder(false)
	paymentsetting0, err := server.store.FindByIdPaymentSetting(paymentsetting.ID)
	if err != nil {
		t.Errorf("this is the error getting one user: %v\n", err)
		return
	}
	assert.Equal(t, paymentsetting0.UserID, paymentsetting.UserID)
	assert.Equal(t, paymentsetting0.Country, paymentsetting.Country)
	assert.Equal(t, paymentsetting0.State, paymentsetting.State)
	assert.Equal(t, paymentsetting0.City, paymentsetting.City)
	assert.Equal(t, paymentsetting0.Street, paymentsetting.Street)
	assert.Equal(t, paymentsetting0.Postal_Code, paymentsetting.Postal_Code)
	assert.Equal(t, paymentsetting0.Promocode, paymentsetting.Promocode)
}

func TestFindPaymentSettingByID(t *testing.T) {
	paymentsetting := PaymentSetting{
		UserID:      1,
		Country:     "Nepal",
		State:       "Bagmati",
		City:        "Kathmandu",
		Street:      "Lainchour",
		Postal_Code: "44600",
		Promocode:   "New year offer",
	}

	server.Mock.ExpectQuery(regexp.QuoteMeta(`SELECT`)).WillReturnRows(sqlmock.NewRows([]string{"id", "user_id", "country", "state", "city", "street", "postal_code", "promocode", "CreatedAt", "UpdatedAt"}).AddRow(1, 1, "Nepal", "Bagmati", "Kathmandu", "Lainchour", "44600", "New year offer", time.Now(), time.Now()))
	server.Mock.ExpectCommit()
	server.Mock.MatchExpectationsInOrder(false)
	paymentsetting0, err := server.store.FindByIdPaymentSetting(paymentsetting.ID)
	if err != nil {
		t.Errorf("this is the error getting one user: %v\n", err)
		return
	}
	assert.Equal(t, paymentsetting0.UserID, paymentsetting.UserID)
	assert.Equal(t, paymentsetting0.Country, paymentsetting.Country)
	assert.Equal(t, paymentsetting0.State, paymentsetting.State)
	assert.Equal(t, paymentsetting0.City, paymentsetting.City)
	assert.Equal(t, paymentsetting0.Street, paymentsetting.Street)
	assert.Equal(t, paymentsetting0.Postal_Code, paymentsetting.Postal_Code)
	assert.Equal(t, paymentsetting0.Promocode, paymentsetting.Promocode)
}
