package controllers

import (
	"01cloud-payment/internal/models"
	"01cloud-payment/internal/util"
	"database/sql"
	"fmt"
	"log"

	"github.com/DATA-DOG/go-sqlmock"
	"github.com/jinzhu/gorm"
)

var tserver Server

// func TestMain(m *testing.M) {
// 	err := godotenv.Load(os.ExpandEnv("../../app.env"))
// 	if err != nil {
// 		log.Fatalf("Error getting env %v\n", err)
// 	}
// 	Database()
// 	os.Exit(m.Run())
// }

func Database() {
	var err error
	var db *sql.DB
	TestDbDriver := "postgres"
	db, tserver.Mock, err = sqlmock.New()
	if err != nil {
		panic(err)
	}
	conn, err := gorm.Open(TestDbDriver, db)
	if err != nil {
		fmt.Printf("Cannot connect to %s database\n", TestDbDriver)
		log.Fatal("This is the error:", err)
	} else {
		fmt.Printf("We are connected to the %s database\n", TestDbDriver)
	}
	store := models.NewDB(conn)
	tserver.DB = store
	config := util.Config{}
	NewServer(config, store)
}
