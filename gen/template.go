package gen

import (
	"text/template"
)

// templates exported for the generator
var (
	Registry    *template.Template
	Prolog      *template.Template
	Struct      *template.Template
	HelpersDecl *template.Template
	HelpersBody *template.Template
)

func init() {

	Registry = template.Must(template.New("registry").Funcs(funcs).Parse(`
// Generated with mh-cbon/jedi. Do not edit by hand.
package {{.PackageName}}

import (
	"github.com/mh-cbon/jedi/runtime"
)

//Jedi is a runtime registry to jedi types.
var Jedi []runtime.Setupable
`))

	Prolog = template.Must(template.New("prolog").Funcs(funcs).Parse(`
// Generated with mh-cbon/jedi. Do not edit by hand.
package {{.PackageName}}

import (
	"database/sql"
	"fmt"
	"strings"
	"time"

	"github.com/gocraft/dbr"
	dbrdialect "github.com/gocraft/dbr/dialect"
	"github.com/mh-cbon/jedi/runtime"
	"github.com/mh-cbon/jedi/drivers"
	"github.com/mh-cbon/jedi/builder"
)

var _ = time.Now
var _ = fmt.Sprintf
var _ = dbrdialect.PostgreSQL

func init(){
	{{range $i, $s := $.all}}
	Jedi = append(Jedi, J{{.Name}}Setup)
	{{end}}
}
`))

	Struct = template.Must(template.New("struct").Funcs(funcs).Parse(`


// J{{.current.Name}}Setup helps to create/drop the schema
func J{{.current.Name}}Setup() runtime.Setuper {
	driver := runtime.GetCurrentDriver()

	var create string
	var drop string
	{{if empty .current.SQLViewSelect .current.SQLViewCreate}}
		{{if notEmpty .current.SQLTableCreate}}
			create = {{.current.SQLTableCreate | quote}}
		{{else}}
			if driver==drivers.Sqlite {
				create = {{.current | createTable "sqlite3" | quote}}
			} else if driver==drivers.Mysql {
				create = {{.current | createTable "mysql" | quote}}
			} else if driver==drivers.Pgsql {
				create = {{.current | createTable "postgres" | quote}}
			}
		{{end}}
		{{if notEmpty .current.SQLTableDrop}}
			drop = {{.current.SQLTableDrop | quote}}
		{{else}}
			if driver==drivers.Sqlite {
				drop = {{.current | dropTable "sqlite3" | quote}}
			} else if driver==drivers.Mysql {
				drop = {{.current | dropTable "mysql" | quote}}
			} else if driver==drivers.Pgsql {
				drop = {{.current | dropTable "postgres" | quote}}
			}
		{{end}}
	{{else}}
		{{if notEmpty .current.SQLViewCreate}}
			create = {{.current.SQLViewCreate | quote}}
		{{else if notEmpty .current.SQLViewSelect}}
			if driver==drivers.Sqlite {
				create = {{.current | createView "sqlite3" | quote}}
			} else if driver==drivers.Mysql {
				create = {{.current | createView "mysql" | quote}}
			} else if driver==drivers.Pgsql {
				create = {{.current | createView "postgres" | quote}}
			}
		{{end}}
		{{if notEmpty .current.SQLViewDrop}}
			drop = {{.current.SQLViewDrop | quote}}
		{{else if notEmpty .current.SQLViewSelect}}
			if driver==drivers.Sqlite {
				drop = {{.current | dropView "sqlite3" | quote}}
			} else if driver==drivers.Mysql {
				drop = {{.current | dropView "mysql" | quote}}
			} else if driver==drivers.Pgsql {
				drop = {{.current | dropView "postgres" | quote}}
			}
		{{end}}
	{{end}}

	var indexes []string

	if driver==drivers.Sqlite {
		{{range $i, $index := .current.Indexes}}
		indexes = append(indexes, {{$.current | createIndex "sqlite3" $index | quote}})
		{{end}}
	} else if driver==drivers.Mysql {
		{{range $i, $index := .current.Indexes}}
		indexes = append(indexes, {{$.current | createIndex "mysql" $index | quote}})
		{{end}}
	} else if driver==drivers.Pgsql {
		{{range $i, $index := .current.Indexes}}
		indexes = append(indexes, {{$.current | createIndex "postgres" $index | quote}})
		{{end}}
	}

	return runtime.Table {
		Name: {{.current.SQLName | quote}},
		CreateStmt: create,
		DropStmt: drop,
		View: !{{empty .current.SQLViewSelect .current.SQLViewCreate}},
		Indexes: indexes,
	}
}

// j{{.current.Name}}Model provides helper to work with {{.current.Name}} data provider
type j{{.current.Name}}Model struct {
	as string
	{{range $i, $col := .current.Fields | isExported}}
		{{$col.Name}} builder.ValuePropertyMeta
	{{end}}
	{{range $i, $col := .current.Fields | notExported | isRel}}
		{{$col.Name | ucfirst}} builder.RelPropertyMeta
	{{end}}
}

{{if .current.Fields | hasPkField}}
// Eq provided items.
func (j j{{.current.Name}}Model) Eq(s ...*{{.current.Name}}) dbr.Builder{
	ors := []dbr.Builder{}
	for _, t := range s {
		ors = append(ors, dbr.And(
				{{range $index, $col := .current.Fields | isPk}}
				J{{$.current.Name}}Model.{{$col.Name}}.Eq( t.{{$col.Name}} ),
				{{end}}
		))
	}
	return dbr.Or(ors...)
}
// In provided items.
func (j j{{.current.Name}}Model) In(s ...*{{.current.Name}}) dbr.Builder{
	ors := []dbr.Builder{}
	for _, t := range s {
		ors = append(ors, dbr.And(
				{{range $index, $col := .current.Fields | isPk}}
				J{{$.current.Name}}Model.{{$col.Name}}.Eq( t.{{$col.Name}} ),
				{{end}}
		))
	}
	return dbr.Or(ors...)
}
{{end}}

// As returns a copy with an alias.
func (j j{{.current.Name}}Model) As (as string) j{{.current.Name}}Model{
	j.as = as
	{{range $i, $col := .current.Fields | withSQLType}}
		j.{{$col.Name}}.TableAlias = as
	{{end}}
	{{range $i, $col := .current.Fields | withoutSQLType}}
		// j.{{$col.Name | ucfirst}}.TableAlias = as
	{{end}}
	return j
}
// Table returns the sql table name
func (j j{{.current.Name}}Model) Table() string {
	return "{{$.current.SQLName}}"
}
// Alias returns the current alias
func (j j{{.current.Name}}Model) Alias() string {
	if j.as == "" {
		return j.Table()
	}
	return j.as
}
// Properties returns a map of property name => meta
func (j j{{.current.Name}}Model) Properties() map[string]builder.MetaProvider {
	ret := map[string]builder.MetaProvider{}
	{{range $i, $col := .current.Fields | withSQLType}}
		ret["{{$col.Name}}"] = j.{{$col.Name}}
	{{end}}
	{{range $i, $col := .current.Fields | withoutSQLType}}
		ret["{{$col.Name | ucfirst}}"] = j.{{$col.Name | ucfirst}}
	{{end}}
	return ret
}
// Fields returns given sql fields with appropriate aliasing.
func (j j{{.current.Name}}Model) Fields(ins ...string) []string {
	dialect := runtime.GetDialect()
	if len(ins)==0 {
		ins = append(ins, "*")
	}
	for i, in := range ins {
		p := j.Table()
		if j.as != "" {
			p = j.as
		}
		if in == "*" {
			ins[i] = fmt.Sprintf("%v.%v", dialect.QuoteIdent(p), in)
		} else {
			ins[i] = fmt.Sprintf("%v.%v", dialect.QuoteIdent(p), dialect.QuoteIdent(in))
		}
	}
	return ins
}

// J{{.current.Name}}Model provides helper to work with {{.current.Name}} data provider
var J{{.current.Name}}Model = j{{.current.Name}}Model{
	{{range $i, $col := .current.Fields | withSQLType}}
		{{$col.Name}}: builder.NewValueMeta(
				{{$col.SQLName | quote}}, {{$col.SQLType | quote}},
				{{$col.Name | quote}}, {{$col.GoType | quote}},
				{{$col.IsPk}}, {{$col.IsAI}},
		),
	{{end}}
	{{range $i, $col := .current.Fields | withoutSQLType}}
		{{$col.Name | ucfirst}}: builder.NewRelMeta(
				{{$col.Name | quote}}, {{$col.GoType | quote}},
				{{$col.HasOne | quote}}, {{$col.HasMany | quote}}, {{$col.On | quote}},
				{{$col.RelType | quote}},
		),
	{{end}}
}


{{if .current.IsAutoGoType }}
// {{.current.Name}} is automatically generated to handle a many to many relationship.
type {{.current.Name}} struct {
	{{range $i, $col := .current.Fields | withGoName}}
		{{$col.Name}} {{$col.GoType}}
	{{end}}
}
{{end}}

type j{{.current.Name}}DeleteBuilder struct {
	*builder.DeleteBuilder
}
// //Build builds the sql string into given buffer using current dialect
// func (c *j{{.current.Name}}DeleteBuilder) Build(b dbr.Buffer) error {
// 	return c.DeleteBuilder.Build(runtime.GetDialect(), b)
// }
// //String returns the sql string for current dialect. It returns empty string if the build returns an error.
// func (c *j{{.current.Name}}DeleteBuilder) String() string {
// 	b := dbr.NewBuffer()
// 	if err := c.Build(b); err != nil {
// 		return ""
// 	}
// 	return b.String()
// }
//Where returns a j{{.current.Name}}DeleteBuilder instead of builder.DeleteBuilder.
func (c *j{{.current.Name}}DeleteBuilder) Where(query interface{}, value ...interface{}) *j{{.current.Name}}DeleteBuilder {
	c.DeleteBuilder.Where(query, value...)
	return c
}

type j{{.current.Name}}SelectBuilder struct {
	as string
	*builder.SelectBuilder
}
// //Build builds the sql string using current dialect into given bufer
// func (c *j{{.current.Name}}SelectBuilder) Build(b dbr.Buffer) error {
// 	return c.SelectBuilder.Build(runtime.GetDialect(), b)
// }
// //String returns the sql string for current dialect. It returns empty string if the build returns an error.
// func (c *j{{.current.Name}}SelectBuilder) String() string {
// 	b := dbr.NewBuffer()
// 	if err := c.Build(b); err != nil {
// 		return ""
// 	}
// 	return b.String()
// }

//Read evaluates current select query and load the results into a {{.current.Name}}
func (c *j{{.current.Name}}SelectBuilder) Read() (*{{.current.Name}}, error) {
	var one {{.current.Name}}
	err := c.LoadStruct(&one)
	return &one, err
}
//ReadAll evaluates current select query and load the results into a slice of {{.current.Name}}
func (c *j{{.current.Name}}SelectBuilder) ReadAll() ([]*{{.current.Name}}, error) {
	var all []*{{.current.Name}}
	_, err := c.LoadStructs(&all)
	return all, err
}
//Where returns a j{{.current.Name}}SelectBuilder instead of builder.SelectBuilder.
func (c *j{{.current.Name}}SelectBuilder) Where(query interface{}, value ...interface{}) *j{{.current.Name}}SelectBuilder {
	c.SelectBuilder.Where(query, value...)
	return c
}
//GroupBy returns a j{{.current.Name}}SelectBuilder instead of builder.SelectBuilder.
func (c *j{{.current.Name}}SelectBuilder) GroupBy(col ...string) *j{{.current.Name}}SelectBuilder {
	c.SelectBuilder.GroupBy(col...)
	return c
}
//Having returns a j{{.current.Name}}SelectBuilder instead of builder.SelectBuilder.
func (c *j{{.current.Name}}SelectBuilder) Having(query interface{}, value ...interface{}) *j{{.current.Name}}SelectBuilder {
	c.SelectBuilder.Having(query, value...)
	return c
}
//Limit returns a j{{.current.Name}}SelectBuilder instead of builder.SelectBuilder.
func (c *j{{.current.Name}}SelectBuilder) Limit(n uint64) *j{{.current.Name}}SelectBuilder {
	c.SelectBuilder.Limit(n)
	return c
}
//Offset returns a j{{.current.Name}}SelectBuilder instead of builder.SelectBuilder.
func (c *j{{.current.Name}}SelectBuilder) Offset(n uint64) *j{{.current.Name}}SelectBuilder {
	c.SelectBuilder.Offset(n)
	return c
}
//OrderAsc returns a j{{.current.Name}}SelectBuilder instead of builder.SelectBuilder.
func (c *j{{.current.Name}}SelectBuilder) OrderAsc(col string) *j{{.current.Name}}SelectBuilder {
	c.SelectBuilder.OrderAsc(col)
	return c
}
//OrderDesc returns a j{{.current.Name}}SelectBuilder instead of builder.SelectBuilder.
func (c *j{{.current.Name}}SelectBuilder) OrderDesc(col string) *j{{.current.Name}}SelectBuilder {
	c.SelectBuilder.OrderDesc(col)
	return c
}
//OrderDir returns a j{{.current.Name}}SelectBuilder instead of builder.SelectBuilder.
func (c *j{{.current.Name}}SelectBuilder) OrderDir(col string, isAsc bool) *j{{.current.Name}}SelectBuilder {
	c.SelectBuilder.OrderDir(col, isAsc)
	return c
}
//OrderBy returns a j{{.current.Name}}SelectBuilder instead of builder.SelectBuilder.
func (c *j{{.current.Name}}SelectBuilder) OrderBy(col string) *j{{.current.Name}}SelectBuilder {
	c.SelectBuilder.OrderBy(col)
	return c
}
//Paginate returns a j{{.current.Name}}SelectBuilder instead of builder.SelectBuilder.
func (c *j{{.current.Name}}SelectBuilder) Paginate(page, perPage uint64) *j{{.current.Name}}SelectBuilder {
	c.SelectBuilder.Paginate(page, perPage)
	return c
}
//Join returns a j{{.current.Name}}SelectBuilder instead of builder.SelectBuilder.
func (c *j{{.current.Name}}SelectBuilder) Join(table, on interface{}) *j{{.current.Name}}SelectBuilder {
	c.SelectBuilder.Join(table, on)
	return c
}
//LeftJoin returns a j{{.current.Name}}SelectBuilder instead of builder.SelectBuilder.
func (c *j{{.current.Name}}SelectBuilder) LeftJoin(table, on interface{}) *j{{.current.Name}}SelectBuilder {
	c.SelectBuilder.LeftJoin(table, on)
	return c
}
//RightJoin returns a j{{.current.Name}}SelectBuilder instead of builder.SelectBuilder.
func (c *j{{.current.Name}}SelectBuilder) RightJoin(table, on interface{}) *j{{.current.Name}}SelectBuilder {
	c.SelectBuilder.RightJoin(table, on)
	return c
}
//FullJoin returns a j{{.current.Name}}SelectBuilder instead of builder.SelectBuilder.
func (c *j{{.current.Name}}SelectBuilder) FullJoin(table, on interface{}) *j{{.current.Name}}SelectBuilder {
	c.SelectBuilder.FullJoin(table, on)
	return c
}
//Distinct returns a j{{.current.Name}}SelectBuilder instead of builder.SelectBuilder.
func (c *j{{.current.Name}}SelectBuilder) Distinct() *j{{.current.Name}}SelectBuilder {
	c.SelectBuilder.Distinct()
	return c
}

// J{{.current.Name}} provides a basic querier
func J{{.current.Name}}(db dbr.SessionRunner) j{{.current.Name}}Querier {
	return j{{.current.Name}}Querier {
		db: db,
	}
}

type j{{.current.Name}}Querier struct {
	db dbr.SessionRunner
	as string
}

//As set alias prior building.
func (c j{{.current.Name}}Querier) As(as string) j{{.current.Name}}Querier {
	c.as = as
	return c
}

//Model returns a model with appropriate aliasing.
func (c j{{.current.Name}}Querier) Model() j{{.current.Name}}Model {
	return J{{.current.Name}}Model.As(c.as)
}


//Select returns a {{.current.Name}} Select Builder.
func (c j{{.current.Name}}Querier) Select(what ...string) *j{{.current.Name}}SelectBuilder {
	m := c.Model()
	dialect := runtime.GetDialect()
	from := dialect.QuoteIdent(m.Table())
	if m.Alias()!="" && m.Alias()!=m.Table() {
		from = fmt.Sprintf("%v as %v", from, dialect.QuoteIdent(m.Alias()))
	}
	if len(what) == 0 {
		// alias := m.Table()
		// if m.Alias()!="" && m.Alias()!=m.Table() {
		// 	alias = m.Alias()
		// }
		// what = m.Fields(alias+".*")
		what = m.Fields("*")
	}
	return &j{{.current.Name}}SelectBuilder{
		as: c.as,
		SelectBuilder: &builder.SelectBuilder{
			SelectBuilder: c.db.Select(what...).From(from),
		},
	}
}

//Where returns a {{.current.Name}} Select Builder.
func (c j{{.current.Name}}Querier) Where(query interface{}, value ...interface{}) *j{{.current.Name}}SelectBuilder {
	return c.Select().Where(query, value...)
}

//Count returns a {{.current.Name}} Select Builder to count given expressions.
func (c j{{.current.Name}}Querier) Count(what ...string) *j{{.current.Name}}SelectBuilder {
	if len(what) == 0 {
			what = append(what, "*")
	}
	return c.Select("COUNT("+strings.Join(what, ",")+")")
}

{{if empty .current.SQLViewSelect .current.SQLViewCreate}}

	// Insert a new {{.current.Name}}, if it has autoincrement primary key, the value will be set.
	// It stops on first error.
	func (c j{{.current.Name}}Querier) Insert(items ...*{{.current.Name}}) (sql.Result,error) {
		var res sql.Result
		var err error
		for _, data := range items {
			{{range $i, $col := .current.Fields | dateTypes}}
				{{if $col.LastUpdated}}
					{{if $col.IsStar}}
					if data.{{$col.Name}} == nil {
						x := time.Now()
						data.{{$col.Name}} = &x
					}
					{
						x := data.{{$col.Name}}.Truncate(time.Microsecond)
						data.{{$col.Name}} = &x
					}
					{{else}}
					data.{{$col.Name}} = data.{{$col.Name}}.Truncate(time.Microsecond)
					{{end}}
				{{else if $col.Insert}}
					{{if $col.IsStar}}
					if data.{{$col.Name}}==nil {
						x := time.Now()
						data.{{$col.Name}} = &x
					}
					{{else}}
					data.{{$col.Name}} = time.Now()
					{{end}}
				{{end}}
			{{end}}
			{{range $i, $col := .current.Fields | dateTypes}}
				{{if $col.UTC}}
					{{if $col.IsStar}}
					if data.{{$col.Name}} != nil {
						x := data.{{$col.Name}}.UTC()
						data.{{$col.Name}} = &x
					}
					{{else}}
					data.{{$col.Name}} = data.{{$col.Name}}.UTC()
					{{end}}
				{{end}}
			{{end}}
			{{range $i, $col := .current.Fields | isPk | isSQLText}}
				if runtime.Runs(drivers.Mysql) {
					{{if $col.IsStar}}
					if data.{{$col.Name}} != nil {
						x := *data.{{$col.Name}}
						if len(x)>255 {
							return nil, fmt.Errorf("{{$col.Name}}: PRIMARY KEY length exceeded max=255, got=%v", len(x))
						}
						data.{{$col.Name}} = &x
					}
					{{else}}
					if len(data.{{$col.Name}})>255 {
						return nil, fmt.Errorf("{{$col.Name}}: PRIMARY KEY length exceeded max=255, got=%v", len(data.{{$col.Name}}))
					}
					{{end}}
				}
			{{end}}
			{{if .current.InsertHook}}
			err = data.beforeInsert()
			if err != nil {
				return nil, err
			}
			{{end}}
			query := c.db.InsertInto(J{{.current.Name}}Model.Table()).Columns(
				{{range $i, $col := .current.Fields | notAI | withSQLType | withGoName}}
					{{$col.SQLName | quote}},
				{{end}}
			).Record(data)
			if runtime.Runs(drivers.Pgsql) {
				{{if .current.Fields | isAI}}
					query = query.Returning(
						{{range $i, $col := .current.Fields | isAI}}
						{{$col.SQLName | quote}},
						{{end}}
					)
					{{range $i, $col := .current.Fields | isAI}}
					var auto{{$i}} {{$col.GoType | itemGoType}}
					{{end}}
					err = query.Load(
						{{range $i, $col := .current.Fields | isAI}}
						&auto{{$i}},
						{{end}}
					)
					{{range $i, $col := .current.Fields | isAI}}
						{{if $col.IsStar}}
						data.{{$col.Name}} = &auto{{$i}}
						{{else}}
						data.{{$col.Name}} = auto{{$i}}
						{{end}}
					{{end}}
				{{else}}
					res, err = query.Exec()
				{{end}}
			} else {
				res, err = query.Exec()
				{{if notEmpty (.current.Fields | isAI | getPkFieldName)}}
				if err == nil {
					id, err2 := res.LastInsertId()
					if err2 != nil {
						return res, err2
					}
					data.{{.current.Fields | isAI | getPkFieldName}} = id
				}
				{{end}}
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
	func (c j{{.current.Name}}Querier) InsertBulk(items ...*{{.current.Name}}) (error) {
		panic("todo")
	}

	{{ if .current.Fields | hasNonPkField }}
		// Update a {{.current.Name}}. It stops on first error.
		func (c j{{.current.Name}}Querier) Update(items ...*{{.current.Name}}) (sql.Result,error) {
			var res sql.Result
			var err error
			for _, data := range items {
				{{range $i, $col := .current.Fields | dateTypes}}
					{{if $col.UTC}}
						{{if $col.IsStar}}
						if data.{{$col.Name}} != nil{
							x := data.{{$col.Name}}.UTC()
							data.{{$col.Name}} = &x
						}
						{{else}}
						data.{{$col.Name}} = data.{{$col.Name}}.UTC()
						{{end}}
					{{end}}
				{{end}}
				{{range $i, $col := .current.Fields | dateTypes}}
					{{if $col.LastUpdated}}
						currentDate := data.{{$col.Name}}
						newDate := time.Now().UTC().Truncate(time.Microsecond)
						{{if $col.IsStar}}
						if currentDate != nil {
							y := currentDate.Truncate(time.Microsecond)
							currentDate = &y
						}
						{{else}}
						currentDate = currentDate.Truncate(time.Microsecond)
						{{end}}
					{{end}}
				{{end}}
				{{if .current.UpdateHook}}
				err = data.beforeUpdate()
				if err != nil {
					return nil, err
				}
				{{end}}
				query := c.db.Update(J{{.current.Name}}Model.Table())
				{{range $i, $col := .current.Fields | notPk | withSQLType | withGoName}}
					{{if $col.LastUpdated}}
						query = query.Set({{.SQLName | quote}}, newDate)
					{{else}}
						query = query.Set({{.SQLName | quote}}, data.{{.Name}})
					{{end}}
				{{end}}
				{{range $i, $col := .current.Fields | isPk}}
					query = query.Where("{{$col.SQLName}} = ?", data.{{$col.Name}})
				{{end}}
				{{range $i, $col := .current.Fields | dateTypes}}
					{{if $col.LastUpdated}}
						{{if $col.IsStar}}
						if currentDate == nil {
							query = query.Where("{{$col.SQLName}} IS NULL")
						} else {
							query = query.Where("{{$col.SQLName}} = ?", currentDate)
						}
						{{else}}
						query = query.Where("{{$col.SQLName}} = ?", currentDate)
						{{end}}
					{{end}}
				{{end}}
				res, err = query.Exec()

				{{range $i, $col := .current.Fields | dateTypes}}
					{{if $col.LastUpdated}}
						if err == nil {
							{{if $col.IsStar}}
							data.{{$col.Name}} = &newDate
							{{else}}
							data.{{$col.Name}} = newDate
							{{end}}
						}
					{{end}}
				{{end}}

				if err != nil {
					return res, err
				}
			}
			return res, err
		}

		// MustUpdate a {{.current.Name}}. It stops on first error. It errors if an update query does not affect row.
		func (c j{{.current.Name}}Querier) MustUpdate(items ...*{{.current.Name}}) (sql.Result,error) {
			var res sql.Result
			var err error
			for _, data := range items {
				{{range $i, $col := .current.Fields | dateTypes}}
					{{if $col.UTC}}
						{{if $col.IsStar}}
						if data.{{$col.Name}} != nil{
							x := data.{{$col.Name}}.UTC()
							data.{{$col.Name}} = &x
						}
						{{else}}
						data.{{$col.Name}} = data.{{$col.Name}}.UTC()
						{{end}}
					{{end}}
				{{end}}
				{{range $i, $col := .current.Fields | dateTypes}}
					{{if $col.LastUpdated}}
						currentDate := data.{{$col.Name}}
						newDate := time.Now().UTC().Truncate(time.Microsecond)
						{{if $col.IsStar}}
						if currentDate != nil {
							y := currentDate.Truncate(time.Microsecond)
							currentDate = &y
						}
						{{else}}
						currentDate = currentDate.Truncate(time.Microsecond)
						{{end}}
					{{end}}
				{{end}}
				{{if .current.UpdateHook}}
				err = data.beforeUpdate()
				if err != nil {
					return nil, err
				}
				{{end}}
				query := c.db.Update(J{{.current.Name}}Model.Table())
				{{range $i, $col := .current.Fields | notPk | withSQLType | withGoName}}
					{{if $col.LastUpdated}}
						query = query.Set({{.SQLName | quote}}, newDate)
					{{else}}
						query = query.Set({{.SQLName | quote}}, data.{{.Name}})
					{{end}}
				{{end}}
				{{range $i, $col := .current.Fields | isPk}}
					query = query.Where("{{$col.SQLName}} = ?", data.{{$col.Name}})
				{{end}}
				{{range $i, $col := .current.Fields | dateTypes}}
					{{if $col.LastUpdated}}
						{{if $col.IsStar}}
						if currentDate == nil {
							query = query.Where("{{$col.SQLName}} IS NULL")
						} else {
							query = query.Where("{{$col.SQLName}} = ?", currentDate)
						}
						{{else}}
						query = query.Where("{{$col.SQLName}} = ?", currentDate)
						{{end}}
					{{end}}
				{{end}}
				res, err = query.Exec()

				if err == nil {
					if n, _ := res.RowsAffected(); n == 0 {
						x := &builder.UpdateBuilder{UpdateBuilder: query}
						err = runtime.NewNoRowsAffected(x.String())
					}
				}

				{{range $i, $col := .current.Fields | dateTypes}}
					{{if $col.LastUpdated}}
						if err == nil {
							{{if $col.IsStar}}
							data.{{$col.Name}} = &newDate
							{{else}}
							data.{{$col.Name}} = newDate
							{{end}}
						}
					{{end}}
				{{end}}

				if err != nil {
					return res, err
				}
			}
			return res, err
		}

		// UpdateBulk updates multiple items into the database.
		// It builds an update query of multiple rows and send it on the underlying connection.
		func (c j{{.current.Name}}Querier) UpdateBulk(items ...*{{.current.Name}}) (error) {
			panic("todo")
		}
	{{end}}


	//Delete returns a delete builder
	func (c j{{.current.Name}}Querier) Delete() *j{{.current.Name}}DeleteBuilder {
		return &j{{.current.Name}}DeleteBuilder{
			&builder.DeleteBuilder{
				DeleteBuilder: c.db.DeleteFrom(J{{.current.Name}}Model.Table()),
			},
		}
	}

	// MustDelete requires the query to affeect rows.
	func (c j{{.current.Name}}Querier) MustDelete() *j{{.current.Name}}DeleteBuilder {
		ret := &j{{.current.Name}}DeleteBuilder{
			&builder.DeleteBuilder{
				DeleteBuilder: c.db.DeleteFrom(J{{.current.Name}}Model.Table()),
			},
		}
		ret.MustDelete()
		return ret
	}

	{{ if .current.Fields | hasPkField }}
		//DeleteByPk deletes one {{.current.Name}} by its PKs
		func (c j{{.current.Name}}Querier) DeleteByPk({{.current.Fields | isPk | AsMethodParams}}) error {
			_, err := c.Delete().Where(
				{{range $i, $col := .current.Fields | isPk}}
				J{{$.current.Name}}Model.{{$col.Name}}.Eq( {{$col.Name}} ),
				{{end}}
			).Exec()
			return err
		}

		// DeleteAll given {{.current.Name}}
		func (c j{{.current.Name}}Querier) DeleteAll(items ...*{{.current.Name}}) (sql.Result, error) {
			q := c.Delete().Where(
				J{{$.current.Name}}Model.In(items...),
			)
			return q.Exec()
		}

		// MustDeleteAll given {{.current.Name}}
		func (c j{{.current.Name}}Querier) MustDeleteAll(items ...*{{.current.Name}}) (sql.Result, error) {
			var res sql.Result
			var err error
			for _, d := range items {
				res, err = c.DeleteAll(d)
				if err != nil {
					return res, err
				}
				if n, e := res.RowsAffected(); e!=nil{
					return res, e
				} else if n == 0 {
					q := c.Delete().Where(
						J{{$.current.Name}}Model.In(items...),
					)
					err = runtime.NewNoRowsAffected(q.String())
					return res, err
				}
			}
			return res, err
		}

		//Find one {{.current.Name}} using its PKs
		func (c j{{.current.Name}}Querier) Find({{.current.Fields | isPk | AsMethodParams}}) (*{{.current.Name}}, error) {
			return c.Select().Where(
				{{range $i, $col := .current.Fields | isPk}}
				J{{$.current.Name}}Model.{{$col.Name}}.Eq( {{$col.Name}} ),
				{{end}}
			).Read()
		}
	{{end}}

{{end}}


{{range $i, $f := .current.Fields | isHasOne}}

	{{$hasOne := getHasOne $.all $.current $f}}

	// Join{{$f.Name | ucfirst}} adds a JOIN to {{$hasOne.Local.Name}}.{{$f.Name | ucfirst}}
	func (c *j{{$.current.Name}}SelectBuilder) Join{{$f.Name | ucfirst}}(
		As{{$f.Name | ucfirst}} string,
	) *j{{$.current.Name}}SelectBuilder {
		dialect := runtime.GetDialect()
		on := ""
		localTable := dialect.QuoteIdent(J{{$hasOne.Local.Name}}Model.Table())
		if c.as != "" {
			localTable = dialect.QuoteIdent(c.as)
		}
		foreiTable := dialect.QuoteIdent(J{{$hasOne.Foreign.Name}}Model.Table())
		if As{{$f.Name | ucfirst}} != "" {
			foreiTable = dialect.QuoteIdent(As{{$f.Name | ucfirst}})
		}
		{{range $i, $col := $hasOne.Fields}}
		on += fmt.Sprintf("%v.%v = %v.%v",
			localTable, dialect.QuoteIdent("{{$col.LocalField.SQLName}}"),
			foreiTable, dialect.QuoteIdent("{{$col.ForeignField.SQLName}}"),
		)
		{{end}}
		if As{{$f.Name | ucfirst}} == "" {
			return c.Join(dbr.I(J{{$hasOne.Foreign.Name}}Model.Table()), on)
		}
		return c.Join(dbr.I(J{{$hasOne.Foreign.Name}}Model.Table()).As(As{{$f.Name | ucfirst}}), on)
	}

	// LeftJoin{{$f.Name | ucfirst}} adds a LEFT JOIN to {{$hasOne.Local.Name}}.{{$f.Name | ucfirst}}
	func (c *j{{$.current.Name}}SelectBuilder) LeftJoin{{$f.Name | ucfirst}}(
		As{{$f.Name | ucfirst}} string,
	) *j{{$.current.Name}}SelectBuilder {
		dialect := runtime.GetDialect()
		on := ""
		localTable := J{{$hasOne.Local.Name}}Model.Table()
		if c.as != "" {
			localTable = dialect.QuoteIdent(c.as)
		}
		foreiTable := dialect.QuoteIdent(J{{$hasOne.Foreign.Name}}Model.Table())
		if As{{$f.Name | ucfirst}} != "" {
			foreiTable = dialect.QuoteIdent(As{{$f.Name | ucfirst}})
		}
		{{range $i, $col := $hasOne.Fields}}
		on += fmt.Sprintf("%v.%v = %v.%v",
			localTable, dialect.QuoteIdent("{{$col.LocalField.SQLName}}"),
			foreiTable, dialect.QuoteIdent("{{$col.ForeignField.SQLName}}"),
		)
		{{end}}
		if As{{$f.Name | ucfirst}} == "" {
			return c.LeftJoin(dbr.I(J{{$hasOne.Foreign.Name}}Model.Table()), on)
		}
		return c.LeftJoin(dbr.I(J{{$hasOne.Foreign.Name}}Model.Table()).As(As{{$f.Name | ucfirst}}), on)
	}

	// // RightJoin{{$f.Name | ucfirst}} adds a RIGHT JOIN to {{$hasOne.Local.Name}}.{{$f.Name | ucfirst}}
	// func (c *j{{$.current.Name}}SelectBuilder) RightJoin{{$f.Name | ucfirst}}(
	// 	As{{$f.Name | ucfirst}} string,
	// ) *j{{$.current.Name}}SelectBuilder {
	// 	dialect := runtime.GetDialect()
	// 	on := ""
	// 	localTable := dialect.QuoteIdent(J{{$hasOne.Local.Name}}Model.Table())
	// 	if c.as != "" {
	// 		localTable = dialect.QuoteIdent(c.as)
	// 	}
	// 	foreiTable := dialect.QuoteIdent(J{{$hasOne.Foreign.Name}}Model.Table())
	// 	if As{{$f.Name | ucfirst}} != "" {
	// 		foreiTable = dialect.QuoteIdent(As{{$f.Name | ucfirst}})
	// 	}
	// 	{{range $i, $col := $hasOne.Fields}}
	// 	on += fmt.Sprintf("%v.%v = %v.%v",
	// 		localTable, dialect.QuoteIdent("{{$col.LocalField.SQLName}}"),
	// 		foreiTable, dialect.QuoteIdent("{{$col.ForeignField.SQLName}}"),
	// 	)
	// 	{{end}}
	// 	if As{{$f.Name | ucfirst}} == "" {
	// 		return c.RightJoin(dbr.I(J{{$hasOne.Foreign.Name}}Model.Table()), on)
	// 	}
	// 	return c.RightJoin(dbr.I(J{{$hasOne.Foreign.Name}}Model.Table()).As(As{{$f.Name | ucfirst}}), on)
	// }

	// FullJoin{{$f.Name | ucfirst}} adds a FULL JOIN to {{$hasOne.Local.Name}}.{{$f.Name | ucfirst}}
	func (c *j{{$.current.Name}}SelectBuilder) FullJoin{{$f.Name | ucfirst}}(
		As{{$f.Name | ucfirst}} string,
	) *j{{$.current.Name}}SelectBuilder {
		dialect := runtime.GetDialect()
		on := ""
		localTable := dialect.QuoteIdent(J{{$hasOne.Local.Name}}Model.Table())
		if c.as != "" {
			localTable = dialect.QuoteIdent(c.as)
		}
		foreiTable := dialect.QuoteIdent(J{{$hasOne.Foreign.Name}}Model.Table())
		if As{{$f.Name | ucfirst}} != "" {
			foreiTable = dialect.QuoteIdent(As{{$f.Name | ucfirst}})
		}
		{{range $i, $col := $hasOne.Fields}}
		on += fmt.Sprintf("%v.%v = %v.%v",
			localTable, dialect.QuoteIdent("{{$col.LocalField.SQLName}}"),
			foreiTable, dialect.QuoteIdent("{{$col.ForeignField.SQLName}}"),
		)
		{{end}}
		if As{{$f.Name | ucfirst}} == "" {
			return c.FullJoin(dbr.I(J{{$hasOne.Foreign.Name}}Model.Table()), on)
		}
		return c.FullJoin(dbr.I(J{{$hasOne.Foreign.Name}}Model.Table()).As(As{{$f.Name | ucfirst}}), on)
	}

	{{$foreign := $f.HasOne | findStruct $.all}}
	// {{$f.Name | ucfirst}} reads associated object
	func (g *{{$.current.Name}}) {{$f.Name | ucfirst}}(db dbr.SessionRunner) ({{$f.GoType}}, error) {
		q := J{{$f.GoType | itemGoType}}(db).Select()
		q = q.Where(
			{{range $i, $col := $foreign.Fields | isPk}}
			J{{$f.GoType | itemGoType}}Model.{{$col.Name}}.Eq(g.{{$f.Name |ucfirst}}{{$col.Name}}),
			{{end}}
		)
		return q.Read()
	}

	// Set{{$f.Name | ucfirst}} copies pk values to this object properties
	func (g *{{$.current.Name | itemGoType}}) Set{{$f.Name | ucfirst}}(o {{$f.GoType}}) *{{$.current.Name | itemGoType}} {
		{{range $i, $col := $foreign.Fields | isPk}}
			{{if $f.IsStar}}
			if o == nil {
				g.{{$f.Name | ucfirst}}{{$col.Name}} = nil
			} else {
				{{if $col.IsStar}}
					g.{{$f.Name | ucfirst}}{{$col.Name}} = o.{{$col.Name}}
				{{else}}
				g.{{$f.Name | ucfirst}}{{$col.Name}} = &o.{{$col.Name}}
				{{end}}
			}
			{{end}}
			{{if not $f.IsStar}}
				{{if $col.IsStar}}
				if o.{{$col.Name}} == nil {
					g.{{$f.Name | ucfirst}}{{$col.Name}} = nil
					var def {{$f.GoType}}
					g.{{$f.Name | ucfirst}}{{$col.Name}} = def
				} else {
					g.{{$f.Name | ucfirst}}{{$col.Name}} = *o.{{$col.Name}}
				}
				{{else}}
				g.{{$f.Name | ucfirst}}{{$col.Name}} = &o.{{$col.Name}}
				{{end}}
			{{end}}
		{{end}}
		return g
	}

	// Unset{{$f.Name | ucfirst}} set defaults values to this object properties
	func (g *{{$.current.Name | itemGoType}}) Unset{{$f.Name | ucfirst}}() *{{$.current.Name | itemGoType}}{
		{{$localFields := findLocals $.all $.current $f}}
		{{range $i, $col := $localFields}}
		var def{{$i}} {{$col.GoType}}
		{{end}}
		{{range $i, $col := $localFields}}
		g.{{$col.Name}} = def{{$i}}
		{{end}}
		{{if $f.IsStar}}
		g.{{$f.Name}} = nil
		{{end}}
		return g
	}
{{end}}

{{range $i, $f := .current.Fields | isHasMany2Many .all}}

	{{$m2m := getMany2Many $.all $.current $f}}

	// {{$f.Name | ucfirst}} returns a query builder to select {{$f.Name | ucfirst}} linked to this {{$.current.Name}}
	func (g *{{$.current.Name}}) {{$f.Name | ucfirst}}(db dbr.SessionRunner,
		As{{$m2m.Foreign.Name}}, As{{$m2m.Middle.Name}} , As{{$m2m.Local.Name}} string,
	) *j{{$m2m.Foreign.Name}}SelectBuilder {

		leftTable := J{{$m2m.Foreign.Name}}Model.Table()
		var query *j{{$m2m.Foreign.Name}}SelectBuilder
		if As{{$m2m.Foreign.Name}} != "" {
			leftTable = As{{$m2m.Foreign.Name}}
			query = J{{$m2m.Foreign.Name}}(db).As(As{{$m2m.Foreign.Name}}).Select(As{{$m2m.Foreign.Name}} + ".*")
		} else {
			query = J{{$m2m.Foreign.Name}}(db).Select(leftTable + ".*")
		}

		midTable := J{{$m2m.Middle.Name}}Model.Table()
		{
			on := ""
			if As{{$m2m.Middle.Name}} != "" {
				midTable = As{{$m2m.Middle.Name}}
			}
			{{range $i, $j := $m2m.FMFields}}
			on += fmt.Sprintf("%v.%v = %v.%v",
				midTable, "{{$j.ForeignField.SQLName}}",
				leftTable, "{{$j.LocalField.SQLName}}",
				)
			{{end}}

			if As{{$m2m.Middle.Name}} == "" {
				query = query.Join(dbr.I(J{{$m2m.Middle.Name}}Model.Table()), on)
			} else {
				query = query.Join(dbr.I(J{{$m2m.Middle.Name}}Model.Table()).As(As{{$m2m.Middle.Name}}), on)
			}
		}

		rightTable := J{{$m2m.Local.Name}}Model.Table()
		{
			on := ""
			if As{{$m2m.Local.Name}} != "" {
				rightTable = As{{$m2m.Local.Name}}
			}
			{{range $i, $j := $m2m.LMFields}}
			on += fmt.Sprintf("%v.%v = %v.%v",
				midTable, "{{$j.ForeignField.SQLName}}",
				rightTable, "{{$j.LocalField.SQLName}}",
				)
			{{end}}

			if As{{$m2m.Local.Name}} == "" {
				query = query.Join(dbr.I(J{{$m2m.Local.Name}}Model.Table()), on)
			} else {
				query = query.Join(dbr.I(J{{$m2m.Local.Name}}Model.Table()).As(As{{$m2m.Local.Name}}), on)
			}
		}

		{
			m := J{{$m2m.Local.Name}}Model
			if As{{$m2m.Local.Name}} != "" {
				 m = m.As(As{{$m2m.Local.Name}})
			} else {
				m = m.As(m.Table())
			}
			query = query.Where(
				{{range $i, $col := $m2m.Local.Pks}}
				m.{{$col.Name}}.Eq(g.{{$col.Name}}),
				{{end}}
			)
		}

		return query
	}

	//LinkWith{{$f.Name | ucfirst}} writes new links with {{$m2m.Local.Name}}.
	func (g *{{$.current.Name}}) LinkWith{{$f.Name | ucfirst}}(db dbr.SessionRunner, items ...*{{$f.GoType | itemGoType}}) (sql.Result, error) {
		toInsert := []*{{$m2m.Middle.Name}}{}
		for _, item := range items {
			toInsert = append(toInsert, &{{$m2m.Middle.Name}}{
				{{range $i, $j := $m2m.FMFields}}
					{{$j.ForeignField.Name}}: item.{{$j.LocalField.Name}},
				{{end}}
				{{range $i, $j := $m2m.LMFields}}
					{{$j.ForeignField.Name}}: g.{{$j.LocalField.Name}},
				{{end}}
			})
		}
		return J{{$m2m.Middle.Name}}(db).Insert(toInsert...)
	}

	//UnlinkWith{{$f.Name | ucfirst}} deletes given existing links with {{$m2m.Local.Name}}.
	func (g *{{$.current.Name}}) UnlinkWith{{$f.Name | ucfirst}}(db dbr.SessionRunner, items ...*{{$f.GoType | itemGoType}}) (sql.Result, error) {
		toDelete := []*{{$m2m.Middle.Name}}{}
		for _, item := range items {
			toDelete = append(toDelete, &{{$m2m.Middle.Name}}{
				{{range $i, $j := $m2m.FMFields}}
					{{$j.ForeignField.Name}}: item.{{$j.LocalField.Name}},
				{{end}}
				{{range $i, $j := $m2m.LMFields}}
					{{$j.ForeignField.Name}}: g.{{$j.LocalField.Name}},
				{{end}}
			})
		}
		return J{{$m2m.Middle.Name}}(db).DeleteAll(toDelete...)
	}

	//UnlinkAll{{$f.Name | ucfirst}} deletes all existing links with {{$m2m.Local.Name}}.
	func (g *{{$.current.Name}}) UnlinkAll{{$f.Name | ucfirst}}(db dbr.SessionRunner) (sql.Result, error) {
		return J{{$m2m.Middle.Name}}(db).Delete().Where(
			{{range $i, $j := $m2m.LMFields}}
				J{{$m2m.Middle.Name}}Model.{{$j.ForeignField.Name}}.Eq(g.{{$j.LocalField.Name}}),
			{{end}}
		).Exec()
	}

	//Set{{$f.Name | ucfirst}} replaces existing links with {{$m2m.Local.Name}}.
	func (g *{{$.current.Name}}) Set{{$f.Name | ucfirst}}(db dbr.SessionRunner, items ...*{{$f.GoType | itemGoType}}) (sql.Result, error) {
		if res, err := g.UnlinkAll{{$f.Name | ucfirst}}(db); err != nil {
			return res, err
		}
		return g.LinkWith{{$f.Name | ucfirst}}(db, items...)
	}

	// Join{{$f.Name | ucfirst}} adds a JOIN to {{$m2m.Local.Name}}.{{$f.Name | ucfirst}}
	func (c *j{{$.current.Name}}SelectBuilder) Join{{$f.Name | ucfirst}}(
		As{{$m2m.Middle.Name}}, As{{$m2m.Foreign.Name}} string,
	) *j{{$m2m.Local.Name}}SelectBuilder {

		query := c

		leftTable := J{{$m2m.Local.Name}}Model.Table()
		if c.as != "" {
			leftTable = c.as
		}

		midTable := J{{$m2m.Middle.Name}}Model.Table()
		if As{{$m2m.Middle.Name}} != "" {
			midTable = As{{$m2m.Middle.Name}}
		}

		{
			on := ""
			{{range $i, $j := $m2m.LMFields}}
			on += fmt.Sprintf("%v.%v = %v.%v",
				midTable, "{{$j.ForeignField.SQLName}}",
				leftTable, "{{$j.LocalField.SQLName}}",
				)
			{{end}}

			if As{{$m2m.Middle.Name}} == "" {
				query = query.Join(dbr.I(J{{$m2m.Middle.Name}}Model.Table()), on)
			} else {
				query = query.Join(dbr.I(J{{$m2m.Middle.Name}}Model.Table()).As(As{{$m2m.Middle.Name}}), on)
			}
		}

		{
			rightTable := J{{$m2m.Foreign.Name}}Model.Table()
			if As{{$m2m.Foreign.Name}} != "" {
				rightTable = As{{$m2m.Foreign.Name}}
			}
			on := ""
			{{range $i, $j := $m2m.FMFields}}
			on += fmt.Sprintf("%v.%v = %v.%v",
				midTable, "{{$j.ForeignField.SQLName}}",
				rightTable, "{{$j.LocalField.SQLName}}",
				)
			{{end}}

			if As{{$m2m.Foreign.Name}} == "" {
				query = query.Join(dbr.I(J{{$m2m.Foreign.Name}}Model.Table()), on)
			} else {
				query = query.Join(dbr.I(J{{$m2m.Foreign.Name}}Model.Table()).As(As{{$m2m.Foreign.Name}}), on)
			}
		}

		return query
	}

	// LeftJoin{{$f.Name | ucfirst}} adds a LEFT JOIN to {{$m2m.Local.Name}}.{{$f.Name | ucfirst}}
	func (c *j{{$.current.Name}}SelectBuilder) LeftJoin{{$f.Name | ucfirst}}(
		As{{$m2m.Middle.Name}}, As{{$m2m.Foreign.Name}} string,
	) *j{{$m2m.Local.Name}}SelectBuilder {

		query := c

		leftTable := J{{$m2m.Local.Name}}Model.Table()
		if c.as != "" {
			leftTable = c.as
		}

		midTable := J{{$m2m.Middle.Name}}Model.Table()
		if As{{$m2m.Middle.Name}} != "" {
			midTable = As{{$m2m.Middle.Name}}
		}

		{
			on := ""
			{{range $i, $j := $m2m.LMFields}}
			on += fmt.Sprintf("%v.%v = %v.%v",
				midTable, "{{$j.ForeignField.SQLName}}",
				leftTable, "{{$j.LocalField.SQLName}}",
				)
			{{end}}

			if As{{$m2m.Middle.Name}} == "" {
				query = query.LeftJoin(dbr.I(J{{$m2m.Middle.Name}}Model.Table()), on)
			} else {
				query = query.LeftJoin(dbr.I(J{{$m2m.Middle.Name}}Model.Table()).As(As{{$m2m.Middle.Name}}), on)
			}
		}

		{
			rightTable := J{{$m2m.Foreign.Name}}Model.Table()
			if As{{$m2m.Foreign.Name}} != "" {
				rightTable = As{{$m2m.Foreign.Name}}
			}
			on := ""
			{{range $i, $j := $m2m.FMFields}}
			on += fmt.Sprintf("%v.%v = %v.%v",
				midTable, "{{$j.ForeignField.SQLName}}",
				rightTable, "{{$j.LocalField.SQLName}}",
				)
			{{end}}

			if As{{$m2m.Foreign.Name}} == "" {
				query = query.LeftJoin(dbr.I(J{{$m2m.Foreign.Name}}Model.Table()), on)
			} else {
				query = query.LeftJoin(dbr.I(J{{$m2m.Foreign.Name}}Model.Table()).As(As{{$m2m.Foreign.Name}}), on)
			}
		}

		return query
	}

	// // RightJoin{{$f.Name | ucfirst}} adds a RIGHT JOIN to {{$m2m.Local.Name}}.{{$f.Name | ucfirst}}
	// func (c *j{{$.current.Name}}SelectBuilder) RightJoin{{$f.Name | ucfirst}}(
	// 	As{{$m2m.Middle.Name}}, As{{$m2m.Foreign.Name}} string,
	// ) *j{{$m2m.Local.Name}}SelectBuilder {
	//
	// 	query := c
	//
	// 	leftTable := J{{$m2m.Local.Name}}Model.Table()
	// 	if c.as != "" {
	// 		leftTable = c.as
	// 	}
	//
	// 	midTable := J{{$m2m.Middle.Name}}Model.Table()
	// 	if As{{$m2m.Middle.Name}} != "" {
	// 		midTable = As{{$m2m.Middle.Name}}
	// 	}
	//
	// 	{
	// 		on := ""
	// 		{{range $i, $j := $m2m.LMFields}}
	// 		on += fmt.Sprintf("%v.%v = %v.%v",
	// 			midTable, "{{$j.ForeignField.SQLName}}",
	// 			leftTable, "{{$j.LocalField.SQLName}}",
	// 			)
	// 		{{end}}
	//
	// 		if As{{$m2m.Middle.Name}} == "" {
	// 			query = query.RightJoin(dbr.I(J{{$m2m.Middle.Name}}Model.Table()), on)
	// 		} else {
	// 			query = query.RightJoin(dbr.I(J{{$m2m.Middle.Name}}Model.Table()).As(As{{$m2m.Middle.Name}}), on)
	// 		}
	// 	}
	//
	// 	{
	// 		rightTable := J{{$m2m.Foreign.Name}}Model.Table()
	// 		if As{{$m2m.Foreign.Name}} != "" {
	// 			rightTable = As{{$m2m.Foreign.Name}}
	// 		}
	// 		on := ""
	// 		{{range $i, $j := $m2m.FMFields}}
	// 		on += fmt.Sprintf("%v.%v = %v.%v",
	// 			midTable, "{{$j.ForeignField.SQLName}}",
	// 			rightTable, "{{$j.LocalField.SQLName}}",
	// 			)
	// 		{{end}}
	//
	// 		if As{{$m2m.Foreign.Name}} == "" {
	// 			query = query.RightJoin(dbr.I(J{{$m2m.Foreign.Name}}Model.Table()), on)
	// 		} else {
	// 			query = query.RightJoin(dbr.I(J{{$m2m.Foreign.Name}}Model.Table()).As(As{{$m2m.Foreign.Name}}), on)
	// 		}
	// 	}
	//
	// 	return query
	// }

{{end}}

{{range $i, $f := .current.Fields | isHasMany2One .all}}

	{{$m2o := getMany2One $.all $.current $f}}

	// {{$f.Name | ucfirst}} returns a query builder to select {{$f.Name | ucfirst}} linked to this {{$.current.Name}}
	func (g *{{$.current.Name}}) {{$f.Name | ucfirst}}(db dbr.SessionRunner,
		As{{$m2o.ForeignField.Name | ucfirst}}, As{{$m2o.LocalField.Name | ucfirst}} string,
	) *j{{$m2o.Foreign.Name}}SelectBuilder {

		var query *j{{$m2o.Foreign.Name}}SelectBuilder

		leftTable := J{{$m2o.Foreign.Name}}Model.Table()
		if As{{$m2o.ForeignField.Name | ucfirst}} != "" {
			leftTable = As{{$m2o.ForeignField.Name | ucfirst}}
			query = J{{$m2o.Foreign.Name}}(db).As(As{{$m2o.ForeignField.Name | ucfirst}}).Select(leftTable + ".*")
		} else {
			query = J{{$m2o.Foreign.Name}}(db).Select(leftTable + ".*")
		}

		rightTable := J{{$m2o.Local.Name}}Model.Table()
		if As{{$m2o.LocalField.Name | ucfirst}} != "" {
			rightTable = As{{$m2o.LocalField.Name | ucfirst}}
		}

		on := ""
		{{range $i, $j := $m2o.Fields}}
		on += fmt.Sprintf("%v.%v = %v.%v",
			leftTable, "{{$j.ForeignField.SQLName}}",
			rightTable, "{{$j.LocalField.SQLName}}",
			)
		{{end}}

		if As{{$m2o.LocalField.Name | ucfirst}} == "" {
			return query.Join(dbr.I(J{{$m2o.Local.Name}}Model.Table()), on)
		}
		return query.Join(dbr.I(J{{$m2o.Local.Name}}Model.Table()).As(As{{$m2o.LocalField.Name | ucfirst}}), on)
	}


	// Join{{$f.Name | ucfirst}} adds a JOIN to {{$m2o.Local.Name}}.{{$f.Name | ucfirst}}
	func (c *j{{$.current.Name}}SelectBuilder) Join{{$f.Name | ucfirst}}(
		As{{$f.Name | ucfirst}} string,
	) *j{{$.current.Name}}SelectBuilder {
		dialect := runtime.GetDialect()
		on := ""
		localTable := dialect.QuoteIdent(J{{$m2o.Local.Name}}Model.Table())
		if c.as != "" {
			localTable = dialect.QuoteIdent(c.as)
		}
		foreiTable := dialect.QuoteIdent(J{{$m2o.Foreign.Name}}Model.Table())
		if As{{$f.Name | ucfirst}} != "" {
			foreiTable = dialect.QuoteIdent(As{{$f.Name | ucfirst}})
		}
		{{range $i, $col := $m2o.Fields}}
		on += fmt.Sprintf("%v.%v = %v.%v",
			localTable, dialect.QuoteIdent("{{$col.LocalField.SQLName}}"),
			foreiTable, dialect.QuoteIdent("{{$col.ForeignField.SQLName}}"),
		)
		{{end}}
		if As{{$f.Name | ucfirst}} == "" {
			return c.Join(dbr.I(J{{$m2o.Foreign.Name}}Model.Table()), on)
		}
		return c.Join(dbr.I(J{{$m2o.Foreign.Name}}Model.Table()).As(As{{$f.Name | ucfirst}}), on)
	}

	// LeftJoin{{$f.Name | ucfirst}} adds a LEFT JOIN to {{$m2o.Local.Name}}.{{$f.Name | ucfirst}}
	func (c *j{{$.current.Name}}SelectBuilder) LeftJoin{{$f.Name | ucfirst}}(
		As{{$f.Name | ucfirst}} string,
	) *j{{$.current.Name}}SelectBuilder {
		dialect := runtime.GetDialect()
		on := ""
		localTable := dialect.QuoteIdent(J{{$m2o.Local.Name}}Model.Table())
		if c.as != "" {
			localTable = dialect.QuoteIdent(c.as)
		}
		foreiTable := dialect.QuoteIdent(J{{$m2o.Foreign.Name}}Model.Table())
		if As{{$f.Name | ucfirst}} != "" {
			foreiTable = dialect.QuoteIdent(As{{$f.Name | ucfirst}})
		}
		{{range $i, $col := $m2o.Fields}}
		on += fmt.Sprintf("%v.%v = %v.%v",
			localTable, dialect.QuoteIdent("{{$col.LocalField.SQLName}}"),
			foreiTable, dialect.QuoteIdent("{{$col.ForeignField.SQLName}}"),
		)
		{{end}}
		if As{{$f.Name | ucfirst}} == "" {
			return c.LeftJoin(dbr.I(J{{$m2o.Foreign.Name}}Model.Table()), on)
		}
		return c.LeftJoin(dbr.I(J{{$m2o.Foreign.Name}}Model.Table()).As(As{{$f.Name | ucfirst}}), on)
	}

	// // RightJoin{{$f.Name | ucfirst}} adds a Right JOIN to {{$m2o.Local.Name}}.{{$f.Name | ucfirst}}
	// func (c *j{{$.current.Name}}SelectBuilder) RightJoin{{$f.Name | ucfirst}}(
	// 	As{{$f.Name | ucfirst}} string,
	// ) *j{{$.current.Name}}SelectBuilder {
	// 	dialect := runtime.GetDialect()
	// 	on := ""
	// 	localTable := dialect.QuoteIdent(J{{$m2o.Local.Name}}Model.Table())
	// 	if c.as != "" {
	// 		localTable = dialect.QuoteIdent(c.as)
	// 	}
	// 	foreiTable := dialect.QuoteIdent(J{{$m2o.Foreign.Name}}Model.Table())
	// 	if As{{$f.Name | ucfirst}} != "" {
	// 		foreiTable = dialect.QuoteIdent(As{{$f.Name | ucfirst}})
	// 	}
	// 	{{range $i, $col := $m2o.Fields}}
	// 	on += fmt.Sprintf("%v.%v = %v.%v",
	// 		localTable, dialect.QuoteIdent("{{$col.LocalField.SQLName}}"),
	// 		foreiTable, dialect.QuoteIdent("{{$col.ForeignField.SQLName}}"),
	// 	)
	// 	{{end}}
	// 	if As{{$f.Name | ucfirst}} == "" {
	// 		return c.RightJoin(dbr.I(J{{$m2o.Foreign.Name}}Model.Table()), on)
	// 	}
	// 	return c.RightJoin(dbr.I(J{{$m2o.Foreign.Name}}Model.Table()).As(As{{$f.Name | ucfirst}}), on)
	// }

{{end}}
`))

}
