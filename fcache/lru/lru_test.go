package lru

import (
	"fmt"
	"strconv"
	"testing"
)

type SS string

func (s SS) Len() int {
	return len(s)
}

func TestCache_Add(t *testing.T) {
	c := New(10, func(s string, value Value) {
		fmt.Println(s, value)
	})

	for i := 0; i < 10; i++ {
		t := strconv.Itoa(i)
		c.Add(t, SS(t))
		fmt.Printf("len:%d,bytes:%d\n", c.Len(), c.nBytes)
	}
	fmt.Println(c.Get("0"))
}

func TestCache_Get(t *testing.T) {
	for i := 0; i < 10; i++ {
		go func(a int) {
			fmt.Println(a)
		}(i)
	}
}
