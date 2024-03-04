package strUtil

import (
	"crypto/md5"
	"encoding/hex"
	"math/rand"
	"strconv"
	"strings"
	"time"
	"unicode"
	"unsafe"
)

const (
	letters = "abcdefghijklmnopqrstuvwxyzABCDEFGHIJKLMNOPQRSTUVWXYZ1234567890"
	// 6 bits to represent a letter index
	letterIdBits = 6
	// All 1-bits as many as letterIdBits
	letterIdMask = 1<<letterIdBits - 1
	letterIdMax  = 63 / letterIdBits
)

var src = rand.NewSource(time.Now().UnixNano())

// SnakeToCamel 字符串蛇形转驼峰,如需翻译单词请输入追加
func SnakeToCamel(str string, translate ...string) string {
	for _, s := range strings.Split(str, "_") {
		str += strings.ToUpper(s[:1]) + s[1:]
	}
	for _, s := range translate {
		str = strings.ReplaceAll(str, ToUpper(ToLower(s), 1), ToUpper(s))
	}
	return str
}

// CamelToSnake 字符串驼峰转蛇形,如需翻译单词请输入追加
func CamelToSnake(str string, translate ...string) string {
	for _, s := range translate {
		str = strings.ReplaceAll(str, ToUpper(s), ToUpper(ToLower(s), 1))
	}
	strArr := make([]string, 1)
	for i, s := range str {
		if i != 0 && unicode.IsUpper(s) {
			strArr = append(strArr, ToLower(string(s)))
		} else {
			strArr[len(strArr)-1] += string(s)
		}
	}
	return ToLower(strings.Join(strArr, "_"))
}

// ToUpper 字符串转大写，默认将全部字符串转成大写
func ToUpper(str string, v ...int) string {
	if len(v) == 0 || len(str) < v[0] || v[0] < 1 {
		return strings.ToUpper(str)
	}
	length := v[0]
	return strings.ToUpper(str[:length]) + str[length:]
}

// ToLower 字符串转小写，默认将全部字符串转成小写
func ToLower(str string, v ...int) string {
	if len(v) == 0 || len(str) < v[0] || v[0] < 1 {
		return strings.ToLower(str)
	}
	length := v[0]
	return strings.ToLower(str[:length]) + str[length:]
}

// RandStr 获取随机字符串
func RandStr(n int, str ...string) string {
	b := make([]byte, n)
	var strTmp string
	if len(str) == 0 {
		strTmp = letters
	} else {
		strTmp = strings.Join(str, "")
	}
	// A rand.Int63() generates 63 random bits, enough for letterIdMax letters!
	for i, cache, remain := n-1, src.Int63(), letterIdMax; i >= 0; {
		if remain == 0 {
			cache, remain = src.Int63(), letterIdMax
		}
		if idx := int(cache & letterIdMask); idx < len(strTmp) {
			b[i] = strTmp[idx]
			i--
		}
		cache >>= letterIdBits
		remain--
	}
	return *(*string)(unsafe.Pointer(&b))
}

// SplitInt 分割字符串并且转int分片
func SplitInt(str, sep string) ([]int, error) {
	var err error
	strArr := strings.Split(str, sep)
	intArr := make([]int, len(strArr))
	for i, s := range strArr {
		if intArr[i], err = strconv.Atoi(s); err != nil {
			return nil, err
		}
	}
	return intArr, nil
}

// ToInt 字符串转数值，失败返回 -1
func ToInt(str string) int {
	v, err := strconv.Atoi(str)
	if err != nil {
		return -1
	}
	return v
}

// Md5 获取字符串md5
func Md5(str string) string {
	hash := md5.Sum([]byte(str))
	return hex.EncodeToString(hash[:])
}

//s
