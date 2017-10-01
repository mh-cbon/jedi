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

func (c jProductSetup) Create(db *dbr.Connection) error {
	_, err := db.Exec(c.CreateStmt)
	return err
}
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
brand Brand,
brand_id INTEGER

)`
	} else if driver == drivers.Mysql {
		create = `CREATE TABLE IF NOT EXISTS product (
id INTEGER NOT NULL AUTO_INCREMENT,
sku TEXT,
brand Brand,
brand_id INTEGER,
PRIMARY KEY (id) 

)`
	} else if driver == drivers.Pgsql {
		create = `CREATE TABLE IF NOT EXISTS product (
id INTEGER,
sku TEXT,
brand Brand,
brand_id INTEGER

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

var JProductModel = struct {
	Eq func(...*Product) dbr.Builder
	In func(...*Product) dbr.Builder

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

	SKU struct {
		Name string
		Eq   func(interface{}) dbr.Builder
		In   func(...interface{}) dbr.Builder
		Like func(interface{}) dbr.Builder
	}

	BrandID struct {
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
}{

	Eq: func(s ...*Product) dbr.Builder {
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
	In: func(s ...*Product) dbr.Builder {
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

	SKU: struct {
		Name string
		Eq   func(interface{}) dbr.Builder
		In   func(...interface{}) dbr.Builder
		Like func(interface{}) dbr.Builder
	}{
		Name: `sku`,
		Eq: func(v interface{}) dbr.Builder {
			return dbr.Eq(`sku`, v)
		},
		In: func(v ...interface{}) dbr.Builder {
			return dbr.Eq(`sku`, v)
		},
		Like: func(v interface{}) dbr.Builder {
			return builder.Like(`sku`, v)
		},
	},

	BrandID: struct {
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
		Name: `brand_id`,
		IsPk: false,
		IsAI: false,
		Eq: func(v interface{}) dbr.Builder {
			return dbr.Eq(`brand_id`, v)
		},
		In: func(v ...interface{}) dbr.Builder {
			return dbr.Eq(`brand_id`, v)
		},
		Gt: func(v interface{}) dbr.Builder {
			return dbr.Gt(`brand_id`, v)
		},
		Gte: func(v interface{}) dbr.Builder {
			return dbr.Gte(`brand_id`, v)
		},
		Lt: func(v interface{}) dbr.Builder {
			return dbr.Lt(`brand_id`, v)
		},
		Lte: func(v interface{}) dbr.Builder {
			return dbr.Lte(`brand_id`, v)
		},
	},
}

// JProduct provides a basic querier
func JProduct(db dbr.SessionRunner) jProductQuerier {
	return jProductQuerier{
		name: `product`,
		db:   db,
	}
}

type jProductSelectBuilder struct {
	*builder.SelectBuilder
}

func (c *jProductSelectBuilder) Read() (*Product, error) {
	var one Product
	err := c.LoadStruct(&one)
	return &one, err
}
func (c *jProductSelectBuilder) ReadAll() ([]*Product, error) {
	var all []*Product
	_, err := c.LoadStructs(&all)
	return all, err
}
func (c *jProductSelectBuilder) Where(query interface{}, value ...interface{}) *jProductSelectBuilder {
	c.SelectBuilder.Where(query, value...)
	return c
}

type jProductQuerier struct {
	name string
	db   dbr.SessionRunner
}

func (c jProductQuerier) Select(what ...string) *jProductSelectBuilder {
	if len(what) == 0 {
		what = append(what, "*")
	}
	return &jProductSelectBuilder{
		&builder.SelectBuilder{
			SelectBuilder: c.db.Select(what...).From(c.name),
		},
	}
}
func (c jProductQuerier) Count(what ...string) *jProductSelectBuilder {
	if len(what) == 0 {
		what = append(what, "*")
	}
	return c.Select("COUNT(" + strings.Join(what, ",") + ")")
}

func (c jProductQuerier) Insert(data *Product) (sql.Result, error) {
	res, err := c.db.InsertInto(c.name).Columns(

		`sku`,

		`brand`,

		`brand_id`,
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

func (c jProductQuerier) Update(data *Product) (sql.Result, error) {
	res, err := c.db.Update(c.name).
		Set(`sku`, data.SKU).
		Set(`brand`, data.brand).
		Set(`brand_id`, data.BrandID).
		Where("id = ?", data.ID).
		Exec()
	return res, err
}

func (c jProductQuerier) Delete() *builder.DeleteBuilder {
	return &builder.DeleteBuilder{
		DeleteBuilder: c.db.DeleteFrom(c.name),
	}
}

func (c jProductQuerier) DeleteByPk(id int64) error {
	_, err := c.Delete().Where(

		JProductModel.ID.Eq(id),
	).Exec()
	return err
}
func (c jProductQuerier) DeleteAll(items ...*Product) error {
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
	_, err := q.Exec()
	return err
}
func (c jProductQuerier) Find(id int64) (*Product, error) {
	return c.Select().Where(

		JProductModel.ID.Eq(id),
	).Read()
}

func (g Product) Brand(db dbr.SessionRunner) (*Brand, error) {
	q := JBrand(db).Select()

	q = q.Where(

		JBrandModel.ID.Eq(g.BrandID),
	)
	return q.Read()
}

type jCategorySetup struct {
	Name       string
	CreateStmt string
	DropStmt   string
}

func (c jCategorySetup) Create(db *dbr.Connection) error {
	_, err := db.Exec(c.CreateStmt)
	return err
}
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

var JCategoryModel = struct {
	Eq func(...*Category) dbr.Builder
	In func(...*Category) dbr.Builder

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
}{

	Eq: func(s ...*Category) dbr.Builder {
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
	In: func(s ...*Category) dbr.Builder {
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
}

// JCategory provides a basic querier
func JCategory(db dbr.SessionRunner) jCategoryQuerier {
	return jCategoryQuerier{
		name: `category`,
		db:   db,
	}
}

type jCategorySelectBuilder struct {
	*builder.SelectBuilder
}

func (c *jCategorySelectBuilder) Read() (*Category, error) {
	var one Category
	err := c.LoadStruct(&one)
	return &one, err
}
func (c *jCategorySelectBuilder) ReadAll() ([]*Category, error) {
	var all []*Category
	_, err := c.LoadStructs(&all)
	return all, err
}
func (c *jCategorySelectBuilder) Where(query interface{}, value ...interface{}) *jCategorySelectBuilder {
	c.SelectBuilder.Where(query, value...)
	return c
}

type jCategoryQuerier struct {
	name string
	db   dbr.SessionRunner
}

func (c jCategoryQuerier) Select(what ...string) *jCategorySelectBuilder {
	if len(what) == 0 {
		what = append(what, "*")
	}
	return &jCategorySelectBuilder{
		&builder.SelectBuilder{
			SelectBuilder: c.db.Select(what...).From(c.name),
		},
	}
}
func (c jCategoryQuerier) Count(what ...string) *jCategorySelectBuilder {
	if len(what) == 0 {
		what = append(what, "*")
	}
	return c.Select("COUNT(" + strings.Join(what, ",") + ")")
}

func (c jCategoryQuerier) Insert(data *Category) (sql.Result, error) {
	res, err := c.db.InsertInto(c.name).Columns(

		`name`,
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

func (c jCategoryQuerier) Update(data *Category) (sql.Result, error) {
	res, err := c.db.Update(c.name).
		Set(`name`, data.Name).
		Where("id = ?", data.ID).
		Exec()
	return res, err
}

func (c jCategoryQuerier) Delete() *builder.DeleteBuilder {
	return &builder.DeleteBuilder{
		DeleteBuilder: c.db.DeleteFrom(c.name),
	}
}

func (c jCategoryQuerier) DeleteByPk(id int64) error {
	_, err := c.Delete().Where(

		JCategoryModel.ID.Eq(id),
	).Exec()
	return err
}
func (c jCategoryQuerier) DeleteAll(items ...*Category) error {
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
	_, err := q.Exec()
	return err
}
func (c jCategoryQuerier) Find(id int64) (*Category, error) {
	return c.Select().Where(

		JCategoryModel.ID.Eq(id),
	).Read()
}

type jBrandSetup struct {
	Name       string
	CreateStmt string
	DropStmt   string
}

func (c jBrandSetup) Create(db *dbr.Connection) error {
	_, err := db.Exec(c.CreateStmt)
	return err
}
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

var JBrandModel = struct {
	Eq func(...*Brand) dbr.Builder
	In func(...*Brand) dbr.Builder

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
}{

	Eq: func(s ...*Brand) dbr.Builder {
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
	In: func(s ...*Brand) dbr.Builder {
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
}

// JBrand provides a basic querier
func JBrand(db dbr.SessionRunner) jBrandQuerier {
	return jBrandQuerier{
		name: `brand`,
		db:   db,
	}
}

type jBrandSelectBuilder struct {
	*builder.SelectBuilder
}

func (c *jBrandSelectBuilder) Read() (*Brand, error) {
	var one Brand
	err := c.LoadStruct(&one)
	return &one, err
}
func (c *jBrandSelectBuilder) ReadAll() ([]*Brand, error) {
	var all []*Brand
	_, err := c.LoadStructs(&all)
	return all, err
}
func (c *jBrandSelectBuilder) Where(query interface{}, value ...interface{}) *jBrandSelectBuilder {
	c.SelectBuilder.Where(query, value...)
	return c
}

type jBrandQuerier struct {
	name string
	db   dbr.SessionRunner
}

func (c jBrandQuerier) Select(what ...string) *jBrandSelectBuilder {
	if len(what) == 0 {
		what = append(what, "*")
	}
	return &jBrandSelectBuilder{
		&builder.SelectBuilder{
			SelectBuilder: c.db.Select(what...).From(c.name),
		},
	}
}
func (c jBrandQuerier) Count(what ...string) *jBrandSelectBuilder {
	if len(what) == 0 {
		what = append(what, "*")
	}
	return c.Select("COUNT(" + strings.Join(what, ",") + ")")
}

func (c jBrandQuerier) Insert(data *Brand) (sql.Result, error) {
	res, err := c.db.InsertInto(c.name).Columns(

		`name`,
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

func (c jBrandQuerier) Update(data *Brand) (sql.Result, error) {
	res, err := c.db.Update(c.name).
		Set(`name`, data.Name).
		Where("id = ?", data.ID).
		Exec()
	return res, err
}

func (c jBrandQuerier) Delete() *builder.DeleteBuilder {
	return &builder.DeleteBuilder{
		DeleteBuilder: c.db.DeleteFrom(c.name),
	}
}

func (c jBrandQuerier) DeleteByPk(id int64) error {
	_, err := c.Delete().Where(

		JBrandModel.ID.Eq(id),
	).Exec()
	return err
}
func (c jBrandQuerier) DeleteAll(items ...*Brand) error {
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
	_, err := q.Exec()
	return err
}
func (c jBrandQuerier) Find(id int64) (*Brand, error) {
	return c.Select().Where(

		JBrandModel.ID.Eq(id),
	).Read()
}

type jCategoryproductsToProductcategoriesSetup struct {
	Name       string
	CreateStmt string
	DropStmt   string
}

func (c jCategoryproductsToProductcategoriesSetup) Create(db *dbr.Connection) error {
	_, err := db.Exec(c.CreateStmt)
	return err
}
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
		create = `CREATE TABLE IF NOT EXISTS category_products_product_categories (
product_id INTEGER,
category_id INTEGER,
PRIMARY KEY (product_id,category_id) 

)`
	} else if driver == drivers.Mysql {
		create = `CREATE TABLE IF NOT EXISTS category_products_product_categories (
product_id INTEGER NOT NULL,
category_id INTEGER NOT NULL,
PRIMARY KEY (product_id,category_id) 

)`
	} else if driver == drivers.Pgsql {
		create = `CREATE TABLE IF NOT EXISTS category_products_product_categories (
product_id INTEGER,
category_id INTEGER

)`
	}

	if driver == drivers.Sqlite {
		drop = `DROP TABLE IF EXISTS category_products_product_categories`
	} else if driver == drivers.Mysql {
		drop = `DROP TABLE IF EXISTS category_products_product_categories`
	} else if driver == drivers.Pgsql {
		drop = `DROP TABLE IF EXISTS category_products_product_categories`
	}

	return jCategoryproductsToProductcategoriesSetup{
		Name:       `category_products_product_categories`,
		CreateStmt: create,
		DropStmt:   drop,
	}
}

var JCategoryproductsToProductcategoriesModel = struct {
	Eq func(...*CategoryproductsToProductcategories) dbr.Builder
	In func(...*CategoryproductsToProductcategories) dbr.Builder

	ProductID struct {
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

	CategoryID struct {
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
}{

	Eq: func(s ...*CategoryproductsToProductcategories) dbr.Builder {
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
	},
	In: func(s ...*CategoryproductsToProductcategories) dbr.Builder {
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
	},

	ProductID: struct {
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
		Name: `product_id`,
		IsPk: true,
		IsAI: true,
		Eq: func(v interface{}) dbr.Builder {
			return dbr.Eq(`product_id`, v)
		},
		In: func(v ...interface{}) dbr.Builder {
			return dbr.Eq(`product_id`, v)
		},
		Gt: func(v interface{}) dbr.Builder {
			return dbr.Gt(`product_id`, v)
		},
		Gte: func(v interface{}) dbr.Builder {
			return dbr.Gte(`product_id`, v)
		},
		Lt: func(v interface{}) dbr.Builder {
			return dbr.Lt(`product_id`, v)
		},
		Lte: func(v interface{}) dbr.Builder {
			return dbr.Lte(`product_id`, v)
		},
	},

	CategoryID: struct {
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
		Name: `category_id`,
		IsPk: true,
		IsAI: true,
		Eq: func(v interface{}) dbr.Builder {
			return dbr.Eq(`category_id`, v)
		},
		In: func(v ...interface{}) dbr.Builder {
			return dbr.Eq(`category_id`, v)
		},
		Gt: func(v interface{}) dbr.Builder {
			return dbr.Gt(`category_id`, v)
		},
		Gte: func(v interface{}) dbr.Builder {
			return dbr.Gte(`category_id`, v)
		},
		Lt: func(v interface{}) dbr.Builder {
			return dbr.Lt(`category_id`, v)
		},
		Lte: func(v interface{}) dbr.Builder {
			return dbr.Lte(`category_id`, v)
		},
	},
}

// CategoryproductsToProductcategories ...
type CategoryproductsToProductcategories struct {
	ProductID int64

	CategoryID int64
}

// JCategoryproductsToProductcategories provides a basic querier
func JCategoryproductsToProductcategories(db dbr.SessionRunner) jCategoryproductsToProductcategoriesQuerier {
	return jCategoryproductsToProductcategoriesQuerier{
		name: `category_products_product_categories`,
		db:   db,
	}
}

type jCategoryproductsToProductcategoriesSelectBuilder struct {
	*builder.SelectBuilder
}

func (c *jCategoryproductsToProductcategoriesSelectBuilder) Read() (*CategoryproductsToProductcategories, error) {
	var one CategoryproductsToProductcategories
	err := c.LoadStruct(&one)
	return &one, err
}
func (c *jCategoryproductsToProductcategoriesSelectBuilder) ReadAll() ([]*CategoryproductsToProductcategories, error) {
	var all []*CategoryproductsToProductcategories
	_, err := c.LoadStructs(&all)
	return all, err
}
func (c *jCategoryproductsToProductcategoriesSelectBuilder) Where(query interface{}, value ...interface{}) *jCategoryproductsToProductcategoriesSelectBuilder {
	c.SelectBuilder.Where(query, value...)
	return c
}

type jCategoryproductsToProductcategoriesQuerier struct {
	name string
	db   dbr.SessionRunner
}

func (c jCategoryproductsToProductcategoriesQuerier) Select(what ...string) *jCategoryproductsToProductcategoriesSelectBuilder {
	if len(what) == 0 {
		what = append(what, "*")
	}
	return &jCategoryproductsToProductcategoriesSelectBuilder{
		&builder.SelectBuilder{
			SelectBuilder: c.db.Select(what...).From(c.name),
		},
	}
}
func (c jCategoryproductsToProductcategoriesQuerier) Count(what ...string) *jCategoryproductsToProductcategoriesSelectBuilder {
	if len(what) == 0 {
		what = append(what, "*")
	}
	return c.Select("COUNT(" + strings.Join(what, ",") + ")")
}

func (c jCategoryproductsToProductcategoriesQuerier) Insert(data *CategoryproductsToProductcategories) (sql.Result, error) {
	res, err := c.db.InsertInto(c.name).Columns(

		`product_id`,

		`category_id`,
	).Record(data).Exec()

	return res, err
}

func (c jCategoryproductsToProductcategoriesQuerier) Delete() *builder.DeleteBuilder {
	return &builder.DeleteBuilder{
		DeleteBuilder: c.db.DeleteFrom(c.name),
	}
}

func (c jCategoryproductsToProductcategoriesQuerier) DeleteByPk(product_id int64, category_id int64) error {
	_, err := c.Delete().Where(

		JCategoryproductsToProductcategoriesModel.ProductID.Eq(product_id),

		JCategoryproductsToProductcategoriesModel.CategoryID.Eq(category_id),
	).Exec()
	return err
}
func (c jCategoryproductsToProductcategoriesQuerier) DeleteAll(items ...*CategoryproductsToProductcategories) error {
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
	_, err := q.Exec()
	return err
}
func (c jCategoryproductsToProductcategoriesQuerier) Find(product_id int64, category_id int64) (*CategoryproductsToProductcategories, error) {
	return c.Select().Where(

		JCategoryproductsToProductcategoriesModel.ProductID.Eq(product_id),

		JCategoryproductsToProductcategoriesModel.CategoryID.Eq(category_id),
	).Read()
}
