package stringutils

import (
	"fmt"
	"math/rand"
	"strings"
)

const (
	alphabeticalLetters = "abcdefghijklmnopqrstuvwxyzABCDEFGHIJKLMNOPQRSTUVWXYZ"
	alphanumericLetters = "1234567890abcdefghijklmnopqrstuvwxyzABCDEFGHIJKLMNOPQRSTUVWXYZ"
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

func Contains(a []string, x string) bool {
	for _, n := range a {
		if x == n {
			return true
		}
	}
	return false
}

func RandAlphabeticalString(n int) string {
	b := make([]byte, n)
	for i := range b {
		b[i] = alphabeticalLetters[rand.Intn(len(alphabeticalLetters))]
	}
	return string(b)
}

func RandAlphanumericString(n int) string {
	b := make([]byte, n)
	for i := range b {
		b[i] = alphanumericLetters[rand.Intn(len(alphanumericLetters))]
	}
	return string(b)
}
