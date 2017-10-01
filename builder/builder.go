package builder

import (
	"fmt"

	"github.com/gocraft/dbr"
)

// type Cond struct {
// 	Query  interface{}
// 	Values []interface{}
// }
//
// func (c Cond) Apply(f interface{}) {
// 	if x, ok := f.(*dbr.SelectBuilder); ok {
// 		x.Where(c.Query, c.Values...)
// 	}
// 	if x, ok := f.(*dbr.DeleteBuilder); ok {
// 		x.Where(c.Query, c.Values...)
// 	}
// 	if x, ok := f.(*dbr.UpdateBuilder); ok {
// 		x.Where(c.Query, c.Values...)
// 	}
// }
//
// func Or(conds ...CondApplier) OrCond {
// 	return OrCond{Conds: conds}
// }
//
// type OrCond struct {
// 	Conds []CondApplier
// }
//
// func (c OrCond) Apply(f interface{}) {
// 	for _, cc := range c.Conds {
// 		cc.Apply(f)
// 	}
// }
//
// func And(conds ...CondApplier) OrCond {
// 	return OrCond{Conds: conds}
// }
//
// type AndCond struct {
// 	Conds []CondApplier
// }
//
// func (c AndCond) Apply(f interface{}) {
// 	for _, cc := range c.Conds {
// 		cc.Apply(f)
// 	}
// }
//
// type CondApplier interface {
// 	Apply(f interface{})
// }

type SelectBuilder struct {
	*dbr.SelectBuilder
}

// func (b *SelectBuilder) With(conds ...CondApplier) *SelectBuilder {
// 	for _, c := range conds {
// 		c.Apply(b.SelectBuilder)
// 	}
// 	return b
// }

// Like is `Like`.
func Like(column string, value interface{}) dbr.Builder {
	return dbr.BuildFunc(func(d dbr.Dialect, buf dbr.Buffer) error {
		buf.WriteString(d.QuoteIdent(column))
		buf.WriteString(" ")
		buf.WriteString("LIKE")
		buf.WriteString(" ")
		buf.WriteString(fmt.Sprintf("%%%v%%", value))
		return nil
	})
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

// func (b *DeleteBuilder) With(conds ...CondApplier) *DeleteBuilder {
// 	for _, c := range conds {
// 		c.Apply(b.DeleteBuilder)
// 	}
// 	return b
// }
