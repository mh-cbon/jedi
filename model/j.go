package model

type JoinFields struct {
	LocalField   *Field
	ForeignField *Field
}

type HasOne struct {
	Local   *Struct
	Foreign *Struct
	Fields  []*JoinFields
}

type Many2One struct {
	Local   *Struct
	Foreign *Struct
	Fields  []*JoinFields
}

type Many2Many struct {
	Local    *Struct
	Middle   *Struct
	Foreign  *Struct
	LMFields []*JoinFields
	FMFields []*JoinFields
}
