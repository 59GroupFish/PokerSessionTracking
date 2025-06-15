package db

import (
	"fmt"
	"log"
	"os"
	"reflect"

	"github.com/jmoiron/sqlx"
)

var DB *sqlx.DB

func InitDB() error {
	connStr := os.Getenv("DATABASE_CONN")
	if connStr == "" {
		return fmt.Errorf("DATABASE_CONN environment variable is not set")
	}

	var err error
	DB, err = sqlx.Open("pgx", connStr)
	if err != nil {
		return fmt.Errorf("error opening database connection: %w", err)
	}

	err = DB.Ping()
	if err != nil {
		// If ping fails, close the connection to release resources.
		DB.Close()
		return fmt.Errorf("error pinging database: %w", err)
	}

	log.Println("Successfully connected to the PostgreSQL database!")
	return nil
}

// Get the value of the db tag from a struct, can specify a prefix and suffix to prevent column name colliding with key words of sql
//
//	"ColumnName" for postgres
//	[ColumnName] for mssql
//
// asPrefix will be appended to the start of the column name
//
//	e.g. asPrefix = "xyz"
//	"ColumnName AS xyz_ColumnName"
func GetColumnNamesFromStruct(v interface{}, ignoreColumns []string, headingOpen string, headingClose string, asPrefix string) ([]string, error) {
	var columns []string
	val := reflect.ValueOf(v)
	if val.Kind() == reflect.Ptr {
		val = val.Elem()
	}

	typ := val.Type()
	for i := 0; i < val.NumField(); i++ {
		field := typ.Field(i)
		tag := field.Tag.Get("db")
		if tag != "" && !contains(ignoreColumns, tag) {
			column := headingOpen + tag + headingClose
			if asPrefix != "" {
				column = column + " AS \"" + asPrefix + tag + "\""
			}
			columns = append(columns, column)
		}
	}
	return columns, nil
}

func contains(slice []string, item string) bool {
	for _, s := range slice {
		if s == item {
			return true
		}
	}
	return false
}
