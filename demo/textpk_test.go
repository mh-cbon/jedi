package main

import (
	"os"
	"strings"
	"testing"

	"github.com/mh-cbon/jedi/drivers"
	jedi "github.com/mh-cbon/jedi/runtime"
)

func TestTextPk(t *testing.T) {

	conn := getConnFromEnv(t, true)
	defer conn.Close()
	defer func() {
		if jedi.Runs(drivers.Sqlite) {
			os.Remove(os.Getenv("JDSN"))
		}
	}()

	sess := conn.NewSession(&jedi.EventReceiver{W: os.Stderr})

	t.Run("can create the schema", func(t *testing.T) {
		if err := JTextPkSetup().Create(conn); err != nil {
			t.Fatal(err)
		}
	})
	t.Run("can insert data", func(t *testing.T) {
		v := &TextPk{Description: "description", Name: "text1"}
		_, err := JTextPk(sess).Insert(v)
		if err != nil {
			t.Fatal(err)
		}
	})
	t.Run("won t insert same pk twice", func(t *testing.T) {
		v := &TextPk{Name: "text1"}
		_, err := JTextPk(sess).Insert(v)
		if err == nil {
			t.Fatal("must return error")
		}
	})
	t.Run("find by pk", func(t *testing.T) {
		v, err := JTextPk(sess).Find("text1")
		if err != nil {
			t.Fatal(err)
		}
		if v.Description != "description" {
			t.Fatal("invalid Description value got=", v.Description)
		}
	})
	t.Run("update a record", func(t *testing.T) {
		_, err := JTextPk(sess).Update(&TextPk{Name: "text1", Description: "blah"})
		if err != nil {
			t.Fatal(err)
		}
		v, err := JTextPk(sess).Find("text1")
		if err != nil {
			t.Fatal(err)
		}
		if v.Description != "blah" {
			t.Fatal("invalid Description value got=", v.Description)
		}
	})
	t.Run("delete a record", func(t *testing.T) {
		err := JTextPk(sess).DeleteByPk("text1")
		if err != nil {
			t.Fatal(err)
		}
		_, err = JTextPk(sess).Find("text1")
		if err == nil {
			t.Fatal("err must not be nil")
		}
	})
	t.Run("trigger error if the value len > 255 when mysql is the db engine", func(t *testing.T) {
		_, err := JTextPk(sess).Insert(&TextPk{Name: strings.Repeat("A", 256)})
		if jedi.Runs(drivers.Mysql) {
			if err == nil {
				t.Fatal("err must not be nil, the value len exceed VARCHAR max (255)")
			}
		} else {
			if err != nil {
				t.Fatal(err)
			}
		}
	})
	t.Run("can drop the schema", func(t *testing.T) {
		if err := JTextPkSetup().Drop(conn); err != nil {
			t.Fatal(err)
		}
	})
}
