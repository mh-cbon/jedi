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

		JProductSetup,

		JCategorySetup,

		JBrandSetup,

		JCategoryproductsToProductcategoriesSetup,
	)
}

type jProductSetup struct {
	Name       string
	CreateStmt string
	DropStmt   string
}

//Create applies the create table command to te underlying connection.
func (c jProductSetup) Create(db *dbr.Connection) error {
	_, err := db.Exec(c.CreateStmt)
	return err
}

//Drop applies the drop table command to te underlying connection.
func (c jProductSetup) Drop(db *dbr.Connection) error {
	_, err := db.Exec(c.DropStmt)
	return err
}

// JProductSetup helps to create/drop the schema
func JProductSetup() runtime.Setuper {
	driver := runtime.GetCurrentDriver()

	var create string
	var drop string

	if driver == drivers.Sqlite {
		create = `CREATE TABLE IF NOT EXISTS product (
id INTEGER PRIMARY KEY AUTOINCREMENT,
sku TEXT,
brand_id INTEGER,
brand2_id INTEGER,
master_id INTEGER

)`
	} else if driver == drivers.Mysql {
		create = `CREATE TABLE IF NOT EXISTS product (
id INTEGER NOT NULL AUTO_INCREMENT,
sku TEXT,
brand_id INTEGER,
brand2_id INTEGER,
master_id INTEGER,
PRIMARY KEY (id) 

)`
	} else if driver == drivers.Pgsql {
		create = `CREATE TABLE IF NOT EXISTS product (
id INTEGER,
sku TEXT,
brand_id INTEGER,
brand2_id INTEGER,
master_id INTEGER

)`
	}

	if driver == drivers.Sqlite {
		drop = `DROP TABLE IF EXISTS product`
	} else if driver == drivers.Mysql {
		drop = `DROP TABLE IF EXISTS product`
	} else if driver == drivers.Pgsql {
		drop = `DROP TABLE IF EXISTS product`
	}

	return jProductSetup{
		Name:       `product`,
		CreateStmt: create,
		DropStmt:   drop,
	}
}

// jProductModel provides helper to work with Product data provider
type jProductModel struct {
	as string

	ID builder.ValuePropertyMeta

	SKU builder.ValuePropertyMeta

	BrandID builder.ValuePropertyMeta

	Brand2ID builder.ValuePropertyMeta

	MasterID builder.ValuePropertyMeta

	Categories builder.RelPropertyMeta

	Brand builder.RelPropertyMeta

	Brand2 builder.RelPropertyMeta

	Master builder.RelPropertyMeta
}

