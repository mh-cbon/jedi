package model

import "strings"

//FindStruct search for the struct matching n, where n can be [struct name] or [struct.field name]
func FindStruct(all []*Struct, n string) *Struct {
	if strings.Index(n, ".") > -1 {
		n = strings.Split(n, ".")[0]
	}
	for _, a := range all {
		if a.Name == n {
			return a
		}
	}
	return nil
}

//FindProp search for the property matching n, where n is [struct.field name]
func FindProp(all []*Struct, n string) *Field {
	var nprop string
	if strings.Index(n, ".") == -1 {
		return nil
	}
	nprop = strings.Split(n, ".")[1]
	n = strings.Split(n, ".")[0]
	for _, a := range all {
		if a.Name == n {
			for _, f := range a.Fields {
				if f.Name == nprop {
					return f
				}
			}
		}
	}
	return nil
}

// HasMany2ManyGoTypeName returns the go type name for a many 2 many relation
func HasMany2ManyGoTypeName(left, right *Struct, leftP, rightP *Field) string {
	if leftP.On != "" {
		return leftP.On
	}
	if rightP.On != "" {
		return rightP.On
	}
	leftName := left.Name + "." + leftP.Name
	rightName := right.Name + "." + rightP.Name
	var n string
	if leftName > rightName {
		n = rightName + "To" + leftName
	} else {
		n = leftName + "To" + rightName
	}
	return strings.Replace(n, ".", "", -1)
}

// HasMany2ManyTableName returns the table name for a many 2 many relation
func HasMany2ManyTableName(left, right *Struct, leftP, rightP *Field) string {
	if leftP.On != "" {
		return leftP.On
	}
	if rightP.On != "" {
		return rightP.On
	}
	leftName := left.Name + "." + leftP.Name
	rightName := right.Name + "." + rightP.Name
	var n string
	if leftName > rightName {
		n = rightName + "To" + leftName
	} else {
		n = leftName + "To" + rightName
	}
	n = strings.Replace(n, ".", "_", -1)
	return strings.ToLower(n)
}
