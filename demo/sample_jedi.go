// Generated with mh-cbon/jedi. Do not edit by hand.
package main

import (
	"database/sql"
	"github.com/gocraft/dbr"
	"github.com/mh-cbon/jedi/builder"
	"github.com/mh-cbon/jedi/drivers"
	"github.com/mh-cbon/jedi/runtime"
	"strings"
)

func init() {
	runtime.Register(

		JSampleSetup,

		JSample2Setup,

		JSampleViewSetup,
	)
}

type jSampleSetup struct {
	Name       string
	CreateStmt string
	DropStmt   string
}

func (c jSampleSetup) Create(db *dbr.Connection) error {
	_, err := db.Exec(c.CreateStmt)
	return err
}
func (c jSampleSetup) Drop(db *dbr.Connection) error {
	_, err := db.Exec(c.DropStmt)
	return err
}

// JSampleSetup helps to create/drop the schema
func JSampleSetup() runtime.Setuper {
	driver := runtime.GetCurrentDriver()

	var create string
	var drop string

	if driver == drivers.Sqlite {
		create = `CREATE TABLE IF NOT EXISTS sample (
id INTEGER PRIMARY KEY AUTOINCREMENT,
name TEXT,
description TEXT,
update_date datetime,
removal_date datetime

)`
	} else if driver == drivers.Mysql {
		create = `CREATE TABLE IF NOT EXISTS sample (
id INTEGER NOT NULL AUTO_INCREMENT,
name TEXT,
description TEXT,
update_date datetime,
removal_date datetime,
PRIMARY KEY (id) 

)`
	} else if driver == drivers.Pgsql {
		create = `CREATE TABLE IF NOT EXISTS sample (
id INTEGER,
name TEXT,
description TEXT,
update_date datetime,
removal_date datetime

)`
	}

	if driver == drivers.Sqlite {
		drop = `DROP TABLE IF EXISTS sample`
	} else if driver == drivers.Mysql {
		drop = `DROP TABLE IF EXISTS sample`
	} else if driver == drivers.Pgsql {
		drop = `DROP TABLE IF EXISTS sample`
	}

	return jSampleSetup{
		Name:       `sample`,
		CreateStmt: create,
		DropStmt:   drop,
	}
}

var JSampleModel = struct {
	Eq func(...*Sample) dbr.Builder
	In func(...*Sample) dbr.Builder

	ID struct {
		Name string
		IsPk bool
		IsAI bool
		Eq   func(interface{}) dbr.Builder
		In   func(...interface{}) dbr.Builder
		Gt   func(interface{}) dbr.Builder
		Gte  func(interface{}) dbr.Builder
		Lt   func(interface{}) dbr.Builder
		Lte  func(interface{}) dbr.Builder
	}

	Name struct {
		Name string
		Eq   func(interface{}) dbr.Builder
		In   func(...interface{}) dbr.Builder
		Like func(interface{}) dbr.Builder
	}

	Description struct {
		Name string
		Eq   func(interface{}) dbr.Builder
		In   func(...interface{}) dbr.Builder
		Like func(interface{}) dbr.Builder
	}
}{

	Eq: func(s ...*Sample) dbr.Builder {
		ors := []dbr.Builder{}
		for _, t := range s {
			ors = append(ors, dbr.Or(
				dbr.And(

					dbr.Eq(`id`, t.ID),
				),
			))
		}
		return dbr.And(ors...)
	},
	In: func(s ...*Sample) dbr.Builder {
		ors := []dbr.Builder{}
		for _, t := range s {
			ors = append(ors, dbr.Or(
				dbr.And(

					dbr.Eq(`id`, t.ID),
				),
			))
		}
		return dbr.And(ors...)
	},

	ID: struct {
		Name string
		IsPk bool
		IsAI bool
		Eq   func(interface{}) dbr.Builder
		In   func(...interface{}) dbr.Builder
		Gt   func(interface{}) dbr.Builder
		Gte  func(interface{}) dbr.Builder
		Lt   func(interface{}) dbr.Builder
		Lte  func(interface{}) dbr.Builder
	}{
		Name: `id`,
		IsPk: true,
		IsAI: true,
		Eq: func(v interface{}) dbr.Builder {
			return dbr.Eq(`id`, v)
		},
		In: func(v ...interface{}) dbr.Builder {
			return dbr.Eq(`id`, v)
		},
		Gt: func(v interface{}) dbr.Builder {
			return dbr.Gt(`id`, v)
		},
		Gte: func(v interface{}) dbr.Builder {
			return dbr.Gte(`id`, v)
		},
		Lt: func(v interface{}) dbr.Builder {
			return dbr.Lt(`id`, v)
		},
		Lte: func(v interface{}) dbr.Builder {
			return dbr.Lte(`id`, v)
		},
	},

	Name: struct {
		Name string
		Eq   func(interface{}) dbr.Builder
		In   func(...interface{}) dbr.Builder
		Like func(interface{}) dbr.Builder
	}{
		Name: `name`,
		Eq: func(v interface{}) dbr.Builder {
			return dbr.Eq(`name`, v)
		},
		In: func(v ...interface{}) dbr.Builder {
			return dbr.Eq(`name`, v)
		},
		Like: func(v interface{}) dbr.Builder {
			return builder.Like(`name`, v)
		},
	},

	Description: struct {
		Name string
		Eq   func(interface{}) dbr.Builder
		In   func(...interface{}) dbr.Builder
		Like func(interface{}) dbr.Builder
	}{
		Name: `description`,
		Eq: func(v interface{}) dbr.Builder {
			return dbr.Eq(`description`, v)
		},
		In: func(v ...interface{}) dbr.Builder {
			return dbr.Eq(`description`, v)
		},
		Like: func(v interface{}) dbr.Builder {
			return builder.Like(`description`, v)
		},
	},
}

