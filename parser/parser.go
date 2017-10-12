package parser

import (
	"bytes"
	"fmt"
	"go/ast"
	"go/parser"
	"go/token"
	"log"
	"reflect"
	"regexp"
	"strings"
	"unicode"

	"github.com/mh-cbon/jedi/model"
)

var jediComment = regexp.MustCompile(`jedi:([0-9A-Za-z_\.]+)?`)
var anyComment = regexp.MustCompile(`^([A-Za-z_\.]+):(.+)`)
var viewComment = regexp.MustCompile(`view:(.+)`)
var viewFullComment = regexp.MustCompile(`viewfull:(.+)`)

// search for a multiline jedi: prefixed comment, where jedi: is a configurable prefix.
// The comment ends at empty line, or a new prefix is found.
func getMLComment(from string, prefix string) string {
	rec := false
	var res string
	for _, l := range strings.Split(from, "\n") {
		if !rec && strings.HasPrefix(l, prefix) {
			rec = true
			l = strings.TrimLeft(l[len(prefix):], " ")
		} else if anyComment.MatchString(l) {
			rec = false
		} else if strings.TrimSpace(l) == "" {
			rec = false
		}

		if rec {
			res += l + "\n"
		}
	}
	return res
}

var integerSQLString = "INTEGER"

// Parse given file returns the jedi structs found.
func Parse(file string) ([]*model.Struct, error) {
	// parse file
	fset := token.NewFileSet()
	fileNode, err := parser.ParseFile(fset, file, nil, parser.ParseComments)
	if err != nil {
		return nil, err
	}

	// consider only top-level struct type declarations with magic comment
	var res []*model.Struct
	for _, decl := range fileNode.Decls {
		gd, ok := decl.(*ast.GenDecl)
		if !ok {
			continue
		}
		for _, spec := range gd.Specs {
			ts, ok := spec.(*ast.TypeSpec)
			if !ok {
				continue
			}

			// magic comment may be attached to "type Foo struct" (TypeSpec)
			// or to "type (" (GenDecl)
			doc := ts.Doc
			if doc == nil && len(gd.Specs) == 1 {
				doc = gd.Doc
			}
			if doc == nil {
				continue
			}

			var table string
			sm := jediComment.FindStringSubmatch(doc.Text())
			if len(sm) != 2 {
				continue
			}

			if sm[1] == "" {
				table = camelCaseToSnakeCase(ts.Name.String())
			} else {
				table = sm[1]
			}

			var viewSelect string
			var viewCreate string
			var viewDrop string
			var tableCreate string
			var tableDrop string
			viewSelect = getMLComment(doc.Text(), "view-select:")
			viewCreate = getMLComment(doc.Text(), "view-create:")
			viewDrop = getMLComment(doc.Text(), "view-drop:")
			tableCreate = getMLComment(doc.Text(), "table-create:")
			tableDrop = getMLComment(doc.Text(), "table-drop:")

			// log.Printf("jedi:%v table:%q viewSelect:%v viewCreate:%v viewDrop:%v tableCreate:%v tableDrop:%v\n",
			// 	ts.Name.String(), table,
			// 	viewSelect != "", viewCreate != "", viewDrop != "", tableCreate != "", tableDrop != "")

			str, ok := ts.Type.(*ast.StructType)
			if !ok {
				continue
			}
			if str.Incomplete {
				continue
			}
			// ast.Print(fset, ts)
			s, err := parseStructTypeSpec(ts, str)
			if err != nil {
				return nil, err
			}
			s.SQLName = table
			s.SQLViewSelect = viewSelect
			s.SQLViewCreate = viewCreate
			s.SQLViewDrop = viewDrop
			s.SQLTableCreate = tableCreate
			s.SQLTableDrop = tableDrop
			res = append(res, s)
		}
	}

	// setup auto AI
	for _, r := range res {
		if len(r.Pks()) == 1 && r.Pks()[0].SQLType == integerSQLString {
			r.Pks()[0].IsAI = true
		}
	}

	// install has_one properties
	for _, r := range res {
		for _, f := range r.Fields {
			if f.HasOne != "" {
				foreign := findStruct(res, f.HasOne)
				if foreign == nil {
					log.Printf("jedi: has_one not found %#v in %#v\n", f.HasOne, r.Name)
					continue
				}
				for _, pk := range foreign.Pks() {
					rRequiredField := strings.Title(f.Name) + pk.Name
					if g := r.GetFieldByName(rRequiredField); g == nil {
						log.Printf("jedi: %v is missing a field %v.%v because it @has_one=%v\n",
							r.Name, r.Name, rRequiredField, f.HasOne)
					} else if g.IsStar() != f.IsStar() {
						log.Printf("jedi: %v.%v / %v.%v must be both go pointer or both value\n",
							r.Name, g.Name, r.Name, f.Name)
					}
				}
			}
		}
	}

	// install has_many properties
	todos := map[string]*model.Struct{}
	for _, r := range res {
		for _, f := range r.Fields {
			if f.HasMany != "" {

				if strings.Index(f.HasMany, ".") < 1 {
					log.Printf("jedi: %v.%v must target a reverse property in @has_many=%q\n",
						r.Name, f.Name, f.HasMany)
					// clean the field so it is not generated becasue invalid.
					f.HasMany = ""
					continue
				}

				foreign := findStruct(res, f.HasMany)
				if foreign == nil {
					log.Printf("jedi: %v.%v target type not found in @has_many=%q\n", r.Name, f.Name, f.HasMany)
					// clean the field so it is not generated becasue invalid.
					f.HasMany = ""
					continue
				}

				foreignProp := findProp(res, f.HasMany)
				if foreignProp == nil {
					log.Printf("jedi: %v.%v target property not found in @has_many=%q\n", r.Name, f.Name, f.HasMany)
					// clean the field so it is not generated becasue invalid.
					f.HasMany = ""
					continue
				}

				if foreignProp.HasOne != "" {
					// looking for many2many relations only
					continue
				}

				if foreignProp.HasMany == "" {
					log.Printf("jedi: %v.%v: is missing a reverse property: @has_many=%q\n",
						foreign.Name, foreignProp.Name,
						r.Name+"."+f.Name,
					)
					// clean the field so it is not generated becasue invalid.
					f.HasMany = ""
					continue
				}

				IsAutoGoType := foreignProp.On == "" && f.On == ""
				if !IsAutoGoType && f.On != foreignProp.On {
					log.Printf("jedi: %v.%v(@has_many=%q @on=%q) / %v.%v(@has_many=%q @on=%q): values must match\n",
						r.Name, f.Name, f.HasMany, f.On,
						foreign.Name, foreignProp.Name, foreignProp.HasMany, foreignProp.On)
				}

				gType := model.HasMany2ManyGoTypeName(r, foreign, f, foreignProp)
				if _, ok := todos[gType]; ok == false {
					if IsAutoGoType {
						todos[gType] = &model.Struct{
							SQLName:      model.HasMany2ManyTableName(r, foreign, f, foreignProp),
							Name:         gType,
							IsAutoGoType: IsAutoGoType,
						}
						hasManyAddColumns(todos[gType], r, foreign, f, foreignProp)
					}
				}
			}
		}
	}
	for _, t := range todos {
		res = append(res, t)
	}

	return res, nil
}

