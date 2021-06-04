package models

import (
	"fmt"
	"regexp"
	"testing"
	"time"

	"github.com/DATA-DOG/go-sqlmock"
	"github.com/stretchr/testify/assert"
)

var promoCodeInstance = PromoCode{}

func TestPromoCode_Save(t *testing.T) {
	ed := time.Now()
	pcd := PromoCode{
		Title:      "test",
		Code:       111,
		IsPercent:  true,
		ExpiryDate: ed,
		Discount:   11,
		Limit:      3,
		Count:      1,
		Active:     true,
	}
	server.Mock.ExpectQuery(regexp.QuoteMeta(
		`INSERT`)).WillReturnRows(sqlmock.NewRows([]string{"id"}).AddRow(1))
	server.Mock.ExpectCommit()
	server.Mock.ExpectBegin()
	server.Mock.MatchExpectationsInOrder(false)
	fmt.Println(server.Mock.ExpectationsWereMet())
	saved, err := pcd.Save(server.DB)
	if err != nil {
		t.Errorf("this is the error getting the PromoCode: %v\n", err)
		return
	}
	assert.Equal(t, saved.Title, pcd.Title)
	assert.Equal(t, saved.Code, pcd.Code)
	assert.Equal(t, saved.IsPercent, pcd.IsPercent)
	assert.Equal(t, saved.Discount, pcd.Discount)
	assert.Equal(t, saved.ExpiryDate, pcd.ExpiryDate)
	assert.Equal(t, saved.Limit, pcd.Limit)
	assert.Equal(t, saved.Count, pcd.Count)
	assert.Equal(t, saved.Active, pcd.Active)
}

func TestPromoCode_FindAll(t *testing.T) {
	ed := time.Now()
	server.Mock.ExpectQuery(regexp.QuoteMeta(`SELECT`)).WillReturnRows(
		sqlmock.NewRows([]string{"id", "created_at", "updated_at", "title", "code", "is_percent", "discount",
			"expiry_date", "limit", "count", "active"}).AddRow(11, time.Now(), time.Now(), "test", 111, true, 11,
			ed, 3, 1, true))
	pc, err := promoCodeInstance.FindAll(server.DB)
	fmt.Println(server.Mock.ExpectationsWereMet())
	if err != nil {
		t.Errorf("this is the error getting the promocodes: %v\n", err)
		return
	}
	assert.Equal(t, len(*pc), 1)
}

func TestPromoCode_Find(t *testing.T) {
	ed := time.Now()
	pcd := PromoCode{
		Title:      "findByIdTest",
		Code:       11,
		IsPercent:  true,
		Discount:   11,
		ExpiryDate: ed,
		Limit:      1,
		Count:      1,
		Active:     true,
	}
	server.Mock.ExpectQuery(regexp.QuoteMeta(`SELECT`)).WillReturnRows(sqlmock.NewRows([]string{"id", "created_at", "updated_at", "title", "code", "is_percent", "discount",
		"expiry_date", "limit", "count", "active"}).AddRow(1, time.Now(), time.Now(), pcd.Title, pcd.Code, pcd.IsPercent, pcd.Discount, pcd.ExpiryDate, pcd.Limit, pcd.Count, pcd.Active))
	server.Mock.ExpectCommit()
	server.Mock.MatchExpectationsInOrder(false)
	fpc, err := promoCodeInstance.Find(server.DB, 9)
	fmt.Println(fpc)
	if err != nil {
		t.Errorf("this is the error getting one PromoCode: %v\n", err)
		return
	}
	assert.NotNil(t, fpc)
	assert.NoError(t, err)
}

func TestPromoCode_Delete(t *testing.T) {
	ed := time.Now()
	server.Mock.ExpectBegin()
	server.Mock.ExpectQuery(regexp.QuoteMeta(`SELECT`)).WillReturnRows(
		sqlmock.NewRows([]string{"id", "title", "code", "is_percent", "discount", "expiry_date", "limit", "count",
			"active", "created_at", "updated_at"}).
			AddRow(1, "test", 111, true, 11, ed, 3, 1, true, time.Now(), time.Now()))
	server.Mock.ExpectExec(regexp.QuoteMeta(`UPDATE`)).WillReturnResult(
		sqlmock.NewResult(0, 1))
	server.Mock.ExpectCommit()
	server.Mock.MatchExpectationsInOrder(false)
	isDeleted, err := promoCodeInstance.Delete(server.DB)
	if err != nil {
		t.Errorf("this is the error deleting the user: %v\n", err)
		return
	}
	assert.Equal(t, int64(isDeleted.ID), int64(1))
}

func TestPromoCode_Update(t *testing.T) {
	ed := time.Now()
	pc := PromoCode{
		Title:      "test",
		Code:       111,
		IsPercent:  true,
		Discount:   11,
		ExpiryDate: ed,
		Limit:      3,
		Count:      1,
		Active:     true,
	}
	server.Mock.ExpectBegin()
	server.Mock.ExpectQuery(regexp.QuoteMeta(`SELECT`)).WillReturnRows(
		sqlmock.NewRows([]string{"id", "title", "code", "is_percent", "discount", "expiry_date", "limit", "count",
			"active", "created_at", "updated_at"}).
			AddRow(111, "test", 111, true, 11, ed, 3, 1, true, time.Now(), time.Now()))
	server.Mock.ExpectExec(regexp.QuoteMeta(`UPDATE`)).WillReturnResult(
		sqlmock.NewResult(0, 1))
	server.Mock.ExpectQuery(regexp.QuoteMeta(`SELECT`)).WillReturnRows(
		sqlmock.NewRows([]string{"id", "name", "value", "is_percent", "country", "create_at", "updated_at"}).
			AddRow(111, "name", 11, true, "test", time.Now(), time.Now()))
	server.Mock.ExpectCommit()
	server.Mock.MatchExpectationsInOrder(false)
	pc.ID = 1
	upc, err := pc.Update(server.DB)
	if err != nil {
		t.Errorf("this is the error updating the user: %v\n", err)
		return
	}
	assert.Equal(t, upc.Title, pc.Title)
	assert.Equal(t, upc.Code, pc.Code)
	assert.Equal(t, upc.IsPercent, pc.IsPercent)
	assert.Equal(t, upc.Discount, pc.Discount)
	assert.Equal(t, upc.ExpiryDate, pc.ExpiryDate)
	assert.Equal(t, upc.Limit, pc.Limit)
	assert.Equal(t, upc.Count, pc.Count)
	assert.Equal(t, upc.Active, pc.Active)
}