// JSample provides a basic querier
func JSample(db dbr.SessionRunner) jSampleQuerier {
	return jSampleQuerier{
		name: `sample`,
		db:   db,
	}
}

type jSampleSelectBuilder struct {
	*builder.SelectBuilder
}

func (c *jSampleSelectBuilder) Read() (*Sample, error) {
	var one Sample
	err := c.LoadStruct(&one)
	return &one, err
}
func (c *jSampleSelectBuilder) ReadAll() ([]*Sample, error) {
	var all []*Sample
	_, err := c.LoadStructs(&all)
	return all, err
}
func (c *jSampleSelectBuilder) Where(query interface{}, value ...interface{}) *jSampleSelectBuilder {
	c.SelectBuilder.Where(query, value...)
	return c
}

type jSampleQuerier struct {
	name string
	db   dbr.SessionRunner
}

func (c jSampleQuerier) Select(what ...string) *jSampleSelectBuilder {
	if len(what) == 0 {
		what = append(what, "*")
	}
	return &jSampleSelectBuilder{
		&builder.SelectBuilder{
			SelectBuilder: c.db.Select(what...).From(c.name),
		},
	}
}
func (c jSampleQuerier) Count(what ...string) *jSampleSelectBuilder {
	if len(what) == 0 {
		what = append(what, "*")
	}
	return c.Select("COUNT(" + strings.Join(what, ",") + ")")
}

func (c jSampleQuerier) Insert(data *Sample) (sql.Result, error) {
	res, err := c.db.InsertInto(c.name).Columns(

		`name`,

		`description`,

		`update_date`,

		`removal_date`,
	).Record(data).Exec()

	if err == nil {
		id, err2 := res.LastInsertId()
		if err2 != nil {
			return res, err2
		}
		data.ID = id
	}

	return res, err
}

func (c jSampleQuerier) Update(data *Sample) (sql.Result, error) {
	res, err := c.db.Update(c.name).
		Set(`name`, data.Name).
		Set(`description`, data.Description).
		Set(`update_date`, data.UpdateDate).
		Set(`removal_date`, data.RemovalDate).
		Where("id = ?", data.ID).
		Exec()
	return res, err
}

func (c jSampleQuerier) Delete() *builder.DeleteBuilder {
	return &builder.DeleteBuilder{
		DeleteBuilder: c.db.DeleteFrom(c.name),
	}
}

func (c jSampleQuerier) DeleteByPk(id int64) error {
	_, err := c.Delete().Where(

		JSampleModel.ID.Eq(id),
	).Exec()
	return err
}
func (c jSampleQuerier) DeleteAll(items ...*Sample) error {
	q := c.Delete()
	for _, item := range items {
		q = q.Where(
			dbr.Or(
				dbr.And(

					JSampleModel.ID.Eq(item.ID),
				),
			),
		)
	}
	_, err := q.Exec()
	return err
}
func (c jSampleQuerier) Find(id int64) (*Sample, error) {
	return c.Select().Where(

		JSampleModel.ID.Eq(id),
	).Read()
}

