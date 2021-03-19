package helperfunc

import (
	"crypto/md5"
	"encoding/hex"
	"strconv"
	"fmt"
	"reflect"
	"strings"
)

//create md5 string
func Strtomd5(s string) string {
	h := md5.New()
	h.Write([]byte(s))
	rs := hex.EncodeToString(h.Sum(nil))
	return rs
}

//password hash function
func Pwdhash(str string) string {
	return Strtomd5(str)
}

func StringsToJson(str string) string {
	rs := []rune(str)
	jsons := ""
	for _, r := range rs {
		rint := int(r)
		if rint < 128 {
			jsons += string(r)
		} else {
			jsons += "\\u" + strconv.FormatInt(int64(rint), 16) // json
		}
	}

	return jsons
}

//三元表达式
func If(condition bool, trueVal int8, falseVal int8) int8 {
	if condition {
		return trueVal
	}
	return falseVal
}

//索引加1
func Index(index int) (out int) {
	out = index+1
	return
}

//格式字符串
func FmtSprintf(value string, printfStr string) (out string) {
	out = fmt.Sprintf(value, printfStr)
	return
}

//值是否存在数组中
func InArray(val interface{}, array interface{}) (exists bool) {
	exists = false
	s := reflect.ValueOf(array)
	switch reflect.TypeOf(array).Kind() {
	case reflect.Slice, reflect.Array:
		for i := 0; i < s.Len(); i++ {
			if reflect.DeepEqual(val, s.Index(i).Interface()) == true {
				exists = true
				return
			}
		}
	case reflect.Map:
		if s.MapIndex(reflect.ValueOf(val)).IsValid() {
			exists = true
			return
		}
	}
	return
}

//面包屑
func Breadcrumb(menus []string, index int) (path string) {
	index++
	paths := menus[:index]
	path = strings.Join(paths, "/")
	return
}