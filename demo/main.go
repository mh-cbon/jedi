package main

import (
	"flag"
	"fmt"
	"os"
	"time"

	_ "github.com/go-sql-driver/mysql"
	"github.com/gocraft/dbr"
	_ "github.com/lib/pq"
	_ "github.com/mattn/go-sqlite3"
	jedi "github.com/mh-cbon/jedi/runtime"
)

// go run demo/* --driver postgres --dsn="user=test2 pwd=test2 dbname=test2 sslmode=disable"
// go run demo/* --driver=mysql --dsn="root:@/test?parseTime=true"
func main() {

	var dsn string
	var driver string
	flag.StringVar(&dsn, "dsn", "s.db", "dsn")
	flag.StringVar(&driver, "driver", "sqlite3", "db driver")
	flag.Parse()

	defer func() {
		if driver == "sqlite3" {
			os.Remove(dsn)
		}
	}()

	// open dbr conn
	conn, err := dbr.Open(driver, dsn, nil)
	fail(err)
	defer conn.Close()

	// prepare the schema
	{
		// setup jedi
		err = jedi.Setup(conn, true)
		fail(err)

		//fyi
		fmt.Printf("model %#v\n", JSampleSetup())
	}

	// make a dbr session
	sess := conn.NewSession(nil)
	defer sess.Close()

	t1(sess)
	t2(sess)
}

func t2(sess *dbr.Session) {

}

func t1(sess *dbr.Session) {
	u := time.Now().UTC()
	s1 := &Sample{Name: "Tomate", UpdateDate: time.Now().UTC(), RemovalDate: &u}
	// using tx
	{
		tx, err := sess.Begin()
		fail(err)
		defer tx.Rollback()

		// using an access layer
		qSample := JSample(tx)
		fmt.Printf("querier %#v\n", qSample)
		fmt.Println()

		// create some new data
		_, err = qSample.Insert(s1)
		fail(err)
		fmt.Printf("after insert %#v\n", s1)
		fmt.Println()

		s2 := &Sample{Name: "Tomate2"}
		_, err = qSample.Insert(s2)
		fail(err)
		fmt.Printf("after insert %#v\n", s2)
		fmt.Println()

		// delete some data
		_, err = qSample.Delete().Where("id > ?", 1).Exec()
		fail(err)

		err = tx.Commit()
		fail(err)
	}

	// without tx
	qSample := JSample(sess)

	// select some data
	all, err := qSample.Select().ReadAll()
	fail(err)
	for _, a := range all {
		fmt.Printf("after select %#v\n", a)
	}
	fmt.Println()

	// update some data
	s1.Name = "Poivron"
	_, err = qSample.Update(s1)
	fail(err)
	fmt.Printf("after update %#v\n", s1)
	fmt.Println()

	// select some more data
	one, err4 := qSample.Select().Where(JSampleModel.Eq(s1)).Read()
	fail(err4)
	fmt.Printf("after select %#v\n", one)
}

func fail(err error) {
	if err != nil {
		panic(err)
	}
}
