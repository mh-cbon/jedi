package builder

import (
	"fmt"

	"github.com/gocraft/dbr"
)

// Like is `Like` condition.
func Like(column string, value interface{}) dbr.Builder {
	return dbr.BuildFunc(func(d dbr.Dialect, buf dbr.Buffer) error {
		buf.WriteString(d.QuoteIdent(column))
		buf.WriteString(" ")
		buf.WriteString("LIKE")
		buf.WriteString(" ")
		buf.WriteString(fmt.Sprintf(`'%%%v%%'`, value)) //todo: make secure, add tests
		return nil
	})
}