// Eq provided items.
func (j jProductModel) Eq(s ...*Product) dbr.Builder {
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
func (j jProductModel) In(s ...*Product) dbr.Builder {
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
func (j jProductModel) As(as string) jProductModel {
	j.as = as

	j.ID.TableAlias = as

	j.SKU.TableAlias = as

	j.BrandID.TableAlias = as

	j.Brand2ID.TableAlias = as

	j.MasterID.TableAlias = as

	// j.Categories.TableAlias = as

	// j.Brand.TableAlias = as

	// j.Brand2.TableAlias = as

	// j.Master.TableAlias = as

	return j
}

// Table returns the sql table name
func (j jProductModel) Table() string {
	return "product"
}

// Alias returns the current alias
func (j jProductModel) Alias() string {
	if j.as == "" {
		return j.Table()
	}
	return j.as
}

// Properties returns a map of property name => meta
func (j jProductModel) Properties() map[string]builder.MetaProvider {
	ret := map[string]builder.MetaProvider{}

	ret["ID"] = j.ID

	ret["SKU"] = j.SKU

	ret["BrandID"] = j.BrandID

	ret["Brand2ID"] = j.Brand2ID

	ret["MasterID"] = j.MasterID

	ret["Categories"] = j.Categories

	ret["Brand"] = j.Brand

	ret["Brand2"] = j.Brand2

	ret["Master"] = j.Master

	return ret
}

// Fields returns given sql fields with appropriate aliasing.
func (j jProductModel) Fields(ins ...string) []string {
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

// JProductModel provides helper to work with Product data provider
var JProductModel = jProductModel{

	ID: builder.NewValueMeta(
		`id`, `INTEGER`,
		`ID`, `int64`,
		true, true,
	),

	SKU: builder.NewValueMeta(
		`sku`, `TEXT`,
		`SKU`, `string`,
		false, false,
	),

	BrandID: builder.NewValueMeta(
		`brand_id`, `INTEGER`,
		`BrandID`, `*int64`,
		false, false,
	),

	Brand2ID: builder.NewValueMeta(
		`brand2_id`, `INTEGER`,
		`Brand2ID`, `*int64`,
		false, false,
	),

	MasterID: builder.NewValueMeta(
		`master_id`, `INTEGER`,
		`MasterID`, `*int64`,
		false, false,
	),

	Categories: builder.NewRelMeta(
		`categories`, `[]*Category`,
		``, `Category.products`, ``,
	),

	Brand: builder.NewRelMeta(
		`brand`, `*Brand`,
		`Brand.products`, ``, ``,
	),

	Brand2: builder.NewRelMeta(
		`brand2`, `*Brand`,
		`Brand.products2`, ``, ``,
	),

	Master: builder.NewRelMeta(
		`master`, `*Product`,
		`Product.master`, ``, ``,
	),
}

type jProductDeleteBuilder struct {
	*builder.DeleteBuilder
}

//Build builds the sql string using current dialect into given bufer
func (c *jProductDeleteBuilder) Build(b dbr.Buffer) error {
	return c.DeleteBuilder.Build(runtime.GetDialect(), b)
}

//String returns the sql string for current dialect. It returns empty string if the build returns an error.
func (c *jProductDeleteBuilder) String() string {
	b := dbr.NewBuffer()
	if err := c.Build(b); err != nil {
		return ""
	}
	return b.String()
}

//Where returns a jProductDeleteBuilder instead of builder.DeleteBuilder.
func (c *jProductDeleteBuilder) Where(query interface{}, value ...interface{}) *jProductDeleteBuilder {
	c.DeleteBuilder.Where(query, value...)
	return c
}

type jProductSelectBuilder struct {
	as string
	*builder.SelectBuilder
}

//Build builds the sql string using current dialect into given bufer
func (c *jProductSelectBuilder) Build(b dbr.Buffer) error {
	return c.SelectBuilder.Build(runtime.GetDialect(), b)
}

//String returns the sql string for current dialect. It returns empty string if the build returns an error.
func (c *jProductSelectBuilder) String() string {
	b := dbr.NewBuffer()
	if err := c.Build(b); err != nil {
		return ""
	}
	return b.String()
}

//Read evaluates current select query and load the results into a Product
func (c *jProductSelectBuilder) Read() (*Product, error) {
	var one Product
	err := c.LoadStruct(&one)
	return &one, err
}

//ReadAll evaluates current select query and load the results into a slice of Product
func (c *jProductSelectBuilder) ReadAll() ([]*Product, error) {
	var all []*Product
	_, err := c.LoadStructs(&all)
	return all, err
}

//Where returns a jProductSelectBuilder instead of builder.SelectBuilder.
func (c *jProductSelectBuilder) Where(query interface{}, value ...interface{}) *jProductSelectBuilder {
	c.SelectBuilder.Where(query, value...)
	return c
}

//Join returns a jProductSelectBuilder instead of builder.SelectBuilder.
func (c *jProductSelectBuilder) Join(table, on interface{}) *jProductSelectBuilder {
	c.SelectBuilder.Join(table, on)
	return c
}

//LeftJoin returns a jProductSelectBuilder instead of builder.SelectBuilder.
func (c *jProductSelectBuilder) LeftJoin(table, on interface{}) *jProductSelectBuilder {
	c.SelectBuilder.LeftJoin(table, on)
	return c
}

//RightJoin returns a jProductSelectBuilder instead of builder.SelectBuilder.
func (c *jProductSelectBuilder) RightJoin(table, on interface{}) *jProductSelectBuilder {
	c.SelectBuilder.RightJoin(table, on)
	return c
}

//FullJoin returns a jProductSelectBuilder instead of builder.SelectBuilder.
func (c *jProductSelectBuilder) FullJoin(table, on interface{}) *jProductSelectBuilder {
	c.SelectBuilder.FullJoin(table, on)
	return c
}

//Distinct returns a jProductSelectBuilder instead of builder.SelectBuilder.
func (c *jProductSelectBuilder) Distinct() *jProductSelectBuilder {
	c.SelectBuilder.Distinct()
	return c
}

// JProduct provides a basic querier
func JProduct(db dbr.SessionRunner) jProductQuerier {
	return jProductQuerier{
		db: db,
	}
}

type jProductQuerier struct {
	db dbr.SessionRunner
	as string
}

//As set alias prior building.
func (c jProductQuerier) As(as string) jProductQuerier {
	c.as = as
	return c
}

//Model returns a model with appropriate aliasing.
func (c jProductQuerier) Model() jProductModel {
	return JProductModel.As(c.as)
}

//Select returns a Product Select Builder.
func (c jProductQuerier) Select(what ...string) *jProductSelectBuilder {
	m := c.Model()
	if len(what) == 0 {
		what = m.Fields("*")
	}
	dialect := runtime.GetDialect()
	from := dialect.QuoteIdent(m.Table())
	if m.Alias() != "" && m.Alias() != m.Table() {
		from = fmt.Sprintf("%v as %v", from, dialect.QuoteIdent(m.Alias()))
	}
	return &jProductSelectBuilder{
		as: c.as,
		SelectBuilder: &builder.SelectBuilder{
			SelectBuilder: c.db.Select(what...).From(from),
		},
	}
}

//Count returns a Product Select Builder to count given expressions.
func (c jProductQuerier) Count(what ...string) *jProductSelectBuilder {
	if len(what) == 0 {
		what = append(what, "*")
	}
	return c.Select("COUNT(" + strings.Join(what, ",") + ")")
}

// Insert a new Product, if it has autoincrement primary key, the value will be set.
// It stops on first error.
func (c jProductQuerier) Insert(items ...*Product) (sql.Result, error) {
	var res sql.Result
	var err error
	for _, data := range items {
		res, err = c.db.InsertInto(JProductModel.Table()).Columns(

			`sku`,

			`brand_id`,

			`brand2_id`,

			`master_id`,
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
func (c jProductQuerier) InsertBulk(items ...*Product) error {
	panic("todo")
}

// Update a Product. It stops on first error.
func (c jProductQuerier) Update(items ...*Product) (sql.Result, error) {
	var res sql.Result
	var err error
	for _, data := range items {
		res, err = c.db.Update(JProductModel.Table()).
			Set(`sku`, data.SKU).
			Set(`brand_id`, data.BrandID).
			Set(`brand2_id`, data.Brand2ID).
			Set(`master_id`, data.MasterID).
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
func (c jProductQuerier) UpdateBulk(items ...*Product) error {
	panic("todo")
}

//Delete returns a delete builder
func (c jProductQuerier) Delete() *jProductDeleteBuilder {
	return &jProductDeleteBuilder{
		&builder.DeleteBuilder{
			DeleteBuilder: c.db.DeleteFrom(JProductModel.Table()),
		},
	}
}

//DeleteByPk deletes one Product by its PKs
func (c jProductQuerier) DeleteByPk(ID int64) error {
	_, err := c.Delete().Where(

		JProductModel.ID.Eq(ID),
	).Exec()
	return err
}

// DeleteAll given Product
func (c jProductQuerier) DeleteAll(items ...*Product) (sql.Result, error) {
	q := c.Delete()
	for _, item := range items {
		q = q.Where(
			dbr.Or(
				dbr.And(

					JProductModel.ID.Eq(item.ID),
				),
			),
		)
	}
	return q.Exec()
}

//Find one Product using its PKs
func (c jProductQuerier) Find(ID int64) (*Product, error) {
	return c.Select().Where(

		JProductModel.ID.Eq(ID),
	).Read()
}

// JoinBrand adds a JOIN to Product.Brand
func (c *jProductSelectBuilder) JoinBrand(
	AsBrand string,
) *jProductSelectBuilder {
	dialect := runtime.GetDialect()
	on := ""
	localTable := dialect.QuoteIdent(JProductModel.Table())
	if c.as != "" {
		localTable = dialect.QuoteIdent(c.as)
	}
	foreiTable := dialect.QuoteIdent(JBrandModel.Table())
	if AsBrand != "" {
		foreiTable = dialect.QuoteIdent(AsBrand)
	}

	on += fmt.Sprintf("%v.%v = %v.%v",
		localTable, dialect.QuoteIdent("brand_id"),
		foreiTable, dialect.QuoteIdent("id"),
	)

	if AsBrand == "" {
		return c.Join(dbr.I(JBrandModel.Table()), on)
	}
	return c.Join(dbr.I(JBrandModel.Table()).As(AsBrand), on)
}

// LeftJoinBrand adds a LEFT JOIN to Product.Brand
func (c *jProductSelectBuilder) LeftJoinBrand(
	AsBrand string,
) *jProductSelectBuilder {
	dialect := runtime.GetDialect()
	on := ""
	localTable := JProductModel.Table()
	if c.as != "" {
		localTable = dialect.QuoteIdent(c.as)
	}
	foreiTable := dialect.QuoteIdent(JBrandModel.Table())
	if AsBrand != "" {
		foreiTable = dialect.QuoteIdent(AsBrand)
	}

	on += fmt.Sprintf("%v.%v = %v.%v",
		localTable, dialect.QuoteIdent("brand_id"),
		foreiTable, dialect.QuoteIdent("id"),
	)

	if AsBrand == "" {
		return c.LeftJoin(dbr.I(JBrandModel.Table()), on)
	}
	return c.LeftJoin(dbr.I(JBrandModel.Table()).As(AsBrand), on)
}

// RightJoinBrand adds a RIGHT JOIN to Product.Brand
func (c *jProductSelectBuilder) RightJoinBrand(
	AsBrand string,
) *jProductSelectBuilder {
	dialect := runtime.GetDialect()
	on := ""
	localTable := dialect.QuoteIdent(JProductModel.Table())
	if c.as != "" {
		localTable = dialect.QuoteIdent(c.as)
	}
	foreiTable := dialect.QuoteIdent(JBrandModel.Table())
	if AsBrand != "" {
		foreiTable = dialect.QuoteIdent(AsBrand)
	}

	on += fmt.Sprintf("%v.%v = %v.%v",
		localTable, dialect.QuoteIdent("brand_id"),
		foreiTable, dialect.QuoteIdent("id"),
	)

	if AsBrand == "" {
		return c.RightJoin(dbr.I(JBrandModel.Table()), on)
	}
	return c.RightJoin(dbr.I(JBrandModel.Table()).As(AsBrand), on)
}

// FullJoinBrand adds a FULL JOIN to Product.Brand
func (c *jProductSelectBuilder) FullJoinBrand(
	AsBrand string,
) *jProductSelectBuilder {
	dialect := runtime.GetDialect()
	on := ""
	localTable := dialect.QuoteIdent(JProductModel.Table())
	if c.as != "" {
		localTable = dialect.QuoteIdent(c.as)
	}
	foreiTable := dialect.QuoteIdent(JBrandModel.Table())
	if AsBrand != "" {
		foreiTable = dialect.QuoteIdent(AsBrand)
	}

	on += fmt.Sprintf("%v.%v = %v.%v",
		localTable, dialect.QuoteIdent("brand_id"),
		foreiTable, dialect.QuoteIdent("id"),
	)

	if AsBrand == "" {
		return c.FullJoin(dbr.I(JBrandModel.Table()), on)
	}
	return c.FullJoin(dbr.I(JBrandModel.Table()).As(AsBrand), on)
}

// Brand reads associated object
func (g *Product) Brand(db dbr.SessionRunner) (*Brand, error) {
	q := JBrand(db).Select()
	q = q.Where(

		JBrandModel.ID.Eq(g.BrandID),
	)
	return q.Read()
}

// SetBrand copies pk values to this object properties
func (g *Product) SetBrand(o *Brand) {

	g.BrandID = &o.ID

}

// UnsetBrand set defaults values to this object properties
func (g *Product) UnsetBrand() {

	var def0 *int64

	g.BrandID = def0

	g.brand = nil

}

// JoinBrand2 adds a JOIN to Product.Brand2
func (c *jProductSelectBuilder) JoinBrand2(
	AsBrand2 string,
) *jProductSelectBuilder {
	dialect := runtime.GetDialect()
	on := ""
	localTable := dialect.QuoteIdent(JProductModel.Table())
	if c.as != "" {
		localTable = dialect.QuoteIdent(c.as)
	}
	foreiTable := dialect.QuoteIdent(JBrandModel.Table())
	if AsBrand2 != "" {
		foreiTable = dialect.QuoteIdent(AsBrand2)
	}

	on += fmt.Sprintf("%v.%v = %v.%v",
		localTable, dialect.QuoteIdent("brand2_id"),
		foreiTable, dialect.QuoteIdent("id"),
	)

	if AsBrand2 == "" {
		return c.Join(dbr.I(JBrandModel.Table()), on)
	}
	return c.Join(dbr.I(JBrandModel.Table()).As(AsBrand2), on)
}

// LeftJoinBrand2 adds a LEFT JOIN to Product.Brand2
func (c *jProductSelectBuilder) LeftJoinBrand2(
	AsBrand2 string,
) *jProductSelectBuilder {
	dialect := runtime.GetDialect()
	on := ""
	localTable := JProductModel.Table()
	if c.as != "" {
		localTable = dialect.QuoteIdent(c.as)
	}
	foreiTable := dialect.QuoteIdent(JBrandModel.Table())
	if AsBrand2 != "" {
		foreiTable = dialect.QuoteIdent(AsBrand2)
	}

	on += fmt.Sprintf("%v.%v = %v.%v",
		localTable, dialect.QuoteIdent("brand2_id"),
		foreiTable, dialect.QuoteIdent("id"),
	)

	if AsBrand2 == "" {
		return c.LeftJoin(dbr.I(JBrandModel.Table()), on)
	}
	return c.LeftJoin(dbr.I(JBrandModel.Table()).As(AsBrand2), on)
}

// RightJoinBrand2 adds a RIGHT JOIN to Product.Brand2
func (c *jProductSelectBuilder) RightJoinBrand2(
	AsBrand2 string,
) *jProductSelectBuilder {
	dialect := runtime.GetDialect()
	on := ""
	localTable := dialect.QuoteIdent(JProductModel.Table())
	if c.as != "" {
		localTable = dialect.QuoteIdent(c.as)
	}
	foreiTable := dialect.QuoteIdent(JBrandModel.Table())
	if AsBrand2 != "" {
		foreiTable = dialect.QuoteIdent(AsBrand2)
	}

	on += fmt.Sprintf("%v.%v = %v.%v",
		localTable, dialect.QuoteIdent("brand2_id"),
		foreiTable, dialect.QuoteIdent("id"),
	)

	if AsBrand2 == "" {
		return c.RightJoin(dbr.I(JBrandModel.Table()), on)
	}
	return c.RightJoin(dbr.I(JBrandModel.Table()).As(AsBrand2), on)
}

// FullJoinBrand2 adds a FULL JOIN to Product.Brand2
func (c *jProductSelectBuilder) FullJoinBrand2(
	AsBrand2 string,
) *jProductSelectBuilder {
	dialect := runtime.GetDialect()
	on := ""
	localTable := dialect.QuoteIdent(JProductModel.Table())
	if c.as != "" {
		localTable = dialect.QuoteIdent(c.as)
	}
	foreiTable := dialect.QuoteIdent(JBrandModel.Table())
	if AsBrand2 != "" {
		foreiTable = dialect.QuoteIdent(AsBrand2)
	}

	on += fmt.Sprintf("%v.%v = %v.%v",
		localTable, dialect.QuoteIdent("brand2_id"),
		foreiTable, dialect.QuoteIdent("id"),
	)

	if AsBrand2 == "" {
		return c.FullJoin(dbr.I(JBrandModel.Table()), on)
	}
	return c.FullJoin(dbr.I(JBrandModel.Table()).As(AsBrand2), on)
}

// Brand2 reads associated object
func (g *Product) Brand2(db dbr.SessionRunner) (*Brand, error) {
	q := JBrand(db).Select()
	q = q.Where(

		JBrandModel.ID.Eq(g.Brand2ID),
	)
	return q.Read()
}

// SetBrand2 copies pk values to this object properties
func (g *Product) SetBrand2(o *Brand) {

	g.Brand2ID = &o.ID

}

// UnsetBrand2 set defaults values to this object properties
func (g *Product) UnsetBrand2() {

	var def0 *int64

	g.Brand2ID = def0

	g.brand2 = nil

}

// JoinMaster adds a JOIN to Product.Master
func (c *jProductSelectBuilder) JoinMaster(
	AsMaster string,
) *jProductSelectBuilder {
	dialect := runtime.GetDialect()
	on := ""
	localTable := dialect.QuoteIdent(JProductModel.Table())
	if c.as != "" {
		localTable = dialect.QuoteIdent(c.as)
	}
	foreiTable := dialect.QuoteIdent(JProductModel.Table())
	if AsMaster != "" {
		foreiTable = dialect.QuoteIdent(AsMaster)
	}

	on += fmt.Sprintf("%v.%v = %v.%v",
		localTable, dialect.QuoteIdent("master_id"),
		foreiTable, dialect.QuoteIdent("id"),
	)

	if AsMaster == "" {
		return c.Join(dbr.I(JProductModel.Table()), on)
	}
	return c.Join(dbr.I(JProductModel.Table()).As(AsMaster), on)
}

// LeftJoinMaster adds a LEFT JOIN to Product.Master
func (c *jProductSelectBuilder) LeftJoinMaster(
	AsMaster string,
) *jProductSelectBuilder {
	dialect := runtime.GetDialect()
	on := ""
	localTable := JProductModel.Table()
	if c.as != "" {
		localTable = dialect.QuoteIdent(c.as)
	}
	foreiTable := dialect.QuoteIdent(JProductModel.Table())
	if AsMaster != "" {
		foreiTable = dialect.QuoteIdent(AsMaster)
	}

	on += fmt.Sprintf("%v.%v = %v.%v",
		localTable, dialect.QuoteIdent("master_id"),
		foreiTable, dialect.QuoteIdent("id"),
	)

	if AsMaster == "" {
		return c.LeftJoin(dbr.I(JProductModel.Table()), on)
	}
	return c.LeftJoin(dbr.I(JProductModel.Table()).As(AsMaster), on)
}

// RightJoinMaster adds a RIGHT JOIN to Product.Master
func (c *jProductSelectBuilder) RightJoinMaster(
	AsMaster string,
) *jProductSelectBuilder {
	dialect := runtime.GetDialect()
	on := ""
	localTable := dialect.QuoteIdent(JProductModel.Table())
	if c.as != "" {
		localTable = dialect.QuoteIdent(c.as)
	}
	foreiTable := dialect.QuoteIdent(JProductModel.Table())
	if AsMaster != "" {
		foreiTable = dialect.QuoteIdent(AsMaster)
	}

	on += fmt.Sprintf("%v.%v = %v.%v",
		localTable, dialect.QuoteIdent("master_id"),
		foreiTable, dialect.QuoteIdent("id"),
	)

	if AsMaster == "" {
		return c.RightJoin(dbr.I(JProductModel.Table()), on)
	}
	return c.RightJoin(dbr.I(JProductModel.Table()).As(AsMaster), on)
}

// FullJoinMaster adds a FULL JOIN to Product.Master
func (c *jProductSelectBuilder) FullJoinMaster(
	AsMaster string,
) *jProductSelectBuilder {
	dialect := runtime.GetDialect()
	on := ""
	localTable := dialect.QuoteIdent(JProductModel.Table())
	if c.as != "" {
		localTable = dialect.QuoteIdent(c.as)
	}
	foreiTable := dialect.QuoteIdent(JProductModel.Table())
	if AsMaster != "" {
		foreiTable = dialect.QuoteIdent(AsMaster)
	}

	on += fmt.Sprintf("%v.%v = %v.%v",
		localTable, dialect.QuoteIdent("master_id"),
		foreiTable, dialect.QuoteIdent("id"),
	)

	if AsMaster == "" {
		return c.FullJoin(dbr.I(JProductModel.Table()), on)
	}
	return c.FullJoin(dbr.I(JProductModel.Table()).As(AsMaster), on)
}

// Master reads associated object
func (g *Product) Master(db dbr.SessionRunner) (*Product, error) {
	q := JProduct(db).Select()
	q = q.Where(

		JProductModel.ID.Eq(g.MasterID),
	)
	return q.Read()
}

// SetMaster copies pk values to this object properties
func (g *Product) SetMaster(o *Product) {

	g.MasterID = &o.ID

}

// UnsetMaster set defaults values to this object properties
func (g *Product) UnsetMaster() {

	var def0 *int64

	g.MasterID = def0

	g.master = nil

}

// Categories returns a query builder to select Categories linked to this Product
func (g *Product) Categories(db dbr.SessionRunner,
	AsCategory, AsCategoryproductsToProductcategories, AsProduct string,
) *jCategorySelectBuilder {

	leftTable := JCategoryModel.Table()
	if AsCategory != "" {
		leftTable = AsCategory
	}

	query := JCategory(db).Select(leftTable + ".*")

	midTable := JCategoryproductsToProductcategoriesModel.Table()
	midWhat := midTable
	if AsCategoryproductsToProductcategories != "" {
		midWhat = fmt.Sprintf("%v as %v", midTable, AsCategoryproductsToProductcategories)
	}

	{
		on := ""
		if AsCategoryproductsToProductcategories != "" {
			midTable = AsCategoryproductsToProductcategories
		}

		on += fmt.Sprintf("%v.%v = %v.%v",
			midTable, "category_id",
			leftTable, "id",
		)

		query = query.Join(midWhat, on)
	}

	rightTable := JProductModel.Table()
	rightWhat := rightTable
	if AsProduct != "" {
		rightWhat = fmt.Sprintf("%v as %v", rightTable, AsProduct)
	}
	{
		on := ""
		if AsProduct != "" {
			rightTable = AsProduct
		}

		on += fmt.Sprintf("%v.%v = %v.%v",
			midTable, "product_id",
			rightTable, "id",
		)

		query = query.Join(rightWhat, on)
	}

	return query
}

//LinkWithCategories writes new links with Product.
func (g *Product) LinkWithCategories(db dbr.SessionRunner, items ...*Category) (sql.Result, error) {
	toInsert := []*CategoryproductsToProductcategories{}
	for _, item := range items {
		toInsert = append(toInsert, &CategoryproductsToProductcategories{

			CategoryID: item.ID,

			ProductID: g.ID,
		})
	}
	return JCategoryproductsToProductcategories(db).Insert(toInsert...)
}

//UnlinkWithCategories deletes given existing links with Product.
func (g *Product) UnlinkWithCategories(db dbr.SessionRunner, items ...*Category) (sql.Result, error) {
	toDelete := []*CategoryproductsToProductcategories{}
	for _, item := range items {
		toDelete = append(toDelete, &CategoryproductsToProductcategories{

			CategoryID: item.ID,

			ProductID: g.ID,
		})
	}
	return JCategoryproductsToProductcategories(db).DeleteAll(toDelete...)
}

//UnlinkAllCategories deletes all existing links with Product.
func (g *Product) UnlinkAllCategories(db dbr.SessionRunner) (sql.Result, error) {
	return JCategoryproductsToProductcategories(db).Delete().Where(

		JCategoryproductsToProductcategoriesModel.ProductID.Eq(g.ID),
	).Exec()
}

//SetCategories replaces existing links with Product.
func (g *Product) SetCategories(db dbr.SessionRunner, items ...*Category) (sql.Result, error) {
	if res, err := g.UnlinkAllCategories(db); err != nil {
		return res, err
	}
	return g.LinkWithCategories(db, items...)
}

type jCategorySetup struct {
	Name       string
	CreateStmt string
	DropStmt   string
}

//Create applies the create table command to te underlying connection.
func (c jCategorySetup) Create(db *dbr.Connection) error {
	_, err := db.Exec(c.CreateStmt)
	return err
}

//Drop applies the drop table command to te underlying connection.
func (c jCategorySetup) Drop(db *dbr.Connection) error {
	_, err := db.Exec(c.DropStmt)
	return err
}

// JCategorySetup helps to create/drop the schema
func JCategorySetup() runtime.Setuper {
	driver := runtime.GetCurrentDriver()

	var create string
	var drop string

	if driver == drivers.Sqlite {
		create = `CREATE TABLE IF NOT EXISTS category (
id INTEGER PRIMARY KEY AUTOINCREMENT,
name TEXT

)`
	} else if driver == drivers.Mysql {
		create = `CREATE TABLE IF NOT EXISTS category (
id INTEGER NOT NULL AUTO_INCREMENT,
name TEXT,
PRIMARY KEY (id) 

)`
	} else if driver == drivers.Pgsql {
		create = `CREATE TABLE IF NOT EXISTS category (
id INTEGER,
name TEXT

)`
	}

	if driver == drivers.Sqlite {
		drop = `DROP TABLE IF EXISTS category`
	} else if driver == drivers.Mysql {
		drop = `DROP TABLE IF EXISTS category`
	} else if driver == drivers.Pgsql {
		drop = `DROP TABLE IF EXISTS category`
	}

	return jCategorySetup{
		Name:       `category`,
		CreateStmt: create,
		DropStmt:   drop,
	}
}

// jCategoryModel provides helper to work with Category data provider
type jCategoryModel struct {
	as string

	ID builder.ValuePropertyMeta

	Name builder.ValuePropertyMeta

	Products builder.RelPropertyMeta
}

// Eq provided items.
func (j jCategoryModel) Eq(s ...*Category) dbr.Builder {
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
func (j jCategoryModel) In(s ...*Category) dbr.Builder {
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
func (j jCategoryModel) As(as string) jCategoryModel {
	j.as = as

	j.ID.TableAlias = as

	j.Name.TableAlias = as

	// j.Products.TableAlias = as

	return j
}

// Table returns the sql table name
func (j jCategoryModel) Table() string {
	return "category"
}

// Alias returns the current alias
func (j jCategoryModel) Alias() string {
	if j.as == "" {
		return j.Table()
	}
	return j.as
}

// Properties returns a map of property name => meta
func (j jCategoryModel) Properties() map[string]builder.MetaProvider {
	ret := map[string]builder.MetaProvider{}

	ret["ID"] = j.ID

	ret["Name"] = j.Name

	ret["Products"] = j.Products

	return ret
}

// Fields returns given sql fields with appropriate aliasing.
func (j jCategoryModel) Fields(ins ...string) []string {
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

// JCategoryModel provides helper to work with Category data provider
var JCategoryModel = jCategoryModel{

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

	Products: builder.NewRelMeta(
		`products`, `[]*Product`,
		``, `Product.categories`, ``,
	),
}

type jCategoryDeleteBuilder struct {
	*builder.DeleteBuilder
}

//Build builds the sql string using current dialect into given bufer
func (c *jCategoryDeleteBuilder) Build(b dbr.Buffer) error {
	return c.DeleteBuilder.Build(runtime.GetDialect(), b)
}

//String returns the sql string for current dialect. It returns empty string if the build returns an error.
func (c *jCategoryDeleteBuilder) String() string {
	b := dbr.NewBuffer()
	if err := c.Build(b); err != nil {
		return ""
	}
	return b.String()
}

//Where returns a jCategoryDeleteBuilder instead of builder.DeleteBuilder.
func (c *jCategoryDeleteBuilder) Where(query interface{}, value ...interface{}) *jCategoryDeleteBuilder {
	c.DeleteBuilder.Where(query, value...)
	return c
}

type jCategorySelectBuilder struct {
	as string
	*builder.SelectBuilder
}

//Build builds the sql string using current dialect into given bufer
func (c *jCategorySelectBuilder) Build(b dbr.Buffer) error {
	return c.SelectBuilder.Build(runtime.GetDialect(), b)
}

//String returns the sql string for current dialect. It returns empty string if the build returns an error.
func (c *jCategorySelectBuilder) String() string {
	b := dbr.NewBuffer()
	if err := c.Build(b); err != nil {
		return ""
	}
	return b.String()
}

//Read evaluates current select query and load the results into a Category
func (c *jCategorySelectBuilder) Read() (*Category, error) {
	var one Category
	err := c.LoadStruct(&one)
	return &one, err
}

//ReadAll evaluates current select query and load the results into a slice of Category
func (c *jCategorySelectBuilder) ReadAll() ([]*Category, error) {
	var all []*Category
	_, err := c.LoadStructs(&all)
	return all, err
}

//Where returns a jCategorySelectBuilder instead of builder.SelectBuilder.
func (c *jCategorySelectBuilder) Where(query interface{}, value ...interface{}) *jCategorySelectBuilder {
	c.SelectBuilder.Where(query, value...)
	return c
}

//Join returns a jCategorySelectBuilder instead of builder.SelectBuilder.
func (c *jCategorySelectBuilder) Join(table, on interface{}) *jCategorySelectBuilder {
	c.SelectBuilder.Join(table, on)
	return c
}

//LeftJoin returns a jCategorySelectBuilder instead of builder.SelectBuilder.
func (c *jCategorySelectBuilder) LeftJoin(table, on interface{}) *jCategorySelectBuilder {
	c.SelectBuilder.LeftJoin(table, on)
	return c
}

//RightJoin returns a jCategorySelectBuilder instead of builder.SelectBuilder.
func (c *jCategorySelectBuilder) RightJoin(table, on interface{}) *jCategorySelectBuilder {
	c.SelectBuilder.RightJoin(table, on)
	return c
}

//FullJoin returns a jCategorySelectBuilder instead of builder.SelectBuilder.
func (c *jCategorySelectBuilder) FullJoin(table, on interface{}) *jCategorySelectBuilder {
	c.SelectBuilder.FullJoin(table, on)
	return c
}

//Distinct returns a jCategorySelectBuilder instead of builder.SelectBuilder.
func (c *jCategorySelectBuilder) Distinct() *jCategorySelectBuilder {
	c.SelectBuilder.Distinct()
	return c
}

// JCategory provides a basic querier
func JCategory(db dbr.SessionRunner) jCategoryQuerier {
	return jCategoryQuerier{
		db: db,
	}
}

type jCategoryQuerier struct {
	db dbr.SessionRunner
	as string
}

//As set alias prior building.
func (c jCategoryQuerier) As(as string) jCategoryQuerier {
	c.as = as
	return c
}

//Model returns a model with appropriate aliasing.
func (c jCategoryQuerier) Model() jCategoryModel {
	return JCategoryModel.As(c.as)
}

//Select returns a Category Select Builder.
func (c jCategoryQuerier) Select(what ...string) *jCategorySelectBuilder {
	m := c.Model()
	if len(what) == 0 {
		what = m.Fields("*")
	}
	dialect := runtime.GetDialect()
	from := dialect.QuoteIdent(m.Table())
	if m.Alias() != "" && m.Alias() != m.Table() {
		from = fmt.Sprintf("%v as %v", from, dialect.QuoteIdent(m.Alias()))
	}
	return &jCategorySelectBuilder{
		as: c.as,
		SelectBuilder: &builder.SelectBuilder{
			SelectBuilder: c.db.Select(what...).From(from),
		},
	}
}

//Count returns a Category Select Builder to count given expressions.
func (c jCategoryQuerier) Count(what ...string) *jCategorySelectBuilder {
	if len(what) == 0 {
		what = append(what, "*")
	}
	return c.Select("COUNT(" + strings.Join(what, ",") + ")")
}

// Insert a new Category, if it has autoincrement primary key, the value will be set.
// It stops on first error.
func (c jCategoryQuerier) Insert(items ...*Category) (sql.Result, error) {
	var res sql.Result
	var err error
	for _, data := range items {
		res, err = c.db.InsertInto(JCategoryModel.Table()).Columns(

			`name`,
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
func (c jCategoryQuerier) InsertBulk(items ...*Category) error {
	panic("todo")
}

// Update a Category. It stops on first error.
func (c jCategoryQuerier) Update(items ...*Category) (sql.Result, error) {
	var res sql.Result
	var err error
	for _, data := range items {
		res, err = c.db.Update(JCategoryModel.Table()).
			Set(`name`, data.Name).
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
func (c jCategoryQuerier) UpdateBulk(items ...*Category) error {
	panic("todo")
}

//Delete returns a delete builder
func (c jCategoryQuerier) Delete() *jCategoryDeleteBuilder {
	return &jCategoryDeleteBuilder{
		&builder.DeleteBuilder{
			DeleteBuilder: c.db.DeleteFrom(JCategoryModel.Table()),
		},
	}
}

//DeleteByPk deletes one Category by its PKs
func (c jCategoryQuerier) DeleteByPk(ID int64) error {
	_, err := c.Delete().Where(

		JCategoryModel.ID.Eq(ID),
	).Exec()
	return err
}

// DeleteAll given Category
func (c jCategoryQuerier) DeleteAll(items ...*Category) (sql.Result, error) {
	q := c.Delete()
	for _, item := range items {
		q = q.Where(
			dbr.Or(
				dbr.And(

					JCategoryModel.ID.Eq(item.ID),
				),
			),
		)
	}
	return q.Exec()
}

//Find one Category using its PKs
func (c jCategoryQuerier) Find(ID int64) (*Category, error) {
	return c.Select().Where(

		JCategoryModel.ID.Eq(ID),
	).Read()
}

// Products returns a query builder to select Products linked to this Category
func (g *Category) Products(db dbr.SessionRunner,
	AsProduct, AsCategoryproductsToProductcategories, AsCategory string,
) *jProductSelectBuilder {

	leftTable := JProductModel.Table()
	if AsProduct != "" {
		leftTable = AsProduct
	}

	query := JProduct(db).Select(leftTable + ".*")

	midTable := JCategoryproductsToProductcategoriesModel.Table()
	midWhat := midTable
	if AsCategoryproductsToProductcategories != "" {
		midWhat = fmt.Sprintf("%v as %v", midTable, AsCategoryproductsToProductcategories)
	}

	{
		on := ""
		if AsCategoryproductsToProductcategories != "" {
			midTable = AsCategoryproductsToProductcategories
		}

		on += fmt.Sprintf("%v.%v = %v.%v",
			midTable, "product_id",
			leftTable, "id",
		)

		query = query.Join(midWhat, on)
	}

	rightTable := JCategoryModel.Table()
	rightWhat := rightTable
	if AsCategory != "" {
		rightWhat = fmt.Sprintf("%v as %v", rightTable, AsCategory)
	}
	{
		on := ""
		if AsCategory != "" {
			rightTable = AsCategory
		}

		on += fmt.Sprintf("%v.%v = %v.%v",
			midTable, "category_id",
			rightTable, "id",
		)

		query = query.Join(rightWhat, on)
	}

	return query
}

//LinkWithProducts writes new links with Category.
func (g *Category) LinkWithProducts(db dbr.SessionRunner, items ...*Product) (sql.Result, error) {
	toInsert := []*CategoryproductsToProductcategories{}
	for _, item := range items {
		toInsert = append(toInsert, &CategoryproductsToProductcategories{

			ProductID: item.ID,

			CategoryID: g.ID,
		})
	}
	return JCategoryproductsToProductcategories(db).Insert(toInsert...)
}

//UnlinkWithProducts deletes given existing links with Category.
func (g *Category) UnlinkWithProducts(db dbr.SessionRunner, items ...*Product) (sql.Result, error) {
	toDelete := []*CategoryproductsToProductcategories{}
	for _, item := range items {
		toDelete = append(toDelete, &CategoryproductsToProductcategories{

			ProductID: item.ID,

			CategoryID: g.ID,
		})
	}
	return JCategoryproductsToProductcategories(db).DeleteAll(toDelete...)
}

//UnlinkAllProducts deletes all existing links with Category.
func (g *Category) UnlinkAllProducts(db dbr.SessionRunner) (sql.Result, error) {
	return JCategoryproductsToProductcategories(db).Delete().Where(

		JCategoryproductsToProductcategoriesModel.CategoryID.Eq(g.ID),
	).Exec()
}

//SetProducts replaces existing links with Category.
func (g *Category) SetProducts(db dbr.SessionRunner, items ...*Product) (sql.Result, error) {
	if res, err := g.UnlinkAllProducts(db); err != nil {
		return res, err
	}
	return g.LinkWithProducts(db, items...)
}

type jBrandSetup struct {
	Name       string
	CreateStmt string
	DropStmt   string
}

//Create applies the create table command to te underlying connection.
func (c jBrandSetup) Create(db *dbr.Connection) error {
	_, err := db.Exec(c.CreateStmt)
	return err
}

//Drop applies the drop table command to te underlying connection.
func (c jBrandSetup) Drop(db *dbr.Connection) error {
	_, err := db.Exec(c.DropStmt)
	return err
}

// JBrandSetup helps to create/drop the schema
func JBrandSetup() runtime.Setuper {
	driver := runtime.GetCurrentDriver()

	var create string
	var drop string

	if driver == drivers.Sqlite {
		create = `CREATE TABLE IF NOT EXISTS brand (
id INTEGER PRIMARY KEY AUTOINCREMENT,
name TEXT

)`
	} else if driver == drivers.Mysql {
		create = `CREATE TABLE IF NOT EXISTS brand (
id INTEGER NOT NULL AUTO_INCREMENT,
name TEXT,
PRIMARY KEY (id) 

)`
	} else if driver == drivers.Pgsql {
		create = `CREATE TABLE IF NOT EXISTS brand (
id INTEGER,
name TEXT

)`
	}

	if driver == drivers.Sqlite {
		drop = `DROP TABLE IF EXISTS brand`
	} else if driver == drivers.Mysql {
		drop = `DROP TABLE IF EXISTS brand`
	} else if driver == drivers.Pgsql {
		drop = `DROP TABLE IF EXISTS brand`
	}

	return jBrandSetup{
		Name:       `brand`,
		CreateStmt: create,
		DropStmt:   drop,
	}
}

// jBrandModel provides helper to work with Brand data provider
type jBrandModel struct {
	as string

	ID builder.ValuePropertyMeta

	Name builder.ValuePropertyMeta

	Products builder.RelPropertyMeta

	Products2 builder.RelPropertyMeta
}

// Eq provided items.
func (j jBrandModel) Eq(s ...*Brand) dbr.Builder {
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
func (j jBrandModel) In(s ...*Brand) dbr.Builder {
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
func (j jBrandModel) As(as string) jBrandModel {
	j.as = as

	j.ID.TableAlias = as

	j.Name.TableAlias = as

	// j.Products.TableAlias = as

	// j.Products2.TableAlias = as

	return j
}

// Table returns the sql table name
func (j jBrandModel) Table() string {
	return "brand"
}

// Alias returns the current alias
func (j jBrandModel) Alias() string {
	if j.as == "" {
		return j.Table()
	}
	return j.as
}

// Properties returns a map of property name => meta
func (j jBrandModel) Properties() map[string]builder.MetaProvider {
	ret := map[string]builder.MetaProvider{}

	ret["ID"] = j.ID

	ret["Name"] = j.Name

	ret["Products"] = j.Products

	ret["Products2"] = j.Products2

	return ret
}

// Fields returns given sql fields with appropriate aliasing.
func (j jBrandModel) Fields(ins ...string) []string {
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

// JBrandModel provides helper to work with Brand data provider
var JBrandModel = jBrandModel{

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

	Products: builder.NewRelMeta(
		`products`, `[]*Product`,
		``, `Product.brand`, ``,
	),

	Products2: builder.NewRelMeta(
		`products2`, `[]*Product`,
		``, `Product.brand2`, ``,
	),
}

type jBrandDeleteBuilder struct {
	*builder.DeleteBuilder
}

//Build builds the sql string using current dialect into given bufer
func (c *jBrandDeleteBuilder) Build(b dbr.Buffer) error {
	return c.DeleteBuilder.Build(runtime.GetDialect(), b)
}

//String returns the sql string for current dialect. It returns empty string if the build returns an error.
func (c *jBrandDeleteBuilder) String() string {
	b := dbr.NewBuffer()
	if err := c.Build(b); err != nil {
		return ""
	}
	return b.String()
}

//Where returns a jBrandDeleteBuilder instead of builder.DeleteBuilder.
func (c *jBrandDeleteBuilder) Where(query interface{}, value ...interface{}) *jBrandDeleteBuilder {
	c.DeleteBuilder.Where(query, value...)
	return c
}

type jBrandSelectBuilder struct {
	as string
	*builder.SelectBuilder
}

//Build builds the sql string using current dialect into given bufer
func (c *jBrandSelectBuilder) Build(b dbr.Buffer) error {
	return c.SelectBuilder.Build(runtime.GetDialect(), b)
}

//String returns the sql string for current dialect. It returns empty string if the build returns an error.
func (c *jBrandSelectBuilder) String() string {
	b := dbr.NewBuffer()
	if err := c.Build(b); err != nil {
		return ""
	}
	return b.String()
}

//Read evaluates current select query and load the results into a Brand
func (c *jBrandSelectBuilder) Read() (*Brand, error) {
	var one Brand
	err := c.LoadStruct(&one)
	return &one, err
}

//ReadAll evaluates current select query and load the results into a slice of Brand
func (c *jBrandSelectBuilder) ReadAll() ([]*Brand, error) {
	var all []*Brand
	_, err := c.LoadStructs(&all)
	return all, err
}

//Where returns a jBrandSelectBuilder instead of builder.SelectBuilder.
func (c *jBrandSelectBuilder) Where(query interface{}, value ...interface{}) *jBrandSelectBuilder {
	c.SelectBuilder.Where(query, value...)
	return c
}

//Join returns a jBrandSelectBuilder instead of builder.SelectBuilder.
func (c *jBrandSelectBuilder) Join(table, on interface{}) *jBrandSelectBuilder {
	c.SelectBuilder.Join(table, on)
	return c
}

//LeftJoin returns a jBrandSelectBuilder instead of builder.SelectBuilder.
func (c *jBrandSelectBuilder) LeftJoin(table, on interface{}) *jBrandSelectBuilder {
	c.SelectBuilder.LeftJoin(table, on)
	return c
}

//RightJoin returns a jBrandSelectBuilder instead of builder.SelectBuilder.
func (c *jBrandSelectBuilder) RightJoin(table, on interface{}) *jBrandSelectBuilder {
	c.SelectBuilder.RightJoin(table, on)
	return c
}

//FullJoin returns a jBrandSelectBuilder instead of builder.SelectBuilder.
func (c *jBrandSelectBuilder) FullJoin(table, on interface{}) *jBrandSelectBuilder {
	c.SelectBuilder.FullJoin(table, on)
	return c
}

//Distinct returns a jBrandSelectBuilder instead of builder.SelectBuilder.
func (c *jBrandSelectBuilder) Distinct() *jBrandSelectBuilder {
	c.SelectBuilder.Distinct()
	return c
}

// JBrand provides a basic querier
func JBrand(db dbr.SessionRunner) jBrandQuerier {
	return jBrandQuerier{
		db: db,
	}
}

type jBrandQuerier struct {
	db dbr.SessionRunner
	as string
}

//As set alias prior building.
func (c jBrandQuerier) As(as string) jBrandQuerier {
	c.as = as
	return c
}

//Model returns a model with appropriate aliasing.
func (c jBrandQuerier) Model() jBrandModel {
	return JBrandModel.As(c.as)
}

//Select returns a Brand Select Builder.
func (c jBrandQuerier) Select(what ...string) *jBrandSelectBuilder {
	m := c.Model()
	if len(what) == 0 {
		what = m.Fields("*")
	}
	dialect := runtime.GetDialect()
	from := dialect.QuoteIdent(m.Table())
	if m.Alias() != "" && m.Alias() != m.Table() {
		from = fmt.Sprintf("%v as %v", from, dialect.QuoteIdent(m.Alias()))
	}
	return &jBrandSelectBuilder{
		as: c.as,
		SelectBuilder: &builder.SelectBuilder{
			SelectBuilder: c.db.Select(what...).From(from),
		},
	}
}

//Count returns a Brand Select Builder to count given expressions.
func (c jBrandQuerier) Count(what ...string) *jBrandSelectBuilder {
	if len(what) == 0 {
		what = append(what, "*")
	}
	return c.Select("COUNT(" + strings.Join(what, ",") + ")")
}

// Insert a new Brand, if it has autoincrement primary key, the value will be set.
// It stops on first error.
func (c jBrandQuerier) Insert(items ...*Brand) (sql.Result, error) {
	var res sql.Result
	var err error
	for _, data := range items {
		res, err = c.db.InsertInto(JBrandModel.Table()).Columns(

			`name`,
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
func (c jBrandQuerier) InsertBulk(items ...*Brand) error {
	panic("todo")
}

// Update a Brand. It stops on first error.
func (c jBrandQuerier) Update(items ...*Brand) (sql.Result, error) {
	var res sql.Result
	var err error
	for _, data := range items {
		res, err = c.db.Update(JBrandModel.Table()).
			Set(`name`, data.Name).
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
func (c jBrandQuerier) UpdateBulk(items ...*Brand) error {
	panic("todo")
}

//Delete returns a delete builder
func (c jBrandQuerier) Delete() *jBrandDeleteBuilder {
	return &jBrandDeleteBuilder{
		&builder.DeleteBuilder{
			DeleteBuilder: c.db.DeleteFrom(JBrandModel.Table()),
		},
	}
}

//DeleteByPk deletes one Brand by its PKs
func (c jBrandQuerier) DeleteByPk(ID int64) error {
	_, err := c.Delete().Where(

		JBrandModel.ID.Eq(ID),
	).Exec()
	return err
}

// DeleteAll given Brand
func (c jBrandQuerier) DeleteAll(items ...*Brand) (sql.Result, error) {
	q := c.Delete()
	for _, item := range items {
		q = q.Where(
			dbr.Or(
				dbr.And(

					JBrandModel.ID.Eq(item.ID),
				),
			),
		)
	}
	return q.Exec()
}

//Find one Brand using its PKs
func (c jBrandQuerier) Find(ID int64) (*Brand, error) {
	return c.Select().Where(

		JBrandModel.ID.Eq(ID),
	).Read()
}

// Products returns a query builder to select Products linked to this Brand
func (g *Brand) Products(db dbr.SessionRunner,
	AsProduct, AsBrand string,
) *jProductSelectBuilder {

	leftTable := JProductModel.Table()
	if AsProduct != "" {
		leftTable = AsProduct
	}

	query := JProduct(db).Select(leftTable + ".*")

	rightTable := JBrandModel.Table()
	rightWhat := rightTable
	if AsBrand != "" {
		rightWhat = fmt.Sprintf("%v as %v", rightTable, AsBrand)
	}

	on := ""
	if AsBrand != "" {
		rightTable = AsBrand
	}

	on += fmt.Sprintf("%v.%v = %v.%v",
		leftTable, "brand_id",
		rightTable, "id",
	)

	query = query.Join(rightWhat, on)

	return query
}

// Products2 returns a query builder to select Products2 linked to this Brand
func (g *Brand) Products2(db dbr.SessionRunner,
	AsProduct, AsBrand string,
) *jProductSelectBuilder {

	leftTable := JProductModel.Table()
	if AsProduct != "" {
		leftTable = AsProduct
	}

	query := JProduct(db).Select(leftTable + ".*")

	rightTable := JBrandModel.Table()
	rightWhat := rightTable
	if AsBrand != "" {
		rightWhat = fmt.Sprintf("%v as %v", rightTable, AsBrand)
	}

	on := ""
	if AsBrand != "" {
		rightTable = AsBrand
	}

	on += fmt.Sprintf("%v.%v = %v.%v",
		leftTable, "brand2_id",
		rightTable, "id",
	)

	query = query.Join(rightWhat, on)

	return query
}

type jCategoryproductsToProductcategoriesSetup struct {
	Name       string
	CreateStmt string
	DropStmt   string
}

//Create applies the create table command to te underlying connection.
func (c jCategoryproductsToProductcategoriesSetup) Create(db *dbr.Connection) error {
	_, err := db.Exec(c.CreateStmt)
	return err
}

//Drop applies the drop table command to te underlying connection.
func (c jCategoryproductsToProductcategoriesSetup) Drop(db *dbr.Connection) error {
	_, err := db.Exec(c.DropStmt)
	return err
}

// JCategoryproductsToProductcategoriesSetup helps to create/drop the schema
func JCategoryproductsToProductcategoriesSetup() runtime.Setuper {
	driver := runtime.GetCurrentDriver()

	var create string
	var drop string

	if driver == drivers.Sqlite {
		create = `CREATE TABLE IF NOT EXISTS category_productstoproduct_categories (
product_id INTEGER,
category_id INTEGER,
PRIMARY KEY (product_id,category_id) 

)`
	} else if driver == drivers.Mysql {
		create = `CREATE TABLE IF NOT EXISTS category_productstoproduct_categories (
product_id INTEGER NOT NULL,
category_id INTEGER NOT NULL,
PRIMARY KEY (product_id,category_id) 

)`
	} else if driver == drivers.Pgsql {
		create = `CREATE TABLE IF NOT EXISTS category_productstoproduct_categories (
product_id INTEGER,
category_id INTEGER

)`
	}

	if driver == drivers.Sqlite {
		drop = `DROP TABLE IF EXISTS category_productstoproduct_categories`
	} else if driver == drivers.Mysql {
		drop = `DROP TABLE IF EXISTS category_productstoproduct_categories`
	} else if driver == drivers.Pgsql {
		drop = `DROP TABLE IF EXISTS category_productstoproduct_categories`
	}

	return jCategoryproductsToProductcategoriesSetup{
		Name:       `category_productstoproduct_categories`,
		CreateStmt: create,
		DropStmt:   drop,
	}
}

// jCategoryproductsToProductcategoriesModel provides helper to work with CategoryproductsToProductcategories data provider
type jCategoryproductsToProductcategoriesModel struct {
	as string

	ProductID builder.ValuePropertyMeta

	CategoryID builder.ValuePropertyMeta
}

// Eq provided items.
func (j jCategoryproductsToProductcategoriesModel) Eq(s ...*CategoryproductsToProductcategories) dbr.Builder {
	ors := []dbr.Builder{}
	for _, t := range s {
		ors = append(ors, dbr.Or(
			dbr.And(

				dbr.Eq(`product_id`, t.ProductID),

				dbr.Eq(`category_id`, t.CategoryID),
			),
		))
	}
	return dbr.And(ors...)
}

// In provided items.
func (j jCategoryproductsToProductcategoriesModel) In(s ...*CategoryproductsToProductcategories) dbr.Builder {
	ors := []dbr.Builder{}
	for _, t := range s {
		ors = append(ors, dbr.Or(
			dbr.And(

				dbr.Eq(`product_id`, t.ProductID),

				dbr.Eq(`category_id`, t.CategoryID),
			),
		))
	}
	return dbr.And(ors...)
}

// As returns a copy with an alias.
func (j jCategoryproductsToProductcategoriesModel) As(as string) jCategoryproductsToProductcategoriesModel {
	j.as = as

	j.ProductID.TableAlias = as

	j.CategoryID.TableAlias = as

	return j
}

// Table returns the sql table name
func (j jCategoryproductsToProductcategoriesModel) Table() string {
	return "category_productstoproduct_categories"
}

// Alias returns the current alias
func (j jCategoryproductsToProductcategoriesModel) Alias() string {
	if j.as == "" {
		return j.Table()
	}
	return j.as
}

// Properties returns a map of property name => meta
func (j jCategoryproductsToProductcategoriesModel) Properties() map[string]builder.MetaProvider {
	ret := map[string]builder.MetaProvider{}

	ret["ProductID"] = j.ProductID

	ret["CategoryID"] = j.CategoryID

	return ret
}

// Fields returns given sql fields with appropriate aliasing.
func (j jCategoryproductsToProductcategoriesModel) Fields(ins ...string) []string {
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

// JCategoryproductsToProductcategoriesModel provides helper to work with CategoryproductsToProductcategories data provider
var JCategoryproductsToProductcategoriesModel = jCategoryproductsToProductcategoriesModel{

	ProductID: builder.NewValueMeta(
		`product_id`, `INTEGER`,
		`ProductID`, `int64`,
		true, false,
	),

	CategoryID: builder.NewValueMeta(
		`category_id`, `INTEGER`,
		`CategoryID`, `int64`,
		true, false,
	),
}

// CategoryproductsToProductcategories is automatically generated to handle a many to many relationship.
type CategoryproductsToProductcategories struct {
	ProductID int64

	CategoryID int64
}

type jCategoryproductsToProductcategoriesDeleteBuilder struct {
	*builder.DeleteBuilder
}

//Build builds the sql string using current dialect into given bufer
func (c *jCategoryproductsToProductcategoriesDeleteBuilder) Build(b dbr.Buffer) error {
	return c.DeleteBuilder.Build(runtime.GetDialect(), b)
}

//String returns the sql string for current dialect. It returns empty string if the build returns an error.
func (c *jCategoryproductsToProductcategoriesDeleteBuilder) String() string {
	b := dbr.NewBuffer()
	if err := c.Build(b); err != nil {
		return ""
	}
	return b.String()
}

//Where returns a jCategoryproductsToProductcategoriesDeleteBuilder instead of builder.DeleteBuilder.
func (c *jCategoryproductsToProductcategoriesDeleteBuilder) Where(query interface{}, value ...interface{}) *jCategoryproductsToProductcategoriesDeleteBuilder {
	c.DeleteBuilder.Where(query, value...)
	return c
}

type jCategoryproductsToProductcategoriesSelectBuilder struct {
	as string
	*builder.SelectBuilder
}

//Build builds the sql string using current dialect into given bufer
func (c *jCategoryproductsToProductcategoriesSelectBuilder) Build(b dbr.Buffer) error {
	return c.SelectBuilder.Build(runtime.GetDialect(), b)
}

//String returns the sql string for current dialect. It returns empty string if the build returns an error.
func (c *jCategoryproductsToProductcategoriesSelectBuilder) String() string {
	b := dbr.NewBuffer()
	if err := c.Build(b); err != nil {
		return ""
	}
	return b.String()
}

//Read evaluates current select query and load the results into a CategoryproductsToProductcategories
func (c *jCategoryproductsToProductcategoriesSelectBuilder) Read() (*CategoryproductsToProductcategories, error) {
	var one CategoryproductsToProductcategories
	err := c.LoadStruct(&one)
	return &one, err
}

//ReadAll evaluates current select query and load the results into a slice of CategoryproductsToProductcategories
func (c *jCategoryproductsToProductcategoriesSelectBuilder) ReadAll() ([]*CategoryproductsToProductcategories, error) {
	var all []*CategoryproductsToProductcategories
	_, err := c.LoadStructs(&all)
	return all, err
}

//Where returns a jCategoryproductsToProductcategoriesSelectBuilder instead of builder.SelectBuilder.
func (c *jCategoryproductsToProductcategoriesSelectBuilder) Where(query interface{}, value ...interface{}) *jCategoryproductsToProductcategoriesSelectBuilder {
	c.SelectBuilder.Where(query, value...)
	return c
}

//Join returns a jCategoryproductsToProductcategoriesSelectBuilder instead of builder.SelectBuilder.
func (c *jCategoryproductsToProductcategoriesSelectBuilder) Join(table, on interface{}) *jCategoryproductsToProductcategoriesSelectBuilder {
	c.SelectBuilder.Join(table, on)
	return c
}

//LeftJoin returns a jCategoryproductsToProductcategoriesSelectBuilder instead of builder.SelectBuilder.
func (c *jCategoryproductsToProductcategoriesSelectBuilder) LeftJoin(table, on interface{}) *jCategoryproductsToProductcategoriesSelectBuilder {
	c.SelectBuilder.LeftJoin(table, on)
	return c
}

//RightJoin returns a jCategoryproductsToProductcategoriesSelectBuilder instead of builder.SelectBuilder.
func (c *jCategoryproductsToProductcategoriesSelectBuilder) RightJoin(table, on interface{}) *jCategoryproductsToProductcategoriesSelectBuilder {
	c.SelectBuilder.RightJoin(table, on)
	return c
}

//FullJoin returns a jCategoryproductsToProductcategoriesSelectBuilder instead of builder.SelectBuilder.
func (c *jCategoryproductsToProductcategoriesSelectBuilder) FullJoin(table, on interface{}) *jCategoryproductsToProductcategoriesSelectBuilder {
	c.SelectBuilder.FullJoin(table, on)
	return c
}

//Distinct returns a jCategoryproductsToProductcategoriesSelectBuilder instead of builder.SelectBuilder.
func (c *jCategoryproductsToProductcategoriesSelectBuilder) Distinct() *jCategoryproductsToProductcategoriesSelectBuilder {
	c.SelectBuilder.Distinct()
	return c
}

// JCategoryproductsToProductcategories provides a basic querier
func JCategoryproductsToProductcategories(db dbr.SessionRunner) jCategoryproductsToProductcategoriesQuerier {
	return jCategoryproductsToProductcategoriesQuerier{
		db: db,
	}
}

type jCategoryproductsToProductcategoriesQuerier struct {
	db dbr.SessionRunner
	as string
}

//As set alias prior building.
func (c jCategoryproductsToProductcategoriesQuerier) As(as string) jCategoryproductsToProductcategoriesQuerier {
	c.as = as
	return c
}

//Model returns a model with appropriate aliasing.
func (c jCategoryproductsToProductcategoriesQuerier) Model() jCategoryproductsToProductcategoriesModel {
	return JCategoryproductsToProductcategoriesModel.As(c.as)
}

//Select returns a CategoryproductsToProductcategories Select Builder.
func (c jCategoryproductsToProductcategoriesQuerier) Select(what ...string) *jCategoryproductsToProductcategoriesSelectBuilder {
	m := c.Model()
	if len(what) == 0 {
		what = m.Fields("*")
	}
	dialect := runtime.GetDialect()
	from := dialect.QuoteIdent(m.Table())
	if m.Alias() != "" && m.Alias() != m.Table() {
		from = fmt.Sprintf("%v as %v", from, dialect.QuoteIdent(m.Alias()))
	}
	return &jCategoryproductsToProductcategoriesSelectBuilder{
		as: c.as,
		SelectBuilder: &builder.SelectBuilder{
			SelectBuilder: c.db.Select(what...).From(from),
		},
	}
}

//Count returns a CategoryproductsToProductcategories Select Builder to count given expressions.
func (c jCategoryproductsToProductcategoriesQuerier) Count(what ...string) *jCategoryproductsToProductcategoriesSelectBuilder {
	if len(what) == 0 {
		what = append(what, "*")
	}
	return c.Select("COUNT(" + strings.Join(what, ",") + ")")
}

// Insert a new CategoryproductsToProductcategories, if it has autoincrement primary key, the value will be set.
// It stops on first error.
func (c jCategoryproductsToProductcategoriesQuerier) Insert(items ...*CategoryproductsToProductcategories) (sql.Result, error) {
	var res sql.Result
	var err error
	for _, data := range items {
		res, err = c.db.InsertInto(JCategoryproductsToProductcategoriesModel.Table()).Columns(

			`product_id`,

			`category_id`,
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
func (c jCategoryproductsToProductcategoriesQuerier) InsertBulk(items ...*CategoryproductsToProductcategories) error {
	panic("todo")
}

//Delete returns a delete builder
func (c jCategoryproductsToProductcategoriesQuerier) Delete() *jCategoryproductsToProductcategoriesDeleteBuilder {
	return &jCategoryproductsToProductcategoriesDeleteBuilder{
		&builder.DeleteBuilder{
			DeleteBuilder: c.db.DeleteFrom(JCategoryproductsToProductcategoriesModel.Table()),
		},
	}
}

//DeleteByPk deletes one CategoryproductsToProductcategories by its PKs
func (c jCategoryproductsToProductcategoriesQuerier) DeleteByPk(ProductID int64, CategoryID int64) error {
	_, err := c.Delete().Where(

		JCategoryproductsToProductcategoriesModel.ProductID.Eq(ProductID),

		JCategoryproductsToProductcategoriesModel.CategoryID.Eq(CategoryID),
	).Exec()
	return err
}

// DeleteAll given CategoryproductsToProductcategories
func (c jCategoryproductsToProductcategoriesQuerier) DeleteAll(items ...*CategoryproductsToProductcategories) (sql.Result, error) {
	q := c.Delete()
	for _, item := range items {
		q = q.Where(
			dbr.Or(
				dbr.And(

					JCategoryproductsToProductcategoriesModel.ProductID.Eq(item.ProductID),

					JCategoryproductsToProductcategoriesModel.CategoryID.Eq(item.CategoryID),
				),
			),
		)
	}
	return q.Exec()
}

//Find one CategoryproductsToProductcategories using its PKs
func (c jCategoryproductsToProductcategoriesQuerier) Find(ProductID int64, CategoryID int64) (*CategoryproductsToProductcategories, error) {
	return c.Select().Where(

		JCategoryproductsToProductcategoriesModel.ProductID.Eq(ProductID),

		JCategoryproductsToProductcategoriesModel.CategoryID.Eq(CategoryID),
	).Read()
}