type jSample2Setup struct {
	Name       string
	CreateStmt string
	DropStmt   string
}

func (c jSample2Setup) Create(db *dbr.Connection) error {
	_, err := db.Exec(c.CreateStmt)
	return err
}
func (c jSample2Setup) Drop(db *dbr.Connection) error {
	_, err := db.Exec(c.DropStmt)
	return err
}

// JSample2Setup helps to create/drop the schema
func JSample2Setup() runtime.Setuper {
	driver := runtime.GetCurrentDriver()

	var create string
	var drop string

	if driver == drivers.Sqlite {
		create = `CREATE TABLE IF NOT EXISTS second_sample (
name TEXT,
description TEXT,
PRIMARY KEY (name) 

)`
	} else if driver == drivers.Mysql {
		create = `CREATE TABLE IF NOT EXISTS second_sample (
name TEXT NOT NULL,
description TEXT,
PRIMARY KEY (name) 

)`
	} else if driver == drivers.Pgsql {
		create = `CREATE TABLE IF NOT EXISTS second_sample (
name TEXT,
description TEXT

)`
	}

	if driver == drivers.Sqlite {
		drop = `DROP TABLE IF EXISTS second_sample`
	} else if driver == drivers.Mysql {
		drop = `DROP TABLE IF EXISTS second_sample`
	} else if driver == drivers.Pgsql {
		drop = `DROP TABLE IF EXISTS second_sample`
	}

	return jSample2Setup{
		Name:       `second_sample`,
		CreateStmt: create,
		DropStmt:   drop,
	}
}

var JSample2Model = struct {
	Eq func(...*Sample2) dbr.Builder
	In func(...*Sample2) dbr.Builder

	Name struct {
		Name string
		Eq   func(interface{}) dbr.Builder
		In   func(...interface{}) dbr.Builder
		Like func(interface{}) dbr.Builder
	}

	Description struct {
		Name string
		Eq   func(interface{}) dbr.Builder
		In   func(...interface{}) dbr.Builder
		Like func(interface{}) dbr.Builder
	}
}{

	Eq: func(s ...*Sample2) dbr.Builder {
		ors := []dbr.Builder{}
		for _, t := range s {
			ors = append(ors, dbr.Or(
				dbr.And(

					dbr.Eq(`name`, t.Name),
				),
			))
		}
		return dbr.And(ors...)
	},
	In: func(s ...*Sample2) dbr.Builder {
		ors := []dbr.Builder{}
		for _, t := range s {
			ors = append(ors, dbr.Or(
				dbr.And(

					dbr.Eq(`name`, t.Name),
				),
			))
		}
		return dbr.And(ors...)
	},

	Name: struct {
		Name string
		Eq   func(interface{}) dbr.Builder
		In   func(...interface{}) dbr.Builder
		Like func(interface{}) dbr.Builder
	}{
		Name: `name`,
		Eq: func(v interface{}) dbr.Builder {
			return dbr.Eq(`name`, v)
		},
		In: func(v ...interface{}) dbr.Builder {
			return dbr.Eq(`name`, v)
		},
		Like: func(v interface{}) dbr.Builder {
			return builder.Like(`name`, v)
		},
	},

	Description: struct {
		Name string
		Eq   func(interface{}) dbr.Builder
		In   func(...interface{}) dbr.Builder
		Like func(interface{}) dbr.Builder
	}{
		Name: `description`,
		Eq: func(v interface{}) dbr.Builder {
			return dbr.Eq(`description`, v)
		},
		In: func(v ...interface{}) dbr.Builder {
			return dbr.Eq(`description`, v)
		},
		Like: func(v interface{}) dbr.Builder {
			return builder.Like(`description`, v)
		},
	},
}

// JSample2 provides a basic querier
func JSample2(db dbr.SessionRunner) jSample2Querier {
	return jSample2Querier{
		name: `second_sample`,
		db:   db,
	}
}

type jSample2SelectBuilder struct {
	*builder.SelectBuilder
}

func (c *jSample2SelectBuilder) Read() (*Sample2, error) {
	var one Sample2
	err := c.LoadStruct(&one)
	return &one, err
}
func (c *jSample2SelectBuilder) ReadAll() ([]*Sample2, error) {
	var all []*Sample2
	_, err := c.LoadStructs(&all)
	return all, err
}
func (c *jSample2SelectBuilder) Where(query interface{}, value ...interface{}) *jSample2SelectBuilder {
	c.SelectBuilder.Where(query, value...)
	return c
}

