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
		} else if !d.LastUpdated.Equal(*t1.LastUpdated) {
			t.Fatal("invalid date d.LastUpdated =", d.LastUpdated.Format(time.RFC3339Nano),
				"t1=", t1.LastUpdated.Format(time.RFC3339Nano))
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
		// if d.LastUpdated.Format(time.RFC3339Nano) == c.Format(time.RFC3339Nano) {
		if d.LastUpdated.Equal(c) {
			t.Fatal("they must mismatch d.LastUpdated", d.LastUpdated.Format(time.RFC3339Nano),
				"c", c.Format(time.RFC3339Nano),
			)
		}
	})
	t.Run("update silently fails if the query does not affect rows", func(t *testing.T) {
		t1, err := JDateType(sess).Find(1)
		if err != nil {
			t.Fatal(err)
		}
		t1.LastUpdated = nil
		_, err = JDateType(sess).Update(t1)
		if err != nil {
			t.Fatalf("Data update must not fail, got err= %v", err)
		}
	})
	t.Run("MustUpdate returns an error if the query does not affect rows", func(t *testing.T) {
		t1, err := JDateType(sess).Find(1)
		if err != nil {
			t.Fatal(err)
		}
		t1.LastUpdated = nil
		_, err = JDateType(sess).MustUpdate(t1)
		if err == nil {
			t.Fatalf("Data update must fail, got err= %v", err)
		}
	})
	t.Run("update truncates @last_updated properties to 6 digits", func(t *testing.T) {
		t1, err := JDateType(sess).Find(1)
		if err != nil {
			t.Fatal(err)
		}
		c := time.Now()
		t1.LastUpdated = &c
		_, err = JDateType(sess).Update(t1)
		if err != nil {
			t.Fatalf("Data update must not fail, got err= %v", err)
		}
		if t1.LastUpdated.Equal(c.UTC()) {
			t.Fatalf("Values must not match %v %v", t1.LastUpdated, c.UTC())
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
