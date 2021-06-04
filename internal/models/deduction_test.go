package models

import (
	"fmt"
	"regexp"
	"testing"
	"time"

	"github.com/DATA-DOG/go-sqlmock"
	"github.com/stretchr/testify/assert"
)

var deductionInstance = Deduction{}

func TestDeduction_Save(t *testing.T) {
	ed := time.Now()
	nd := Deduction{
		Name:          "Tax",
		Value:         11,
		IsPercent:     true,
		Country:       "Nepal",
		Description:   "test description",
		Attributes:    "test attributes",
		EffectiveDate: ed,
		//StartDate: time.Date(time.Now().Year(), time.Now().Month(), 1, 0, 0, 0, 0, time.Local),
	}
	server.Mock.ExpectQuery(regexp.QuoteMeta(
		`INSERT`)).WillReturnRows(sqlmock.NewRows([]string{"id"}).AddRow(1))
	server.Mock.ExpectCommit()
	server.Mock.ExpectBegin()
	server.Mock.MatchExpectationsInOrder(false)
	fmt.Println(server.Mock.ExpectationsWereMet())

	saved, err := nd.Save(server.DB)
	if err != nil {
		t.Errorf("this is the error getting the Deductions: %v\n", err)
		return
	}
	assert.Equal(t, saved.Name, nd.Name)
	assert.Equal(t, saved.Value, nd.Value)
	assert.Equal(t, saved.IsPercent, nd.IsPercent)
	assert.Equal(t, saved.Country, nd.Country)
	assert.Equal(t, saved.Description, nd.Description)
	assert.Equal(t, saved.Attributes, nd.Attributes)
	assert.Equal(t, saved.EffectiveDate, nd.EffectiveDate)
}

func TestGateway_FindAll(t *testing.T) {
	ed := time.Now()
	server.Mock.ExpectQuery(regexp.QuoteMeta(`SELECT`)).WillReturnRows(
		sqlmock.NewRows([]string{"id", "created_at", "updated_at", "name", "value", "is_percent", "country", "description", "attributes", "effective_date"}).
			AddRow(1, time.Now(), time.Now(), "test", 11, true, "test", "test description", "test attributes", ed))
	d, err := deductionInstance.FindAll(server.DB)
	fmt.Println(server.Mock.ExpectationsWereMet())
	if err != nil {
		t.Errorf("this is the error getting the deductions: %v\n", err)
		return
	}
	assert.Equal(t, len(*d), 1)
}

func TestDeduction_Find(t *testing.T) {
	ed := time.Now()
	d := Deduction{
		Name:          "test",
		Value:         11,
		IsPercent:     true,
		Country:       "test",
		Description:   "test description",
		Attributes:    "test attributes",
		EffectiveDate: ed,
	}

	server.Mock.ExpectQuery(regexp.QuoteMeta(`SELECT`)).WillReturnRows(
		sqlmock.NewRows([]string{"id", "name", "value", "is_percent", "country", "description", "attributes", "effective_date", "created_at", "updated_at"}).
			AddRow(1, "test", 11, true, "test", "test description", "test attributes", ed, time.Now(), time.Now()))
	server.Mock.ExpectCommit()
	server.Mock.MatchExpectationsInOrder(false)
	fd, err := deductionInstance.Find(server.DB, 1)
	if err != nil {
		t.Errorf("this is the error getting one deduction %v\n", err)
		return
	}
	assert.Equal(t, fd.Name, d.Name)
	assert.Equal(t, fd.Value, d.Value)
	assert.Equal(t, fd.IsPercent, d.IsPercent)
	assert.Equal(t, fd.Country, d.Country)
	assert.Equal(t, fd.Description, d.Description)
	assert.Equal(t, fd.Attributes, d.Attributes)
	assert.Equal(t, fd.EffectiveDate, d.EffectiveDate)

}

func TestDeduction_Delete(t *testing.T) {
	ed := time.Now()
	server.Mock.ExpectBegin()
	server.Mock.ExpectQuery(regexp.QuoteMeta(`SELECT`)).WillReturnRows(
		sqlmock.NewRows([]string{"id", "name", "value", "is_percent", "country", "description", "attributes", "effective_date", "created_at", "updated_at"}).
			AddRow(1, "test", 11, true, "test", "test description", "attributes", ed, time.Now(), time.Now()))
	server.Mock.ExpectExec(regexp.QuoteMeta(`UPDATE`)).WillReturnResult(
		sqlmock.NewResult(0, 1))
	server.Mock.ExpectCommit()
	server.Mock.MatchExpectationsInOrder(false)
	isDeleted, err := deductionInstance.Delete(server.DB)
	if err != nil {
		t.Errorf("this is the error deleting the user: %v\n", err)
		return
	}
	assert.Equal(t, int64(isDeleted.ID), int64(0))
}

func TestDeduction_Update(t *testing.T) {
	ed := time.Now()
	d := Deduction{
		Name:          "test",
		Value:         11,
		IsPercent:     true,
		Country:       "test",
		Description:   "test description",
		Attributes:    "test attributes",
		EffectiveDate: ed,
	}
	server.Mock.ExpectBegin()
	server.Mock.ExpectQuery(regexp.QuoteMeta(`SELECT`)).WillReturnRows(
		sqlmock.NewRows([]string{"id", "name", "value", "is_percent", "country", "description", "attributes", "effective_date", "create_at", "updated_at"}).
			AddRow(1, "test", 11, true, "test", "test description", "test attributes", ed, time.Now(), time.Now()))
	server.Mock.ExpectExec(regexp.QuoteMeta(`UPDATE`)).WillReturnResult(
		sqlmock.NewResult(0, 1))
	server.Mock.ExpectQuery(regexp.QuoteMeta(`SELECT`)).WillReturnRows(
		sqlmock.NewRows([]string{"id", "name", "value", "is_percent", "country", "create_at", "description", "attributes", "effective_date", "updated_at"}).
			AddRow(1, "name", 11, true, "test", "test description", "test attributes", ed, time.Now(), time.Now()))
	server.Mock.ExpectCommit()
	server.Mock.MatchExpectationsInOrder(false)
	d.ID = 1
	ud, err := d.Update(server.DB)
	if err != nil {
		t.Errorf("this is the error updating the user: %v\n", err)
		return
	}
	assert.Equal(t, ud.ID, d.ID)
	assert.Equal(t, ud.Name, d.Name)
	assert.Equal(t, ud.Value, d.Value)
	assert.Equal(t, ud.IsPercent, d.IsPercent)
	assert.Equal(t, ud.Country, d.Country)
	assert.Equal(t, ud.Description, d.Description)
	assert.Equal(t, ud.Attributes, d.Attributes)
	assert.Equal(t, ud.EffectiveDate, d.EffectiveDate)
}
