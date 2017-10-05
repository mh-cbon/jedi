package model

import (
	"strings"
)

// Struct represents struct properties.
type Struct struct {
	SQLName        string
	SQLTableCreate string
	SQLTableDrop   string
	SQLViewSelect  string
	SQLViewCreate  string
	SQLViewDrop    string
	Name           string
	IsAutoGoType   bool
	Fields         []*Field
}

//Pks returns only primary keys
func (s *Struct) Pks() []*Field {
	ret := []*Field{}
	for _, a := range s.Fields {
		if a.IsPk {
			ret = append(ret, a)
		}
	}
	return ret
}

//GetFieldByName finds the property with given name
func (s *Struct) GetFieldByName(n string) *Field {
	for _, a := range s.Fields {
		if a.Name == n {
			return a
		}
	}
	return nil
}

// Field represents computed properties from struct tags.
type Field struct {
	Name       string
	SQLName    string
	SQLType    string
	GoType     string
	On         string
	HasMany    string
	HasOne     string
	IsPk       bool
	IsAI       bool
	IsNullable bool
}

// IsStar returns true if the go type is pointer
func (f *Field) IsStar() bool {
	return strings.HasPrefix(f.GoType, "*")
}
