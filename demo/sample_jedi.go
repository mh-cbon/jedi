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

		JBasicTypesSetup,

		JTextPkSetup,

		JCompositePkSetup,

		JDateTypeSetup,

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
	return runtime.NewSQLError(err, c.CreateStmt)
}

//Drop applies the drop table command to te underlying connection.
func (c jSampleSetup) Drop(db *dbr.Connection) error {
	_, err := db.Exec(c.DropStmt)
	return runtime.NewSQLError(err, c.DropStmt)
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
update_date timestamp,
removal_date timestamp

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

		{
			x := data.RemovalDate.UTC()
			data.RemovalDate = &x
		}

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

		data.UpdateDate = data.UpdateDate.UTC()

		{
			x := data.RemovalDate.UTC()
			data.RemovalDate = &x
		}

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
	q := c.Delete().Where(
		JSampleModel.In(items...),
	)
	return q.Exec()
}

//Find one Sample using its PKs
func (c jSampleQuerier) Find(ID int64) (*Sample, error) {
	return c.Select().Where(

		JSampleModel.ID.Eq(ID),
	).Read()
}

type jBasicTypesSetup struct {
	Name       string
	CreateStmt string
	DropStmt   string
}

//Create applies the create table command to te underlying connection.
func (c jBasicTypesSetup) Create(db *dbr.Connection) error {
	_, err := db.Exec(c.CreateStmt)
	return runtime.NewSQLError(err, c.CreateStmt)
}

