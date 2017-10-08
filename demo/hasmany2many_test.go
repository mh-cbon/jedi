package main

import (
	"os"
	"testing"

	"github.com/mh-cbon/jedi/drivers"
	jedi "github.com/mh-cbon/jedi/runtime"
)

func TestHasMany2Many(t *testing.T) {

	conn := getConnFromEnv(t, true)
	defer conn.Close()
	defer func() {
		if jedi.Runs(drivers.Sqlite) {
			os.Remove(os.Getenv("JDSN"))
		}
	}()

	sess := conn.NewSession(&jedi.EventReceiver{W: os.Stderr})

	pz := &Product{SKU: "Prodz"}
	JProduct(sess).Insert(pz)
	JProduct(sess).Insert(&Product{SKU: "Prodx"})
	JProduct(sess).Insert(&Product{SKU: "Prody"})
	cz := &Category{Name: "Catz"}
	JCategory(sess).Insert(cz)
	JCategory(sess).Insert(&Category{Name: "Catx"})
	JCategory(sess).Insert(&Category{Name: "Caty"})
	cz.LinkWithProducts(sess, pz)

	c := &Category{Name: "CatA"}
	JCategory(sess).Insert(c)
	p := &Product{SKU: "ProdA"}
	JProduct(sess).Insert(p)
	p.LinkWithCategories(sess, c)

	t.Run("provides Read method to get related instances", func(t *testing.T) {
		products, err := c.Products(sess, "p", "", "c").ReadAll()
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
	t.Run("provides Join helper", func(t *testing.T) {
		cats, err := JCategory(sess).Select().JoinProducts("", "p").ReadAll()
		if err != nil {
			t.Fatal(err)
		}
		if len(cats) != 2 {
			t.Fatal("invalid res len wanted", 1, "got", len(cats))
		}
		{
			wanted := cz.ID
			got := cats[0].ID
			if wanted != got {
				t.Fatal("invalid ID wanted", wanted, "got", got)
			}
		}
	})
	t.Run("provides Left Join helper", func(t *testing.T) {
		cats, err := JCategory(sess).Select().LeftJoinProducts("", "p").ReadAll()
		if err != nil {
			t.Fatal(err)
		}
		if len(cats) != 4 {
			t.Fatal("invalid res len wanted", 1, "got", len(cats))
		}
		{
			wanted := cz.ID
			got := cats[0].ID
			if wanted != got {
				t.Fatal("invalid ID wanted", wanted, "got", got)
			}
		}
	})
	// t.Run("provides Right Join helper", func(t *testing.T) {
	// 	log.Println(
	// 		JCategory(sess).As("c").Select("c.*").RightJoinProducts("", "p").String(),
	// 	)
	// 	if jedi.Runs(drivers.Mysql, drivers.Pgsql) {
	// 		cats, err := JCategory(sess).As("c").Select("c.*").RightJoinProducts("", "p").ReadAll()
	// 		if err != nil {
	// 			t.Fatal(err)
	// 		}
	// 		if len(cats) != 4 {
	// 			t.Fatal("invalid res len wanted", 1, "got", len(cats))
	// 		}
	// 		{
	// 			wanted := cz.ID
	// 			got := cats[0].ID
	// 			if wanted != got {
	// 				t.Fatal("invalid ID wanted", wanted, "got", got)
	// 			}
	// 		}
	// 	}
	// })
	t.Run("provides Link helper", func(t *testing.T) {
		if jedi.Runs(drivers.Mysql, drivers.Pgsql) {
			if _, err := c.LinkWithProducts(sess, pz); err != nil {
				t.Fatal(err)
			}
			res, err := c.Products(sess, "p", "", "c").ReadAll()
			if err != nil {
				t.Fatal(err)
			}
			if len(res) != 2 {
				t.Fatal("invalid res len wanted", 2, "got", len(res))
			}
			{
				wanted := cz.ID
				got := res[0].ID
				if wanted != got {
					t.Fatal("invalid ID wanted", wanted, "got", got)
				}
			}
			{
				wanted := c.ID
				got := res[1].ID
				if wanted != got {
					t.Fatal("invalid ID wanted", wanted, "got", got)
				}
			}
		}
	})
	t.Run("provides Unlink helper", func(t *testing.T) {
		if jedi.Runs(drivers.Mysql, drivers.Pgsql) {
			if _, err := c.UnlinkWithProducts(sess, pz); err != nil {
				t.Fatal(err)
			}
			res, err := c.Products(sess, "p", "", "c").ReadAll()
			if err != nil {
				t.Fatal(err)
			}
			if len(res) != 1 {
				t.Fatal("invalid res len wanted", 1, "got", len(res))
			}
			{
				wanted := c.ID
				got := res[0].ID
				if wanted != got {
					t.Fatal("invalid ID wanted", wanted, "got", got)
				}
			}
		}
	})
	t.Run("provides UnlinkAll helper", func(t *testing.T) {
		if jedi.Runs(drivers.Mysql, drivers.Pgsql) {
			if _, err := c.UnlinkAllProducts(sess); err != nil {
				t.Fatal(err)
			}
			res, err := c.Products(sess, "p", "", "c").ReadAll()
			if err != nil {
				t.Fatal(err)
			}
			if len(res) != 0 {
				t.Fatal("invalid res len wanted", 0, "got", len(res))
			}
		}
	})
}
