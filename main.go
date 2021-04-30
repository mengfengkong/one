package main

import (
	"fcache"
	"fmt"
	"log"
	"net/http"
)

var db = map[string]string{
	"a": "aa",
}

func main() {
	fcache.NewGroup("store", 2<<10, fcache.GetterFunc(
		func(key string) ([]byte, error) {
			if v, ok := db[key]; ok {
				return []byte(v), nil
			}

			return nil, fmt.Errorf("%s not exist", key)
		}))
	addr := "localhost:8888"
	log.Println("f is running at", addr)
	log.Fatal(http.ListenAndServe(addr, fcache.NewHttpPool(addr)))
}
