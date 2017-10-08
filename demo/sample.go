package main

import "time"

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
	Name        string `jedi:"@pk"` // PK on non integer
	Description string
}

//CompositePk have a text primary key column.
//jedi:
type CompositePk struct {
	P           string `jedi:"@pk"`
	K           string `jedi:"@pk"`
	Description string
}

//DateType have dates properties.
//jedi:
type DateType struct {
	ID int64 `jedi:"@pk"`
	T  time.Time
	TP *time.Time
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
