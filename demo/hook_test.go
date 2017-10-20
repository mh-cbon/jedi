package main

import (
	"os"
	"testing"

	"github.com/mh-cbon/jedi/drivers"
	jedi "github.com/mh-cbon/jedi/runtime"
)

func TestHook(t *testing.T) {

	conn := getConnFromEnv(t, true)
	defer conn.Close()
	defer func() {
		if jedi.Runs(drivers.Sqlite) {
			os.Remove(os.Getenv("JDSN"))
		}
	}()

	sess := conn.NewSession(&jedi.EventReceiver{W: os.Stderr})

	t.Run("fails to Insert because of the hook", func(t *testing.T) {
		_, err := JHookDemo(sess).Insert(&HookDemo{})
		if err == nil {
			t.Fatal("wanted err != nil")
		} else if err.Error() != "It won t happen" {
			t.Fatal("wanted err =It won t happen")
		}
	})
	t.Run("fails to Update because of the hook", func(t *testing.T) {
		_, err := JHookDemo(sess).Update(&HookDemo{})
		if err == nil {
			t.Fatal("wanted err != nil")
		} else if err.Error() != "It won t happen" {
			t.Fatal("wanted err =It won t happen")
		}
	})
	t.Run("fails to MustUpdate because of the hook", func(t *testing.T) {
		_, err := JHookDemo(sess).MustUpdate(&HookDemo{})
		if err == nil {
			t.Fatal("wanted err != nil")
		} else if err.Error() != "It won t happen" {
			t.Fatal("wanted err =It won t happen")
		}
	})

}
