package main

import (
	"net/http"
	"strings"
)

type tree struct {
	node *node
}

type node struct {
	label    string
	actions  map[string]*action
	children map[string]*node
}

type action struct {
	handler http.Handler
}

type result struct {
	actions *action
}

func newResult() *result {
	return &result{}
}

func NewTree() *tree {
	return &tree{
		node: &node{
			label:    pathRoot,
			actions:  make(map[string]*action),
			children: make(map[string]*node),
		},
	}
}

const (
	pathRoot      string = "/"
	pathDelimiter string = "/"
)

func (t *tree) Insert(methods []string, path string, handler http.Handler) {
	curNode := t.node
	if path == pathRoot {
		curNode.label = path
		for _, method := range methods {
			curNode.actions[method] = &action{
				handler: handler,
			}
		}
		return nil
	}
	ep := explodePath(path)

}

func explodePath(path string) []string {
	s := strings.Split(path, pathDelimiter)
	var r []string
	for _, str := range s {
		if str != "" {
			r = append(r, str)
		}
	}
	return r
}
