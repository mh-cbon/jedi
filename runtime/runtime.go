package runtime

import (
	"fmt"

	"github.com/gocraft/dbr"
	"github.com/mh-cbon/jedi/drivers"
)

// CurrentDriver user has setup at runtime.
var CurrentDriver string

var toSetup []Setupable

// Setup the current driver at runtime
func Setup(conn *dbr.Connection, force bool) error {
	driver := fmt.Sprintf("%T", conn.Driver())
	CurrentDriver = drivers.Drivers[driver] //must panic if not found.
	if force {
		for _, t := range toSetup {
			k := t()
			if err := k.Drop(conn); err != nil {
				return err
			}
			if err := k.Create(conn); err != nil {
				return err
			}
		}
	}
	return nil
}

type Setuper interface {
	Create(db *dbr.Connection) error
	Drop(db *dbr.Connection) error
}

type Setupable func() Setuper

// Setup the current driver at runtime
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
