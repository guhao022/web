package web

import (
	"encoding/json"
	"net/http"
)

const (
	BAD_REQUEST_ERROR = iota + 100000
)

type Error struct {
	ID     string      `json:"id,omitempty"`
	Links  *ErrLinks   `json:"links"`
	Status int         `json:"_"`
	Code   string      `json:"code"`
	Title  string      `json:"title"`
	Detail string      `json:"detail"`
	Source *ErrSource  `json:"source"`
	Meta   interface{} `json:"meta"`
}

type ErrLinks struct {
	About string `json:"about"`
}

type ErrSource struct {
	Pointer   string `json:"pointer"`
	Parameter string `json:"parameter"`
}

func NewError(w http.ResponseWriter, err *Error) error {
	w.Header().Set("Content-Type", "application/json; charset=UTF-8")
	w.WriteHeader(err.Status)
	return json.NewEncoder(w).Encode(map[string][]*Error{"errors": []*Error{err}})
}
