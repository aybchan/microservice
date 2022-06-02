package handlers

import (
	"log"
	"net/http"
	"os"
)

type Bye struct {
	l *log.Logger
}

func NewBye(l *log.Logger) *Bye {
	return &Bye{l}
}

func (b *Bye) ServeHTTP(rw http.ResponseWriter, r *http.Request) {
	b.l.Println("Bye World")
	os.Exit(3)
}
