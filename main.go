package main

import (
	"fmt"
	"log"
	"net/http"
	"os"

	"github.com/hpareek07/microservices-with-golang/handlers"
)

func main() {

	l := log.New(os.Stdout, "INFO::", log.LstdFlags)

	helloHandler := handlers.NewHelloHandler(l)

	sm := http.NewServeMux()
	sm.Handle("/", helloHandler)
	fmt.Println("Starting the server at port 9090")

	http.ListenAndServe(":9090", nil)
}
