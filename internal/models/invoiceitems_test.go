package models

import (
	"fmt"
	"regexp"
	"testing"
	"time"

	"github.com/DATA-DOG/go-sqlmock"
	"github.com/stretchr/testify/assert"
)

var iiInstance = InvoiceItems{}

func TestInvoiceItems_Save(t *testing.T) {
	ii := InvoiceItems{
		InvoiceID:  1,
		UserID:     1,
		Particular: "test",
		Rate:       11,
		Days:       11,
		Total:      111,
	}
	server.Mock.ExpectQuery(regexp.QuoteMeta(`INSERT`)).
		WillReturnRows(sqlmock.NewRows([]string{"id"}).AddRow(1))
	server.Mock.ExpectCommit()
	server.Mock.ExpectBegin()
	server.Mock.MatchExpectationsInOrder(false)
	fmt.Println(server.Mock.ExpectationsWereMet())
	s, err := ii.Save(server.DB)
	if err != nil {
		t.Errorf("this is the error getting the invoiceItems: %v\n", err)
		return
	}
	assert.Equal(t, s.InvoiceID, ii.InvoiceID)
	assert.Equal(t, s.UserID, ii.UserID)
	assert.Equal(t, s.Particular, ii.Particular)
	assert.Equal(t, s.Rate, ii.Rate)
	assert.Equal(t, s.Days, ii.Days)
	assert.Equal(t, s.Total, ii.Total)
}

func TestInvoiceItems_FindAll(t *testing.T) {
	server.Mock.ExpectQuery(regexp.QuoteMeta(`SELECT`)).WillReturnRows(
		sqlmock.NewRows([]string{"id", "created_at", "updated_at", "invoice_id", "user_id", "particular", "rate", "days", "total"}).
			AddRow(1, time.Now(), time.Now(), 1, 1, "test", 11, 11, 111))
	d, err := iiInstance.FindAll(server.DB)
	fmt.Println(server.Mock.ExpectationsWereMet())
	if err != nil {
		t.Errorf("this is the error getting the invoiceItems: %v\n", err)
		return
	}
	assert.Equal(t, len(*d), 1)
}

func TestInvoiceItems_Find(t *testing.T) {
	ii := InvoiceItems{
		InvoiceID:  1,
		UserID:     1,
		Particular: "test",
		Rate:       11,
		Days:       11,
		Total:      111,
	}
	server.Mock.ExpectQuery(regexp.QuoteMeta(`SELECT`)).WillReturnRows(
		sqlmock.NewRows([]string{"id", "created_at", "updated_at", "invoice_id", "user_id", "particular", "rate", "days", "total"}).
			AddRow(1, time.Now(), time.Now(), 1, 1, "test", 11, 11, 111))
	server.Mock.ExpectCommit()
	server.Mock.MatchExpectationsInOrder(false)
	s, err := ii.Find(server.DB, 1)
	if err != nil {
		t.Errorf("this is the error getting one invoiceItems: %v\n", err)
		return
	}
	assert.Equal(t, s.InvoiceID, ii.InvoiceID)
	assert.Equal(t, s.UserID, ii.UserID)
	assert.Equal(t, s.Particular, ii.Particular)
	assert.Equal(t, s.Rate, ii.Rate)
	assert.Equal(t, s.Days, ii.Days)
	assert.Equal(t, s.Total, ii.Total)
}

func TestInvoiceItems_Delete(t *testing.T) {
	server.Mock.ExpectBegin()
	server.Mock.ExpectQuery(regexp.QuoteMeta(`SELECT`)).WillReturnRows(
		sqlmock.NewRows([]string{"id", "created_at", "updated_at", "invoice_id", "user_id", "particular", "rate", "days", "total"}).
			AddRow(1, time.Now(), time.Now(), 1, 1, "test", 11, 11, 111))
	server.Mock.ExpectExec(regexp.QuoteMeta(`UPDATE`)).WillReturnResult(
		sqlmock.NewResult(0, 1))
	server.Mock.ExpectCommit()
	server.Mock.MatchExpectationsInOrder(false)
	isDel, err := deductionInstance.Delete(server.DB)
	if err != nil {
		t.Errorf("this is the error deleting the invoiceItems: %v\n", err)
		return
	}
	assert.Equal(t, int64(isDel.ID), int64(1))
}

func TestInvoiceItems_Update(t *testing.T) {
	ii := InvoiceItems{
		InvoiceID:  1,
		UserID:     1,
		Particular: "test",
		Rate:       11,
		Days:       11,
		Total:      111,
	}
	server.Mock.ExpectBegin()
	server.Mock.ExpectQuery(regexp.QuoteMeta(`SELECT`)).WillReturnRows(
		sqlmock.NewRows([]string{"id", "created_at", "updated_at", "invoice_id", "user_id", "particular", "rate", "days", "total"}).
			AddRow(1, time.Now(), time.Now(), 1, 1, "test", 11, 11, 111))
	server.Mock.ExpectExec(regexp.QuoteMeta(`UPDATE`)).WillReturnResult(
		sqlmock.NewResult(0, 1))
	server.Mock.ExpectQuery(regexp.QuoteMeta(`SELECT`)).WillReturnRows(
		sqlmock.NewRows([]string{"id", "created_at", "updated_at", "invoice_id", "user_id", "particular", "rate", "days", "total"}).
			AddRow(1, time.Now(), time.Now(), 1, 1, "test", 11, 11, 111))
	server.Mock.ExpectCommit()
	server.Mock.MatchExpectationsInOrder(false)
	ii.ID = 1
	s, err := ii.Update(server.DB)
	if err != nil {
		t.Errorf("this is the error updating the invoiceItems: %v\n", err)
		return
	}
	assert.Equal(t, s.InvoiceID, ii.InvoiceID)
	assert.Equal(t, s.UserID, ii.UserID)
	assert.Equal(t, s.Particular, ii.Particular)
	assert.Equal(t, s.Rate, ii.Rate)
	assert.Equal(t, s.Days, ii.Days)
	assert.Equal(t, s.Total, ii.Total)
}
