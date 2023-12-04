package sql_config

import (
	"context"
	"database/sql"
	"fmt"
	_ "github.com/go-sql-driver/mysql"
	"log"
	"time"
)

const (
	username = "gotest"
	password = "1234"
	hostname = "localhost:3306"
	dbname   = "questionnairesDB"
)

type Sqlconfiguration struct {
	DBHost                        string
	DBPort                        string
	DBUserName                    string
	DBPass                        string
	DBDatabaseName                string
	maxIdleConnection             int
	maxOpenConnection             int
	connectionMaxLifetimeInSecond int
}
type mysqlOption func(*Sqlconfiguration)

func Connect(DBHost string,
	DBPort string,
	DBUserName string,
	DBPass string,
	DBDatabaseName string) *Sqlconfiguration {
	configuration := &Sqlconfiguration{
		DBHost:                        DBHost,
		DBPort:                        DBPort,
		DBUserName:                    DBUserName,
		DBPass:                        DBPass,
		DBDatabaseName:                DBDatabaseName,
		maxIdleConnection:             5,
		maxOpenConnection:             10,
		connectionMaxLifetimeInSecond: 60,
	}
	return configuration
}

func dsn(dsName string, param *Sqlconfiguration) string {
	return fmt.Sprintf("%s:%s@tcp(%s)/%s", username, password, hostname, dbname)
}

func connect(param *Sqlconfiguration) (*sql.DB, error) {
	db, err := sql.Open("mysql", dsn("", param))

	if err != nil {
		log.Printf("Error %s when opening DB\n", err)
		return nil, err
	}
	_, cancelfunc := context.WithTimeout(context.Background(), 5*time.Second)
	defer cancelfunc()

	fmt.Println(">>>>>>>>>>>>>>>>>>>>>>>>>>>>>>>>>>>>>>>>>", dsn("", param))
	res, err := db.ExecContext(context.Background(), "CREATE DATABASE IF NOT EXISTS "+dbname)
	if err != nil {
		log.Printf("Error %s when creating DBOMO\n", err)
		return nil, err
	}
	no, err := res.RowsAffected()
	if err != nil {
		log.Printf("Error %s when fetching rows", err)
		return nil, err
	}
	log.Printf("rows affected: %d\n", no)
	defer db.Close()

	db, err = sql.Open("mysql", dsn("Survey", param))
	if err != nil {
		log.Printf("Error %s when opening DB with database", err)
		return nil, err
	}

	// set configuration pooling connection
	db.SetMaxOpenConns(param.maxOpenConnection)
	db.SetConnMaxLifetime(time.Duration(param.connectionMaxLifetimeInSecond) * time.Minute)
	db.SetMaxIdleConns(param.maxIdleConnection)

	// fmt.Println("HERE >>>>>>>>>>>>>>>>>>>>>>>>>>",db)
	// var version string
	// err2 := db.QueryRow("SELECT VERSION()").Scan(&version)
	// if err2 != nil {
	// 	log.Fatal(err2)
	// }
	// ctx, cancelfunc = context.WithTimeout(context.Background(), 5*time.Second)
	// defer cancelfunc()
	// err = db.PingContext(ctx)
	// if err != nil {
	// 	log.Printf("Errors %s pinging DB", err)
	// 	return nil, err
	// }

	log.Printf("Connected to DB %s successfully\n", param.DBDatabaseName)

	return db, nil
}

func SetMaxIdleConns(conns int) mysqlOption {
	return func(c *Sqlconfiguration) {
		if conns > 0 {
			c.maxIdleConnection = conns
		}
	}
}

func SetMaxOpenConns(conns int) mysqlOption {
	return func(c *Sqlconfiguration) {
		if conns > 0 {
			c.maxOpenConnection = conns
		}
	}
}

func SetConnMaxLifetime(conns int) mysqlOption {
	return func(c *Sqlconfiguration) {
		if conns > 0 {
			c.connectionMaxLifetimeInSecond = conns
		}
	}
}
