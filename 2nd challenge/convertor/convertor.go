package convertor

import (
	validator "CodeChallenge/sanitizer"
	"CodeChallenge/schema"
	"encoding/json"
	"strconv"
	"strings"
)

func ConvertJsonToStruct(inp string) schema.MyJsonName {
	output := schema.MyJsonName{}
	json.Unmarshal([]byte(inp), &output)
	return output
}
func ConvertStructToJson(inp schema.OutputFile) (string, error) {
	var out []byte
	out, err := json.Marshal(inp)
	if err != nil {
		return "", err
	}
	return string(out), nil
}

func Transformer(inpForm schema.MyJsonName) (schema.OutputFile, error) {
	response := make(schema.OutputFile, 1)

	response[0].Number1, _ = strconv.ParseFloat(validator.ValidateNumeric(inpForm.Number1.N), 64)
	response[0].String1 = strings.TrimSpace(inpForm.String1.S)

	str1, err := validator.ValidateString(inpForm.String1.S)
	if err != nil {
		return response, nil
	}
	response[0].String2 = str1

	rest := []interface{}{}
	for _, v := range inpForm.Map1.M.List1.L {
		rest = append(rest, validator.ValidateBool(v.Bool))
		rest = append(rest, validator.ValidateNumeric(v.N))
		rest = append(rest, validator.ValidateNull(v.Null))
	}
	response[0].Map1.List1 = rest
	response[0].Map1.Null1 = validator.ValidateNull(inpForm.Map1.M.Null1.NULL)

	return response, nil
}
