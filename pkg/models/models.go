package models

// Basic data structure to define the fields that just has name anda description
type Basic struct {
	Name string `json:"name"`
	Desc string `json:"description"`
}

// Columns data structure to define each columns of table
type Columns struct {
	Column  string `json:"column"`
	Type    string `json:"type"`
	Allow   string `json:"allow"`
	Comment string `json:"comment"`
}

// Table data structure that defines a table
type Table struct {
	Name    string    `json:"name"`
	Desc    string    `json:"description"`
	Columns []Columns `json:"columns"`
}

// Describe is the data structure description to the database and schema
type Describe struct {
	Database Basic   `json:"database"`
	Schema   Basic   `json:"schema"`
	Tables   []Table `json:"tables"`
}
