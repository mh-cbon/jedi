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
	t.Run("can read has_one to many relations", func(t *testing.T) {
		n := "rr"
		JTextPk(sess).Insert(&TextPk{Name: n})
		p := &HasOneTextPk{RelatedName: &n}
		_, err := p.Related(sess)
		if err != nil {
			t.Fatal(err)
		}
	})
	t.Run("can link some related objects", func(t *testing.T) {
		r := &HasManyTextPk{}
		r1 := &HasManyTextPk{}
		JHasManyTextPk(sess).Insert(r, r1)
		p := &TextPk{Name: "rr"}
		_, err := p.LinkWithRelateds(sess, r, r1)
		if err != nil {
			t.Fatal(err)
		}
	})
	t.Run("can read has_many to one relations", func(t *testing.T) {
		p := &TextPk{Name: "rr"}
		res, err := p.Relateds(sess, "", "", "").ReadAll()
		if err != nil {
			t.Fatal(err)
		}
		if len(res) != 2 {
			t.Fatal("must get 2 results")
		}
	})
	t.Run("fails to link same objects twice", func(t *testing.T) {
		p := &TextPk{}
		_, err := p.LinkWithRelateds(sess, &HasManyTextPk{ID: 3}, &HasManyTextPk{ID: 3})
		if err == nil {
			t.Fatal("must fail: UNIQUE constraint failed")
		}
	})
	t.Run("can unlink some related objects", func(t *testing.T) {
		p := &TextPk{}
		_, err := p.UnlinkWithRelateds(sess, &HasManyTextPk{ID: 1}, &HasManyTextPk{ID: 2})
		if err != nil {
			t.Fatal(err)
		}
	})
	t.Run("can drop the schema", func(t *testing.T) {
		if err := JTextPkSetup().Drop(conn); err != nil {
			t.Fatal(err)
		}
	})
}
