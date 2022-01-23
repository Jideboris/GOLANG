package main

import (
	
)

type DatabaseConnection struct {
	Host     string
	Port     string
	Database string
	Username string
	Password string
}

func (db DatabaseConnection) connect() {
	db, err := sql.Open("mysql", credentials.MongoDbDsn())

	if err != nil {
		panic(err.Error())
	}
}
