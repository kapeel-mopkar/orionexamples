package service

import (
	"fmt"
	"log"

	"github.com/carousell/Orion/orion"

	"github.com/jinzhu/gorm"
	_ "github.com/jinzhu/gorm/dialects/postgres" // Need for connecting to POSTGRES DB
)

type svcFactory struct{}

var db *gorm.DB

const (
	dbhost = "DBHOST"
	dbport = "DBPORT"
	dbuser = "DBUSER"
	dbpass = "DBPASS"
	dbname = "DBNAME"
)

func (s *svcFactory) NewService(orion.Server) interface{} {
	initDb()
	db.AutoMigrate(&Consignment{})
	return &svc{
		db: db,
	}
}

func (s *svcFactory) DisposeService(svc interface{}) {
	log.Println("Disposing SVCFactory service")
}

func GetFactory() orion.ServiceFactory {
	return &svcFactory{}
}

func initDb() {
	config := dbConfig()
	var err error
	psqlInfo := fmt.Sprintf("host=%s port=%s user=%s password=%s dbname=%s sslmode=disable",
		config[dbhost], config[dbport],
		config[dbuser], config[dbpass], config[dbname])

	db, err = gorm.Open("postgres", psqlInfo)
	if err != nil {
		panic(err)
	}
	log.Println("Successfully connected!")

	if !db.HasTable(&Consignment{}) {
		db.CreateTable(&Consignment{})
	}
}

func dbConfig() map[string]string {
	conf := make(map[string]string)
	// host, ok := os.LookupEnv(dbhost)
	// if !ok {
	// 	panic("DBHOST environment variable required but not set")
	// }
	// port, ok := os.LookupEnv(dbport)
	// if !ok {
	// 	panic("DBPORT environment variable required but not set")
	// }
	// user, ok := os.LookupEnv(dbuser)
	// if !ok {
	// 	panic("DBUSER environment variable required but not set")
	// }
	// password, ok := os.LookupEnv(dbpass)
	// if !ok {
	// 	panic("DBPASS environment variable required but not set")
	// }
	// name, ok := os.LookupEnv(dbname)
	// if !ok {
	// 	panic("DBNAME environment variable required but not set")
	// }
	conf[dbhost] = "localhost"//"localhost"
	conf[dbport] = "5432"
	conf[dbuser] = "postgres"
	conf[dbpass] = "postgres"
	conf[dbname] = "shippingapp"
	return conf
}
