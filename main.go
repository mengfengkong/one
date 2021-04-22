package main

import (
	"fmt"
	"frame"
)

func main() {
	f := frame.New()
	f.Get("/", func(c *frame.Context) {
		fmt.Println(1)
	})
	f.Get("/a/:name", func(c *frame.Context) {
		fmt.Println(c.Params)
	})
	f.Run(":8888")
}
