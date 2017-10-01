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
	// TableVar  string
}

func (s *Struct) Pks() []*Field {
	ret := []*Field{}
	for _, a := range s.Fields {
		if a.IsPk {
			ret = append(ret, a)
		}
	}
	return ret
}
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

func (f *Field) IsStar() bool {
	return strings.HasPrefix(f.GoType, "*")
}