//Drop applies the drop table command to te underlying connection.
func (c jBasicTypesSetup) Drop(db *dbr.Connection) error {
	_, err := db.Exec(c.DropStmt)
	return runtime.NewSQLError(err, c.DropStmt)
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
string_p TEXT,
intfield INTEGER,
int_p INTEGER,
int32 INTEGER,
int32_p INTEGER,
int64 INTEGER,
int64_p INTEGER,
u_int INTEGER,
u_int_p INTEGER,
u_int32 INTEGER,
u_int32_p INTEGER,
u_int64 INTEGER,
u_int64_p INTEGER,
bool INTEGER,
bool_p INTEGER,
float32 FLOAT,
float32_p FLOAT,
float64 FLOAT,
float64_p FLOAT

)`
	} else if driver == drivers.Mysql {
		create = `CREATE TABLE IF NOT EXISTS basic_types (
id INTEGER NOT NULL AUTO_INCREMENT,
string TEXT,
string_p TEXT,
intfield INTEGER,
int_p INTEGER,
int32 INTEGER,
int32_p INTEGER,
int64 INTEGER,
int64_p INTEGER,
u_int INTEGER,
u_int_p INTEGER,
u_int32 INTEGER,
u_int32_p INTEGER,
u_int64 INTEGER,
u_int64_p INTEGER,
bool INTEGER,
bool_p INTEGER,
float32 FLOAT,
float32_p FLOAT,
float64 FLOAT,
float64_p FLOAT,
PRIMARY KEY (id) 

)`
	} else if driver == drivers.Pgsql {
		create = `CREATE TABLE IF NOT EXISTS basic_types (
id INTEGER,
string TEXT,
string_p TEXT,
intfield INTEGER,
int_p INTEGER,
int32 INTEGER,
int32_p INTEGER,
int64 INTEGER,
int64_p INTEGER,
u_int INTEGER,
u_int_p INTEGER,
u_int32 INTEGER,
u_int32_p INTEGER,
u_int64 INTEGER,
u_int64_p INTEGER,
bool INTEGER,
bool_p INTEGER,
float32 FLOAT,
float32_p FLOAT,
float64 FLOAT,
float64_p FLOAT

)`
	}

	if driver == drivers.Sqlite {
		drop = `DROP TABLE IF EXISTS basic_types`
	} else if driver == drivers.Mysql {
		drop = `DROP TABLE IF EXISTS basic_types`
	} else if driver == drivers.Pgsql {
		drop = `DROP TABLE IF EXISTS basic_types`
	}

	return jBasicTypesSetup{
		Name:       `basic_types`,
		CreateStmt: create,
		DropStmt:   drop,
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

//Build builds the sql string using current dialect into given bufer
func (c *jBasicTypesDeleteBuilder) Build(b dbr.Buffer) error {
	return c.DeleteBuilder.Build(runtime.GetDialect(), b)
}

//String returns the sql string for current dialect. It returns empty string if the build returns an error.
func (c *jBasicTypesDeleteBuilder) String() string {
	b := dbr.NewBuffer()
	if err := c.Build(b); err != nil {
		return ""
	}
	return b.String()
}

//Where returns a jBasicTypesDeleteBuilder instead of builder.DeleteBuilder.
func (c *jBasicTypesDeleteBuilder) Where(query interface{}, value ...interface{}) *jBasicTypesDeleteBuilder {
	c.DeleteBuilder.Where(query, value...)
	return c
}

type jBasicTypesSelectBuilder struct {
	as string
	*builder.SelectBuilder
}

//Build builds the sql string using current dialect into given bufer
func (c *jBasicTypesSelectBuilder) Build(b dbr.Buffer) error {
	return c.SelectBuilder.Build(runtime.GetDialect(), b)
}

//String returns the sql string for current dialect. It returns empty string if the build returns an error.
func (c *jBasicTypesSelectBuilder) String() string {
	b := dbr.NewBuffer()
	if err := c.Build(b); err != nil {
		return ""
	}
	return b.String()
}

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

		res, err = c.db.InsertInto(JBasicTypesModel.Table()).Columns(

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
func (c jBasicTypesQuerier) InsertBulk(items ...*BasicTypes) error {
	panic("todo")
}

// Update a BasicTypes. It stops on first error.
func (c jBasicTypesQuerier) Update(items ...*BasicTypes) (sql.Result, error) {
	var res sql.Result
	var err error
	for _, data := range items {

		res, err = c.db.Update(JBasicTypesModel.Table()).
			Set(`string`, data.String).
			Set(`string_p`, data.StringP).
			Set(`intfield`, data.Int).
			Set(`int_p`, data.IntP).
			Set(`int32`, data.Int32).
			Set(`int32_p`, data.Int32P).
			Set(`int64`, data.Int64).
			Set(`int64_p`, data.Int64P).
			Set(`u_int`, data.UInt).
			Set(`u_int_p`, data.UIntP).
			Set(`u_int32`, data.UInt32).
			Set(`u_int32_p`, data.UInt32P).
			Set(`u_int64`, data.UInt64).
			Set(`u_int64_p`, data.UInt64P).
			Set(`bool`, data.Bool).
			Set(`bool_p`, data.BoolP).
			Set(`float32`, data.Float32).
			Set(`float32_p`, data.Float32P).
			Set(`float64`, data.Float64).
			Set(`float64_p`, data.Float64P).
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

//Find one BasicTypes using its PKs
func (c jBasicTypesQuerier) Find(ID int64) (*BasicTypes, error) {
	return c.Select().Where(

		JBasicTypesModel.ID.Eq(ID),
	).Read()
}

type jTextPkSetup struct {
	Name       string
	CreateStmt string
	DropStmt   string
}

//Create applies the create table command to te underlying connection.
func (c jTextPkSetup) Create(db *dbr.Connection) error {
	_, err := db.Exec(c.CreateStmt)
	return runtime.NewSQLError(err, c.CreateStmt)
}

//Drop applies the drop table command to te underlying connection.
func (c jTextPkSetup) Drop(db *dbr.Connection) error {
	_, err := db.Exec(c.DropStmt)
	return runtime.NewSQLError(err, c.DropStmt)
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

	return jTextPkSetup{
		Name:       `second_sample`,
		CreateStmt: create,
		DropStmt:   drop,
	}
}

// jTextPkModel provides helper to work with TextPk data provider
type jTextPkModel struct {
	as string

	Name builder.ValuePropertyMeta

	Description builder.ValuePropertyMeta
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
}

type jTextPkDeleteBuilder struct {
	*builder.DeleteBuilder
}

//Build builds the sql string using current dialect into given bufer
func (c *jTextPkDeleteBuilder) Build(b dbr.Buffer) error {
	return c.DeleteBuilder.Build(runtime.GetDialect(), b)
}

//String returns the sql string for current dialect. It returns empty string if the build returns an error.
func (c *jTextPkDeleteBuilder) String() string {
	b := dbr.NewBuffer()
	if err := c.Build(b); err != nil {
		return ""
	}
	return b.String()
}

//Where returns a jTextPkDeleteBuilder instead of builder.DeleteBuilder.
func (c *jTextPkDeleteBuilder) Where(query interface{}, value ...interface{}) *jTextPkDeleteBuilder {
	c.DeleteBuilder.Where(query, value...)
	return c
}

type jTextPkSelectBuilder struct {
	as string
	*builder.SelectBuilder
}

//Build builds the sql string using current dialect into given bufer
func (c *jTextPkSelectBuilder) Build(b dbr.Buffer) error {
	return c.SelectBuilder.Build(runtime.GetDialect(), b)
}

//String returns the sql string for current dialect. It returns empty string if the build returns an error.
func (c *jTextPkSelectBuilder) String() string {
	b := dbr.NewBuffer()
	if err := c.Build(b); err != nil {
		return ""
	}
	return b.String()
}

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

		res, err = c.db.InsertInto(JTextPkModel.Table()).Columns(

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
func (c jTextPkQuerier) InsertBulk(items ...*TextPk) error {
	panic("todo")
}

// Update a TextPk. It stops on first error.
func (c jTextPkQuerier) Update(items ...*TextPk) (sql.Result, error) {
	var res sql.Result
	var err error
	for _, data := range items {

		res, err = c.db.Update(JTextPkModel.Table()).
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

//Find one TextPk using its PKs
func (c jTextPkQuerier) Find(Name string) (*TextPk, error) {
	return c.Select().Where(

		JTextPkModel.Name.Eq(Name),
	).Read()
}

type jCompositePkSetup struct {
	Name       string
	CreateStmt string
	DropStmt   string
}

//Create applies the create table command to te underlying connection.
func (c jCompositePkSetup) Create(db *dbr.Connection) error {
	_, err := db.Exec(c.CreateStmt)
	return runtime.NewSQLError(err, c.CreateStmt)
}

//Drop applies the drop table command to te underlying connection.
func (c jCompositePkSetup) Drop(db *dbr.Connection) error {
	_, err := db.Exec(c.DropStmt)
	return runtime.NewSQLError(err, c.DropStmt)
}

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
description TEXT

)`
	}

	if driver == drivers.Sqlite {
		drop = `DROP TABLE IF EXISTS composite_pk`
	} else if driver == drivers.Mysql {
		drop = `DROP TABLE IF EXISTS composite_pk`
	} else if driver == drivers.Pgsql {
		drop = `DROP TABLE IF EXISTS composite_pk`
	}

	return jCompositePkSetup{
		Name:       `composite_pk`,
		CreateStmt: create,
		DropStmt:   drop,
	}
}

