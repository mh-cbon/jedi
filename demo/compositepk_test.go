package main

import (
	"os"
	"strings"
	"testing"

	"github.com/mh-cbon/jedi/drivers"
	jedi "github.com/mh-cbon/jedi/runtime"
)

func TestCompositePk(t *testing.T) {

	conn := getConnFromEnv(t, true)
	defer conn.Close()
	defer func() {
		if jedi.Runs(drivers.Sqlite) {
			os.Remove(os.Getenv("JDSN"))
		}
	}()

	sess := conn.NewSession(nil)

	t.Run("can create the schema", func(t *testing.T) {
		if err := JCompositePkSetup().Create(conn); err != nil {
			t.Fatal(err)
		}
	})
	t.Run("can insert data", func(t *testing.T) {
		v := &CompositePk{Description: "description", P: "text", K: "1"}
		_, err := JCompositePk(sess).Insert(v)
		if err != nil {
			t.Fatal(err)
		}
	})
	t.Run("won t insert same pk twice", func(t *testing.T) {
		v := &CompositePk{P: "text", K: "1"}
		_, err := JCompositePk(sess).Insert(v)
		if err == nil {
			t.Fatal("must return error")
		}
	})
	t.Run("find by pk", func(t *testing.T) {
		v, err := JCompositePk(sess).Find("text", "1")
		if err != nil {
			t.Fatal(err)
		}
		if v.Description != "description" {
			t.Fatal("invalid Description value got=", v.Description)
		}
	})
	t.Run("update a record", func(t *testing.T) {
		_, err := JCompositePk(sess).Update(&CompositePk{P: "text", K: "1", Description: "blah"})
		if err != nil {
			t.Fatal(err)
		}
		v, err := JCompositePk(sess).Find("text", "1")
		if err != nil {
			t.Fatal(err)
		}
		if v.Description != "blah" {
			t.Fatal("invalid Description value got=", v.Description)
		}
	})
	t.Run("delete a record", func(t *testing.T) {
		err := JCompositePk(sess).DeleteByPk("text", "1")
		if err != nil {
			t.Fatal(err)
		}
		_, err = JCompositePk(sess).Find("text", "1")
		if err == nil {
			t.Fatal("err must not be nil")
		}
	})
	t.Run("trigger error if the value len > 255 when mysql is the db engine", func(t *testing.T) {
		_, err := JCompositePk(sess).Insert(&CompositePk{P: strings.Repeat("A", 256), K: strings.Repeat("B", 256)})
		if jedi.Runs(drivers.Mysql) {
			if err != nil {
				t.Fatal("err must not be nil, the value len exceed VARCHAR max (255)")
			}
		} else {
			if err != nil {
				t.Fatal(err)
			}
		}
	})
	t.Run("can drop the schema", func(t *testing.T) {
		if err := JCompositePkSetup().Drop(conn); err != nil {
			t.Fatal(err)
		}
	})
}
