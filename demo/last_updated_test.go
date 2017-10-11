package main

import (
	"os"
	"testing"
	"time"

	"github.com/mh-cbon/jedi/drivers"
	jedi "github.com/mh-cbon/jedi/runtime"
)

func TestLastUpdatedDate(t *testing.T) {

	conn := getConnFromEnv(t, true)
	defer conn.Close()
	defer func() {
		if jedi.Runs(drivers.Sqlite) {
			os.Remove(os.Getenv("JDSN"))
		}
	}()

	sess := conn.NewSession(&jedi.EventReceiver{W: os.Stderr})

	t.Run("insert automatically writes nil date", func(t *testing.T) {
		t1 := &DateType{}
		_, err := JDateType(sess).Insert(t1)
		if err != nil {
			t.Fatalf("Data insert failed: %v", err)
		}
		d, err := JDateType(sess).Find(t1.ID)
		if err != nil {
			t.Fatal(err)
		}
		if d.LastUpdated == nil {
			t.Fatal("LastUpdated date not written, must not be nil")
		}
	})
	t.Run("insert won t automatically set last updated date if not nil", func(t *testing.T) {
		c := time.Now()
		t1 := &DateType{LastUpdated: &c}
		_, err := JDateType(sess).Insert(t1)
		if err != nil {
			t.Fatalf("Data insert failed: %v", err)
		}
		d, err := JDateType(sess).Find(t1.ID)
		if err != nil {
			t.Fatal(err)
		}
		if d.LastUpdated == nil {
			t.Fatal("LastUpdated date not written, must not be nil")
		} else if d.LastUpdated.Format(time.RFC3339) != t1.LastUpdated.Format(time.RFC3339) {
			t.Fatal("invalid date d.LastUpdated wanted", d.LastUpdated.Format(time.RFC3339),
				"got", t1.LastUpdated.Format(time.RFC3339))
		}
	})
	t.Run("update data automatically update LastUpdated property", func(t *testing.T) {
		t1, err := JDateType(sess).Find(1)
		if err != nil {
			t.Fatal(err)
		}
		c := *t1.LastUpdated
		<-time.After(1 * time.Millisecond)
		_, err = JDateType(sess).Update(t1)
		if err != nil {
			t.Fatalf("Data update failed: %v", err)
		}
		d, err := JDateType(sess).Find(t1.ID)
		if err != nil {
			t.Fatal(err)
		}
		if d.LastUpdated.Format(time.RFC3339Nano) == c.Format(time.RFC3339Nano) {
			t.Fatal("they must mismatch d.LastUpdated", d.LastUpdated.Format(time.RFC3339Nano),
				"c", c.Format(time.RFC3339Nano),
			)
		}
	})
	t.Run("update returns error if the LastUpdated value does not match", func(t *testing.T) {
		t1, err := JDateType(sess).Find(1)
		if err != nil {
			t.Fatal(err)
		}
		t1.LastUpdated = nil
		_, err = JDateType(sess).Update(t1)
		if err == nil {
			t.Fatalf("Data update must fail, got err= %v", err)
		}
	})
	t.Run("update returns error if the LastUpdated value does not match 2", func(t *testing.T) {
		t1, err := JDateType(sess).Find(1)
		if err != nil {
			t.Fatal(err)
		}
		<-time.After(time.Millisecond)
		c := time.Now()
		t1.LastUpdated = &c
		_, err = JDateType(sess).Update(t1)
		if err == nil {
			t.Fatalf("Data update must fail, got err= %v", err)
		}
	})
	t.Run("insert then update", func(t *testing.T) {
		t1 := &DateType{}
		_, err := JDateType(sess).Insert(t1)
		if err != nil {
			t.Fatalf("Data insert failed: %v", err)
		}
		t1.T = time.Now()
		_, err = JDateType(sess).Update(t1)
		if err != nil {
			t.Fatalf("Data update failed: %v", err)
		}
	})
}
