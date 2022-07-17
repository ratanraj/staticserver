package main

import (
	"fmt"
	"log"
	"net/http"
	"os"
	"time"

	"github.com/ratanraj/staticserver"
)

func main() {
	conf := staticserver.LoadConfig(os.Args[1])

	server := &http.Server{
		Addr:           fmt.Sprintf(":%d", conf.Port),
		Handler:        conf,
		ReadTimeout:    10 * time.Second,
		WriteTimeout:   10 * time.Second,
		MaxHeaderBytes: 1 << 20,
	}

	log.Fatal(server.ListenAndServe())
}
