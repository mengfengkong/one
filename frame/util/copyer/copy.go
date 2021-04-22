package copyer

import "frame/util/json"

func Copy(origin interface{}, target interface{}) error {
	originStr := json.ToJson(origin)
	return json.FromJson(originStr, target)
}