type jSample2Querier struct {
	name string
	db   dbr.SessionRunner
}

func (c jSample2Querier) Select(what ...string) *jSample2SelectBuilder {
	if len(what) == 0 {
		what = append(what, "*")
	}
	return &jSample2SelectBuilder{
		&builder.SelectBuilder{
			SelectBuilder: c.db.Select(what...).From(c.name),
		},
	}
}
func (c jSample2Querier) Count(what ...string) *jSample2SelectBuilder {
	if len(what) == 0 {
		what = append(what, "*")
	}
	return c.Select("COUNT(" + strings.Join(what, ",") + ")")
}

func (c jSample2Querier) Insert(data *Sample2) (sql.Result, error) {
	res, err := c.db.InsertInto(c.name).Columns(

		`name`,

		`description`,
	).Record(data).Exec()

	return res, err
}

func (c jSample2Querier) Update(data *Sample2) (sql.Result, error) {
	res, err := c.db.Update(c.name).
		Set(`description`, data.Description).
		Where("name = ?", data.Name).
		Exec()
	return res, err
}

func (c jSample2Querier) Delete() *builder.DeleteBuilder {
	return &builder.DeleteBuilder{
		DeleteBuilder: c.db.DeleteFrom(c.name),
	}
}

func (c jSample2Querier) DeleteByPk(name string) error {
	_, err := c.Delete().Where(

		JSample2Model.Name.Eq(name),
	).Exec()
	return err
}
func (c jSample2Querier) DeleteAll(items ...*Sample2) error {
	q := c.Delete()
	for _, item := range items {
		q = q.Where(
			dbr.Or(
				dbr.And(

					JSample2Model.Name.Eq(item.Name),
				),
			),
		)
	}
	_, err := q.Exec()
	return err
}
func (c jSample2Querier) Find(name string) (*Sample2, error) {
	return c.Select().Where(

		JSample2Model.Name.Eq(name),
	).Read()
}

type jSampleViewSetup struct {
	Name       string
	CreateStmt string
	DropStmt   string
}

func (c jSampleViewSetup) Create(db *dbr.Connection) error {
	_, err := db.Exec(c.CreateStmt)
	return err
}
func (c jSampleViewSetup) Drop(db *dbr.Connection) error {
	_, err := db.Exec(c.DropStmt)
	return err
}

// JSampleViewSetup helps to create/drop the schema
func JSampleViewSetup() runtime.Setuper {
	driver := runtime.GetCurrentDriver()

	var create string
	var drop string

	if driver == drivers.Sqlite {
		create = `CREATE VIEW IF NOT EXISTS sample_view AS 
	SELECT *
	FROM sample
	WHERE id > 1
	-- comments finish on new empty line
`
	} else if driver == drivers.Mysql {
		create = `CREATE VIEW IF NOT EXISTS sample_view AS 
	SELECT *
	FROM sample
	WHERE id > 1
	-- comments finish on new empty line
`
	} else if driver == drivers.Pgsql {
		create = `CREATE VIEW IF NOT EXISTS sample_view AS 
	SELECT *
	FROM sample
	WHERE id > 1
	-- comments finish on new empty line
`
	}

	if driver == drivers.Sqlite {
		drop = `DROP VIEW IF EXISTS sample_view`
	} else if driver == drivers.Mysql {
		drop = `DROP VIEW IF EXISTS sample_view`
	} else if driver == drivers.Pgsql {
		drop = `DROP VIEW IF EXISTS sample_view`
	}

	return jSampleViewSetup{
		Name:       `sample_view`,
		CreateStmt: create,
		DropStmt:   drop,
	}
}

