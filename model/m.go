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
	Name        string
	SQLName     string
	SQLType     string
	GoType      string
	On          string // for has many 2 many, define the middle table
	HasMany     string // for has many, define the target property
	HasOne      string // for has one, define the target property
	IsPk        bool   // the field is part of the primary key
	IsAI        bool   // the field is auto incremented.
	IsNullable  bool   // the field is nullable because its go type is nilable.
	UTC         bool   // for datetime it will automatically turn it into UTC before Insert / Update.
	LastUpdated bool   // for datetime, such field is added as a condition of the update query and always update its value to NOW()
}

// IsStar returns true if the go type is pointer
func (f *Field) IsStar() bool {
	return strings.HasPrefix(f.GoType, "*")
}

// Unexported returns true if the proeprty is lowercase
func (f *Field) Unexported() bool {
	return f.Name[0] >= 'a' && f.Name[0] <= 'z'
}
