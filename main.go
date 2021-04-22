package main

import (
	"frame"
)

func main() {
	f := frame.New()
	f.Get("/", func(c *frame.Context) {
		c.Json(200, "hello")
	})
	g := f.Group("/g")
	g.Get("/a", func(c *frame.Context) {
		c.Json(200, "/g/a")
	})
	f.Run(":8888")
}
