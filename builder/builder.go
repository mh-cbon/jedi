package builder

import (
	"github.com/gocraft/dbr"
)

type SelectBuilder struct {
	*dbr.SelectBuilder
}

func (b *SelectBuilder) ReadInt() (int, error) {
	var ret int64
	return int(ret), b.LoadValue(&ret)
}
func (b *SelectBuilder) ReadInt64() (int64, error) {
	var ret int64
	return ret, b.LoadValue(&ret)
}

type DeleteBuilder struct {
	*dbr.DeleteBuilder
}

func (c *DeleteBuilder) Where(query interface{}, value ...interface{}) *DeleteBuilder {
	c.DeleteBuilder.Where(query, value...)
	return c
}
