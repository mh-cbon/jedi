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
	Indexes        []*Index
	InsertHook     bool
	UpdateHook     bool
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

//GetIndexByName finds the index with given name
func (s *Struct) GetIndexByName(n string) *Index {
	for _, a := range s.Indexes {
		if a.Name == n {
			return a
		}
	}
	return nil
}

//IsIndexed tests if given property name is part of an index.
func (s *Struct) IsIndexed(n string) bool {
	for _, a := range s.Indexes {
		for _, c := range a.Fields {
			if c == n {
				return true
			}
		}
	}
	return false
}

// Field represents computed properties from struct tags.
type Field struct {
	Name        string
	SQLName     string
	SQLType     string
	GoType      string
	On          string   // @on for has many 2 many, define the middle table
	HasMany     string   // @has_many for has many, define the target property
	HasOne      string   // @has_one for has one, define the target property
	RelType     string   // @has_one / @has_many / has_many2many
	IsPk        bool     // @pk the field is part of the primary key
	IsAI        bool     // @pk the field is auto incremented.
	IsNullable  bool     // the field is nullable because its go type is nilable.
	UTC         bool     // @utc for datetime it will automatically turn it into UTC before Insert / Update.
	LastUpdated bool     // @last_updated for datetime, such field is added as a condition of the update query and always update its value to NOW()
	Insert      bool     // @insert for datetime, such field is automatically set to NOW() on insert.
	Index       []string // The indexes name the prop is member.
	Unique      []string // The unique index name the prop is member.
}

// IsStar returns true if the go type is pointer
func (f *Field) IsStar() bool {
	return strings.HasPrefix(f.GoType, "*")
}

// Exported returns true if the proeprty is lowercase
func (f *Field) Exported() bool {
	return f.Name[0] >= 'A' && f.Name[0] <= 'Z' //todo: std api has better.
}

// Index of a table
type Index struct {
	Unique bool
	Name   string
	Fields []string //The go properties name
}
