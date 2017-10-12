---
License: MIT
LicenseFile: LICENSE
LicenseColor: yellow
---
# {{.Name}}

{{template "badge/travis" .}} {{template "badge/goreport" .}} {{template "badge/godoc" .}} {{template "license/shields" .}}

A golang database generator on top of [dbr](https://github.com/gocraft/dbr)

Compatibility
- ‎✔ SQLite3
- ‎✔ MySQL
- ‎‎✔ PostgreSQL

Features
- ‎✔ Table create / drop
- ‎✔ View create / drop
- ‎✔ CRUD operations
- ‎‎✔ auto increment support
- ‎‎✔ Always UTC date
- ‎‎✔ text pk
- ‎‎✔/‎- composite pk
- ‎✔ hasOne relation helper
- ‎✔ hasMany2One relation helper
- ‎‎✔/‎- hasMany2Many relation helper

__"✔/-"__ are items in progress, check the [CI](https://travis-ci.org/mh-cbon/jedi)

`jedi` is a go generator, read more about [go generate](https://blog.golang.org/generate)

# {{toc 5}}

# Install

{{template "go/install" .}}

# Implementing a model

## Declaration

```go
package whatever

// You should add a `go:generate jedi` comment
//go:generate jedi

// Todo is the type to record in database.
// You need to add a jedi: [table name] annotation to enable jedi support on this type.
//jedi:
type Todo struct {
	// jedi annotations are either
	// - @name=string for string values
	// - @name for thruthy values
	// - columnName to change the column name
	ID         int64       `jedi:"@pk"`
	Task       string      `jedi:"description"` // set a different column name
}
```

Run `go generate [package]` to generate the helpers.

The generated go code is written into files such as `<original file name>_jedi.go`.

# Jedi setup

Every `jedi` types is automatically registered at runtime.

Call `jedi.Setup(conn, forceSchemaReset)` to setup jedi driver and schema at runtime.

When `forceSchemaReset=true` the schema is dropped then created.

```go
package main

import (
	_ "github.com/mattn/go-sqlite3"
	"github.com/gocraft/dbr"
	jedi "github.com/mh-cbon/jedi/runtime"
)

func main () {
	dsn := "schema.db"
	conn, err := dbr.Open("sqlite3", dsn, nil)
	if err != nil {
		t.Fatalf("Connection setup failed: %v", err)
	}
	defer conn.Close()
	defer os.Remove(dsn)

	forceSchemaReset := true
	if err := jedi.Setup(conn, forceSchemaReset); err != nil {
		panic(err)
	}
	//...
}
```

## Schema

`jedi` creates for you a `jedi.Setuper` type for each `jedi` types.

```go
type jTodoSetup struct {
	Name       string
	CreateStmt string
	DropStmt   string
}

//Create applies the create table command to te underlying connection.
func (c jProductSetup) Create(db *dbr.Connection) {}

//Drop applies the drop table command to te underlying connection.
func (c jProductSetup) Drop(db *dbr.Connection) {}

func JTodoSetup() runtime.Setuper {}
```

To manually setup a type

```go
func main() {
	//...
	if err := JTodoSetup().Create(conn); err != nil {
		panic(err)
	}
}
```

# Jedi model

`jedi` types are translated into `jedi` models.

`jedi` models are runtime translations of the declared `jedi` types.

Models are struct of properties.

```go
// jTodoModel provides helper to work with Todo data provider
type jTodoModel struct {
	as string
	ID builder.ValuePropertyMeta
	Task builder.ValuePropertyMeta
}

var JTodoModel = jProductModel{
	ID: builder.NewValueMeta(
		`id`, `INTEGER`,
		`ID`, `int64`,
		true, true,
	),
	Task: builder.NewValueMeta(
		`task`, `TEXT`,
		`Task`, `string`,
		false, false,
	),
}
```

Models are useful to create condition without using raw text identifiers.

```go
JTodoModel.ID.Eq(1)
JTodoModel.ID.In(1, 2)
JTodoModel.ID.Gte(1)
// more in the documentation

JTodoModel.Task.Like("r%")
JTodoModel.Task.In("t", "r")
// more in the documentation
```

You can also get metadata
```go
JTodoModel.ID.IsPk()
JTodoModel.ID.IsAI()
// more in the documentation
```

## Tags

#### pk

`jedi:"@pk"` tag defines a property as being part of the primary key.

```go
type CompositePk struct {
	P           string `jedi:"@pk"`
	K           string `jedi:"@pk"`
	//...
}
```

#### has_one

`jedi:"@has_one"` tag defines a property as having `One <go type`.

```go
type Product struct {
	ID         int64       `jedi:"@pk"`
	brand      *Brand      `jedi:"@has_one=Brand.products"`
	BrandID    *int64
	//...
}
type Brand struct {
	ID        int64      `jedi:"@pk"`
	products  []*Product `jedi:"@has_many=Product.brand"`
	//...
}
```

#### has_many

`jedi:"@has_many"` tag defines a property as having `Many <go type`.

```go
type Product struct {
	ID       int64      `jedi:"@pk"`
	categories []*Category `jedi:"@has_many=Category.products"`
	//...
}
type Category struct {
	ID       int64      `jedi:"@pk"`
	products []*Product `jedi:"@has_many=Product.categories"`
	//...
}
```

#### on

`jedi:"@on"` tag defines the middle type being used for a `many2many` relation.

```go
type Product struct {
	ID       int64      		`jedi:"@pk"`
	categories []*Category 	`jedi:"@has_many=Category.products, @on=CatToProd"`
	//...
}
type Category struct {
	ID       int64      `jedi:"@pk"`
	products []*Product `jedi:"@has_many=Product.categories"`
	//...
}
type CatToProd struct {
	ProductsID       		int64      `jedi:"@pk"`
	CategoriesID       	int64      `jedi:"@pk"`
	//...
}
```

#### utc

`jedi:"@utc=false"` tag defines a `time.Time` property that must not be automatically turned into UTC before `Insert`/`Update`.

```go
type DateType struct {
	//...
	NotUTC      *time.Time `jedi:"@utc=false"`
}
```

#### last_updated

`jedi:"@last_updated"` tag defines the `time.Time` property to automatically being set when the struct is inserted, and added as a condition for an update query.

```go
type DateType struct {
	//...
	LastUpdated *time.Time `jedi:"@last_updated"`
}
```

# Jedi CRUD

`jedi` provides `CRUD` and more via a specialized querier type.

```go
type jTodoQuerier struct {
	db dbr.SessionRunner
	as string
}

// JTodo provides a todo querier
func JTodo(db dbr.SessionRunner) jTodoQuerier {
	return jTodoQuerier{
		db: db,
	}
}
```

To use a querier you need to create a `dbr.Session`

```go
package main

import (
	_ "github.com/mattn/go-sqlite3"
	jedi "github.com/mh-cbon/jedi/runtime"
)

func main () {
	// ...
	sess := conn.NewSession(nil)
	defer sess.Close()
	// ...
}
```

### Find

Every querier about types having primary keys implements a `Find(...pk) (type, error)` method.

It reads one instance by its pk.

```go
func main () {
	// ...
	todo, err := JTodo(sess).Find(1)
	if err != nil {
		panic(err)
	}
	log.Println(todo)
}
```

### Insert

To insert data in the database, the type declared must have primary keys.

The `Insert(obj type) (sql.Result, error)` method attempts to write given object into the database.

If the object has declared an `AUTO INCREMENT` field, the property is updated.

An `integer primary key` field is `AUTO INCREMENT`.

```go
func main () {
	// ...
	t := &Todo{}
	res, err := JTodo(sess).Insert(t)
	if err != nil {
		panic(err)
	}
	log.Println(t.ID)
	log.Println(res.LastInsertedID())
	log.Println(res.RowsAffected())
}
```

### Update

To update data in the database, the type declared must have primary keys.

The `Update(obj type) (sql.Result, error)` method attempts to write existing object into the database.

```go
func main () {
	// ...
	res, err := JTodo(sess).Update(&Todo{ID:1})
	if err != nil {
		panic(err)
	}
	log.Println(res.RowsAffected())
}
```

### DeleteByPk

To delete data by pk in the database, the type declared must have primary keys.

The `DeleteByPk(pk type) error` method attempts to delete the row matching given primary keys from the database.

```go
func main () {
	// ...
	res, err := JTodo(sess).DeleteByPk(1)
	if err != nil {
		panic(err)
	}
	log.Println(res.RowsAffected())
}
```

### DeleteAll

To delete many instances within the database, the type declared must have primary keys.

The `DeleteAll(items ...type) error` method attempts to delete the rows matching given instances from the database.

```go
func main () {
	// ...
	t := &Todo{ID:1}
	res, err := JTodo(sess).DeleteAll(t,t,t)
	if err != nil {
		panic(err)
	}
	log.Println(res.RowsAffected())
}
```

## query builder

`jedi` provides query building capabilities on top of `dbr`.

The query builder should help to build queries programmatically and
improve application maintenance.

### Select

The `Select(what ...string) <type>SelectBuilder` returns a select query builder of given columns.

If there is no given columns it defaults to `alias.*`.

All `Select()` call should end with an execution call to consume the query.

A short list of available actions is
- `Read() (*type, error)`
- `ReadAll() ([]*type, error)`
- `ReadInt() (int, error)`
- `ReadInt64() (int64, error)`

See also dbr documentation for `Load`/`LoadStructs` etc.

```go
func main () {
	// ...
	todo, err := JTodo(sess).
		Select("task.*"). // input sql values.
		Where(JTodoModel.Task.Like("%whatever%")). // set some conditions
		Read() // get all results found
	if err != nil {
		panic(err)
	}
	log.Println(todo)
}
```

See also `Limit`/`Offset` ect in dbr documentation.

### Where

The `Where(query interface{}, value ...interface{}) <type>SelectBuilder` is a shorthand for `Querier(sess).Select().Where()`

```go
func main () {
	// ...
	todo, err := JTodo(sess).Where(JTodoModel.Task.Like("%whatever%")).Read()
	if err != nil {
		panic(err)
	}
	log.Println(todo)
}
```

### Delete

The `Delete() <type>DeleteBuilder` is a query builder to remove data.

It s a shorthand for `DELETE FROM XXXX`

```go
func main () {
	// ...
	res, err := JTodo(sess).
		Delete().Where(JTodoModel.Task.Like("%whatever%")).Exec()
	if err != nil {
		panic(err)
	}
	log.Println(res.RowsAffected())
}
```

# Working with Basic types

You can work with those basic types, they might be pointer too,

- string
- int / int8 / int16 / int32 / int64
- uint / uint8 / uint16 / uint32 / uint64
- float32 / float64
- time.Time

# Working with Dates

`jedi` recognizes fields of type `time.Time` or `*time.Time`.

Unless its tags defines `jedi:"@utc=false"`, it will automatically be turned into UTC before `Insert` and `Update` queries.

# Working with text PK

When you define a string field as being part of the PK,
`jedi` will use a `VARCHAR(255)` sql type when the driver is `mysql`.

Error reference : `BLOB/TEXT column 'XXX' used in key specification without a key length`

Ass an addition `jedi` will add special checks in the `Insert`/`Update` procedure to trigger an error
if you pass in a string with a length greater than 255 when `mysql` is the driver being used.

This is to ensure consistency independently of the underlying driver.

## Fractionnal seconds

`time.Time` of `@last_updated` properties are always truncated to the microseond (6 digits).

__In general__ if you stumble upon a case where `refTime.Equal(y.JustInsertedWhateverTime)`
is unexpectedly `false`, try to Round the time values such as `y.Truncate(time.Microsecond)`.

# Working with Relations

`jedi` supports `@has_one` and `@has_many` property tags to implement `oneToMany`, `manyToOne` and `manyToMany` relationships.

Those attributes (`@has_one`/`@has_many`) takes a value, the foreign reverse property.

Whe you declare a relationship, you must do it on a `private/unexported` property.

The type might be a `pointer`, or a `value`.

```go
//Product is a sku representation.
//jedi:
type Product struct {
	ID         int64       `jedi:"@pk"`
	brand      *Brand      `jedi:"@has_one=Brand.products"`
	categories []*Category `jedi:"@has_many=Category.products"`
}
```

## Has One

The `@has_one` tag attribute defines a `one to one` relationship.

The type declaring the `@has_one` attribute must also declare the
imported primary keys respecting the convention `<local property name | ucfirst><foreign primary key name>`

```go
//Product is a sku representation.
//jedi:
type Product struct {
	ID         int64       `jedi:"@pk"`
	brand2      *Brand      `jedi:"@has_one=Brand.products"`
	Brand2ID    *int64      // the imported primary key of Brand.ID on Product.brand2
}

//Brand is a product brand representation.
//jedi:
type Brand struct {
	ID        int64      `jedi:"@pk"`
	products  []*Product `jedi:"@has_many=Product.brand2"`
	Name      string
}
```

#### Set / Unset / Read

Because `Product` has a property `has_one` of type `Brand`, `jedi` automatically adds
approperiate methods to `Set<PropertyName>` / `Unset<Property name>`, such as:

```go
func (p *Product) SetBrand(o *Brand) *Product {}
func (p *Product) UnsetBrand() *Product {}
```

It also creates a method to `read` the related object from the database:

```go
func (p *Product) Brand(db dbr.SessionRunner) (*Brand, error)  {}
```

#### Join

For every `has_one` properties the related querier gets
new `Join<PropertyName>` / `LeftJoin<Property name>` methods:

```go
func (c *jProductSelectBuilder) JoinBrand( AsBrand string) *jProductSelectBuilder { }
func (c *jProductSelectBuilder) LeftJoinBrand( AsBrand string) *jProductSelectBuilder { }
```

## Has Many 2 One

The `@has_many` tag attribute defines a `one to many` relationship.

If the foreign reverse property defines an `@has_one` tag,
then this property become a specific `has_many_to_one` relationship.

Unlike `@has_one` attribute, it does not require additional properties.

```go
type Brand struct {
	ID        int64      `jedi:"@pk"`
	products  []*Product `jedi:"@has_many=Product.brand"`
	Name      string
}

type Product struct {
	ID         int64       `jedi:"@pk"`
	brand      *Brand      `jedi:"@has_one=Brand.products"`
	BrandID    *int64      // the imported primary key of Brand.ID on Product.brand2
}
```

#### Read

For every `@has_many` properties `jedi` adds method to select related objects

```go
func (b *Brand) Products(db dbr.SessionRunner, AsBrand, AsProducts string) *jProductSelectBuilder {}
```

#### Join

The querier also gets
new `Join<PropertyName>` / `LeftJoin<Property name>` methods:

```go
func (c *jBrandSelectBuilder) JoinProducts(AsProducts string) *jBrandSelectBuilder {}
func (c *jBrandSelectBuilder) LeftJoinProducts(AsProducts string) *jBrandSelectBuilder {}

func (c *jProductSelectBuilder) JoinBrand(AsBrand string) *jProductSelectBuilder {}
func (c *jProductSelectBuilder) LeftJoinBrand(AsBrand string) *jProductSelectBuilder {}
```

## Has Many 2 Many

If the foreign reverse property of an `@has_many` also defines an `@has_many` tag,
then this property become a specific `has_many_to_many` relationship.

Such relationship are handled by a middle table.

The middle table is automatically deduced and registered at runtime.

```go
//Product is a sku representation.
//jedi:
type Product struct {
	ID         int64       `jedi:"@pk"`
	SKU        string      //todo: see if can be pk (string not int)
	categories []*Category `jedi:"@has_many=Category.products"`
}

//Category is a product category representation.
//jedi:
type Category struct {
	ID       int64      `jedi:"@pk"`
	products []*Product `jedi:"@has_many=Product.categories"`
	Name     string
}
```

Alternatively, you can define a specific middle type by defining an `@on=<type>` extra tag attribute on one of the side.

That table must declare appropriate properties refering to the primary keys of each table involved.

It can declare additionnal properties.

```go
//Product is a sku representation.
//jedi:
type Product struct {
	ID         int64       `jedi:"@pk"`
	SKU        string      //todo: see if can be pk (string not int)
	categories []*Category `jedi:"@has_many=Category.products, @on=CategoryToProduct"`
}

//CategoryToProduct is a product to category relationship.
//jedi:
type CategoryToProduct struct {
	CategoriesID int64
	ProductsID int64
	ProductOrder int
}

//Category is a product category representation.
//jedi:
type Category struct {
	ID       int64      `jedi:"@pk"`
	products []*Product `jedi:"@has_many=Product.categories"`
	Name     string
}
```

#### LinkWith / UnlinkWith

For every `many2many` relations, `jedi` will create appropriatem methods on the jedified type to
`LinkWith<PropertyName>` / `UnlinkWith<PropertyName>` related object.

```go
func (p *Product) LinkWithCategories(db dbr.SessionRunner, items ...*Category) (sql.Result, error) {}
func (p *Product) UnlinkWithCategories(db dbr.SessionRunner, items ...*Category) (sql.Result, error) {}

func (c *Category) LinkWithProducts(db dbr.SessionRunner, items ...*Product) (sql.Result, error) {}
func (c *Category) UnlinkWithProducts(db dbr.SessionRunner, items ...*Product) (sql.Result, error) {}
```

#### Select

For every `many2many` relations, `jedi` will create appropriatem methods on the jedified type to
get related related objects.

```go
func (g *Product) Categories(
	db dbr.SessionRunner,
	AsCategory,
	AsCategoryproductsToProductcategories,
	AsProduct string,
) *jCategorySelectBuilder {}

func (c *Category) Products(
	db dbr.SessionRunner,
	AsProduct,
	AsCategoryproductsToProductcategories,
	AsCategory string,
	) *jProductSelectBuilder {}
```

#### Join

The querier also gets
new `Join<PropertyName>` / `LeftJoin<Property name>` methods:

```go
func (p *jProductSelectBuilder) JoinCategories(
	AsCategoryproductsToProductcategories, AsCategory string,
) *jProductSelectBuilder {}

func (p *jProductSelectBuilder) LeftJoinCategories(
	AsCategoryproductsToProductcategories, AsCategory string,
) *jProductSelectBuilder {}



func (c *jCategorySelectBuilder) JoinProducts(
	AsCategoryproductsToProductcategories, AsProduct string,
) *jCategorySelectBuilder {}

func (c *jCategorySelectBuilder) LeftJoinProducts(
	AsCategoryproductsToProductcategories, AsProduct string,
) *jCategorySelectBuilder {}
```

# Working with views

Tow ork with views, or define the `CREATE` and `DROP` queries,
it is possible to define multiples attributes on the `jedi` types.

### view-select
The `view-select` let you define the `SELECT` query to generate the data related to the type,

```go
//SampleView is view of samples.
//jedi:
//view-select:
//	SELECT *
//	FROM sample
//	WHERE id > 1
//
// regular commets can restart here.
type SampleView struct {
	ID          int64 `jedi:"@pk"` //you can configure the ok on the view, it adds some methods.
	Name        string
	Description string
}
```

The `SELECT` query is automatically wrapped with `CREATE VIEW IF NOT EXISTS...`

### view-create

The `view-create` let you define the `CREATE` query of the view,

```go
//jedi:
//view-create:
//	CREATE VIEW xyz...
//
// regular commets can restart here.
type W struct {
	ID          int64 `jedi:"@pk"` //you can configure the ok on the view, it adds some methods.
	Name        string
	Description string
}
```

### view-drop

### view-create

The `view-drop` let you define the `DROP` query of the view,

```go
//jedi:
//view-drop:
//	DROP...
//
// regular commets can restart here.
type W struct {
	ID          int64 `jedi:"@pk"` //you can configure the ok on the view, it adds some methods.
	Name        string
	Description string
}
```

### table-create

The `table-create` let you define the `CREATE` query of the table,

```go
//jedi:
//table-create:
//	CREATE TABLE...
//
// regular commets can restart here.
type W struct {
	ID          int64 `jedi:"@pk"` //you can configure the ok on the view, it adds some methods.
	Name        string
	Description string
}
```

### table-drop

The `table-drop` let you define the `DROP` query of the table,

```go
//jedi:
//table-drop:
//	DROP TABLE...
//
// regular commets can restart here.
type W struct {
	ID          int64 `jedi:"@pk"` //you can configure the ok on the view, it adds some methods.
	Name        string
	Description string
}
```

# cli

#### $ {{exec "jedi" "-help" | color "sh"}}

# credits

inspiration from [reform](https://github.com/go-reform/reform), [dbr](https://github.com/gocraft/dbr).
