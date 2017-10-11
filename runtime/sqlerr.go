package runtime

import "fmt"

//SQLError embeds its query
type SQLError struct {
	error
	sql string
}

func (s SQLError) Error() string {
	return fmt.Sprintf("%v\n---\n%v", s.error, s.sql)
}

// GetQuery returns the query.
func (s SQLError) GetQuery() string {
	return s.sql
}

//NewSQLError returns nil or an SQL error
func NewSQLError(err error, sql string) error {
	if err == nil {
		return nil
	}
	return SQLError{error: err, sql: sql}
}

//NoRowsAffected embeds its query
type NoRowsAffected struct {
	sql string
}

// GetQuery returns the query.
func (s NoRowsAffected) GetQuery() string {
	return s.sql
}

func (s NoRowsAffected) Error() string {
	return "the query did not affect any rows"
}

//NewNoRowsAffected returns a no rows affected error
func NewNoRowsAffected(sql string) error {
	return NoRowsAffected{sql: sql}
}
