package structUtil

import (
	"reflect"
	"strings"
)

type Struct struct {
	Fields  []*StructField
	typ     reflect.Type
	comment string
}

func (s *Struct) Names() StructNames {
	return NewStructNames(s.typ.Name())
}
func (s *Struct) Comment() string {
	return s.comment
}

// NewStruct 传入结构体
func NewStruct(v any, comment ...string) *Struct {
	var s Struct
	if len(comment) != 0 {
		s.comment = strings.Join(comment, " ")
	}
	s.typ = reflect.TypeOf(v)
	s.Fields = NewStructFields(s.typ)
	return &s
}
