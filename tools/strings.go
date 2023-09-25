package tools

import (
	"strings"
	"unicode"
)

// Capitalize 首字母大写
func Capitalize(str string) string {
	if str == "" {
		return ""
	}
	str = strings.ToLower(str)
	vv := []rune(str)
	return strings.ToUpper(string(vv[0])) + string(vv[1:])
}

// UnderlineToUpperCamelCase 下划线单词转为大写驼峰单词
func UnderlineToUpperCamelCase(s string) string {
	vv := strings.Split(s, "_")
	for i := 0; i < len(vv); i++ {
		vv[i] = Capitalize(vv[i])
	}
	return strings.Join(vv, "")
}

// UnderlineToLowerCamelCase 下划线单词转为小写驼峰单词
func UnderlineToLowerCamelCase(s string) string {
	vv := strings.Split(s, "_")
	vv[0] = strings.ToLower(vv[0])
	for i := 1; i < len(vv); i++ {
		vv[i] = Capitalize(vv[i])
	}
	return strings.Join(vv, "")
}

// CamelCaseToUnderline 驼峰单词转下划线单词
func CamelCaseToUnderline(s string) string {
	var output []rune
	for i, r := range s {
		if i == 0 {
			output = append(output, unicode.ToLower(r))
		} else {
			if unicode.IsUpper(r) {
				output = append(output, '_')
			}

			output = append(output, unicode.ToLower(r))
		}
	}
	return string(output)
}
