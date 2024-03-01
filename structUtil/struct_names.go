package structUtil

import (
	"go-utils/strUtil"
	"strings"
	"unicode"
)

type StructNames []string

func NewStructNames(str string) StructNames {
	names := make(StructNames, 1)
	for i, s := range str {
		if unicode.IsUpper(s) && i != 0 {
			names = append(names, "")
		}
		names[len(names)-1] += string(s)
	}
	return names
}
func (s StructNames) String() string {
	return strings.Join(s, "")
}

// Snake 转蛇形
func (s StructNames) Snake(translate ...string) string {
	return strUtil.CamelToSnake(s.String(), translate...)
}

// LowerCamel Camel 转驼峰
func (s StructNames) LowerCamel() string {
	return strUtil.ToLower(s.String(), 1)
}

// RawJoin 原样拼接
func (s StructNames) RawJoin(sep string) string {
	return strUtil.ToLower(strings.Join(s, sep))
}

// Join 转小写拼接
func (s StructNames) Join(sep string) string {
	return strUtil.ToLower(strings.Join(s, sep))
}
