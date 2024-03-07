package validator

import (
	"reflect"
	"sort"
	"strings"
	"time"
)

func ValidateBool(str string) bool {
	str = strings.TrimSpace(str)
	switch str {
	case "1", "t", "T", "TRUE", "true", "True":
		return true
	case "0", "f", "F", "FALSE", "false", "False":
		return false
	}
	return false
}
func ValidateNull(str string) string {
	str = strings.TrimSpace(str)
	switch str {
	case "1", "t", "T", "TRUE", "true", "True":
		return "Null"
	case "0", "f", "F", "FALSE", "false", "False":
		return ""
	}
	return ""
}

func ValidateNumeric(str string) string {
	str = strings.TrimSpace(str)
	str = stripLeadingZeros(str)
	return str
}
func stripLeadingZeros(str string) string {
	var output string
	removeZero := true
	for _, ch := range str {
		if ch != '0' || !removeZero {
			output += string(ch)
			removeZero = false
		}
	}
	return output
}
func ValidateString(str string) (int64, error) {
	str = strings.TrimSpace(str)
	t, err := time.Parse(time.RFC3339, str)
	if err != nil {
		return 0, err
	}
	return t.Unix(), nil
}

func ValidateMap(m interface{}) (keyList []string) {
	keys := reflect.ValueOf(m).MapKeys()

	for _, key := range keys {
		keyList = append(keyList, key.Interface().(string))
	}
	sort.Strings(keyList)
	return
}
