package gen

import (
	"bytes"
	"strings"

	"github.com/mh-cbon/jedi/model"
)

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

func itemGoType(s string) string {
	s = strings.TrimLeft(s, "*")
	if strings.HasPrefix(s, "[]") {
		s = s[2:]
	}
	s = strings.TrimLeft(s, "*")
	return s
}

var funcs = map[string]interface{}{
	"quote": func(a string) string {
		return "`" + strings.Replace(a, "`", "\\`", -1) + "`"
	},
	"notEmpty": func(a ...string) bool {
		for _, b := range a {
			if len(b) == 0 {
				return false
			}
		}
		return true
	},
	"empty": func(a ...string) bool {
		for _, b := range a {
			if len(b) > 0 {
				return false
			}
		}
		return true
	},
	"trim": func(a string) string {
		return strings.TrimSpace(a)
	},
	"ucfirst": func(a string) string {
		return strings.Title(a)
	},
	"trail": func(w, a string) string {
		if a != "" {
			a += w
		}
		return a
	},
	"itemGoType": itemGoType,
	"findStruct": findStruct,
	"findProp":   findProp,
	"notPk": func(fields []*model.Field) []*model.Field {
		var ret []*model.Field
		for _, c := range fields {
			if c.IsPk == false {
				ret = append(ret, c)
			}
		}
		return ret
	},
	"notAI": func(fields []*model.Field) []*model.Field {
		var ret []*model.Field
		for _, c := range fields {
			if !c.IsAI {
				ret = append(ret, c)
			}
		}
		return ret
	},
	"isAI": func(fields []*model.Field) []*model.Field {
		var ret []*model.Field
		for _, c := range fields {
			if c.IsAI {
				ret = append(ret, c)
			}
		}
		return ret
	},
	"isPk": func(fields []*model.Field) []*model.Field {
		var ret []*model.Field
		for _, c := range fields {
			if c.IsPk == true {
				ret = append(ret, c)
			}
		}
		return ret
	},
	"isHasOne": func(fields []*model.Field) []*model.Field {
		var ret []*model.Field
		for _, c := range fields {
			if c.HasOne != "" {
				ret = append(ret, c)
			}
		}
		return ret
	},
	"isHasMany2Many": func(fields []*model.Field) []*model.Field {
		var ret []*model.Field
		for _, c := range fields {
			if c.IsPk == true {
				ret = append(ret, c)
			}
		}
		return ret
	},
	"isHasMany2One": func(fields []*model.Field) []*model.Field {
		var ret []*model.Field
		for _, c := range fields {
			if c.IsPk == true {
				ret = append(ret, c)
			}
		}
		return ret
	},
	"withSQLType": func(fields []*model.Field) []*model.Field {
		var ret []*model.Field
		for _, c := range fields {
			if c.SQLType != "" {
				ret = append(ret, c)
			}
		}
		return ret
	},
	"withGoName": func(fields []*model.Field) []*model.Field {
		var ret []*model.Field
		for _, c := range fields {
			if c.Name != "" {
				ret = append(ret, c)
			}
		}
		return ret
	},
	"hasNonPkField": func(fields []*model.Field) bool {
		for _, c := range fields {
			if !c.IsPk {
				return true
			}
		}
		return false
	},
	"hasPkField": func(fields []*model.Field) bool {
		for _, c := range fields {
			if c.IsPk {
				return true
			}
		}
		return false
	},
	"getPkFieldName": func(fields []*model.Field) string {
		for _, c := range fields {
			if c.IsPk {
				return c.Name
			}
		}
		return ""
	},
	"getPkFieldGoType": func(fields []*model.Field) string {
		for _, c := range fields {
			if c.IsPk {
				return c.GoType
			}
		}
		return ""
	},
	"getPkSQLType": func(fields []*model.Field) string {
		for _, c := range fields {
			if c.IsPk {
				return c.SQLType
			}
		}
		return ""
	},
	"getPkSQLName": func(fields []*model.Field) string {
		for _, c := range fields {
			if c.IsPk {
				return c.SQLName
			}
		}
		return ""
	},
	"AsMethodParams": func(fields []*model.Field) string {
		ret := []string{}
		for _, c := range fields {
			if c.IsPk {
				ret = append(ret, c.SQLName+" "+c.GoType)
			}
		}
		return strings.Join(ret, ",")
	},
	"AsVariadicMethodParams": func(fields []*model.Field) string {
		ret := []string{}
		for _, c := range fields {
			if c.IsPk {
				ret = append(ret, c.SQLName+" ..."+c.GoType)
			}
		}
		return strings.Join(ret, ",")
	},
	"createTable": func(driver string, s model.Struct) string {
		cols := ""
		hasPk := []string{}
		for _, f := range s.Fields {
			if f.SQLType != "" {
				cols += f.SQLName
				cols += " " + f.SQLType
				if f.IsPk {
					if driver == "sqlite3" {
						if f.IsAI && f.IsPk {
							cols += " PRIMARY KEY"
						} else {
							hasPk = append(hasPk, f.SQLName)
						}
						if f.IsAI {
							cols += " AUTOINCREMENT"
						}
					} else if driver == "mysql" {
						hasPk = append(hasPk, f.SQLName)
						cols += " NOT NULL"
						if f.IsAI {
							cols += " AUTO_INCREMENT"
						}
					}
				}
				cols += ",\n"
			}
		}
		if driver == "mysql" && len(hasPk) > 0 {
			cols += "PRIMARY KEY (" + strings.Join(hasPk, ",") + ") ,\n"
		} else if driver == "sqlite3" && len(hasPk) > 0 {
			cols += "PRIMARY KEY (" + strings.Join(hasPk, ",") + ") ,\n"
		}
		if len(cols) > 1 {
			cols = cols[:len(cols)-2]
		}
		cols += "\n"
		return "CREATE TABLE IF NOT EXISTS " + s.SQLName + " (\n" + cols + "\n)"
	},
	"dropTable": func(driver string, s model.Struct) string {
		return "DROP TABLE IF EXISTS " + s.SQLName
	},
	"createView": func(driver string, s model.Struct) string {
		return "CREATE VIEW IF NOT EXISTS " + s.SQLName + " AS " + s.SQLViewSelect + ""
	},
	"dropView": func(driver string, s model.Struct) string {
		return "DROP VIEW IF EXISTS " + s.SQLName
	},
	"printHelperDecl": func(field model.Field) (string, error) {
		var out bytes.Buffer
		err := HelpersDecl.Execute(&out, field)
		return (&out).String(), err
	},
	"printHelperBody": func(field model.Field) (string, error) {
		var out bytes.Buffer
		err := HelpersBody.Execute(&out, field)
		return (&out).String(), err
	},
}
