package main

import (
	"log"
	"net/http"
	"os"

	"github.com/nilJackson/handlers"
)

func main() {
	l := log.New(os.Stdout, "product-api", log.LstdFlags)
	hh := handlers.NewHello(l)

	sm := http.NewServeMux()
	sm.Handle("/", hh)

	s := &http.Server{}

	http.ListenAndServe(":8080", sm)
}
