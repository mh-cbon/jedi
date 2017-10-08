package main

import (
	"os"
	"testing"

	"github.com/mh-cbon/jedi/drivers"
	jedi "github.com/mh-cbon/jedi/runtime"
)

func TestSchema(t *testing.T) {

	conn := getConnFromEnv(t, true)
	defer conn.Close()
	defer func() {
		if jedi.Runs(drivers.Sqlite) {
			os.Remove(os.Getenv("JDSN"))
		}
	}()

	t.Run("create table", func(t *testing.T) {
		if err := JProductSetup().Create(conn); err != nil {
			t.Fatalf("Create table must not fail, got err=%v", err)
		}
	})
	t.Run("does not fail to create table", func(t *testing.T) {
		if err := JProductSetup().Create(conn); err != nil {
			t.Fatalf("Create table must not fail, got err=%v", err)
		}
	})
	t.Run("drop table", func(t *testing.T) {
		if err := JProductSetup().Drop(conn); err != nil {
			t.Fatalf("Drop table must not fail, got err=%v", err)
		}
	})
	t.Run("does not fail to drop table", func(t *testing.T) {
		if err := JProductSetup().Drop(conn); err != nil {
			t.Fatalf("Drop table must not fail, got err=%v", err)
		}
	})
	t.Skip("create view")
	t.Skip("does not fail to create view")
	t.Skip("drop view")
	t.Skip("does not fail to drop view")
}