func itemGoType(s string) string {
	s = strings.TrimLeft(s, "*")
	if strings.HasPrefix(s, "[]") {
		s = s[2:]
	}
	s = strings.TrimLeft(s, "*")
	return s
}

func hasManyAddColumns(join, left, right *model.Struct, leftProp, rightProp *model.Field) {
	for _, r := range right.Pks() {
		gType := itemGoType(right.Name)
		sName := fmt.Sprintf("%v_%v", camelCaseToSnakeCase(gType), strings.ToLower(r.Name))
		gName := fmt.Sprintf("%v%v", gType, r.Name)
		join.Fields = append(join.Fields, &model.Field{
			Name:    gName,
			GoType:  r.GoType,
			SQLName: sName,
			SQLType: r.SQLType,
			IsPk:    true,
		})
	}
	for _, l := range left.Pks() {
		gType := itemGoType(left.Name)
		sName := fmt.Sprintf("%v_%v", camelCaseToSnakeCase(gType), strings.ToLower(l.Name))
		gName := fmt.Sprintf("%v%v", gType, l.Name)
		join.Fields = append(join.Fields, &model.Field{
			Name:    gName,
			GoType:  l.GoType,
			SQLName: sName,
			SQLType: l.SQLType,
			IsPk:    true,
		})
	}
}

func findStruct(all []*model.Struct, n string) *model.Struct {
	if strings.Index(n, ".") > -1 {
		n = strings.Split(n, ".")[0]
	}
	for _, a := range all {
		if a.Name == n {
			return a
		}
	}
	return nil
}

func findProp(all []*model.Struct, n string) *model.Field {
	var nprop string
	if strings.Index(n, ".") == -1 {
		return nil
	}
	nprop = strings.Split(n, ".")[1]
	n = strings.Split(n, ".")[0]
	for _, a := range all {
		if a.Name == n {
			for _, f := range a.Fields {
				if f.Name == nprop {
					return f
				}
			}
		}
	}
	return nil
}

var trueString = "true"
var byteString = "byte"
var uint8String = "uint8"

