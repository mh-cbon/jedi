package builder

import (
	"github.com/gocraft/dbr"
	"github.com/mh-cbon/jedi/runtime"
)

//SelectBuilder ...
type SelectBuilder struct {
	*dbr.SelectBuilder
}

//ReadInt tries to load an int
func (s *SelectBuilder) ReadInt() (int, error) {
	var ret int64
	return int(ret), s.LoadValue(&ret)
}

//ReadInt64 tries to load an int64
func (s *SelectBuilder) ReadInt64() (int64, error) {
	var ret int64
	return ret, s.LoadValue(&ret)
}

//Build builds the sql string into given buffer using current dialect
func (s *SelectBuilder) Build(b dbr.Buffer) error {
	return s.SelectBuilder.Build(runtime.GetDialect(), b)
}

//String returns the sql string for current dialect. It returns empty string if the build returns an error.
func (s *SelectBuilder) String() string {
	b := dbr.NewBuffer()
	if err := s.Build(b); err != nil {
		return ""
	}
	return b.String()
}

//DeleteBuilder ...
type DeleteBuilder struct {
	*dbr.DeleteBuilder
}

//String returns the sql string for current dialect. It returns empty string if the build returns an error.
func (d *DeleteBuilder) String() string {
	b := dbr.NewBuffer()
	if err := d.Build(b); err != nil {
		return ""
	}
	return b.String()
}

//Build builds the sql string into given buffer using current dialect
func (d *DeleteBuilder) Build(b dbr.Buffer) error {
	return d.DeleteBuilder.Build(runtime.GetDialect(), b)
}

// Where add conditions
func (d *DeleteBuilder) Where(query interface{}, value ...interface{}) *DeleteBuilder {
	d.DeleteBuilder.Where(query, value...)
	return d
}

//UpdateBuilder ...
type UpdateBuilder struct {
	*dbr.UpdateBuilder
}

//String returns the sql string for current dialect. It returns empty string if the build returns an error.
func (d *UpdateBuilder) String() string {
	b := dbr.NewBuffer()
	if err := d.Build(b); err != nil {
		return ""
	}
	return b.String()
}

//Build builds the sql string into given buffer using current dialect
func (d *UpdateBuilder) Build(b dbr.Buffer) error {
	return d.UpdateBuilder.Build(runtime.GetDialect(), b)
}
