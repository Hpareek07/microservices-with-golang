package handlers

import (
	"log"
	"net/http"
)

type Hello struct {
	l *log.Logger
}

func NewHelloHandler(log *log.Logger) *Hello {
	return &Hello{log}
}

func (h *Hello) ServeHTTP(rw http.ResponseWriter, req *http.Request) {
	rw.Write([]byte("Hello World !!"))
}
