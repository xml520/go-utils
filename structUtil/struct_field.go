package structUtil

import "reflect"

type StructField struct {
	field reflect.StructField
}

func NewStructField(field reflect.StructField) *StructField {
	return &StructField{
		field: field,
	}
}
func NewStructFields(typ reflect.Type) []*StructField {
	fields := make([]*StructField, 0)
	for _, field := range reflect.VisibleFields(typ) {
		if field.Type.Kind() == reflect.Struct && field.Anonymous {
			fields = append(fields, NewStructFields(field.Type)...)
		} else {
			fields = append(fields, NewStructField(field))
		}
	}
	return fields
}
func (f *StructField) Names() StructNames {
	return NewStructNames(f.field.Name)
}
func (f *StructField) Tag() StructFieldTag {
	return StructFieldTag(f.field.Tag)
}
func (f *StructField) Type() reflect.Type {
	return f.field.Type
}
