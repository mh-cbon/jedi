package runtime

import (
	"fmt"

	"github.com/gocraft/dbr"
	"github.com/mh-cbon/jedi/drivers"
)

// Rand is `Rand()` condition.
func Rand() string {
	if Runs(drivers.Mysql) {
		return "RAND()"
	}
	return "RANDOM()"
}

// Like is `Like` condition.
func Like(column string, value interface{}) dbr.Builder {
	return dbr.BuildFunc(func(d dbr.Dialect, buf dbr.Buffer) error {
		buf.WriteString(d.QuoteIdent(column))
		buf.WriteString(" ")
		buf.WriteString("LIKE")
		buf.WriteString(" ")
		buf.WriteString(fmt.Sprintf(`'%v'`, value)) //todo: make secure, add tests
		return nil
	})
}

// IsNull is `IS NULL` condition.
func IsNull(column string) dbr.Builder {
	return dbr.BuildFunc(func(d dbr.Dialect, buf dbr.Buffer) error {
		buf.WriteString(d.QuoteIdent(column))
		buf.WriteString(" ")
		buf.WriteString("IS NULL")
		return nil
	})
}

// IsNotNull is `IS NOT NULL` condition.
func IsNotNull(column string) dbr.Builder {
	return dbr.BuildFunc(func(d dbr.Dialect, buf dbr.Buffer) error {
		buf.WriteString(d.QuoteIdent(column))
		buf.WriteString(" ")
		buf.WriteString("IS NOT NULL")
		return nil
	})
}
