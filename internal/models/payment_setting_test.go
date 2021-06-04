package models

import (
	"fmt"
	"regexp"
	"testing"
	"time"

	"github.com/DATA-DOG/go-sqlmock"
	"github.com/stretchr/testify/assert"
)

func TestSavePaymentSetting(t *testing.T) {
	paymentsetting0 := PaymentSetting{
		UserID:      1,
		Country:     "Nepal",
		State:       "Bagmati",
		City:        "Kathmandu",
		Street:      "Lainchour",
		Postal_Code: "44600",
		Promocode:   "New year offer",
	}
	server.Mock.ExpectQuery(regexp.QuoteMeta(
		`INSERT`)).WillReturnRows(sqlmock.NewRows([]string{"id"}).AddRow(1))
	server.Mock.ExpectCommit()
	server.Mock.ExpectBegin()
	server.Mock.MatchExpectationsInOrder(false)
	fmt.Println(server.Mock.ExpectationsWereMet())

	err := server.DB.Save(paymentsetting0)
	if err != nil {
		t.Errorf("this is the error getting the Payment Setting: %v\n", err)
		return
	}
	// assert.Equal(t, paymentsetting0.UserID, paymentsetting.UserID)
	// assert.Equal(t, paymentsetting0.Country, paymentsetting.Country)
	// assert.Equal(t, paymentsetting0.State, paymentsetting.State)
	// assert.Equal(t, paymentsetting0.City, paymentsetting.City)
	// assert.Equal(t, paymentsetting0.Street, paymentsetting.Street)
	// assert.Equal(t, paymentsetting0.Postal_Code, paymentsetting.Postal_Code)
	// assert.Equal(t, paymentsetting0.Promocode, paymentsetting.Promocode)
}

func TestFindAllPaymentSetting(t *testing.T) {
	var paymentsetting = PaymentSetting{}
	server.Mock.ExpectQuery(regexp.QuoteMeta(
		`SELECT`)).WillReturnRows(sqlmock.NewRows([]string{"id", "Country", "State", "City", "Street", "Postal_Code", "created_at", "updated_at"}).AddRow(1, "Nepal", "Bagmati", "Kathmandu", "Lainchour", "44600", time.Now(), time.Now()))
	paymentsetting0, err := paymentsetting.FindAll(server.DB)
	fmt.Println(server.Mock.ExpectationsWereMet())

	if err != nil {
		t.Errorf("this is the error getting the users: %v\n", err)
		return
	}
	assert.Equal(t, len(*paymentsetting0), 1)
}

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
	paymentsetting0, err := paymentsetting.Find(server.DB, 1)
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

func TestUpdatePaymentSetting(t *testing.T) {
	paymentsetting := PaymentSetting{
		UserID:      1,
		Country:     "Nepal",
		State:       "Bagmati",
		City:        "Kathmandu",
		Street:      "Lainchour",
		Postal_Code: "44600",
		Promocode:   "New year offer",
	}
	server.Mock.ExpectBegin()
	server.Mock.ExpectQuery(regexp.QuoteMeta(
		`SELECT`)).WillReturnRows(sqlmock.NewRows([]string{"id", "user_id", "country", "state", "city", "street", "postal_code", "promocode", "CreatedAt", "UpdatedAt"}).AddRow(1, paymentsetting.UserID, paymentsetting.Country, paymentsetting.State, paymentsetting.City, paymentsetting.Street, paymentsetting.Postal_Code, paymentsetting.Promocode, time.Now(), time.Now()))
	server.Mock.ExpectExec(regexp.QuoteMeta(
		`UPDATE`)).WillReturnResult(sqlmock.NewResult(0, 1))
	server.Mock.ExpectQuery(regexp.QuoteMeta(
		`SELECT`)).WillReturnRows(sqlmock.NewRows([]string{"id", "user_id", "country", "state", "city", "street", "postal_code", "promocode", "CreatedAt", "UpdatedAt"}).AddRow(1, 2, "India", "UtterPardesh", "Delhi", "delhi", "55600", "New year offer", time.Now(), time.Now()))
	server.Mock.ExpectCommit()
	server.Mock.MatchExpectationsInOrder(false)
	paymentsetting.ID = 1
	updatedInvoice, err := paymentsetting.Update(server.DB)
	if err != nil {
		t.Errorf("this is the error updating the user: %v\n", err)
		return
	}
	assert.Equal(t, updatedInvoice.UserID, paymentsetting.UserID)
	assert.Equal(t, updatedInvoice.Country, paymentsetting.Country)
}

func TestDeletePaymentSetting(t *testing.T) {
	paymentsetting := PaymentSetting{
		UserID:      1,
		Country:     "Nepal",
		State:       "Bagmati",
		City:        "Kathmandu",
		Street:      "Lainchour",
		Postal_Code: "44600",
		Promocode:   "New year offer",
	}
	server.Mock.ExpectBegin()
	server.Mock.ExpectQuery(regexp.QuoteMeta(
		`SELECT`)).WillReturnRows(sqlmock.NewRows([]string{"id", "user_id", "country", "state", "city", "street", "postal_code", "promocode", "CreatedAt", "UpdatedAt"}).AddRow(1, paymentsetting.UserID, paymentsetting.Country, paymentsetting.State, paymentsetting.City, paymentsetting.Street, paymentsetting.Postal_Code, paymentsetting.Postal_Code, time.Now(), time.Now()))
	server.Mock.ExpectExec(regexp.QuoteMeta(
		`UPDATE`)).WillReturnResult(sqlmock.NewResult(0, 1))
	server.Mock.ExpectCommit()
	server.Mock.MatchExpectationsInOrder(false)
	paymentsetting0, err := paymentsetting.Delete(server.DB)
	if err != nil {
		t.Errorf("this is the error updating the user: %v\n", err)
		return
	}
	assert.Equal(t, int64(paymentsetting0.ID), int64(0))
}