var JSampleViewModel = struct {
	Eq func(...*SampleView) dbr.Builder
	In func(...*SampleView) dbr.Builder

	ID struct {
		Name string
		IsPk bool
		IsAI bool
		Eq   func(interface{}) dbr.Builder
		In   func(...interface{}) dbr.Builder
		Gt   func(interface{}) dbr.Builder
		Gte  func(interface{}) dbr.Builder
		Lt   func(interface{}) dbr.Builder
		Lte  func(interface{}) dbr.Builder
	}

	Name struct {
		Name string
		Eq   func(interface{}) dbr.Builder
		In   func(...interface{}) dbr.Builder
		Like func(interface{}) dbr.Builder
	}

	Description struct {
		Name string
		Eq   func(interface{}) dbr.Builder
		In   func(...interface{}) dbr.Builder
		Like func(interface{}) dbr.Builder
	}
}{

	Eq: func(s ...*SampleView) dbr.Builder {
		ors := []dbr.Builder{}
		for _, t := range s {
			ors = append(ors, dbr.Or(
				dbr.And(

					dbr.Eq(`id`, t.ID),
				),
			))
		}
		return dbr.And(ors...)
	},
	In: func(s ...*SampleView) dbr.Builder {
		ors := []dbr.Builder{}
		for _, t := range s {
			ors = append(ors, dbr.Or(
				dbr.And(

					dbr.Eq(`id`, t.ID),
				),
			))
		}
		return dbr.And(ors...)
	},

	ID: struct {
		Name string
		IsPk bool
		IsAI bool
		Eq   func(interface{}) dbr.Builder
		In   func(...interface{}) dbr.Builder
		Gt   func(interface{}) dbr.Builder
		Gte  func(interface{}) dbr.Builder
		Lt   func(interface{}) dbr.Builder
		Lte  func(interface{}) dbr.Builder
	}{
		Name: `id`,
		IsPk: true,
		IsAI: true,
		Eq: func(v interface{}) dbr.Builder {
			return dbr.Eq(`id`, v)
		},
		In: func(v ...interface{}) dbr.Builder {
			return dbr.Eq(`id`, v)
		},
		Gt: func(v interface{}) dbr.Builder {
			return dbr.Gt(`id`, v)
		},
		Gte: func(v interface{}) dbr.Builder {
			return dbr.Gte(`id`, v)
		},
		Lt: func(v interface{}) dbr.Builder {
			return dbr.Lt(`id`, v)
		},
		Lte: func(v interface{}) dbr.Builder {
			return dbr.Lte(`id`, v)
		},
	},

	Name: struct {
		Name string
		Eq   func(interface{}) dbr.Builder
		In   func(...interface{}) dbr.Builder
		Like func(interface{}) dbr.Builder
	}{
		Name: `name`,
		Eq: func(v interface{}) dbr.Builder {
			return dbr.Eq(`name`, v)
		},
		In: func(v ...interface{}) dbr.Builder {
			return dbr.Eq(`name`, v)
		},
		Like: func(v interface{}) dbr.Builder {
			return builder.Like(`name`, v)
		},
	},

	Description: struct {
		Name string
		Eq   func(interface{}) dbr.Builder
		In   func(...interface{}) dbr.Builder
		Like func(interface{}) dbr.Builder
	}{
		Name: `description`,
		Eq: func(v interface{}) dbr.Builder {
			return dbr.Eq(`description`, v)
		},
		In: func(v ...interface{}) dbr.Builder {
			return dbr.Eq(`description`, v)
		},
		Like: func(v interface{}) dbr.Builder {
			return builder.Like(`description`, v)
		},
	},
}

// JSampleView provides a basic querier
func JSampleView(db dbr.SessionRunner) jSampleViewQuerier {
	return jSampleViewQuerier{
		name: `sample_view`,
		db:   db,
	}
}

type jSampleViewSelectBuilder struct {
	*builder.SelectBuilder
}

func (c *jSampleViewSelectBuilder) Read() (*SampleView, error) {
	var one SampleView
	err := c.LoadStruct(&one)
	return &one, err
}
func (c *jSampleViewSelectBuilder) ReadAll() ([]*SampleView, error) {
	var all []*SampleView
	_, err := c.LoadStructs(&all)
	return all, err
}
func (c *jSampleViewSelectBuilder) Where(query interface{}, value ...interface{}) *jSampleViewSelectBuilder {
	c.SelectBuilder.Where(query, value...)
	return c
}

type jSampleViewQuerier struct {
	name string
	db   dbr.SessionRunner
}

func (c jSampleViewQuerier) Select(what ...string) *jSampleViewSelectBuilder {
	if len(what) == 0 {
		what = append(what, "*")
	}
	return &jSampleViewSelectBuilder{
		&builder.SelectBuilder{
			SelectBuilder: c.db.Select(what...).From(c.name),
		},
	}
}
func (c jSampleViewQuerier) Count(what ...string) *jSampleViewSelectBuilder {
	if len(what) == 0 {
		what = append(what, "*")
	}
	return c.Select("COUNT(" + strings.Join(what, ",") + ")")
}
