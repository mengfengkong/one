package fcache

import "testing"

func TestGetterFunc_Get(t *testing.T) {
	var fun GetterFunc = func(key string) (bytes []byte, err error) {
		return []byte(key), nil
	}
	v, _ := fun.Get("key")
	t.Log(string(v))
}
