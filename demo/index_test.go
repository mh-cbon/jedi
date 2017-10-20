package main

import (
	"os"
	"testing"

	"github.com/mh-cbon/jedi/drivers"
	jedi "github.com/mh-cbon/jedi/runtime"
)

func TestIndex(t *testing.T) {

	conn := getConnFromEnv(t, true)
	defer conn.Close()
	defer func() {
		if jedi.Runs(drivers.Sqlite) {
			os.Remove(os.Getenv("JDSN"))
		}
	}()

	sess := conn.NewSession(&jedi.EventReceiver{W: os.Stderr})

	t.Run("test index", func(t *testing.T) {
		t.Skip("todo")
	})
	t.Run("unique index is created", func(t *testing.T) {
		_, err := JProduct(sess).Insert(&Product{SKU: "A"})
		if err != nil {
			t.Fatal(err)
		}
		_, err = JProduct(sess).Insert(&Product{SKU: "A"})
		if err == nil {
			t.Fatal("wanted err != nil")
		}
	})

}
