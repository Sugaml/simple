package models

import (
	"database/sql"
	"fmt"
	"log"
	"os"
	"testing"

	"github.com/DATA-DOG/go-sqlmock"
	"github.com/jinzhu/gorm"
	"github.com/joho/godotenv"
)

type TestServer struct {
	DB    *gorm.DB
	Mock  sqlmock.Sqlmock
	store Store
}

var server TestServer
var invoiceInstance = Invoice{}

func TestMain(m *testing.M) {
	err = godotenv.Load(os.ExpandEnv("../../app.env"))
	if err != nil {
		log.Fatalf("Error getting env %v\n", err)
	}
	Database()
	os.Exit(m.Run())
}

func Database() {
	var err error
	var db *sql.DB
	TestDbDriver := os.Getenv("TestDbDriver")
	db, server.Mock, err = sqlmock.New()
	if err != nil {
		panic(err)
	}
	server.DB, err = gorm.Open(TestDbDriver, db)
	if err != nil {
		fmt.Printf("Cannot connect to %s database\n", TestDbDriver)
		log.Fatal("This is the error:", err)
	} else {
		fmt.Printf("We are connected to the %s database\n", TestDbDriver)
	}
	server.store = NewDB(server.DB)
}
