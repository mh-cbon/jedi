package main

import (
	"os"
	"testing"
	"time"

	"github.com/mh-cbon/jedi/drivers"
	jedi "github.com/mh-cbon/jedi/runtime"
)

func TestDate(t *testing.T) {

	conn := getConnFromEnv(t, true)
	defer conn.Close()
	defer func() {
		if jedi.Runs(drivers.Sqlite) {
			os.Remove(os.Getenv("JDSN"))
		}
	}()

	sess := conn.NewSession(&jedi.EventReceiver{W: os.Stderr})

	t.Run("insert data", func(t *testing.T) {
		t1 := &DateType{T: time.Now()}
		t1.TP = &t1.T
		_, err := JDateType(sess).Insert(t1)
		if err != nil {
			t.Fatalf("Data insert failed: %v", err)
		}
		d, err := JDateType(sess).Find(t1.ID)
		if err != nil {
			t.Fatal(err)
		}
		if d.T.Format(time.RFC3339) != t1.T.Format(time.RFC3339) {
			t.Fatal("invalid date d.T =", d.T.Format(time.RFC3339),
				"wanted", t1.T.Format(time.RFC3339))
		}
		if d.TP.Format(time.RFC3339) != t1.TP.Format(time.RFC3339) {
			t.Fatal("invalid date d.TP =", d.TP.Format(time.RFC3339),
				"wanted", t1.TP.Format(time.RFC3339))
		}
	})
	t.Run("update data", func(t *testing.T) {
		t1, err := JDateType(sess).Find(1)
		if err != nil {
			t.Fatal(err)
		}
		t1.T = time.Now()
		t1.TP = &t1.T
		<-time.After(1 * time.Millisecond)
		_, err = JDateType(sess).Update(t1)
		if err != nil {
			t.Fatalf("Data update failed: %v", err)
		}
		d, err := JDateType(sess).Find(t1.ID)
		if err != nil {
			t.Fatal(err)
		}
		if d.T.Format(time.RFC3339) != t1.T.Format(time.RFC3339) {
			t.Fatal("invalid date d.T =", d.T.Format(time.RFC3339),
				"wanted", t1.T.Format(time.RFC3339))
		}
		if d.TP.Format(time.RFC3339) != t1.TP.Format(time.RFC3339) {
			t.Fatal("invalid date d.TP =", d.TP.Format(time.RFC3339),
				"wanted", t1.TP.Format(time.RFC3339))
		}
	})
}
