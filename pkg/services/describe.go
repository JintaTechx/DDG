package services

import (
	"github.com/JintaTechx/DDG/tree/main/pkg/database"
	"github.com/JintaTechx/DDG/tree/main/pkg/models"
)

// Describe function that has the main objective create the describe data structure that
// represents all elements of database at provided schema
func Describe(uri string, db string, schema string) (*models.Describe, error) {
	desc := &models.Describe{}

	conn, err := database.Connect("postgres", uri)
	if err != nil {
		return nil, err
	}
	defer conn.Close()

	// get Database Info
	dbDesc, err := database.GetDatabaseComment(conn, db)
	if err != nil {
		return nil, err
	}

	desc.Database = models.Basic{
		Name: db,
		Desc: dbDesc,
	}

	//get schema info
	scDesc, err := database.GetSchemaComment(conn, schema)
	if err != nil {
		return nil, err
	}

	desc.Schema = models.Basic{
		Name: schema,
		Desc: scDesc,
	}

	tables, err := database.GetAllTables(conn, schema)
	if err != nil {
		return nil, err
	}

	desc.Tables = tables

	for i := range desc.Tables {
		columns, err := database.GetTableColumns(conn, schema, desc.Tables[i].Name)
		if err != nil {
			return nil, err
		}
		desc.Tables[i].Columns = columns
	}

	return desc, nil
}
