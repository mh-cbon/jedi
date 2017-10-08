package runtime

import "fmt"

//SQLError embeds its query
type SQLError struct {
	error
	SQL string
}

func (s SQLError) Error() string {
	return fmt.Sprintf("%v\n---\n%v", s.error, s.SQL)
}

//NewSQLError returns nil or an ssql error
func NewSQLError(err error, sql string) error {
	if err == nil {
		return nil
	}
	return SQLError{error: err, SQL: sql}
}
