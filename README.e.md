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
- ‎‎✔/‎- text pk
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
func (c jProductSetup) Create(db *dbr.Connection) ...

//Drop applies the drop table command to te underlying connection.
func (c jProductSetup) Drop(db *dbr.Connection) ...

func JTodoSetup() runtime.Setuper {
}
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

Models are useful to create condition without using wihtout raw text.

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

# Jedi crud

`jedi` provides CRUD and more via a specialized querier type.

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

The `Update(obj type) (sql.Result, error)` method attempts to write given object into the database.

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
		Select("task"). // input sql values.
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
	products  []*Product `jedi:"@has_many=Product.brand"`
	Name      string
}
```

#### Set
#### Unset
#### Read
#### Join

## Has Many 2 One

The `@has_many` tag attribute defines a `one to many` relationship.

If the foreign reverse property defines an `@has_one` tag,
then this property become a specific `has_many_to_one` relationship.

Unlike `@has_one` attribute, it does not require additional properties.

```go
//Brand is a product brand representation.
//jedi:
type Brand struct {
	ID        int64      `jedi:"@pk"`
	products  []*Product `jedi:"@has_many=Product.brand"`
	Name      string
}

//Product is a sku representation.
//jedi:
type Product struct {
	ID         int64       `jedi:"@pk"`
	brand2      *Brand      `jedi:"@has_one=Brand.products"`
	Brand2ID    *int64      // the imported primary key of Brand.ID on Product.brand2
}
```

#### Set
#### Unset
#### Read
#### Join

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

#### Link
#### Unlink
#### Read
#### Join

# cli

#### $ {{exec "jedi" "-help" | color "sh"}}

# credits

inspiration from [reform](https://github.com/go-reform/reform), [dbr](https://github.com/gocraft/dbr).
