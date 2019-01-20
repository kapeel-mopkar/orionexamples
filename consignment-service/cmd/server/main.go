package main

import (
	"database/sql"
	"fmt"

	"github.com/carousell/Orion/orion"
	proto "github.com/carousell/orionexamples/consignment-service/pkg/api/v1"
	"github.com/carousell/orionexamples/consignment-service/pkg/service"
	_ "github.com/lib/pq"
)

var db *sql.DB

const (
	dbhost = "DBHOST"
	dbport = "DBPORT"
	dbuser = "DBUSER"
	dbpass = "DBPASS"
	dbname = "DBNAME"
)

func main() {
	initDb()
	defer db.Close()

	server := orion.GetDefaultServer("ShippingService")               //.GetDefaultServer("StringService")
	proto.RegisterShippingServiceServer(service.GetFactory(), server) //RegisterStringServiceOrionServer(service.GetFactory(), server)
	server.Start()
	server.Wait()
}

func initDb() {
	config := dbConfig()
	var err error
	psqlInfo := fmt.Sprintf("host=%s port=%s user=%s password=%s dbname=%s sslmode=disable",
		config[dbhost], config[dbport],
		config[dbuser], config[dbpass], config[dbname])

	db, err = sql.Open("postgres", psqlInfo)
	if err != nil {
		panic(err)
	}
	err = db.Ping()
	if err != nil {
		panic(err)
	}
	fmt.Println("Successfully connected!")
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
	conf[dbhost] = "localhost"
	conf[dbport] = "5432"
	conf[dbuser] = "postgres"
	conf[dbpass] = "postgres"
	conf[dbname] = "shippingapp"
	return conf
}
