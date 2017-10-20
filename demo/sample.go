package main

import (
	"errors"
	"time"
)

//go:generate jedi

//Sample is whatever.
//jedi:
type Sample struct {
	ID          int64 `jedi:"@pk"`
	Name        string
	Description string
	UpdateDate  time.Time
	RemovalDate *time.Time
}

//BasicPK is a type with onlya primary interger auto increment key.
// useful to test vendors.
//jedi:
type BasicPK struct {
	ID       int64 `jedi:"@pk"`
	Whatever string
}

//BasicTypes is a type with various possible basic type properties.
//jedi:
type BasicTypes struct {
	ID       int64 `jedi:"@pk"`
	String   string
	StringP  *string
	Int      int `jedi:"intfield"` // in mysql int is reserved kw
	IntP     *int
	Int32    int32
	Int32P   *int32
	Int64    int64
	Int64P   *int64
	UInt     uint
	UIntP    *uint
	UInt32   uint32
	UInt32P  *uint32
	UInt64   uint64
	UInt64P  *uint64
	Bool     bool
	BoolP    *bool
	Float32  float32
	Float32P *float32
	Float64  float64
	Float64P *float64
}

//TextPk have a text primary key column.
//jedi:second_sample
type TextPk struct {
	Name                string `jedi:"@pk"` // PK on non integer
	Description         string
	hasManyHasOneTextPk []*HasOneTextPk  `jedi:"@has_many=HasOneTextPk.related"`
	relateds            []*HasManyTextPk `jedi:"@has_many=HasManyTextPk.relateds"`
}

//HasOneTextPk ...
//todo: write tests
//jedi:
type HasOneTextPk struct {
	ID          int64 `jedi:"@pk"`
	X           string
	related     *TextPk `jedi:"@has_one=TextPk"`
	RelatedName *string
}

//HasManyTextPk ...
//todo: write tests
//jedi:
type HasManyTextPk struct {
	ID       int64 `jedi:"@pk"`
	X        string
	relateds []*TextPk `jedi:"@has_many=TextPk.relateds"`
}

//CompositePk have a text primary key column.
//jedi:
type CompositePk struct {
	P                        string `jedi:"@pk"`
	K                        string `jedi:"@pk"`
	Description              string
	hasManyHasOneCompositePk []*HasOneCompositePk  `jedi:"@has_many=HasOneCompositePk.related"`
	relateds                 []*HasManyCompositePk `jedi:"@has_many=HasManyCompositePk.relateds"`
}

//HasOneCompositePk ...
//todo: write tests
//jedi:
type HasOneCompositePk struct {
	ID       int64 `jedi:"@pk"`
	X        string
	related  *CompositePk `jedi:"@has_one=CompositePk"`
	RelatedP *string
	RelatedK *string
}

//HasManyCompositePk ...
//todo: write tests
//jedi:
type HasManyCompositePk struct {
	ID       int64 `jedi:"@pk"`
	X        string
	relateds []*CompositePk `jedi:"@has_many=CompositePk.relateds"`
}

//DateType have dates properties.
//jedi:
type DateType struct {
	ID          int64 `jedi:"@pk"`
	T           time.Time
	TP          *time.Time
	NotUTC      *time.Time `jedi:"@utc=false"`
	LastUpdated *time.Time `jedi:"@last_updated"`
}

//SampleView is view of samples.
//jedi:
//view-select:
//	SELECT *
//	FROM sample
//	WHERE id > 1
//
// regular commets can restart.
type SampleView struct {
	ID          int64 `jedi:"@pk"` //you can configure the ok on the view, it adds some methods.
	Name        string
	Description string
}

//HookDemo demosntrates hooks.
//jedi:
type HookDemo struct {
	ID   int64 `jedi:"@pk"`
	Name string
}

func (p *HookDemo) beforeInsert() error {
	return errors.New("It won t happen")
}
func (p *HookDemo) beforeUpdate() error {
	return errors.New("It won t happen")
}
