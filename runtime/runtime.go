package runtime

import (
	"fmt"
	"io"

	"github.com/gocraft/dbr"
	"github.com/gocraft/dbr/dialect"
	"github.com/mh-cbon/jedi/drivers"
)

// CurrentDriver user has setup at runtime.
var CurrentDriver string

// Open a connection, shorthand for dbr.Open
func Open(driver, dsn string, log dbr.EventReceiver, force bool) (*dbr.Connection, error) {
	conn, err := dbr.Open(driver, dsn, log)
	if err != nil {
		return conn, err
	}
	return conn, Setup(conn, force)
}

var toSetup []Setupable

// Setup the current driver at runtime
func Setup(conn *dbr.Connection, force bool) error {
	driver := fmt.Sprintf("%T", conn.Driver())
	CurrentDriver = drivers.Drivers[driver] //must panic if not found.
	if force {
		for _, t := range toSetup {
			k := t()
			if k.IsView() {
				if err := k.Drop(conn); err != nil {
					return err
				}
			}
		}
		for _, t := range toSetup {
			k := t()
			if !k.IsView() {
				if err := k.Drop(conn); err != nil {
					return err
				}
				if err := k.Create(conn); err != nil {
					return err
				}
			}
		}
		for _, t := range toSetup {
			k := t()
			if k.IsView() {
				if err := k.Create(conn); err != nil {
					return err
				}
			}
		}
	}
	return nil
}

// Setuper can create/drop tables.
type Setuper interface {
	IsView() bool
	Create(db *dbr.Connection) error
	Drop(db *dbr.Connection) error
}

// Setupable returns a Setuper
type Setupable func() Setuper

// Register the current driver at runtime
func Register(m ...Setupable) {
	toSetup = append(toSetup, m...)
}

// GetCurrentDriver panics if the driver is not setup
func GetCurrentDriver() string {
	if CurrentDriver == "" {
		panic("Did you configured jedi ? See drivers.Setup()")
	}
	return CurrentDriver
}

// Runs returns true if the current connection is of given type.
func Runs(t ...string) bool {
	for _, tt := range t {
		if tt == CurrentDriver {
			return true
		}
	}
	return false
}

// GetDialect returns current dbr dialct
func GetDialect() dbr.Dialect {
	switch GetCurrentDriver() {
	case drivers.Sqlite:
		return dialect.SQLite3
	case drivers.Mysql:
		return dialect.MySQL
	case drivers.Pgsql:
		return dialect.PostgreSQL
	}
	return nil
}

//DumpSchema generated automatically.
func DumpSchema(out io.Writer) {

}
