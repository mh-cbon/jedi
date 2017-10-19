package main

import (
	"os"
	"testing"

	"github.com/mh-cbon/jedi/drivers"
	jedi "github.com/mh-cbon/jedi/runtime"
)

func TestCRUD(t *testing.T) {

	conn := getConnFromEnv(t, true)
	defer conn.Close()
	defer func() {
		if jedi.Runs(drivers.Sqlite) {
			os.Remove(os.Getenv("JDSN"))
		}
	}()

	sess := conn.NewSession(&jedi.EventReceiver{W: os.Stderr})

	t.Run("insert data", func(t *testing.T) {
		p1 := &Product{SKU: "test"}
		_, err := JProduct(sess).Insert(p1)
		if err != nil {
			t.Fatalf("Data insert failed: %v", err)
		}
		if p1.ID != 1 {
			t.Fatalf("Data insert did not update ID auto increment property got ID=%v, want ID=1", p1.ID)
		}
	})
	t.Run("insert data always autoincrement ID", func(t *testing.T) {
		p1 := &Product{SKU: "test2", ID: 3}
		_, err := JProduct(sess).Insert(p1)
		if err != nil {
			t.Fatalf("Data insert failed: %v", err)
		}
		if p1.ID != 2 {
			t.Fatalf("Data insert did not update ID auto increment property got ID=%v, want ID=2", p1.ID)
		}
	})
	t.Run("insert multiple data and updates autoincrement properties", func(t *testing.T) {
		p3 := &Product{SKU: "test3"}
		p4 := &Product{SKU: "test4"}
		_, err := JProduct(sess).Insert(p3, p4)
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
	t.Run("select data", func(t *testing.T) {
		p1 := new(Product)
		err := JProduct(sess).Select().Where("id = ?", 1).LoadStruct(p1)
		if err != nil {
			t.Fatalf("Data select failed: %v", err)
		}
		if p1.ID != 1 {
			t.Fatalf("Data select is invalid got ID=%v, want ID=1", p1.ID)
		}
		if p1.SKU != "test" {
			t.Fatalf("Data select is invalid got SKU=%v, want SKU=test", p1.SKU)
		}
	})
	t.Run("fail to select data", func(t *testing.T) {
		p1 := new(Product)
		err := JProduct(sess).Select().Where("id = ?", 10).LoadStruct(p1)
		if err == nil {
			t.Fatalf("Data select must fail: err was %v", nil)
		}
	})
	t.Run("find data", func(t *testing.T) {
		p1, err := JProduct(sess).Find(2)
		if err != nil {
			t.Fatalf("Data find failed: %v", err)
		}
		if p1.ID != 2 {
			t.Fatalf("Data find is invalid got ID=%v, want ID=1", p1.ID)
		}
		if p1.SKU != "test2" {
			t.Fatalf("Data find is invalid got SKU=%v, want SKU=test", p1.SKU)
		}
	})
	t.Run("fail to find data", func(t *testing.T) {
		_, err := JProduct(sess).Find(20)
		if err == nil {
			t.Fatalf("Data find must fail: err was %v", nil)
		}
	})
	t.Run("read data", func(t *testing.T) {
		p1, err := JProduct(sess).Select().Where("id = ?", 1).Read()
		if err != nil {
			t.Fatalf("Data read failed: %v", err)
		}
		if p1.ID != 1 {
			t.Fatalf("Data read is inalid got ID=%v, want ID=1", p1.ID)
		}
		if p1.SKU != "test" {
			t.Fatalf("Data read is invalid got SKU=%v, want SKU=test", p1.SKU)
		}
	})
	t.Run("fail to read data", func(t *testing.T) {
		_, err := JProduct(sess).Select().Where("id = ?", 10).Read()
		if err == nil {
			t.Fatalf("Data find must fail: err was %v", nil)
		}
	})
	t.Run("read many data", func(t *testing.T) {
		ps, err := JProduct(sess).Select().OrderAsc("id").ReadAll()
		if err != nil {
			t.Fatalf("Data read failed: %v", err)
		}
		if len(ps) != 4 {
			t.Fatalf("Data length is invalid, got len=%v, want len=4", len(ps))
		}
		p1 := ps[0]
		if p1.ID != 1 {
			t.Fatalf("Data insert did not update ID auto increment property got ID=%v, want ID=1", p1.ID)
		}
		if p1.SKU != "test" {
			t.Fatalf("Data insert did not update ID auto increment property got SKU=%v, want SKU=test", p1.SKU)
		}
	})
	t.Run("fail to read many data", func(t *testing.T) {
		_, err := JProduct(sess).Select("t").ReadAll()
		if err == nil {
			t.Fatalf("Data find must fail: err was %v", nil)
		}
	})
	t.Run("update data", func(t *testing.T) {
		p1 := &Product{SKU: "update1", ID: 1}
		_, err := JProduct(sess).Update(p1)
		if err != nil {
			t.Fatalf("Data update failed: %v", err)
		}
		p1, err2 := JProduct(sess).Find(p1.ID)
		if err2 != nil {
			t.Fatalf("Data update failed: %v", err2)
		}
		if p1.SKU != "update1" {
			t.Fatalf("Data update is invalid got SKU=%v, want SKU=update1", p1.SKU)
		}
	})
	t.Run("update multiple data", func(t *testing.T) {
		p2 := &Product{SKU: "update2", ID: 2}
		p3 := &Product{SKU: "update3", ID: 3}
		_, err := JProduct(sess).Update(p2, p3)
		if err != nil {
			t.Fatalf("Data update failed: %v", err)
		}
		p2, err2 := JProduct(sess).Find(p2.ID)
		if err2 != nil {
			t.Fatalf("Data update failed: %v", err2)
		}
		if p2.SKU != "update2" {
			t.Fatalf("Data update is invalid got SKU=%v, want SKU=update2", p2.SKU)
		}
		p3, err3 := JProduct(sess).Find(p3.ID)
		if err3 != nil {
			t.Fatalf("Data update failed: %v", err3)
		}
		if p3.SKU != "update3" {
			t.Fatalf("Data update is invalid got SKU=%v, want SKU=update3", p3.SKU)
		}
	})
	t.Run("update does not fails if the query did not update rows", func(t *testing.T) {
		p1 := &Product{SKU: "update1", ID: 10}
		res, err := JProduct(sess).Update(p1)
		if n, _ := res.RowsAffected(); n != 0 {
			t.Fatalf("update affected %v rows", n)
		}
		if err != nil {
			t.Fatalf("Data update must not fail: err was %v", err)
		}
	})
	t.Run("MustUpdate fails if the query did not update rows", func(t *testing.T) {
		p1 := &Product{SKU: "update1", ID: 10}
		res, err := JProduct(sess).MustUpdate(p1)
		if n, _ := res.RowsAffected(); n != 0 {
			t.Fatalf("update affected %v rows", n)
		}
		if err == nil {
			t.Fatalf("Data update must fail: err was %v", err)
		}
	})
	t.Run("delete by primary key", func(t *testing.T) {
		if _, err := JProduct(sess).Find(4); err != nil {
			t.Fatalf("Data read must not fail got err=%v", err)
		}
		if err := JProduct(sess).DeleteByPk(4); err != nil {
			t.Fatalf("Data delete must not fail, got err=%v", err)
		}
		if _, err := JProduct(sess).Find(4); err == nil {
			t.Fatalf("Data read must fail, got err=%v", err)
		}
	})
	t.Run("delete data", func(t *testing.T) {
		if _, err := JProduct(sess).Find(3); err != nil {
			t.Fatalf("Data read must not fail got err=%v", err)
		}
		if res, err := JProduct(sess).Delete().Where("id = ?", 3).Exec(); err != nil {
			t.Fatalf("Data delete must not fail, got err=%v", err)
		} else if n, _ := res.RowsAffected(); n != 1 {
			t.Fatalf("Data delete must affect 1 rows, got n=%v", n)
		}
		if _, err := JProduct(sess).Find(3); err == nil {
			t.Fatalf("Data read must fail, got err=%v", err)
		}
	})
	t.Run("MustDelete data raise an error if the query did not affect rows", func(t *testing.T) {
		if _, err := JProduct(sess).Find(3); err == nil {
			t.Fatalf("Data read must fail got err=%v", err)
		}
		if res, err := JProduct(sess).MustDelete().Where("id = ?", 3).Exec(); err == nil {
			t.Fatalf("Data delete must fail, got err=%v", err)
		} else if n, _ := res.RowsAffected(); n != 0 {
			t.Fatalf("Data delete must affect 0 rows, got n=%v", n)
		}
	})
	t.Run("DeleteAll many data", func(t *testing.T) {
		if _, err := JProduct(sess).Find(1); err != nil {
			t.Fatalf("Data read must not fail got err=%v", err)
		}
		if _, err := JProduct(sess).Find(2); err != nil {
			t.Fatalf("Data read must not fail got err=%v", err)
		}
		if res, err := JProduct(sess).DeleteAll(&Product{ID: 1}, &Product{ID: 2}); err != nil {
			t.Fatalf("Data delete must not fail, got err=%v", err)
		} else if n, _ := res.RowsAffected(); n != 2 {
			t.Fatalf("Data delete must affect 2 rows, got n=%v", n)
		}
		if _, err := JProduct(sess).Find(1); err == nil {
			t.Fatalf("Data read must fail, got err=%v", err)
		}
		if _, err := JProduct(sess).Find(2); err == nil {
			t.Fatalf("Data read must fail, got err=%v", err)
		}
	})
	t.Run("DeleteAll does not fail if no rows were affected", func(t *testing.T) {
		if _, err := JProduct(sess).Find(100); err == nil {
			t.Fatalf("Data read must fail got err=%v", err)
		}
		if res, err := JProduct(sess).DeleteAll(&Product{ID: 100}); err != nil {
			t.Fatalf("DeleteAll must not fail, got err=%v", err)
		} else if n, _ := res.RowsAffected(); n != 0 {
			t.Fatalf("DeleteAll must affect 0 rows, got n=%v", n)
		}
	})
	t.Run("MustDeleteAll does not fail if no rows were affected", func(t *testing.T) {
		if _, err := JProduct(sess).Find(100); err == nil {
			t.Fatalf("Data read must fail got err=%v", err)
		}
		if _, err := JProduct(sess).MustDeleteAll(&Product{ID: 100}); err == nil {
			t.Fatalf("MustDeleteAll must  fail, got err=%v", err)
		}
	})
}
