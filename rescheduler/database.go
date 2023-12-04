 package main

// import (
// 	"database/sql"
// 	"fmt"
// 	"log"
// )

// const (
// 	username = "root"
// 	password = "password"
// 	hostname = "127.0.0.1:3306"
// 	dbname   = "ecommerce"
// )

// type DatabaseConnection struct {
// 	Host     string
// 	Port     string
// 	Database string
// 	Username string
// 	Password string
// }

// func MongoDbDsn(dbName string) string {
// 	return fmt.Sprintf("%s:%s@tcp(%s)/%s", username, password, hostname, dbName)
// }

// func (db DatabaseConnection) connect() {
// 	dbb, err := sql.Open("mysql", "user7:s$cret@tcp(127.0.0.1:3306)/testdb")
//     defer db.Close()

//     if err != nil {
//         log.Fatal(err)
//     }

//     var version string

//     err2 := db.QueryRow("SELECT VERSION()").Scan(&version)

//     if err2 != nil {
//         log.Fatal(err2)
//     }

//     fmt.Println(version)
// }
