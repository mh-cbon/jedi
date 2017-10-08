package main

import (
	"os"
	"testing"

	"github.com/mh-cbon/jedi/drivers"
	jedi "github.com/mh-cbon/jedi/runtime"
)

func TestAI(t *testing.T) {

	conn := getConnFromEnv(t, true)
	defer conn.Close()
	defer func() {
		if jedi.Runs(drivers.Sqlite) {
			os.Remove(os.Getenv("JDSN"))
		}
	}()

	sess := conn.NewSession(&jedi.EventReceiver{W: os.Stderr})

	t.Run("insert data", func(t *testing.T) {
		p1 := &BasicPK{}
		_, err := JBasicPK(sess).Insert(p1)
		if err != nil {
			t.Fatalf("Data insert failed: %v", err)
		}
		if p1.ID != 1 {
			t.Fatalf("Data insert did not update ID auto increment property got ID=%v, want ID=1", p1.ID)
		}
	})
	t.Run("insert data always autoincrement ID", func(t *testing.T) {
		p1 := &BasicPK{ID: 3}
		_, err := JBasicPK(sess).Insert(p1)
		if err != nil {
			t.Fatalf("Data insert failed: %v", err)
		}
		if p1.ID != 2 {
			t.Fatalf("Data insert did not update ID auto increment property got ID=%v, want ID=2", p1.ID)
		}
	})
	t.Run("insert multiple data and updates autoincrement properties", func(t *testing.T) {
		p3 := &BasicPK{}
		p4 := &BasicPK{}
		_, err := JBasicPK(sess).Insert(p3, p4)
		if err != nil {
			t.Fatalf("Data insert multiple failed: %v", err)
		}
		if p3.ID != 3 {
			t.Fatalf("Data insert multiple did not update ID auto increment property got ID=%v, want ID=3", p3.ID)
		}
		if p4.ID != 4 {
			t.Fatalf("Data insert multiple did not update ID auto increment property got ID=%v, want ID=4", p4.ID)
		}
	})
}
