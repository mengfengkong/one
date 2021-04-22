package json

import "encoding/json"

//notice:Private variables are not printable
func ToJson(obj interface{}) string {
	if obj == nil {
		return ""
	}
	b, err := json.Marshal(obj)
	if err != nil {
		return ""
	}
	return string(b)
}

func FromJson(jsonStr string, obj interface{}) error {
	return json.Unmarshal([]byte(jsonStr), obj)
}
