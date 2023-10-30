package DBConnect

import (
	"fmt"
	"gorm.io/driver/postgres"
	"gorm.io/gorm"
)

type DBconnect struct {
	Db  *gorm.DB
	Err error
}

var singleConnect DBconnect

func GetDB() DBconnect {
	if singleConnect.Db == nil {
		fmt.Println("Connect to database now.")
		singleConnect.Db, singleConnect.Err = gorm.Open(postgres.New(postgres.Config{
			DSN: "host=localhost user=postgres password=1 dbname=webapi port=5432 sslmode=disable",
		}), &gorm.Config{})
		if singleConnect.Err != nil {
			panic("failed to connect database")
		} else {
			fmt.Println("connect success")
		}
	} else {
		fmt.Println("Database connected.")
	}

	return singleConnect
}
