package main

import (
	"frame"
	"strconv"
)

func main() {
	f := frame.Default()
	f.Get("/bb", func(c *frame.Context) {
		c.Json(200, "hello")
	})
	f.Get("/panic/:name", func(c *frame.Context) {
		names := []string{"geektutu"}
		index, _ := strconv.Atoi(c.Param("name"))
		c.Json(200, names[index])
	})
	g := f.Group("/g")
	g.Get("/a", func(c *frame.Context) {
		c.Json(200, "/g/a")
	})
	f.Run(":8888")
}
