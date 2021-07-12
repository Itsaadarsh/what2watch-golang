package main

import (
	"flag"
	"fmt"
	"log"
	"net/http"
)

const version = "1.0.0"

type config struct {
	port int
	env  string
}

func main() {
	var cfg config

	flag.IntVar(&cfg.port, "port", 4000, "Server port to listen on")
	flag.StringVar(&cfg.env, "env", "dev", "App env (dev or prod)")
	flag.Parse()

	fmt.Println("Starting...")

	http.HandleFunc("/status", func(w http.ResponseWriter, r *http.Request) {
		fmt.Fprint(w, "STATUS")
	})

	err := http.ListenAndServe(fmt.Sprintf(":%d", cfg.port), nil)
	if err != nil {
		log.Println(err)
	}
}
