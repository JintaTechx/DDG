package database

import (
	"database/sql"
	"errors"
	"fmt"

	"github.com/JintaTechx/DDG/tree/main/pkg/models"
)

// undefinedQueryError function to create a new undefined error instance
func undefinedQueryError() error {
	return errors.New("undefined query error, try again or contact your DBA")
}

// Connect function that will make a connection to the database
func Connect(driver string, uri string) (*sql.DB, error) {
	var err error

	db, err := sql.Open(driver, uri)

	if err != nil {
		return nil, errors.New(`error when try to connect to the database"`)
	}

	if err = db.Ping(); err != nil {
		return nil, errors.New("connection was created but ping fail, so no content is accessible")
	}

	return db, nil
}

// GetDatabaseComment returns the database comment string or error case some occur
func GetDatabaseComment(db *sql.DB, database string) (string, error) {
	var desc string
	row := db.QueryRow(selectDatabaseComment, database)

	switch err := row.Scan(&desc); err {
	case nil:
		return desc, nil
	default:
		return "", fmt.Errorf("there no database named %s", database)
	}
}

// GetSchemaComment returns the schema comment string or error case some occur
func GetSchemaComment(db *sql.DB, schema string) (string, error) {
	var desc string
	row := db.QueryRow(selectSchemaComment, schema)
	switch err := row.Scan(&desc); err {
	case nil:
		return desc, nil
	default:
		return "", fmt.Errorf("there no schema named %s", schema)
	}
}

// GetAllTables returns all tables structure and comments of a determined schema or error
// case some occur
func GetAllTables(db *sql.DB, schema string) ([]models.Table, error) {
	var tbl []models.Table

	rows, err := db.Query(selectAllTables, schema)
	if err != nil {
		return nil, undefinedQueryError()
	}

	for rows.Next() {
		var table models.Table

		if err := rows.Scan(&table.Name, &table.Desc); err != nil {
			return nil, undefinedQueryError()
		}

		tbl = append(tbl, table)
	}

	return tbl, nil
}

// GetTableColumns returns the data from the table passed by parameter.
func GetTableColumns(db *sql.DB, schema string, table string) ([]models.Columns, error) {
	var tbl []models.Columns

	rows, err := db.Query(selectTable, schema, table)
	if err != nil {
		return nil, undefinedQueryError()
	}

	for rows.Next() {
		var c models.Columns

		err := rows.Scan(&c.Column, &c.Type, &c.Allow, &c.Comment)
		if err != nil {
			return nil, undefinedQueryError()
		}

		tbl = append(tbl, c)
	}

	return tbl, nil
}
