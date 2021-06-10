package models

import (
	"fmt"
	"regexp"
	"testing"
	"time"

	"github.com/DATA-DOG/go-sqlmock"
	"github.com/stretchr/testify/assert"
)

func TestThresholdCreate(t *testing.T) {
	threshold := PaymentThreshold{
		UserID:         1,
		ThresholdLimit: 2000,
		Email:          "test@gmail.com",
		Active:         true,
	}
	server.Mock.ExpectQuery(regexp.QuoteMeta(
		`INSERT`)).WillReturnRows(sqlmock.NewRows([]string{"id"}).AddRow(1))
	server.Mock.ExpectCommit()
	server.Mock.ExpectBegin()
	server.Mock.MatchExpectationsInOrder(false)
	fmt.Println(server.Mock.ExpectationsWereMet())
	saved, err := server.store.CreateThreshold(threshold)
	if err != nil {
		t.Errorf("this is the error getting the threshold: %v\n", err)
		return
	}
	assert.Equal(t, saved.UserID, threshold.UserID)
	assert.Equal(t, saved.ThresholdLimit, threshold.ThresholdLimit)
	assert.Equal(t, saved.Email, threshold.Email)
	assert.Equal(t, saved.Active, threshold.Active)
}

func TestFindAllThreshold(t *testing.T) {
	// thresholds := []PaymentThreshold{
	// 	{
	// 		UserID:         1,
	// 		ThresholdLimit: 2000,
	// 		Email:          "test@gmail.com",
	// 		Active:         true,
	// 	},
	// 	{
	// 		UserID:         2,
	// 		ThresholdLimit: 1000,
	// 		Email:          "test1@gmail.com",
	// 		Active:         false,
	// 	},
	// }
	//var testthresholds *[]PaymentThreshold
	server.Mock.ExpectQuery(regexp.QuoteMeta(
		`SELECT`)).WillReturnRows(sqlmock.NewRows([]string{"id", "user_id", "thresholdlimit", "email", "active", "created_at", "updated_at"}).AddRow(1, 2, 1000, "test1@gmail.com", false, time.Now(), time.Now()))
	testthresholds, err := server.store.FindAllThreshold()
	fmt.Println(server.Mock.ExpectationsWereMet())

	if err != nil {
		t.Errorf("this is the error getting the users: %v\n", err)
		return
	}
	assert.Equal(t, len(testthresholds), 1)
}

func TestFindThresholdByID(t *testing.T) {
	threshold := PaymentThreshold{
		UserID:         2,
		ThresholdLimit: 1000,
		Email:          "test1@gmail.com",
		Active:         false,
	}
	threshold.ID = 1
	server.Mock.ExpectQuery(regexp.QuoteMeta(`SELECT`)).WillReturnRows(sqlmock.NewRows([]string{"id", "user_id", "thresholdlimit", "email", "active", "created_at", "updated_at"}).AddRow(1, 2, 1000, "test1@gmail.com", false, time.Now(), time.Now()))
	server.Mock.ExpectCommit()
	server.Mock.MatchExpectationsInOrder(false)
	threshold0, err := server.store.FindByIdThreshold(threshold.ID)
	if err != nil {
		t.Errorf("this is the error getting one user: %v\n", err)
		return
	}
	assert.Equal(t, threshold0.ID, threshold.ID)
	assert.Equal(t, threshold0.UserID, threshold.UserID)
	assert.Equal(t, threshold0.Email, threshold.Email)
}

func TestUpdateThreshold(t *testing.T) {
	threshold := PaymentThreshold{
		UserID:         1,
		ThresholdLimit: 2000,
		Email:          "test@gmail.com",
		Active:         true,
	}
	server.Mock.ExpectBegin()
	server.Mock.ExpectQuery(regexp.QuoteMeta(
		`SELECT`)).WillReturnRows(sqlmock.NewRows([]string{"id", "user_id", "thresholdlimit", "email", "active", "CreatedAt", "UpdatedAt"}).AddRow(1, threshold.UserID, threshold.ThresholdLimit, threshold.Email, threshold.Active, time.Now(), time.Now()))
	server.Mock.ExpectExec(regexp.QuoteMeta(
		`UPDATE`)).WillReturnResult(sqlmock.NewResult(0, 1))
	server.Mock.ExpectQuery(regexp.QuoteMeta(
		`SELECT`)).WillReturnRows(sqlmock.NewRows([]string{"id", "user_id", "thresholdlimit", "email", "active", "CreatedAt", "UpdatedAt"}).AddRow(1, 2, 2000, "myemail@gmail.com", true, time.Now(), time.Now()))
	server.Mock.ExpectCommit()
	server.Mock.MatchExpectationsInOrder(false)
	threshold.ID = 1
	updatedThreshold, err := server.store.UpdateThreshold(threshold)
	if err != nil {
		t.Errorf("this is the error updating the user: %v\n", err)
		return
	}
	assert.Equal(t, updatedThreshold.UserID, threshold.UserID)
	assert.Equal(t, updatedThreshold.Email, threshold.Email)
	assert.Equal(t, updatedThreshold.Active, threshold.Active)
}

func TestDeleteThreshold(t *testing.T) {
	threshold := PaymentThreshold{
		UserID:         1,
		ThresholdLimit: 2000,
		Email:          "test@gmail.com",
		Active:         true,
	}
	server.Mock.ExpectBegin()
	server.Mock.ExpectQuery(regexp.QuoteMeta(
		`SELECT`)).WillReturnRows(sqlmock.NewRows([]string{"id", "user_id", "thresholdlimit", "email", "active", "CreatedAt", "UpdatedAt"}).AddRow(1, threshold.UserID, threshold.ThresholdLimit, threshold.Email, threshold.Active, time.Now(), time.Now()))
	server.Mock.ExpectExec(regexp.QuoteMeta(
		`UPDATE`)).WillReturnResult(sqlmock.NewResult(0, 1))
	server.Mock.ExpectCommit()
	server.Mock.MatchExpectationsInOrder(false)
	id, err := server.store.DeleteThreshold(threshold.ID)
	if err != nil {
		t.Errorf("this is the error updating the user: %v\n", err)
		return
	}
	assert.Equal(t, int64(id), int64(1))
}