// jCompositePkModel provides helper to work with CompositePk data provider
type jCompositePkModel struct {
	as string

	P builder.ValuePropertyMeta

	K builder.ValuePropertyMeta

	Description builder.ValuePropertyMeta
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
}

type jCompositePkDeleteBuilder struct {
	*builder.DeleteBuilder
}

//Build builds the sql string using current dialect into given bufer
func (c *jCompositePkDeleteBuilder) Build(b dbr.Buffer) error {
	return c.DeleteBuilder.Build(runtime.GetDialect(), b)
}

//String returns the sql string for current dialect. It returns empty string if the build returns an error.
func (c *jCompositePkDeleteBuilder) String() string {
	b := dbr.NewBuffer()
	if err := c.Build(b); err != nil {
		return ""
	}
	return b.String()
}

//Where returns a jCompositePkDeleteBuilder instead of builder.DeleteBuilder.
func (c *jCompositePkDeleteBuilder) Where(query interface{}, value ...interface{}) *jCompositePkDeleteBuilder {
	c.DeleteBuilder.Where(query, value...)
	return c
}

type jCompositePkSelectBuilder struct {
	as string
	*builder.SelectBuilder
}

//Build builds the sql string using current dialect into given bufer
func (c *jCompositePkSelectBuilder) Build(b dbr.Buffer) error {
	return c.SelectBuilder.Build(runtime.GetDialect(), b)
}

