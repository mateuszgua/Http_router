package trie

import (
	my_err "mateusz/http_router/errors"
	"net/http"
	"strings"
)

type Tree struct {
	node *node
}

type node struct {
	label    string
	Actions  map[string]*action
	children map[string]*node
}

type action struct {
	Handler http.Handler
}

type result struct {
	Actions *action
}

const (
	pathRoot      string = "/"
	pathDelimiter string = "/"
)

func newResult() *result {
	return &result{}
}

func NewTree() *Tree {
	return &Tree{
		node: &node{
			label:    pathRoot,
			Actions:  make(map[string]*action),
			children: make(map[string]*node),
		},
	}
}

func (t *Tree) Insert(methods []string, path string, handler http.Handler) error {
	curNode := t.node
	if path == pathRoot {
		curNode.label = path
		for _, method := range methods {
			curNode.Actions[method] = &action{
				Handler: handler,
			}
		}
		return nil
	}
	ep := explodePath(path)

	for i, p := range ep {
		nextNode, ok := curNode.children[p]
		if ok {
			curNode = nextNode
		}

		if !ok {
			curNode.children[p] = &node{
				label:    p,
				Actions:  make(map[string]*action),
				children: make(map[string]*node),
			}
			curNode = curNode.children[p]
		}

		if i == len(ep)-1 {
			curNode.label = p
			for _, method := range methods {
				curNode.Actions[method] = &action{
					Handler: handler,
				}
			}
			break
		}
	}
	return nil
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

func (t *Tree) Search(method string, path string) (*result, error) {
	result := newResult()
	curNode := t.node
	if path != pathRoot {
		for _, p := range explodePath(path) {
			nextNode, ok := curNode.children[p]
			if !ok {
				if p == curNode.label {
					break
				} else {
					return nil, my_err.ErrNotFound
				}
			}
			curNode = nextNode
			continue
		}
	}
	result.Actions = curNode.Actions[method]
	if result.Actions == nil {
		return nil, my_err.ErrMethodNotAllowed
	}
	return result, nil
}
