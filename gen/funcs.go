package gen

import (
	"fmt"
	"log"
	"strings"

	"github.com/mh-cbon/jedi/drivers"
	"github.com/mh-cbon/jedi/model"
)

func findLocals(all []*model.Struct, s *model.Struct, f *model.Field) []*model.Field {
	var ret []*model.Field
	if f.HasOne != "" {
		foreign := model.FindStruct(all, f.HasOne)
		for _, pk := range foreign.Pks() {
			local := s.GetFieldByName(strings.Title(f.Name) + "" + pk.Name)
			if local != nil {
				ret = append(ret, local)
			}
		}
	}
	return ret
}

func itemGoType(s string) string {
	s = strings.TrimLeft(s, "*")
	if strings.HasPrefix(s, "[]") {
		s = s[2:]
	}
	s = strings.TrimLeft(s, "*")
	return s
}

var timeString = "time.Time"
var timePString = "*time.Time"

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
	"findStruct": model.FindStruct,
	"findProp":   model.FindProp,
	"findLocals": findLocals,
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
	"notExported": func(fields []*model.Field) []*model.Field {
		var ret []*model.Field
		for _, c := range fields {
			if !c.Exported() {
				ret = append(ret, c)
			}
		}
		return ret
	},
	"isExported": func(fields []*model.Field) []*model.Field {
		var ret []*model.Field
		for _, c := range fields {
			if c.Exported() {
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
	"isSQLText": func(fields []*model.Field) []*model.Field {
		var ret []*model.Field
		for _, c := range fields {
			if c.SQLType == "TEXT" {
				ret = append(ret, c)
			}
		}
		return ret
	},
	"isRel": func(fields []*model.Field) []*model.Field {
		var ret []*model.Field
		for _, c := range fields {
			if c.HasOne != "" || c.HasMany != "" {
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
	"isHasMany2Many": func(all []*model.Struct, fields []*model.Field) []*model.Field {
		var ret []*model.Field
		for _, c := range fields {
			if c.HasMany != "" {
				foreign := model.FindProp(all, c.HasMany)
				if foreign == nil {
					ret = append(ret, c)
				} else if foreign.HasMany != "" {
					ret = append(ret, c)
				}
			}
		}
		return ret
	},
	"isHasMany2One": func(all []*model.Struct, fields []*model.Field) []*model.Field {
		var ret []*model.Field
		for _, c := range fields {
			if c.HasMany != "" {
				foreign := model.FindProp(all, c.HasMany)
				if foreign != nil && foreign.HasOne != "" {
					ret = append(ret, c)
				}
			}
		}
		return ret
	},
	"getHasOne": func(all []*model.Struct, aboutStruct *model.Struct, aboutField *model.Field) *model.HasOne {
		var ret *model.HasOne

		foreign := model.FindStruct(all, aboutField.HasOne)
		if foreign == nil {
			return ret
		}

		ret = &model.HasOne{
			Local:   aboutStruct,
			Foreign: foreign,
		}

		for _, pk := range foreign.Pks() {
			localName := strings.Title(aboutField.Name) + pk.Name
			localImported := aboutStruct.GetFieldByName(localName)
			if localImported == nil {
				log.Printf("jedi: %q is missing a field %v.%v", aboutStruct.Name, aboutStruct.Name, localName)
				continue
			}
			j := &model.JoinFields{
				LocalField:   aboutStruct.GetFieldByName(strings.Title(aboutField.Name) + pk.Name),
				ForeignField: pk,
			}
			ret.Fields = append(ret.Fields, j)
		}

		return ret
	},
	"getMany2One": func(all []*model.Struct, aboutStruct *model.Struct, aboutField *model.Field) *model.Many2One {
		var ret *model.Many2One

		foreign := model.FindStruct(all, aboutField.HasMany)
		if foreign == nil {
			return ret
		}
		foreignProp := model.FindProp(all, aboutField.HasMany)
		if foreignProp == nil {
			return ret
		}

		ret = &model.Many2One{
			Local:        aboutStruct,
			Foreign:      foreign,
			LocalField:   aboutField,
			ForeignField: foreignProp,
		}

		for _, pk := range aboutStruct.Pks() {
			j := &model.JoinFields{
				LocalField:   pk,
				ForeignField: foreign.GetFieldByName(strings.Title(foreignProp.Name) + pk.Name),
			}
			ret.Fields = append(ret.Fields, j)
		}

		return ret
	},
	"getMany2Many": func(all []*model.Struct, aboutStruct *model.Struct, aboutField *model.Field) *model.Many2Many {
		var ret *model.Many2Many

		foreign := model.FindStruct(all, aboutField.HasMany)
		if foreign == nil {
			return ret
		}
		foreignProp := model.FindProp(all, aboutField.HasMany)
		if foreignProp == nil {
			return ret
		}

		n := model.HasMany2ManyGoTypeName(aboutStruct, foreign, aboutField, foreignProp)
		midType := model.FindStruct(all, n)
		ret = &model.Many2Many{
			Local:   aboutStruct,
			Foreign: foreign,
			Middle:  midType,
		}

		for _, pk := range ret.Local.Pks() {
			foreiName := strings.Title(aboutStruct.Name) + pk.Name
			foreiField := ret.Middle.GetFieldByName(foreiName)
			j := &model.JoinFields{
				LocalField:   pk,
				ForeignField: foreiField,
			}
			ret.LMFields = append(ret.LMFields, j)
		}

		for _, pk := range ret.Foreign.Pks() {
			j := &model.JoinFields{
				LocalField:   pk,
				ForeignField: ret.Middle.GetFieldByName(strings.Title(foreign.Name) + pk.Name),
			}
			ret.FMFields = append(ret.FMFields, j)
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
	"withoutSQLType": func(fields []*model.Field) []*model.Field {
		var ret []*model.Field
		for _, c := range fields {
			if c.SQLType == "" {
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
	"dateTypes": func(fields []*model.Field) []*model.Field {
		var ret []*model.Field
		for _, c := range fields {
			if c.GoType == timeString || c.GoType == timePString {
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
				ret = append(ret, c.Name+" "+c.GoType)
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
	"createIndex": func(driver string, index *model.Index, s model.Struct) string {
		q := fmt.Sprintf("CREATE INDEX %v ", index.Name)
		if index.Unique {
			q = fmt.Sprintf("CREATE UNIQUE INDEX %v ", index.Name)
		}
		q += fmt.Sprintf("ON %v ", s.SQLName)
		t := []string{}
		for _, c := range index.Fields {
			f := s.GetFieldByName(c)
			if strings.HasSuffix(f.GoType, "string") && driver == drivers.Mysql {
				t = append(t, f.SQLName+"(255)")
			} else {
				t = append(t, f.SQLName)
			}
		}
		q += fmt.Sprintf("(%v) ", strings.Join(t, ","))
		return q
	},
	"createTable": func(driver string, s model.Struct) string {
		cols := ""
		hasPk := []string{}
		for _, f := range s.Fields {
			if f.SQLType != "" {
				cols += f.SQLName
				if f.IsPk && driver == drivers.Mysql && f.SQLType == "TEXT" {
					cols += " VARCHAR(255)"
				} else if (f.GoType == timeString || f.GoType == timePString) && driver == drivers.Pgsql {
					cols += " timestamp(6)"
				} else if (f.GoType == timeString || f.GoType == timePString) && driver == drivers.Mysql {
					cols += " datetime(6)"
				} else if f.IsAI && f.SQLType == "INTEGER" && driver == drivers.Pgsql {
					cols += " SERIAL"
				} else {
					cols += " " + f.SQLType
				}
				if !f.IsPk && f.IsStar() {
					cols += " NULL"
				}
				if f.IsPk {
					if driver == drivers.Sqlite {
						if f.IsAI && f.IsPk {
							cols += " PRIMARY KEY"
						} else {
							hasPk = append(hasPk, f.SQLName)
						}
						if f.IsAI {
							cols += " AUTOINCREMENT"
						}
					} else if driver == drivers.Mysql {
						hasPk = append(hasPk, f.SQLName)
						cols += " NOT NULL"
						if f.IsAI {
							cols += " AUTO_INCREMENT"
						}
					} else if driver == drivers.Pgsql {
						// cols += " NOT NULL"
						if f.IsAI {
							cols += " PRIMARY KEY"
						} else {
							hasPk = append(hasPk, f.SQLName)
						}
					}
				}
				cols += ",\n"
			}
		}
		if len(hasPk) > 0 {
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
		if driver == drivers.Mysql || driver == drivers.Pgsql {
			return "CREATE OR REPLACE VIEW " + s.SQLName + " AS " + s.SQLViewSelect + ""
		}
		return "CREATE VIEW IF NOT EXISTS " + s.SQLName + " AS " + s.SQLViewSelect + ""
	},
	"dropView": func(driver string, s model.Struct) string {
		return "DROP VIEW IF EXISTS " + s.SQLName
	},
}
