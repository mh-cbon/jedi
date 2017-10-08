package main

import (
	"os"
	"testing"

	"github.com/mh-cbon/jedi/drivers"
	jedi "github.com/mh-cbon/jedi/runtime"
)

func TestHasOne(t *testing.T) {

	conn := getConnFromEnv(t, true)
	defer conn.Close()
	defer func() {
		if jedi.Runs(drivers.Sqlite) {
			os.Remove(os.Getenv("JDSN"))
		}
	}()

	sess := conn.NewSession(nil)

	//Prepare some data
	b := &Brand{Name: "z"}
	p := &Product{SKU: "B"}
	m := &Product{SKU: "Master"}
	JBrand(sess).Insert(b)
	JProduct(sess).Insert(&Product{SKU: "A"})
	JProduct(sess).Insert(m)
	p = p.SetBrand(b)
	p.SetMaster(m)
	JProduct(sess).Insert(p)

	t.Run("provides Join helper", func(t *testing.T) {
		res, err := JProduct(sess).Select().JoinBrand("b").ReadAll()
		if err != nil {
			t.Fatal(err)
		}
		if len(res) != 1 {
			t.Fatalf("invalid length wanted=%v, got=%v", 1, len(res))
		}
		if res[0].SKU != "B" {
			t.Fatalf("inavlid object, wanted=%#v got=%#v", p, res[0])
		}
	})
	t.Run("provides Left Join helper", func(t *testing.T) {
		res, err := JProduct(sess).As("p").Select("p.*").LeftJoinBrand("b").ReadAll()
		if err != nil {
			t.Fatal(err)
		}
		if len(res) != 3 {
			t.Fatalf("invalid length wanted=%v, got=%v", 3, len(res))
		}
		if res[0].SKU != "A" {
			t.Fatalf("inavlid object, wanted=%#v got=%#v", p, res[0])
		}
	})
	t.Run("provides Right Join helper", func(t *testing.T) {
		if jedi.Runs(drivers.Mysql, drivers.Pgsql) {
			// This does not run on sqlite at all
			res, err := JProduct(sess).As("p").Select("p.*").RightJoinBrand("b").ReadAll()
			if err != nil {
				t.Fatal(err)
			}
			if len(res) != 1 {
				t.Fatalf("invalid length wanted=%v, got=%v", 1, len(res))
			}
			if res[0].SKU != "A" {
				t.Fatalf("inavlid object, wanted=%#v got=%#v", p, res[0])
			}
		} else {
			t.Skip("sqlite can t do that")
		}
	})
	t.Run("provides Set method to update imported primary keys", func(t *testing.T) {
		b1 := &Brand{Name: "adidas"}
		if _, err := JBrand(sess).Insert(b1); err != nil {
			t.Fatal(err)
		}
		p1 := &Product{SKU: "B"}
		if p1.BrandID != nil {
			t.Fatal("BrandID must be nil")
		}
		p1.SetBrand(b1)
		if p1.BrandID == nil {
			t.Fatal("BrandID must not be nil")
		}
		if *p1.BrandID != b1.ID {
			t.Fatal("BrandID must eq Brand.ID")
		}
	})
	t.Run("Set can handle nil values", func(t *testing.T) {
		p1 := &Product{SKU: "B"}
		p1.SetBrand(b)
		if p1.BrandID == nil {
			t.Fatal("BrandID must not be nil")
		}

		p1.SetBrand(nil)
		if p1.BrandID != nil {
			t.Fatal("BrandID must be nil")
		}
	})
	t.Run("provides Read method to get a remote value", func(t *testing.T) {
		b1, err := p.Brand(sess)
		if err != nil {
			t.Fatal(err)
		}
		if b1 == nil {
			t.Fatal("remote object must no be nil")
		}
		if b1.ID != b.ID {
			t.Fatal("got invalid ID")
		}
	})
	t.Run("handles self-referencing read", func(t *testing.T) {
		mt, err := p.Master(sess)
		if err != nil {
			t.Fatal(err)
		}
		if mt == nil {
			t.Fatal("remote object must no be nil")
		}
		if mt.ID != m.ID {
			t.Fatal("got invalid ID")
		}
	})
	t.Run("handles self-referencing join", func(t *testing.T) {
		res, err := JProduct(sess).Select().JoinMaster("m").Where(JProductModel.As("m").SKU.Like("mas")).ReadAll()
		if err != nil {
			t.Fatal(err)
		}
		if len(res) != 1 {
			t.Fatal("invalid res length, wanted 1, got", len(res))
		}
		if res[0].ID != p.ID {
			t.Fatal("got invalid ID")
		}
		if res[0].SKU != p.SKU {
			t.Fatal("got invalid SKU")
		}
	})
}
