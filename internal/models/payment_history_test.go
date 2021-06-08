package models

import (
	"fmt"
	"regexp"
	"testing"
	"time"

	"github.com/DATA-DOG/go-sqlmock"
	"github.com/stretchr/testify/assert"
)

func TestSavePaymentHistory(t *testing.T) {
	paymenthistory0 := PaymentHistory{
		UserID:        1,
		Credit:        500,
		Debit:         500,
		Balance:       0,
		InvoiceID:     1,
		TransactionID: 1,
	}
	paymenthistory0.ID = 1
	server.Mock.ExpectQuery(regexp.QuoteMeta(
		`INSERT`)).WillReturnRows(sqlmock.NewRows([]string{"id"}).AddRow(1))
	server.Mock.ExpectCommit()
	server.Mock.ExpectBegin()
	server.Mock.MatchExpectationsInOrder(false)
	fmt.Println(server.Mock.ExpectationsWereMet())

	paymenthistory, err := server.store.CreatePaymentHistory(&paymenthistory0)
	if err != nil {
		t.Errorf("this is the error getting the Payment Setting: %v\n", err)
		return
	}
	assert.Equal(t, paymenthistory0.ID, paymenthistory.ID)
	assert.Equal(t, paymenthistory0.UserID, paymenthistory.UserID)
	assert.Equal(t, paymenthistory0.Credit, paymenthistory.Credit)
	assert.Equal(t, paymenthistory0.Debit, paymenthistory.Debit)
	assert.Equal(t, paymenthistory0.Balance, paymenthistory.Balance)
	assert.Equal(t, paymenthistory0.TransactionID, paymenthistory.TransactionID)
	assert.Equal(t, paymenthistory0.InvoiceID, paymenthistory.InvoiceID)
}

func TestFindAllPaymentHistory(t *testing.T) {
	paymenthistories := []PaymentHistory{
		{
			UserID:        1,
			Credit:        500,
			Debit:         500,
			Balance:       0,
			TransactionID: 1,
			InvoiceID:     1,
		},
		{
			UserID:        2,
			Credit:        300,
			Debit:         300,
			Balance:       0,
			TransactionID: 2,
			InvoiceID:     2,
		},
	}
	fmt.Println(len(paymenthistories))
	var testpaymenthisotries *[]PaymentHistory
	server.Mock.ExpectQuery(regexp.QuoteMeta(
		`SELECT`)).WillReturnRows(sqlmock.NewRows([]string{"id", "user_id", "debit", "Credit", "balance", "invoice_id", "transaction_id", "created_at", "updated_at"}).AddRow(1, 1, 500, 500, 0, 1, 1, time.Now(), time.Now()))
	testpaymenthisotries, err = server.store.FindAllPaymentHistory(&paymenthistories)
	fmt.Println(server.Mock.ExpectationsWereMet())

	if err != nil {
		t.Errorf("this is the error getting the users: %v\n", err)
		return
	}
	assert.Equal(t, len(*testpaymenthisotries), len(paymenthistories))
}

func TestGetPaymentHistoryByID(t *testing.T) {
	paymenthistory := PaymentHistory{
		UserID:        1,
		Credit:        500,
		Debit:         500,
		Balance:       0,
		TransactionID: 1,
		InvoiceID:     1,
	}
	paymenthistory.ID = 1
	server.Mock.ExpectQuery(regexp.QuoteMeta(`SELECT`)).WillReturnRows(sqlmock.NewRows([]string{"id", "user_id", "credit", "debit", "balance", "transaction_id", "invoice_id", "CreatedAt", "UpdatedAt"}).AddRow(1, 1, 500, 500, 0, 1, 1, time.Now(), time.Now()))
	server.Mock.ExpectCommit()
	server.Mock.MatchExpectationsInOrder(false)
	paymenthistory0, err := server.store.FindByIdPaymentHistory(paymenthistory.ID)
	if err != nil {
		t.Errorf("this is the error getting one user: %v\n", err)
		return
	}
	assert.Equal(t, paymenthistory0.ID, paymenthistory.ID)
	assert.Equal(t, paymenthistory0.UserID, paymenthistory.UserID)
	assert.Equal(t, paymenthistory0.Credit, paymenthistory.Credit)
	assert.Equal(t, paymenthistory0.Debit, paymenthistory.Debit)
	assert.Equal(t, paymenthistory0.Balance, paymenthistory.Balance)
	assert.Equal(t, paymenthistory0.TransactionID, paymenthistory.TransactionID)
	assert.Equal(t, paymenthistory0.InvoiceID, paymenthistory.InvoiceID)
}

func TestUpdatePaymentHistory(t *testing.T) {
	paymenthistory := PaymentHistory{
		UserID:        1,
		Credit:        500,
		Debit:         500,
		Balance:       0,
		InvoiceID:     1,
		TransactionID: 1,
	}
	server.Mock.ExpectBegin()
	server.Mock.ExpectQuery(regexp.QuoteMeta(
		`SELECT`)).WillReturnRows(sqlmock.NewRows([]string{"id", "user_id", "credit", "debit", "balance", "transaction_id", "invoice_id", "CreatedAt", "UpdatedAt"}).AddRow(1, paymenthistory.UserID, paymenthistory.Credit, paymenthistory.Debit, paymenthistory.Balance, paymenthistory.TransactionID, paymenthistory.InvoiceID, time.Now(), time.Now()))
	server.Mock.ExpectExec(regexp.QuoteMeta(
		`UPDATE`)).WillReturnResult(sqlmock.NewResult(0, 1))
	server.Mock.ExpectQuery(regexp.QuoteMeta(
		`SELECT`)).WillReturnRows(sqlmock.NewRows([]string{"id", "user_id", "credit", "debit", "balance", "transaction_id", "invoice_id", "CreatedAt", "UpdatedAt"}).AddRow(1, 2, 500, 500, 0, 1, 1, time.Now(), time.Now()))
	server.Mock.ExpectCommit()
	server.Mock.MatchExpectationsInOrder(false)
	paymenthistory.ID = 1
	updatedPaymentHistory, err := server.store.UpdatePaymentHistory(&paymenthistory)
	if err != nil {
		t.Errorf("this is the error updating the user: %v\n", err)
		return
	}
	assert.Equal(t, updatedPaymentHistory.UserID, paymenthistory.UserID)

}

func TestDeletePaymentHistory(t *testing.T) {
	paymenthistory := PaymentHistory{
		UserID:        1,
		Credit:        500,
		Debit:         500,
		Balance:       0,
		InvoiceID:     1,
		TransactionID: 1,
	}
	paymenthistory.ID = 1
	server.Mock.ExpectBegin()
	server.Mock.ExpectQuery(regexp.QuoteMeta(
		`SELECT`)).WillReturnRows(sqlmock.NewRows([]string{"id", "user_id", "credit", "debit", "balance", "transaction_id", "invoice_id", "CreatedAt", "UpdatedAt"}).AddRow(1, paymenthistory.UserID, paymenthistory.Credit, paymenthistory.Debit, paymenthistory.Balance, paymenthistory.TransactionID, paymenthistory.InvoiceID, time.Now(), time.Now()))
	server.Mock.ExpectExec(regexp.QuoteMeta(
		`UPDATE`)).WillReturnResult(sqlmock.NewResult(0, 1))
	server.Mock.ExpectCommit()
	server.Mock.MatchExpectationsInOrder(false)
	id, err := server.store.DeletePaymentHistory(paymenthistory.ID)
	if err != nil {
		t.Errorf("this is the error updating the user: %v\n", err)
		return
	}
	assert.Equal(t, id, int64(1))
}
