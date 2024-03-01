package structUtil

import (
	"reflect"
	"regexp"
)

type StructFieldTag string

func (t StructFieldTag) String() string {
	return string(t)
}

// Get 获取标签内容
func (t StructFieldTag) Get(key string) string {
	return reflect.StructTag(t).Get(key)
}

// Match 正则查找返回括号内字符串 `xorm:"comment(好吧)"`  reg = ”comment\\(.(.?)\\)“
func (t StructFieldTag) Match(reg string, key ...string) string {
	var str string
	if len(key) == 1 {
		str = t.Get(key[0])
		if str == "" {
			return ""
		}
	} else {
		str = string(t)
	}
	if r := regexp.MustCompile(reg).FindStringSubmatch(str); len(r) == 2 {
		return r[1]
	}
	return ""
}

// Exist 正则查找关键词是否存在
func (t StructFieldTag) Exist(reg string, key ...string) bool {
	var str string
	if len(key) == 1 {
		str = t.Get(key[0])
		if str == "" {
			return false
		}
	} else {
		str = string(t)
	}
	if r := regexp.MustCompile(reg).FindString(str); len(r) != 0 {
		return true
	}
	return false
}
