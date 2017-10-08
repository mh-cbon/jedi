package main

import (
	"os"
	"testing"

	"github.com/mh-cbon/jedi/drivers"
	jedi "github.com/mh-cbon/jedi/runtime"
)

func TestHasMany2One(t *testing.T) {

	conn := getConnFromEnv(t, true)
	defer conn.Close()
	defer func() {
		if jedi.Runs(drivers.Sqlite) {
			os.Remove(os.Getenv("JDSN"))
		}
	}()

	sess := conn.NewSession(&jedi.EventReceiver{W: os.Stderr})

	//Prepare some data
	JBrand(sess).Insert(&Brand{Name: "b"})
	JBrand(sess).Insert(&Brand{Name: "d"})
	b := &Brand{Name: "a"}
	JBrand(sess).Insert(b)
	p := &Product{SKU: "A"}
	p.SetBrand(b)
	JProduct(sess).Insert(p)
	v := &Product{SKU: "V"}
	v.SetMaster(p)
	JProduct(sess).Insert(v)

	t.Run("provides Join helper", func(t *testing.T) {
		brands, err := JBrand(sess).Select().JoinProducts("p").Where(JProductModel.As("p").SKU.Like("A")).ReadAll()
		if err != nil {
			t.Fatal(err)
		}
		if len(brands) != 1 {
			t.Fatal("invalid res length, wanted=1 got", len(brands))
		}
		if brands[0].Name != b.Name {
			t.Fatal("invalid brand Name expected=", b.Name, "got", brands[0].Name)
		}
	})
	t.Run("provides LeftJoin helper", func(t *testing.T) {
		brands, err := JBrand(sess).Select().LeftJoinProducts("p").Where(JProductModel.As("p").SKU.Like("A")).ReadAll()
		if err != nil {
			t.Fatal(err)
		}
		if len(brands) != 1 {
			t.Fatal("invalid res length, wanted=1 got", len(brands))
		}
		if brands[0].Name != b.Name {
			t.Fatal("invalid brand Name expected=", b.Name, "got", brands[0].Name)
		}
	})
	// t.Run("provides RightJoin helper", func(t *testing.T) {
	// 	if jedi.Runs(drivers.Mysql, drivers.Pgsql) {
	// 		brands, err := JBrand(sess).Select().RightJoinProducts("p").Where(JProductModel.As("p").SKU.Like("A")).ReadAll()
	// 		if err != nil {
	// 			t.Fatal(err)
	// 		}
	// 		if len(brands) != 1 {
	// 			t.Fatal("invalid res length, wanted=1 got", len(brands))
	// 		}
	// 		if brands[0].Name != b.Name {
	// 			t.Fatal("invalid brand Name expected=", b.Name, "got", brands[0].Name)
	// 		}
	// 	} else {
	// 		t.Skip("sqlite can t do that")
	// 	}
	// })
	t.Run("provides Read method to get related instances", func(t *testing.T) {
		products, err := b.Products(sess, "p", "b").ReadAll()
		if err != nil {
			t.Fatal(err)
		}
		if len(products) != 1 {
			t.Fatal("invalid res len wanted", 1, "got", len(products))
		}
		{
			wanted := p.ID
			got := products[0].ID
			if wanted != got {
				t.Fatal("invalid ID wanted", wanted, "got", got)
			}
		}
	})
	t.Run("provides Read method to get self-related instances", func(t *testing.T) {
		products, err := p.Variances(sess, "p", "v").ReadAll()
		if err != nil {
			t.Fatal(err)
		}
		if len(products) != 1 {
			t.Fatal("invalid res len wanted", 1, "got", len(products))
		}
		{
			wanted := v.ID
			got := products[0].ID
			if wanted != got {
				t.Fatal("invalid ID wanted", wanted, "got", got)
			}
		}
	})
}
