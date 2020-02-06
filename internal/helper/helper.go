package helper

import "encoding/json"

// ToJSON return the JSON represenation of an object as a string
func ToJSON(i interface{}) string {
	j, _ := json.Marshal(i)
	return string(j)
}
