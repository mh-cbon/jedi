package model

// JoinFields represents the join between a local field and the foreign field.
type JoinFields struct {
	LocalField   *Field
	ForeignField *Field
}

//HasOne represents a has_one relation.
type HasOne struct {
	Local   *Struct
	Foreign *Struct
	Fields  []*JoinFields
}

//Many2One represents a has_many relation.
type Many2One struct {
	Local        *Struct
	Foreign      *Struct
	LocalField   *Field
	ForeignField *Field
	Fields       []*JoinFields
}

//Many2Many represents a has_many to has_many relation.
type Many2Many struct {
	Local    *Struct
	Middle   *Struct
	Foreign  *Struct
	LMFields []*JoinFields
	FMFields []*JoinFields
}
