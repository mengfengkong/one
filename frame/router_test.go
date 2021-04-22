package frame

import (
	"fmt"
	"frame/util/json"
	"testing"
)

type NodeRepeat struct {
	Pattern  string        `json:"pattern"`
	Part     string        `json:"part"`
	Children []*NodeRepeat `json:"children"`
	IsWild   bool          `json:"isWild"`
}

func TestRouteInsert(t *testing.T) {
	n := new(node)
	n.insert("/a/b/c", []string{"a", "b", "c"}, 0)
	//print:{}, attribute is private
	fmt.Println(json.ToJson(*n))
	res := &NodeRepeat{
		Pattern:  n.pattern,
		Part:     n.part,
		Children: nil,
		IsWild:   n.isWild,
	}
	for {
		if len(n.children) > 0 {
			//for _, child := range n.children {
			//
			//}
		}
	}
	fmt.Println(json.ToJson(res))
}

func TestParsePattern(t *testing.T) {
	t.Log(json.ToJson(parsePattern("/a")))
	t.Log(json.ToJson(parsePattern("/a/:b/")))
	t.Log(json.ToJson(parsePattern("/a/*")))
	t.Log(json.ToJson(parsePattern("/a/*/c")))
}

func initRouter() *router {
	r := newRoute()

	//r.addRoute("get", "/", nil)
	r.addRoute("get", "/a", nil)
	r.addRoute("get", "/a/:name", nil)
	//r.addRoute("get", "/a/:b/c", nil)
	//r.addRoute("get", "/*", nil)
	r.addRoute("get", "/b/*name", nil)
	return r
}

func TestGetRoute(t *testing.T) {
	r := initRouter()
	_, params := r.getRoute("get", "/b/mm/x")
	//fmt.Println(n.pattern)
	fmt.Println(params)
	//fmt.Printf("matched path: %s, params['name']: %s\n", n.pattern, params["name"])
}
