package models

import (
	_ "github.com/jinzhu/gorm/dialects/postgres"
	"gorm.io/gorm"
)

type MyInterface interface {
	Create(value interface{}) *gorm.DB
}

func New(db MyInterface) *Queries {
	return &Queries{db: db}
}

type Queries struct {
	db MyInterface
}

func (q *Queries) WithTx(tx *gorm.DB) *Queries {
	return &Queries{
		db: tx,
	}
}

// func Init() {

// 	var err error
// 	dsn := "host=localhost user=root password=secret dbname=paymentdb port=5432 sslmode=disable TimeZone=Asia/Shanghai"

// 	db, err := gorm.Open(postgres.Open(dsn), &gorm.Config{})
// 	if err != nil {
// 		log.Fatalf("Cannot connect to  database: %v", err)
// 	}
// 	r = Repository{}
// 	r.DB = db
// }
