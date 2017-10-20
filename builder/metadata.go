package builder

import (
	"fmt"

	"github.com/gocraft/dbr"
	"github.com/mh-cbon/jedi/runtime"
)

//MetaProvider is an interface to mix Value and Rel Property meta.
type MetaProvider interface {
	IsRel() bool
	IsPk() bool
	IsAI() bool
	Type() string
	GoType() string
	Name() string
	GoName() string
}

// PropertyMeta is the common information about any property.
type PropertyMeta struct {
	sqlName string
	sqlType string
	goName  string
	goType  string
	isPk    bool
	isAI    bool
	hasOne  string
	hasMany string
	On      string
	RelType string
}

// IsRel returns true when thep roperty is a relation ship.
func (i PropertyMeta) IsRel() bool { return i.hasOne != "" || i.hasMany != "" }

// IsPk returns true for a primary key
func (i PropertyMeta) IsPk() bool { return i.isPk }

// IsAI returns true for an autoincrement
func (i PropertyMeta) IsAI() bool { return i.isAI }

// Name returns the sql name (it is empty for a relationship)
func (i PropertyMeta) Name() string { return i.sqlName }

// Type returns the sql type (it is empty for a relationship)
func (i PropertyMeta) Type() string { return i.sqlType }

// GoName returns the go name
func (i PropertyMeta) GoName() string { return i.goName }

// GoType returns the go type
func (i PropertyMeta) GoType() string { return i.goType }

//NewValueMeta returns a new meta for a value property
func NewValueMeta(
	sqlName, sqlType,
	goName, goType string,
	isPk, isAI bool,
) ValuePropertyMeta {
	return ValuePropertyMeta{
		PropertyMeta: PropertyMeta{
			sqlName: sqlName,
			sqlType: sqlType,
			goName:  goName,
			goType:  goType,
			isPk:    isPk,
			isAI:    isAI,
		},
	}
}

// ValuePropertyMeta represents a value property.
type ValuePropertyMeta struct {
	PropertyMeta
	TableAlias string
}

//Aliased returns [table.]column if tablealias is not empty.
func (i ValuePropertyMeta) Aliased() string {
	ret := i.sqlName
	if i.TableAlias != "" {
		ret = fmt.Sprintf("%v.%v", i.TableAlias, ret)
	}
	return ret
}

// IsNull is `col IS NULL`.
func (i ValuePropertyMeta) IsNull() dbr.Builder {
	return runtime.IsNull(i.Aliased())
}

// IsNotNull is `col IS NOT NULL`.
func (i ValuePropertyMeta) IsNotNull() dbr.Builder {
	return runtime.IsNotNull(i.Aliased())
}

// Eq is `col = value`.
func (i ValuePropertyMeta) Eq(v interface{}) dbr.Builder {
	return dbr.Eq(i.Aliased(), v)
}

// In is `col IN (...values)`.
func (i ValuePropertyMeta) In(v ...interface{}) dbr.Builder {
	return dbr.Eq(i.Aliased(), v)
}

// Gt is `col > value`.
func (i ValuePropertyMeta) Gt(v interface{}) dbr.Builder {
	return dbr.Gt(i.Aliased(), v)
}

// Gte is `col >= value`.
func (i ValuePropertyMeta) Gte(v interface{}) dbr.Builder {
	return dbr.Gte(i.Aliased(), v)
}

// Lt is `col < value`.
func (i ValuePropertyMeta) Lt(v interface{}) dbr.Builder {
	return dbr.Lt(i.Aliased(), v)
}

// Lte is `col <= value`.
func (i ValuePropertyMeta) Lte(v interface{}) dbr.Builder {
	return dbr.Lte(i.Aliased(), v)
}

// Like is `col LIKE %value%`.
func (i ValuePropertyMeta) Like(v interface{}) dbr.Builder {
	return runtime.Like(i.Aliased(), v)
}

// RelPropertyMeta represents a relationship property.
type RelPropertyMeta struct {
	PropertyMeta
}

//NewRelMeta returns a new meta for a relationship property
func NewRelMeta(
	goName, goType string,
	hasOne, hasMany, On string,
	RelType string,
) RelPropertyMeta {
	return RelPropertyMeta{
		PropertyMeta: PropertyMeta{
			goName:  goName,
			goType:  goType,
			hasOne:  hasOne,
			hasMany: hasMany,
			On:      On,
		},
	}
}
