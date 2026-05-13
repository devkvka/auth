package main

import (
	"fmt"
	"flag"
	"net/http"
	"log"

	"auth/internal/handler"
)

func main() {
	fmt.Println(" !-- Booting auth service ... ")

	port := flag.String("-port", ":8080", "port to serve on")
	flag.Parse()

	fmt.Println(" !-- Setting up handlers ... ")
	http.HandleFunc("/verify", handler.Verify)
	http.HandleFunc("/login", handler.Login)
	http.HandleFunc("/create-account", handler.CreateAccount)
	
	fmt.Printf(" !-- Setup ready, serving on: localhost%s \n", *port)
	log.Fatal(http.ListenAndServe(*port, nil))
}
