package main

import (
	"frame"
)

func main()  {
	f := frame.New()
	f.Get("/", func(c *frame.Context) {
		c.Data(200, []byte("abc"))
	})
	f.Run(":8888")
}
