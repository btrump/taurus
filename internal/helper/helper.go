package helper

import "encoding/json"

func ToJSON(i interface{}) string {
	j, _ := json.Marshal(i)
	return string(j)
}
