package models

import (
	"fmt"
	"regexp"
	"testing"
	"time"

	"github.com/DATA-DOG/go-sqlmock"
	"github.com/stretchr/testify/assert"
)

func TestPromoCodeCreate(t *testing.T) {
	ed := time.Now()
	promocode := PromoCode{
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
	saved, err := server.store.CreatePromocode(promocode)
	if err != nil {
		t.Errorf("this is the error getting the PromoCode: %v\n", err)
		return
	}
	assert.Equal(t, saved.Title, promocode.Title)
	assert.Equal(t, saved.Code, promocode.Code)
	assert.Equal(t, saved.IsPercent, promocode.IsPercent)
	assert.Equal(t, saved.Discount, promocode.Discount)
	assert.Equal(t, saved.ExpiryDate, promocode.ExpiryDate)
	assert.Equal(t, saved.Limit, promocode.Limit)
	assert.Equal(t, saved.Count, promocode.Count)
	assert.Equal(t, saved.Active, promocode.Active)
}

func TestPromoCode_FindAll(t *testing.T) {
	ed := time.Now()
	// promocodes := []PromoCode{
	// 	{
	// 		Title:      "test",
	// 		Code:       111,
	// 		IsPercent:  true,
	// 		ExpiryDate: ed,
	// 		Discount:   11,
	// 		Limit:      3,
	// 		Count:      1,
	// 		Active:     true,
	// 	},
	// 	{
	// 		Title:      "test1",
	// 		Code:       1111,
	// 		IsPercent:  true,
	// 		ExpiryDate: ed,
	// 		Discount:   10,
	// 		Limit:      2,
	// 		Count:      1,
	// 		Active:     false,
	// 	},
	// }

	server.Mock.ExpectQuery(regexp.QuoteMeta(`SELECT`)).WillReturnRows(
		sqlmock.NewRows([]string{"id", "created_at", "updated_at", "title", "code", "is_percent", "discount",
			"expiry_date", "limit", "count", "active"}).AddRow(11, time.Now(), time.Now(), "test", 111, true, 11,
			ed, 3, 1, true))
	pc, err := server.store.FindAllPromocode()
	fmt.Println(server.Mock.ExpectationsWereMet())
	if err != nil {
		t.Errorf("this is the error getting the promocodes: %v\n", err)
		return
	}
	assert.Equal(t, len(pc), 1)
}

func TestPromoCode_Find(t *testing.T) {
	ed := time.Now()
	promocode := PromoCode{
		Title:      "findByIdTest",
		Code:       11,
		IsPercent:  true,
		Discount:   11,
		ExpiryDate: ed,
		Limit:      1,
		Count:      1,
		Active:     true,
	}
	promocode.ID = 1
	server.Mock.ExpectQuery(regexp.QuoteMeta(`SELECT`)).WillReturnRows(sqlmock.NewRows([]string{"id", "created_at", "updated_at", "title", "code", "is_percent", "discount",
		"expiry_date", "limit", "count", "active"}).AddRow(1, time.Now(), time.Now(), promocode.Title, promocode.Code, promocode.IsPercent, promocode.Discount, promocode.ExpiryDate, promocode.Limit, promocode.Count, promocode.Active))
	server.Mock.ExpectCommit()
	server.Mock.MatchExpectationsInOrder(false)
	fpc, err := server.store.FindByIdPromocode(promocode.ID)
	if err != nil {
		t.Errorf("this is the error getting one PromoCode: %v\n", err)
		return
	}
	assert.NotNil(t, fpc)
	assert.NoError(t, err)
}

func TestPromoCode_Delete(t *testing.T) {
	ed := time.Now()
	promocode := PromoCode{
		Title:      "findByIdTest",
		Code:       11,
		IsPercent:  true,
		Discount:   11,
		ExpiryDate: ed,
		Limit:      1,
		Count:      1,
		Active:     true,
	}
	promocode.ID = 1
	server.Mock.ExpectBegin()
	server.Mock.ExpectQuery(regexp.QuoteMeta(`SELECT`)).WillReturnRows(
		sqlmock.NewRows([]string{"id", "title", "code", "is_percent", "discount", "expiry_date", "limit", "count",
			"active", "created_at", "updated_at"}).
			AddRow(1, "test", 111, true, 11, ed, 3, 1, true, time.Now(), time.Now()))
	server.Mock.ExpectExec(regexp.QuoteMeta(`UPDATE`)).WillReturnResult(
		sqlmock.NewResult(0, 1))
	server.Mock.ExpectCommit()
	server.Mock.MatchExpectationsInOrder(false)
	isDeleted, err := server.store.DeletePromocode(promocode.ID)
	if err != nil {
		t.Errorf("this is the error deleting the user: %v\n", err)
		return
	}
	assert.Equal(t, isDeleted, int64(1))
}

func TestPromoCode_Update(t *testing.T) {
	ed := time.Now()
	promocode := PromoCode{
		Title:      "test",
		Code:       111,
		IsPercent:  true,
		Discount:   11,
		ExpiryDate: ed,
		Limit:      3,
		Count:      1,
		Active:     true,
	}
	promocode.ID = 1
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
	upc, err := server.store.UpdatePromocode(promocode)
	if err != nil {
		t.Errorf("this is the error updating the user: %v\n", err)
		return
	}
	assert.Equal(t, upc.Title, promocode.Title)
	assert.Equal(t, upc.Code, promocode.Code)
	assert.Equal(t, upc.IsPercent, promocode.IsPercent)
	assert.Equal(t, upc.Discount, promocode.Discount)
	assert.Equal(t, upc.ExpiryDate, promocode.ExpiryDate)
	assert.Equal(t, upc.Limit, promocode.Limit)
	assert.Equal(t, upc.Count, promocode.Count)
	assert.Equal(t, upc.Active, promocode.Active)
}
