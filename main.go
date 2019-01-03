package main

import (
	"flag"
	"fmt"
	"log"
	"net/http"
	"os/exec"
)

func cmdShutdown() error {
	return exec.Command("shutdown").Run()
}

func handleShutdown(securityCode string) http.HandlerFunc {
	return func(w http.ResponseWriter, r *http.Request) {
		if r.Method != "GET" {
			log.Printf("incorrect method received: %v", r.Method)
			w.WriteHeader(404)
			return
		}

		sKey := r.URL.Query().Get("s")

		if sKey == "" {
			log.Println("no security key provided")
			return
		}

		if sKey != securityCode {
			log.Println("security codes don't match")
			return
		}

		err := cmdShutdown()
		if err != nil {
			w.WriteHeader(500)
			msg := "Failed to shutdown system"
			log.Printf("%s - %v", msg, err)
			fmt.Fprintf(w, msg)
			return
		}

		msg := "System is going to shutdown soon"
		log.Println(msg)
		fmt.Fprintf(w, msg)
	}
}

// This application shutdowns the **Linux server** with **shutdown** command.
// For additional security, you should genereate some secret code (string) and start your application with it.
// You should provide this secret code in the request to `GET /shutdown` in **s** parameter.
func main() {
	var (
		securityCode = flag.String("sec-code", "", "Security code")
		port         = flag.String("port", "9898", "Port to listen")
	)
	flag.Parse()

	http.HandleFunc("/shutdown", handleShutdown(*securityCode))
	log.Fatal(http.ListenAndServe(":"+*port, nil))
}
