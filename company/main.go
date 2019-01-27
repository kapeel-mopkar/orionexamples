package main

import (
	"database/sql"
	"time"

	"github.com/jinzhu/gorm"
	_ "github.com/lib/pq"
)

type Company struct {
	ID        int        `sql:"AUTO_INCREMENT" gorm:"primary_key"`
	Name      string     `sql:"size:255;unique;index"`
	Employees []Employee // one-to-many relationship
	Address   Address    // one-to-one relationship
}

type Employee struct {
	FirstName        string    `sql:"size:255;index:name_idx"`
	LastName         string    `sql:"size:255;index:name_idx"`
	SocialSecurityNo string    `sql:"type:varchar(100);unique" gorm:"column:ssn"`
	DateOfBirth      time.Time `sql:"DEFAULT:current_timestamp"`
	Address          *Address  // one-to-one relationship
	Deleted          bool      `sql:"DEFAULT:false"`
}

type Address struct {
	Country  string `gorm:"primary_key"`
	City     string `gorm:"primary_key"`
	PostCode string `gorm:"primary_key"`
	Line1    sql.NullString
	Line2    sql.NullString
}

func main() {
	// conf[dbhost] = "localhost"
	// conf[dbport] = "5432"
	// conf[dbuser] = "postgres"
	// conf[dbpass] = "pwd123"
	// conf[dbname] = "postgres"  host=%s port=%s user=%s password=%s dbname=%s sslmode=disable
	db, err := gorm.Open("postgres", "host=localhost port=5432 user=postgres password=pwd123 dbname=shippingapp sslmode=disable")
	if err != nil {
		panic(err)
	}

	// Ping function checks the database connectivity
	err = db.DB().Ping()
	if err != nil {
		panic(err)
	}

	if !db.HasTable(&Address{}) {
		db.CreateTable(&Address{})
	}

	if !db.HasTable(&Company{}) {
		db.CreateTable(&Company{})
	}

	if !db.HasTable(&Employee{}) {
		db.CreateTable(&Employee{})
	}

}
