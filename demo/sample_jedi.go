// Generated with mh-cbon/jedi. Do not edit by hand.
package main

import (
	"database/sql"
	"fmt"
	"strings"

	"github.com/gocraft/dbr"
	"github.com/mh-cbon/jedi/builder"
	"github.com/mh-cbon/jedi/drivers"
	"github.com/mh-cbon/jedi/runtime"
)

var _ = fmt.Sprintf

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

//Create applies the create table command to te underlying connection.
func (c jSampleSetup) Create(db *dbr.Connection) error {
	_, err := db.Exec(c.CreateStmt)
	return err
}

//Drop applies the drop table command to te underlying connection.
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
		ors = append(ors, dbr.Or(
			dbr.And(

				dbr.Eq(`id`, t.ID),
			),
		))
	}
	return dbr.And(ors...)
}

// In provided items.
func (j jSampleModel) In(s ...*Sample) dbr.Builder {
	ors := []dbr.Builder{}
	for _, t := range s {
		ors = append(ors, dbr.Or(
			dbr.And(

				dbr.Eq(`id`, t.ID),
			),
		))
	}
	return dbr.And(ors...)
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

//Build builds the sql string using current dialect into given bufer
func (c *jSampleDeleteBuilder) Build(b dbr.Buffer) error {
	return c.DeleteBuilder.Build(runtime.GetDialect(), b)
}

//String returns the sql string for current dialect. It returns empty string if the build returns an error.
func (c *jSampleDeleteBuilder) String() string {
	b := dbr.NewBuffer()
	if err := c.Build(b); err != nil {
		return ""
	}
	return b.String()
}

//Where returns a jSampleDeleteBuilder instead of builder.DeleteBuilder.
func (c *jSampleDeleteBuilder) Where(query interface{}, value ...interface{}) *jSampleDeleteBuilder {
	c.DeleteBuilder.Where(query, value...)
	return c
}

type jSampleSelectBuilder struct {
	as string
	*builder.SelectBuilder
}

//Build builds the sql string using current dialect into given bufer
func (c *jSampleSelectBuilder) Build(b dbr.Buffer) error {
	return c.SelectBuilder.Build(runtime.GetDialect(), b)
}

//String returns the sql string for current dialect. It returns empty string if the build returns an error.
func (c *jSampleSelectBuilder) String() string {
	b := dbr.NewBuffer()
	if err := c.Build(b); err != nil {
		return ""
	}
	return b.String()
}

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
	if len(what) == 0 {
		what = m.Fields("*")
	}
	dialect := runtime.GetDialect()
	from := dialect.QuoteIdent(m.Table())
	if m.Alias() != "" && m.Alias() != m.Table() {
		from = fmt.Sprintf("%v as %v", from, dialect.QuoteIdent(m.Alias()))
	}
	return &jSampleSelectBuilder{
		as: c.as,
		SelectBuilder: &builder.SelectBuilder{
			SelectBuilder: c.db.Select(what...).From(from),
		},
	}
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
		res, err = c.db.InsertInto(JSampleModel.Table()).Columns(

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
		res, err = c.db.Update(JSampleModel.Table()).
			Set(`name`, data.Name).
			Set(`description`, data.Description).
			Set(`update_date`, data.UpdateDate).
			Set(`removal_date`, data.RemovalDate).
			Where("id = ?", data.ID).
			Exec()
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

//DeleteByPk deletes one Sample by its PKs
func (c jSampleQuerier) DeleteByPk(ID int64) error {
	_, err := c.Delete().Where(

		JSampleModel.ID.Eq(ID),
	).Exec()
	return err
}

// DeleteAll given Sample
func (c jSampleQuerier) DeleteAll(items ...*Sample) (sql.Result, error) {
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
	return q.Exec()
}

//Find one Sample using its PKs
func (c jSampleQuerier) Find(ID int64) (*Sample, error) {
	return c.Select().Where(

		JSampleModel.ID.Eq(ID),
	).Read()
}

type jSample2Setup struct {
	Name       string
	CreateStmt string
	DropStmt   string
}

//Create applies the create table command to te underlying connection.
func (c jSample2Setup) Create(db *dbr.Connection) error {
	_, err := db.Exec(c.CreateStmt)
	return err
}

//Drop applies the drop table command to te underlying connection.
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
name VARCHAR(255) NOT NULL,
description TEXT,
PRIMARY KEY (name) 

)`
	} else if driver == drivers.Pgsql {
		create = `CREATE TABLE IF NOT EXISTS second_sample (
name VARCHAR(255),
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

// jSample2Model provides helper to work with Sample2 data provider
type jSample2Model struct {
	as string

	Name builder.ValuePropertyMeta

	Description builder.ValuePropertyMeta
}

// Eq provided items.
func (j jSample2Model) Eq(s ...*Sample2) dbr.Builder {
	ors := []dbr.Builder{}
	for _, t := range s {
		ors = append(ors, dbr.Or(
			dbr.And(

				dbr.Eq(`name`, t.Name),
			),
		))
	}
	return dbr.And(ors...)
}

// In provided items.
func (j jSample2Model) In(s ...*Sample2) dbr.Builder {
	ors := []dbr.Builder{}
	for _, t := range s {
		ors = append(ors, dbr.Or(
			dbr.And(

				dbr.Eq(`name`, t.Name),
			),
		))
	}
	return dbr.And(ors...)
}

// As returns a copy with an alias.
func (j jSample2Model) As(as string) jSample2Model {
	j.as = as

	j.Name.TableAlias = as

	j.Description.TableAlias = as

	return j
}

// Table returns the sql table name
func (j jSample2Model) Table() string {
	return "second_sample"
}

// Alias returns the current alias
func (j jSample2Model) Alias() string {
	if j.as == "" {
		return j.Table()
	}
	return j.as
}

// Properties returns a map of property name => meta
func (j jSample2Model) Properties() map[string]builder.MetaProvider {
	ret := map[string]builder.MetaProvider{}

	ret["Name"] = j.Name

	ret["Description"] = j.Description

	return ret
}

// Fields returns given sql fields with appropriate aliasing.
func (j jSample2Model) Fields(ins ...string) []string {
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

// JSample2Model provides helper to work with Sample2 data provider
var JSample2Model = jSample2Model{

	Name: builder.NewValueMeta(
		`name`, `VARCHAR(255)`,
		`Name`, `string`,
		true, false,
	),

	Description: builder.NewValueMeta(
		`description`, `TEXT`,
		`Description`, `string`,
		false, false,
	),
}

type jSample2DeleteBuilder struct {
	*builder.DeleteBuilder
}

//Build builds the sql string using current dialect into given bufer
func (c *jSample2DeleteBuilder) Build(b dbr.Buffer) error {
	return c.DeleteBuilder.Build(runtime.GetDialect(), b)
}

//String returns the sql string for current dialect. It returns empty string if the build returns an error.
func (c *jSample2DeleteBuilder) String() string {
	b := dbr.NewBuffer()
	if err := c.Build(b); err != nil {
		return ""
	}
	return b.String()
}

//Where returns a jSample2DeleteBuilder instead of builder.DeleteBuilder.
func (c *jSample2DeleteBuilder) Where(query interface{}, value ...interface{}) *jSample2DeleteBuilder {
	c.DeleteBuilder.Where(query, value...)
	return c
}

type jSample2SelectBuilder struct {
	as string
	*builder.SelectBuilder
}

//Build builds the sql string using current dialect into given bufer
func (c *jSample2SelectBuilder) Build(b dbr.Buffer) error {
	return c.SelectBuilder.Build(runtime.GetDialect(), b)
}

//String returns the sql string for current dialect. It returns empty string if the build returns an error.
func (c *jSample2SelectBuilder) String() string {
	b := dbr.NewBuffer()
	if err := c.Build(b); err != nil {
		return ""
	}
	return b.String()
}

//Read evaluates current select query and load the results into a Sample2
func (c *jSample2SelectBuilder) Read() (*Sample2, error) {
	var one Sample2
	err := c.LoadStruct(&one)
	return &one, err
}

//ReadAll evaluates current select query and load the results into a slice of Sample2
func (c *jSample2SelectBuilder) ReadAll() ([]*Sample2, error) {
	var all []*Sample2
	_, err := c.LoadStructs(&all)
	return all, err
}

//Where returns a jSample2SelectBuilder instead of builder.SelectBuilder.
func (c *jSample2SelectBuilder) Where(query interface{}, value ...interface{}) *jSample2SelectBuilder {
	c.SelectBuilder.Where(query, value...)
	return c
}

//Join returns a jSample2SelectBuilder instead of builder.SelectBuilder.
func (c *jSample2SelectBuilder) Join(table, on interface{}) *jSample2SelectBuilder {
	c.SelectBuilder.Join(table, on)
	return c
}

//LeftJoin returns a jSample2SelectBuilder instead of builder.SelectBuilder.
func (c *jSample2SelectBuilder) LeftJoin(table, on interface{}) *jSample2SelectBuilder {
	c.SelectBuilder.LeftJoin(table, on)
	return c
}

//RightJoin returns a jSample2SelectBuilder instead of builder.SelectBuilder.
func (c *jSample2SelectBuilder) RightJoin(table, on interface{}) *jSample2SelectBuilder {
	c.SelectBuilder.RightJoin(table, on)
	return c
}

//FullJoin returns a jSample2SelectBuilder instead of builder.SelectBuilder.
func (c *jSample2SelectBuilder) FullJoin(table, on interface{}) *jSample2SelectBuilder {
	c.SelectBuilder.FullJoin(table, on)
	return c
}

//Distinct returns a jSample2SelectBuilder instead of builder.SelectBuilder.
func (c *jSample2SelectBuilder) Distinct() *jSample2SelectBuilder {
	c.SelectBuilder.Distinct()
	return c
}

// JSample2 provides a basic querier
func JSample2(db dbr.SessionRunner) jSample2Querier {
	return jSample2Querier{
		db: db,
	}
}

type jSample2Querier struct {
	db dbr.SessionRunner
	as string
}

//As set alias prior building.
func (c jSample2Querier) As(as string) jSample2Querier {
	c.as = as
	return c
}

//Model returns a model with appropriate aliasing.
func (c jSample2Querier) Model() jSample2Model {
	return JSample2Model.As(c.as)
}

//Select returns a Sample2 Select Builder.
func (c jSample2Querier) Select(what ...string) *jSample2SelectBuilder {
	m := c.Model()
	if len(what) == 0 {
		what = m.Fields("*")
	}
	dialect := runtime.GetDialect()
	from := dialect.QuoteIdent(m.Table())
	if m.Alias() != "" && m.Alias() != m.Table() {
		from = fmt.Sprintf("%v as %v", from, dialect.QuoteIdent(m.Alias()))
	}
	return &jSample2SelectBuilder{
		as: c.as,
		SelectBuilder: &builder.SelectBuilder{
			SelectBuilder: c.db.Select(what...).From(from),
		},
	}
}

//Count returns a Sample2 Select Builder to count given expressions.
func (c jSample2Querier) Count(what ...string) *jSample2SelectBuilder {
	if len(what) == 0 {
		what = append(what, "*")
	}
	return c.Select("COUNT(" + strings.Join(what, ",") + ")")
}

// Insert a new Sample2, if it has autoincrement primary key, the value will be set.
// It stops on first error.
func (c jSample2Querier) Insert(items ...*Sample2) (sql.Result, error) {
	var res sql.Result
	var err error
	for _, data := range items {
		res, err = c.db.InsertInto(JSample2Model.Table()).Columns(

			`name`,

			`description`,
		).Record(data).Exec()

		if err != nil {
			return res, err
		}
	}
	return res, err
}

// InsertBulk inserts multiple items into the database.
// It does not post update any auto increment field.
// It builds an insert query of multiple rows and send it on the underlying connection.
func (c jSample2Querier) InsertBulk(items ...*Sample2) error {
	panic("todo")
}

// Update a Sample2. It stops on first error.
func (c jSample2Querier) Update(items ...*Sample2) (sql.Result, error) {
	var res sql.Result
	var err error
	for _, data := range items {
		res, err = c.db.Update(JSample2Model.Table()).
			Set(`description`, data.Description).
			Where("name = ?", data.Name).
			Exec()
		if err != nil {
			return res, err
		}
	}
	return res, err
}

// UpdateBulk updates multiple items into the database.
// It builds an update query of multiple rows and send it on the underlying connection.
func (c jSample2Querier) UpdateBulk(items ...*Sample2) error {
	panic("todo")
}

//Delete returns a delete builder
func (c jSample2Querier) Delete() *jSample2DeleteBuilder {
	return &jSample2DeleteBuilder{
		&builder.DeleteBuilder{
			DeleteBuilder: c.db.DeleteFrom(JSample2Model.Table()),
		},
	}
}

//DeleteByPk deletes one Sample2 by its PKs
func (c jSample2Querier) DeleteByPk(Name string) error {
	_, err := c.Delete().Where(

		JSample2Model.Name.Eq(Name),
	).Exec()
	return err
}

// DeleteAll given Sample2
func (c jSample2Querier) DeleteAll(items ...*Sample2) (sql.Result, error) {
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
	return q.Exec()
}

//Find one Sample2 using its PKs
func (c jSample2Querier) Find(Name string) (*Sample2, error) {
	return c.Select().Where(

		JSample2Model.Name.Eq(Name),
	).Read()
}

type jSampleViewSetup struct {
	Name       string
	CreateStmt string
	DropStmt   string
}

//Create applies the create table command to te underlying connection.
func (c jSampleViewSetup) Create(db *dbr.Connection) error {
	_, err := db.Exec(c.CreateStmt)
	return err
}

//Drop applies the drop table command to te underlying connection.
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
		ors = append(ors, dbr.Or(
			dbr.And(

				dbr.Eq(`id`, t.ID),
			),
		))
	}
	return dbr.And(ors...)
}

// In provided items.
func (j jSampleViewModel) In(s ...*SampleView) dbr.Builder {
	ors := []dbr.Builder{}
	for _, t := range s {
		ors = append(ors, dbr.Or(
			dbr.And(

				dbr.Eq(`id`, t.ID),
			),
		))
	}
	return dbr.And(ors...)
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

//Build builds the sql string using current dialect into given bufer
func (c *jSampleViewDeleteBuilder) Build(b dbr.Buffer) error {
	return c.DeleteBuilder.Build(runtime.GetDialect(), b)
}

//String returns the sql string for current dialect. It returns empty string if the build returns an error.
func (c *jSampleViewDeleteBuilder) String() string {
	b := dbr.NewBuffer()
	if err := c.Build(b); err != nil {
		return ""
	}
	return b.String()
}

//Where returns a jSampleViewDeleteBuilder instead of builder.DeleteBuilder.
func (c *jSampleViewDeleteBuilder) Where(query interface{}, value ...interface{}) *jSampleViewDeleteBuilder {
	c.DeleteBuilder.Where(query, value...)
	return c
}

type jSampleViewSelectBuilder struct {
	as string
	*builder.SelectBuilder
}

//Build builds the sql string using current dialect into given bufer
func (c *jSampleViewSelectBuilder) Build(b dbr.Buffer) error {
	return c.SelectBuilder.Build(runtime.GetDialect(), b)
}

//String returns the sql string for current dialect. It returns empty string if the build returns an error.
func (c *jSampleViewSelectBuilder) String() string {
	b := dbr.NewBuffer()
	if err := c.Build(b); err != nil {
		return ""
	}
	return b.String()
}

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
	if len(what) == 0 {
		what = m.Fields("*")
	}
	dialect := runtime.GetDialect()
	from := dialect.QuoteIdent(m.Table())
	if m.Alias() != "" && m.Alias() != m.Table() {
		from = fmt.Sprintf("%v as %v", from, dialect.QuoteIdent(m.Alias()))
	}
	return &jSampleViewSelectBuilder{
		as: c.as,
		SelectBuilder: &builder.SelectBuilder{
			SelectBuilder: c.db.Select(what...).From(from),
		},
	}
}

//Count returns a SampleView Select Builder to count given expressions.
func (c jSampleViewQuerier) Count(what ...string) *jSampleViewSelectBuilder {
	if len(what) == 0 {
		what = append(what, "*")
	}
	return c.Select("COUNT(" + strings.Join(what, ",") + ")")
}
