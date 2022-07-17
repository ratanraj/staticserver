package main

import (
	"fmt"
	"os"
	"encoding/json"
	"log"
	"net/http"
	"time"
	"github.com/ratanraj/staticserver"
)


func main() {
	fp, err := os.Open(os.Args[1])
	if err != nil {
		log.Fatal(err)
	}
	defer fp.Close()

	d := json.NewDecoder(fp)
	var conf staticserver.Config
	err = d.Decode(&conf)
	if err != nil {
		log.Fatal(err)
	}

	server := &http.Server{
		Addr: fmt.Sprintf(":%d", conf.Port),
		Handler: conf,
		ReadTimeout: 10 * time.Second,
		WriteTimeout: 10 * time.Second,
		MaxHeaderBytes: 1 << 20,
	}

	log.Fatal(server.ListenAndServe())
}
