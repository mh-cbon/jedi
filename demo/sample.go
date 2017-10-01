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

//Sample2 is whatever else.
//jedi:second_sample
type Sample2 struct {
	Name        string `jedi:"@pk"` // PK on non integer
	Description string
}

//SampleView is view of samples.
//jedi:
//view-select:
//	SELECT *
//	FROM sample
//	WHERE id > 1
//	-- comments finish on new empty line
//
// regular commets can restart.
type SampleView struct {
	ID          int64 `jedi:"@pk"` //you can configure the ok on the view, it adds some methods.
	Name        string
	Description string
}
