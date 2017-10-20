// Generated with mh-cbon/jedi. Do not edit by hand.
package main

import (
	"database/sql"
	"fmt"
	"strings"
	"time"

	"github.com/gocraft/dbr"
	dbrdialect "github.com/gocraft/dbr/dialect"
	"github.com/mh-cbon/jedi/builder"
	"github.com/mh-cbon/jedi/drivers"
	"github.com/mh-cbon/jedi/runtime"
)

var _ = time.Now
var _ = fmt.Sprintf
var _ = dbrdialect.PostgreSQL

func init() {

	Jedi = append(Jedi, JSampleSetup)

	Jedi = append(Jedi, JBasicPKSetup)

	Jedi = append(Jedi, JBasicTypesSetup)

	Jedi = append(Jedi, JTextPkSetup)

	Jedi = append(Jedi, JHasOneTextPkSetup)

	Jedi = append(Jedi, JHasManyTextPkSetup)

	Jedi = append(Jedi, JCompositePkSetup)

	Jedi = append(Jedi, JHasOneCompositePkSetup)

	Jedi = append(Jedi, JHasManyCompositePkSetup)

	Jedi = append(Jedi, JDateTypeSetup)

	Jedi = append(Jedi, JSampleViewSetup)

	Jedi = append(Jedi, JHookDemoSetup)

	Jedi = append(Jedi, JHasManyTextPkrelatedsToTextPkrelatedsSetup)

	Jedi = append(Jedi, JCompositePkrelatedsToHasManyCompositePkrelatedsSetup)

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
removal_date datetime NULL

)`
	} else if driver == drivers.Mysql {
		create = `CREATE TABLE IF NOT EXISTS sample (
id INTEGER NOT NULL AUTO_INCREMENT,
name TEXT,
description TEXT,
update_date datetime(6),
removal_date datetime(6) NULL,
PRIMARY KEY (id) 

)`
	} else if driver == drivers.Pgsql {
		create = `CREATE TABLE IF NOT EXISTS sample (
id SERIAL PRIMARY KEY,
name TEXT,
description TEXT,
update_date timestamp(6),
removal_date timestamp(6) NULL

)`
	}

	if driver == drivers.Sqlite {
		drop = `DROP TABLE IF EXISTS sample`
	} else if driver == drivers.Mysql {
		drop = `DROP TABLE IF EXISTS sample`
	} else if driver == drivers.Pgsql {
		drop = `DROP TABLE IF EXISTS sample`
	}

	var indexes []string

	if driver == drivers.Sqlite {

	} else if driver == drivers.Mysql {

	} else if driver == drivers.Pgsql {

	}

	return runtime.Table{
		Name:       `sample`,
		CreateStmt: create,
		DropStmt:   drop,
		View:       !true,
		Indexes:    indexes,
	}
}

// jSampleModel provides helper to work with Sample data provider
type jSampleModel struct {
	as string

	ID builder.ValuePropertyMeta

	Name builder.ValuePropertyMeta

	Description builder.ValuePropertyMeta

	UpdateDate builder.ValuePropertyMeta

	RemovalDate builder.ValuePropertyMeta
}

// Eq provided items.
func (j jSampleModel) Eq(s ...*Sample) dbr.Builder {
	ors := []dbr.Builder{}
	for _, t := range s {
		ors = append(ors, dbr.And(

			JSampleModel.ID.Eq(t.ID),
		))
	}
	return dbr.Or(ors...)
}

// In provided items.
func (j jSampleModel) In(s ...*Sample) dbr.Builder {
	ors := []dbr.Builder{}
	for _, t := range s {
		ors = append(ors, dbr.And(

			JSampleModel.ID.Eq(t.ID),
		))
	}
	return dbr.Or(ors...)
}

// As returns a copy with an alias.
func (j jSampleModel) As(as string) jSampleModel {
	j.as = as

	j.ID.TableAlias = as

	j.Name.TableAlias = as

	j.Description.TableAlias = as

	j.UpdateDate.TableAlias = as

	j.RemovalDate.TableAlias = as

	return j
}

// Table returns the sql table name
func (j jSampleModel) Table() string {
	return "sample"
}

// Alias returns the current alias
func (j jSampleModel) Alias() string {
	if j.as == "" {
		return j.Table()
	}
	return j.as
}

// Properties returns a map of property name => meta
func (j jSampleModel) Properties() map[string]builder.MetaProvider {
	ret := map[string]builder.MetaProvider{}

	ret["ID"] = j.ID

	ret["Name"] = j.Name

	ret["Description"] = j.Description

	ret["UpdateDate"] = j.UpdateDate

	ret["RemovalDate"] = j.RemovalDate

	return ret
}

// Fields returns given sql fields with appropriate aliasing.
func (j jSampleModel) Fields(ins ...string) []string {
	dialect := runtime.GetDialect()
	if len(ins) == 0 {
		ins = append(ins, "*")
	}
	for i, in := range ins {
		if j.as != "" {
			if in == "*" {
				ins[i] = fmt.Sprintf("%v.%v", dialect.QuoteIdent(j.as), in)
			} else {
				ins[i] = fmt.Sprintf("%v.%v", dialect.QuoteIdent(j.as), dialect.QuoteIdent(in))
			}
		}
	}
	return ins
}

// JSampleModel provides helper to work with Sample data provider
var JSampleModel = jSampleModel{

	ID: builder.NewValueMeta(
		`id`, `INTEGER`,
		`ID`, `int64`,
		true, true,
	),

	Name: builder.NewValueMeta(
		`name`, `TEXT`,
		`Name`, `string`,
		false, false,
	),

	Description: builder.NewValueMeta(
		`description`, `TEXT`,
		`Description`, `string`,
		false, false,
	),

	UpdateDate: builder.NewValueMeta(
		`update_date`, `datetime`,
		`UpdateDate`, `time.Time`,
		false, false,
	),

	RemovalDate: builder.NewValueMeta(
		`removal_date`, `datetime`,
		`RemovalDate`, `*time.Time`,
		false, false,
	),
}

type jSampleDeleteBuilder struct {
	*builder.DeleteBuilder
}

// //Build builds the sql string into given buffer using current dialect
// func (c *jSampleDeleteBuilder) Build(b dbr.Buffer) error {
// 	return c.DeleteBuilder.Build(runtime.GetDialect(), b)
// }
// //String returns the sql string for current dialect. It returns empty string if the build returns an error.
// func (c *jSampleDeleteBuilder) String() string {
// 	b := dbr.NewBuffer()
// 	if err := c.Build(b); err != nil {
// 		return ""
// 	}
// 	return b.String()
// }
//Where returns a jSampleDeleteBuilder instead of builder.DeleteBuilder.
func (c *jSampleDeleteBuilder) Where(query interface{}, value ...interface{}) *jSampleDeleteBuilder {
	c.DeleteBuilder.Where(query, value...)
	return c
}

type jSampleSelectBuilder struct {
	as string
	*builder.SelectBuilder
}

// //Build builds the sql string using current dialect into given bufer
// func (c *jSampleSelectBuilder) Build(b dbr.Buffer) error {
// 	return c.SelectBuilder.Build(runtime.GetDialect(), b)
// }
// //String returns the sql string for current dialect. It returns empty string if the build returns an error.
// func (c *jSampleSelectBuilder) String() string {
// 	b := dbr.NewBuffer()
// 	if err := c.Build(b); err != nil {
// 		return ""
// 	}
// 	return b.String()
// }

//Read evaluates current select query and load the results into a Sample
func (c *jSampleSelectBuilder) Read() (*Sample, error) {
	var one Sample
	err := c.LoadStruct(&one)
	return &one, err
}

//ReadAll evaluates current select query and load the results into a slice of Sample
func (c *jSampleSelectBuilder) ReadAll() ([]*Sample, error) {
	var all []*Sample
	_, err := c.LoadStructs(&all)
	return all, err
}

//Where returns a jSampleSelectBuilder instead of builder.SelectBuilder.
func (c *jSampleSelectBuilder) Where(query interface{}, value ...interface{}) *jSampleSelectBuilder {
	c.SelectBuilder.Where(query, value...)
	return c
}

//GroupBy returns a jSampleSelectBuilder instead of builder.SelectBuilder.
func (c *jSampleSelectBuilder) GroupBy(col ...string) *jSampleSelectBuilder {
	c.SelectBuilder.GroupBy(col...)
	return c
}

//Having returns a jSampleSelectBuilder instead of builder.SelectBuilder.
func (c *jSampleSelectBuilder) Having(query interface{}, value ...interface{}) *jSampleSelectBuilder {
	c.SelectBuilder.Having(query, value...)
	return c
}

//Limit returns a jSampleSelectBuilder instead of builder.SelectBuilder.
func (c *jSampleSelectBuilder) Limit(n uint64) *jSampleSelectBuilder {
	c.SelectBuilder.Limit(n)
	return c
}

//Offset returns a jSampleSelectBuilder instead of builder.SelectBuilder.
func (c *jSampleSelectBuilder) Offset(n uint64) *jSampleSelectBuilder {
	c.SelectBuilder.Offset(n)
	return c
}

//OrderAsc returns a jSampleSelectBuilder instead of builder.SelectBuilder.
func (c *jSampleSelectBuilder) OrderAsc(col string) *jSampleSelectBuilder {
	c.SelectBuilder.OrderAsc(col)
	return c
}

//OrderDesc returns a jSampleSelectBuilder instead of builder.SelectBuilder.
func (c *jSampleSelectBuilder) OrderDesc(col string) *jSampleSelectBuilder {
	c.SelectBuilder.OrderDesc(col)
	return c
}

//OrderDir returns a jSampleSelectBuilder instead of builder.SelectBuilder.
func (c *jSampleSelectBuilder) OrderDir(col string, isAsc bool) *jSampleSelectBuilder {
	c.SelectBuilder.OrderDir(col, isAsc)
	return c
}

//OrderBy returns a jSampleSelectBuilder instead of builder.SelectBuilder.
func (c *jSampleSelectBuilder) OrderBy(col string) *jSampleSelectBuilder {
	c.SelectBuilder.OrderBy(col)
	return c
}

//Paginate returns a jSampleSelectBuilder instead of builder.SelectBuilder.
func (c *jSampleSelectBuilder) Paginate(page, perPage uint64) *jSampleSelectBuilder {
	c.SelectBuilder.Paginate(page, perPage)
	return c
}

//Join returns a jSampleSelectBuilder instead of builder.SelectBuilder.
func (c *jSampleSelectBuilder) Join(table, on interface{}) *jSampleSelectBuilder {
	c.SelectBuilder.Join(table, on)
	return c
}

//LeftJoin returns a jSampleSelectBuilder instead of builder.SelectBuilder.
func (c *jSampleSelectBuilder) LeftJoin(table, on interface{}) *jSampleSelectBuilder {
	c.SelectBuilder.LeftJoin(table, on)
	return c
}

//RightJoin returns a jSampleSelectBuilder instead of builder.SelectBuilder.
func (c *jSampleSelectBuilder) RightJoin(table, on interface{}) *jSampleSelectBuilder {
	c.SelectBuilder.RightJoin(table, on)
	return c
}

//FullJoin returns a jSampleSelectBuilder instead of builder.SelectBuilder.
func (c *jSampleSelectBuilder) FullJoin(table, on interface{}) *jSampleSelectBuilder {
	c.SelectBuilder.FullJoin(table, on)
	return c
}

//Distinct returns a jSampleSelectBuilder instead of builder.SelectBuilder.
func (c *jSampleSelectBuilder) Distinct() *jSampleSelectBuilder {
	c.SelectBuilder.Distinct()
	return c
}

// JSample provides a basic querier
func JSample(db dbr.SessionRunner) jSampleQuerier {
	return jSampleQuerier{
		db: db,
	}
}

type jSampleQuerier struct {
	db dbr.SessionRunner
	as string
}

//As set alias prior building.
func (c jSampleQuerier) As(as string) jSampleQuerier {
	c.as = as
	return c
}

//Model returns a model with appropriate aliasing.
func (c jSampleQuerier) Model() jSampleModel {
	return JSampleModel.As(c.as)
}

//Select returns a Sample Select Builder.
func (c jSampleQuerier) Select(what ...string) *jSampleSelectBuilder {
	m := c.Model()
	dialect := runtime.GetDialect()
	from := dialect.QuoteIdent(m.Table())
	if m.Alias() != "" && m.Alias() != m.Table() {
		from = fmt.Sprintf("%v as %v", from, dialect.QuoteIdent(m.Alias()))
	}
	if len(what) == 0 {
		alias := m.Table()
		if m.Alias() != "" && m.Alias() != m.Table() {
			alias = m.Alias()
		}
		what = m.Fields(alias + ".*")
	}
	return &jSampleSelectBuilder{
		as: c.as,
		SelectBuilder: &builder.SelectBuilder{
			SelectBuilder: c.db.Select(what...).From(from),
		},
	}
}

//Where returns a Sample Select Builder.
func (c jSampleQuerier) Where(query interface{}, value ...interface{}) *jSampleSelectBuilder {
	return c.Select().Where(query, value...)
}

//Count returns a Sample Select Builder to count given expressions.
func (c jSampleQuerier) Count(what ...string) *jSampleSelectBuilder {
	if len(what) == 0 {
		what = append(what, "*")
	}
	return c.Select("COUNT(" + strings.Join(what, ",") + ")")
}

// Insert a new Sample, if it has autoincrement primary key, the value will be set.
// It stops on first error.
func (c jSampleQuerier) Insert(items ...*Sample) (sql.Result, error) {
	var res sql.Result
	var err error
	for _, data := range items {

		data.UpdateDate = data.UpdateDate.UTC()

		if data.RemovalDate != nil {
			x := data.RemovalDate.UTC()
			data.RemovalDate = &x
		}

		query := c.db.InsertInto(JSampleModel.Table()).Columns(

			`name`,

			`description`,

			`update_date`,

			`removal_date`,
		).Record(data)
		if runtime.Runs(drivers.Pgsql) {

			query = query.Returning(

				`id`,
			)

			var auto0 int64

			err = query.Load(

				&auto0,
			)

			data.ID = auto0

		} else {
			res, err = query.Exec()

			if err == nil {
				id, err2 := res.LastInsertId()
				if err2 != nil {
					return res, err2
				}
				data.ID = id
			}

		}
		if err != nil {
			return res, err
		}
	}
	return res, err
}

// InsertBulk inserts multiple items into the database.
// It does not post update any auto increment field.
// It builds an insert query of multiple rows and send it on the underlying connection.
func (c jSampleQuerier) InsertBulk(items ...*Sample) error {
	panic("todo")
}

// Update a Sample. It stops on first error.
func (c jSampleQuerier) Update(items ...*Sample) (sql.Result, error) {
	var res sql.Result
	var err error
	for _, data := range items {

		data.UpdateDate = data.UpdateDate.UTC()

		if data.RemovalDate != nil {
			x := data.RemovalDate.UTC()
			data.RemovalDate = &x
		}

		query := c.db.Update(JSampleModel.Table())

		query = query.Set(`name`, data.Name)

		query = query.Set(`description`, data.Description)

		query = query.Set(`update_date`, data.UpdateDate)

		query = query.Set(`removal_date`, data.RemovalDate)

		query = query.Where("id = ?", data.ID)

		res, err = query.Exec()

		if err != nil {
			return res, err
		}
	}
	return res, err
}

// MustUpdate a Sample. It stops on first error. It errors if an update query does not affect row.
func (c jSampleQuerier) MustUpdate(items ...*Sample) (sql.Result, error) {
	var res sql.Result
	var err error
	for _, data := range items {

		data.UpdateDate = data.UpdateDate.UTC()

		if data.RemovalDate != nil {
			x := data.RemovalDate.UTC()
			data.RemovalDate = &x
		}

		query := c.db.Update(JSampleModel.Table())

		query = query.Set(`name`, data.Name)

		query = query.Set(`description`, data.Description)

		query = query.Set(`update_date`, data.UpdateDate)

		query = query.Set(`removal_date`, data.RemovalDate)

		query = query.Where("id = ?", data.ID)

		res, err = query.Exec()

		if err == nil {
			if n, _ := res.RowsAffected(); n == 0 {
				x := &builder.UpdateBuilder{UpdateBuilder: query}
				err = runtime.NewNoRowsAffected(x.String())
			}
		}

		if err != nil {
			return res, err
		}
	}
	return res, err
}

// UpdateBulk updates multiple items into the database.
// It builds an update query of multiple rows and send it on the underlying connection.
func (c jSampleQuerier) UpdateBulk(items ...*Sample) error {
	panic("todo")
}

//Delete returns a delete builder
func (c jSampleQuerier) Delete() *jSampleDeleteBuilder {
	return &jSampleDeleteBuilder{
		&builder.DeleteBuilder{
			DeleteBuilder: c.db.DeleteFrom(JSampleModel.Table()),
		},
	}
}

// MustDelete requires the query to affeect rows.
func (c jSampleQuerier) MustDelete() *jSampleDeleteBuilder {
	ret := &jSampleDeleteBuilder{
		&builder.DeleteBuilder{
			DeleteBuilder: c.db.DeleteFrom(JSampleModel.Table()),
		},
	}
	ret.MustDelete()
	return ret
}

//DeleteByPk deletes one Sample by its PKs
func (c jSampleQuerier) DeleteByPk(ID int64) error {
	_, err := c.Delete().Where(

		JSampleModel.ID.Eq(ID),
	).Exec()
	return err
}

// DeleteAll given Sample
func (c jSampleQuerier) DeleteAll(items ...*Sample) (sql.Result, error) {
	q := c.Delete().Where(
		JSampleModel.In(items...),
	)
	return q.Exec()
}

// MustDeleteAll given Sample
func (c jSampleQuerier) MustDeleteAll(items ...*Sample) (sql.Result, error) {
	var res sql.Result
	var err error
	for _, d := range items {
		res, err = c.DeleteAll(d)
		if err != nil {
			return res, err
		}
		if n, e := res.RowsAffected(); e != nil {
			return res, e
		} else if n == 0 {
			q := c.Delete().Where(
				JSampleModel.In(items...),
			)
			err = runtime.NewNoRowsAffected(q.String())
			return res, err
		}
	}
	return res, err
}

//Find one Sample using its PKs
func (c jSampleQuerier) Find(ID int64) (*Sample, error) {
	return c.Select().Where(

		JSampleModel.ID.Eq(ID),
	).Read()
}

// JBasicPKSetup helps to create/drop the schema
func JBasicPKSetup() runtime.Setuper {
	driver := runtime.GetCurrentDriver()

	var create string
	var drop string

	if driver == drivers.Sqlite {
		create = `CREATE TABLE IF NOT EXISTS basic_pk (
id INTEGER PRIMARY KEY AUTOINCREMENT,
whatever TEXT

)`
	} else if driver == drivers.Mysql {
		create = `CREATE TABLE IF NOT EXISTS basic_pk (
id INTEGER NOT NULL AUTO_INCREMENT,
whatever TEXT,
PRIMARY KEY (id) 

)`
	} else if driver == drivers.Pgsql {
		create = `CREATE TABLE IF NOT EXISTS basic_pk (
id SERIAL PRIMARY KEY,
whatever TEXT

)`
	}

	if driver == drivers.Sqlite {
		drop = `DROP TABLE IF EXISTS basic_pk`
	} else if driver == drivers.Mysql {
		drop = `DROP TABLE IF EXISTS basic_pk`
	} else if driver == drivers.Pgsql {
		drop = `DROP TABLE IF EXISTS basic_pk`
	}

	var indexes []string

	if driver == drivers.Sqlite {

	} else if driver == drivers.Mysql {

	} else if driver == drivers.Pgsql {

	}

	return runtime.Table{
		Name:       `basic_pk`,
		CreateStmt: create,
		DropStmt:   drop,
		View:       !true,
		Indexes:    indexes,
	}
}

// jBasicPKModel provides helper to work with BasicPK data provider
type jBasicPKModel struct {
	as string

	ID builder.ValuePropertyMeta

	Whatever builder.ValuePropertyMeta
}

// Eq provided items.
func (j jBasicPKModel) Eq(s ...*BasicPK) dbr.Builder {
	ors := []dbr.Builder{}
	for _, t := range s {
		ors = append(ors, dbr.And(

			JBasicPKModel.ID.Eq(t.ID),
		))
	}
	return dbr.Or(ors...)
}

// In provided items.
func (j jBasicPKModel) In(s ...*BasicPK) dbr.Builder {
	ors := []dbr.Builder{}
	for _, t := range s {
		ors = append(ors, dbr.And(

			JBasicPKModel.ID.Eq(t.ID),
		))
	}
	return dbr.Or(ors...)
}

// As returns a copy with an alias.
func (j jBasicPKModel) As(as string) jBasicPKModel {
	j.as = as

	j.ID.TableAlias = as

	j.Whatever.TableAlias = as

	return j
}

// Table returns the sql table name
func (j jBasicPKModel) Table() string {
	return "basic_pk"
}

// Alias returns the current alias
func (j jBasicPKModel) Alias() string {
	if j.as == "" {
		return j.Table()
	}
	return j.as
}

// Properties returns a map of property name => meta
func (j jBasicPKModel) Properties() map[string]builder.MetaProvider {
	ret := map[string]builder.MetaProvider{}

	ret["ID"] = j.ID

	ret["Whatever"] = j.Whatever

	return ret
}

// Fields returns given sql fields with appropriate aliasing.
func (j jBasicPKModel) Fields(ins ...string) []string {
	dialect := runtime.GetDialect()
	if len(ins) == 0 {
		ins = append(ins, "*")
	}
	for i, in := range ins {
		if j.as != "" {
			if in == "*" {
				ins[i] = fmt.Sprintf("%v.%v", dialect.QuoteIdent(j.as), in)
			} else {
				ins[i] = fmt.Sprintf("%v.%v", dialect.QuoteIdent(j.as), dialect.QuoteIdent(in))
			}
		}
	}
	return ins
}

// JBasicPKModel provides helper to work with BasicPK data provider
var JBasicPKModel = jBasicPKModel{

	ID: builder.NewValueMeta(
		`id`, `INTEGER`,
		`ID`, `int64`,
		true, true,
	),

	Whatever: builder.NewValueMeta(
		`whatever`, `TEXT`,
		`Whatever`, `string`,
		false, false,
	),
}

type jBasicPKDeleteBuilder struct {
	*builder.DeleteBuilder
}

// //Build builds the sql string into given buffer using current dialect
// func (c *jBasicPKDeleteBuilder) Build(b dbr.Buffer) error {
// 	return c.DeleteBuilder.Build(runtime.GetDialect(), b)
// }
// //String returns the sql string for current dialect. It returns empty string if the build returns an error.
// func (c *jBasicPKDeleteBuilder) String() string {
// 	b := dbr.NewBuffer()
// 	if err := c.Build(b); err != nil {
// 		return ""
// 	}
// 	return b.String()
// }
//Where returns a jBasicPKDeleteBuilder instead of builder.DeleteBuilder.
func (c *jBasicPKDeleteBuilder) Where(query interface{}, value ...interface{}) *jBasicPKDeleteBuilder {
	c.DeleteBuilder.Where(query, value...)
	return c
}

type jBasicPKSelectBuilder struct {
	as string
	*builder.SelectBuilder
}

// //Build builds the sql string using current dialect into given bufer
// func (c *jBasicPKSelectBuilder) Build(b dbr.Buffer) error {
// 	return c.SelectBuilder.Build(runtime.GetDialect(), b)
// }
// //String returns the sql string for current dialect. It returns empty string if the build returns an error.
// func (c *jBasicPKSelectBuilder) String() string {
// 	b := dbr.NewBuffer()
// 	if err := c.Build(b); err != nil {
// 		return ""
// 	}
// 	return b.String()
// }

//Read evaluates current select query and load the results into a BasicPK
func (c *jBasicPKSelectBuilder) Read() (*BasicPK, error) {
	var one BasicPK
	err := c.LoadStruct(&one)
	return &one, err
}

//ReadAll evaluates current select query and load the results into a slice of BasicPK
func (c *jBasicPKSelectBuilder) ReadAll() ([]*BasicPK, error) {
	var all []*BasicPK
	_, err := c.LoadStructs(&all)
	return all, err
}

//Where returns a jBasicPKSelectBuilder instead of builder.SelectBuilder.
func (c *jBasicPKSelectBuilder) Where(query interface{}, value ...interface{}) *jBasicPKSelectBuilder {
	c.SelectBuilder.Where(query, value...)
	return c
}

//GroupBy returns a jBasicPKSelectBuilder instead of builder.SelectBuilder.
func (c *jBasicPKSelectBuilder) GroupBy(col ...string) *jBasicPKSelectBuilder {
	c.SelectBuilder.GroupBy(col...)
	return c
}

//Having returns a jBasicPKSelectBuilder instead of builder.SelectBuilder.
func (c *jBasicPKSelectBuilder) Having(query interface{}, value ...interface{}) *jBasicPKSelectBuilder {
	c.SelectBuilder.Having(query, value...)
	return c
}

//Limit returns a jBasicPKSelectBuilder instead of builder.SelectBuilder.
func (c *jBasicPKSelectBuilder) Limit(n uint64) *jBasicPKSelectBuilder {
	c.SelectBuilder.Limit(n)
	return c
}

//Offset returns a jBasicPKSelectBuilder instead of builder.SelectBuilder.
func (c *jBasicPKSelectBuilder) Offset(n uint64) *jBasicPKSelectBuilder {
	c.SelectBuilder.Offset(n)
	return c
}

//OrderAsc returns a jBasicPKSelectBuilder instead of builder.SelectBuilder.
func (c *jBasicPKSelectBuilder) OrderAsc(col string) *jBasicPKSelectBuilder {
	c.SelectBuilder.OrderAsc(col)
	return c
}

//OrderDesc returns a jBasicPKSelectBuilder instead of builder.SelectBuilder.
func (c *jBasicPKSelectBuilder) OrderDesc(col string) *jBasicPKSelectBuilder {
	c.SelectBuilder.OrderDesc(col)
	return c
}

//OrderDir returns a jBasicPKSelectBuilder instead of builder.SelectBuilder.
func (c *jBasicPKSelectBuilder) OrderDir(col string, isAsc bool) *jBasicPKSelectBuilder {
	c.SelectBuilder.OrderDir(col, isAsc)
	return c
}

//OrderBy returns a jBasicPKSelectBuilder instead of builder.SelectBuilder.
func (c *jBasicPKSelectBuilder) OrderBy(col string) *jBasicPKSelectBuilder {
	c.SelectBuilder.OrderBy(col)
	return c
}

//Paginate returns a jBasicPKSelectBuilder instead of builder.SelectBuilder.
func (c *jBasicPKSelectBuilder) Paginate(page, perPage uint64) *jBasicPKSelectBuilder {
	c.SelectBuilder.Paginate(page, perPage)
	return c
}

//Join returns a jBasicPKSelectBuilder instead of builder.SelectBuilder.
func (c *jBasicPKSelectBuilder) Join(table, on interface{}) *jBasicPKSelectBuilder {
	c.SelectBuilder.Join(table, on)
	return c
}

//LeftJoin returns a jBasicPKSelectBuilder instead of builder.SelectBuilder.
func (c *jBasicPKSelectBuilder) LeftJoin(table, on interface{}) *jBasicPKSelectBuilder {
	c.SelectBuilder.LeftJoin(table, on)
	return c
}

//RightJoin returns a jBasicPKSelectBuilder instead of builder.SelectBuilder.
func (c *jBasicPKSelectBuilder) RightJoin(table, on interface{}) *jBasicPKSelectBuilder {
	c.SelectBuilder.RightJoin(table, on)
	return c
}

//FullJoin returns a jBasicPKSelectBuilder instead of builder.SelectBuilder.
func (c *jBasicPKSelectBuilder) FullJoin(table, on interface{}) *jBasicPKSelectBuilder {
	c.SelectBuilder.FullJoin(table, on)
	return c
}

//Distinct returns a jBasicPKSelectBuilder instead of builder.SelectBuilder.
func (c *jBasicPKSelectBuilder) Distinct() *jBasicPKSelectBuilder {
	c.SelectBuilder.Distinct()
	return c
}

// JBasicPK provides a basic querier
func JBasicPK(db dbr.SessionRunner) jBasicPKQuerier {
	return jBasicPKQuerier{
		db: db,
	}
}

type jBasicPKQuerier struct {
	db dbr.SessionRunner
	as string
}

//As set alias prior building.
func (c jBasicPKQuerier) As(as string) jBasicPKQuerier {
	c.as = as
	return c
}

//Model returns a model with appropriate aliasing.
func (c jBasicPKQuerier) Model() jBasicPKModel {
	return JBasicPKModel.As(c.as)
}

//Select returns a BasicPK Select Builder.
func (c jBasicPKQuerier) Select(what ...string) *jBasicPKSelectBuilder {
	m := c.Model()
	dialect := runtime.GetDialect()
	from := dialect.QuoteIdent(m.Table())
	if m.Alias() != "" && m.Alias() != m.Table() {
		from = fmt.Sprintf("%v as %v", from, dialect.QuoteIdent(m.Alias()))
	}
	if len(what) == 0 {
		alias := m.Table()
		if m.Alias() != "" && m.Alias() != m.Table() {
			alias = m.Alias()
		}
		what = m.Fields(alias + ".*")
	}
	return &jBasicPKSelectBuilder{
		as: c.as,
		SelectBuilder: &builder.SelectBuilder{
			SelectBuilder: c.db.Select(what...).From(from),
		},
	}
}

//Where returns a BasicPK Select Builder.
func (c jBasicPKQuerier) Where(query interface{}, value ...interface{}) *jBasicPKSelectBuilder {
	return c.Select().Where(query, value...)
}

//Count returns a BasicPK Select Builder to count given expressions.
func (c jBasicPKQuerier) Count(what ...string) *jBasicPKSelectBuilder {
	if len(what) == 0 {
		what = append(what, "*")
	}
	return c.Select("COUNT(" + strings.Join(what, ",") + ")")
}

// Insert a new BasicPK, if it has autoincrement primary key, the value will be set.
// It stops on first error.
func (c jBasicPKQuerier) Insert(items ...*BasicPK) (sql.Result, error) {
	var res sql.Result
	var err error
	for _, data := range items {

		query := c.db.InsertInto(JBasicPKModel.Table()).Columns(

			`whatever`,
		).Record(data)
		if runtime.Runs(drivers.Pgsql) {

			query = query.Returning(

				`id`,
			)

			var auto0 int64

			err = query.Load(

				&auto0,
			)

			data.ID = auto0

		} else {
			res, err = query.Exec()

			if err == nil {
				id, err2 := res.LastInsertId()
				if err2 != nil {
					return res, err2
				}
				data.ID = id
			}

		}
		if err != nil {
			return res, err
		}
	}
	return res, err
}

// InsertBulk inserts multiple items into the database.
// It does not post update any auto increment field.
// It builds an insert query of multiple rows and send it on the underlying connection.
func (c jBasicPKQuerier) InsertBulk(items ...*BasicPK) error {
	panic("todo")
}

// Update a BasicPK. It stops on first error.
func (c jBasicPKQuerier) Update(items ...*BasicPK) (sql.Result, error) {
	var res sql.Result
	var err error
	for _, data := range items {

		query := c.db.Update(JBasicPKModel.Table())

		query = query.Set(`whatever`, data.Whatever)

		query = query.Where("id = ?", data.ID)

		res, err = query.Exec()

		if err != nil {
			return res, err
		}
	}
	return res, err
}

// MustUpdate a BasicPK. It stops on first error. It errors if an update query does not affect row.
func (c jBasicPKQuerier) MustUpdate(items ...*BasicPK) (sql.Result, error) {
	var res sql.Result
	var err error
	for _, data := range items {

		query := c.db.Update(JBasicPKModel.Table())

		query = query.Set(`whatever`, data.Whatever)

		query = query.Where("id = ?", data.ID)

		res, err = query.Exec()

		if err == nil {
			if n, _ := res.RowsAffected(); n == 0 {
				x := &builder.UpdateBuilder{UpdateBuilder: query}
				err = runtime.NewNoRowsAffected(x.String())
			}
		}

		if err != nil {
			return res, err
		}
	}
	return res, err
}

// UpdateBulk updates multiple items into the database.
// It builds an update query of multiple rows and send it on the underlying connection.
func (c jBasicPKQuerier) UpdateBulk(items ...*BasicPK) error {
	panic("todo")
}

//Delete returns a delete builder
func (c jBasicPKQuerier) Delete() *jBasicPKDeleteBuilder {
	return &jBasicPKDeleteBuilder{
		&builder.DeleteBuilder{
			DeleteBuilder: c.db.DeleteFrom(JBasicPKModel.Table()),
		},
	}
}

// MustDelete requires the query to affeect rows.
func (c jBasicPKQuerier) MustDelete() *jBasicPKDeleteBuilder {
	ret := &jBasicPKDeleteBuilder{
		&builder.DeleteBuilder{
			DeleteBuilder: c.db.DeleteFrom(JBasicPKModel.Table()),
		},
	}
	ret.MustDelete()
	return ret
}

//DeleteByPk deletes one BasicPK by its PKs
func (c jBasicPKQuerier) DeleteByPk(ID int64) error {
	_, err := c.Delete().Where(

		JBasicPKModel.ID.Eq(ID),
	).Exec()
	return err
}

// DeleteAll given BasicPK
func (c jBasicPKQuerier) DeleteAll(items ...*BasicPK) (sql.Result, error) {
	q := c.Delete().Where(
		JBasicPKModel.In(items...),
	)
	return q.Exec()
}

// MustDeleteAll given BasicPK
func (c jBasicPKQuerier) MustDeleteAll(items ...*BasicPK) (sql.Result, error) {
	var res sql.Result
	var err error
	for _, d := range items {
		res, err = c.DeleteAll(d)
		if err != nil {
			return res, err
		}
		if n, e := res.RowsAffected(); e != nil {
			return res, e
		} else if n == 0 {
			q := c.Delete().Where(
				JBasicPKModel.In(items...),
			)
			err = runtime.NewNoRowsAffected(q.String())
			return res, err
		}
	}
	return res, err
}

//Find one BasicPK using its PKs
func (c jBasicPKQuerier) Find(ID int64) (*BasicPK, error) {
	return c.Select().Where(

		JBasicPKModel.ID.Eq(ID),
	).Read()
}

// JBasicTypesSetup helps to create/drop the schema
func JBasicTypesSetup() runtime.Setuper {
	driver := runtime.GetCurrentDriver()

	var create string
	var drop string

	if driver == drivers.Sqlite {
		create = `CREATE TABLE IF NOT EXISTS basic_types (
id INTEGER PRIMARY KEY AUTOINCREMENT,
string TEXT,
string_p TEXT NULL,
intfield INTEGER,
int_p INTEGER NULL,
int32 INTEGER,
int32_p INTEGER NULL,
int64 INTEGER,
int64_p INTEGER NULL,
u_int INTEGER,
u_int_p INTEGER NULL,
u_int32 INTEGER,
u_int32_p INTEGER NULL,
u_int64 INTEGER,
u_int64_p INTEGER NULL,
bool INTEGER,
bool_p INTEGER NULL,
float32 FLOAT,
float32_p FLOAT NULL,
float64 FLOAT,
float64_p FLOAT NULL

)`
	} else if driver == drivers.Mysql {
		create = `CREATE TABLE IF NOT EXISTS basic_types (
id INTEGER NOT NULL AUTO_INCREMENT,
string TEXT,
string_p TEXT NULL,
intfield INTEGER,
int_p INTEGER NULL,
int32 INTEGER,
int32_p INTEGER NULL,
int64 INTEGER,
int64_p INTEGER NULL,
u_int INTEGER,
u_int_p INTEGER NULL,
u_int32 INTEGER,
u_int32_p INTEGER NULL,
u_int64 INTEGER,
u_int64_p INTEGER NULL,
bool INTEGER,
bool_p INTEGER NULL,
float32 FLOAT,
float32_p FLOAT NULL,
float64 FLOAT,
float64_p FLOAT NULL,
PRIMARY KEY (id) 

)`
	} else if driver == drivers.Pgsql {
		create = `CREATE TABLE IF NOT EXISTS basic_types (
id SERIAL PRIMARY KEY,
string TEXT,
string_p TEXT NULL,
intfield INTEGER,
int_p INTEGER NULL,
int32 INTEGER,
int32_p INTEGER NULL,
int64 INTEGER,
int64_p INTEGER NULL,
u_int INTEGER,
u_int_p INTEGER NULL,
u_int32 INTEGER,
u_int32_p INTEGER NULL,
u_int64 INTEGER,
u_int64_p INTEGER NULL,
bool INTEGER,
bool_p INTEGER NULL,
float32 FLOAT,
float32_p FLOAT NULL,
float64 FLOAT,
float64_p FLOAT NULL

)`
	}

	if driver == drivers.Sqlite {
		drop = `DROP TABLE IF EXISTS basic_types`
	} else if driver == drivers.Mysql {
		drop = `DROP TABLE IF EXISTS basic_types`
	} else if driver == drivers.Pgsql {
		drop = `DROP TABLE IF EXISTS basic_types`
	}

	var indexes []string

	if driver == drivers.Sqlite {

	} else if driver == drivers.Mysql {

	} else if driver == drivers.Pgsql {

	}

	return runtime.Table{
		Name:       `basic_types`,
		CreateStmt: create,
		DropStmt:   drop,
		View:       !true,
		Indexes:    indexes,
	}
}

// jBasicTypesModel provides helper to work with BasicTypes data provider
type jBasicTypesModel struct {
	as string

	ID builder.ValuePropertyMeta

	String builder.ValuePropertyMeta

	StringP builder.ValuePropertyMeta

	Int builder.ValuePropertyMeta

	IntP builder.ValuePropertyMeta

	Int32 builder.ValuePropertyMeta

	Int32P builder.ValuePropertyMeta

	Int64 builder.ValuePropertyMeta

	Int64P builder.ValuePropertyMeta

	UInt builder.ValuePropertyMeta

	UIntP builder.ValuePropertyMeta

	UInt32 builder.ValuePropertyMeta

	UInt32P builder.ValuePropertyMeta

	UInt64 builder.ValuePropertyMeta

	UInt64P builder.ValuePropertyMeta

	Bool builder.ValuePropertyMeta

	BoolP builder.ValuePropertyMeta

	Float32 builder.ValuePropertyMeta

	Float32P builder.ValuePropertyMeta

	Float64 builder.ValuePropertyMeta

	Float64P builder.ValuePropertyMeta
}

// Eq provided items.
func (j jBasicTypesModel) Eq(s ...*BasicTypes) dbr.Builder {
	ors := []dbr.Builder{}
	for _, t := range s {
		ors = append(ors, dbr.And(

			JBasicTypesModel.ID.Eq(t.ID),
		))
	}
	return dbr.Or(ors...)
}

// In provided items.
func (j jBasicTypesModel) In(s ...*BasicTypes) dbr.Builder {
	ors := []dbr.Builder{}
	for _, t := range s {
		ors = append(ors, dbr.And(

			JBasicTypesModel.ID.Eq(t.ID),
		))
	}
	return dbr.Or(ors...)
}

// As returns a copy with an alias.
func (j jBasicTypesModel) As(as string) jBasicTypesModel {
	j.as = as

	j.ID.TableAlias = as

	j.String.TableAlias = as

	j.StringP.TableAlias = as

	j.Int.TableAlias = as

	j.IntP.TableAlias = as

	j.Int32.TableAlias = as

	j.Int32P.TableAlias = as

	j.Int64.TableAlias = as

	j.Int64P.TableAlias = as

	j.UInt.TableAlias = as

	j.UIntP.TableAlias = as

	j.UInt32.TableAlias = as

	j.UInt32P.TableAlias = as

	j.UInt64.TableAlias = as

	j.UInt64P.TableAlias = as

	j.Bool.TableAlias = as

	j.BoolP.TableAlias = as

	j.Float32.TableAlias = as

	j.Float32P.TableAlias = as

	j.Float64.TableAlias = as

	j.Float64P.TableAlias = as

	return j
}

// Table returns the sql table name
func (j jBasicTypesModel) Table() string {
	return "basic_types"
}

// Alias returns the current alias
func (j jBasicTypesModel) Alias() string {
	if j.as == "" {
		return j.Table()
	}
	return j.as
}

// Properties returns a map of property name => meta
func (j jBasicTypesModel) Properties() map[string]builder.MetaProvider {
	ret := map[string]builder.MetaProvider{}

	ret["ID"] = j.ID

	ret["String"] = j.String

	ret["StringP"] = j.StringP

	ret["Int"] = j.Int

	ret["IntP"] = j.IntP

	ret["Int32"] = j.Int32

	ret["Int32P"] = j.Int32P

	ret["Int64"] = j.Int64

	ret["Int64P"] = j.Int64P

	ret["UInt"] = j.UInt

	ret["UIntP"] = j.UIntP

	ret["UInt32"] = j.UInt32

	ret["UInt32P"] = j.UInt32P

	ret["UInt64"] = j.UInt64

	ret["UInt64P"] = j.UInt64P

	ret["Bool"] = j.Bool

	ret["BoolP"] = j.BoolP

	ret["Float32"] = j.Float32

	ret["Float32P"] = j.Float32P

	ret["Float64"] = j.Float64

	ret["Float64P"] = j.Float64P

	return ret
}

// Fields returns given sql fields with appropriate aliasing.
func (j jBasicTypesModel) Fields(ins ...string) []string {
	dialect := runtime.GetDialect()
	if len(ins) == 0 {
		ins = append(ins, "*")
	}
	for i, in := range ins {
		if j.as != "" {
			if in == "*" {
				ins[i] = fmt.Sprintf("%v.%v", dialect.QuoteIdent(j.as), in)
			} else {
				ins[i] = fmt.Sprintf("%v.%v", dialect.QuoteIdent(j.as), dialect.QuoteIdent(in))
			}
		}
	}
	return ins
}

// JBasicTypesModel provides helper to work with BasicTypes data provider
var JBasicTypesModel = jBasicTypesModel{

	ID: builder.NewValueMeta(
		`id`, `INTEGER`,
		`ID`, `int64`,
		true, true,
	),

	String: builder.NewValueMeta(
		`string`, `TEXT`,
		`String`, `string`,
		false, false,
	),

	StringP: builder.NewValueMeta(
		`string_p`, `TEXT`,
		`StringP`, `*string`,
		false, false,
	),

	Int: builder.NewValueMeta(
		`intfield`, `INTEGER`,
		`Int`, `int`,
		false, false,
	),

	IntP: builder.NewValueMeta(
		`int_p`, `INTEGER`,
		`IntP`, `*int`,
		false, false,
	),

	Int32: builder.NewValueMeta(
		`int32`, `INTEGER`,
		`Int32`, `int32`,
		false, false,
	),

	Int32P: builder.NewValueMeta(
		`int32_p`, `INTEGER`,
		`Int32P`, `*int32`,
		false, false,
	),

	Int64: builder.NewValueMeta(
		`int64`, `INTEGER`,
		`Int64`, `int64`,
		false, false,
	),

	Int64P: builder.NewValueMeta(
		`int64_p`, `INTEGER`,
		`Int64P`, `*int64`,
		false, false,
	),

	UInt: builder.NewValueMeta(
		`u_int`, `INTEGER`,
		`UInt`, `uint`,
		false, false,
	),

	UIntP: builder.NewValueMeta(
		`u_int_p`, `INTEGER`,
		`UIntP`, `*uint`,
		false, false,
	),

	UInt32: builder.NewValueMeta(
		`u_int32`, `INTEGER`,
		`UInt32`, `uint32`,
		false, false,
	),

	UInt32P: builder.NewValueMeta(
		`u_int32_p`, `INTEGER`,
		`UInt32P`, `*uint32`,
		false, false,
	),

	UInt64: builder.NewValueMeta(
		`u_int64`, `INTEGER`,
		`UInt64`, `uint64`,
		false, false,
	),

	UInt64P: builder.NewValueMeta(
		`u_int64_p`, `INTEGER`,
		`UInt64P`, `*uint64`,
		false, false,
	),

	Bool: builder.NewValueMeta(
		`bool`, `INTEGER`,
		`Bool`, `bool`,
		false, false,
	),

	BoolP: builder.NewValueMeta(
		`bool_p`, `INTEGER`,
		`BoolP`, `*bool`,
		false, false,
	),

	Float32: builder.NewValueMeta(
		`float32`, `FLOAT`,
		`Float32`, `float32`,
		false, false,
	),

	Float32P: builder.NewValueMeta(
		`float32_p`, `FLOAT`,
		`Float32P`, `*float32`,
		false, false,
	),

	Float64: builder.NewValueMeta(
		`float64`, `FLOAT`,
		`Float64`, `float64`,
		false, false,
	),

	Float64P: builder.NewValueMeta(
		`float64_p`, `FLOAT`,
		`Float64P`, `*float64`,
		false, false,
	),
}

type jBasicTypesDeleteBuilder struct {
	*builder.DeleteBuilder
}

// //Build builds the sql string into given buffer using current dialect
// func (c *jBasicTypesDeleteBuilder) Build(b dbr.Buffer) error {
// 	return c.DeleteBuilder.Build(runtime.GetDialect(), b)
// }
// //String returns the sql string for current dialect. It returns empty string if the build returns an error.
// func (c *jBasicTypesDeleteBuilder) String() string {
// 	b := dbr.NewBuffer()
// 	if err := c.Build(b); err != nil {
// 		return ""
// 	}
// 	return b.String()
// }
//Where returns a jBasicTypesDeleteBuilder instead of builder.DeleteBuilder.
func (c *jBasicTypesDeleteBuilder) Where(query interface{}, value ...interface{}) *jBasicTypesDeleteBuilder {
	c.DeleteBuilder.Where(query, value...)
	return c
}

type jBasicTypesSelectBuilder struct {
	as string
	*builder.SelectBuilder
}

// //Build builds the sql string using current dialect into given bufer
// func (c *jBasicTypesSelectBuilder) Build(b dbr.Buffer) error {
// 	return c.SelectBuilder.Build(runtime.GetDialect(), b)
// }
// //String returns the sql string for current dialect. It returns empty string if the build returns an error.
// func (c *jBasicTypesSelectBuilder) String() string {
// 	b := dbr.NewBuffer()
// 	if err := c.Build(b); err != nil {
// 		return ""
// 	}
// 	return b.String()
// }

//Read evaluates current select query and load the results into a BasicTypes
func (c *jBasicTypesSelectBuilder) Read() (*BasicTypes, error) {
	var one BasicTypes
	err := c.LoadStruct(&one)
	return &one, err
}

//ReadAll evaluates current select query and load the results into a slice of BasicTypes
func (c *jBasicTypesSelectBuilder) ReadAll() ([]*BasicTypes, error) {
	var all []*BasicTypes
	_, err := c.LoadStructs(&all)
	return all, err
}

//Where returns a jBasicTypesSelectBuilder instead of builder.SelectBuilder.
func (c *jBasicTypesSelectBuilder) Where(query interface{}, value ...interface{}) *jBasicTypesSelectBuilder {
	c.SelectBuilder.Where(query, value...)
	return c
}

//GroupBy returns a jBasicTypesSelectBuilder instead of builder.SelectBuilder.
func (c *jBasicTypesSelectBuilder) GroupBy(col ...string) *jBasicTypesSelectBuilder {
	c.SelectBuilder.GroupBy(col...)
	return c
}

//Having returns a jBasicTypesSelectBuilder instead of builder.SelectBuilder.
func (c *jBasicTypesSelectBuilder) Having(query interface{}, value ...interface{}) *jBasicTypesSelectBuilder {
	c.SelectBuilder.Having(query, value...)
	return c
}

//Limit returns a jBasicTypesSelectBuilder instead of builder.SelectBuilder.
func (c *jBasicTypesSelectBuilder) Limit(n uint64) *jBasicTypesSelectBuilder {
	c.SelectBuilder.Limit(n)
	return c
}

//Offset returns a jBasicTypesSelectBuilder instead of builder.SelectBuilder.
func (c *jBasicTypesSelectBuilder) Offset(n uint64) *jBasicTypesSelectBuilder {
	c.SelectBuilder.Offset(n)
	return c
}

//OrderAsc returns a jBasicTypesSelectBuilder instead of builder.SelectBuilder.
func (c *jBasicTypesSelectBuilder) OrderAsc(col string) *jBasicTypesSelectBuilder {
	c.SelectBuilder.OrderAsc(col)
	return c
}

//OrderDesc returns a jBasicTypesSelectBuilder instead of builder.SelectBuilder.
func (c *jBasicTypesSelectBuilder) OrderDesc(col string) *jBasicTypesSelectBuilder {
	c.SelectBuilder.OrderDesc(col)
	return c
}

//OrderDir returns a jBasicTypesSelectBuilder instead of builder.SelectBuilder.
func (c *jBasicTypesSelectBuilder) OrderDir(col string, isAsc bool) *jBasicTypesSelectBuilder {
	c.SelectBuilder.OrderDir(col, isAsc)
	return c
}

//OrderBy returns a jBasicTypesSelectBuilder instead of builder.SelectBuilder.
func (c *jBasicTypesSelectBuilder) OrderBy(col string) *jBasicTypesSelectBuilder {
	c.SelectBuilder.OrderBy(col)
	return c
}

//Paginate returns a jBasicTypesSelectBuilder instead of builder.SelectBuilder.
func (c *jBasicTypesSelectBuilder) Paginate(page, perPage uint64) *jBasicTypesSelectBuilder {
	c.SelectBuilder.Paginate(page, perPage)
	return c
}

//Join returns a jBasicTypesSelectBuilder instead of builder.SelectBuilder.
func (c *jBasicTypesSelectBuilder) Join(table, on interface{}) *jBasicTypesSelectBuilder {
	c.SelectBuilder.Join(table, on)
	return c
}

//LeftJoin returns a jBasicTypesSelectBuilder instead of builder.SelectBuilder.
func (c *jBasicTypesSelectBuilder) LeftJoin(table, on interface{}) *jBasicTypesSelectBuilder {
	c.SelectBuilder.LeftJoin(table, on)
	return c
}

//RightJoin returns a jBasicTypesSelectBuilder instead of builder.SelectBuilder.
func (c *jBasicTypesSelectBuilder) RightJoin(table, on interface{}) *jBasicTypesSelectBuilder {
	c.SelectBuilder.RightJoin(table, on)
	return c
}

//FullJoin returns a jBasicTypesSelectBuilder instead of builder.SelectBuilder.
func (c *jBasicTypesSelectBuilder) FullJoin(table, on interface{}) *jBasicTypesSelectBuilder {
	c.SelectBuilder.FullJoin(table, on)
	return c
}

//Distinct returns a jBasicTypesSelectBuilder instead of builder.SelectBuilder.
func (c *jBasicTypesSelectBuilder) Distinct() *jBasicTypesSelectBuilder {
	c.SelectBuilder.Distinct()
	return c
}

// JBasicTypes provides a basic querier
func JBasicTypes(db dbr.SessionRunner) jBasicTypesQuerier {
	return jBasicTypesQuerier{
		db: db,
	}
}

type jBasicTypesQuerier struct {
	db dbr.SessionRunner
	as string
}

//As set alias prior building.
func (c jBasicTypesQuerier) As(as string) jBasicTypesQuerier {
	c.as = as
	return c
}

//Model returns a model with appropriate aliasing.
func (c jBasicTypesQuerier) Model() jBasicTypesModel {
	return JBasicTypesModel.As(c.as)
}

//Select returns a BasicTypes Select Builder.
func (c jBasicTypesQuerier) Select(what ...string) *jBasicTypesSelectBuilder {
	m := c.Model()
	dialect := runtime.GetDialect()
	from := dialect.QuoteIdent(m.Table())
	if m.Alias() != "" && m.Alias() != m.Table() {
		from = fmt.Sprintf("%v as %v", from, dialect.QuoteIdent(m.Alias()))
	}
	if len(what) == 0 {
		alias := m.Table()
		if m.Alias() != "" && m.Alias() != m.Table() {
			alias = m.Alias()
		}
		what = m.Fields(alias + ".*")
	}
	return &jBasicTypesSelectBuilder{
		as: c.as,
		SelectBuilder: &builder.SelectBuilder{
			SelectBuilder: c.db.Select(what...).From(from),
		},
	}
}

//Where returns a BasicTypes Select Builder.
func (c jBasicTypesQuerier) Where(query interface{}, value ...interface{}) *jBasicTypesSelectBuilder {
	return c.Select().Where(query, value...)
}

//Count returns a BasicTypes Select Builder to count given expressions.
func (c jBasicTypesQuerier) Count(what ...string) *jBasicTypesSelectBuilder {
	if len(what) == 0 {
		what = append(what, "*")
	}
	return c.Select("COUNT(" + strings.Join(what, ",") + ")")
}

// Insert a new BasicTypes, if it has autoincrement primary key, the value will be set.
// It stops on first error.
func (c jBasicTypesQuerier) Insert(items ...*BasicTypes) (sql.Result, error) {
	var res sql.Result
	var err error
	for _, data := range items {

		query := c.db.InsertInto(JBasicTypesModel.Table()).Columns(

			`string`,

			`string_p`,

			`intfield`,

			`int_p`,

			`int32`,

			`int32_p`,

			`int64`,

			`int64_p`,

			`u_int`,

			`u_int_p`,

			`u_int32`,

			`u_int32_p`,

			`u_int64`,

			`u_int64_p`,

			`bool`,

			`bool_p`,

			`float32`,

			`float32_p`,

			`float64`,

			`float64_p`,
		).Record(data)
		if runtime.Runs(drivers.Pgsql) {

			query = query.Returning(

				`id`,
			)

			var auto0 int64

			err = query.Load(

				&auto0,
			)

			data.ID = auto0

		} else {
			res, err = query.Exec()

			if err == nil {
				id, err2 := res.LastInsertId()
				if err2 != nil {
					return res, err2
				}
				data.ID = id
			}

		}
		if err != nil {
			return res, err
		}
	}
	return res, err
}

// InsertBulk inserts multiple items into the database.
// It does not post update any auto increment field.
// It builds an insert query of multiple rows and send it on the underlying connection.
func (c jBasicTypesQuerier) InsertBulk(items ...*BasicTypes) error {
	panic("todo")
}

// Update a BasicTypes. It stops on first error.
func (c jBasicTypesQuerier) Update(items ...*BasicTypes) (sql.Result, error) {
	var res sql.Result
	var err error
	for _, data := range items {

		query := c.db.Update(JBasicTypesModel.Table())

		query = query.Set(`string`, data.String)

		query = query.Set(`string_p`, data.StringP)

		query = query.Set(`intfield`, data.Int)

		query = query.Set(`int_p`, data.IntP)

		query = query.Set(`int32`, data.Int32)

		query = query.Set(`int32_p`, data.Int32P)

		query = query.Set(`int64`, data.Int64)

		query = query.Set(`int64_p`, data.Int64P)

		query = query.Set(`u_int`, data.UInt)

		query = query.Set(`u_int_p`, data.UIntP)

		query = query.Set(`u_int32`, data.UInt32)

		query = query.Set(`u_int32_p`, data.UInt32P)

		query = query.Set(`u_int64`, data.UInt64)

		query = query.Set(`u_int64_p`, data.UInt64P)

		query = query.Set(`bool`, data.Bool)

		query = query.Set(`bool_p`, data.BoolP)

		query = query.Set(`float32`, data.Float32)

		query = query.Set(`float32_p`, data.Float32P)

		query = query.Set(`float64`, data.Float64)

		query = query.Set(`float64_p`, data.Float64P)

		query = query.Where("id = ?", data.ID)

		res, err = query.Exec()

		if err != nil {
			return res, err
		}
	}
	return res, err
}

// MustUpdate a BasicTypes. It stops on first error. It errors if an update query does not affect row.
func (c jBasicTypesQuerier) MustUpdate(items ...*BasicTypes) (sql.Result, error) {
	var res sql.Result
	var err error
	for _, data := range items {

		query := c.db.Update(JBasicTypesModel.Table())

		query = query.Set(`string`, data.String)

		query = query.Set(`string_p`, data.StringP)

		query = query.Set(`intfield`, data.Int)

		query = query.Set(`int_p`, data.IntP)

		query = query.Set(`int32`, data.Int32)

		query = query.Set(`int32_p`, data.Int32P)

		query = query.Set(`int64`, data.Int64)

		query = query.Set(`int64_p`, data.Int64P)

		query = query.Set(`u_int`, data.UInt)

		query = query.Set(`u_int_p`, data.UIntP)

		query = query.Set(`u_int32`, data.UInt32)

		query = query.Set(`u_int32_p`, data.UInt32P)

		query = query.Set(`u_int64`, data.UInt64)

		query = query.Set(`u_int64_p`, data.UInt64P)

		query = query.Set(`bool`, data.Bool)

		query = query.Set(`bool_p`, data.BoolP)

		query = query.Set(`float32`, data.Float32)

		query = query.Set(`float32_p`, data.Float32P)

		query = query.Set(`float64`, data.Float64)

		query = query.Set(`float64_p`, data.Float64P)

		query = query.Where("id = ?", data.ID)

		res, err = query.Exec()

		if err == nil {
			if n, _ := res.RowsAffected(); n == 0 {
				x := &builder.UpdateBuilder{UpdateBuilder: query}
				err = runtime.NewNoRowsAffected(x.String())
			}
		}

		if err != nil {
			return res, err
		}
	}
	return res, err
}

// UpdateBulk updates multiple items into the database.
// It builds an update query of multiple rows and send it on the underlying connection.
func (c jBasicTypesQuerier) UpdateBulk(items ...*BasicTypes) error {
	panic("todo")
}

//Delete returns a delete builder
func (c jBasicTypesQuerier) Delete() *jBasicTypesDeleteBuilder {
	return &jBasicTypesDeleteBuilder{
		&builder.DeleteBuilder{
			DeleteBuilder: c.db.DeleteFrom(JBasicTypesModel.Table()),
		},
	}
}

// MustDelete requires the query to affeect rows.
func (c jBasicTypesQuerier) MustDelete() *jBasicTypesDeleteBuilder {
	ret := &jBasicTypesDeleteBuilder{
		&builder.DeleteBuilder{
			DeleteBuilder: c.db.DeleteFrom(JBasicTypesModel.Table()),
		},
	}
	ret.MustDelete()
	return ret
}

//DeleteByPk deletes one BasicTypes by its PKs
func (c jBasicTypesQuerier) DeleteByPk(ID int64) error {
	_, err := c.Delete().Where(

		JBasicTypesModel.ID.Eq(ID),
	).Exec()
	return err
}

// DeleteAll given BasicTypes
func (c jBasicTypesQuerier) DeleteAll(items ...*BasicTypes) (sql.Result, error) {
	q := c.Delete().Where(
		JBasicTypesModel.In(items...),
	)
	return q.Exec()
}

// MustDeleteAll given BasicTypes
func (c jBasicTypesQuerier) MustDeleteAll(items ...*BasicTypes) (sql.Result, error) {
	var res sql.Result
	var err error
	for _, d := range items {
		res, err = c.DeleteAll(d)
		if err != nil {
			return res, err
		}
		if n, e := res.RowsAffected(); e != nil {
			return res, e
		} else if n == 0 {
			q := c.Delete().Where(
				JBasicTypesModel.In(items...),
			)
			err = runtime.NewNoRowsAffected(q.String())
			return res, err
		}
	}
	return res, err
}

//Find one BasicTypes using its PKs
func (c jBasicTypesQuerier) Find(ID int64) (*BasicTypes, error) {
	return c.Select().Where(

		JBasicTypesModel.ID.Eq(ID),
	).Read()
}

// JTextPkSetup helps to create/drop the schema
func JTextPkSetup() runtime.Setuper {
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
name VARCHAR(255) NOT NULL,
description TEXT,
PRIMARY KEY (name) 

)`
	} else if driver == drivers.Pgsql {
		create = `CREATE TABLE IF NOT EXISTS second_sample (
name TEXT,
description TEXT,
PRIMARY KEY (name) 

)`
	}

	if driver == drivers.Sqlite {
		drop = `DROP TABLE IF EXISTS second_sample`
	} else if driver == drivers.Mysql {
		drop = `DROP TABLE IF EXISTS second_sample`
	} else if driver == drivers.Pgsql {
		drop = `DROP TABLE IF EXISTS second_sample`
	}

	var indexes []string

	if driver == drivers.Sqlite {

	} else if driver == drivers.Mysql {

	} else if driver == drivers.Pgsql {

	}

	return runtime.Table{
		Name:       `second_sample`,
		CreateStmt: create,
		DropStmt:   drop,
		View:       !true,
		Indexes:    indexes,
	}
}

// jTextPkModel provides helper to work with TextPk data provider
type jTextPkModel struct {
	as string

	Name builder.ValuePropertyMeta

	Description builder.ValuePropertyMeta

	HasManyHasOneTextPk builder.RelPropertyMeta

	Relateds builder.RelPropertyMeta
}

// Eq provided items.
func (j jTextPkModel) Eq(s ...*TextPk) dbr.Builder {
	ors := []dbr.Builder{}
	for _, t := range s {
		ors = append(ors, dbr.And(

			JTextPkModel.Name.Eq(t.Name),
		))
	}
	return dbr.Or(ors...)
}

// In provided items.
func (j jTextPkModel) In(s ...*TextPk) dbr.Builder {
	ors := []dbr.Builder{}
	for _, t := range s {
		ors = append(ors, dbr.And(

			JTextPkModel.Name.Eq(t.Name),
		))
	}
	return dbr.Or(ors...)
}

// As returns a copy with an alias.
func (j jTextPkModel) As(as string) jTextPkModel {
	j.as = as

	j.Name.TableAlias = as

	j.Description.TableAlias = as

	// j.HasManyHasOneTextPk.TableAlias = as

	// j.Relateds.TableAlias = as

	return j
}

// Table returns the sql table name
func (j jTextPkModel) Table() string {
	return "second_sample"
}

// Alias returns the current alias
func (j jTextPkModel) Alias() string {
	if j.as == "" {
		return j.Table()
	}
	return j.as
}

// Properties returns a map of property name => meta
func (j jTextPkModel) Properties() map[string]builder.MetaProvider {
	ret := map[string]builder.MetaProvider{}

	ret["Name"] = j.Name

	ret["Description"] = j.Description

	ret["HasManyHasOneTextPk"] = j.HasManyHasOneTextPk

	ret["Relateds"] = j.Relateds

	return ret
}

// Fields returns given sql fields with appropriate aliasing.
func (j jTextPkModel) Fields(ins ...string) []string {
	dialect := runtime.GetDialect()
	if len(ins) == 0 {
		ins = append(ins, "*")
	}
	for i, in := range ins {
		if j.as != "" {
			if in == "*" {
				ins[i] = fmt.Sprintf("%v.%v", dialect.QuoteIdent(j.as), in)
			} else {
				ins[i] = fmt.Sprintf("%v.%v", dialect.QuoteIdent(j.as), dialect.QuoteIdent(in))
			}
		}
	}
	return ins
}

// JTextPkModel provides helper to work with TextPk data provider
var JTextPkModel = jTextPkModel{

	Name: builder.NewValueMeta(
		`name`, `TEXT`,
		`Name`, `string`,
		true, false,
	),

	Description: builder.NewValueMeta(
		`description`, `TEXT`,
		`Description`, `string`,
		false, false,
	),

	HasManyHasOneTextPk: builder.NewRelMeta(
		`hasManyHasOneTextPk`, `[]*HasOneTextPk`,
		``, `HasOneTextPk.related`, ``,
		`has_many`,
	),

	Relateds: builder.NewRelMeta(
		`relateds`, `[]*HasManyTextPk`,
		``, `HasManyTextPk.relateds`, ``,
		`has_many2many`,
	),
}

type jTextPkDeleteBuilder struct {
	*builder.DeleteBuilder
}

// //Build builds the sql string into given buffer using current dialect
// func (c *jTextPkDeleteBuilder) Build(b dbr.Buffer) error {
// 	return c.DeleteBuilder.Build(runtime.GetDialect(), b)
// }
// //String returns the sql string for current dialect. It returns empty string if the build returns an error.
// func (c *jTextPkDeleteBuilder) String() string {
// 	b := dbr.NewBuffer()
// 	if err := c.Build(b); err != nil {
// 		return ""
// 	}
// 	return b.String()
// }
//Where returns a jTextPkDeleteBuilder instead of builder.DeleteBuilder.
func (c *jTextPkDeleteBuilder) Where(query interface{}, value ...interface{}) *jTextPkDeleteBuilder {
	c.DeleteBuilder.Where(query, value...)
	return c
}

type jTextPkSelectBuilder struct {
	as string
	*builder.SelectBuilder
}

// //Build builds the sql string using current dialect into given bufer
// func (c *jTextPkSelectBuilder) Build(b dbr.Buffer) error {
// 	return c.SelectBuilder.Build(runtime.GetDialect(), b)
// }
// //String returns the sql string for current dialect. It returns empty string if the build returns an error.
// func (c *jTextPkSelectBuilder) String() string {
// 	b := dbr.NewBuffer()
// 	if err := c.Build(b); err != nil {
// 		return ""
// 	}
// 	return b.String()
// }

//Read evaluates current select query and load the results into a TextPk
func (c *jTextPkSelectBuilder) Read() (*TextPk, error) {
	var one TextPk
	err := c.LoadStruct(&one)
	return &one, err
}

//ReadAll evaluates current select query and load the results into a slice of TextPk
func (c *jTextPkSelectBuilder) ReadAll() ([]*TextPk, error) {
	var all []*TextPk
	_, err := c.LoadStructs(&all)
	return all, err
}

//Where returns a jTextPkSelectBuilder instead of builder.SelectBuilder.
func (c *jTextPkSelectBuilder) Where(query interface{}, value ...interface{}) *jTextPkSelectBuilder {
	c.SelectBuilder.Where(query, value...)
	return c
}

//GroupBy returns a jTextPkSelectBuilder instead of builder.SelectBuilder.
func (c *jTextPkSelectBuilder) GroupBy(col ...string) *jTextPkSelectBuilder {
	c.SelectBuilder.GroupBy(col...)
	return c
}

//Having returns a jTextPkSelectBuilder instead of builder.SelectBuilder.
func (c *jTextPkSelectBuilder) Having(query interface{}, value ...interface{}) *jTextPkSelectBuilder {
	c.SelectBuilder.Having(query, value...)
	return c
}

//Limit returns a jTextPkSelectBuilder instead of builder.SelectBuilder.
func (c *jTextPkSelectBuilder) Limit(n uint64) *jTextPkSelectBuilder {
	c.SelectBuilder.Limit(n)
	return c
}

//Offset returns a jTextPkSelectBuilder instead of builder.SelectBuilder.
func (c *jTextPkSelectBuilder) Offset(n uint64) *jTextPkSelectBuilder {
	c.SelectBuilder.Offset(n)
	return c
}

//OrderAsc returns a jTextPkSelectBuilder instead of builder.SelectBuilder.
func (c *jTextPkSelectBuilder) OrderAsc(col string) *jTextPkSelectBuilder {
	c.SelectBuilder.OrderAsc(col)
	return c
}

//OrderDesc returns a jTextPkSelectBuilder instead of builder.SelectBuilder.
func (c *jTextPkSelectBuilder) OrderDesc(col string) *jTextPkSelectBuilder {
	c.SelectBuilder.OrderDesc(col)
	return c
}

//OrderDir returns a jTextPkSelectBuilder instead of builder.SelectBuilder.
func (c *jTextPkSelectBuilder) OrderDir(col string, isAsc bool) *jTextPkSelectBuilder {
	c.SelectBuilder.OrderDir(col, isAsc)
	return c
}

//OrderBy returns a jTextPkSelectBuilder instead of builder.SelectBuilder.
func (c *jTextPkSelectBuilder) OrderBy(col string) *jTextPkSelectBuilder {
	c.SelectBuilder.OrderBy(col)
	return c
}

//Paginate returns a jTextPkSelectBuilder instead of builder.SelectBuilder.
func (c *jTextPkSelectBuilder) Paginate(page, perPage uint64) *jTextPkSelectBuilder {
	c.SelectBuilder.Paginate(page, perPage)
	return c
}

//Join returns a jTextPkSelectBuilder instead of builder.SelectBuilder.
func (c *jTextPkSelectBuilder) Join(table, on interface{}) *jTextPkSelectBuilder {
	c.SelectBuilder.Join(table, on)
	return c
}

//LeftJoin returns a jTextPkSelectBuilder instead of builder.SelectBuilder.
func (c *jTextPkSelectBuilder) LeftJoin(table, on interface{}) *jTextPkSelectBuilder {
	c.SelectBuilder.LeftJoin(table, on)
	return c
}

//RightJoin returns a jTextPkSelectBuilder instead of builder.SelectBuilder.
func (c *jTextPkSelectBuilder) RightJoin(table, on interface{}) *jTextPkSelectBuilder {
	c.SelectBuilder.RightJoin(table, on)
	return c
}

//FullJoin returns a jTextPkSelectBuilder instead of builder.SelectBuilder.
func (c *jTextPkSelectBuilder) FullJoin(table, on interface{}) *jTextPkSelectBuilder {
	c.SelectBuilder.FullJoin(table, on)
	return c
}

//Distinct returns a jTextPkSelectBuilder instead of builder.SelectBuilder.
func (c *jTextPkSelectBuilder) Distinct() *jTextPkSelectBuilder {
	c.SelectBuilder.Distinct()
	return c
}

// JTextPk provides a basic querier
func JTextPk(db dbr.SessionRunner) jTextPkQuerier {
	return jTextPkQuerier{
		db: db,
	}
}

type jTextPkQuerier struct {
	db dbr.SessionRunner
	as string
}

//As set alias prior building.
func (c jTextPkQuerier) As(as string) jTextPkQuerier {
	c.as = as
	return c
}

//Model returns a model with appropriate aliasing.
func (c jTextPkQuerier) Model() jTextPkModel {
	return JTextPkModel.As(c.as)
}

//Select returns a TextPk Select Builder.
func (c jTextPkQuerier) Select(what ...string) *jTextPkSelectBuilder {
	m := c.Model()
	dialect := runtime.GetDialect()
	from := dialect.QuoteIdent(m.Table())
	if m.Alias() != "" && m.Alias() != m.Table() {
		from = fmt.Sprintf("%v as %v", from, dialect.QuoteIdent(m.Alias()))
	}
	if len(what) == 0 {
		alias := m.Table()
		if m.Alias() != "" && m.Alias() != m.Table() {
			alias = m.Alias()
		}
		what = m.Fields(alias + ".*")
	}
	return &jTextPkSelectBuilder{
		as: c.as,
		SelectBuilder: &builder.SelectBuilder{
			SelectBuilder: c.db.Select(what...).From(from),
		},
	}
}

//Where returns a TextPk Select Builder.
func (c jTextPkQuerier) Where(query interface{}, value ...interface{}) *jTextPkSelectBuilder {
	return c.Select().Where(query, value...)
}

//Count returns a TextPk Select Builder to count given expressions.
func (c jTextPkQuerier) Count(what ...string) *jTextPkSelectBuilder {
	if len(what) == 0 {
		what = append(what, "*")
	}
	return c.Select("COUNT(" + strings.Join(what, ",") + ")")
}

// Insert a new TextPk, if it has autoincrement primary key, the value will be set.
// It stops on first error.
func (c jTextPkQuerier) Insert(items ...*TextPk) (sql.Result, error) {
	var res sql.Result
	var err error
	for _, data := range items {

		if runtime.Runs(drivers.Mysql) {

			if len(data.Name) > 255 {
				return nil, fmt.Errorf("Name: PRIMARY KEY length exceeded max=255, got=%v", len(data.Name))
			}

		}

		query := c.db.InsertInto(JTextPkModel.Table()).Columns(

			`name`,

			`description`,
		).Record(data)
		if runtime.Runs(drivers.Pgsql) {

			res, err = query.Exec()

		} else {
			res, err = query.Exec()

		}
		if err != nil {
			return res, err
		}
	}
	return res, err
}

// InsertBulk inserts multiple items into the database.
// It does not post update any auto increment field.
// It builds an insert query of multiple rows and send it on the underlying connection.
func (c jTextPkQuerier) InsertBulk(items ...*TextPk) error {
	panic("todo")
}

// Update a TextPk. It stops on first error.
func (c jTextPkQuerier) Update(items ...*TextPk) (sql.Result, error) {
	var res sql.Result
	var err error
	for _, data := range items {

		query := c.db.Update(JTextPkModel.Table())

		query = query.Set(`description`, data.Description)

		query = query.Where("name = ?", data.Name)

		res, err = query.Exec()

		if err != nil {
			return res, err
		}
	}
	return res, err
}

// MustUpdate a TextPk. It stops on first error. It errors if an update query does not affect row.
func (c jTextPkQuerier) MustUpdate(items ...*TextPk) (sql.Result, error) {
	var res sql.Result
	var err error
	for _, data := range items {

		query := c.db.Update(JTextPkModel.Table())

		query = query.Set(`description`, data.Description)

		query = query.Where("name = ?", data.Name)

		res, err = query.Exec()

		if err == nil {
			if n, _ := res.RowsAffected(); n == 0 {
				x := &builder.UpdateBuilder{UpdateBuilder: query}
				err = runtime.NewNoRowsAffected(x.String())
			}
		}

		if err != nil {
			return res, err
		}
	}
	return res, err
}

// UpdateBulk updates multiple items into the database.
// It builds an update query of multiple rows and send it on the underlying connection.
func (c jTextPkQuerier) UpdateBulk(items ...*TextPk) error {
	panic("todo")
}

//Delete returns a delete builder
func (c jTextPkQuerier) Delete() *jTextPkDeleteBuilder {
	return &jTextPkDeleteBuilder{
		&builder.DeleteBuilder{
			DeleteBuilder: c.db.DeleteFrom(JTextPkModel.Table()),
		},
	}
}

// MustDelete requires the query to affeect rows.
func (c jTextPkQuerier) MustDelete() *jTextPkDeleteBuilder {
	ret := &jTextPkDeleteBuilder{
		&builder.DeleteBuilder{
			DeleteBuilder: c.db.DeleteFrom(JTextPkModel.Table()),
		},
	}
	ret.MustDelete()
	return ret
}

//DeleteByPk deletes one TextPk by its PKs
func (c jTextPkQuerier) DeleteByPk(Name string) error {
	_, err := c.Delete().Where(

		JTextPkModel.Name.Eq(Name),
	).Exec()
	return err
}

// DeleteAll given TextPk
func (c jTextPkQuerier) DeleteAll(items ...*TextPk) (sql.Result, error) {
	q := c.Delete().Where(
		JTextPkModel.In(items...),
	)
	return q.Exec()
}

// MustDeleteAll given TextPk
func (c jTextPkQuerier) MustDeleteAll(items ...*TextPk) (sql.Result, error) {
	var res sql.Result
	var err error
	for _, d := range items {
		res, err = c.DeleteAll(d)
		if err != nil {
			return res, err
		}
		if n, e := res.RowsAffected(); e != nil {
			return res, e
		} else if n == 0 {
			q := c.Delete().Where(
				JTextPkModel.In(items...),
			)
			err = runtime.NewNoRowsAffected(q.String())
			return res, err
		}
	}
	return res, err
}

//Find one TextPk using its PKs
func (c jTextPkQuerier) Find(Name string) (*TextPk, error) {
	return c.Select().Where(

		JTextPkModel.Name.Eq(Name),
	).Read()
}

// Relateds returns a query builder to select Relateds linked to this TextPk
func (g *TextPk) Relateds(db dbr.SessionRunner,
	AsHasManyTextPk, AsHasManyTextPkrelatedsToTextPkrelateds, AsTextPk string,
) *jHasManyTextPkSelectBuilder {

	leftTable := JHasManyTextPkModel.Table()
	var query *jHasManyTextPkSelectBuilder
	if AsHasManyTextPk != "" {
		leftTable = AsHasManyTextPk
		query = JHasManyTextPk(db).As(AsHasManyTextPk).Select(AsHasManyTextPk + ".*")
	} else {
		query = JHasManyTextPk(db).Select(leftTable + ".*")
	}

	midTable := JHasManyTextPkrelatedsToTextPkrelatedsModel.Table()
	{
		on := ""
		if AsHasManyTextPkrelatedsToTextPkrelateds != "" {
			midTable = AsHasManyTextPkrelatedsToTextPkrelateds
		}

		on += fmt.Sprintf("%v.%v = %v.%v",
			midTable, "has_many_text_pk_id",
			leftTable, "id",
		)

		if AsHasManyTextPkrelatedsToTextPkrelateds == "" {
			query = query.Join(dbr.I(JHasManyTextPkrelatedsToTextPkrelatedsModel.Table()), on)
		} else {
			query = query.Join(dbr.I(JHasManyTextPkrelatedsToTextPkrelatedsModel.Table()).As(AsHasManyTextPkrelatedsToTextPkrelateds), on)
		}
	}

	rightTable := JTextPkModel.Table()
	{
		on := ""
		if AsTextPk != "" {
			rightTable = AsTextPk
		}

		on += fmt.Sprintf("%v.%v = %v.%v",
			midTable, "text_pk_name",
			rightTable, "name",
		)

		if AsTextPk == "" {
			query = query.Join(dbr.I(JTextPkModel.Table()), on)
		} else {
			query = query.Join(dbr.I(JTextPkModel.Table()).As(AsTextPk), on)
		}
	}

	{
		m := JTextPkModel
		if AsTextPk != "" {
			m = m.As(AsTextPk)
		} else {
			m = m.As(m.Table())
		}
		query = query.Where(

			m.Name.Eq(g.Name),
		)
	}

	return query
}

//LinkWithRelateds writes new links with TextPk.
func (g *TextPk) LinkWithRelateds(db dbr.SessionRunner, items ...*HasManyTextPk) (sql.Result, error) {
	toInsert := []*HasManyTextPkrelatedsToTextPkrelateds{}
	for _, item := range items {
		toInsert = append(toInsert, &HasManyTextPkrelatedsToTextPkrelateds{

			HasManyTextPkID: item.ID,

			TextPkName: g.Name,
		})
	}
	return JHasManyTextPkrelatedsToTextPkrelateds(db).Insert(toInsert...)
}

//UnlinkWithRelateds deletes given existing links with TextPk.
func (g *TextPk) UnlinkWithRelateds(db dbr.SessionRunner, items ...*HasManyTextPk) (sql.Result, error) {
	toDelete := []*HasManyTextPkrelatedsToTextPkrelateds{}
	for _, item := range items {
		toDelete = append(toDelete, &HasManyTextPkrelatedsToTextPkrelateds{

			HasManyTextPkID: item.ID,

			TextPkName: g.Name,
		})
	}
	return JHasManyTextPkrelatedsToTextPkrelateds(db).DeleteAll(toDelete...)
}

//UnlinkAllRelateds deletes all existing links with TextPk.
func (g *TextPk) UnlinkAllRelateds(db dbr.SessionRunner) (sql.Result, error) {
	return JHasManyTextPkrelatedsToTextPkrelateds(db).Delete().Where(

		JHasManyTextPkrelatedsToTextPkrelatedsModel.TextPkName.Eq(g.Name),
	).Exec()
}

//SetRelateds replaces existing links with TextPk.
func (g *TextPk) SetRelateds(db dbr.SessionRunner, items ...*HasManyTextPk) (sql.Result, error) {
	if res, err := g.UnlinkAllRelateds(db); err != nil {
		return res, err
	}
	return g.LinkWithRelateds(db, items...)
}

// JoinRelateds adds a JOIN to TextPk.Relateds
func (c *jTextPkSelectBuilder) JoinRelateds(
	AsHasManyTextPkrelatedsToTextPkrelateds, AsHasManyTextPk string,
) *jTextPkSelectBuilder {

	query := c

	leftTable := JTextPkModel.Table()
	if c.as != "" {
		leftTable = c.as
	}

	midTable := JHasManyTextPkrelatedsToTextPkrelatedsModel.Table()
	if AsHasManyTextPkrelatedsToTextPkrelateds != "" {
		midTable = AsHasManyTextPkrelatedsToTextPkrelateds
	}

	{
		on := ""

		on += fmt.Sprintf("%v.%v = %v.%v",
			midTable, "text_pk_name",
			leftTable, "name",
		)

		if AsHasManyTextPkrelatedsToTextPkrelateds == "" {
			query = query.Join(dbr.I(JHasManyTextPkrelatedsToTextPkrelatedsModel.Table()), on)
		} else {
			query = query.Join(dbr.I(JHasManyTextPkrelatedsToTextPkrelatedsModel.Table()).As(AsHasManyTextPkrelatedsToTextPkrelateds), on)
		}
	}

	{
		rightTable := JHasManyTextPkModel.Table()
		if AsHasManyTextPk != "" {
			rightTable = AsHasManyTextPk
		}
		on := ""

		on += fmt.Sprintf("%v.%v = %v.%v",
			midTable, "has_many_text_pk_id",
			rightTable, "id",
		)

		if AsHasManyTextPk == "" {
			query = query.Join(dbr.I(JHasManyTextPkModel.Table()), on)
		} else {
			query = query.Join(dbr.I(JHasManyTextPkModel.Table()).As(AsHasManyTextPk), on)
		}
	}

	return query
}

// LeftJoinRelateds adds a LEFT JOIN to TextPk.Relateds
func (c *jTextPkSelectBuilder) LeftJoinRelateds(
	AsHasManyTextPkrelatedsToTextPkrelateds, AsHasManyTextPk string,
) *jTextPkSelectBuilder {

	query := c

	leftTable := JTextPkModel.Table()
	if c.as != "" {
		leftTable = c.as
	}

	midTable := JHasManyTextPkrelatedsToTextPkrelatedsModel.Table()
	if AsHasManyTextPkrelatedsToTextPkrelateds != "" {
		midTable = AsHasManyTextPkrelatedsToTextPkrelateds
	}

	{
		on := ""

		on += fmt.Sprintf("%v.%v = %v.%v",
			midTable, "text_pk_name",
			leftTable, "name",
		)

		if AsHasManyTextPkrelatedsToTextPkrelateds == "" {
			query = query.LeftJoin(dbr.I(JHasManyTextPkrelatedsToTextPkrelatedsModel.Table()), on)
		} else {
			query = query.LeftJoin(dbr.I(JHasManyTextPkrelatedsToTextPkrelatedsModel.Table()).As(AsHasManyTextPkrelatedsToTextPkrelateds), on)
		}
	}

	{
		rightTable := JHasManyTextPkModel.Table()
		if AsHasManyTextPk != "" {
			rightTable = AsHasManyTextPk
		}
		on := ""

		on += fmt.Sprintf("%v.%v = %v.%v",
			midTable, "has_many_text_pk_id",
			rightTable, "id",
		)

		if AsHasManyTextPk == "" {
			query = query.LeftJoin(dbr.I(JHasManyTextPkModel.Table()), on)
		} else {
			query = query.LeftJoin(dbr.I(JHasManyTextPkModel.Table()).As(AsHasManyTextPk), on)
		}
	}

	return query
}

// // RightJoinRelateds adds a RIGHT JOIN to TextPk.Relateds
// func (c *jTextPkSelectBuilder) RightJoinRelateds(
// 	AsHasManyTextPkrelatedsToTextPkrelateds, AsHasManyTextPk string,
// ) *jTextPkSelectBuilder {
//
// 	query := c
//
// 	leftTable := JTextPkModel.Table()
// 	if c.as != "" {
// 		leftTable = c.as
// 	}
//
// 	midTable := JHasManyTextPkrelatedsToTextPkrelatedsModel.Table()
// 	if AsHasManyTextPkrelatedsToTextPkrelateds != "" {
// 		midTable = AsHasManyTextPkrelatedsToTextPkrelateds
// 	}
//
// 	{
// 		on := ""
//
// 		on += fmt.Sprintf("%v.%v = %v.%v",
// 			midTable, "text_pk_name",
// 			leftTable, "name",
// 			)
//
//
// 		if AsHasManyTextPkrelatedsToTextPkrelateds == "" {
// 			query = query.RightJoin(dbr.I(JHasManyTextPkrelatedsToTextPkrelatedsModel.Table()), on)
// 		} else {
// 			query = query.RightJoin(dbr.I(JHasManyTextPkrelatedsToTextPkrelatedsModel.Table()).As(AsHasManyTextPkrelatedsToTextPkrelateds), on)
// 		}
// 	}
//
// 	{
// 		rightTable := JHasManyTextPkModel.Table()
// 		if AsHasManyTextPk != "" {
// 			rightTable = AsHasManyTextPk
// 		}
// 		on := ""
//
// 		on += fmt.Sprintf("%v.%v = %v.%v",
// 			midTable, "has_many_text_pk_id",
// 			rightTable, "id",
// 			)
//
//
// 		if AsHasManyTextPk == "" {
// 			query = query.RightJoin(dbr.I(JHasManyTextPkModel.Table()), on)
// 		} else {
// 			query = query.RightJoin(dbr.I(JHasManyTextPkModel.Table()).As(AsHasManyTextPk), on)
// 		}
// 	}
//
// 	return query
// }

// HasManyHasOneTextPk returns a query builder to select HasManyHasOneTextPk linked to this TextPk
func (g *TextPk) HasManyHasOneTextPk(db dbr.SessionRunner,
	AsRelated, AsHasManyHasOneTextPk string,
) *jHasOneTextPkSelectBuilder {

	var query *jHasOneTextPkSelectBuilder

	leftTable := JHasOneTextPkModel.Table()
	if AsRelated != "" {
		leftTable = AsRelated
		query = JHasOneTextPk(db).As(AsRelated).Select(leftTable + ".*")
	} else {
		query = JHasOneTextPk(db).Select(leftTable + ".*")
	}

	rightTable := JTextPkModel.Table()
	if AsHasManyHasOneTextPk != "" {
		rightTable = AsHasManyHasOneTextPk
	}

	on := ""

	on += fmt.Sprintf("%v.%v = %v.%v",
		leftTable, "related_name",
		rightTable, "name",
	)

	if AsHasManyHasOneTextPk == "" {
		return query.Join(dbr.I(JTextPkModel.Table()), on)
	}
	return query.Join(dbr.I(JTextPkModel.Table()).As(AsHasManyHasOneTextPk), on)
}

// JoinHasManyHasOneTextPk adds a JOIN to TextPk.HasManyHasOneTextPk
func (c *jTextPkSelectBuilder) JoinHasManyHasOneTextPk(
	AsHasManyHasOneTextPk string,
) *jTextPkSelectBuilder {
	dialect := runtime.GetDialect()
	on := ""
	localTable := dialect.QuoteIdent(JTextPkModel.Table())
	if c.as != "" {
		localTable = dialect.QuoteIdent(c.as)
	}
	foreiTable := dialect.QuoteIdent(JHasOneTextPkModel.Table())
	if AsHasManyHasOneTextPk != "" {
		foreiTable = dialect.QuoteIdent(AsHasManyHasOneTextPk)
	}

	on += fmt.Sprintf("%v.%v = %v.%v",
		localTable, dialect.QuoteIdent("name"),
		foreiTable, dialect.QuoteIdent("related_name"),
	)

	if AsHasManyHasOneTextPk == "" {
		return c.Join(dbr.I(JHasOneTextPkModel.Table()), on)
	}
	return c.Join(dbr.I(JHasOneTextPkModel.Table()).As(AsHasManyHasOneTextPk), on)
}

// LeftJoinHasManyHasOneTextPk adds a LEFT JOIN to TextPk.HasManyHasOneTextPk
func (c *jTextPkSelectBuilder) LeftJoinHasManyHasOneTextPk(
	AsHasManyHasOneTextPk string,
) *jTextPkSelectBuilder {
	dialect := runtime.GetDialect()
	on := ""
	localTable := dialect.QuoteIdent(JTextPkModel.Table())
	if c.as != "" {
		localTable = dialect.QuoteIdent(c.as)
	}
	foreiTable := dialect.QuoteIdent(JHasOneTextPkModel.Table())
	if AsHasManyHasOneTextPk != "" {
		foreiTable = dialect.QuoteIdent(AsHasManyHasOneTextPk)
	}

	on += fmt.Sprintf("%v.%v = %v.%v",
		localTable, dialect.QuoteIdent("name"),
		foreiTable, dialect.QuoteIdent("related_name"),
	)

	if AsHasManyHasOneTextPk == "" {
		return c.LeftJoin(dbr.I(JHasOneTextPkModel.Table()), on)
	}
	return c.LeftJoin(dbr.I(JHasOneTextPkModel.Table()).As(AsHasManyHasOneTextPk), on)
}

// // RightJoinHasManyHasOneTextPk adds a Right JOIN to TextPk.HasManyHasOneTextPk
// func (c *jTextPkSelectBuilder) RightJoinHasManyHasOneTextPk(
// 	AsHasManyHasOneTextPk string,
// ) *jTextPkSelectBuilder {
// 	dialect := runtime.GetDialect()
// 	on := ""
// 	localTable := dialect.QuoteIdent(JTextPkModel.Table())
// 	if c.as != "" {
// 		localTable = dialect.QuoteIdent(c.as)
// 	}
// 	foreiTable := dialect.QuoteIdent(JHasOneTextPkModel.Table())
// 	if AsHasManyHasOneTextPk != "" {
// 		foreiTable = dialect.QuoteIdent(AsHasManyHasOneTextPk)
// 	}
//
// 	on += fmt.Sprintf("%v.%v = %v.%v",
// 		localTable, dialect.QuoteIdent("name"),
// 		foreiTable, dialect.QuoteIdent("related_name"),
// 	)
//
// 	if AsHasManyHasOneTextPk == "" {
// 		return c.RightJoin(dbr.I(JHasOneTextPkModel.Table()), on)
// 	}
// 	return c.RightJoin(dbr.I(JHasOneTextPkModel.Table()).As(AsHasManyHasOneTextPk), on)
// }

// JHasOneTextPkSetup helps to create/drop the schema
func JHasOneTextPkSetup() runtime.Setuper {
	driver := runtime.GetCurrentDriver()

	var create string
	var drop string

	if driver == drivers.Sqlite {
		create = `CREATE TABLE IF NOT EXISTS has_one_text_pk (
id INTEGER PRIMARY KEY AUTOINCREMENT,
x TEXT,
related_name TEXT NULL

)`
	} else if driver == drivers.Mysql {
		create = `CREATE TABLE IF NOT EXISTS has_one_text_pk (
id INTEGER NOT NULL AUTO_INCREMENT,
x TEXT,
related_name TEXT NULL,
PRIMARY KEY (id) 

)`
	} else if driver == drivers.Pgsql {
		create = `CREATE TABLE IF NOT EXISTS has_one_text_pk (
id SERIAL PRIMARY KEY,
x TEXT,
related_name TEXT NULL

)`
	}

	if driver == drivers.Sqlite {
		drop = `DROP TABLE IF EXISTS has_one_text_pk`
	} else if driver == drivers.Mysql {
		drop = `DROP TABLE IF EXISTS has_one_text_pk`
	} else if driver == drivers.Pgsql {
		drop = `DROP TABLE IF EXISTS has_one_text_pk`
	}

	var indexes []string

	if driver == drivers.Sqlite {

	} else if driver == drivers.Mysql {

	} else if driver == drivers.Pgsql {

	}

	return runtime.Table{
		Name:       `has_one_text_pk`,
		CreateStmt: create,
		DropStmt:   drop,
		View:       !true,
		Indexes:    indexes,
	}
}

// jHasOneTextPkModel provides helper to work with HasOneTextPk data provider
type jHasOneTextPkModel struct {
	as string

	ID builder.ValuePropertyMeta

	X builder.ValuePropertyMeta

	RelatedName builder.ValuePropertyMeta

	Related builder.RelPropertyMeta
}

// Eq provided items.
func (j jHasOneTextPkModel) Eq(s ...*HasOneTextPk) dbr.Builder {
	ors := []dbr.Builder{}
	for _, t := range s {
		ors = append(ors, dbr.And(

			JHasOneTextPkModel.ID.Eq(t.ID),
		))
	}
	return dbr.Or(ors...)
}

// In provided items.
func (j jHasOneTextPkModel) In(s ...*HasOneTextPk) dbr.Builder {
	ors := []dbr.Builder{}
	for _, t := range s {
		ors = append(ors, dbr.And(

			JHasOneTextPkModel.ID.Eq(t.ID),
		))
	}
	return dbr.Or(ors...)
}

// As returns a copy with an alias.
func (j jHasOneTextPkModel) As(as string) jHasOneTextPkModel {
	j.as = as

	j.ID.TableAlias = as

	j.X.TableAlias = as

	j.RelatedName.TableAlias = as

	// j.Related.TableAlias = as

	return j
}

// Table returns the sql table name
func (j jHasOneTextPkModel) Table() string {
	return "has_one_text_pk"
}

// Alias returns the current alias
func (j jHasOneTextPkModel) Alias() string {
	if j.as == "" {
		return j.Table()
	}
	return j.as
}

// Properties returns a map of property name => meta
func (j jHasOneTextPkModel) Properties() map[string]builder.MetaProvider {
	ret := map[string]builder.MetaProvider{}

	ret["ID"] = j.ID

	ret["X"] = j.X

	ret["RelatedName"] = j.RelatedName

	ret["Related"] = j.Related

	return ret
}

// Fields returns given sql fields with appropriate aliasing.
func (j jHasOneTextPkModel) Fields(ins ...string) []string {
	dialect := runtime.GetDialect()
	if len(ins) == 0 {
		ins = append(ins, "*")
	}
	for i, in := range ins {
		if j.as != "" {
			if in == "*" {
				ins[i] = fmt.Sprintf("%v.%v", dialect.QuoteIdent(j.as), in)
			} else {
				ins[i] = fmt.Sprintf("%v.%v", dialect.QuoteIdent(j.as), dialect.QuoteIdent(in))
			}
		}
	}
	return ins
}

// JHasOneTextPkModel provides helper to work with HasOneTextPk data provider
var JHasOneTextPkModel = jHasOneTextPkModel{

	ID: builder.NewValueMeta(
		`id`, `INTEGER`,
		`ID`, `int64`,
		true, true,
	),

	X: builder.NewValueMeta(
		`x`, `TEXT`,
		`X`, `string`,
		false, false,
	),

	RelatedName: builder.NewValueMeta(
		`related_name`, `TEXT`,
		`RelatedName`, `*string`,
		false, false,
	),

	Related: builder.NewRelMeta(
		`related`, `*TextPk`,
		`TextPk`, ``, ``,
		`has_one`,
	),
}

type jHasOneTextPkDeleteBuilder struct {
	*builder.DeleteBuilder
}

// //Build builds the sql string into given buffer using current dialect
// func (c *jHasOneTextPkDeleteBuilder) Build(b dbr.Buffer) error {
// 	return c.DeleteBuilder.Build(runtime.GetDialect(), b)
// }
// //String returns the sql string for current dialect. It returns empty string if the build returns an error.
// func (c *jHasOneTextPkDeleteBuilder) String() string {
// 	b := dbr.NewBuffer()
// 	if err := c.Build(b); err != nil {
// 		return ""
// 	}
// 	return b.String()
// }
//Where returns a jHasOneTextPkDeleteBuilder instead of builder.DeleteBuilder.
func (c *jHasOneTextPkDeleteBuilder) Where(query interface{}, value ...interface{}) *jHasOneTextPkDeleteBuilder {
	c.DeleteBuilder.Where(query, value...)
	return c
}

type jHasOneTextPkSelectBuilder struct {
	as string
	*builder.SelectBuilder
}

// //Build builds the sql string using current dialect into given bufer
// func (c *jHasOneTextPkSelectBuilder) Build(b dbr.Buffer) error {
// 	return c.SelectBuilder.Build(runtime.GetDialect(), b)
// }
// //String returns the sql string for current dialect. It returns empty string if the build returns an error.
// func (c *jHasOneTextPkSelectBuilder) String() string {
// 	b := dbr.NewBuffer()
// 	if err := c.Build(b); err != nil {
// 		return ""
// 	}
// 	return b.String()
// }

//Read evaluates current select query and load the results into a HasOneTextPk
func (c *jHasOneTextPkSelectBuilder) Read() (*HasOneTextPk, error) {
	var one HasOneTextPk
	err := c.LoadStruct(&one)
	return &one, err
}

//ReadAll evaluates current select query and load the results into a slice of HasOneTextPk
func (c *jHasOneTextPkSelectBuilder) ReadAll() ([]*HasOneTextPk, error) {
	var all []*HasOneTextPk
	_, err := c.LoadStructs(&all)
	return all, err
}

//Where returns a jHasOneTextPkSelectBuilder instead of builder.SelectBuilder.
func (c *jHasOneTextPkSelectBuilder) Where(query interface{}, value ...interface{}) *jHasOneTextPkSelectBuilder {
	c.SelectBuilder.Where(query, value...)
	return c
}

//GroupBy returns a jHasOneTextPkSelectBuilder instead of builder.SelectBuilder.
func (c *jHasOneTextPkSelectBuilder) GroupBy(col ...string) *jHasOneTextPkSelectBuilder {
	c.SelectBuilder.GroupBy(col...)
	return c
}

//Having returns a jHasOneTextPkSelectBuilder instead of builder.SelectBuilder.
func (c *jHasOneTextPkSelectBuilder) Having(query interface{}, value ...interface{}) *jHasOneTextPkSelectBuilder {
	c.SelectBuilder.Having(query, value...)
	return c
}

//Limit returns a jHasOneTextPkSelectBuilder instead of builder.SelectBuilder.
func (c *jHasOneTextPkSelectBuilder) Limit(n uint64) *jHasOneTextPkSelectBuilder {
	c.SelectBuilder.Limit(n)
	return c
}

//Offset returns a jHasOneTextPkSelectBuilder instead of builder.SelectBuilder.
func (c *jHasOneTextPkSelectBuilder) Offset(n uint64) *jHasOneTextPkSelectBuilder {
	c.SelectBuilder.Offset(n)
	return c
}

//OrderAsc returns a jHasOneTextPkSelectBuilder instead of builder.SelectBuilder.
func (c *jHasOneTextPkSelectBuilder) OrderAsc(col string) *jHasOneTextPkSelectBuilder {
	c.SelectBuilder.OrderAsc(col)
	return c
}

//OrderDesc returns a jHasOneTextPkSelectBuilder instead of builder.SelectBuilder.
func (c *jHasOneTextPkSelectBuilder) OrderDesc(col string) *jHasOneTextPkSelectBuilder {
	c.SelectBuilder.OrderDesc(col)
	return c
}

//OrderDir returns a jHasOneTextPkSelectBuilder instead of builder.SelectBuilder.
func (c *jHasOneTextPkSelectBuilder) OrderDir(col string, isAsc bool) *jHasOneTextPkSelectBuilder {
	c.SelectBuilder.OrderDir(col, isAsc)
	return c
}

//OrderBy returns a jHasOneTextPkSelectBuilder instead of builder.SelectBuilder.
func (c *jHasOneTextPkSelectBuilder) OrderBy(col string) *jHasOneTextPkSelectBuilder {
	c.SelectBuilder.OrderBy(col)
	return c
}

//Paginate returns a jHasOneTextPkSelectBuilder instead of builder.SelectBuilder.
func (c *jHasOneTextPkSelectBuilder) Paginate(page, perPage uint64) *jHasOneTextPkSelectBuilder {
	c.SelectBuilder.Paginate(page, perPage)
	return c
}

//Join returns a jHasOneTextPkSelectBuilder instead of builder.SelectBuilder.
func (c *jHasOneTextPkSelectBuilder) Join(table, on interface{}) *jHasOneTextPkSelectBuilder {
	c.SelectBuilder.Join(table, on)
	return c
}

//LeftJoin returns a jHasOneTextPkSelectBuilder instead of builder.SelectBuilder.
func (c *jHasOneTextPkSelectBuilder) LeftJoin(table, on interface{}) *jHasOneTextPkSelectBuilder {
	c.SelectBuilder.LeftJoin(table, on)
	return c
}

//RightJoin returns a jHasOneTextPkSelectBuilder instead of builder.SelectBuilder.
func (c *jHasOneTextPkSelectBuilder) RightJoin(table, on interface{}) *jHasOneTextPkSelectBuilder {
	c.SelectBuilder.RightJoin(table, on)
	return c
}

//FullJoin returns a jHasOneTextPkSelectBuilder instead of builder.SelectBuilder.
func (c *jHasOneTextPkSelectBuilder) FullJoin(table, on interface{}) *jHasOneTextPkSelectBuilder {
	c.SelectBuilder.FullJoin(table, on)
	return c
}

//Distinct returns a jHasOneTextPkSelectBuilder instead of builder.SelectBuilder.
func (c *jHasOneTextPkSelectBuilder) Distinct() *jHasOneTextPkSelectBuilder {
	c.SelectBuilder.Distinct()
	return c
}

// JHasOneTextPk provides a basic querier
func JHasOneTextPk(db dbr.SessionRunner) jHasOneTextPkQuerier {
	return jHasOneTextPkQuerier{
		db: db,
	}
}

type jHasOneTextPkQuerier struct {
	db dbr.SessionRunner
	as string
}

//As set alias prior building.
func (c jHasOneTextPkQuerier) As(as string) jHasOneTextPkQuerier {
	c.as = as
	return c
}

//Model returns a model with appropriate aliasing.
func (c jHasOneTextPkQuerier) Model() jHasOneTextPkModel {
	return JHasOneTextPkModel.As(c.as)
}

//Select returns a HasOneTextPk Select Builder.
func (c jHasOneTextPkQuerier) Select(what ...string) *jHasOneTextPkSelectBuilder {
	m := c.Model()
	dialect := runtime.GetDialect()
	from := dialect.QuoteIdent(m.Table())
	if m.Alias() != "" && m.Alias() != m.Table() {
		from = fmt.Sprintf("%v as %v", from, dialect.QuoteIdent(m.Alias()))
	}
	if len(what) == 0 {
		alias := m.Table()
		if m.Alias() != "" && m.Alias() != m.Table() {
			alias = m.Alias()
		}
		what = m.Fields(alias + ".*")
	}
	return &jHasOneTextPkSelectBuilder{
		as: c.as,
		SelectBuilder: &builder.SelectBuilder{
			SelectBuilder: c.db.Select(what...).From(from),
		},
	}
}

//Where returns a HasOneTextPk Select Builder.
func (c jHasOneTextPkQuerier) Where(query interface{}, value ...interface{}) *jHasOneTextPkSelectBuilder {
	return c.Select().Where(query, value...)
}

//Count returns a HasOneTextPk Select Builder to count given expressions.
func (c jHasOneTextPkQuerier) Count(what ...string) *jHasOneTextPkSelectBuilder {
	if len(what) == 0 {
		what = append(what, "*")
	}
	return c.Select("COUNT(" + strings.Join(what, ",") + ")")
}

// Insert a new HasOneTextPk, if it has autoincrement primary key, the value will be set.
// It stops on first error.
func (c jHasOneTextPkQuerier) Insert(items ...*HasOneTextPk) (sql.Result, error) {
	var res sql.Result
	var err error
	for _, data := range items {

		query := c.db.InsertInto(JHasOneTextPkModel.Table()).Columns(

			`x`,

			`related_name`,
		).Record(data)
		if runtime.Runs(drivers.Pgsql) {

			query = query.Returning(

				`id`,
			)

			var auto0 int64

			err = query.Load(

				&auto0,
			)

			data.ID = auto0

		} else {
			res, err = query.Exec()

			if err == nil {
				id, err2 := res.LastInsertId()
				if err2 != nil {
					return res, err2
				}
				data.ID = id
			}

		}
		if err != nil {
			return res, err
		}
	}
	return res, err
}

// InsertBulk inserts multiple items into the database.
// It does not post update any auto increment field.
// It builds an insert query of multiple rows and send it on the underlying connection.
func (c jHasOneTextPkQuerier) InsertBulk(items ...*HasOneTextPk) error {
	panic("todo")
}

// Update a HasOneTextPk. It stops on first error.
func (c jHasOneTextPkQuerier) Update(items ...*HasOneTextPk) (sql.Result, error) {
	var res sql.Result
	var err error
	for _, data := range items {

		query := c.db.Update(JHasOneTextPkModel.Table())

		query = query.Set(`x`, data.X)

		query = query.Set(`related_name`, data.RelatedName)

		query = query.Where("id = ?", data.ID)

		res, err = query.Exec()

		if err != nil {
			return res, err
		}
	}
	return res, err
}

// MustUpdate a HasOneTextPk. It stops on first error. It errors if an update query does not affect row.
func (c jHasOneTextPkQuerier) MustUpdate(items ...*HasOneTextPk) (sql.Result, error) {
	var res sql.Result
	var err error
	for _, data := range items {

		query := c.db.Update(JHasOneTextPkModel.Table())

		query = query.Set(`x`, data.X)

		query = query.Set(`related_name`, data.RelatedName)

		query = query.Where("id = ?", data.ID)

		res, err = query.Exec()

		if err == nil {
			if n, _ := res.RowsAffected(); n == 0 {
				x := &builder.UpdateBuilder{UpdateBuilder: query}
				err = runtime.NewNoRowsAffected(x.String())
			}
		}

		if err != nil {
			return res, err
		}
	}
	return res, err
}

// UpdateBulk updates multiple items into the database.
// It builds an update query of multiple rows and send it on the underlying connection.
func (c jHasOneTextPkQuerier) UpdateBulk(items ...*HasOneTextPk) error {
	panic("todo")
}

//Delete returns a delete builder
func (c jHasOneTextPkQuerier) Delete() *jHasOneTextPkDeleteBuilder {
	return &jHasOneTextPkDeleteBuilder{
		&builder.DeleteBuilder{
			DeleteBuilder: c.db.DeleteFrom(JHasOneTextPkModel.Table()),
		},
	}
}

// MustDelete requires the query to affeect rows.
func (c jHasOneTextPkQuerier) MustDelete() *jHasOneTextPkDeleteBuilder {
	ret := &jHasOneTextPkDeleteBuilder{
		&builder.DeleteBuilder{
			DeleteBuilder: c.db.DeleteFrom(JHasOneTextPkModel.Table()),
		},
	}
	ret.MustDelete()
	return ret
}

//DeleteByPk deletes one HasOneTextPk by its PKs
func (c jHasOneTextPkQuerier) DeleteByPk(ID int64) error {
	_, err := c.Delete().Where(

		JHasOneTextPkModel.ID.Eq(ID),
	).Exec()
	return err
}

// DeleteAll given HasOneTextPk
func (c jHasOneTextPkQuerier) DeleteAll(items ...*HasOneTextPk) (sql.Result, error) {
	q := c.Delete().Where(
		JHasOneTextPkModel.In(items...),
	)
	return q.Exec()
}

// MustDeleteAll given HasOneTextPk
func (c jHasOneTextPkQuerier) MustDeleteAll(items ...*HasOneTextPk) (sql.Result, error) {
	var res sql.Result
	var err error
	for _, d := range items {
		res, err = c.DeleteAll(d)
		if err != nil {
			return res, err
		}
		if n, e := res.RowsAffected(); e != nil {
			return res, e
		} else if n == 0 {
			q := c.Delete().Where(
				JHasOneTextPkModel.In(items...),
			)
			err = runtime.NewNoRowsAffected(q.String())
			return res, err
		}
	}
	return res, err
}

//Find one HasOneTextPk using its PKs
func (c jHasOneTextPkQuerier) Find(ID int64) (*HasOneTextPk, error) {
	return c.Select().Where(

		JHasOneTextPkModel.ID.Eq(ID),
	).Read()
}

// JoinRelated adds a JOIN to HasOneTextPk.Related
func (c *jHasOneTextPkSelectBuilder) JoinRelated(
	AsRelated string,
) *jHasOneTextPkSelectBuilder {
	dialect := runtime.GetDialect()
	on := ""
	localTable := dialect.QuoteIdent(JHasOneTextPkModel.Table())
	if c.as != "" {
		localTable = dialect.QuoteIdent(c.as)
	}
	foreiTable := dialect.QuoteIdent(JTextPkModel.Table())
	if AsRelated != "" {
		foreiTable = dialect.QuoteIdent(AsRelated)
	}

	on += fmt.Sprintf("%v.%v = %v.%v",
		localTable, dialect.QuoteIdent("related_name"),
		foreiTable, dialect.QuoteIdent("name"),
	)

	if AsRelated == "" {
		return c.Join(dbr.I(JTextPkModel.Table()), on)
	}
	return c.Join(dbr.I(JTextPkModel.Table()).As(AsRelated), on)
}

// LeftJoinRelated adds a LEFT JOIN to HasOneTextPk.Related
func (c *jHasOneTextPkSelectBuilder) LeftJoinRelated(
	AsRelated string,
) *jHasOneTextPkSelectBuilder {
	dialect := runtime.GetDialect()
	on := ""
	localTable := JHasOneTextPkModel.Table()
	if c.as != "" {
		localTable = dialect.QuoteIdent(c.as)
	}
	foreiTable := dialect.QuoteIdent(JTextPkModel.Table())
	if AsRelated != "" {
		foreiTable = dialect.QuoteIdent(AsRelated)
	}

	on += fmt.Sprintf("%v.%v = %v.%v",
		localTable, dialect.QuoteIdent("related_name"),
		foreiTable, dialect.QuoteIdent("name"),
	)

	if AsRelated == "" {
		return c.LeftJoin(dbr.I(JTextPkModel.Table()), on)
	}
	return c.LeftJoin(dbr.I(JTextPkModel.Table()).As(AsRelated), on)
}

// // RightJoinRelated adds a RIGHT JOIN to HasOneTextPk.Related
// func (c *jHasOneTextPkSelectBuilder) RightJoinRelated(
// 	AsRelated string,
// ) *jHasOneTextPkSelectBuilder {
// 	dialect := runtime.GetDialect()
// 	on := ""
// 	localTable := dialect.QuoteIdent(JHasOneTextPkModel.Table())
// 	if c.as != "" {
// 		localTable = dialect.QuoteIdent(c.as)
// 	}
// 	foreiTable := dialect.QuoteIdent(JTextPkModel.Table())
// 	if AsRelated != "" {
// 		foreiTable = dialect.QuoteIdent(AsRelated)
// 	}
//
// 	on += fmt.Sprintf("%v.%v = %v.%v",
// 		localTable, dialect.QuoteIdent("related_name"),
// 		foreiTable, dialect.QuoteIdent("name"),
// 	)
//
// 	if AsRelated == "" {
// 		return c.RightJoin(dbr.I(JTextPkModel.Table()), on)
// 	}
// 	return c.RightJoin(dbr.I(JTextPkModel.Table()).As(AsRelated), on)
// }

// FullJoinRelated adds a FULL JOIN to HasOneTextPk.Related
func (c *jHasOneTextPkSelectBuilder) FullJoinRelated(
	AsRelated string,
) *jHasOneTextPkSelectBuilder {
	dialect := runtime.GetDialect()
	on := ""
	localTable := dialect.QuoteIdent(JHasOneTextPkModel.Table())
	if c.as != "" {
		localTable = dialect.QuoteIdent(c.as)
	}
	foreiTable := dialect.QuoteIdent(JTextPkModel.Table())
	if AsRelated != "" {
		foreiTable = dialect.QuoteIdent(AsRelated)
	}

	on += fmt.Sprintf("%v.%v = %v.%v",
		localTable, dialect.QuoteIdent("related_name"),
		foreiTable, dialect.QuoteIdent("name"),
	)

	if AsRelated == "" {
		return c.FullJoin(dbr.I(JTextPkModel.Table()), on)
	}
	return c.FullJoin(dbr.I(JTextPkModel.Table()).As(AsRelated), on)
}

// Related reads associated object
func (g *HasOneTextPk) Related(db dbr.SessionRunner) (*TextPk, error) {
	q := JTextPk(db).Select()
	q = q.Where(

		JTextPkModel.Name.Eq(g.RelatedName),
	)
	return q.Read()
}

// SetRelated copies pk values to this object properties
func (g *HasOneTextPk) SetRelated(o *TextPk) *HasOneTextPk {

	if o == nil {
		g.RelatedName = nil
	} else {

		g.RelatedName = &o.Name

	}

	return g
}

// UnsetRelated set defaults values to this object properties
func (g *HasOneTextPk) UnsetRelated() *HasOneTextPk {

	var def0 *string

	g.RelatedName = def0

	g.related = nil

	return g
}

// JHasManyTextPkSetup helps to create/drop the schema
func JHasManyTextPkSetup() runtime.Setuper {
	driver := runtime.GetCurrentDriver()

	var create string
	var drop string

	if driver == drivers.Sqlite {
		create = `CREATE TABLE IF NOT EXISTS has_many_text_pk (
id INTEGER PRIMARY KEY AUTOINCREMENT,
x TEXT

)`
	} else if driver == drivers.Mysql {
		create = `CREATE TABLE IF NOT EXISTS has_many_text_pk (
id INTEGER NOT NULL AUTO_INCREMENT,
x TEXT,
PRIMARY KEY (id) 

)`
	} else if driver == drivers.Pgsql {
		create = `CREATE TABLE IF NOT EXISTS has_many_text_pk (
id SERIAL PRIMARY KEY,
x TEXT

)`
	}

	if driver == drivers.Sqlite {
		drop = `DROP TABLE IF EXISTS has_many_text_pk`
	} else if driver == drivers.Mysql {
		drop = `DROP TABLE IF EXISTS has_many_text_pk`
	} else if driver == drivers.Pgsql {
		drop = `DROP TABLE IF EXISTS has_many_text_pk`
	}

	var indexes []string

	if driver == drivers.Sqlite {

	} else if driver == drivers.Mysql {

	} else if driver == drivers.Pgsql {

	}

	return runtime.Table{
		Name:       `has_many_text_pk`,
		CreateStmt: create,
		DropStmt:   drop,
		View:       !true,
		Indexes:    indexes,
	}
}

// jHasManyTextPkModel provides helper to work with HasManyTextPk data provider
type jHasManyTextPkModel struct {
	as string

	ID builder.ValuePropertyMeta

	X builder.ValuePropertyMeta

	Relateds builder.RelPropertyMeta
}

// Eq provided items.
func (j jHasManyTextPkModel) Eq(s ...*HasManyTextPk) dbr.Builder {
	ors := []dbr.Builder{}
	for _, t := range s {
		ors = append(ors, dbr.And(

			JHasManyTextPkModel.ID.Eq(t.ID),
		))
	}
	return dbr.Or(ors...)
}

// In provided items.
func (j jHasManyTextPkModel) In(s ...*HasManyTextPk) dbr.Builder {
	ors := []dbr.Builder{}
	for _, t := range s {
		ors = append(ors, dbr.And(

			JHasManyTextPkModel.ID.Eq(t.ID),
		))
	}
	return dbr.Or(ors...)
}

// As returns a copy with an alias.
func (j jHasManyTextPkModel) As(as string) jHasManyTextPkModel {
	j.as = as

	j.ID.TableAlias = as

	j.X.TableAlias = as

	// j.Relateds.TableAlias = as

	return j
}

// Table returns the sql table name
func (j jHasManyTextPkModel) Table() string {
	return "has_many_text_pk"
}

// Alias returns the current alias
func (j jHasManyTextPkModel) Alias() string {
	if j.as == "" {
		return j.Table()
	}
	return j.as
}

// Properties returns a map of property name => meta
func (j jHasManyTextPkModel) Properties() map[string]builder.MetaProvider {
	ret := map[string]builder.MetaProvider{}

	ret["ID"] = j.ID

	ret["X"] = j.X

	ret["Relateds"] = j.Relateds

	return ret
}

// Fields returns given sql fields with appropriate aliasing.
func (j jHasManyTextPkModel) Fields(ins ...string) []string {
	dialect := runtime.GetDialect()
	if len(ins) == 0 {
		ins = append(ins, "*")
	}
	for i, in := range ins {
		if j.as != "" {
			if in == "*" {
				ins[i] = fmt.Sprintf("%v.%v", dialect.QuoteIdent(j.as), in)
			} else {
				ins[i] = fmt.Sprintf("%v.%v", dialect.QuoteIdent(j.as), dialect.QuoteIdent(in))
			}
		}
	}
	return ins
}

// JHasManyTextPkModel provides helper to work with HasManyTextPk data provider
var JHasManyTextPkModel = jHasManyTextPkModel{

	ID: builder.NewValueMeta(
		`id`, `INTEGER`,
		`ID`, `int64`,
		true, true,
	),

	X: builder.NewValueMeta(
		`x`, `TEXT`,
		`X`, `string`,
		false, false,
	),

	Relateds: builder.NewRelMeta(
		`relateds`, `[]*TextPk`,
		``, `TextPk.relateds`, ``,
		`has_many2many`,
	),
}

type jHasManyTextPkDeleteBuilder struct {
	*builder.DeleteBuilder
}

// //Build builds the sql string into given buffer using current dialect
// func (c *jHasManyTextPkDeleteBuilder) Build(b dbr.Buffer) error {
// 	return c.DeleteBuilder.Build(runtime.GetDialect(), b)
// }
// //String returns the sql string for current dialect. It returns empty string if the build returns an error.
// func (c *jHasManyTextPkDeleteBuilder) String() string {
// 	b := dbr.NewBuffer()
// 	if err := c.Build(b); err != nil {
// 		return ""
// 	}
// 	return b.String()
// }
//Where returns a jHasManyTextPkDeleteBuilder instead of builder.DeleteBuilder.
func (c *jHasManyTextPkDeleteBuilder) Where(query interface{}, value ...interface{}) *jHasManyTextPkDeleteBuilder {
	c.DeleteBuilder.Where(query, value...)
	return c
}

type jHasManyTextPkSelectBuilder struct {
	as string
	*builder.SelectBuilder
}

// //Build builds the sql string using current dialect into given bufer
// func (c *jHasManyTextPkSelectBuilder) Build(b dbr.Buffer) error {
// 	return c.SelectBuilder.Build(runtime.GetDialect(), b)
// }
// //String returns the sql string for current dialect. It returns empty string if the build returns an error.
// func (c *jHasManyTextPkSelectBuilder) String() string {
// 	b := dbr.NewBuffer()
// 	if err := c.Build(b); err != nil {
// 		return ""
// 	}
// 	return b.String()
// }

//Read evaluates current select query and load the results into a HasManyTextPk
func (c *jHasManyTextPkSelectBuilder) Read() (*HasManyTextPk, error) {
	var one HasManyTextPk
	err := c.LoadStruct(&one)
	return &one, err
}

//ReadAll evaluates current select query and load the results into a slice of HasManyTextPk
func (c *jHasManyTextPkSelectBuilder) ReadAll() ([]*HasManyTextPk, error) {
	var all []*HasManyTextPk
	_, err := c.LoadStructs(&all)
	return all, err
}

//Where returns a jHasManyTextPkSelectBuilder instead of builder.SelectBuilder.
func (c *jHasManyTextPkSelectBuilder) Where(query interface{}, value ...interface{}) *jHasManyTextPkSelectBuilder {
	c.SelectBuilder.Where(query, value...)
	return c
}

//GroupBy returns a jHasManyTextPkSelectBuilder instead of builder.SelectBuilder.
func (c *jHasManyTextPkSelectBuilder) GroupBy(col ...string) *jHasManyTextPkSelectBuilder {
	c.SelectBuilder.GroupBy(col...)
	return c
}

//Having returns a jHasManyTextPkSelectBuilder instead of builder.SelectBuilder.
func (c *jHasManyTextPkSelectBuilder) Having(query interface{}, value ...interface{}) *jHasManyTextPkSelectBuilder {
	c.SelectBuilder.Having(query, value...)
	return c
}

//Limit returns a jHasManyTextPkSelectBuilder instead of builder.SelectBuilder.
func (c *jHasManyTextPkSelectBuilder) Limit(n uint64) *jHasManyTextPkSelectBuilder {
	c.SelectBuilder.Limit(n)
	return c
}

//Offset returns a jHasManyTextPkSelectBuilder instead of builder.SelectBuilder.
func (c *jHasManyTextPkSelectBuilder) Offset(n uint64) *jHasManyTextPkSelectBuilder {
	c.SelectBuilder.Offset(n)
	return c
}

//OrderAsc returns a jHasManyTextPkSelectBuilder instead of builder.SelectBuilder.
func (c *jHasManyTextPkSelectBuilder) OrderAsc(col string) *jHasManyTextPkSelectBuilder {
	c.SelectBuilder.OrderAsc(col)
	return c
}

//OrderDesc returns a jHasManyTextPkSelectBuilder instead of builder.SelectBuilder.
func (c *jHasManyTextPkSelectBuilder) OrderDesc(col string) *jHasManyTextPkSelectBuilder {
	c.SelectBuilder.OrderDesc(col)
	return c
}

//OrderDir returns a jHasManyTextPkSelectBuilder instead of builder.SelectBuilder.
func (c *jHasManyTextPkSelectBuilder) OrderDir(col string, isAsc bool) *jHasManyTextPkSelectBuilder {
	c.SelectBuilder.OrderDir(col, isAsc)
	return c
}

//OrderBy returns a jHasManyTextPkSelectBuilder instead of builder.SelectBuilder.
func (c *jHasManyTextPkSelectBuilder) OrderBy(col string) *jHasManyTextPkSelectBuilder {
	c.SelectBuilder.OrderBy(col)
	return c
}

//Paginate returns a jHasManyTextPkSelectBuilder instead of builder.SelectBuilder.
func (c *jHasManyTextPkSelectBuilder) Paginate(page, perPage uint64) *jHasManyTextPkSelectBuilder {
	c.SelectBuilder.Paginate(page, perPage)
	return c
}

//Join returns a jHasManyTextPkSelectBuilder instead of builder.SelectBuilder.
func (c *jHasManyTextPkSelectBuilder) Join(table, on interface{}) *jHasManyTextPkSelectBuilder {
	c.SelectBuilder.Join(table, on)
	return c
}

//LeftJoin returns a jHasManyTextPkSelectBuilder instead of builder.SelectBuilder.
func (c *jHasManyTextPkSelectBuilder) LeftJoin(table, on interface{}) *jHasManyTextPkSelectBuilder {
	c.SelectBuilder.LeftJoin(table, on)
	return c
}

//RightJoin returns a jHasManyTextPkSelectBuilder instead of builder.SelectBuilder.
func (c *jHasManyTextPkSelectBuilder) RightJoin(table, on interface{}) *jHasManyTextPkSelectBuilder {
	c.SelectBuilder.RightJoin(table, on)
	return c
}

//FullJoin returns a jHasManyTextPkSelectBuilder instead of builder.SelectBuilder.
func (c *jHasManyTextPkSelectBuilder) FullJoin(table, on interface{}) *jHasManyTextPkSelectBuilder {
	c.SelectBuilder.FullJoin(table, on)
	return c
}

//Distinct returns a jHasManyTextPkSelectBuilder instead of builder.SelectBuilder.
func (c *jHasManyTextPkSelectBuilder) Distinct() *jHasManyTextPkSelectBuilder {
	c.SelectBuilder.Distinct()
	return c
}

// JHasManyTextPk provides a basic querier
func JHasManyTextPk(db dbr.SessionRunner) jHasManyTextPkQuerier {
	return jHasManyTextPkQuerier{
		db: db,
	}
}

type jHasManyTextPkQuerier struct {
	db dbr.SessionRunner
	as string
}

//As set alias prior building.
func (c jHasManyTextPkQuerier) As(as string) jHasManyTextPkQuerier {
	c.as = as
	return c
}

//Model returns a model with appropriate aliasing.
func (c jHasManyTextPkQuerier) Model() jHasManyTextPkModel {
	return JHasManyTextPkModel.As(c.as)
}

//Select returns a HasManyTextPk Select Builder.
func (c jHasManyTextPkQuerier) Select(what ...string) *jHasManyTextPkSelectBuilder {
	m := c.Model()
	dialect := runtime.GetDialect()
	from := dialect.QuoteIdent(m.Table())
	if m.Alias() != "" && m.Alias() != m.Table() {
		from = fmt.Sprintf("%v as %v", from, dialect.QuoteIdent(m.Alias()))
	}
	if len(what) == 0 {
		alias := m.Table()
		if m.Alias() != "" && m.Alias() != m.Table() {
			alias = m.Alias()
		}
		what = m.Fields(alias + ".*")
	}
	return &jHasManyTextPkSelectBuilder{
		as: c.as,
		SelectBuilder: &builder.SelectBuilder{
			SelectBuilder: c.db.Select(what...).From(from),
		},
	}
}

//Where returns a HasManyTextPk Select Builder.
func (c jHasManyTextPkQuerier) Where(query interface{}, value ...interface{}) *jHasManyTextPkSelectBuilder {
	return c.Select().Where(query, value...)
}

//Count returns a HasManyTextPk Select Builder to count given expressions.
func (c jHasManyTextPkQuerier) Count(what ...string) *jHasManyTextPkSelectBuilder {
	if len(what) == 0 {
		what = append(what, "*")
	}
	return c.Select("COUNT(" + strings.Join(what, ",") + ")")
}

// Insert a new HasManyTextPk, if it has autoincrement primary key, the value will be set.
// It stops on first error.
func (c jHasManyTextPkQuerier) Insert(items ...*HasManyTextPk) (sql.Result, error) {
	var res sql.Result
	var err error
	for _, data := range items {

		query := c.db.InsertInto(JHasManyTextPkModel.Table()).Columns(

			`x`,
		).Record(data)
		if runtime.Runs(drivers.Pgsql) {

			query = query.Returning(

				`id`,
			)

			var auto0 int64

			err = query.Load(

				&auto0,
			)

			data.ID = auto0

		} else {
			res, err = query.Exec()

			if err == nil {
				id, err2 := res.LastInsertId()
				if err2 != nil {
					return res, err2
				}
				data.ID = id
			}

		}
		if err != nil {
			return res, err
		}
	}
	return res, err
}

// InsertBulk inserts multiple items into the database.
// It does not post update any auto increment field.
// It builds an insert query of multiple rows and send it on the underlying connection.
func (c jHasManyTextPkQuerier) InsertBulk(items ...*HasManyTextPk) error {
	panic("todo")
}

// Update a HasManyTextPk. It stops on first error.
func (c jHasManyTextPkQuerier) Update(items ...*HasManyTextPk) (sql.Result, error) {
	var res sql.Result
	var err error
	for _, data := range items {

		query := c.db.Update(JHasManyTextPkModel.Table())

		query = query.Set(`x`, data.X)

		query = query.Where("id = ?", data.ID)

		res, err = query.Exec()

		if err != nil {
			return res, err
		}
	}
	return res, err
}

// MustUpdate a HasManyTextPk. It stops on first error. It errors if an update query does not affect row.
func (c jHasManyTextPkQuerier) MustUpdate(items ...*HasManyTextPk) (sql.Result, error) {
	var res sql.Result
	var err error
	for _, data := range items {

		query := c.db.Update(JHasManyTextPkModel.Table())

		query = query.Set(`x`, data.X)

		query = query.Where("id = ?", data.ID)

		res, err = query.Exec()

		if err == nil {
			if n, _ := res.RowsAffected(); n == 0 {
				x := &builder.UpdateBuilder{UpdateBuilder: query}
				err = runtime.NewNoRowsAffected(x.String())
			}
		}

		if err != nil {
			return res, err
		}
	}
	return res, err
}

// UpdateBulk updates multiple items into the database.
// It builds an update query of multiple rows and send it on the underlying connection.
func (c jHasManyTextPkQuerier) UpdateBulk(items ...*HasManyTextPk) error {
	panic("todo")
}

//Delete returns a delete builder
func (c jHasManyTextPkQuerier) Delete() *jHasManyTextPkDeleteBuilder {
	return &jHasManyTextPkDeleteBuilder{
		&builder.DeleteBuilder{
			DeleteBuilder: c.db.DeleteFrom(JHasManyTextPkModel.Table()),
		},
	}
}

// MustDelete requires the query to affeect rows.
func (c jHasManyTextPkQuerier) MustDelete() *jHasManyTextPkDeleteBuilder {
	ret := &jHasManyTextPkDeleteBuilder{
		&builder.DeleteBuilder{
			DeleteBuilder: c.db.DeleteFrom(JHasManyTextPkModel.Table()),
		},
	}
	ret.MustDelete()
	return ret
}

//DeleteByPk deletes one HasManyTextPk by its PKs
func (c jHasManyTextPkQuerier) DeleteByPk(ID int64) error {
	_, err := c.Delete().Where(

		JHasManyTextPkModel.ID.Eq(ID),
	).Exec()
	return err
}

// DeleteAll given HasManyTextPk
func (c jHasManyTextPkQuerier) DeleteAll(items ...*HasManyTextPk) (sql.Result, error) {
	q := c.Delete().Where(
		JHasManyTextPkModel.In(items...),
	)
	return q.Exec()
}

// MustDeleteAll given HasManyTextPk
func (c jHasManyTextPkQuerier) MustDeleteAll(items ...*HasManyTextPk) (sql.Result, error) {
	var res sql.Result
	var err error
	for _, d := range items {
		res, err = c.DeleteAll(d)
		if err != nil {
			return res, err
		}
		if n, e := res.RowsAffected(); e != nil {
			return res, e
		} else if n == 0 {
			q := c.Delete().Where(
				JHasManyTextPkModel.In(items...),
			)
			err = runtime.NewNoRowsAffected(q.String())
			return res, err
		}
	}
	return res, err
}

//Find one HasManyTextPk using its PKs
func (c jHasManyTextPkQuerier) Find(ID int64) (*HasManyTextPk, error) {
	return c.Select().Where(

		JHasManyTextPkModel.ID.Eq(ID),
	).Read()
}

// Relateds returns a query builder to select Relateds linked to this HasManyTextPk
func (g *HasManyTextPk) Relateds(db dbr.SessionRunner,
	AsTextPk, AsHasManyTextPkrelatedsToTextPkrelateds, AsHasManyTextPk string,
) *jTextPkSelectBuilder {

	leftTable := JTextPkModel.Table()
	var query *jTextPkSelectBuilder
	if AsTextPk != "" {
		leftTable = AsTextPk
		query = JTextPk(db).As(AsTextPk).Select(AsTextPk + ".*")
	} else {
		query = JTextPk(db).Select(leftTable + ".*")
	}

	midTable := JHasManyTextPkrelatedsToTextPkrelatedsModel.Table()
	{
		on := ""
		if AsHasManyTextPkrelatedsToTextPkrelateds != "" {
			midTable = AsHasManyTextPkrelatedsToTextPkrelateds
		}

		on += fmt.Sprintf("%v.%v = %v.%v",
			midTable, "text_pk_name",
			leftTable, "name",
		)

		if AsHasManyTextPkrelatedsToTextPkrelateds == "" {
			query = query.Join(dbr.I(JHasManyTextPkrelatedsToTextPkrelatedsModel.Table()), on)
		} else {
			query = query.Join(dbr.I(JHasManyTextPkrelatedsToTextPkrelatedsModel.Table()).As(AsHasManyTextPkrelatedsToTextPkrelateds), on)
		}
	}

	rightTable := JHasManyTextPkModel.Table()
	{
		on := ""
		if AsHasManyTextPk != "" {
			rightTable = AsHasManyTextPk
		}

		on += fmt.Sprintf("%v.%v = %v.%v",
			midTable, "has_many_text_pk_id",
			rightTable, "id",
		)

		if AsHasManyTextPk == "" {
			query = query.Join(dbr.I(JHasManyTextPkModel.Table()), on)
		} else {
			query = query.Join(dbr.I(JHasManyTextPkModel.Table()).As(AsHasManyTextPk), on)
		}
	}

	{
		m := JHasManyTextPkModel
		if AsHasManyTextPk != "" {
			m = m.As(AsHasManyTextPk)
		} else {
			m = m.As(m.Table())
		}
		query = query.Where(

			m.ID.Eq(g.ID),
		)
	}

	return query
}

//LinkWithRelateds writes new links with HasManyTextPk.
func (g *HasManyTextPk) LinkWithRelateds(db dbr.SessionRunner, items ...*TextPk) (sql.Result, error) {
	toInsert := []*HasManyTextPkrelatedsToTextPkrelateds{}
	for _, item := range items {
		toInsert = append(toInsert, &HasManyTextPkrelatedsToTextPkrelateds{

			TextPkName: item.Name,

			HasManyTextPkID: g.ID,
		})
	}
	return JHasManyTextPkrelatedsToTextPkrelateds(db).Insert(toInsert...)
}

//UnlinkWithRelateds deletes given existing links with HasManyTextPk.
func (g *HasManyTextPk) UnlinkWithRelateds(db dbr.SessionRunner, items ...*TextPk) (sql.Result, error) {
	toDelete := []*HasManyTextPkrelatedsToTextPkrelateds{}
	for _, item := range items {
		toDelete = append(toDelete, &HasManyTextPkrelatedsToTextPkrelateds{

			TextPkName: item.Name,

			HasManyTextPkID: g.ID,
		})
	}
	return JHasManyTextPkrelatedsToTextPkrelateds(db).DeleteAll(toDelete...)
}

//UnlinkAllRelateds deletes all existing links with HasManyTextPk.
func (g *HasManyTextPk) UnlinkAllRelateds(db dbr.SessionRunner) (sql.Result, error) {
	return JHasManyTextPkrelatedsToTextPkrelateds(db).Delete().Where(

		JHasManyTextPkrelatedsToTextPkrelatedsModel.HasManyTextPkID.Eq(g.ID),
	).Exec()
}

//SetRelateds replaces existing links with HasManyTextPk.
func (g *HasManyTextPk) SetRelateds(db dbr.SessionRunner, items ...*TextPk) (sql.Result, error) {
	if res, err := g.UnlinkAllRelateds(db); err != nil {
		return res, err
	}
	return g.LinkWithRelateds(db, items...)
}

// JoinRelateds adds a JOIN to HasManyTextPk.Relateds
func (c *jHasManyTextPkSelectBuilder) JoinRelateds(
	AsHasManyTextPkrelatedsToTextPkrelateds, AsTextPk string,
) *jHasManyTextPkSelectBuilder {

	query := c

	leftTable := JHasManyTextPkModel.Table()
	if c.as != "" {
		leftTable = c.as
	}

	midTable := JHasManyTextPkrelatedsToTextPkrelatedsModel.Table()
	if AsHasManyTextPkrelatedsToTextPkrelateds != "" {
		midTable = AsHasManyTextPkrelatedsToTextPkrelateds
	}

	{
		on := ""

		on += fmt.Sprintf("%v.%v = %v.%v",
			midTable, "has_many_text_pk_id",
			leftTable, "id",
		)

		if AsHasManyTextPkrelatedsToTextPkrelateds == "" {
			query = query.Join(dbr.I(JHasManyTextPkrelatedsToTextPkrelatedsModel.Table()), on)
		} else {
			query = query.Join(dbr.I(JHasManyTextPkrelatedsToTextPkrelatedsModel.Table()).As(AsHasManyTextPkrelatedsToTextPkrelateds), on)
		}
	}

	{
		rightTable := JTextPkModel.Table()
		if AsTextPk != "" {
			rightTable = AsTextPk
		}
		on := ""

		on += fmt.Sprintf("%v.%v = %v.%v",
			midTable, "text_pk_name",
			rightTable, "name",
		)

		if AsTextPk == "" {
			query = query.Join(dbr.I(JTextPkModel.Table()), on)
		} else {
			query = query.Join(dbr.I(JTextPkModel.Table()).As(AsTextPk), on)
		}
	}

	return query
}

// LeftJoinRelateds adds a LEFT JOIN to HasManyTextPk.Relateds
func (c *jHasManyTextPkSelectBuilder) LeftJoinRelateds(
	AsHasManyTextPkrelatedsToTextPkrelateds, AsTextPk string,
) *jHasManyTextPkSelectBuilder {

	query := c

	leftTable := JHasManyTextPkModel.Table()
	if c.as != "" {
		leftTable = c.as
	}

	midTable := JHasManyTextPkrelatedsToTextPkrelatedsModel.Table()
	if AsHasManyTextPkrelatedsToTextPkrelateds != "" {
		midTable = AsHasManyTextPkrelatedsToTextPkrelateds
	}

	{
		on := ""

		on += fmt.Sprintf("%v.%v = %v.%v",
			midTable, "has_many_text_pk_id",
			leftTable, "id",
		)

		if AsHasManyTextPkrelatedsToTextPkrelateds == "" {
			query = query.LeftJoin(dbr.I(JHasManyTextPkrelatedsToTextPkrelatedsModel.Table()), on)
		} else {
			query = query.LeftJoin(dbr.I(JHasManyTextPkrelatedsToTextPkrelatedsModel.Table()).As(AsHasManyTextPkrelatedsToTextPkrelateds), on)
		}
	}

	{
		rightTable := JTextPkModel.Table()
		if AsTextPk != "" {
			rightTable = AsTextPk
		}
		on := ""

		on += fmt.Sprintf("%v.%v = %v.%v",
			midTable, "text_pk_name",
			rightTable, "name",
		)

		if AsTextPk == "" {
			query = query.LeftJoin(dbr.I(JTextPkModel.Table()), on)
		} else {
			query = query.LeftJoin(dbr.I(JTextPkModel.Table()).As(AsTextPk), on)
		}
	}

	return query
}

// // RightJoinRelateds adds a RIGHT JOIN to HasManyTextPk.Relateds
// func (c *jHasManyTextPkSelectBuilder) RightJoinRelateds(
// 	AsHasManyTextPkrelatedsToTextPkrelateds, AsTextPk string,
// ) *jHasManyTextPkSelectBuilder {
//
// 	query := c
//
// 	leftTable := JHasManyTextPkModel.Table()
// 	if c.as != "" {
// 		leftTable = c.as
// 	}
//
// 	midTable := JHasManyTextPkrelatedsToTextPkrelatedsModel.Table()
// 	if AsHasManyTextPkrelatedsToTextPkrelateds != "" {
// 		midTable = AsHasManyTextPkrelatedsToTextPkrelateds
// 	}
//
// 	{
// 		on := ""
//
// 		on += fmt.Sprintf("%v.%v = %v.%v",
// 			midTable, "has_many_text_pk_id",
// 			leftTable, "id",
// 			)
//
//
// 		if AsHasManyTextPkrelatedsToTextPkrelateds == "" {
// 			query = query.RightJoin(dbr.I(JHasManyTextPkrelatedsToTextPkrelatedsModel.Table()), on)
// 		} else {
// 			query = query.RightJoin(dbr.I(JHasManyTextPkrelatedsToTextPkrelatedsModel.Table()).As(AsHasManyTextPkrelatedsToTextPkrelateds), on)
// 		}
// 	}
//
// 	{
// 		rightTable := JTextPkModel.Table()
// 		if AsTextPk != "" {
// 			rightTable = AsTextPk
// 		}
// 		on := ""
//
// 		on += fmt.Sprintf("%v.%v = %v.%v",
// 			midTable, "text_pk_name",
// 			rightTable, "name",
// 			)
//
//
// 		if AsTextPk == "" {
// 			query = query.RightJoin(dbr.I(JTextPkModel.Table()), on)
// 		} else {
// 			query = query.RightJoin(dbr.I(JTextPkModel.Table()).As(AsTextPk), on)
// 		}
// 	}
//
// 	return query
// }

// JCompositePkSetup helps to create/drop the schema
func JCompositePkSetup() runtime.Setuper {
	driver := runtime.GetCurrentDriver()

	var create string
	var drop string

	if driver == drivers.Sqlite {
		create = `CREATE TABLE IF NOT EXISTS composite_pk (
p TEXT,
k TEXT,
description TEXT,
PRIMARY KEY (p,k) 

)`
	} else if driver == drivers.Mysql {
		create = `CREATE TABLE IF NOT EXISTS composite_pk (
p VARCHAR(255) NOT NULL,
k VARCHAR(255) NOT NULL,
description TEXT,
PRIMARY KEY (p,k) 

)`
	} else if driver == drivers.Pgsql {
		create = `CREATE TABLE IF NOT EXISTS composite_pk (
p TEXT,
k TEXT,
description TEXT,
PRIMARY KEY (p,k) 

)`
	}

	if driver == drivers.Sqlite {
		drop = `DROP TABLE IF EXISTS composite_pk`
	} else if driver == drivers.Mysql {
		drop = `DROP TABLE IF EXISTS composite_pk`
	} else if driver == drivers.Pgsql {
		drop = `DROP TABLE IF EXISTS composite_pk`
	}

	var indexes []string

	if driver == drivers.Sqlite {

	} else if driver == drivers.Mysql {

	} else if driver == drivers.Pgsql {

	}

	return runtime.Table{
		Name:       `composite_pk`,
		CreateStmt: create,
		DropStmt:   drop,
		View:       !true,
		Indexes:    indexes,
	}
}

// jCompositePkModel provides helper to work with CompositePk data provider
type jCompositePkModel struct {
	as string

	P builder.ValuePropertyMeta

	K builder.ValuePropertyMeta

	Description builder.ValuePropertyMeta

	HasManyHasOneCompositePk builder.RelPropertyMeta

	Relateds builder.RelPropertyMeta
}

// Eq provided items.
func (j jCompositePkModel) Eq(s ...*CompositePk) dbr.Builder {
	ors := []dbr.Builder{}
	for _, t := range s {
		ors = append(ors, dbr.And(

			JCompositePkModel.P.Eq(t.P),

			JCompositePkModel.K.Eq(t.K),
		))
	}
	return dbr.Or(ors...)
}

// In provided items.
func (j jCompositePkModel) In(s ...*CompositePk) dbr.Builder {
	ors := []dbr.Builder{}
	for _, t := range s {
		ors = append(ors, dbr.And(

			JCompositePkModel.P.Eq(t.P),

			JCompositePkModel.K.Eq(t.K),
		))
	}
	return dbr.Or(ors...)
}

// As returns a copy with an alias.
func (j jCompositePkModel) As(as string) jCompositePkModel {
	j.as = as

	j.P.TableAlias = as

	j.K.TableAlias = as

	j.Description.TableAlias = as

	// j.HasManyHasOneCompositePk.TableAlias = as

	// j.Relateds.TableAlias = as

	return j
}

// Table returns the sql table name
func (j jCompositePkModel) Table() string {
	return "composite_pk"
}

// Alias returns the current alias
func (j jCompositePkModel) Alias() string {
	if j.as == "" {
		return j.Table()
	}
	return j.as
}

// Properties returns a map of property name => meta
func (j jCompositePkModel) Properties() map[string]builder.MetaProvider {
	ret := map[string]builder.MetaProvider{}

	ret["P"] = j.P

	ret["K"] = j.K

	ret["Description"] = j.Description

	ret["HasManyHasOneCompositePk"] = j.HasManyHasOneCompositePk

	ret["Relateds"] = j.Relateds

	return ret
}

// Fields returns given sql fields with appropriate aliasing.
func (j jCompositePkModel) Fields(ins ...string) []string {
	dialect := runtime.GetDialect()
	if len(ins) == 0 {
		ins = append(ins, "*")
	}
	for i, in := range ins {
		if j.as != "" {
			if in == "*" {
				ins[i] = fmt.Sprintf("%v.%v", dialect.QuoteIdent(j.as), in)
			} else {
				ins[i] = fmt.Sprintf("%v.%v", dialect.QuoteIdent(j.as), dialect.QuoteIdent(in))
			}
		}
	}
	return ins
}

// JCompositePkModel provides helper to work with CompositePk data provider
var JCompositePkModel = jCompositePkModel{

	P: builder.NewValueMeta(
		`p`, `TEXT`,
		`P`, `string`,
		true, false,
	),

	K: builder.NewValueMeta(
		`k`, `TEXT`,
		`K`, `string`,
		true, false,
	),

	Description: builder.NewValueMeta(
		`description`, `TEXT`,
		`Description`, `string`,
		false, false,
	),

	HasManyHasOneCompositePk: builder.NewRelMeta(
		`hasManyHasOneCompositePk`, `[]*HasOneCompositePk`,
		``, `HasOneCompositePk.related`, ``,
		`has_many`,
	),

	Relateds: builder.NewRelMeta(
		`relateds`, `[]*HasManyCompositePk`,
		``, `HasManyCompositePk.relateds`, ``,
		`has_many2many`,
	),
}

type jCompositePkDeleteBuilder struct {
	*builder.DeleteBuilder
}

// //Build builds the sql string into given buffer using current dialect
// func (c *jCompositePkDeleteBuilder) Build(b dbr.Buffer) error {
// 	return c.DeleteBuilder.Build(runtime.GetDialect(), b)
// }
// //String returns the sql string for current dialect. It returns empty string if the build returns an error.
// func (c *jCompositePkDeleteBuilder) String() string {
// 	b := dbr.NewBuffer()
// 	if err := c.Build(b); err != nil {
// 		return ""
// 	}
// 	return b.String()
// }
//Where returns a jCompositePkDeleteBuilder instead of builder.DeleteBuilder.
func (c *jCompositePkDeleteBuilder) Where(query interface{}, value ...interface{}) *jCompositePkDeleteBuilder {
	c.DeleteBuilder.Where(query, value...)
	return c
}

type jCompositePkSelectBuilder struct {
	as string
	*builder.SelectBuilder
}

// //Build builds the sql string using current dialect into given bufer
// func (c *jCompositePkSelectBuilder) Build(b dbr.Buffer) error {
// 	return c.SelectBuilder.Build(runtime.GetDialect(), b)
// }
// //String returns the sql string for current dialect. It returns empty string if the build returns an error.
// func (c *jCompositePkSelectBuilder) String() string {
// 	b := dbr.NewBuffer()
// 	if err := c.Build(b); err != nil {
// 		return ""
// 	}
// 	return b.String()
// }

//Read evaluates current select query and load the results into a CompositePk
func (c *jCompositePkSelectBuilder) Read() (*CompositePk, error) {
	var one CompositePk
	err := c.LoadStruct(&one)
	return &one, err
}

//ReadAll evaluates current select query and load the results into a slice of CompositePk
func (c *jCompositePkSelectBuilder) ReadAll() ([]*CompositePk, error) {
	var all []*CompositePk
	_, err := c.LoadStructs(&all)
	return all, err
}

//Where returns a jCompositePkSelectBuilder instead of builder.SelectBuilder.
func (c *jCompositePkSelectBuilder) Where(query interface{}, value ...interface{}) *jCompositePkSelectBuilder {
	c.SelectBuilder.Where(query, value...)
	return c
}

//GroupBy returns a jCompositePkSelectBuilder instead of builder.SelectBuilder.
func (c *jCompositePkSelectBuilder) GroupBy(col ...string) *jCompositePkSelectBuilder {
	c.SelectBuilder.GroupBy(col...)
	return c
}

//Having returns a jCompositePkSelectBuilder instead of builder.SelectBuilder.
func (c *jCompositePkSelectBuilder) Having(query interface{}, value ...interface{}) *jCompositePkSelectBuilder {
	c.SelectBuilder.Having(query, value...)
	return c
}

//Limit returns a jCompositePkSelectBuilder instead of builder.SelectBuilder.
func (c *jCompositePkSelectBuilder) Limit(n uint64) *jCompositePkSelectBuilder {
	c.SelectBuilder.Limit(n)
	return c
}

//Offset returns a jCompositePkSelectBuilder instead of builder.SelectBuilder.
func (c *jCompositePkSelectBuilder) Offset(n uint64) *jCompositePkSelectBuilder {
	c.SelectBuilder.Offset(n)
	return c
}

//OrderAsc returns a jCompositePkSelectBuilder instead of builder.SelectBuilder.
func (c *jCompositePkSelectBuilder) OrderAsc(col string) *jCompositePkSelectBuilder {
	c.SelectBuilder.OrderAsc(col)
	return c
}

//OrderDesc returns a jCompositePkSelectBuilder instead of builder.SelectBuilder.
func (c *jCompositePkSelectBuilder) OrderDesc(col string) *jCompositePkSelectBuilder {
	c.SelectBuilder.OrderDesc(col)
	return c
}

//OrderDir returns a jCompositePkSelectBuilder instead of builder.SelectBuilder.
func (c *jCompositePkSelectBuilder) OrderDir(col string, isAsc bool) *jCompositePkSelectBuilder {
	c.SelectBuilder.OrderDir(col, isAsc)
	return c
}

//OrderBy returns a jCompositePkSelectBuilder instead of builder.SelectBuilder.
func (c *jCompositePkSelectBuilder) OrderBy(col string) *jCompositePkSelectBuilder {
	c.SelectBuilder.OrderBy(col)
	return c
}

//Paginate returns a jCompositePkSelectBuilder instead of builder.SelectBuilder.
func (c *jCompositePkSelectBuilder) Paginate(page, perPage uint64) *jCompositePkSelectBuilder {
	c.SelectBuilder.Paginate(page, perPage)
	return c
}

//Join returns a jCompositePkSelectBuilder instead of builder.SelectBuilder.
func (c *jCompositePkSelectBuilder) Join(table, on interface{}) *jCompositePkSelectBuilder {
	c.SelectBuilder.Join(table, on)
	return c
}

//LeftJoin returns a jCompositePkSelectBuilder instead of builder.SelectBuilder.
func (c *jCompositePkSelectBuilder) LeftJoin(table, on interface{}) *jCompositePkSelectBuilder {
	c.SelectBuilder.LeftJoin(table, on)
	return c
}

//RightJoin returns a jCompositePkSelectBuilder instead of builder.SelectBuilder.
func (c *jCompositePkSelectBuilder) RightJoin(table, on interface{}) *jCompositePkSelectBuilder {
	c.SelectBuilder.RightJoin(table, on)
	return c
}

//FullJoin returns a jCompositePkSelectBuilder instead of builder.SelectBuilder.
func (c *jCompositePkSelectBuilder) FullJoin(table, on interface{}) *jCompositePkSelectBuilder {
	c.SelectBuilder.FullJoin(table, on)
	return c
}

//Distinct returns a jCompositePkSelectBuilder instead of builder.SelectBuilder.
func (c *jCompositePkSelectBuilder) Distinct() *jCompositePkSelectBuilder {
	c.SelectBuilder.Distinct()
	return c
}

// JCompositePk provides a basic querier
func JCompositePk(db dbr.SessionRunner) jCompositePkQuerier {
	return jCompositePkQuerier{
		db: db,
	}
}

type jCompositePkQuerier struct {
	db dbr.SessionRunner
	as string
}

//As set alias prior building.
func (c jCompositePkQuerier) As(as string) jCompositePkQuerier {
	c.as = as
	return c
}

//Model returns a model with appropriate aliasing.
func (c jCompositePkQuerier) Model() jCompositePkModel {
	return JCompositePkModel.As(c.as)
}

//Select returns a CompositePk Select Builder.
func (c jCompositePkQuerier) Select(what ...string) *jCompositePkSelectBuilder {
	m := c.Model()
	dialect := runtime.GetDialect()
	from := dialect.QuoteIdent(m.Table())
	if m.Alias() != "" && m.Alias() != m.Table() {
		from = fmt.Sprintf("%v as %v", from, dialect.QuoteIdent(m.Alias()))
	}
	if len(what) == 0 {
		alias := m.Table()
		if m.Alias() != "" && m.Alias() != m.Table() {
			alias = m.Alias()
		}
		what = m.Fields(alias + ".*")
	}
	return &jCompositePkSelectBuilder{
		as: c.as,
		SelectBuilder: &builder.SelectBuilder{
			SelectBuilder: c.db.Select(what...).From(from),
		},
	}
}

//Where returns a CompositePk Select Builder.
func (c jCompositePkQuerier) Where(query interface{}, value ...interface{}) *jCompositePkSelectBuilder {
	return c.Select().Where(query, value...)
}

//Count returns a CompositePk Select Builder to count given expressions.
func (c jCompositePkQuerier) Count(what ...string) *jCompositePkSelectBuilder {
	if len(what) == 0 {
		what = append(what, "*")
	}
	return c.Select("COUNT(" + strings.Join(what, ",") + ")")
}

// Insert a new CompositePk, if it has autoincrement primary key, the value will be set.
// It stops on first error.
func (c jCompositePkQuerier) Insert(items ...*CompositePk) (sql.Result, error) {
	var res sql.Result
	var err error
	for _, data := range items {

		if runtime.Runs(drivers.Mysql) {

			if len(data.P) > 255 {
				return nil, fmt.Errorf("P: PRIMARY KEY length exceeded max=255, got=%v", len(data.P))
			}

		}

		if runtime.Runs(drivers.Mysql) {

			if len(data.K) > 255 {
				return nil, fmt.Errorf("K: PRIMARY KEY length exceeded max=255, got=%v", len(data.K))
			}

		}

		query := c.db.InsertInto(JCompositePkModel.Table()).Columns(

			`p`,

			`k`,

			`description`,
		).Record(data)
		if runtime.Runs(drivers.Pgsql) {

			res, err = query.Exec()

		} else {
			res, err = query.Exec()

		}
		if err != nil {
			return res, err
		}
	}
	return res, err
}

// InsertBulk inserts multiple items into the database.
// It does not post update any auto increment field.
// It builds an insert query of multiple rows and send it on the underlying connection.
func (c jCompositePkQuerier) InsertBulk(items ...*CompositePk) error {
	panic("todo")
}

// Update a CompositePk. It stops on first error.
func (c jCompositePkQuerier) Update(items ...*CompositePk) (sql.Result, error) {
	var res sql.Result
	var err error
	for _, data := range items {

		query := c.db.Update(JCompositePkModel.Table())

		query = query.Set(`description`, data.Description)

		query = query.Where("p = ?", data.P)

		query = query.Where("k = ?", data.K)

		res, err = query.Exec()

		if err != nil {
			return res, err
		}
	}
	return res, err
}

// MustUpdate a CompositePk. It stops on first error. It errors if an update query does not affect row.
func (c jCompositePkQuerier) MustUpdate(items ...*CompositePk) (sql.Result, error) {
	var res sql.Result
	var err error
	for _, data := range items {

		query := c.db.Update(JCompositePkModel.Table())

		query = query.Set(`description`, data.Description)

		query = query.Where("p = ?", data.P)

		query = query.Where("k = ?", data.K)

		res, err = query.Exec()

		if err == nil {
			if n, _ := res.RowsAffected(); n == 0 {
				x := &builder.UpdateBuilder{UpdateBuilder: query}
				err = runtime.NewNoRowsAffected(x.String())
			}
		}

		if err != nil {
			return res, err
		}
	}
	return res, err
}

// UpdateBulk updates multiple items into the database.
// It builds an update query of multiple rows and send it on the underlying connection.
func (c jCompositePkQuerier) UpdateBulk(items ...*CompositePk) error {
	panic("todo")
}

//Delete returns a delete builder
func (c jCompositePkQuerier) Delete() *jCompositePkDeleteBuilder {
	return &jCompositePkDeleteBuilder{
		&builder.DeleteBuilder{
			DeleteBuilder: c.db.DeleteFrom(JCompositePkModel.Table()),
		},
	}
}

// MustDelete requires the query to affeect rows.
func (c jCompositePkQuerier) MustDelete() *jCompositePkDeleteBuilder {
	ret := &jCompositePkDeleteBuilder{
		&builder.DeleteBuilder{
			DeleteBuilder: c.db.DeleteFrom(JCompositePkModel.Table()),
		},
	}
	ret.MustDelete()
	return ret
}

//DeleteByPk deletes one CompositePk by its PKs
func (c jCompositePkQuerier) DeleteByPk(P string, K string) error {
	_, err := c.Delete().Where(

		JCompositePkModel.P.Eq(P),

		JCompositePkModel.K.Eq(K),
	).Exec()
	return err
}

// DeleteAll given CompositePk
func (c jCompositePkQuerier) DeleteAll(items ...*CompositePk) (sql.Result, error) {
	q := c.Delete().Where(
		JCompositePkModel.In(items...),
	)
	return q.Exec()
}

// MustDeleteAll given CompositePk
func (c jCompositePkQuerier) MustDeleteAll(items ...*CompositePk) (sql.Result, error) {
	var res sql.Result
	var err error
	for _, d := range items {
		res, err = c.DeleteAll(d)
		if err != nil {
			return res, err
		}
		if n, e := res.RowsAffected(); e != nil {
			return res, e
		} else if n == 0 {
			q := c.Delete().Where(
				JCompositePkModel.In(items...),
			)
			err = runtime.NewNoRowsAffected(q.String())
			return res, err
		}
	}
	return res, err
}

//Find one CompositePk using its PKs
func (c jCompositePkQuerier) Find(P string, K string) (*CompositePk, error) {
	return c.Select().Where(

		JCompositePkModel.P.Eq(P),

		JCompositePkModel.K.Eq(K),
	).Read()
}

// Relateds returns a query builder to select Relateds linked to this CompositePk
func (g *CompositePk) Relateds(db dbr.SessionRunner,
	AsHasManyCompositePk, AsCompositePkrelatedsToHasManyCompositePkrelateds, AsCompositePk string,
) *jHasManyCompositePkSelectBuilder {

	leftTable := JHasManyCompositePkModel.Table()
	var query *jHasManyCompositePkSelectBuilder
	if AsHasManyCompositePk != "" {
		leftTable = AsHasManyCompositePk
		query = JHasManyCompositePk(db).As(AsHasManyCompositePk).Select(AsHasManyCompositePk + ".*")
	} else {
		query = JHasManyCompositePk(db).Select(leftTable + ".*")
	}

	midTable := JCompositePkrelatedsToHasManyCompositePkrelatedsModel.Table()
	{
		on := ""
		if AsCompositePkrelatedsToHasManyCompositePkrelateds != "" {
			midTable = AsCompositePkrelatedsToHasManyCompositePkrelateds
		}

		on += fmt.Sprintf("%v.%v = %v.%v",
			midTable, "has_many_composite_pk_id",
			leftTable, "id",
		)

		if AsCompositePkrelatedsToHasManyCompositePkrelateds == "" {
			query = query.Join(dbr.I(JCompositePkrelatedsToHasManyCompositePkrelatedsModel.Table()), on)
		} else {
			query = query.Join(dbr.I(JCompositePkrelatedsToHasManyCompositePkrelatedsModel.Table()).As(AsCompositePkrelatedsToHasManyCompositePkrelateds), on)
		}
	}

	rightTable := JCompositePkModel.Table()
	{
		on := ""
		if AsCompositePk != "" {
			rightTable = AsCompositePk
		}

		on += fmt.Sprintf("%v.%v = %v.%v",
			midTable, "composite_pk_p",
			rightTable, "p",
		)

		on += fmt.Sprintf("%v.%v = %v.%v",
			midTable, "composite_pk_k",
			rightTable, "k",
		)

		if AsCompositePk == "" {
			query = query.Join(dbr.I(JCompositePkModel.Table()), on)
		} else {
			query = query.Join(dbr.I(JCompositePkModel.Table()).As(AsCompositePk), on)
		}
	}

	{
		m := JCompositePkModel
		if AsCompositePk != "" {
			m = m.As(AsCompositePk)
		} else {
			m = m.As(m.Table())
		}
		query = query.Where(

			m.P.Eq(g.P),

			m.K.Eq(g.K),
		)
	}

	return query
}

//LinkWithRelateds writes new links with CompositePk.
func (g *CompositePk) LinkWithRelateds(db dbr.SessionRunner, items ...*HasManyCompositePk) (sql.Result, error) {
	toInsert := []*CompositePkrelatedsToHasManyCompositePkrelateds{}
	for _, item := range items {
		toInsert = append(toInsert, &CompositePkrelatedsToHasManyCompositePkrelateds{

			HasManyCompositePkID: item.ID,

			CompositePkP: g.P,

			CompositePkK: g.K,
		})
	}
	return JCompositePkrelatedsToHasManyCompositePkrelateds(db).Insert(toInsert...)
}

//UnlinkWithRelateds deletes given existing links with CompositePk.
func (g *CompositePk) UnlinkWithRelateds(db dbr.SessionRunner, items ...*HasManyCompositePk) (sql.Result, error) {
	toDelete := []*CompositePkrelatedsToHasManyCompositePkrelateds{}
	for _, item := range items {
		toDelete = append(toDelete, &CompositePkrelatedsToHasManyCompositePkrelateds{

			HasManyCompositePkID: item.ID,

			CompositePkP: g.P,

			CompositePkK: g.K,
		})
	}
	return JCompositePkrelatedsToHasManyCompositePkrelateds(db).DeleteAll(toDelete...)
}

//UnlinkAllRelateds deletes all existing links with CompositePk.
func (g *CompositePk) UnlinkAllRelateds(db dbr.SessionRunner) (sql.Result, error) {
	return JCompositePkrelatedsToHasManyCompositePkrelateds(db).Delete().Where(

		JCompositePkrelatedsToHasManyCompositePkrelatedsModel.CompositePkP.Eq(g.P),

		JCompositePkrelatedsToHasManyCompositePkrelatedsModel.CompositePkK.Eq(g.K),
	).Exec()
}

//SetRelateds replaces existing links with CompositePk.
func (g *CompositePk) SetRelateds(db dbr.SessionRunner, items ...*HasManyCompositePk) (sql.Result, error) {
	if res, err := g.UnlinkAllRelateds(db); err != nil {
		return res, err
	}
	return g.LinkWithRelateds(db, items...)
}

// JoinRelateds adds a JOIN to CompositePk.Relateds
func (c *jCompositePkSelectBuilder) JoinRelateds(
	AsCompositePkrelatedsToHasManyCompositePkrelateds, AsHasManyCompositePk string,
) *jCompositePkSelectBuilder {

	query := c

	leftTable := JCompositePkModel.Table()
	if c.as != "" {
		leftTable = c.as
	}

	midTable := JCompositePkrelatedsToHasManyCompositePkrelatedsModel.Table()
	if AsCompositePkrelatedsToHasManyCompositePkrelateds != "" {
		midTable = AsCompositePkrelatedsToHasManyCompositePkrelateds
	}

	{
		on := ""

		on += fmt.Sprintf("%v.%v = %v.%v",
			midTable, "composite_pk_p",
			leftTable, "p",
		)

		on += fmt.Sprintf("%v.%v = %v.%v",
			midTable, "composite_pk_k",
			leftTable, "k",
		)

		if AsCompositePkrelatedsToHasManyCompositePkrelateds == "" {
			query = query.Join(dbr.I(JCompositePkrelatedsToHasManyCompositePkrelatedsModel.Table()), on)
		} else {
			query = query.Join(dbr.I(JCompositePkrelatedsToHasManyCompositePkrelatedsModel.Table()).As(AsCompositePkrelatedsToHasManyCompositePkrelateds), on)
		}
	}

	{
		rightTable := JHasManyCompositePkModel.Table()
		if AsHasManyCompositePk != "" {
			rightTable = AsHasManyCompositePk
		}
		on := ""

		on += fmt.Sprintf("%v.%v = %v.%v",
			midTable, "has_many_composite_pk_id",
			rightTable, "id",
		)

		if AsHasManyCompositePk == "" {
			query = query.Join(dbr.I(JHasManyCompositePkModel.Table()), on)
		} else {
			query = query.Join(dbr.I(JHasManyCompositePkModel.Table()).As(AsHasManyCompositePk), on)
		}
	}

	return query
}

// LeftJoinRelateds adds a LEFT JOIN to CompositePk.Relateds
func (c *jCompositePkSelectBuilder) LeftJoinRelateds(
	AsCompositePkrelatedsToHasManyCompositePkrelateds, AsHasManyCompositePk string,
) *jCompositePkSelectBuilder {

	query := c

	leftTable := JCompositePkModel.Table()
	if c.as != "" {
		leftTable = c.as
	}

	midTable := JCompositePkrelatedsToHasManyCompositePkrelatedsModel.Table()
	if AsCompositePkrelatedsToHasManyCompositePkrelateds != "" {
		midTable = AsCompositePkrelatedsToHasManyCompositePkrelateds
	}

	{
		on := ""

		on += fmt.Sprintf("%v.%v = %v.%v",
			midTable, "composite_pk_p",
			leftTable, "p",
		)

		on += fmt.Sprintf("%v.%v = %v.%v",
			midTable, "composite_pk_k",
			leftTable, "k",
		)

		if AsCompositePkrelatedsToHasManyCompositePkrelateds == "" {
			query = query.LeftJoin(dbr.I(JCompositePkrelatedsToHasManyCompositePkrelatedsModel.Table()), on)
		} else {
			query = query.LeftJoin(dbr.I(JCompositePkrelatedsToHasManyCompositePkrelatedsModel.Table()).As(AsCompositePkrelatedsToHasManyCompositePkrelateds), on)
		}
	}

	{
		rightTable := JHasManyCompositePkModel.Table()
		if AsHasManyCompositePk != "" {
			rightTable = AsHasManyCompositePk
		}
		on := ""

		on += fmt.Sprintf("%v.%v = %v.%v",
			midTable, "has_many_composite_pk_id",
			rightTable, "id",
		)

		if AsHasManyCompositePk == "" {
			query = query.LeftJoin(dbr.I(JHasManyCompositePkModel.Table()), on)
		} else {
			query = query.LeftJoin(dbr.I(JHasManyCompositePkModel.Table()).As(AsHasManyCompositePk), on)
		}
	}

	return query
}

// // RightJoinRelateds adds a RIGHT JOIN to CompositePk.Relateds
// func (c *jCompositePkSelectBuilder) RightJoinRelateds(
// 	AsCompositePkrelatedsToHasManyCompositePkrelateds, AsHasManyCompositePk string,
// ) *jCompositePkSelectBuilder {
//
// 	query := c
//
// 	leftTable := JCompositePkModel.Table()
// 	if c.as != "" {
// 		leftTable = c.as
// 	}
//
// 	midTable := JCompositePkrelatedsToHasManyCompositePkrelatedsModel.Table()
// 	if AsCompositePkrelatedsToHasManyCompositePkrelateds != "" {
// 		midTable = AsCompositePkrelatedsToHasManyCompositePkrelateds
// 	}
//
// 	{
// 		on := ""
//
// 		on += fmt.Sprintf("%v.%v = %v.%v",
// 			midTable, "composite_pk_p",
// 			leftTable, "p",
// 			)
//
// 		on += fmt.Sprintf("%v.%v = %v.%v",
// 			midTable, "composite_pk_k",
// 			leftTable, "k",
// 			)
//
//
// 		if AsCompositePkrelatedsToHasManyCompositePkrelateds == "" {
// 			query = query.RightJoin(dbr.I(JCompositePkrelatedsToHasManyCompositePkrelatedsModel.Table()), on)
// 		} else {
// 			query = query.RightJoin(dbr.I(JCompositePkrelatedsToHasManyCompositePkrelatedsModel.Table()).As(AsCompositePkrelatedsToHasManyCompositePkrelateds), on)
// 		}
// 	}
//
// 	{
// 		rightTable := JHasManyCompositePkModel.Table()
// 		if AsHasManyCompositePk != "" {
// 			rightTable = AsHasManyCompositePk
// 		}
// 		on := ""
//
// 		on += fmt.Sprintf("%v.%v = %v.%v",
// 			midTable, "has_many_composite_pk_id",
// 			rightTable, "id",
// 			)
//
//
// 		if AsHasManyCompositePk == "" {
// 			query = query.RightJoin(dbr.I(JHasManyCompositePkModel.Table()), on)
// 		} else {
// 			query = query.RightJoin(dbr.I(JHasManyCompositePkModel.Table()).As(AsHasManyCompositePk), on)
// 		}
// 	}
//
// 	return query
// }

// HasManyHasOneCompositePk returns a query builder to select HasManyHasOneCompositePk linked to this CompositePk
func (g *CompositePk) HasManyHasOneCompositePk(db dbr.SessionRunner,
	AsRelated, AsHasManyHasOneCompositePk string,
) *jHasOneCompositePkSelectBuilder {

	var query *jHasOneCompositePkSelectBuilder

	leftTable := JHasOneCompositePkModel.Table()
	if AsRelated != "" {
		leftTable = AsRelated
		query = JHasOneCompositePk(db).As(AsRelated).Select(leftTable + ".*")
	} else {
		query = JHasOneCompositePk(db).Select(leftTable + ".*")
	}

	rightTable := JCompositePkModel.Table()
	if AsHasManyHasOneCompositePk != "" {
		rightTable = AsHasManyHasOneCompositePk
	}

	on := ""

	on += fmt.Sprintf("%v.%v = %v.%v",
		leftTable, "related_p",
		rightTable, "p",
	)

	on += fmt.Sprintf("%v.%v = %v.%v",
		leftTable, "related_k",
		rightTable, "k",
	)

	if AsHasManyHasOneCompositePk == "" {
		return query.Join(dbr.I(JCompositePkModel.Table()), on)
	}
	return query.Join(dbr.I(JCompositePkModel.Table()).As(AsHasManyHasOneCompositePk), on)
}

// JoinHasManyHasOneCompositePk adds a JOIN to CompositePk.HasManyHasOneCompositePk
func (c *jCompositePkSelectBuilder) JoinHasManyHasOneCompositePk(
	AsHasManyHasOneCompositePk string,
) *jCompositePkSelectBuilder {
	dialect := runtime.GetDialect()
	on := ""
	localTable := dialect.QuoteIdent(JCompositePkModel.Table())
	if c.as != "" {
		localTable = dialect.QuoteIdent(c.as)
	}
	foreiTable := dialect.QuoteIdent(JHasOneCompositePkModel.Table())
	if AsHasManyHasOneCompositePk != "" {
		foreiTable = dialect.QuoteIdent(AsHasManyHasOneCompositePk)
	}

	on += fmt.Sprintf("%v.%v = %v.%v",
		localTable, dialect.QuoteIdent("p"),
		foreiTable, dialect.QuoteIdent("related_p"),
	)

	on += fmt.Sprintf("%v.%v = %v.%v",
		localTable, dialect.QuoteIdent("k"),
		foreiTable, dialect.QuoteIdent("related_k"),
	)

	if AsHasManyHasOneCompositePk == "" {
		return c.Join(dbr.I(JHasOneCompositePkModel.Table()), on)
	}
	return c.Join(dbr.I(JHasOneCompositePkModel.Table()).As(AsHasManyHasOneCompositePk), on)
}

// LeftJoinHasManyHasOneCompositePk adds a LEFT JOIN to CompositePk.HasManyHasOneCompositePk
func (c *jCompositePkSelectBuilder) LeftJoinHasManyHasOneCompositePk(
	AsHasManyHasOneCompositePk string,
) *jCompositePkSelectBuilder {
	dialect := runtime.GetDialect()
	on := ""
	localTable := dialect.QuoteIdent(JCompositePkModel.Table())
	if c.as != "" {
		localTable = dialect.QuoteIdent(c.as)
	}
	foreiTable := dialect.QuoteIdent(JHasOneCompositePkModel.Table())
	if AsHasManyHasOneCompositePk != "" {
		foreiTable = dialect.QuoteIdent(AsHasManyHasOneCompositePk)
	}

	on += fmt.Sprintf("%v.%v = %v.%v",
		localTable, dialect.QuoteIdent("p"),
		foreiTable, dialect.QuoteIdent("related_p"),
	)

	on += fmt.Sprintf("%v.%v = %v.%v",
		localTable, dialect.QuoteIdent("k"),
		foreiTable, dialect.QuoteIdent("related_k"),
	)

	if AsHasManyHasOneCompositePk == "" {
		return c.LeftJoin(dbr.I(JHasOneCompositePkModel.Table()), on)
	}
	return c.LeftJoin(dbr.I(JHasOneCompositePkModel.Table()).As(AsHasManyHasOneCompositePk), on)
}

// // RightJoinHasManyHasOneCompositePk adds a Right JOIN to CompositePk.HasManyHasOneCompositePk
// func (c *jCompositePkSelectBuilder) RightJoinHasManyHasOneCompositePk(
// 	AsHasManyHasOneCompositePk string,
// ) *jCompositePkSelectBuilder {
// 	dialect := runtime.GetDialect()
// 	on := ""
// 	localTable := dialect.QuoteIdent(JCompositePkModel.Table())
// 	if c.as != "" {
// 		localTable = dialect.QuoteIdent(c.as)
// 	}
// 	foreiTable := dialect.QuoteIdent(JHasOneCompositePkModel.Table())
// 	if AsHasManyHasOneCompositePk != "" {
// 		foreiTable = dialect.QuoteIdent(AsHasManyHasOneCompositePk)
// 	}
//
// 	on += fmt.Sprintf("%v.%v = %v.%v",
// 		localTable, dialect.QuoteIdent("p"),
// 		foreiTable, dialect.QuoteIdent("related_p"),
// 	)
//
// 	on += fmt.Sprintf("%v.%v = %v.%v",
// 		localTable, dialect.QuoteIdent("k"),
// 		foreiTable, dialect.QuoteIdent("related_k"),
// 	)
//
// 	if AsHasManyHasOneCompositePk == "" {
// 		return c.RightJoin(dbr.I(JHasOneCompositePkModel.Table()), on)
// 	}
// 	return c.RightJoin(dbr.I(JHasOneCompositePkModel.Table()).As(AsHasManyHasOneCompositePk), on)
// }

// JHasOneCompositePkSetup helps to create/drop the schema
func JHasOneCompositePkSetup() runtime.Setuper {
	driver := runtime.GetCurrentDriver()

	var create string
	var drop string

	if driver == drivers.Sqlite {
		create = `CREATE TABLE IF NOT EXISTS has_one_composite_pk (
id INTEGER PRIMARY KEY AUTOINCREMENT,
x TEXT,
related_p TEXT NULL,
related_k TEXT NULL

)`
	} else if driver == drivers.Mysql {
		create = `CREATE TABLE IF NOT EXISTS has_one_composite_pk (
id INTEGER NOT NULL AUTO_INCREMENT,
x TEXT,
related_p TEXT NULL,
related_k TEXT NULL,
PRIMARY KEY (id) 

)`
	} else if driver == drivers.Pgsql {
		create = `CREATE TABLE IF NOT EXISTS has_one_composite_pk (
id SERIAL PRIMARY KEY,
x TEXT,
related_p TEXT NULL,
related_k TEXT NULL

)`
	}

	if driver == drivers.Sqlite {
		drop = `DROP TABLE IF EXISTS has_one_composite_pk`
	} else if driver == drivers.Mysql {
		drop = `DROP TABLE IF EXISTS has_one_composite_pk`
	} else if driver == drivers.Pgsql {
		drop = `DROP TABLE IF EXISTS has_one_composite_pk`
	}

	var indexes []string

	if driver == drivers.Sqlite {

	} else if driver == drivers.Mysql {

	} else if driver == drivers.Pgsql {

	}

	return runtime.Table{
		Name:       `has_one_composite_pk`,
		CreateStmt: create,
		DropStmt:   drop,
		View:       !true,
		Indexes:    indexes,
	}
}

// jHasOneCompositePkModel provides helper to work with HasOneCompositePk data provider
type jHasOneCompositePkModel struct {
	as string

	ID builder.ValuePropertyMeta

	X builder.ValuePropertyMeta

	RelatedP builder.ValuePropertyMeta

	RelatedK builder.ValuePropertyMeta

	Related builder.RelPropertyMeta
}

// Eq provided items.
func (j jHasOneCompositePkModel) Eq(s ...*HasOneCompositePk) dbr.Builder {
	ors := []dbr.Builder{}
	for _, t := range s {
		ors = append(ors, dbr.And(

			JHasOneCompositePkModel.ID.Eq(t.ID),
		))
	}
	return dbr.Or(ors...)
}

// In provided items.
func (j jHasOneCompositePkModel) In(s ...*HasOneCompositePk) dbr.Builder {
	ors := []dbr.Builder{}
	for _, t := range s {
		ors = append(ors, dbr.And(

			JHasOneCompositePkModel.ID.Eq(t.ID),
		))
	}
	return dbr.Or(ors...)
}

// As returns a copy with an alias.
func (j jHasOneCompositePkModel) As(as string) jHasOneCompositePkModel {
	j.as = as

	j.ID.TableAlias = as

	j.X.TableAlias = as

	j.RelatedP.TableAlias = as

	j.RelatedK.TableAlias = as

	// j.Related.TableAlias = as

	return j
}

// Table returns the sql table name
func (j jHasOneCompositePkModel) Table() string {
	return "has_one_composite_pk"
}

// Alias returns the current alias
func (j jHasOneCompositePkModel) Alias() string {
	if j.as == "" {
		return j.Table()
	}
	return j.as
}

// Properties returns a map of property name => meta
func (j jHasOneCompositePkModel) Properties() map[string]builder.MetaProvider {
	ret := map[string]builder.MetaProvider{}

	ret["ID"] = j.ID

	ret["X"] = j.X

	ret["RelatedP"] = j.RelatedP

	ret["RelatedK"] = j.RelatedK

	ret["Related"] = j.Related

	return ret
}

// Fields returns given sql fields with appropriate aliasing.
func (j jHasOneCompositePkModel) Fields(ins ...string) []string {
	dialect := runtime.GetDialect()
	if len(ins) == 0 {
		ins = append(ins, "*")
	}
	for i, in := range ins {
		if j.as != "" {
			if in == "*" {
				ins[i] = fmt.Sprintf("%v.%v", dialect.QuoteIdent(j.as), in)
			} else {
				ins[i] = fmt.Sprintf("%v.%v", dialect.QuoteIdent(j.as), dialect.QuoteIdent(in))
			}
		}
	}
	return ins
}

// JHasOneCompositePkModel provides helper to work with HasOneCompositePk data provider
var JHasOneCompositePkModel = jHasOneCompositePkModel{

	ID: builder.NewValueMeta(
		`id`, `INTEGER`,
		`ID`, `int64`,
		true, true,
	),

	X: builder.NewValueMeta(
		`x`, `TEXT`,
		`X`, `string`,
		false, false,
	),

	RelatedP: builder.NewValueMeta(
		`related_p`, `TEXT`,
		`RelatedP`, `*string`,
		false, false,
	),

	RelatedK: builder.NewValueMeta(
		`related_k`, `TEXT`,
		`RelatedK`, `*string`,
		false, false,
	),

	Related: builder.NewRelMeta(
		`related`, `*CompositePk`,
		`CompositePk`, ``, ``,
		`has_one`,
	),
}

type jHasOneCompositePkDeleteBuilder struct {
	*builder.DeleteBuilder
}

// //Build builds the sql string into given buffer using current dialect
// func (c *jHasOneCompositePkDeleteBuilder) Build(b dbr.Buffer) error {
// 	return c.DeleteBuilder.Build(runtime.GetDialect(), b)
// }
// //String returns the sql string for current dialect. It returns empty string if the build returns an error.
// func (c *jHasOneCompositePkDeleteBuilder) String() string {
// 	b := dbr.NewBuffer()
// 	if err := c.Build(b); err != nil {
// 		return ""
// 	}
// 	return b.String()
// }
//Where returns a jHasOneCompositePkDeleteBuilder instead of builder.DeleteBuilder.
func (c *jHasOneCompositePkDeleteBuilder) Where(query interface{}, value ...interface{}) *jHasOneCompositePkDeleteBuilder {
	c.DeleteBuilder.Where(query, value...)
	return c
}

type jHasOneCompositePkSelectBuilder struct {
	as string
	*builder.SelectBuilder
}

// //Build builds the sql string using current dialect into given bufer
// func (c *jHasOneCompositePkSelectBuilder) Build(b dbr.Buffer) error {
// 	return c.SelectBuilder.Build(runtime.GetDialect(), b)
// }
// //String returns the sql string for current dialect. It returns empty string if the build returns an error.
// func (c *jHasOneCompositePkSelectBuilder) String() string {
// 	b := dbr.NewBuffer()
// 	if err := c.Build(b); err != nil {
// 		return ""
// 	}
// 	return b.String()
// }

//Read evaluates current select query and load the results into a HasOneCompositePk
func (c *jHasOneCompositePkSelectBuilder) Read() (*HasOneCompositePk, error) {
	var one HasOneCompositePk
	err := c.LoadStruct(&one)
	return &one, err
}

//ReadAll evaluates current select query and load the results into a slice of HasOneCompositePk
func (c *jHasOneCompositePkSelectBuilder) ReadAll() ([]*HasOneCompositePk, error) {
	var all []*HasOneCompositePk
	_, err := c.LoadStructs(&all)
	return all, err
}

//Where returns a jHasOneCompositePkSelectBuilder instead of builder.SelectBuilder.
func (c *jHasOneCompositePkSelectBuilder) Where(query interface{}, value ...interface{}) *jHasOneCompositePkSelectBuilder {
	c.SelectBuilder.Where(query, value...)
	return c
}

//GroupBy returns a jHasOneCompositePkSelectBuilder instead of builder.SelectBuilder.
func (c *jHasOneCompositePkSelectBuilder) GroupBy(col ...string) *jHasOneCompositePkSelectBuilder {
	c.SelectBuilder.GroupBy(col...)
	return c
}

//Having returns a jHasOneCompositePkSelectBuilder instead of builder.SelectBuilder.
func (c *jHasOneCompositePkSelectBuilder) Having(query interface{}, value ...interface{}) *jHasOneCompositePkSelectBuilder {
	c.SelectBuilder.Having(query, value...)
	return c
}

//Limit returns a jHasOneCompositePkSelectBuilder instead of builder.SelectBuilder.
func (c *jHasOneCompositePkSelectBuilder) Limit(n uint64) *jHasOneCompositePkSelectBuilder {
	c.SelectBuilder.Limit(n)
	return c
}

//Offset returns a jHasOneCompositePkSelectBuilder instead of builder.SelectBuilder.
func (c *jHasOneCompositePkSelectBuilder) Offset(n uint64) *jHasOneCompositePkSelectBuilder {
	c.SelectBuilder.Offset(n)
	return c
}

//OrderAsc returns a jHasOneCompositePkSelectBuilder instead of builder.SelectBuilder.
func (c *jHasOneCompositePkSelectBuilder) OrderAsc(col string) *jHasOneCompositePkSelectBuilder {
	c.SelectBuilder.OrderAsc(col)
	return c
}

//OrderDesc returns a jHasOneCompositePkSelectBuilder instead of builder.SelectBuilder.
func (c *jHasOneCompositePkSelectBuilder) OrderDesc(col string) *jHasOneCompositePkSelectBuilder {
	c.SelectBuilder.OrderDesc(col)
	return c
}

//OrderDir returns a jHasOneCompositePkSelectBuilder instead of builder.SelectBuilder.
func (c *jHasOneCompositePkSelectBuilder) OrderDir(col string, isAsc bool) *jHasOneCompositePkSelectBuilder {
	c.SelectBuilder.OrderDir(col, isAsc)
	return c
}

//OrderBy returns a jHasOneCompositePkSelectBuilder instead of builder.SelectBuilder.
func (c *jHasOneCompositePkSelectBuilder) OrderBy(col string) *jHasOneCompositePkSelectBuilder {
	c.SelectBuilder.OrderBy(col)
	return c
}

//Paginate returns a jHasOneCompositePkSelectBuilder instead of builder.SelectBuilder.
func (c *jHasOneCompositePkSelectBuilder) Paginate(page, perPage uint64) *jHasOneCompositePkSelectBuilder {
	c.SelectBuilder.Paginate(page, perPage)
	return c
}

//Join returns a jHasOneCompositePkSelectBuilder instead of builder.SelectBuilder.
func (c *jHasOneCompositePkSelectBuilder) Join(table, on interface{}) *jHasOneCompositePkSelectBuilder {
	c.SelectBuilder.Join(table, on)
	return c
}

//LeftJoin returns a jHasOneCompositePkSelectBuilder instead of builder.SelectBuilder.
func (c *jHasOneCompositePkSelectBuilder) LeftJoin(table, on interface{}) *jHasOneCompositePkSelectBuilder {
	c.SelectBuilder.LeftJoin(table, on)
	return c
}

//RightJoin returns a jHasOneCompositePkSelectBuilder instead of builder.SelectBuilder.
func (c *jHasOneCompositePkSelectBuilder) RightJoin(table, on interface{}) *jHasOneCompositePkSelectBuilder {
	c.SelectBuilder.RightJoin(table, on)
	return c
}

//FullJoin returns a jHasOneCompositePkSelectBuilder instead of builder.SelectBuilder.
func (c *jHasOneCompositePkSelectBuilder) FullJoin(table, on interface{}) *jHasOneCompositePkSelectBuilder {
	c.SelectBuilder.FullJoin(table, on)
	return c
}

//Distinct returns a jHasOneCompositePkSelectBuilder instead of builder.SelectBuilder.
func (c *jHasOneCompositePkSelectBuilder) Distinct() *jHasOneCompositePkSelectBuilder {
	c.SelectBuilder.Distinct()
	return c
}

// JHasOneCompositePk provides a basic querier
func JHasOneCompositePk(db dbr.SessionRunner) jHasOneCompositePkQuerier {
	return jHasOneCompositePkQuerier{
		db: db,
	}
}

type jHasOneCompositePkQuerier struct {
	db dbr.SessionRunner
	as string
}

//As set alias prior building.
func (c jHasOneCompositePkQuerier) As(as string) jHasOneCompositePkQuerier {
	c.as = as
	return c
}

//Model returns a model with appropriate aliasing.
func (c jHasOneCompositePkQuerier) Model() jHasOneCompositePkModel {
	return JHasOneCompositePkModel.As(c.as)
}

//Select returns a HasOneCompositePk Select Builder.
func (c jHasOneCompositePkQuerier) Select(what ...string) *jHasOneCompositePkSelectBuilder {
	m := c.Model()
	dialect := runtime.GetDialect()
	from := dialect.QuoteIdent(m.Table())
	if m.Alias() != "" && m.Alias() != m.Table() {
		from = fmt.Sprintf("%v as %v", from, dialect.QuoteIdent(m.Alias()))
	}
	if len(what) == 0 {
		alias := m.Table()
		if m.Alias() != "" && m.Alias() != m.Table() {
			alias = m.Alias()
		}
		what = m.Fields(alias + ".*")
	}
	return &jHasOneCompositePkSelectBuilder{
		as: c.as,
		SelectBuilder: &builder.SelectBuilder{
			SelectBuilder: c.db.Select(what...).From(from),
		},
	}
}

//Where returns a HasOneCompositePk Select Builder.
func (c jHasOneCompositePkQuerier) Where(query interface{}, value ...interface{}) *jHasOneCompositePkSelectBuilder {
	return c.Select().Where(query, value...)
}

//Count returns a HasOneCompositePk Select Builder to count given expressions.
func (c jHasOneCompositePkQuerier) Count(what ...string) *jHasOneCompositePkSelectBuilder {
	if len(what) == 0 {
		what = append(what, "*")
	}
	return c.Select("COUNT(" + strings.Join(what, ",") + ")")
}

// Insert a new HasOneCompositePk, if it has autoincrement primary key, the value will be set.
// It stops on first error.
func (c jHasOneCompositePkQuerier) Insert(items ...*HasOneCompositePk) (sql.Result, error) {
	var res sql.Result
	var err error
	for _, data := range items {

		query := c.db.InsertInto(JHasOneCompositePkModel.Table()).Columns(

			`x`,

			`related_p`,

			`related_k`,
		).Record(data)
		if runtime.Runs(drivers.Pgsql) {

			query = query.Returning(

				`id`,
			)

			var auto0 int64

			err = query.Load(

				&auto0,
			)

			data.ID = auto0

		} else {
			res, err = query.Exec()

			if err == nil {
				id, err2 := res.LastInsertId()
				if err2 != nil {
					return res, err2
				}
				data.ID = id
			}

		}
		if err != nil {
			return res, err
		}
	}
	return res, err
}

// InsertBulk inserts multiple items into the database.
// It does not post update any auto increment field.
// It builds an insert query of multiple rows and send it on the underlying connection.
func (c jHasOneCompositePkQuerier) InsertBulk(items ...*HasOneCompositePk) error {
	panic("todo")
}

// Update a HasOneCompositePk. It stops on first error.
func (c jHasOneCompositePkQuerier) Update(items ...*HasOneCompositePk) (sql.Result, error) {
	var res sql.Result
	var err error
	for _, data := range items {

		query := c.db.Update(JHasOneCompositePkModel.Table())

		query = query.Set(`x`, data.X)

		query = query.Set(`related_p`, data.RelatedP)

		query = query.Set(`related_k`, data.RelatedK)

		query = query.Where("id = ?", data.ID)

		res, err = query.Exec()

		if err != nil {
			return res, err
		}
	}
	return res, err
}

// MustUpdate a HasOneCompositePk. It stops on first error. It errors if an update query does not affect row.
func (c jHasOneCompositePkQuerier) MustUpdate(items ...*HasOneCompositePk) (sql.Result, error) {
	var res sql.Result
	var err error
	for _, data := range items {

		query := c.db.Update(JHasOneCompositePkModel.Table())

		query = query.Set(`x`, data.X)

		query = query.Set(`related_p`, data.RelatedP)

		query = query.Set(`related_k`, data.RelatedK)

		query = query.Where("id = ?", data.ID)

		res, err = query.Exec()

		if err == nil {
			if n, _ := res.RowsAffected(); n == 0 {
				x := &builder.UpdateBuilder{UpdateBuilder: query}
				err = runtime.NewNoRowsAffected(x.String())
			}
		}

		if err != nil {
			return res, err
		}
	}
	return res, err
}

// UpdateBulk updates multiple items into the database.
// It builds an update query of multiple rows and send it on the underlying connection.
func (c jHasOneCompositePkQuerier) UpdateBulk(items ...*HasOneCompositePk) error {
	panic("todo")
}

//Delete returns a delete builder
func (c jHasOneCompositePkQuerier) Delete() *jHasOneCompositePkDeleteBuilder {
	return &jHasOneCompositePkDeleteBuilder{
		&builder.DeleteBuilder{
			DeleteBuilder: c.db.DeleteFrom(JHasOneCompositePkModel.Table()),
		},
	}
}

// MustDelete requires the query to affeect rows.
func (c jHasOneCompositePkQuerier) MustDelete() *jHasOneCompositePkDeleteBuilder {
	ret := &jHasOneCompositePkDeleteBuilder{
		&builder.DeleteBuilder{
			DeleteBuilder: c.db.DeleteFrom(JHasOneCompositePkModel.Table()),
		},
	}
	ret.MustDelete()
	return ret
}

//DeleteByPk deletes one HasOneCompositePk by its PKs
func (c jHasOneCompositePkQuerier) DeleteByPk(ID int64) error {
	_, err := c.Delete().Where(

		JHasOneCompositePkModel.ID.Eq(ID),
	).Exec()
	return err
}

// DeleteAll given HasOneCompositePk
func (c jHasOneCompositePkQuerier) DeleteAll(items ...*HasOneCompositePk) (sql.Result, error) {
	q := c.Delete().Where(
		JHasOneCompositePkModel.In(items...),
	)
	return q.Exec()
}

// MustDeleteAll given HasOneCompositePk
func (c jHasOneCompositePkQuerier) MustDeleteAll(items ...*HasOneCompositePk) (sql.Result, error) {
	var res sql.Result
	var err error
	for _, d := range items {
		res, err = c.DeleteAll(d)
		if err != nil {
			return res, err
		}
		if n, e := res.RowsAffected(); e != nil {
			return res, e
		} else if n == 0 {
			q := c.Delete().Where(
				JHasOneCompositePkModel.In(items...),
			)
			err = runtime.NewNoRowsAffected(q.String())
			return res, err
		}
	}
	return res, err
}

//Find one HasOneCompositePk using its PKs
func (c jHasOneCompositePkQuerier) Find(ID int64) (*HasOneCompositePk, error) {
	return c.Select().Where(

		JHasOneCompositePkModel.ID.Eq(ID),
	).Read()
}

// JoinRelated adds a JOIN to HasOneCompositePk.Related
func (c *jHasOneCompositePkSelectBuilder) JoinRelated(
	AsRelated string,
) *jHasOneCompositePkSelectBuilder {
	dialect := runtime.GetDialect()
	on := ""
	localTable := dialect.QuoteIdent(JHasOneCompositePkModel.Table())
	if c.as != "" {
		localTable = dialect.QuoteIdent(c.as)
	}
	foreiTable := dialect.QuoteIdent(JCompositePkModel.Table())
	if AsRelated != "" {
		foreiTable = dialect.QuoteIdent(AsRelated)
	}

	on += fmt.Sprintf("%v.%v = %v.%v",
		localTable, dialect.QuoteIdent("related_p"),
		foreiTable, dialect.QuoteIdent("p"),
	)

	on += fmt.Sprintf("%v.%v = %v.%v",
		localTable, dialect.QuoteIdent("related_k"),
		foreiTable, dialect.QuoteIdent("k"),
	)

	if AsRelated == "" {
		return c.Join(dbr.I(JCompositePkModel.Table()), on)
	}
	return c.Join(dbr.I(JCompositePkModel.Table()).As(AsRelated), on)
}

// LeftJoinRelated adds a LEFT JOIN to HasOneCompositePk.Related
func (c *jHasOneCompositePkSelectBuilder) LeftJoinRelated(
	AsRelated string,
) *jHasOneCompositePkSelectBuilder {
	dialect := runtime.GetDialect()
	on := ""
	localTable := JHasOneCompositePkModel.Table()
	if c.as != "" {
		localTable = dialect.QuoteIdent(c.as)
	}
	foreiTable := dialect.QuoteIdent(JCompositePkModel.Table())
	if AsRelated != "" {
		foreiTable = dialect.QuoteIdent(AsRelated)
	}

	on += fmt.Sprintf("%v.%v = %v.%v",
		localTable, dialect.QuoteIdent("related_p"),
		foreiTable, dialect.QuoteIdent("p"),
	)

	on += fmt.Sprintf("%v.%v = %v.%v",
		localTable, dialect.QuoteIdent("related_k"),
		foreiTable, dialect.QuoteIdent("k"),
	)

	if AsRelated == "" {
		return c.LeftJoin(dbr.I(JCompositePkModel.Table()), on)
	}
	return c.LeftJoin(dbr.I(JCompositePkModel.Table()).As(AsRelated), on)
}

// // RightJoinRelated adds a RIGHT JOIN to HasOneCompositePk.Related
// func (c *jHasOneCompositePkSelectBuilder) RightJoinRelated(
// 	AsRelated string,
// ) *jHasOneCompositePkSelectBuilder {
// 	dialect := runtime.GetDialect()
// 	on := ""
// 	localTable := dialect.QuoteIdent(JHasOneCompositePkModel.Table())
// 	if c.as != "" {
// 		localTable = dialect.QuoteIdent(c.as)
// 	}
// 	foreiTable := dialect.QuoteIdent(JCompositePkModel.Table())
// 	if AsRelated != "" {
// 		foreiTable = dialect.QuoteIdent(AsRelated)
// 	}
//
// 	on += fmt.Sprintf("%v.%v = %v.%v",
// 		localTable, dialect.QuoteIdent("related_p"),
// 		foreiTable, dialect.QuoteIdent("p"),
// 	)
//
// 	on += fmt.Sprintf("%v.%v = %v.%v",
// 		localTable, dialect.QuoteIdent("related_k"),
// 		foreiTable, dialect.QuoteIdent("k"),
// 	)
//
// 	if AsRelated == "" {
// 		return c.RightJoin(dbr.I(JCompositePkModel.Table()), on)
// 	}
// 	return c.RightJoin(dbr.I(JCompositePkModel.Table()).As(AsRelated), on)
// }

// FullJoinRelated adds a FULL JOIN to HasOneCompositePk.Related
func (c *jHasOneCompositePkSelectBuilder) FullJoinRelated(
	AsRelated string,
) *jHasOneCompositePkSelectBuilder {
	dialect := runtime.GetDialect()
	on := ""
	localTable := dialect.QuoteIdent(JHasOneCompositePkModel.Table())
	if c.as != "" {
		localTable = dialect.QuoteIdent(c.as)
	}
	foreiTable := dialect.QuoteIdent(JCompositePkModel.Table())
	if AsRelated != "" {
		foreiTable = dialect.QuoteIdent(AsRelated)
	}

	on += fmt.Sprintf("%v.%v = %v.%v",
		localTable, dialect.QuoteIdent("related_p"),
		foreiTable, dialect.QuoteIdent("p"),
	)

	on += fmt.Sprintf("%v.%v = %v.%v",
		localTable, dialect.QuoteIdent("related_k"),
		foreiTable, dialect.QuoteIdent("k"),
	)

	if AsRelated == "" {
		return c.FullJoin(dbr.I(JCompositePkModel.Table()), on)
	}
	return c.FullJoin(dbr.I(JCompositePkModel.Table()).As(AsRelated), on)
}

// Related reads associated object
func (g *HasOneCompositePk) Related(db dbr.SessionRunner) (*CompositePk, error) {
	q := JCompositePk(db).Select()
	q = q.Where(

		JCompositePkModel.P.Eq(g.RelatedP),

		JCompositePkModel.K.Eq(g.RelatedK),
	)
	return q.Read()
}

// SetRelated copies pk values to this object properties
func (g *HasOneCompositePk) SetRelated(o *CompositePk) *HasOneCompositePk {

	if o == nil {
		g.RelatedP = nil
	} else {

		g.RelatedP = &o.P

	}

	if o == nil {
		g.RelatedK = nil
	} else {

		g.RelatedK = &o.K

	}

	return g
}

// UnsetRelated set defaults values to this object properties
func (g *HasOneCompositePk) UnsetRelated() *HasOneCompositePk {

	var def0 *string

	var def1 *string

	g.RelatedP = def0

	g.RelatedK = def1

	g.related = nil

	return g
}

// JHasManyCompositePkSetup helps to create/drop the schema
func JHasManyCompositePkSetup() runtime.Setuper {
	driver := runtime.GetCurrentDriver()

	var create string
	var drop string

	if driver == drivers.Sqlite {
		create = `CREATE TABLE IF NOT EXISTS has_many_composite_pk (
id INTEGER PRIMARY KEY AUTOINCREMENT,
x TEXT

)`
	} else if driver == drivers.Mysql {
		create = `CREATE TABLE IF NOT EXISTS has_many_composite_pk (
id INTEGER NOT NULL AUTO_INCREMENT,
x TEXT,
PRIMARY KEY (id) 

)`
	} else if driver == drivers.Pgsql {
		create = `CREATE TABLE IF NOT EXISTS has_many_composite_pk (
id SERIAL PRIMARY KEY,
x TEXT

)`
	}

	if driver == drivers.Sqlite {
		drop = `DROP TABLE IF EXISTS has_many_composite_pk`
	} else if driver == drivers.Mysql {
		drop = `DROP TABLE IF EXISTS has_many_composite_pk`
	} else if driver == drivers.Pgsql {
		drop = `DROP TABLE IF EXISTS has_many_composite_pk`
	}

	var indexes []string

	if driver == drivers.Sqlite {

	} else if driver == drivers.Mysql {

	} else if driver == drivers.Pgsql {

	}

	return runtime.Table{
		Name:       `has_many_composite_pk`,
		CreateStmt: create,
		DropStmt:   drop,
		View:       !true,
		Indexes:    indexes,
	}
}

// jHasManyCompositePkModel provides helper to work with HasManyCompositePk data provider
type jHasManyCompositePkModel struct {
	as string

	ID builder.ValuePropertyMeta

	X builder.ValuePropertyMeta

	Relateds builder.RelPropertyMeta
}

// Eq provided items.
func (j jHasManyCompositePkModel) Eq(s ...*HasManyCompositePk) dbr.Builder {
	ors := []dbr.Builder{}
	for _, t := range s {
		ors = append(ors, dbr.And(

			JHasManyCompositePkModel.ID.Eq(t.ID),
		))
	}
	return dbr.Or(ors...)
}

// In provided items.
func (j jHasManyCompositePkModel) In(s ...*HasManyCompositePk) dbr.Builder {
	ors := []dbr.Builder{}
	for _, t := range s {
		ors = append(ors, dbr.And(

			JHasManyCompositePkModel.ID.Eq(t.ID),
		))
	}
	return dbr.Or(ors...)
}

// As returns a copy with an alias.
func (j jHasManyCompositePkModel) As(as string) jHasManyCompositePkModel {
	j.as = as

	j.ID.TableAlias = as

	j.X.TableAlias = as

	// j.Relateds.TableAlias = as

	return j
}

// Table returns the sql table name
func (j jHasManyCompositePkModel) Table() string {
	return "has_many_composite_pk"
}

// Alias returns the current alias
func (j jHasManyCompositePkModel) Alias() string {
	if j.as == "" {
		return j.Table()
	}
	return j.as
}

// Properties returns a map of property name => meta
func (j jHasManyCompositePkModel) Properties() map[string]builder.MetaProvider {
	ret := map[string]builder.MetaProvider{}

	ret["ID"] = j.ID

	ret["X"] = j.X

	ret["Relateds"] = j.Relateds

	return ret
}

// Fields returns given sql fields with appropriate aliasing.
func (j jHasManyCompositePkModel) Fields(ins ...string) []string {
	dialect := runtime.GetDialect()
	if len(ins) == 0 {
		ins = append(ins, "*")
	}
	for i, in := range ins {
		if j.as != "" {
			if in == "*" {
				ins[i] = fmt.Sprintf("%v.%v", dialect.QuoteIdent(j.as), in)
			} else {
				ins[i] = fmt.Sprintf("%v.%v", dialect.QuoteIdent(j.as), dialect.QuoteIdent(in))
			}
		}
	}
	return ins
}

// JHasManyCompositePkModel provides helper to work with HasManyCompositePk data provider
var JHasManyCompositePkModel = jHasManyCompositePkModel{

	ID: builder.NewValueMeta(
		`id`, `INTEGER`,
		`ID`, `int64`,
		true, true,
	),

	X: builder.NewValueMeta(
		`x`, `TEXT`,
		`X`, `string`,
		false, false,
	),

	Relateds: builder.NewRelMeta(
		`relateds`, `[]*CompositePk`,
		``, `CompositePk.relateds`, ``,
		`has_many2many`,
	),
}

type jHasManyCompositePkDeleteBuilder struct {
	*builder.DeleteBuilder
}

// //Build builds the sql string into given buffer using current dialect
// func (c *jHasManyCompositePkDeleteBuilder) Build(b dbr.Buffer) error {
// 	return c.DeleteBuilder.Build(runtime.GetDialect(), b)
// }
// //String returns the sql string for current dialect. It returns empty string if the build returns an error.
// func (c *jHasManyCompositePkDeleteBuilder) String() string {
// 	b := dbr.NewBuffer()
// 	if err := c.Build(b); err != nil {
// 		return ""
// 	}
// 	return b.String()
// }
//Where returns a jHasManyCompositePkDeleteBuilder instead of builder.DeleteBuilder.
func (c *jHasManyCompositePkDeleteBuilder) Where(query interface{}, value ...interface{}) *jHasManyCompositePkDeleteBuilder {
	c.DeleteBuilder.Where(query, value...)
	return c
}

type jHasManyCompositePkSelectBuilder struct {
	as string
	*builder.SelectBuilder
}

// //Build builds the sql string using current dialect into given bufer
// func (c *jHasManyCompositePkSelectBuilder) Build(b dbr.Buffer) error {
// 	return c.SelectBuilder.Build(runtime.GetDialect(), b)
// }
// //String returns the sql string for current dialect. It returns empty string if the build returns an error.
// func (c *jHasManyCompositePkSelectBuilder) String() string {
// 	b := dbr.NewBuffer()
// 	if err := c.Build(b); err != nil {
// 		return ""
// 	}
// 	return b.String()
// }

//Read evaluates current select query and load the results into a HasManyCompositePk
func (c *jHasManyCompositePkSelectBuilder) Read() (*HasManyCompositePk, error) {
	var one HasManyCompositePk
	err := c.LoadStruct(&one)
	return &one, err
}

//ReadAll evaluates current select query and load the results into a slice of HasManyCompositePk
func (c *jHasManyCompositePkSelectBuilder) ReadAll() ([]*HasManyCompositePk, error) {
	var all []*HasManyCompositePk
	_, err := c.LoadStructs(&all)
	return all, err
}

//Where returns a jHasManyCompositePkSelectBuilder instead of builder.SelectBuilder.
func (c *jHasManyCompositePkSelectBuilder) Where(query interface{}, value ...interface{}) *jHasManyCompositePkSelectBuilder {
	c.SelectBuilder.Where(query, value...)
	return c
}

//GroupBy returns a jHasManyCompositePkSelectBuilder instead of builder.SelectBuilder.
func (c *jHasManyCompositePkSelectBuilder) GroupBy(col ...string) *jHasManyCompositePkSelectBuilder {
	c.SelectBuilder.GroupBy(col...)
	return c
}

//Having returns a jHasManyCompositePkSelectBuilder instead of builder.SelectBuilder.
func (c *jHasManyCompositePkSelectBuilder) Having(query interface{}, value ...interface{}) *jHasManyCompositePkSelectBuilder {
	c.SelectBuilder.Having(query, value...)
	return c
}

//Limit returns a jHasManyCompositePkSelectBuilder instead of builder.SelectBuilder.
func (c *jHasManyCompositePkSelectBuilder) Limit(n uint64) *jHasManyCompositePkSelectBuilder {
	c.SelectBuilder.Limit(n)
	return c
}

//Offset returns a jHasManyCompositePkSelectBuilder instead of builder.SelectBuilder.
func (c *jHasManyCompositePkSelectBuilder) Offset(n uint64) *jHasManyCompositePkSelectBuilder {
	c.SelectBuilder.Offset(n)
	return c
}

//OrderAsc returns a jHasManyCompositePkSelectBuilder instead of builder.SelectBuilder.
func (c *jHasManyCompositePkSelectBuilder) OrderAsc(col string) *jHasManyCompositePkSelectBuilder {
	c.SelectBuilder.OrderAsc(col)
	return c
}

//OrderDesc returns a jHasManyCompositePkSelectBuilder instead of builder.SelectBuilder.
func (c *jHasManyCompositePkSelectBuilder) OrderDesc(col string) *jHasManyCompositePkSelectBuilder {
	c.SelectBuilder.OrderDesc(col)
	return c
}

//OrderDir returns a jHasManyCompositePkSelectBuilder instead of builder.SelectBuilder.
func (c *jHasManyCompositePkSelectBuilder) OrderDir(col string, isAsc bool) *jHasManyCompositePkSelectBuilder {
	c.SelectBuilder.OrderDir(col, isAsc)
	return c
}

//OrderBy returns a jHasManyCompositePkSelectBuilder instead of builder.SelectBuilder.
func (c *jHasManyCompositePkSelectBuilder) OrderBy(col string) *jHasManyCompositePkSelectBuilder {
	c.SelectBuilder.OrderBy(col)
	return c
}

//Paginate returns a jHasManyCompositePkSelectBuilder instead of builder.SelectBuilder.
func (c *jHasManyCompositePkSelectBuilder) Paginate(page, perPage uint64) *jHasManyCompositePkSelectBuilder {
	c.SelectBuilder.Paginate(page, perPage)
	return c
}

//Join returns a jHasManyCompositePkSelectBuilder instead of builder.SelectBuilder.
func (c *jHasManyCompositePkSelectBuilder) Join(table, on interface{}) *jHasManyCompositePkSelectBuilder {
	c.SelectBuilder.Join(table, on)
	return c
}

//LeftJoin returns a jHasManyCompositePkSelectBuilder instead of builder.SelectBuilder.
func (c *jHasManyCompositePkSelectBuilder) LeftJoin(table, on interface{}) *jHasManyCompositePkSelectBuilder {
	c.SelectBuilder.LeftJoin(table, on)
	return c
}

//RightJoin returns a jHasManyCompositePkSelectBuilder instead of builder.SelectBuilder.
func (c *jHasManyCompositePkSelectBuilder) RightJoin(table, on interface{}) *jHasManyCompositePkSelectBuilder {
	c.SelectBuilder.RightJoin(table, on)
	return c
}

//FullJoin returns a jHasManyCompositePkSelectBuilder instead of builder.SelectBuilder.
func (c *jHasManyCompositePkSelectBuilder) FullJoin(table, on interface{}) *jHasManyCompositePkSelectBuilder {
	c.SelectBuilder.FullJoin(table, on)
	return c
}

//Distinct returns a jHasManyCompositePkSelectBuilder instead of builder.SelectBuilder.
func (c *jHasManyCompositePkSelectBuilder) Distinct() *jHasManyCompositePkSelectBuilder {
	c.SelectBuilder.Distinct()
	return c
}

// JHasManyCompositePk provides a basic querier
func JHasManyCompositePk(db dbr.SessionRunner) jHasManyCompositePkQuerier {
	return jHasManyCompositePkQuerier{
		db: db,
	}
}

type jHasManyCompositePkQuerier struct {
	db dbr.SessionRunner
	as string
}

//As set alias prior building.
func (c jHasManyCompositePkQuerier) As(as string) jHasManyCompositePkQuerier {
	c.as = as
	return c
}

//Model returns a model with appropriate aliasing.
func (c jHasManyCompositePkQuerier) Model() jHasManyCompositePkModel {
	return JHasManyCompositePkModel.As(c.as)
}

//Select returns a HasManyCompositePk Select Builder.
func (c jHasManyCompositePkQuerier) Select(what ...string) *jHasManyCompositePkSelectBuilder {
	m := c.Model()
	dialect := runtime.GetDialect()
	from := dialect.QuoteIdent(m.Table())
	if m.Alias() != "" && m.Alias() != m.Table() {
		from = fmt.Sprintf("%v as %v", from, dialect.QuoteIdent(m.Alias()))
	}
	if len(what) == 0 {
		alias := m.Table()
		if m.Alias() != "" && m.Alias() != m.Table() {
			alias = m.Alias()
		}
		what = m.Fields(alias + ".*")
	}
	return &jHasManyCompositePkSelectBuilder{
		as: c.as,
		SelectBuilder: &builder.SelectBuilder{
			SelectBuilder: c.db.Select(what...).From(from),
		},
	}
}

//Where returns a HasManyCompositePk Select Builder.
func (c jHasManyCompositePkQuerier) Where(query interface{}, value ...interface{}) *jHasManyCompositePkSelectBuilder {
	return c.Select().Where(query, value...)
}

//Count returns a HasManyCompositePk Select Builder to count given expressions.
func (c jHasManyCompositePkQuerier) Count(what ...string) *jHasManyCompositePkSelectBuilder {
	if len(what) == 0 {
		what = append(what, "*")
	}
	return c.Select("COUNT(" + strings.Join(what, ",") + ")")
}

// Insert a new HasManyCompositePk, if it has autoincrement primary key, the value will be set.
// It stops on first error.
func (c jHasManyCompositePkQuerier) Insert(items ...*HasManyCompositePk) (sql.Result, error) {
	var res sql.Result
	var err error
	for _, data := range items {

		query := c.db.InsertInto(JHasManyCompositePkModel.Table()).Columns(

			`x`,
		).Record(data)
		if runtime.Runs(drivers.Pgsql) {

			query = query.Returning(

				`id`,
			)

			var auto0 int64

			err = query.Load(

				&auto0,
			)

			data.ID = auto0

		} else {
			res, err = query.Exec()

			if err == nil {
				id, err2 := res.LastInsertId()
				if err2 != nil {
					return res, err2
				}
				data.ID = id
			}

		}
		if err != nil {
			return res, err
		}
	}
	return res, err
}

// InsertBulk inserts multiple items into the database.
// It does not post update any auto increment field.
// It builds an insert query of multiple rows and send it on the underlying connection.
func (c jHasManyCompositePkQuerier) InsertBulk(items ...*HasManyCompositePk) error {
	panic("todo")
}

// Update a HasManyCompositePk. It stops on first error.
func (c jHasManyCompositePkQuerier) Update(items ...*HasManyCompositePk) (sql.Result, error) {
	var res sql.Result
	var err error
	for _, data := range items {

		query := c.db.Update(JHasManyCompositePkModel.Table())

		query = query.Set(`x`, data.X)

		query = query.Where("id = ?", data.ID)

		res, err = query.Exec()

		if err != nil {
			return res, err
		}
	}
	return res, err
}

// MustUpdate a HasManyCompositePk. It stops on first error. It errors if an update query does not affect row.
func (c jHasManyCompositePkQuerier) MustUpdate(items ...*HasManyCompositePk) (sql.Result, error) {
	var res sql.Result
	var err error
	for _, data := range items {

		query := c.db.Update(JHasManyCompositePkModel.Table())

		query = query.Set(`x`, data.X)

		query = query.Where("id = ?", data.ID)

		res, err = query.Exec()

		if err == nil {
			if n, _ := res.RowsAffected(); n == 0 {
				x := &builder.UpdateBuilder{UpdateBuilder: query}
				err = runtime.NewNoRowsAffected(x.String())
			}
		}

		if err != nil {
			return res, err
		}
	}
	return res, err
}

// UpdateBulk updates multiple items into the database.
// It builds an update query of multiple rows and send it on the underlying connection.
func (c jHasManyCompositePkQuerier) UpdateBulk(items ...*HasManyCompositePk) error {
	panic("todo")
}

//Delete returns a delete builder
func (c jHasManyCompositePkQuerier) Delete() *jHasManyCompositePkDeleteBuilder {
	return &jHasManyCompositePkDeleteBuilder{
		&builder.DeleteBuilder{
			DeleteBuilder: c.db.DeleteFrom(JHasManyCompositePkModel.Table()),
		},
	}
}

// MustDelete requires the query to affeect rows.
func (c jHasManyCompositePkQuerier) MustDelete() *jHasManyCompositePkDeleteBuilder {
	ret := &jHasManyCompositePkDeleteBuilder{
		&builder.DeleteBuilder{
			DeleteBuilder: c.db.DeleteFrom(JHasManyCompositePkModel.Table()),
		},
	}
	ret.MustDelete()
	return ret
}

//DeleteByPk deletes one HasManyCompositePk by its PKs
func (c jHasManyCompositePkQuerier) DeleteByPk(ID int64) error {
	_, err := c.Delete().Where(

		JHasManyCompositePkModel.ID.Eq(ID),
	).Exec()
	return err
}

// DeleteAll given HasManyCompositePk
func (c jHasManyCompositePkQuerier) DeleteAll(items ...*HasManyCompositePk) (sql.Result, error) {
	q := c.Delete().Where(
		JHasManyCompositePkModel.In(items...),
	)
	return q.Exec()
}

// MustDeleteAll given HasManyCompositePk
func (c jHasManyCompositePkQuerier) MustDeleteAll(items ...*HasManyCompositePk) (sql.Result, error) {
	var res sql.Result
	var err error
	for _, d := range items {
		res, err = c.DeleteAll(d)
		if err != nil {
			return res, err
		}
		if n, e := res.RowsAffected(); e != nil {
			return res, e
		} else if n == 0 {
			q := c.Delete().Where(
				JHasManyCompositePkModel.In(items...),
			)
			err = runtime.NewNoRowsAffected(q.String())
			return res, err
		}
	}
	return res, err
}

//Find one HasManyCompositePk using its PKs
func (c jHasManyCompositePkQuerier) Find(ID int64) (*HasManyCompositePk, error) {
	return c.Select().Where(

		JHasManyCompositePkModel.ID.Eq(ID),
	).Read()
}

// Relateds returns a query builder to select Relateds linked to this HasManyCompositePk
func (g *HasManyCompositePk) Relateds(db dbr.SessionRunner,
	AsCompositePk, AsCompositePkrelatedsToHasManyCompositePkrelateds, AsHasManyCompositePk string,
) *jCompositePkSelectBuilder {

	leftTable := JCompositePkModel.Table()
	var query *jCompositePkSelectBuilder
	if AsCompositePk != "" {
		leftTable = AsCompositePk
		query = JCompositePk(db).As(AsCompositePk).Select(AsCompositePk + ".*")
	} else {
		query = JCompositePk(db).Select(leftTable + ".*")
	}

	midTable := JCompositePkrelatedsToHasManyCompositePkrelatedsModel.Table()
	{
		on := ""
		if AsCompositePkrelatedsToHasManyCompositePkrelateds != "" {
			midTable = AsCompositePkrelatedsToHasManyCompositePkrelateds
		}

		on += fmt.Sprintf("%v.%v = %v.%v",
			midTable, "composite_pk_p",
			leftTable, "p",
		)

		on += fmt.Sprintf("%v.%v = %v.%v",
			midTable, "composite_pk_k",
			leftTable, "k",
		)

		if AsCompositePkrelatedsToHasManyCompositePkrelateds == "" {
			query = query.Join(dbr.I(JCompositePkrelatedsToHasManyCompositePkrelatedsModel.Table()), on)
		} else {
			query = query.Join(dbr.I(JCompositePkrelatedsToHasManyCompositePkrelatedsModel.Table()).As(AsCompositePkrelatedsToHasManyCompositePkrelateds), on)
		}
	}

	rightTable := JHasManyCompositePkModel.Table()
	{
		on := ""
		if AsHasManyCompositePk != "" {
			rightTable = AsHasManyCompositePk
		}

		on += fmt.Sprintf("%v.%v = %v.%v",
			midTable, "has_many_composite_pk_id",
			rightTable, "id",
		)

		if AsHasManyCompositePk == "" {
			query = query.Join(dbr.I(JHasManyCompositePkModel.Table()), on)
		} else {
			query = query.Join(dbr.I(JHasManyCompositePkModel.Table()).As(AsHasManyCompositePk), on)
		}
	}

	{
		m := JHasManyCompositePkModel
		if AsHasManyCompositePk != "" {
			m = m.As(AsHasManyCompositePk)
		} else {
			m = m.As(m.Table())
		}
		query = query.Where(

			m.ID.Eq(g.ID),
		)
	}

	return query
}

//LinkWithRelateds writes new links with HasManyCompositePk.
func (g *HasManyCompositePk) LinkWithRelateds(db dbr.SessionRunner, items ...*CompositePk) (sql.Result, error) {
	toInsert := []*CompositePkrelatedsToHasManyCompositePkrelateds{}
	for _, item := range items {
		toInsert = append(toInsert, &CompositePkrelatedsToHasManyCompositePkrelateds{

			CompositePkP: item.P,

			CompositePkK: item.K,

			HasManyCompositePkID: g.ID,
		})
	}
	return JCompositePkrelatedsToHasManyCompositePkrelateds(db).Insert(toInsert...)
}

//UnlinkWithRelateds deletes given existing links with HasManyCompositePk.
func (g *HasManyCompositePk) UnlinkWithRelateds(db dbr.SessionRunner, items ...*CompositePk) (sql.Result, error) {
	toDelete := []*CompositePkrelatedsToHasManyCompositePkrelateds{}
	for _, item := range items {
		toDelete = append(toDelete, &CompositePkrelatedsToHasManyCompositePkrelateds{

			CompositePkP: item.P,

			CompositePkK: item.K,

			HasManyCompositePkID: g.ID,
		})
	}
	return JCompositePkrelatedsToHasManyCompositePkrelateds(db).DeleteAll(toDelete...)
}

//UnlinkAllRelateds deletes all existing links with HasManyCompositePk.
func (g *HasManyCompositePk) UnlinkAllRelateds(db dbr.SessionRunner) (sql.Result, error) {
	return JCompositePkrelatedsToHasManyCompositePkrelateds(db).Delete().Where(

		JCompositePkrelatedsToHasManyCompositePkrelatedsModel.HasManyCompositePkID.Eq(g.ID),
	).Exec()
}

//SetRelateds replaces existing links with HasManyCompositePk.
func (g *HasManyCompositePk) SetRelateds(db dbr.SessionRunner, items ...*CompositePk) (sql.Result, error) {
	if res, err := g.UnlinkAllRelateds(db); err != nil {
		return res, err
	}
	return g.LinkWithRelateds(db, items...)
}

// JoinRelateds adds a JOIN to HasManyCompositePk.Relateds
func (c *jHasManyCompositePkSelectBuilder) JoinRelateds(
	AsCompositePkrelatedsToHasManyCompositePkrelateds, AsCompositePk string,
) *jHasManyCompositePkSelectBuilder {

	query := c

	leftTable := JHasManyCompositePkModel.Table()
	if c.as != "" {
		leftTable = c.as
	}

	midTable := JCompositePkrelatedsToHasManyCompositePkrelatedsModel.Table()
	if AsCompositePkrelatedsToHasManyCompositePkrelateds != "" {
		midTable = AsCompositePkrelatedsToHasManyCompositePkrelateds
	}

	{
		on := ""

		on += fmt.Sprintf("%v.%v = %v.%v",
			midTable, "has_many_composite_pk_id",
			leftTable, "id",
		)

		if AsCompositePkrelatedsToHasManyCompositePkrelateds == "" {
			query = query.Join(dbr.I(JCompositePkrelatedsToHasManyCompositePkrelatedsModel.Table()), on)
		} else {
			query = query.Join(dbr.I(JCompositePkrelatedsToHasManyCompositePkrelatedsModel.Table()).As(AsCompositePkrelatedsToHasManyCompositePkrelateds), on)
		}
	}

	{
		rightTable := JCompositePkModel.Table()
		if AsCompositePk != "" {
			rightTable = AsCompositePk
		}
		on := ""

		on += fmt.Sprintf("%v.%v = %v.%v",
			midTable, "composite_pk_p",
			rightTable, "p",
		)

		on += fmt.Sprintf("%v.%v = %v.%v",
			midTable, "composite_pk_k",
			rightTable, "k",
		)

		if AsCompositePk == "" {
			query = query.Join(dbr.I(JCompositePkModel.Table()), on)
		} else {
			query = query.Join(dbr.I(JCompositePkModel.Table()).As(AsCompositePk), on)
		}
	}

	return query
}

// LeftJoinRelateds adds a LEFT JOIN to HasManyCompositePk.Relateds
func (c *jHasManyCompositePkSelectBuilder) LeftJoinRelateds(
	AsCompositePkrelatedsToHasManyCompositePkrelateds, AsCompositePk string,
) *jHasManyCompositePkSelectBuilder {

	query := c

	leftTable := JHasManyCompositePkModel.Table()
	if c.as != "" {
		leftTable = c.as
	}

	midTable := JCompositePkrelatedsToHasManyCompositePkrelatedsModel.Table()
	if AsCompositePkrelatedsToHasManyCompositePkrelateds != "" {
		midTable = AsCompositePkrelatedsToHasManyCompositePkrelateds
	}

	{
		on := ""

		on += fmt.Sprintf("%v.%v = %v.%v",
			midTable, "has_many_composite_pk_id",
			leftTable, "id",
		)

		if AsCompositePkrelatedsToHasManyCompositePkrelateds == "" {
			query = query.LeftJoin(dbr.I(JCompositePkrelatedsToHasManyCompositePkrelatedsModel.Table()), on)
		} else {
			query = query.LeftJoin(dbr.I(JCompositePkrelatedsToHasManyCompositePkrelatedsModel.Table()).As(AsCompositePkrelatedsToHasManyCompositePkrelateds), on)
		}
	}

	{
		rightTable := JCompositePkModel.Table()
		if AsCompositePk != "" {
			rightTable = AsCompositePk
		}
		on := ""

		on += fmt.Sprintf("%v.%v = %v.%v",
			midTable, "composite_pk_p",
			rightTable, "p",
		)

		on += fmt.Sprintf("%v.%v = %v.%v",
			midTable, "composite_pk_k",
			rightTable, "k",
		)

		if AsCompositePk == "" {
			query = query.LeftJoin(dbr.I(JCompositePkModel.Table()), on)
		} else {
			query = query.LeftJoin(dbr.I(JCompositePkModel.Table()).As(AsCompositePk), on)
		}
	}

	return query
}

// // RightJoinRelateds adds a RIGHT JOIN to HasManyCompositePk.Relateds
// func (c *jHasManyCompositePkSelectBuilder) RightJoinRelateds(
// 	AsCompositePkrelatedsToHasManyCompositePkrelateds, AsCompositePk string,
// ) *jHasManyCompositePkSelectBuilder {
//
// 	query := c
//
// 	leftTable := JHasManyCompositePkModel.Table()
// 	if c.as != "" {
// 		leftTable = c.as
// 	}
//
// 	midTable := JCompositePkrelatedsToHasManyCompositePkrelatedsModel.Table()
// 	if AsCompositePkrelatedsToHasManyCompositePkrelateds != "" {
// 		midTable = AsCompositePkrelatedsToHasManyCompositePkrelateds
// 	}
//
// 	{
// 		on := ""
//
// 		on += fmt.Sprintf("%v.%v = %v.%v",
// 			midTable, "has_many_composite_pk_id",
// 			leftTable, "id",
// 			)
//
//
// 		if AsCompositePkrelatedsToHasManyCompositePkrelateds == "" {
// 			query = query.RightJoin(dbr.I(JCompositePkrelatedsToHasManyCompositePkrelatedsModel.Table()), on)
// 		} else {
// 			query = query.RightJoin(dbr.I(JCompositePkrelatedsToHasManyCompositePkrelatedsModel.Table()).As(AsCompositePkrelatedsToHasManyCompositePkrelateds), on)
// 		}
// 	}
//
// 	{
// 		rightTable := JCompositePkModel.Table()
// 		if AsCompositePk != "" {
// 			rightTable = AsCompositePk
// 		}
// 		on := ""
//
// 		on += fmt.Sprintf("%v.%v = %v.%v",
// 			midTable, "composite_pk_p",
// 			rightTable, "p",
// 			)
//
// 		on += fmt.Sprintf("%v.%v = %v.%v",
// 			midTable, "composite_pk_k",
// 			rightTable, "k",
// 			)
//
//
// 		if AsCompositePk == "" {
// 			query = query.RightJoin(dbr.I(JCompositePkModel.Table()), on)
// 		} else {
// 			query = query.RightJoin(dbr.I(JCompositePkModel.Table()).As(AsCompositePk), on)
// 		}
// 	}
//
// 	return query
// }

// JDateTypeSetup helps to create/drop the schema
func JDateTypeSetup() runtime.Setuper {
	driver := runtime.GetCurrentDriver()

	var create string
	var drop string

	if driver == drivers.Sqlite {
		create = `CREATE TABLE IF NOT EXISTS date_type (
id INTEGER PRIMARY KEY AUTOINCREMENT,
t datetime,
tp datetime NULL,
not_utc datetime NULL,
last_updated datetime NULL

)`
	} else if driver == drivers.Mysql {
		create = `CREATE TABLE IF NOT EXISTS date_type (
id INTEGER NOT NULL AUTO_INCREMENT,
t datetime(6),
tp datetime(6) NULL,
not_utc datetime(6) NULL,
last_updated datetime(6) NULL,
PRIMARY KEY (id) 

)`
	} else if driver == drivers.Pgsql {
		create = `CREATE TABLE IF NOT EXISTS date_type (
id SERIAL PRIMARY KEY,
t timestamp(6),
tp timestamp(6) NULL,
not_utc timestamp(6) NULL,
last_updated timestamp(6) NULL

)`
	}

	if driver == drivers.Sqlite {
		drop = `DROP TABLE IF EXISTS date_type`
	} else if driver == drivers.Mysql {
		drop = `DROP TABLE IF EXISTS date_type`
	} else if driver == drivers.Pgsql {
		drop = `DROP TABLE IF EXISTS date_type`
	}

	var indexes []string

	if driver == drivers.Sqlite {

	} else if driver == drivers.Mysql {

	} else if driver == drivers.Pgsql {

	}

	return runtime.Table{
		Name:       `date_type`,
		CreateStmt: create,
		DropStmt:   drop,
		View:       !true,
		Indexes:    indexes,
	}
}

// jDateTypeModel provides helper to work with DateType data provider
type jDateTypeModel struct {
	as string

	ID builder.ValuePropertyMeta

	T builder.ValuePropertyMeta

	TP builder.ValuePropertyMeta

	NotUTC builder.ValuePropertyMeta

	LastUpdated builder.ValuePropertyMeta
}

// Eq provided items.
func (j jDateTypeModel) Eq(s ...*DateType) dbr.Builder {
	ors := []dbr.Builder{}
	for _, t := range s {
		ors = append(ors, dbr.And(

			JDateTypeModel.ID.Eq(t.ID),
		))
	}
	return dbr.Or(ors...)
}

// In provided items.
func (j jDateTypeModel) In(s ...*DateType) dbr.Builder {
	ors := []dbr.Builder{}
	for _, t := range s {
		ors = append(ors, dbr.And(

			JDateTypeModel.ID.Eq(t.ID),
		))
	}
	return dbr.Or(ors...)
}

// As returns a copy with an alias.
func (j jDateTypeModel) As(as string) jDateTypeModel {
	j.as = as

	j.ID.TableAlias = as

	j.T.TableAlias = as

	j.TP.TableAlias = as

	j.NotUTC.TableAlias = as

	j.LastUpdated.TableAlias = as

	return j
}

// Table returns the sql table name
func (j jDateTypeModel) Table() string {
	return "date_type"
}

// Alias returns the current alias
func (j jDateTypeModel) Alias() string {
	if j.as == "" {
		return j.Table()
	}
	return j.as
}

// Properties returns a map of property name => meta
func (j jDateTypeModel) Properties() map[string]builder.MetaProvider {
	ret := map[string]builder.MetaProvider{}

	ret["ID"] = j.ID

	ret["T"] = j.T

	ret["TP"] = j.TP

	ret["NotUTC"] = j.NotUTC

	ret["LastUpdated"] = j.LastUpdated

	return ret
}

// Fields returns given sql fields with appropriate aliasing.
func (j jDateTypeModel) Fields(ins ...string) []string {
	dialect := runtime.GetDialect()
	if len(ins) == 0 {
		ins = append(ins, "*")
	}
	for i, in := range ins {
		if j.as != "" {
			if in == "*" {
				ins[i] = fmt.Sprintf("%v.%v", dialect.QuoteIdent(j.as), in)
			} else {
				ins[i] = fmt.Sprintf("%v.%v", dialect.QuoteIdent(j.as), dialect.QuoteIdent(in))
			}
		}
	}
	return ins
}

// JDateTypeModel provides helper to work with DateType data provider
var JDateTypeModel = jDateTypeModel{

	ID: builder.NewValueMeta(
		`id`, `INTEGER`,
		`ID`, `int64`,
		true, true,
	),

	T: builder.NewValueMeta(
		`t`, `datetime`,
		`T`, `time.Time`,
		false, false,
	),

	TP: builder.NewValueMeta(
		`tp`, `datetime`,
		`TP`, `*time.Time`,
		false, false,
	),

	NotUTC: builder.NewValueMeta(
		`not_utc`, `datetime`,
		`NotUTC`, `*time.Time`,
		false, false,
	),

	LastUpdated: builder.NewValueMeta(
		`last_updated`, `datetime`,
		`LastUpdated`, `*time.Time`,
		false, false,
	),
}

type jDateTypeDeleteBuilder struct {
	*builder.DeleteBuilder
}

// //Build builds the sql string into given buffer using current dialect
// func (c *jDateTypeDeleteBuilder) Build(b dbr.Buffer) error {
// 	return c.DeleteBuilder.Build(runtime.GetDialect(), b)
// }
// //String returns the sql string for current dialect. It returns empty string if the build returns an error.
// func (c *jDateTypeDeleteBuilder) String() string {
// 	b := dbr.NewBuffer()
// 	if err := c.Build(b); err != nil {
// 		return ""
// 	}
// 	return b.String()
// }
//Where returns a jDateTypeDeleteBuilder instead of builder.DeleteBuilder.
func (c *jDateTypeDeleteBuilder) Where(query interface{}, value ...interface{}) *jDateTypeDeleteBuilder {
	c.DeleteBuilder.Where(query, value...)
	return c
}

type jDateTypeSelectBuilder struct {
	as string
	*builder.SelectBuilder
}

// //Build builds the sql string using current dialect into given bufer
// func (c *jDateTypeSelectBuilder) Build(b dbr.Buffer) error {
// 	return c.SelectBuilder.Build(runtime.GetDialect(), b)
// }
// //String returns the sql string for current dialect. It returns empty string if the build returns an error.
// func (c *jDateTypeSelectBuilder) String() string {
// 	b := dbr.NewBuffer()
// 	if err := c.Build(b); err != nil {
// 		return ""
// 	}
// 	return b.String()
// }

//Read evaluates current select query and load the results into a DateType
func (c *jDateTypeSelectBuilder) Read() (*DateType, error) {
	var one DateType
	err := c.LoadStruct(&one)
	return &one, err
}

//ReadAll evaluates current select query and load the results into a slice of DateType
func (c *jDateTypeSelectBuilder) ReadAll() ([]*DateType, error) {
	var all []*DateType
	_, err := c.LoadStructs(&all)
	return all, err
}

//Where returns a jDateTypeSelectBuilder instead of builder.SelectBuilder.
func (c *jDateTypeSelectBuilder) Where(query interface{}, value ...interface{}) *jDateTypeSelectBuilder {
	c.SelectBuilder.Where(query, value...)
	return c
}

//GroupBy returns a jDateTypeSelectBuilder instead of builder.SelectBuilder.
func (c *jDateTypeSelectBuilder) GroupBy(col ...string) *jDateTypeSelectBuilder {
	c.SelectBuilder.GroupBy(col...)
	return c
}

//Having returns a jDateTypeSelectBuilder instead of builder.SelectBuilder.
func (c *jDateTypeSelectBuilder) Having(query interface{}, value ...interface{}) *jDateTypeSelectBuilder {
	c.SelectBuilder.Having(query, value...)
	return c
}

//Limit returns a jDateTypeSelectBuilder instead of builder.SelectBuilder.
func (c *jDateTypeSelectBuilder) Limit(n uint64) *jDateTypeSelectBuilder {
	c.SelectBuilder.Limit(n)
	return c
}

//Offset returns a jDateTypeSelectBuilder instead of builder.SelectBuilder.
func (c *jDateTypeSelectBuilder) Offset(n uint64) *jDateTypeSelectBuilder {
	c.SelectBuilder.Offset(n)
	return c
}

//OrderAsc returns a jDateTypeSelectBuilder instead of builder.SelectBuilder.
func (c *jDateTypeSelectBuilder) OrderAsc(col string) *jDateTypeSelectBuilder {
	c.SelectBuilder.OrderAsc(col)
	return c
}

//OrderDesc returns a jDateTypeSelectBuilder instead of builder.SelectBuilder.
func (c *jDateTypeSelectBuilder) OrderDesc(col string) *jDateTypeSelectBuilder {
	c.SelectBuilder.OrderDesc(col)
	return c
}

//OrderDir returns a jDateTypeSelectBuilder instead of builder.SelectBuilder.
func (c *jDateTypeSelectBuilder) OrderDir(col string, isAsc bool) *jDateTypeSelectBuilder {
	c.SelectBuilder.OrderDir(col, isAsc)
	return c
}

//OrderBy returns a jDateTypeSelectBuilder instead of builder.SelectBuilder.
func (c *jDateTypeSelectBuilder) OrderBy(col string) *jDateTypeSelectBuilder {
	c.SelectBuilder.OrderBy(col)
	return c
}

//Paginate returns a jDateTypeSelectBuilder instead of builder.SelectBuilder.
func (c *jDateTypeSelectBuilder) Paginate(page, perPage uint64) *jDateTypeSelectBuilder {
	c.SelectBuilder.Paginate(page, perPage)
	return c
}

//Join returns a jDateTypeSelectBuilder instead of builder.SelectBuilder.
func (c *jDateTypeSelectBuilder) Join(table, on interface{}) *jDateTypeSelectBuilder {
	c.SelectBuilder.Join(table, on)
	return c
}

//LeftJoin returns a jDateTypeSelectBuilder instead of builder.SelectBuilder.
func (c *jDateTypeSelectBuilder) LeftJoin(table, on interface{}) *jDateTypeSelectBuilder {
	c.SelectBuilder.LeftJoin(table, on)
	return c
}

//RightJoin returns a jDateTypeSelectBuilder instead of builder.SelectBuilder.
func (c *jDateTypeSelectBuilder) RightJoin(table, on interface{}) *jDateTypeSelectBuilder {
	c.SelectBuilder.RightJoin(table, on)
	return c
}

//FullJoin returns a jDateTypeSelectBuilder instead of builder.SelectBuilder.
func (c *jDateTypeSelectBuilder) FullJoin(table, on interface{}) *jDateTypeSelectBuilder {
	c.SelectBuilder.FullJoin(table, on)
	return c
}

//Distinct returns a jDateTypeSelectBuilder instead of builder.SelectBuilder.
func (c *jDateTypeSelectBuilder) Distinct() *jDateTypeSelectBuilder {
	c.SelectBuilder.Distinct()
	return c
}

// JDateType provides a basic querier
func JDateType(db dbr.SessionRunner) jDateTypeQuerier {
	return jDateTypeQuerier{
		db: db,
	}
}

type jDateTypeQuerier struct {
	db dbr.SessionRunner
	as string
}

//As set alias prior building.
func (c jDateTypeQuerier) As(as string) jDateTypeQuerier {
	c.as = as
	return c
}

//Model returns a model with appropriate aliasing.
func (c jDateTypeQuerier) Model() jDateTypeModel {
	return JDateTypeModel.As(c.as)
}

//Select returns a DateType Select Builder.
func (c jDateTypeQuerier) Select(what ...string) *jDateTypeSelectBuilder {
	m := c.Model()
	dialect := runtime.GetDialect()
	from := dialect.QuoteIdent(m.Table())
	if m.Alias() != "" && m.Alias() != m.Table() {
		from = fmt.Sprintf("%v as %v", from, dialect.QuoteIdent(m.Alias()))
	}
	if len(what) == 0 {
		alias := m.Table()
		if m.Alias() != "" && m.Alias() != m.Table() {
			alias = m.Alias()
		}
		what = m.Fields(alias + ".*")
	}
	return &jDateTypeSelectBuilder{
		as: c.as,
		SelectBuilder: &builder.SelectBuilder{
			SelectBuilder: c.db.Select(what...).From(from),
		},
	}
}

//Where returns a DateType Select Builder.
func (c jDateTypeQuerier) Where(query interface{}, value ...interface{}) *jDateTypeSelectBuilder {
	return c.Select().Where(query, value...)
}

//Count returns a DateType Select Builder to count given expressions.
func (c jDateTypeQuerier) Count(what ...string) *jDateTypeSelectBuilder {
	if len(what) == 0 {
		what = append(what, "*")
	}
	return c.Select("COUNT(" + strings.Join(what, ",") + ")")
}

// Insert a new DateType, if it has autoincrement primary key, the value will be set.
// It stops on first error.
func (c jDateTypeQuerier) Insert(items ...*DateType) (sql.Result, error) {
	var res sql.Result
	var err error
	for _, data := range items {

		if data.LastUpdated == nil {
			x := time.Now()
			data.LastUpdated = &x
		}
		{
			x := data.LastUpdated.Truncate(time.Microsecond)
			data.LastUpdated = &x
		}

		data.T = data.T.UTC()

		if data.TP != nil {
			x := data.TP.UTC()
			data.TP = &x
		}

		if data.LastUpdated != nil {
			x := data.LastUpdated.UTC()
			data.LastUpdated = &x
		}

		query := c.db.InsertInto(JDateTypeModel.Table()).Columns(

			`t`,

			`tp`,

			`not_utc`,

			`last_updated`,
		).Record(data)
		if runtime.Runs(drivers.Pgsql) {

			query = query.Returning(

				`id`,
			)

			var auto0 int64

			err = query.Load(

				&auto0,
			)

			data.ID = auto0

		} else {
			res, err = query.Exec()

			if err == nil {
				id, err2 := res.LastInsertId()
				if err2 != nil {
					return res, err2
				}
				data.ID = id
			}

		}
		if err != nil {
			return res, err
		}
	}
	return res, err
}

// InsertBulk inserts multiple items into the database.
// It does not post update any auto increment field.
// It builds an insert query of multiple rows and send it on the underlying connection.
func (c jDateTypeQuerier) InsertBulk(items ...*DateType) error {
	panic("todo")
}

// Update a DateType. It stops on first error.
func (c jDateTypeQuerier) Update(items ...*DateType) (sql.Result, error) {
	var res sql.Result
	var err error
	for _, data := range items {

		data.T = data.T.UTC()

		if data.TP != nil {
			x := data.TP.UTC()
			data.TP = &x
		}

		if data.LastUpdated != nil {
			x := data.LastUpdated.UTC()
			data.LastUpdated = &x
		}

		currentDate := data.LastUpdated
		newDate := time.Now().UTC().Truncate(time.Microsecond)

		if currentDate != nil {
			y := currentDate.Truncate(time.Microsecond)
			currentDate = &y
		}

		query := c.db.Update(JDateTypeModel.Table())

		query = query.Set(`t`, data.T)

		query = query.Set(`tp`, data.TP)

		query = query.Set(`not_utc`, data.NotUTC)

		query = query.Set(`last_updated`, newDate)

		query = query.Where("id = ?", data.ID)

		if currentDate == nil {
			query = query.Where("last_updated IS NULL")
		} else {
			query = query.Where("last_updated = ?", currentDate)
		}

		res, err = query.Exec()

		if err == nil {

			data.LastUpdated = &newDate

		}

		if err != nil {
			return res, err
		}
	}
	return res, err
}

// MustUpdate a DateType. It stops on first error. It errors if an update query does not affect row.
func (c jDateTypeQuerier) MustUpdate(items ...*DateType) (sql.Result, error) {
	var res sql.Result
	var err error
	for _, data := range items {

		data.T = data.T.UTC()

		if data.TP != nil {
			x := data.TP.UTC()
			data.TP = &x
		}

		if data.LastUpdated != nil {
			x := data.LastUpdated.UTC()
			data.LastUpdated = &x
		}

		currentDate := data.LastUpdated
		newDate := time.Now().UTC().Truncate(time.Microsecond)

		if currentDate != nil {
			y := currentDate.Truncate(time.Microsecond)
			currentDate = &y
		}

		query := c.db.Update(JDateTypeModel.Table())

		query = query.Set(`t`, data.T)

		query = query.Set(`tp`, data.TP)

		query = query.Set(`not_utc`, data.NotUTC)

		query = query.Set(`last_updated`, newDate)

		query = query.Where("id = ?", data.ID)

		if currentDate == nil {
			query = query.Where("last_updated IS NULL")
		} else {
			query = query.Where("last_updated = ?", currentDate)
		}

		res, err = query.Exec()

		if err == nil {
			if n, _ := res.RowsAffected(); n == 0 {
				x := &builder.UpdateBuilder{UpdateBuilder: query}
				err = runtime.NewNoRowsAffected(x.String())
			}
		}

		if err == nil {

			data.LastUpdated = &newDate

		}

		if err != nil {
			return res, err
		}
	}
	return res, err
}

// UpdateBulk updates multiple items into the database.
// It builds an update query of multiple rows and send it on the underlying connection.
func (c jDateTypeQuerier) UpdateBulk(items ...*DateType) error {
	panic("todo")
}

//Delete returns a delete builder
func (c jDateTypeQuerier) Delete() *jDateTypeDeleteBuilder {
	return &jDateTypeDeleteBuilder{
		&builder.DeleteBuilder{
			DeleteBuilder: c.db.DeleteFrom(JDateTypeModel.Table()),
		},
	}
}

// MustDelete requires the query to affeect rows.
func (c jDateTypeQuerier) MustDelete() *jDateTypeDeleteBuilder {
	ret := &jDateTypeDeleteBuilder{
		&builder.DeleteBuilder{
			DeleteBuilder: c.db.DeleteFrom(JDateTypeModel.Table()),
		},
	}
	ret.MustDelete()
	return ret
}

//DeleteByPk deletes one DateType by its PKs
func (c jDateTypeQuerier) DeleteByPk(ID int64) error {
	_, err := c.Delete().Where(

		JDateTypeModel.ID.Eq(ID),
	).Exec()
	return err
}

// DeleteAll given DateType
func (c jDateTypeQuerier) DeleteAll(items ...*DateType) (sql.Result, error) {
	q := c.Delete().Where(
		JDateTypeModel.In(items...),
	)
	return q.Exec()
}

// MustDeleteAll given DateType
func (c jDateTypeQuerier) MustDeleteAll(items ...*DateType) (sql.Result, error) {
	var res sql.Result
	var err error
	for _, d := range items {
		res, err = c.DeleteAll(d)
		if err != nil {
			return res, err
		}
		if n, e := res.RowsAffected(); e != nil {
			return res, e
		} else if n == 0 {
			q := c.Delete().Where(
				JDateTypeModel.In(items...),
			)
			err = runtime.NewNoRowsAffected(q.String())
			return res, err
		}
	}
	return res, err
}

//Find one DateType using its PKs
func (c jDateTypeQuerier) Find(ID int64) (*DateType, error) {
	return c.Select().Where(

		JDateTypeModel.ID.Eq(ID),
	).Read()
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
`
	} else if driver == drivers.Mysql {
		create = `CREATE OR REPLACE VIEW sample_view AS 
	SELECT *
	FROM sample
	WHERE id > 1
`
	} else if driver == drivers.Pgsql {
		create = `CREATE OR REPLACE VIEW sample_view AS 
	SELECT *
	FROM sample
	WHERE id > 1
`
	}

	if driver == drivers.Sqlite {
		drop = `DROP VIEW IF EXISTS sample_view`
	} else if driver == drivers.Mysql {
		drop = `DROP VIEW IF EXISTS sample_view`
	} else if driver == drivers.Pgsql {
		drop = `DROP VIEW IF EXISTS sample_view`
	}

	var indexes []string

	if driver == drivers.Sqlite {

	} else if driver == drivers.Mysql {

	} else if driver == drivers.Pgsql {

	}

	return runtime.Table{
		Name:       `sample_view`,
		CreateStmt: create,
		DropStmt:   drop,
		View:       !false,
		Indexes:    indexes,
	}
}

// jSampleViewModel provides helper to work with SampleView data provider
type jSampleViewModel struct {
	as string

	ID builder.ValuePropertyMeta

	Name builder.ValuePropertyMeta

	Description builder.ValuePropertyMeta
}

// Eq provided items.
func (j jSampleViewModel) Eq(s ...*SampleView) dbr.Builder {
	ors := []dbr.Builder{}
	for _, t := range s {
		ors = append(ors, dbr.And(

			JSampleViewModel.ID.Eq(t.ID),
		))
	}
	return dbr.Or(ors...)
}

// In provided items.
func (j jSampleViewModel) In(s ...*SampleView) dbr.Builder {
	ors := []dbr.Builder{}
	for _, t := range s {
		ors = append(ors, dbr.And(

			JSampleViewModel.ID.Eq(t.ID),
		))
	}
	return dbr.Or(ors...)
}

// As returns a copy with an alias.
func (j jSampleViewModel) As(as string) jSampleViewModel {
	j.as = as

	j.ID.TableAlias = as

	j.Name.TableAlias = as

	j.Description.TableAlias = as

	return j
}

// Table returns the sql table name
func (j jSampleViewModel) Table() string {
	return "sample_view"
}

// Alias returns the current alias
func (j jSampleViewModel) Alias() string {
	if j.as == "" {
		return j.Table()
	}
	return j.as
}

// Properties returns a map of property name => meta
func (j jSampleViewModel) Properties() map[string]builder.MetaProvider {
	ret := map[string]builder.MetaProvider{}

	ret["ID"] = j.ID

	ret["Name"] = j.Name

	ret["Description"] = j.Description

	return ret
}

// Fields returns given sql fields with appropriate aliasing.
func (j jSampleViewModel) Fields(ins ...string) []string {
	dialect := runtime.GetDialect()
	if len(ins) == 0 {
		ins = append(ins, "*")
	}
	for i, in := range ins {
		if j.as != "" {
			if in == "*" {
				ins[i] = fmt.Sprintf("%v.%v", dialect.QuoteIdent(j.as), in)
			} else {
				ins[i] = fmt.Sprintf("%v.%v", dialect.QuoteIdent(j.as), dialect.QuoteIdent(in))
			}
		}
	}
	return ins
}

// JSampleViewModel provides helper to work with SampleView data provider
var JSampleViewModel = jSampleViewModel{

	ID: builder.NewValueMeta(
		`id`, `INTEGER`,
		`ID`, `int64`,
		true, true,
	),

	Name: builder.NewValueMeta(
		`name`, `TEXT`,
		`Name`, `string`,
		false, false,
	),

	Description: builder.NewValueMeta(
		`description`, `TEXT`,
		`Description`, `string`,
		false, false,
	),
}

type jSampleViewDeleteBuilder struct {
	*builder.DeleteBuilder
}

// //Build builds the sql string into given buffer using current dialect
// func (c *jSampleViewDeleteBuilder) Build(b dbr.Buffer) error {
// 	return c.DeleteBuilder.Build(runtime.GetDialect(), b)
// }
// //String returns the sql string for current dialect. It returns empty string if the build returns an error.
// func (c *jSampleViewDeleteBuilder) String() string {
// 	b := dbr.NewBuffer()
// 	if err := c.Build(b); err != nil {
// 		return ""
// 	}
// 	return b.String()
// }
//Where returns a jSampleViewDeleteBuilder instead of builder.DeleteBuilder.
func (c *jSampleViewDeleteBuilder) Where(query interface{}, value ...interface{}) *jSampleViewDeleteBuilder {
	c.DeleteBuilder.Where(query, value...)
	return c
}

type jSampleViewSelectBuilder struct {
	as string
	*builder.SelectBuilder
}

// //Build builds the sql string using current dialect into given bufer
// func (c *jSampleViewSelectBuilder) Build(b dbr.Buffer) error {
// 	return c.SelectBuilder.Build(runtime.GetDialect(), b)
// }
// //String returns the sql string for current dialect. It returns empty string if the build returns an error.
// func (c *jSampleViewSelectBuilder) String() string {
// 	b := dbr.NewBuffer()
// 	if err := c.Build(b); err != nil {
// 		return ""
// 	}
// 	return b.String()
// }

//Read evaluates current select query and load the results into a SampleView
func (c *jSampleViewSelectBuilder) Read() (*SampleView, error) {
	var one SampleView
	err := c.LoadStruct(&one)
	return &one, err
}

//ReadAll evaluates current select query and load the results into a slice of SampleView
func (c *jSampleViewSelectBuilder) ReadAll() ([]*SampleView, error) {
	var all []*SampleView
	_, err := c.LoadStructs(&all)
	return all, err
}

//Where returns a jSampleViewSelectBuilder instead of builder.SelectBuilder.
func (c *jSampleViewSelectBuilder) Where(query interface{}, value ...interface{}) *jSampleViewSelectBuilder {
	c.SelectBuilder.Where(query, value...)
	return c
}

//GroupBy returns a jSampleViewSelectBuilder instead of builder.SelectBuilder.
func (c *jSampleViewSelectBuilder) GroupBy(col ...string) *jSampleViewSelectBuilder {
	c.SelectBuilder.GroupBy(col...)
	return c
}

//Having returns a jSampleViewSelectBuilder instead of builder.SelectBuilder.
func (c *jSampleViewSelectBuilder) Having(query interface{}, value ...interface{}) *jSampleViewSelectBuilder {
	c.SelectBuilder.Having(query, value...)
	return c
}

//Limit returns a jSampleViewSelectBuilder instead of builder.SelectBuilder.
func (c *jSampleViewSelectBuilder) Limit(n uint64) *jSampleViewSelectBuilder {
	c.SelectBuilder.Limit(n)
	return c
}

//Offset returns a jSampleViewSelectBuilder instead of builder.SelectBuilder.
func (c *jSampleViewSelectBuilder) Offset(n uint64) *jSampleViewSelectBuilder {
	c.SelectBuilder.Offset(n)
	return c
}

//OrderAsc returns a jSampleViewSelectBuilder instead of builder.SelectBuilder.
func (c *jSampleViewSelectBuilder) OrderAsc(col string) *jSampleViewSelectBuilder {
	c.SelectBuilder.OrderAsc(col)
	return c
}

//OrderDesc returns a jSampleViewSelectBuilder instead of builder.SelectBuilder.
func (c *jSampleViewSelectBuilder) OrderDesc(col string) *jSampleViewSelectBuilder {
	c.SelectBuilder.OrderDesc(col)
	return c
}

//OrderDir returns a jSampleViewSelectBuilder instead of builder.SelectBuilder.
func (c *jSampleViewSelectBuilder) OrderDir(col string, isAsc bool) *jSampleViewSelectBuilder {
	c.SelectBuilder.OrderDir(col, isAsc)
	return c
}

//OrderBy returns a jSampleViewSelectBuilder instead of builder.SelectBuilder.
func (c *jSampleViewSelectBuilder) OrderBy(col string) *jSampleViewSelectBuilder {
	c.SelectBuilder.OrderBy(col)
	return c
}

//Paginate returns a jSampleViewSelectBuilder instead of builder.SelectBuilder.
func (c *jSampleViewSelectBuilder) Paginate(page, perPage uint64) *jSampleViewSelectBuilder {
	c.SelectBuilder.Paginate(page, perPage)
	return c
}

//Join returns a jSampleViewSelectBuilder instead of builder.SelectBuilder.
func (c *jSampleViewSelectBuilder) Join(table, on interface{}) *jSampleViewSelectBuilder {
	c.SelectBuilder.Join(table, on)
	return c
}

//LeftJoin returns a jSampleViewSelectBuilder instead of builder.SelectBuilder.
func (c *jSampleViewSelectBuilder) LeftJoin(table, on interface{}) *jSampleViewSelectBuilder {
	c.SelectBuilder.LeftJoin(table, on)
	return c
}

//RightJoin returns a jSampleViewSelectBuilder instead of builder.SelectBuilder.
func (c *jSampleViewSelectBuilder) RightJoin(table, on interface{}) *jSampleViewSelectBuilder {
	c.SelectBuilder.RightJoin(table, on)
	return c
}

//FullJoin returns a jSampleViewSelectBuilder instead of builder.SelectBuilder.
func (c *jSampleViewSelectBuilder) FullJoin(table, on interface{}) *jSampleViewSelectBuilder {
	c.SelectBuilder.FullJoin(table, on)
	return c
}

//Distinct returns a jSampleViewSelectBuilder instead of builder.SelectBuilder.
func (c *jSampleViewSelectBuilder) Distinct() *jSampleViewSelectBuilder {
	c.SelectBuilder.Distinct()
	return c
}

// JSampleView provides a basic querier
func JSampleView(db dbr.SessionRunner) jSampleViewQuerier {
	return jSampleViewQuerier{
		db: db,
	}
}

type jSampleViewQuerier struct {
	db dbr.SessionRunner
	as string
}

//As set alias prior building.
func (c jSampleViewQuerier) As(as string) jSampleViewQuerier {
	c.as = as
	return c
}

//Model returns a model with appropriate aliasing.
func (c jSampleViewQuerier) Model() jSampleViewModel {
	return JSampleViewModel.As(c.as)
}

//Select returns a SampleView Select Builder.
func (c jSampleViewQuerier) Select(what ...string) *jSampleViewSelectBuilder {
	m := c.Model()
	dialect := runtime.GetDialect()
	from := dialect.QuoteIdent(m.Table())
	if m.Alias() != "" && m.Alias() != m.Table() {
		from = fmt.Sprintf("%v as %v", from, dialect.QuoteIdent(m.Alias()))
	}
	if len(what) == 0 {
		alias := m.Table()
		if m.Alias() != "" && m.Alias() != m.Table() {
			alias = m.Alias()
		}
		what = m.Fields(alias + ".*")
	}
	return &jSampleViewSelectBuilder{
		as: c.as,
		SelectBuilder: &builder.SelectBuilder{
			SelectBuilder: c.db.Select(what...).From(from),
		},
	}
}

//Where returns a SampleView Select Builder.
func (c jSampleViewQuerier) Where(query interface{}, value ...interface{}) *jSampleViewSelectBuilder {
	return c.Select().Where(query, value...)
}

//Count returns a SampleView Select Builder to count given expressions.
func (c jSampleViewQuerier) Count(what ...string) *jSampleViewSelectBuilder {
	if len(what) == 0 {
		what = append(what, "*")
	}
	return c.Select("COUNT(" + strings.Join(what, ",") + ")")
}

// JHookDemoSetup helps to create/drop the schema
func JHookDemoSetup() runtime.Setuper {
	driver := runtime.GetCurrentDriver()

	var create string
	var drop string

	if driver == drivers.Sqlite {
		create = `CREATE TABLE IF NOT EXISTS hook_demo (
id INTEGER PRIMARY KEY AUTOINCREMENT,
name TEXT

)`
	} else if driver == drivers.Mysql {
		create = `CREATE TABLE IF NOT EXISTS hook_demo (
id INTEGER NOT NULL AUTO_INCREMENT,
name TEXT,
PRIMARY KEY (id) 

)`
	} else if driver == drivers.Pgsql {
		create = `CREATE TABLE IF NOT EXISTS hook_demo (
id SERIAL PRIMARY KEY,
name TEXT

)`
	}

	if driver == drivers.Sqlite {
		drop = `DROP TABLE IF EXISTS hook_demo`
	} else if driver == drivers.Mysql {
		drop = `DROP TABLE IF EXISTS hook_demo`
	} else if driver == drivers.Pgsql {
		drop = `DROP TABLE IF EXISTS hook_demo`
	}

	var indexes []string

	if driver == drivers.Sqlite {

	} else if driver == drivers.Mysql {

	} else if driver == drivers.Pgsql {

	}

	return runtime.Table{
		Name:       `hook_demo`,
		CreateStmt: create,
		DropStmt:   drop,
		View:       !true,
		Indexes:    indexes,
	}
}

// jHookDemoModel provides helper to work with HookDemo data provider
type jHookDemoModel struct {
	as string

	ID builder.ValuePropertyMeta

	Name builder.ValuePropertyMeta
}

// Eq provided items.
func (j jHookDemoModel) Eq(s ...*HookDemo) dbr.Builder {
	ors := []dbr.Builder{}
	for _, t := range s {
		ors = append(ors, dbr.And(

			JHookDemoModel.ID.Eq(t.ID),
		))
	}
	return dbr.Or(ors...)
}

// In provided items.
func (j jHookDemoModel) In(s ...*HookDemo) dbr.Builder {
	ors := []dbr.Builder{}
	for _, t := range s {
		ors = append(ors, dbr.And(

			JHookDemoModel.ID.Eq(t.ID),
		))
	}
	return dbr.Or(ors...)
}

// As returns a copy with an alias.
func (j jHookDemoModel) As(as string) jHookDemoModel {
	j.as = as

	j.ID.TableAlias = as

	j.Name.TableAlias = as

	return j
}

// Table returns the sql table name
func (j jHookDemoModel) Table() string {
	return "hook_demo"
}

// Alias returns the current alias
func (j jHookDemoModel) Alias() string {
	if j.as == "" {
		return j.Table()
	}
	return j.as
}

// Properties returns a map of property name => meta
func (j jHookDemoModel) Properties() map[string]builder.MetaProvider {
	ret := map[string]builder.MetaProvider{}

	ret["ID"] = j.ID

	ret["Name"] = j.Name

	return ret
}

// Fields returns given sql fields with appropriate aliasing.
func (j jHookDemoModel) Fields(ins ...string) []string {
	dialect := runtime.GetDialect()
	if len(ins) == 0 {
		ins = append(ins, "*")
	}
	for i, in := range ins {
		if j.as != "" {
			if in == "*" {
				ins[i] = fmt.Sprintf("%v.%v", dialect.QuoteIdent(j.as), in)
			} else {
				ins[i] = fmt.Sprintf("%v.%v", dialect.QuoteIdent(j.as), dialect.QuoteIdent(in))
			}
		}
	}
	return ins
}

// JHookDemoModel provides helper to work with HookDemo data provider
var JHookDemoModel = jHookDemoModel{

	ID: builder.NewValueMeta(
		`id`, `INTEGER`,
		`ID`, `int64`,
		true, true,
	),

	Name: builder.NewValueMeta(
		`name`, `TEXT`,
		`Name`, `string`,
		false, false,
	),
}

type jHookDemoDeleteBuilder struct {
	*builder.DeleteBuilder
}

// //Build builds the sql string into given buffer using current dialect
// func (c *jHookDemoDeleteBuilder) Build(b dbr.Buffer) error {
// 	return c.DeleteBuilder.Build(runtime.GetDialect(), b)
// }
// //String returns the sql string for current dialect. It returns empty string if the build returns an error.
// func (c *jHookDemoDeleteBuilder) String() string {
// 	b := dbr.NewBuffer()
// 	if err := c.Build(b); err != nil {
// 		return ""
// 	}
// 	return b.String()
// }
//Where returns a jHookDemoDeleteBuilder instead of builder.DeleteBuilder.
func (c *jHookDemoDeleteBuilder) Where(query interface{}, value ...interface{}) *jHookDemoDeleteBuilder {
	c.DeleteBuilder.Where(query, value...)
	return c
}

type jHookDemoSelectBuilder struct {
	as string
	*builder.SelectBuilder
}

// //Build builds the sql string using current dialect into given bufer
// func (c *jHookDemoSelectBuilder) Build(b dbr.Buffer) error {
// 	return c.SelectBuilder.Build(runtime.GetDialect(), b)
// }
// //String returns the sql string for current dialect. It returns empty string if the build returns an error.
// func (c *jHookDemoSelectBuilder) String() string {
// 	b := dbr.NewBuffer()
// 	if err := c.Build(b); err != nil {
// 		return ""
// 	}
// 	return b.String()
// }

//Read evaluates current select query and load the results into a HookDemo
func (c *jHookDemoSelectBuilder) Read() (*HookDemo, error) {
	var one HookDemo
	err := c.LoadStruct(&one)
	return &one, err
}

//ReadAll evaluates current select query and load the results into a slice of HookDemo
func (c *jHookDemoSelectBuilder) ReadAll() ([]*HookDemo, error) {
	var all []*HookDemo
	_, err := c.LoadStructs(&all)
	return all, err
}

//Where returns a jHookDemoSelectBuilder instead of builder.SelectBuilder.
func (c *jHookDemoSelectBuilder) Where(query interface{}, value ...interface{}) *jHookDemoSelectBuilder {
	c.SelectBuilder.Where(query, value...)
	return c
}

//GroupBy returns a jHookDemoSelectBuilder instead of builder.SelectBuilder.
func (c *jHookDemoSelectBuilder) GroupBy(col ...string) *jHookDemoSelectBuilder {
	c.SelectBuilder.GroupBy(col...)
	return c
}

//Having returns a jHookDemoSelectBuilder instead of builder.SelectBuilder.
func (c *jHookDemoSelectBuilder) Having(query interface{}, value ...interface{}) *jHookDemoSelectBuilder {
	c.SelectBuilder.Having(query, value...)
	return c
}

//Limit returns a jHookDemoSelectBuilder instead of builder.SelectBuilder.
func (c *jHookDemoSelectBuilder) Limit(n uint64) *jHookDemoSelectBuilder {
	c.SelectBuilder.Limit(n)
	return c
}

//Offset returns a jHookDemoSelectBuilder instead of builder.SelectBuilder.
func (c *jHookDemoSelectBuilder) Offset(n uint64) *jHookDemoSelectBuilder {
	c.SelectBuilder.Offset(n)
	return c
}

//OrderAsc returns a jHookDemoSelectBuilder instead of builder.SelectBuilder.
func (c *jHookDemoSelectBuilder) OrderAsc(col string) *jHookDemoSelectBuilder {
	c.SelectBuilder.OrderAsc(col)
	return c
}

//OrderDesc returns a jHookDemoSelectBuilder instead of builder.SelectBuilder.
func (c *jHookDemoSelectBuilder) OrderDesc(col string) *jHookDemoSelectBuilder {
	c.SelectBuilder.OrderDesc(col)
	return c
}

//OrderDir returns a jHookDemoSelectBuilder instead of builder.SelectBuilder.
func (c *jHookDemoSelectBuilder) OrderDir(col string, isAsc bool) *jHookDemoSelectBuilder {
	c.SelectBuilder.OrderDir(col, isAsc)
	return c
}

//OrderBy returns a jHookDemoSelectBuilder instead of builder.SelectBuilder.
func (c *jHookDemoSelectBuilder) OrderBy(col string) *jHookDemoSelectBuilder {
	c.SelectBuilder.OrderBy(col)
	return c
}

//Paginate returns a jHookDemoSelectBuilder instead of builder.SelectBuilder.
func (c *jHookDemoSelectBuilder) Paginate(page, perPage uint64) *jHookDemoSelectBuilder {
	c.SelectBuilder.Paginate(page, perPage)
	return c
}

//Join returns a jHookDemoSelectBuilder instead of builder.SelectBuilder.
func (c *jHookDemoSelectBuilder) Join(table, on interface{}) *jHookDemoSelectBuilder {
	c.SelectBuilder.Join(table, on)
	return c
}

//LeftJoin returns a jHookDemoSelectBuilder instead of builder.SelectBuilder.
func (c *jHookDemoSelectBuilder) LeftJoin(table, on interface{}) *jHookDemoSelectBuilder {
	c.SelectBuilder.LeftJoin(table, on)
	return c
}

//RightJoin returns a jHookDemoSelectBuilder instead of builder.SelectBuilder.
func (c *jHookDemoSelectBuilder) RightJoin(table, on interface{}) *jHookDemoSelectBuilder {
	c.SelectBuilder.RightJoin(table, on)
	return c
}

//FullJoin returns a jHookDemoSelectBuilder instead of builder.SelectBuilder.
func (c *jHookDemoSelectBuilder) FullJoin(table, on interface{}) *jHookDemoSelectBuilder {
	c.SelectBuilder.FullJoin(table, on)
	return c
}

//Distinct returns a jHookDemoSelectBuilder instead of builder.SelectBuilder.
func (c *jHookDemoSelectBuilder) Distinct() *jHookDemoSelectBuilder {
	c.SelectBuilder.Distinct()
	return c
}

// JHookDemo provides a basic querier
func JHookDemo(db dbr.SessionRunner) jHookDemoQuerier {
	return jHookDemoQuerier{
		db: db,
	}
}

type jHookDemoQuerier struct {
	db dbr.SessionRunner
	as string
}

//As set alias prior building.
func (c jHookDemoQuerier) As(as string) jHookDemoQuerier {
	c.as = as
	return c
}

//Model returns a model with appropriate aliasing.
func (c jHookDemoQuerier) Model() jHookDemoModel {
	return JHookDemoModel.As(c.as)
}

//Select returns a HookDemo Select Builder.
func (c jHookDemoQuerier) Select(what ...string) *jHookDemoSelectBuilder {
	m := c.Model()
	dialect := runtime.GetDialect()
	from := dialect.QuoteIdent(m.Table())
	if m.Alias() != "" && m.Alias() != m.Table() {
		from = fmt.Sprintf("%v as %v", from, dialect.QuoteIdent(m.Alias()))
	}
	if len(what) == 0 {
		alias := m.Table()
		if m.Alias() != "" && m.Alias() != m.Table() {
			alias = m.Alias()
		}
		what = m.Fields(alias + ".*")
	}
	return &jHookDemoSelectBuilder{
		as: c.as,
		SelectBuilder: &builder.SelectBuilder{
			SelectBuilder: c.db.Select(what...).From(from),
		},
	}
}

//Where returns a HookDemo Select Builder.
func (c jHookDemoQuerier) Where(query interface{}, value ...interface{}) *jHookDemoSelectBuilder {
	return c.Select().Where(query, value...)
}

//Count returns a HookDemo Select Builder to count given expressions.
func (c jHookDemoQuerier) Count(what ...string) *jHookDemoSelectBuilder {
	if len(what) == 0 {
		what = append(what, "*")
	}
	return c.Select("COUNT(" + strings.Join(what, ",") + ")")
}

// Insert a new HookDemo, if it has autoincrement primary key, the value will be set.
// It stops on first error.
func (c jHookDemoQuerier) Insert(items ...*HookDemo) (sql.Result, error) {
	var res sql.Result
	var err error
	for _, data := range items {

		err = data.beforeInsert()
		if err != nil {
			return nil, err
		}

		query := c.db.InsertInto(JHookDemoModel.Table()).Columns(

			`name`,
		).Record(data)
		if runtime.Runs(drivers.Pgsql) {

			query = query.Returning(

				`id`,
			)

			var auto0 int64

			err = query.Load(

				&auto0,
			)

			data.ID = auto0

		} else {
			res, err = query.Exec()

			if err == nil {
				id, err2 := res.LastInsertId()
				if err2 != nil {
					return res, err2
				}
				data.ID = id
			}

		}
		if err != nil {
			return res, err
		}
	}
	return res, err
}

// InsertBulk inserts multiple items into the database.
// It does not post update any auto increment field.
// It builds an insert query of multiple rows and send it on the underlying connection.
func (c jHookDemoQuerier) InsertBulk(items ...*HookDemo) error {
	panic("todo")
}

// Update a HookDemo. It stops on first error.
func (c jHookDemoQuerier) Update(items ...*HookDemo) (sql.Result, error) {
	var res sql.Result
	var err error
	for _, data := range items {

		err = data.beforeUpdate()
		if err != nil {
			return nil, err
		}

		query := c.db.Update(JHookDemoModel.Table())

		query = query.Set(`name`, data.Name)

		query = query.Where("id = ?", data.ID)

		res, err = query.Exec()

		if err != nil {
			return res, err
		}
	}
	return res, err
}

// MustUpdate a HookDemo. It stops on first error. It errors if an update query does not affect row.
func (c jHookDemoQuerier) MustUpdate(items ...*HookDemo) (sql.Result, error) {
	var res sql.Result
	var err error
	for _, data := range items {

		err = data.beforeUpdate()
		if err != nil {
			return nil, err
		}

		query := c.db.Update(JHookDemoModel.Table())

		query = query.Set(`name`, data.Name)

		query = query.Where("id = ?", data.ID)

		res, err = query.Exec()

		if err == nil {
			if n, _ := res.RowsAffected(); n == 0 {
				x := &builder.UpdateBuilder{UpdateBuilder: query}
				err = runtime.NewNoRowsAffected(x.String())
			}
		}

		if err != nil {
			return res, err
		}
	}
	return res, err
}

// UpdateBulk updates multiple items into the database.
// It builds an update query of multiple rows and send it on the underlying connection.
func (c jHookDemoQuerier) UpdateBulk(items ...*HookDemo) error {
	panic("todo")
}

//Delete returns a delete builder
func (c jHookDemoQuerier) Delete() *jHookDemoDeleteBuilder {
	return &jHookDemoDeleteBuilder{
		&builder.DeleteBuilder{
			DeleteBuilder: c.db.DeleteFrom(JHookDemoModel.Table()),
		},
	}
}

// MustDelete requires the query to affeect rows.
func (c jHookDemoQuerier) MustDelete() *jHookDemoDeleteBuilder {
	ret := &jHookDemoDeleteBuilder{
		&builder.DeleteBuilder{
			DeleteBuilder: c.db.DeleteFrom(JHookDemoModel.Table()),
		},
	}
	ret.MustDelete()
	return ret
}

//DeleteByPk deletes one HookDemo by its PKs
func (c jHookDemoQuerier) DeleteByPk(ID int64) error {
	_, err := c.Delete().Where(

		JHookDemoModel.ID.Eq(ID),
	).Exec()
	return err
}

// DeleteAll given HookDemo
func (c jHookDemoQuerier) DeleteAll(items ...*HookDemo) (sql.Result, error) {
	q := c.Delete().Where(
		JHookDemoModel.In(items...),
	)
	return q.Exec()
}

// MustDeleteAll given HookDemo
func (c jHookDemoQuerier) MustDeleteAll(items ...*HookDemo) (sql.Result, error) {
	var res sql.Result
	var err error
	for _, d := range items {
		res, err = c.DeleteAll(d)
		if err != nil {
			return res, err
		}
		if n, e := res.RowsAffected(); e != nil {
			return res, e
		} else if n == 0 {
			q := c.Delete().Where(
				JHookDemoModel.In(items...),
			)
			err = runtime.NewNoRowsAffected(q.String())
			return res, err
		}
	}
	return res, err
}

//Find one HookDemo using its PKs
func (c jHookDemoQuerier) Find(ID int64) (*HookDemo, error) {
	return c.Select().Where(

		JHookDemoModel.ID.Eq(ID),
	).Read()
}

// JHasManyTextPkrelatedsToTextPkrelatedsSetup helps to create/drop the schema
func JHasManyTextPkrelatedsToTextPkrelatedsSetup() runtime.Setuper {
	driver := runtime.GetCurrentDriver()

	var create string
	var drop string

	if driver == drivers.Sqlite {
		create = `CREATE TABLE IF NOT EXISTS hasmanytextpk_relatedstotextpk_relateds (
has_many_text_pk_id INTEGER,
text_pk_name TEXT,
PRIMARY KEY (has_many_text_pk_id,text_pk_name) 

)`
	} else if driver == drivers.Mysql {
		create = `CREATE TABLE IF NOT EXISTS hasmanytextpk_relatedstotextpk_relateds (
has_many_text_pk_id INTEGER NOT NULL,
text_pk_name VARCHAR(255) NOT NULL,
PRIMARY KEY (has_many_text_pk_id,text_pk_name) 

)`
	} else if driver == drivers.Pgsql {
		create = `CREATE TABLE IF NOT EXISTS hasmanytextpk_relatedstotextpk_relateds (
has_many_text_pk_id INTEGER,
text_pk_name TEXT,
PRIMARY KEY (has_many_text_pk_id,text_pk_name) 

)`
	}

	if driver == drivers.Sqlite {
		drop = `DROP TABLE IF EXISTS hasmanytextpk_relatedstotextpk_relateds`
	} else if driver == drivers.Mysql {
		drop = `DROP TABLE IF EXISTS hasmanytextpk_relatedstotextpk_relateds`
	} else if driver == drivers.Pgsql {
		drop = `DROP TABLE IF EXISTS hasmanytextpk_relatedstotextpk_relateds`
	}

	var indexes []string

	if driver == drivers.Sqlite {

	} else if driver == drivers.Mysql {

	} else if driver == drivers.Pgsql {

	}

	return runtime.Table{
		Name:       `hasmanytextpk_relatedstotextpk_relateds`,
		CreateStmt: create,
		DropStmt:   drop,
		View:       !true,
		Indexes:    indexes,
	}
}

// jHasManyTextPkrelatedsToTextPkrelatedsModel provides helper to work with HasManyTextPkrelatedsToTextPkrelateds data provider
type jHasManyTextPkrelatedsToTextPkrelatedsModel struct {
	as string

	HasManyTextPkID builder.ValuePropertyMeta

	TextPkName builder.ValuePropertyMeta
}

// Eq provided items.
func (j jHasManyTextPkrelatedsToTextPkrelatedsModel) Eq(s ...*HasManyTextPkrelatedsToTextPkrelateds) dbr.Builder {
	ors := []dbr.Builder{}
	for _, t := range s {
		ors = append(ors, dbr.And(

			JHasManyTextPkrelatedsToTextPkrelatedsModel.HasManyTextPkID.Eq(t.HasManyTextPkID),

			JHasManyTextPkrelatedsToTextPkrelatedsModel.TextPkName.Eq(t.TextPkName),
		))
	}
	return dbr.Or(ors...)
}

// In provided items.
func (j jHasManyTextPkrelatedsToTextPkrelatedsModel) In(s ...*HasManyTextPkrelatedsToTextPkrelateds) dbr.Builder {
	ors := []dbr.Builder{}
	for _, t := range s {
		ors = append(ors, dbr.And(

			JHasManyTextPkrelatedsToTextPkrelatedsModel.HasManyTextPkID.Eq(t.HasManyTextPkID),

			JHasManyTextPkrelatedsToTextPkrelatedsModel.TextPkName.Eq(t.TextPkName),
		))
	}
	return dbr.Or(ors...)
}

// As returns a copy with an alias.
func (j jHasManyTextPkrelatedsToTextPkrelatedsModel) As(as string) jHasManyTextPkrelatedsToTextPkrelatedsModel {
	j.as = as

	j.HasManyTextPkID.TableAlias = as

	j.TextPkName.TableAlias = as

	return j
}

// Table returns the sql table name
func (j jHasManyTextPkrelatedsToTextPkrelatedsModel) Table() string {
	return "hasmanytextpk_relatedstotextpk_relateds"
}

// Alias returns the current alias
func (j jHasManyTextPkrelatedsToTextPkrelatedsModel) Alias() string {
	if j.as == "" {
		return j.Table()
	}
	return j.as
}

// Properties returns a map of property name => meta
func (j jHasManyTextPkrelatedsToTextPkrelatedsModel) Properties() map[string]builder.MetaProvider {
	ret := map[string]builder.MetaProvider{}

	ret["HasManyTextPkID"] = j.HasManyTextPkID

	ret["TextPkName"] = j.TextPkName

	return ret
}

// Fields returns given sql fields with appropriate aliasing.
func (j jHasManyTextPkrelatedsToTextPkrelatedsModel) Fields(ins ...string) []string {
	dialect := runtime.GetDialect()
	if len(ins) == 0 {
		ins = append(ins, "*")
	}
	for i, in := range ins {
		if j.as != "" {
			if in == "*" {
				ins[i] = fmt.Sprintf("%v.%v", dialect.QuoteIdent(j.as), in)
			} else {
				ins[i] = fmt.Sprintf("%v.%v", dialect.QuoteIdent(j.as), dialect.QuoteIdent(in))
			}
		}
	}
	return ins
}

// JHasManyTextPkrelatedsToTextPkrelatedsModel provides helper to work with HasManyTextPkrelatedsToTextPkrelateds data provider
var JHasManyTextPkrelatedsToTextPkrelatedsModel = jHasManyTextPkrelatedsToTextPkrelatedsModel{

	HasManyTextPkID: builder.NewValueMeta(
		`has_many_text_pk_id`, `INTEGER`,
		`HasManyTextPkID`, `int64`,
		true, false,
	),

	TextPkName: builder.NewValueMeta(
		`text_pk_name`, `TEXT`,
		`TextPkName`, `string`,
		true, false,
	),
}

// HasManyTextPkrelatedsToTextPkrelateds is automatically generated to handle a many to many relationship.
type HasManyTextPkrelatedsToTextPkrelateds struct {
	HasManyTextPkID int64

	TextPkName string
}

type jHasManyTextPkrelatedsToTextPkrelatedsDeleteBuilder struct {
	*builder.DeleteBuilder
}

// //Build builds the sql string into given buffer using current dialect
// func (c *jHasManyTextPkrelatedsToTextPkrelatedsDeleteBuilder) Build(b dbr.Buffer) error {
// 	return c.DeleteBuilder.Build(runtime.GetDialect(), b)
// }
// //String returns the sql string for current dialect. It returns empty string if the build returns an error.
// func (c *jHasManyTextPkrelatedsToTextPkrelatedsDeleteBuilder) String() string {
// 	b := dbr.NewBuffer()
// 	if err := c.Build(b); err != nil {
// 		return ""
// 	}
// 	return b.String()
// }
//Where returns a jHasManyTextPkrelatedsToTextPkrelatedsDeleteBuilder instead of builder.DeleteBuilder.
func (c *jHasManyTextPkrelatedsToTextPkrelatedsDeleteBuilder) Where(query interface{}, value ...interface{}) *jHasManyTextPkrelatedsToTextPkrelatedsDeleteBuilder {
	c.DeleteBuilder.Where(query, value...)
	return c
}

type jHasManyTextPkrelatedsToTextPkrelatedsSelectBuilder struct {
	as string
	*builder.SelectBuilder
}

// //Build builds the sql string using current dialect into given bufer
// func (c *jHasManyTextPkrelatedsToTextPkrelatedsSelectBuilder) Build(b dbr.Buffer) error {
// 	return c.SelectBuilder.Build(runtime.GetDialect(), b)
// }
// //String returns the sql string for current dialect. It returns empty string if the build returns an error.
// func (c *jHasManyTextPkrelatedsToTextPkrelatedsSelectBuilder) String() string {
// 	b := dbr.NewBuffer()
// 	if err := c.Build(b); err != nil {
// 		return ""
// 	}
// 	return b.String()
// }

//Read evaluates current select query and load the results into a HasManyTextPkrelatedsToTextPkrelateds
func (c *jHasManyTextPkrelatedsToTextPkrelatedsSelectBuilder) Read() (*HasManyTextPkrelatedsToTextPkrelateds, error) {
	var one HasManyTextPkrelatedsToTextPkrelateds
	err := c.LoadStruct(&one)
	return &one, err
}

//ReadAll evaluates current select query and load the results into a slice of HasManyTextPkrelatedsToTextPkrelateds
func (c *jHasManyTextPkrelatedsToTextPkrelatedsSelectBuilder) ReadAll() ([]*HasManyTextPkrelatedsToTextPkrelateds, error) {
	var all []*HasManyTextPkrelatedsToTextPkrelateds
	_, err := c.LoadStructs(&all)
	return all, err
}

//Where returns a jHasManyTextPkrelatedsToTextPkrelatedsSelectBuilder instead of builder.SelectBuilder.
func (c *jHasManyTextPkrelatedsToTextPkrelatedsSelectBuilder) Where(query interface{}, value ...interface{}) *jHasManyTextPkrelatedsToTextPkrelatedsSelectBuilder {
	c.SelectBuilder.Where(query, value...)
	return c
}

//GroupBy returns a jHasManyTextPkrelatedsToTextPkrelatedsSelectBuilder instead of builder.SelectBuilder.
func (c *jHasManyTextPkrelatedsToTextPkrelatedsSelectBuilder) GroupBy(col ...string) *jHasManyTextPkrelatedsToTextPkrelatedsSelectBuilder {
	c.SelectBuilder.GroupBy(col...)
	return c
}

//Having returns a jHasManyTextPkrelatedsToTextPkrelatedsSelectBuilder instead of builder.SelectBuilder.
func (c *jHasManyTextPkrelatedsToTextPkrelatedsSelectBuilder) Having(query interface{}, value ...interface{}) *jHasManyTextPkrelatedsToTextPkrelatedsSelectBuilder {
	c.SelectBuilder.Having(query, value...)
	return c
}

//Limit returns a jHasManyTextPkrelatedsToTextPkrelatedsSelectBuilder instead of builder.SelectBuilder.
func (c *jHasManyTextPkrelatedsToTextPkrelatedsSelectBuilder) Limit(n uint64) *jHasManyTextPkrelatedsToTextPkrelatedsSelectBuilder {
	c.SelectBuilder.Limit(n)
	return c
}

//Offset returns a jHasManyTextPkrelatedsToTextPkrelatedsSelectBuilder instead of builder.SelectBuilder.
func (c *jHasManyTextPkrelatedsToTextPkrelatedsSelectBuilder) Offset(n uint64) *jHasManyTextPkrelatedsToTextPkrelatedsSelectBuilder {
	c.SelectBuilder.Offset(n)
	return c
}

//OrderAsc returns a jHasManyTextPkrelatedsToTextPkrelatedsSelectBuilder instead of builder.SelectBuilder.
func (c *jHasManyTextPkrelatedsToTextPkrelatedsSelectBuilder) OrderAsc(col string) *jHasManyTextPkrelatedsToTextPkrelatedsSelectBuilder {
	c.SelectBuilder.OrderAsc(col)
	return c
}

//OrderDesc returns a jHasManyTextPkrelatedsToTextPkrelatedsSelectBuilder instead of builder.SelectBuilder.
func (c *jHasManyTextPkrelatedsToTextPkrelatedsSelectBuilder) OrderDesc(col string) *jHasManyTextPkrelatedsToTextPkrelatedsSelectBuilder {
	c.SelectBuilder.OrderDesc(col)
	return c
}

//OrderDir returns a jHasManyTextPkrelatedsToTextPkrelatedsSelectBuilder instead of builder.SelectBuilder.
func (c *jHasManyTextPkrelatedsToTextPkrelatedsSelectBuilder) OrderDir(col string, isAsc bool) *jHasManyTextPkrelatedsToTextPkrelatedsSelectBuilder {
	c.SelectBuilder.OrderDir(col, isAsc)
	return c
}

//OrderBy returns a jHasManyTextPkrelatedsToTextPkrelatedsSelectBuilder instead of builder.SelectBuilder.
func (c *jHasManyTextPkrelatedsToTextPkrelatedsSelectBuilder) OrderBy(col string) *jHasManyTextPkrelatedsToTextPkrelatedsSelectBuilder {
	c.SelectBuilder.OrderBy(col)
	return c
}

//Paginate returns a jHasManyTextPkrelatedsToTextPkrelatedsSelectBuilder instead of builder.SelectBuilder.
func (c *jHasManyTextPkrelatedsToTextPkrelatedsSelectBuilder) Paginate(page, perPage uint64) *jHasManyTextPkrelatedsToTextPkrelatedsSelectBuilder {
	c.SelectBuilder.Paginate(page, perPage)
	return c
}

//Join returns a jHasManyTextPkrelatedsToTextPkrelatedsSelectBuilder instead of builder.SelectBuilder.
func (c *jHasManyTextPkrelatedsToTextPkrelatedsSelectBuilder) Join(table, on interface{}) *jHasManyTextPkrelatedsToTextPkrelatedsSelectBuilder {
	c.SelectBuilder.Join(table, on)
	return c
}

//LeftJoin returns a jHasManyTextPkrelatedsToTextPkrelatedsSelectBuilder instead of builder.SelectBuilder.
func (c *jHasManyTextPkrelatedsToTextPkrelatedsSelectBuilder) LeftJoin(table, on interface{}) *jHasManyTextPkrelatedsToTextPkrelatedsSelectBuilder {
	c.SelectBuilder.LeftJoin(table, on)
	return c
}

//RightJoin returns a jHasManyTextPkrelatedsToTextPkrelatedsSelectBuilder instead of builder.SelectBuilder.
func (c *jHasManyTextPkrelatedsToTextPkrelatedsSelectBuilder) RightJoin(table, on interface{}) *jHasManyTextPkrelatedsToTextPkrelatedsSelectBuilder {
	c.SelectBuilder.RightJoin(table, on)
	return c
}

//FullJoin returns a jHasManyTextPkrelatedsToTextPkrelatedsSelectBuilder instead of builder.SelectBuilder.
func (c *jHasManyTextPkrelatedsToTextPkrelatedsSelectBuilder) FullJoin(table, on interface{}) *jHasManyTextPkrelatedsToTextPkrelatedsSelectBuilder {
	c.SelectBuilder.FullJoin(table, on)
	return c
}

//Distinct returns a jHasManyTextPkrelatedsToTextPkrelatedsSelectBuilder instead of builder.SelectBuilder.
func (c *jHasManyTextPkrelatedsToTextPkrelatedsSelectBuilder) Distinct() *jHasManyTextPkrelatedsToTextPkrelatedsSelectBuilder {
	c.SelectBuilder.Distinct()
	return c
}

// JHasManyTextPkrelatedsToTextPkrelateds provides a basic querier
func JHasManyTextPkrelatedsToTextPkrelateds(db dbr.SessionRunner) jHasManyTextPkrelatedsToTextPkrelatedsQuerier {
	return jHasManyTextPkrelatedsToTextPkrelatedsQuerier{
		db: db,
	}
}

type jHasManyTextPkrelatedsToTextPkrelatedsQuerier struct {
	db dbr.SessionRunner
	as string
}

//As set alias prior building.
func (c jHasManyTextPkrelatedsToTextPkrelatedsQuerier) As(as string) jHasManyTextPkrelatedsToTextPkrelatedsQuerier {
	c.as = as
	return c
}

//Model returns a model with appropriate aliasing.
func (c jHasManyTextPkrelatedsToTextPkrelatedsQuerier) Model() jHasManyTextPkrelatedsToTextPkrelatedsModel {
	return JHasManyTextPkrelatedsToTextPkrelatedsModel.As(c.as)
}

//Select returns a HasManyTextPkrelatedsToTextPkrelateds Select Builder.
func (c jHasManyTextPkrelatedsToTextPkrelatedsQuerier) Select(what ...string) *jHasManyTextPkrelatedsToTextPkrelatedsSelectBuilder {
	m := c.Model()
	dialect := runtime.GetDialect()
	from := dialect.QuoteIdent(m.Table())
	if m.Alias() != "" && m.Alias() != m.Table() {
		from = fmt.Sprintf("%v as %v", from, dialect.QuoteIdent(m.Alias()))
	}
	if len(what) == 0 {
		alias := m.Table()
		if m.Alias() != "" && m.Alias() != m.Table() {
			alias = m.Alias()
		}
		what = m.Fields(alias + ".*")
	}
	return &jHasManyTextPkrelatedsToTextPkrelatedsSelectBuilder{
		as: c.as,
		SelectBuilder: &builder.SelectBuilder{
			SelectBuilder: c.db.Select(what...).From(from),
		},
	}
}

//Where returns a HasManyTextPkrelatedsToTextPkrelateds Select Builder.
func (c jHasManyTextPkrelatedsToTextPkrelatedsQuerier) Where(query interface{}, value ...interface{}) *jHasManyTextPkrelatedsToTextPkrelatedsSelectBuilder {
	return c.Select().Where(query, value...)
}

//Count returns a HasManyTextPkrelatedsToTextPkrelateds Select Builder to count given expressions.
func (c jHasManyTextPkrelatedsToTextPkrelatedsQuerier) Count(what ...string) *jHasManyTextPkrelatedsToTextPkrelatedsSelectBuilder {
	if len(what) == 0 {
		what = append(what, "*")
	}
	return c.Select("COUNT(" + strings.Join(what, ",") + ")")
}

// Insert a new HasManyTextPkrelatedsToTextPkrelateds, if it has autoincrement primary key, the value will be set.
// It stops on first error.
func (c jHasManyTextPkrelatedsToTextPkrelatedsQuerier) Insert(items ...*HasManyTextPkrelatedsToTextPkrelateds) (sql.Result, error) {
	var res sql.Result
	var err error
	for _, data := range items {

		if runtime.Runs(drivers.Mysql) {

			if len(data.TextPkName) > 255 {
				return nil, fmt.Errorf("TextPkName: PRIMARY KEY length exceeded max=255, got=%v", len(data.TextPkName))
			}

		}

		query := c.db.InsertInto(JHasManyTextPkrelatedsToTextPkrelatedsModel.Table()).Columns(

			`has_many_text_pk_id`,

			`text_pk_name`,
		).Record(data)
		if runtime.Runs(drivers.Pgsql) {

			res, err = query.Exec()

		} else {
			res, err = query.Exec()

		}
		if err != nil {
			return res, err
		}
	}
	return res, err
}

// InsertBulk inserts multiple items into the database.
// It does not post update any auto increment field.
// It builds an insert query of multiple rows and send it on the underlying connection.
func (c jHasManyTextPkrelatedsToTextPkrelatedsQuerier) InsertBulk(items ...*HasManyTextPkrelatedsToTextPkrelateds) error {
	panic("todo")
}

//Delete returns a delete builder
func (c jHasManyTextPkrelatedsToTextPkrelatedsQuerier) Delete() *jHasManyTextPkrelatedsToTextPkrelatedsDeleteBuilder {
	return &jHasManyTextPkrelatedsToTextPkrelatedsDeleteBuilder{
		&builder.DeleteBuilder{
			DeleteBuilder: c.db.DeleteFrom(JHasManyTextPkrelatedsToTextPkrelatedsModel.Table()),
		},
	}
}

// MustDelete requires the query to affeect rows.
func (c jHasManyTextPkrelatedsToTextPkrelatedsQuerier) MustDelete() *jHasManyTextPkrelatedsToTextPkrelatedsDeleteBuilder {
	ret := &jHasManyTextPkrelatedsToTextPkrelatedsDeleteBuilder{
		&builder.DeleteBuilder{
			DeleteBuilder: c.db.DeleteFrom(JHasManyTextPkrelatedsToTextPkrelatedsModel.Table()),
		},
	}
	ret.MustDelete()
	return ret
}

//DeleteByPk deletes one HasManyTextPkrelatedsToTextPkrelateds by its PKs
func (c jHasManyTextPkrelatedsToTextPkrelatedsQuerier) DeleteByPk(HasManyTextPkID int64, TextPkName string) error {
	_, err := c.Delete().Where(

		JHasManyTextPkrelatedsToTextPkrelatedsModel.HasManyTextPkID.Eq(HasManyTextPkID),

		JHasManyTextPkrelatedsToTextPkrelatedsModel.TextPkName.Eq(TextPkName),
	).Exec()
	return err
}

// DeleteAll given HasManyTextPkrelatedsToTextPkrelateds
func (c jHasManyTextPkrelatedsToTextPkrelatedsQuerier) DeleteAll(items ...*HasManyTextPkrelatedsToTextPkrelateds) (sql.Result, error) {
	q := c.Delete().Where(
		JHasManyTextPkrelatedsToTextPkrelatedsModel.In(items...),
	)
	return q.Exec()
}

// MustDeleteAll given HasManyTextPkrelatedsToTextPkrelateds
func (c jHasManyTextPkrelatedsToTextPkrelatedsQuerier) MustDeleteAll(items ...*HasManyTextPkrelatedsToTextPkrelateds) (sql.Result, error) {
	var res sql.Result
	var err error
	for _, d := range items {
		res, err = c.DeleteAll(d)
		if err != nil {
			return res, err
		}
		if n, e := res.RowsAffected(); e != nil {
			return res, e
		} else if n == 0 {
			q := c.Delete().Where(
				JHasManyTextPkrelatedsToTextPkrelatedsModel.In(items...),
			)
			err = runtime.NewNoRowsAffected(q.String())
			return res, err
		}
	}
	return res, err
}

//Find one HasManyTextPkrelatedsToTextPkrelateds using its PKs
func (c jHasManyTextPkrelatedsToTextPkrelatedsQuerier) Find(HasManyTextPkID int64, TextPkName string) (*HasManyTextPkrelatedsToTextPkrelateds, error) {
	return c.Select().Where(

		JHasManyTextPkrelatedsToTextPkrelatedsModel.HasManyTextPkID.Eq(HasManyTextPkID),

		JHasManyTextPkrelatedsToTextPkrelatedsModel.TextPkName.Eq(TextPkName),
	).Read()
}

// JCompositePkrelatedsToHasManyCompositePkrelatedsSetup helps to create/drop the schema
func JCompositePkrelatedsToHasManyCompositePkrelatedsSetup() runtime.Setuper {
	driver := runtime.GetCurrentDriver()

	var create string
	var drop string

	if driver == drivers.Sqlite {
		create = `CREATE TABLE IF NOT EXISTS compositepk_relatedstohasmanycompositepk_relateds (
has_many_composite_pk_id INTEGER,
composite_pk_p TEXT,
composite_pk_k TEXT,
PRIMARY KEY (has_many_composite_pk_id,composite_pk_p,composite_pk_k) 

)`
	} else if driver == drivers.Mysql {
		create = `CREATE TABLE IF NOT EXISTS compositepk_relatedstohasmanycompositepk_relateds (
has_many_composite_pk_id INTEGER NOT NULL,
composite_pk_p VARCHAR(255) NOT NULL,
composite_pk_k VARCHAR(255) NOT NULL,
PRIMARY KEY (has_many_composite_pk_id,composite_pk_p,composite_pk_k) 

)`
	} else if driver == drivers.Pgsql {
		create = `CREATE TABLE IF NOT EXISTS compositepk_relatedstohasmanycompositepk_relateds (
has_many_composite_pk_id INTEGER,
composite_pk_p TEXT,
composite_pk_k TEXT,
PRIMARY KEY (has_many_composite_pk_id,composite_pk_p,composite_pk_k) 

)`
	}

	if driver == drivers.Sqlite {
		drop = `DROP TABLE IF EXISTS compositepk_relatedstohasmanycompositepk_relateds`
	} else if driver == drivers.Mysql {
		drop = `DROP TABLE IF EXISTS compositepk_relatedstohasmanycompositepk_relateds`
	} else if driver == drivers.Pgsql {
		drop = `DROP TABLE IF EXISTS compositepk_relatedstohasmanycompositepk_relateds`
	}

	var indexes []string

	if driver == drivers.Sqlite {

	} else if driver == drivers.Mysql {

	} else if driver == drivers.Pgsql {

	}

	return runtime.Table{
		Name:       `compositepk_relatedstohasmanycompositepk_relateds`,
		CreateStmt: create,
		DropStmt:   drop,
		View:       !true,
		Indexes:    indexes,
	}
}

// jCompositePkrelatedsToHasManyCompositePkrelatedsModel provides helper to work with CompositePkrelatedsToHasManyCompositePkrelateds data provider
type jCompositePkrelatedsToHasManyCompositePkrelatedsModel struct {
	as string

	HasManyCompositePkID builder.ValuePropertyMeta

	CompositePkP builder.ValuePropertyMeta

	CompositePkK builder.ValuePropertyMeta
}

// Eq provided items.
func (j jCompositePkrelatedsToHasManyCompositePkrelatedsModel) Eq(s ...*CompositePkrelatedsToHasManyCompositePkrelateds) dbr.Builder {
	ors := []dbr.Builder{}
	for _, t := range s {
		ors = append(ors, dbr.And(

			JCompositePkrelatedsToHasManyCompositePkrelatedsModel.HasManyCompositePkID.Eq(t.HasManyCompositePkID),

			JCompositePkrelatedsToHasManyCompositePkrelatedsModel.CompositePkP.Eq(t.CompositePkP),

			JCompositePkrelatedsToHasManyCompositePkrelatedsModel.CompositePkK.Eq(t.CompositePkK),
		))
	}
	return dbr.Or(ors...)
}

// In provided items.
func (j jCompositePkrelatedsToHasManyCompositePkrelatedsModel) In(s ...*CompositePkrelatedsToHasManyCompositePkrelateds) dbr.Builder {
	ors := []dbr.Builder{}
	for _, t := range s {
		ors = append(ors, dbr.And(

			JCompositePkrelatedsToHasManyCompositePkrelatedsModel.HasManyCompositePkID.Eq(t.HasManyCompositePkID),

			JCompositePkrelatedsToHasManyCompositePkrelatedsModel.CompositePkP.Eq(t.CompositePkP),

			JCompositePkrelatedsToHasManyCompositePkrelatedsModel.CompositePkK.Eq(t.CompositePkK),
		))
	}
	return dbr.Or(ors...)
}

// As returns a copy with an alias.
func (j jCompositePkrelatedsToHasManyCompositePkrelatedsModel) As(as string) jCompositePkrelatedsToHasManyCompositePkrelatedsModel {
	j.as = as

	j.HasManyCompositePkID.TableAlias = as

	j.CompositePkP.TableAlias = as

	j.CompositePkK.TableAlias = as

	return j
}

// Table returns the sql table name
func (j jCompositePkrelatedsToHasManyCompositePkrelatedsModel) Table() string {
	return "compositepk_relatedstohasmanycompositepk_relateds"
}

// Alias returns the current alias
func (j jCompositePkrelatedsToHasManyCompositePkrelatedsModel) Alias() string {
	if j.as == "" {
		return j.Table()
	}
	return j.as
}

// Properties returns a map of property name => meta
func (j jCompositePkrelatedsToHasManyCompositePkrelatedsModel) Properties() map[string]builder.MetaProvider {
	ret := map[string]builder.MetaProvider{}

	ret["HasManyCompositePkID"] = j.HasManyCompositePkID

	ret["CompositePkP"] = j.CompositePkP

	ret["CompositePkK"] = j.CompositePkK

	return ret
}

// Fields returns given sql fields with appropriate aliasing.
func (j jCompositePkrelatedsToHasManyCompositePkrelatedsModel) Fields(ins ...string) []string {
	dialect := runtime.GetDialect()
	if len(ins) == 0 {
		ins = append(ins, "*")
	}
	for i, in := range ins {
		if j.as != "" {
			if in == "*" {
				ins[i] = fmt.Sprintf("%v.%v", dialect.QuoteIdent(j.as), in)
			} else {
				ins[i] = fmt.Sprintf("%v.%v", dialect.QuoteIdent(j.as), dialect.QuoteIdent(in))
			}
		}
	}
	return ins
}

// JCompositePkrelatedsToHasManyCompositePkrelatedsModel provides helper to work with CompositePkrelatedsToHasManyCompositePkrelateds data provider
var JCompositePkrelatedsToHasManyCompositePkrelatedsModel = jCompositePkrelatedsToHasManyCompositePkrelatedsModel{

	HasManyCompositePkID: builder.NewValueMeta(
		`has_many_composite_pk_id`, `INTEGER`,
		`HasManyCompositePkID`, `int64`,
		true, false,
	),

	CompositePkP: builder.NewValueMeta(
		`composite_pk_p`, `TEXT`,
		`CompositePkP`, `string`,
		true, false,
	),

	CompositePkK: builder.NewValueMeta(
		`composite_pk_k`, `TEXT`,
		`CompositePkK`, `string`,
		true, false,
	),
}

// CompositePkrelatedsToHasManyCompositePkrelateds is automatically generated to handle a many to many relationship.
type CompositePkrelatedsToHasManyCompositePkrelateds struct {
	HasManyCompositePkID int64

	CompositePkP string

	CompositePkK string
}

type jCompositePkrelatedsToHasManyCompositePkrelatedsDeleteBuilder struct {
	*builder.DeleteBuilder
}

// //Build builds the sql string into given buffer using current dialect
// func (c *jCompositePkrelatedsToHasManyCompositePkrelatedsDeleteBuilder) Build(b dbr.Buffer) error {
// 	return c.DeleteBuilder.Build(runtime.GetDialect(), b)
// }
// //String returns the sql string for current dialect. It returns empty string if the build returns an error.
// func (c *jCompositePkrelatedsToHasManyCompositePkrelatedsDeleteBuilder) String() string {
// 	b := dbr.NewBuffer()
// 	if err := c.Build(b); err != nil {
// 		return ""
// 	}
// 	return b.String()
// }
//Where returns a jCompositePkrelatedsToHasManyCompositePkrelatedsDeleteBuilder instead of builder.DeleteBuilder.
func (c *jCompositePkrelatedsToHasManyCompositePkrelatedsDeleteBuilder) Where(query interface{}, value ...interface{}) *jCompositePkrelatedsToHasManyCompositePkrelatedsDeleteBuilder {
	c.DeleteBuilder.Where(query, value...)
	return c
}

type jCompositePkrelatedsToHasManyCompositePkrelatedsSelectBuilder struct {
	as string
	*builder.SelectBuilder
}

// //Build builds the sql string using current dialect into given bufer
// func (c *jCompositePkrelatedsToHasManyCompositePkrelatedsSelectBuilder) Build(b dbr.Buffer) error {
// 	return c.SelectBuilder.Build(runtime.GetDialect(), b)
// }
// //String returns the sql string for current dialect. It returns empty string if the build returns an error.
// func (c *jCompositePkrelatedsToHasManyCompositePkrelatedsSelectBuilder) String() string {
// 	b := dbr.NewBuffer()
// 	if err := c.Build(b); err != nil {
// 		return ""
// 	}
// 	return b.String()
// }

//Read evaluates current select query and load the results into a CompositePkrelatedsToHasManyCompositePkrelateds
func (c *jCompositePkrelatedsToHasManyCompositePkrelatedsSelectBuilder) Read() (*CompositePkrelatedsToHasManyCompositePkrelateds, error) {
	var one CompositePkrelatedsToHasManyCompositePkrelateds
	err := c.LoadStruct(&one)
	return &one, err
}

//ReadAll evaluates current select query and load the results into a slice of CompositePkrelatedsToHasManyCompositePkrelateds
func (c *jCompositePkrelatedsToHasManyCompositePkrelatedsSelectBuilder) ReadAll() ([]*CompositePkrelatedsToHasManyCompositePkrelateds, error) {
	var all []*CompositePkrelatedsToHasManyCompositePkrelateds
	_, err := c.LoadStructs(&all)
	return all, err
}

//Where returns a jCompositePkrelatedsToHasManyCompositePkrelatedsSelectBuilder instead of builder.SelectBuilder.
func (c *jCompositePkrelatedsToHasManyCompositePkrelatedsSelectBuilder) Where(query interface{}, value ...interface{}) *jCompositePkrelatedsToHasManyCompositePkrelatedsSelectBuilder {
	c.SelectBuilder.Where(query, value...)
	return c
}

//GroupBy returns a jCompositePkrelatedsToHasManyCompositePkrelatedsSelectBuilder instead of builder.SelectBuilder.
func (c *jCompositePkrelatedsToHasManyCompositePkrelatedsSelectBuilder) GroupBy(col ...string) *jCompositePkrelatedsToHasManyCompositePkrelatedsSelectBuilder {
	c.SelectBuilder.GroupBy(col...)
	return c
}

//Having returns a jCompositePkrelatedsToHasManyCompositePkrelatedsSelectBuilder instead of builder.SelectBuilder.
func (c *jCompositePkrelatedsToHasManyCompositePkrelatedsSelectBuilder) Having(query interface{}, value ...interface{}) *jCompositePkrelatedsToHasManyCompositePkrelatedsSelectBuilder {
	c.SelectBuilder.Having(query, value...)
	return c
}

//Limit returns a jCompositePkrelatedsToHasManyCompositePkrelatedsSelectBuilder instead of builder.SelectBuilder.
func (c *jCompositePkrelatedsToHasManyCompositePkrelatedsSelectBuilder) Limit(n uint64) *jCompositePkrelatedsToHasManyCompositePkrelatedsSelectBuilder {
	c.SelectBuilder.Limit(n)
	return c
}

//Offset returns a jCompositePkrelatedsToHasManyCompositePkrelatedsSelectBuilder instead of builder.SelectBuilder.
func (c *jCompositePkrelatedsToHasManyCompositePkrelatedsSelectBuilder) Offset(n uint64) *jCompositePkrelatedsToHasManyCompositePkrelatedsSelectBuilder {
	c.SelectBuilder.Offset(n)
	return c
}

//OrderAsc returns a jCompositePkrelatedsToHasManyCompositePkrelatedsSelectBuilder instead of builder.SelectBuilder.
func (c *jCompositePkrelatedsToHasManyCompositePkrelatedsSelectBuilder) OrderAsc(col string) *jCompositePkrelatedsToHasManyCompositePkrelatedsSelectBuilder {
	c.SelectBuilder.OrderAsc(col)
	return c
}

//OrderDesc returns a jCompositePkrelatedsToHasManyCompositePkrelatedsSelectBuilder instead of builder.SelectBuilder.
func (c *jCompositePkrelatedsToHasManyCompositePkrelatedsSelectBuilder) OrderDesc(col string) *jCompositePkrelatedsToHasManyCompositePkrelatedsSelectBuilder {
	c.SelectBuilder.OrderDesc(col)
	return c
}

//OrderDir returns a jCompositePkrelatedsToHasManyCompositePkrelatedsSelectBuilder instead of builder.SelectBuilder.
func (c *jCompositePkrelatedsToHasManyCompositePkrelatedsSelectBuilder) OrderDir(col string, isAsc bool) *jCompositePkrelatedsToHasManyCompositePkrelatedsSelectBuilder {
	c.SelectBuilder.OrderDir(col, isAsc)
	return c
}

//OrderBy returns a jCompositePkrelatedsToHasManyCompositePkrelatedsSelectBuilder instead of builder.SelectBuilder.
func (c *jCompositePkrelatedsToHasManyCompositePkrelatedsSelectBuilder) OrderBy(col string) *jCompositePkrelatedsToHasManyCompositePkrelatedsSelectBuilder {
	c.SelectBuilder.OrderBy(col)
	return c
}

//Paginate returns a jCompositePkrelatedsToHasManyCompositePkrelatedsSelectBuilder instead of builder.SelectBuilder.
func (c *jCompositePkrelatedsToHasManyCompositePkrelatedsSelectBuilder) Paginate(page, perPage uint64) *jCompositePkrelatedsToHasManyCompositePkrelatedsSelectBuilder {
	c.SelectBuilder.Paginate(page, perPage)
	return c
}

//Join returns a jCompositePkrelatedsToHasManyCompositePkrelatedsSelectBuilder instead of builder.SelectBuilder.
func (c *jCompositePkrelatedsToHasManyCompositePkrelatedsSelectBuilder) Join(table, on interface{}) *jCompositePkrelatedsToHasManyCompositePkrelatedsSelectBuilder {
	c.SelectBuilder.Join(table, on)
	return c
}

//LeftJoin returns a jCompositePkrelatedsToHasManyCompositePkrelatedsSelectBuilder instead of builder.SelectBuilder.
func (c *jCompositePkrelatedsToHasManyCompositePkrelatedsSelectBuilder) LeftJoin(table, on interface{}) *jCompositePkrelatedsToHasManyCompositePkrelatedsSelectBuilder {
	c.SelectBuilder.LeftJoin(table, on)
	return c
}

//RightJoin returns a jCompositePkrelatedsToHasManyCompositePkrelatedsSelectBuilder instead of builder.SelectBuilder.
func (c *jCompositePkrelatedsToHasManyCompositePkrelatedsSelectBuilder) RightJoin(table, on interface{}) *jCompositePkrelatedsToHasManyCompositePkrelatedsSelectBuilder {
	c.SelectBuilder.RightJoin(table, on)
	return c
}

//FullJoin returns a jCompositePkrelatedsToHasManyCompositePkrelatedsSelectBuilder instead of builder.SelectBuilder.
func (c *jCompositePkrelatedsToHasManyCompositePkrelatedsSelectBuilder) FullJoin(table, on interface{}) *jCompositePkrelatedsToHasManyCompositePkrelatedsSelectBuilder {
	c.SelectBuilder.FullJoin(table, on)
	return c
}

//Distinct returns a jCompositePkrelatedsToHasManyCompositePkrelatedsSelectBuilder instead of builder.SelectBuilder.
func (c *jCompositePkrelatedsToHasManyCompositePkrelatedsSelectBuilder) Distinct() *jCompositePkrelatedsToHasManyCompositePkrelatedsSelectBuilder {
	c.SelectBuilder.Distinct()
	return c
}

// JCompositePkrelatedsToHasManyCompositePkrelateds provides a basic querier
func JCompositePkrelatedsToHasManyCompositePkrelateds(db dbr.SessionRunner) jCompositePkrelatedsToHasManyCompositePkrelatedsQuerier {
	return jCompositePkrelatedsToHasManyCompositePkrelatedsQuerier{
		db: db,
	}
}

type jCompositePkrelatedsToHasManyCompositePkrelatedsQuerier struct {
	db dbr.SessionRunner
	as string
}

//As set alias prior building.
func (c jCompositePkrelatedsToHasManyCompositePkrelatedsQuerier) As(as string) jCompositePkrelatedsToHasManyCompositePkrelatedsQuerier {
	c.as = as
	return c
}

//Model returns a model with appropriate aliasing.
func (c jCompositePkrelatedsToHasManyCompositePkrelatedsQuerier) Model() jCompositePkrelatedsToHasManyCompositePkrelatedsModel {
	return JCompositePkrelatedsToHasManyCompositePkrelatedsModel.As(c.as)
}

//Select returns a CompositePkrelatedsToHasManyCompositePkrelateds Select Builder.
func (c jCompositePkrelatedsToHasManyCompositePkrelatedsQuerier) Select(what ...string) *jCompositePkrelatedsToHasManyCompositePkrelatedsSelectBuilder {
	m := c.Model()
	dialect := runtime.GetDialect()
	from := dialect.QuoteIdent(m.Table())
	if m.Alias() != "" && m.Alias() != m.Table() {
		from = fmt.Sprintf("%v as %v", from, dialect.QuoteIdent(m.Alias()))
	}
	if len(what) == 0 {
		alias := m.Table()
		if m.Alias() != "" && m.Alias() != m.Table() {
			alias = m.Alias()
		}
		what = m.Fields(alias + ".*")
	}
	return &jCompositePkrelatedsToHasManyCompositePkrelatedsSelectBuilder{
		as: c.as,
		SelectBuilder: &builder.SelectBuilder{
			SelectBuilder: c.db.Select(what...).From(from),
		},
	}
}

//Where returns a CompositePkrelatedsToHasManyCompositePkrelateds Select Builder.
func (c jCompositePkrelatedsToHasManyCompositePkrelatedsQuerier) Where(query interface{}, value ...interface{}) *jCompositePkrelatedsToHasManyCompositePkrelatedsSelectBuilder {
	return c.Select().Where(query, value...)
}

//Count returns a CompositePkrelatedsToHasManyCompositePkrelateds Select Builder to count given expressions.
func (c jCompositePkrelatedsToHasManyCompositePkrelatedsQuerier) Count(what ...string) *jCompositePkrelatedsToHasManyCompositePkrelatedsSelectBuilder {
	if len(what) == 0 {
		what = append(what, "*")
	}
	return c.Select("COUNT(" + strings.Join(what, ",") + ")")
}

// Insert a new CompositePkrelatedsToHasManyCompositePkrelateds, if it has autoincrement primary key, the value will be set.
// It stops on first error.
func (c jCompositePkrelatedsToHasManyCompositePkrelatedsQuerier) Insert(items ...*CompositePkrelatedsToHasManyCompositePkrelateds) (sql.Result, error) {
	var res sql.Result
	var err error
	for _, data := range items {

		if runtime.Runs(drivers.Mysql) {

			if len(data.CompositePkP) > 255 {
				return nil, fmt.Errorf("CompositePkP: PRIMARY KEY length exceeded max=255, got=%v", len(data.CompositePkP))
			}

		}

		if runtime.Runs(drivers.Mysql) {

			if len(data.CompositePkK) > 255 {
				return nil, fmt.Errorf("CompositePkK: PRIMARY KEY length exceeded max=255, got=%v", len(data.CompositePkK))
			}

		}

		query := c.db.InsertInto(JCompositePkrelatedsToHasManyCompositePkrelatedsModel.Table()).Columns(

			`has_many_composite_pk_id`,

			`composite_pk_p`,

			`composite_pk_k`,
		).Record(data)
		if runtime.Runs(drivers.Pgsql) {

			res, err = query.Exec()

		} else {
			res, err = query.Exec()

		}
		if err != nil {
			return res, err
		}
	}
	return res, err
}

// InsertBulk inserts multiple items into the database.
// It does not post update any auto increment field.
// It builds an insert query of multiple rows and send it on the underlying connection.
func (c jCompositePkrelatedsToHasManyCompositePkrelatedsQuerier) InsertBulk(items ...*CompositePkrelatedsToHasManyCompositePkrelateds) error {
	panic("todo")
}

//Delete returns a delete builder
func (c jCompositePkrelatedsToHasManyCompositePkrelatedsQuerier) Delete() *jCompositePkrelatedsToHasManyCompositePkrelatedsDeleteBuilder {
	return &jCompositePkrelatedsToHasManyCompositePkrelatedsDeleteBuilder{
		&builder.DeleteBuilder{
			DeleteBuilder: c.db.DeleteFrom(JCompositePkrelatedsToHasManyCompositePkrelatedsModel.Table()),
		},
	}
}

// MustDelete requires the query to affeect rows.
func (c jCompositePkrelatedsToHasManyCompositePkrelatedsQuerier) MustDelete() *jCompositePkrelatedsToHasManyCompositePkrelatedsDeleteBuilder {
	ret := &jCompositePkrelatedsToHasManyCompositePkrelatedsDeleteBuilder{
		&builder.DeleteBuilder{
			DeleteBuilder: c.db.DeleteFrom(JCompositePkrelatedsToHasManyCompositePkrelatedsModel.Table()),
		},
	}
	ret.MustDelete()
	return ret
}

//DeleteByPk deletes one CompositePkrelatedsToHasManyCompositePkrelateds by its PKs
func (c jCompositePkrelatedsToHasManyCompositePkrelatedsQuerier) DeleteByPk(HasManyCompositePkID int64, CompositePkP string, CompositePkK string) error {
	_, err := c.Delete().Where(

		JCompositePkrelatedsToHasManyCompositePkrelatedsModel.HasManyCompositePkID.Eq(HasManyCompositePkID),

		JCompositePkrelatedsToHasManyCompositePkrelatedsModel.CompositePkP.Eq(CompositePkP),

		JCompositePkrelatedsToHasManyCompositePkrelatedsModel.CompositePkK.Eq(CompositePkK),
	).Exec()
	return err
}

// DeleteAll given CompositePkrelatedsToHasManyCompositePkrelateds
func (c jCompositePkrelatedsToHasManyCompositePkrelatedsQuerier) DeleteAll(items ...*CompositePkrelatedsToHasManyCompositePkrelateds) (sql.Result, error) {
	q := c.Delete().Where(
		JCompositePkrelatedsToHasManyCompositePkrelatedsModel.In(items...),
	)
	return q.Exec()
}

// MustDeleteAll given CompositePkrelatedsToHasManyCompositePkrelateds
func (c jCompositePkrelatedsToHasManyCompositePkrelatedsQuerier) MustDeleteAll(items ...*CompositePkrelatedsToHasManyCompositePkrelateds) (sql.Result, error) {
	var res sql.Result
	var err error
	for _, d := range items {
		res, err = c.DeleteAll(d)
		if err != nil {
			return res, err
		}
		if n, e := res.RowsAffected(); e != nil {
			return res, e
		} else if n == 0 {
			q := c.Delete().Where(
				JCompositePkrelatedsToHasManyCompositePkrelatedsModel.In(items...),
			)
			err = runtime.NewNoRowsAffected(q.String())
			return res, err
		}
	}
	return res, err
}

//Find one CompositePkrelatedsToHasManyCompositePkrelateds using its PKs
func (c jCompositePkrelatedsToHasManyCompositePkrelatedsQuerier) Find(HasManyCompositePkID int64, CompositePkP string, CompositePkK string) (*CompositePkrelatedsToHasManyCompositePkrelateds, error) {
	return c.Select().Where(

		JCompositePkrelatedsToHasManyCompositePkrelatedsModel.HasManyCompositePkID.Eq(HasManyCompositePkID),

		JCompositePkrelatedsToHasManyCompositePkrelatedsModel.CompositePkP.Eq(CompositePkP),

		JCompositePkrelatedsToHasManyCompositePkrelatedsModel.CompositePkK.Eq(CompositePkK),
	).Read()
}
