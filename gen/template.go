package gen

import (
	"text/template"
)

// templates exported for the generator
var (
	Prolog      *template.Template
	Struct      *template.Template
	HelpersDecl *template.Template
	HelpersBody *template.Template
)

func init() {

	Prolog = template.Must(template.New("prolog").Funcs(funcs).Parse(`
// Generated with mh-cbon/jedi. Do not edit by hand.
package {{.PackageName}}

import (
	"strings"
	"database/sql"
	"github.com/gocraft/dbr"
	"github.com/mh-cbon/jedi/runtime"
	"github.com/mh-cbon/jedi/drivers"
	"github.com/mh-cbon/jedi/builder"
)

func init(){
	runtime.Register(
		{{range $i, $s := $.all}}
		J{{.Name}}Setup,
		{{end}}
	)
}

`))

	Struct = template.Must(template.New("struct").Funcs(funcs).Parse(`
type j{{.current.Name}}Setup struct {
	Name string
	CreateStmt string
	DropStmt string
}

func (c j{{.current.Name}}Setup) Create(db *dbr.Connection) error {
	_, err := db.Exec(c.CreateStmt)
	return err
}
func (c j{{.current.Name}}Setup) Drop(db *dbr.Connection) error {
	_, err := db.Exec(c.DropStmt)
	return err
}

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
				create = {{.current | createTable "pgsql" | quote}}
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
				drop = {{.current | dropTable "pgsql" | quote}}
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
				create = {{.current | createView "pgsql" | quote}}
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
				drop = {{.current | dropView "pgsql" | quote}}
			}
		{{end}}
	{{end}}

	return j{{.current.Name}}Setup {
		Name: {{.current.SQLName | quote}},
		CreateStmt: create,
		DropStmt: drop,
	}
}

var J{{.current.Name}}Model = struct {
	{{if notEmpty (.current.Fields | getPkFieldName)}}
	Eq func(...*{{.current.Name}}) dbr.Builder
	In func(...*{{.current.Name}}) dbr.Builder
	{{end}}
	{{range $i, $col := .current.Fields | withGoName}}
		{{$col | printHelperDecl | trim }}
	{{end}}
}{
	{{if notEmpty (.current.Fields | getPkFieldName)}}
	Eq: func(s ...*{{.current.Name}}) dbr.Builder {
		ors := []dbr.Builder{}
		for _, t := range s {
			ors = append(ors, dbr.Or(
				dbr.And(
					{{range $index, $col := .current.Fields | isPk}}
					dbr.Eq({{$col.SQLName | quote}}, t.{{$col.Name}} ),
					{{end}}
				),
			))
		}
		return dbr.And(ors...)
	},
	In: func(s ...*{{.current.Name}}) dbr.Builder {
		ors := []dbr.Builder{}
		for _, t := range s {
			ors = append(ors, dbr.Or(
				dbr.And(
					{{range $index, $col := .current.Fields | isPk}}
					dbr.Eq({{$col.SQLName | quote}}, t.{{$col.Name}} ),
					{{end}}
				),
			))
		}
		return dbr.And(ors...)
	},
	{{end}}
	{{range $i, $col := .current.Fields | withGoName}}
		{{$col | printHelperBody | trim | trail ","}}
	{{end}}
}

{{if .current.IsAutoGoType }}
// {{.current.Name}} ...
type {{.current.Name}} struct {
	{{range $i, $col := .current.Fields | withGoName}}
		{{$col.Name}} {{$col.GoType}}
	{{end}}
}
{{end}}

// J{{.current.Name}} provides a basic querier
func J{{.current.Name}}(db dbr.SessionRunner) j{{.current.Name}}Querier {
	return j{{.current.Name}}Querier {
		name: {{.current.SQLName | quote}},
		db: db,
	}
}

type j{{.current.Name}}SelectBuilder struct {
	*builder.SelectBuilder
}
func (c *j{{.current.Name}}SelectBuilder) Read() (*{{.current.Name}}, error) {
	var one {{.current.Name}}
	err := c.LoadStruct(&one)
	return &one, err
}
func (c *j{{.current.Name}}SelectBuilder) ReadAll() ([]*{{.current.Name}}, error) {
	var all []*{{.current.Name}}
	_, err := c.LoadStructs(&all)
	return all, err
}
func (c *j{{.current.Name}}SelectBuilder) Where(query interface{}, value ...interface{}) *j{{.current.Name}}SelectBuilder {
	c.SelectBuilder.Where(query, value...)
	return c
}

type j{{.current.Name}}Querier struct {
	name string
	db dbr.SessionRunner
}

func (c j{{.current.Name}}Querier) Select(what ...string) *j{{.current.Name}}SelectBuilder {
	if len(what) == 0 {
			what = append(what, "*")
	}
	return &j{{.current.Name}}SelectBuilder{
		&builder.SelectBuilder{
			SelectBuilder: c.db.Select(what...).From(c.name),
		},
	}
}
func (c j{{.current.Name}}Querier) Count(what ...string) *j{{.current.Name}}SelectBuilder {
	if len(what) == 0 {
			what = append(what, "*")
	}
	return c.Select("COUNT("+strings.Join(what, ",")+")")
}

{{if empty .current.SQLViewSelect .current.SQLViewCreate}}
	func (c j{{.current.Name}}Querier) Insert(data *{{.current.Name}}) (sql.Result,error) {
		res, err := c.db.InsertInto(c.name).Columns(
			{{range $i, $col := .current.Fields | notAI | withSQLType | withGoName}}
			{{$col.SQLName | quote}},
			{{end}}
		).Record(data).Exec()
		{{if notEmpty (.current.Fields | isAI | getPkFieldName)}}
		if err == nil {
			id, err2 := res.LastInsertId()
			if err2 != nil {
				return res, err2
			}
			data.{{.current.Fields | isAI | getPkFieldName}} = id
		}
		{{end}}
		return res, err
	}

	{{ if .current.Fields | hasNonPkField }}
		func (c j{{.current.Name}}Querier) Update(data *{{.current.Name}}) (sql.Result,error) {
			res, err := c.db.Update(c.name).
				{{range $i, $col := .current.Fields | notPk | withSQLType | withGoName}}
					Set({{.SQLName | quote}}, data.{{.Name}}).
				{{end}}
				{{range $i, $col := .current.Fields | isPk}}
					Where("{{$col.SQLName}} = ?", data.{{$col.Name}}).
				{{end}}
				Exec()
			return res, err
		}
	{{end}}

	func (c j{{.current.Name}}Querier) Delete() *builder.DeleteBuilder {
		return &builder.DeleteBuilder{
			DeleteBuilder: c.db.DeleteFrom(c.name),
		}
	}

	{{ if .current.Fields | hasPkField }}
		func (c j{{.current.Name}}Querier) DeleteByPk({{.current.Fields | isPk | AsMethodParams}}) error {
			_, err := c.Delete().Where(
				{{range $i, $col := .current.Fields | isPk}}
				J{{$.current.Name}}Model.{{$col.Name}}.Eq( {{$col.SQLName}} ),
				{{end}}
			).Exec()
			return err
		}
		func (c j{{.current.Name}}Querier) DeleteAll(items ...*{{.current.Name}}) error {
			q := c.Delete()
			for _, item := range items {
				q = q.Where(
					dbr.Or(
						dbr.And(
							{{range $i, $col := .current.Fields | isPk}}
							J{{$.current.Name}}Model.{{$col.Name}}.Eq( item.{{$col.Name}} ),
							{{end}}
						),
					),
				)
			}
			_, err := q.Exec()
			return err
		}
		func (c j{{.current.Name}}Querier) Find({{.current.Fields | isPk | AsMethodParams}}) (*{{.current.Name}}, error) {
			return c.Select().Where(
				{{range $i, $col := .current.Fields | isPk}}
				J{{$.current.Name}}Model.{{$col.Name}}.Eq( {{$col.SQLName}} ),
				{{end}}
			).Read()
		}
	{{end}}

{{end}}

{{if gt (len (.current.Fields | isHasOne)) 0}}
	{{range $i, $f := .current.Fields | isHasOne}}
		func (g {{$.current.Name}}) {{$f.Name | ucfirst}}(db dbr.SessionRunner) ({{$f.GoType}}, error) {
			q := J{{$f.GoType | itemGoType}}(db).Select()
			{{$foreign := $f.HasOne | findStruct $.all}}
			q = q.Where(
				{{range $i, $col := $foreign.Fields | isPk}}
				J{{$f.GoType | itemGoType}}Model.{{$col.Name}}.Eq(g.{{$f.Name |ucfirst}}{{$col.Name}}),
				{{end}}
			)
			return q.Read()
		}
	{{end}}
{{end}}
`))

	HelpersDecl = template.Must(template.New("helpersDecl").Funcs(funcs).Parse(`
{{ if eq .SQLType "INTEGER"}}
	{{.Name}} struct {
		Name string
		IsPk bool
		IsAI bool
		Eq func(interface{}) dbr.Builder
		In func(...interface{}) dbr.Builder
		Gt func(interface{}) dbr.Builder
		Gte func(interface{}) dbr.Builder
		Lt func(interface{}) dbr.Builder
		Lte func(interface{}) dbr.Builder
	}
{{end}}
{{ if eq .SQLType "TEXT"}}
	{{.Name}} struct {
		Name string
		Eq func(interface{}) dbr.Builder
		In func(...interface{}) dbr.Builder
		Like func(interface{}) dbr.Builder
	}
{{end}}
`))

	HelpersBody = template.Must(template.New("helpersBody").Funcs(funcs).Parse(`
{{ if eq .SQLType "INTEGER"}}
	{{.Name}}: struct {
		Name string
		IsPk bool
		IsAI bool
		Eq func(interface{}) dbr.Builder
		In func(...interface{}) dbr.Builder
		Gt func(interface{}) dbr.Builder
		Gte func(interface{}) dbr.Builder
		Lt func(interface{}) dbr.Builder
		Lte func(interface{}) dbr.Builder
	}{
		Name: {{.SQLName | quote}},
		IsPk: {{.IsPk}},
		IsAI: {{.IsPk}},
		Eq: func(v interface{}) dbr.Builder {
			return dbr.Eq({{.SQLName | quote}}, v)
		},
		In: func(v ...interface{}) dbr.Builder {
			return dbr.Eq({{.SQLName | quote}}, v)
		},
		Gt: func(v interface{}) dbr.Builder {
			return dbr.Gt({{.SQLName | quote}}, v)
		},
		Gte: func(v interface{}) dbr.Builder {
			return dbr.Gte({{.SQLName | quote}}, v)
		},
		Lt: func(v interface{}) dbr.Builder {
			return dbr.Lt({{.SQLName | quote}}, v)
		},
		Lte: func(v interface{}) dbr.Builder {
			return dbr.Lte({{.SQLName | quote}}, v)
		},
	}
{{end}}
{{ if eq .SQLType "TEXT"}}
	{{.Name}}: struct {
		Name string
		Eq func(interface{}) dbr.Builder
		In func(...interface{}) dbr.Builder
		Like func(interface{}) dbr.Builder
	}{
		Name: {{.SQLName | quote}},
		Eq: func(v interface{}) dbr.Builder {
			return dbr.Eq({{.SQLName | quote}}, v)
		},
		In: func(v ...interface{}) dbr.Builder {
			return dbr.Eq({{.SQLName | quote}}, v)
		},
		Like: func(v interface{}) dbr.Builder {
			return builder.Like({{.SQLName | quote}}, v)
		},
	}
{{end}}
`))

}
