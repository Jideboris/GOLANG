package sql_config

import (
	_ "github.com/go-sql-driver/mysql" 
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