//String returns the sql string for current dialect. It returns empty string if the build returns an error.
func (c *jCompositePkSelectBuilder) String() string {
	b := dbr.NewBuffer()
	if err := c.Build(b); err != nil {
		return ""
	}
	return b.String()
}

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

		res, err = c.db.InsertInto(JCompositePkModel.Table()).Columns(

			`p`,

			`k`,

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
func (c jCompositePkQuerier) InsertBulk(items ...*CompositePk) error {
	panic("todo")
}

// Update a CompositePk. It stops on first error.
func (c jCompositePkQuerier) Update(items ...*CompositePk) (sql.Result, error) {
	var res sql.Result
	var err error
	for _, data := range items {

		res, err = c.db.Update(JCompositePkModel.Table()).
			Set(`description`, data.Description).
			Where("p = ?", data.P).
			Where("k = ?", data.K).
			Exec()
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

//Find one CompositePk using its PKs
func (c jCompositePkQuerier) Find(P string, K string) (*CompositePk, error) {
	return c.Select().Where(

		JCompositePkModel.P.Eq(P),

		JCompositePkModel.K.Eq(K),
	).Read()
}

type jDateTypeSetup struct {
	Name       string
	CreateStmt string
	DropStmt   string
}

//Create applies the create table command to te underlying connection.
func (c jDateTypeSetup) Create(db *dbr.Connection) error {
	_, err := db.Exec(c.CreateStmt)
	return runtime.NewSQLError(err, c.CreateStmt)
}

//Drop applies the drop table command to te underlying connection.
func (c jDateTypeSetup) Drop(db *dbr.Connection) error {
	_, err := db.Exec(c.DropStmt)
	return runtime.NewSQLError(err, c.DropStmt)
}

// JDateTypeSetup helps to create/drop the schema
func JDateTypeSetup() runtime.Setuper {
	driver := runtime.GetCurrentDriver()

	var create string
	var drop string

	if driver == drivers.Sqlite {
		create = `CREATE TABLE IF NOT EXISTS date_type (
id INTEGER PRIMARY KEY AUTOINCREMENT,
t datetime,
tp datetime

)`
	} else if driver == drivers.Mysql {
		create = `CREATE TABLE IF NOT EXISTS date_type (
id INTEGER NOT NULL AUTO_INCREMENT,
t datetime,
tp datetime,
PRIMARY KEY (id) 

)`
	} else if driver == drivers.Pgsql {
		create = `CREATE TABLE IF NOT EXISTS date_type (
id INTEGER,
t timestamp,
tp timestamp

)`
	}

	if driver == drivers.Sqlite {
		drop = `DROP TABLE IF EXISTS date_type`
	} else if driver == drivers.Mysql {
		drop = `DROP TABLE IF EXISTS date_type`
	} else if driver == drivers.Pgsql {
		drop = `DROP TABLE IF EXISTS date_type`
	}

	return jDateTypeSetup{
		Name:       `date_type`,
		CreateStmt: create,
		DropStmt:   drop,
	}
}

// jDateTypeModel provides helper to work with DateType data provider
type jDateTypeModel struct {
	as string

	ID builder.ValuePropertyMeta

	T builder.ValuePropertyMeta

	TP builder.ValuePropertyMeta
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
}

type jDateTypeDeleteBuilder struct {
	*builder.DeleteBuilder
}

//Build builds the sql string using current dialect into given bufer
func (c *jDateTypeDeleteBuilder) Build(b dbr.Buffer) error {
	return c.DeleteBuilder.Build(runtime.GetDialect(), b)
}

//String returns the sql string for current dialect. It returns empty string if the build returns an error.
func (c *jDateTypeDeleteBuilder) String() string {
	b := dbr.NewBuffer()
	if err := c.Build(b); err != nil {
		return ""
	}
	return b.String()
}

//Where returns a jDateTypeDeleteBuilder instead of builder.DeleteBuilder.
func (c *jDateTypeDeleteBuilder) Where(query interface{}, value ...interface{}) *jDateTypeDeleteBuilder {
	c.DeleteBuilder.Where(query, value...)
	return c
}

type jDateTypeSelectBuilder struct {
	as string
	*builder.SelectBuilder
}

//Build builds the sql string using current dialect into given bufer
func (c *jDateTypeSelectBuilder) Build(b dbr.Buffer) error {
	return c.SelectBuilder.Build(runtime.GetDialect(), b)
}

//String returns the sql string for current dialect. It returns empty string if the build returns an error.
func (c *jDateTypeSelectBuilder) String() string {
	b := dbr.NewBuffer()
	if err := c.Build(b); err != nil {
		return ""
	}
	return b.String()
}

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

		data.T = data.T.UTC()

		{
			x := data.TP.UTC()
			data.TP = &x
		}

		res, err = c.db.InsertInto(JDateTypeModel.Table()).Columns(

			`t`,

			`tp`,
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
func (c jDateTypeQuerier) InsertBulk(items ...*DateType) error {
	panic("todo")
}

// Update a DateType. It stops on first error.
func (c jDateTypeQuerier) Update(items ...*DateType) (sql.Result, error) {
	var res sql.Result
	var err error
	for _, data := range items {

		data.T = data.T.UTC()

		{
			x := data.TP.UTC()
			data.TP = &x
		}

		res, err = c.db.Update(JDateTypeModel.Table()).
			Set(`t`, data.T).
			Set(`tp`, data.TP).
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

//Find one DateType using its PKs
func (c jDateTypeQuerier) Find(ID int64) (*DateType, error) {
	return c.Select().Where(

		JDateTypeModel.ID.Eq(ID),
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
	return runtime.NewSQLError(err, c.CreateStmt)
}

//Drop applies the drop table command to te underlying connection.
func (c jSampleViewSetup) Drop(db *dbr.Connection) error {
	_, err := db.Exec(c.DropStmt)
	return runtime.NewSQLError(err, c.DropStmt)
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
		create = `CREATE VIEW IF NOT EXISTS sample_view AS 
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
