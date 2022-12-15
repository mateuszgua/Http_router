package main

import (
	"errors"
	"net/http"
)

var (
	ErrNotFound         = errors.New("no matching route was found")
	ErrMethodNotAllowed = errors.New("methods is not allowed")
)

func handleErr(err error) int {
	var status int
	switch err {
	case ErrMethodNotAllowed:
		status = http.StatusMethodNotAllowed
	case ErrNotFound:
		status = http.StatusNotFound
	}
	return status
}
