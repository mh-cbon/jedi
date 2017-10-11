package main

import (
	"os"
	"testing"
	"time"

	"github.com/mh-cbon/jedi/drivers"
	jedi "github.com/mh-cbon/jedi/runtime"
)

func TestDateNotUTC(t *testing.T) {

	conn := getConnFromEnv(t, true)
	defer conn.Close()
	defer func() {
		if jedi.Runs(drivers.Sqlite) {
			os.Remove(os.Getenv("JDSN"))
		}
	}()

	sess := conn.NewSession(&jedi.EventReceiver{W: os.Stderr})

	t.Run("insert data does not write utc date", func(t *testing.T) {
		loc, err := time.LoadLocation("Asia/Tokyo")
		if err != nil {
			t.Fatal(err)
		}
		c := time.Now().In(loc)
		t1 := &DateType{NotUTC: &c}
		_, err = JDateType(sess).Insert(t1)
		if err != nil {
			t.Fatalf("Data insert failed: %v", err)
		}
		d, err := JDateType(sess).Find(t1.ID)
		if err != nil {
			t.Fatal(err)
		}
		if d.NotUTC.Equal(*t1.NotUTC) {
			t.Fatal("invalid date d.NotUTC, they must mismatch")
		}
	})
	t.Run("update data does not set date to UTC", func(t *testing.T) {
		t1, err := JDateType(sess).Find(1)
		if err != nil {
			t.Fatal(err)
		}
		loc, err := time.LoadLocation("Asia/Tokyo")
		if err != nil {
			t.Fatal(err)
		}
		c := time.Now().In(loc)
		t1.NotUTC = &c
		<-time.After(1 * time.Millisecond)
		_, err = JDateType(sess).Update(t1)
		if err != nil {
			t.Fatalf("Data update failed: %v", err)
		}
		d, err := JDateType(sess).Find(t1.ID)
		if err != nil {
			t.Fatal(err)
		}
		if d.NotUTC.Equal(*t1.NotUTC) {
			t.Fatal("invalid date d.NotUTC, they must mismatch")
		}
	})
}
