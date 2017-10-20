package runtime

import "github.com/gocraft/dbr"

// Table provides statements to create/drop table.
type Table struct {
	Name       string
	CreateStmt string
	DropStmt   string
	View       bool
	Indexes    []string
}

//IndexStatements returns the index string statements.
func (c Table) IndexStatements() []string {
	return c.Indexes
}

//CreateStatement returns the string.
func (c Table) CreateStatement() string {
	return c.CreateStmt
}

//DropStatement returns the string.
func (c Table) DropStatement() string {
	return c.DropStmt
}

//Create applies the create table command to te underlying connection.
func (c Table) Create(db *dbr.Connection) error {
	_, err := db.Exec(c.CreateStmt)
	return NewSQLError(err, c.CreateStmt)
}

//CreateIndexes applies the indexes on the table.
func (c Table) CreateIndexes(db *dbr.Connection) error {
	for _, i := range c.Indexes {
		_, err := db.Exec(i)
		if err != nil {
			return NewSQLError(err, i)
		}
	}
	return nil
}

//Drop applies the drop table command to te underlying connection.
func (c Table) Drop(db *dbr.Connection) error {
	_, err := db.Exec(c.DropStmt)
	return NewSQLError(err, c.DropStmt)
}

//IsView returns true if it is a view.
func (c Table) IsView() bool {
	return c.View
}
