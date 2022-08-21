package base

import (
	"crypto/rand"
	"fmt"
	"regexp"
	"strings"
)

func RandomString(length int) string {
	var chars = "abcdefghijklmnopqrstuvwxyzABCDEFGHIJKLMNOPQRSTUVWXYZ0123456789"
	var bytes = make([]byte, length)
	rand.Read(bytes)
	for k, v := range bytes {
		bytes[k] = chars[v%byte(len(chars))]
	}
	return string(bytes)
}

func Populate(str string, data PopulateFields) string {
	var s = str
	s = strings.Replace(s, "${RANDOM_STRING}", RandomString(16), -1)
	m := regexp.MustCompile(`\$\{[^}]+\}`)
	fmt.Println(data)
	return m.ReplaceAllStringFunc(s, func(t string) string {
		t = strings.Replace(t, "${", "", -1)
		t = strings.Replace(t, "}", "", -1)
		fmt.Println("Replacing " + t + " with " + data[t])
		value, ok := data[t]
		if !ok {
			return t + "_NOT_FOUND"
		}
		return value
	})
}

func GetNames[T Namable](vs []T) []string {
	vsm := make([]string, len(vs))
	for i, v := range vs {
		vsm[i] = v.getName()
	}
	return vsm
}
