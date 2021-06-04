package models

import (
	"fmt"
	"regexp"
	"testing"
	"time"

	"github.com/DATA-DOG/go-sqlmock"
	"github.com/stretchr/testify/assert"
)

func TestSaveInvoice(t *testing.T) {
	d := time.Now()
	startdate := time.Date(d.Year(), d.Month(), 1, 0, 0, 0, 0, time.Local)
	enddate := d
	newInvoice := Invoice{
		UserID:    1,
		StartDate: startdate,
		EndDate:   enddate,
		TotalCost: 200,
	}
	server.Mock.ExpectQuery(regexp.QuoteMeta(
		`INSERT`)).WillReturnRows(sqlmock.NewRows([]string{"id"}).AddRow(1))
	server.Mock.ExpectCommit()
	server.Mock.ExpectBegin()
	server.Mock.MatchExpectationsInOrder(false)
	fmt.Println(server.Mock.ExpectationsWereMet())

	savedInvoice, err := newInvoice.Save(server.DB)
	if err != nil {
		t.Errorf("this is the error getting the Invoices: %v\n", err)
		return
	}
	assert.Equal(t, newInvoice.UserID, savedInvoice.UserID)
	assert.Equal(t, newInvoice.TotalCost, savedInvoice.TotalCost)
	assert.Equal(t, newInvoice.StartDate, savedInvoice.StartDate)
	assert.Equal(t, newInvoice.EndDate, savedInvoice.EndDate)
}
func TestFindAllInvoice(t *testing.T) {
	d := time.Now()
	startdate := time.Date(d.Year(), d.Month(), 1, 0, 0, 0, 0, time.Local)
	enddate := d
	server.Mock.ExpectQuery(regexp.QuoteMeta(
		`SELECT`)).WillReturnRows(sqlmock.NewRows([]string{"id", "user_id", "start_date", "end_date", "total_cost", "CreatedAt", "UpdatedAt"}).AddRow(1, 1, startdate, enddate, 200, time.Now(), time.Now()))
	invoice, err := invoiceInstance.FindAll(server.DB)
	fmt.Println(server.Mock.ExpectationsWereMet())

	if err != nil {
		t.Errorf("this is the error getting the users: %v\n", err)
		return
	}
	assert.Equal(t, len(*invoice), 1)
}

func TestGetInvoiceByID(t *testing.T) {
	d := time.Now()
	startdate := time.Date(d.Year(), d.Month(), 1, 0, 0, 0, 0, time.Local)
	enddate := d
	invoice := Invoice{
		UserID:    1,
		StartDate: startdate,
		EndDate:   enddate,
		TotalCost: 200,
	}
	server.Mock.ExpectQuery(regexp.QuoteMeta(`SELECT`)).WillReturnRows(sqlmock.NewRows([]string{"id", "user_id", "total_cost"}).AddRow(1, invoice.UserID, 200))
	server.Mock.ExpectCommit()
	server.Mock.MatchExpectationsInOrder(false)
	foundInvoice, err := invoiceInstance.Find(server.DB, 1)
	fmt.Println(foundInvoice)
	if err != nil {
		t.Errorf("this is the error getting one user: %v\n", err)
		return
	}
	assert.NotNil(t, foundInvoice)
	assert.NoError(t, err)
}
func TestUpdateInvoice(t *testing.T) {
	d := time.Now()
	startdate := time.Date(d.Year(), d.Month(), 1, 0, 0, 0, 0, time.Local)
	enddate := d
	invoiceUpdate := Invoice{
		UserID:    2,
		StartDate: startdate,
		EndDate:   enddate,
		TotalCost: 500,
	}

	server.Mock.ExpectBegin()
	server.Mock.ExpectQuery(regexp.QuoteMeta(
		`SELECT`)).WillReturnRows(sqlmock.NewRows([]string{"id", "user_id", "start_date", "end_date", "total_cost", "CreatedAt", "UpdatedAt"}).AddRow(1, 1, startdate, enddate, 200, time.Now(), time.Now()))
	server.Mock.ExpectExec(regexp.QuoteMeta(
		`UPDATE`)).WillReturnResult(sqlmock.NewResult(0, 1))
	server.Mock.ExpectQuery(regexp.QuoteMeta(
		`SELECT`)).WillReturnRows(sqlmock.NewRows([]string{"id", "user_id", "start_date", "end_date", "total_cost", "CreatedAt", "UpdatedAt"}).AddRow(1, 2, startdate, enddate, 500, time.Now(), time.Now()))
	server.Mock.ExpectCommit()
	server.Mock.MatchExpectationsInOrder(false)
	invoiceUpdate.ID = 1
	updatedInvoice, err := invoiceUpdate.Update(server.DB)
	if err != nil {
		t.Errorf("this is the error updating the user: %v\n", err)
		return
	}

	assert.Equal(t, updatedInvoice.UserID, invoiceUpdate.UserID)
	assert.Equal(t, updatedInvoice.TotalCost, invoiceUpdate.TotalCost)
}

func TestDeleteInvoice(t *testing.T) {
	d := time.Now()
	startdate := time.Date(d.Year(), d.Month(), 1, 0, 0, 0, 0, time.Local)
	enddate := d
	server.Mock.ExpectBegin()
	server.Mock.ExpectQuery(regexp.QuoteMeta(
		`SELECT`)).WillReturnRows(sqlmock.NewRows([]string{"id", "user_id", "start_date", "end_date", "total_cost", "CreatedAt", "UpdatedAt"}).AddRow(1, 1, startdate, enddate, 200, time.Now(), time.Now()))
	server.Mock.ExpectExec(regexp.QuoteMeta(
		`UPDATE`)).WillReturnResult(sqlmock.NewResult(0, 1))
	server.Mock.ExpectCommit()
	server.Mock.MatchExpectationsInOrder(false)
	isDeleted, err := invoiceInstance.Delete(server.DB, 1)
	if err != nil {
		t.Errorf("this is the error updating the user: %v\n", err)
		return
	}
	assert.Equal(t, isDeleted, int64(1))
}
