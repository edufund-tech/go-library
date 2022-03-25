package stringutils

import (
	"fmt"
	"strings"
)

func ExtractStringToArrayStr(str string) (result []string) {
	if strings.Index(str, "in(") == 0 {
		parseStr := str[3:]
		splitStr := strings.Split(parseStr[:len(parseStr)-1], ",")

		return splitStr
	}
	return
}

func GenerateStringInFromArrayStr(arr []string) (result string) {
	result = "in("
	result = result + strings.Join(arr, ",")
	result = result + ")"
	return result
}

func GenerateStringInFromArrayInt64(arr []int64) (result string) {
	result = "in("
	result = result + strings.Trim(strings.Join(strings.Fields(fmt.Sprint(arr)), ","), "[]")
	result = result + ")"
	return result
}