func parseStructTypeSpec(ts *ast.TypeSpec, str *ast.StructType) (*model.Struct, error) {
	res := &model.Struct{
		Name: ts.Name.Name,
	}

	var n int
	for _, f := range str.Fields.List {

		var tag string
		if f.Tag != nil {
			t := f.Tag.Value
			t = t[1 : len(t)-1] // strip quotes
			tag = reflect.StructTag(t).Get("jedi")
		}

		if tag == "-" {
			continue
		}

		// check for anonymous fields
		if len(f.Names) == 0 {
			return nil, fmt.Errorf(`jedi: %s has anonymous field %s with "jedi:" tag, it is not allowed`, res.Name, f.Type)
		}
		if len(f.Names) != 1 {
			panic(fmt.Sprintf("jedi: %d names: %#v. Please report this bug.", len(f.Names), f.Names))
		}

		// check for exported name
		name := f.Names[0]

		// parse tag and type
		column, props := parseStructFieldTag(tag)
		if column == "" {
			column = camelCaseToSnakeCase(name.Name)
		}
		typ := strGoType(f.Type)
		styp := strSQLType(f.Type)
		isPk := props["pk"] == trueString
		if props["has_many"] == trueString {
			props["has_many"] = strGoItemType(f.Type)
		}
		if props["has_one"] == trueString {
			props["has_one"] = strGoItemType(f.Type)
		}
		res.Fields = append(res.Fields, &model.Field{
			Name:        name.String(),
			GoType:      typ,
			SQLName:     column,
			SQLType:     styp,
			IsPk:        isPk,
			IsAI:        isPk && styp == integerSQLString,
			On:          props["on"],
			HasMany:     props["has_many"],
			HasOne:      props["has_one"],
			IsNullable:  strSQLNullable(f.Type),
			UTC:         props["utc"] == trueString,
			LastUpdated: props["last_updated"] == trueString,
		})
		n++
	}

	if len(res.Fields) == 0 {
		return nil, fmt.Errorf(`jedi: %s has no fields with "jedi:" tag, it is not allowed`, res.Name)
	}

	return res, nil
}

// parseStructFieldTag is used by both file and runtime parsers
func parseStructFieldTag(tag string) (sqlName string, props map[string]string) {
	props = map[string]string{
		"pk":           "false",
		"has_many":     "",
		"has_one":      "",
		"on":           "",
		"utc":          trueString,
		"last_updated": "false",
	}

	parts := strings.Split(tag, ",")
	if len(parts) == 0 || len(parts) > 2 {
		return
	}

	for _, p := range parts {
		p = strings.TrimSpace(p)
		if strings.HasPrefix(p, "@") && len(p) > 1 {
			u := strings.Split(p[1:], "=")
			if len(u) > 1 {
				props[u[0]] = u[1]
			} else {
				props[p[1:]] = trueString
			}
		} else if len(p) > 0 {
			sqlName = p
		}
	}

	return
}

func strGoItemType(x ast.Expr) string {
	switch t := x.(type) {
	case *ast.StarExpr:
		return strGoItemType(t.X)
	case *ast.SelectorExpr:
		return strGoItemType(t.X) + "." + t.Sel.String()
	case *ast.Ident:
		s := t.String()
		if s == byteString {
			return uint8String
		}
		return s
	case *ast.ArrayType:
		return strGoItemType(t.Elt)
	case *ast.BasicLit:
		return t.Value
	case nil:
		return ""
	default:
		panic(fmt.Sprintf("jedi: strGoItemType: unhandled '%s' (%#v). Please report this bug.", x, x))
	}
}
func strGoType(x ast.Expr) string {
	switch t := x.(type) {
	case *ast.StarExpr:
		return "*" + strGoType(t.X)
	case *ast.SelectorExpr:
		return strGoType(t.X) + "." + t.Sel.String()
	case *ast.Ident:
		s := t.String()
		if s == byteString {
			return uint8String
		}
		return s
	case *ast.ArrayType:
		return "[" + strGoType(t.Len) + "]" + strGoType(t.Elt)
	case *ast.BasicLit:
		return t.Value
	case nil:
		return ""
	default:
		panic(fmt.Sprintf("jedi: strGoType: unhandled '%s' (%#v). Please report this bug.", x, x))
	}
}
func strSQLType(x ast.Expr) string {
	switch t := x.(type) {
	case *ast.StarExpr:
		return strSQLType(t.X)
	case *ast.SelectorExpr:
		return "datetime"
	case *ast.Ident:
		s := t.String()
		switch s {
		case "int64", "int32", "int16", "int8", "int":
			return integerSQLString
		case "uint64", "uint32", "uint16", uint8String, byteString, "uint":
			return integerSQLString
		case "bool":
			return integerSQLString
		case "float64", "float32":
			return "FLOAT"
		case "string":
			return "TEXT"
		}
		return ""
	case *ast.BasicLit:
		return t.Value
	default:
		return ""
	}
}
func strSQLNullable(x ast.Expr) bool {
	switch x.(type) {
	case *ast.StarExpr:
		return true
	}
	return false
}

func camelCaseToSnakeCase(name string) string {
	buf := new(bytes.Buffer)

	runes := []rune(name)

	for i := 0; i < len(runes); i++ {
		buf.WriteRune(unicode.ToLower(runes[i]))
		if i != len(runes)-1 && unicode.IsUpper(runes[i+1]) &&
			(unicode.IsLower(runes[i]) || unicode.IsDigit(runes[i]) ||
				(i != len(runes)-2 && unicode.IsLower(runes[i+2]))) {
			buf.WriteRune('_')
		}
	}

	return buf.String()
}
