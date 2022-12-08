package string_utils

import "encoding/json"

func ToJSONPrettier[T any](result T) string {
	jsonResult, _ := json.MarshalIndent(result, "", "    ")
	return string(jsonResult)
}
