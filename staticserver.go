package staticserver

import (
	"encoding/json"
	"fmt"
	"log"
	"net/http"
	"os"
)

type Response struct {
	StatusCode  int    `json:"status_code"`
	Body        string `json:"body"`
	ContentType string `json:"content_type"`
}

type Config struct {
	Port   int                 `json:"port"`
	Routes map[string]Response `json:"routes"`
}

func (c Config) ServeHTTP(w http.ResponseWriter, req *http.Request) {
	for path, resp := range c.Routes {
		if path == req.URL.Path {
			if resp.ContentType != "" {
				w.Header().Set("Content-Type", resp.ContentType)
			}
			w.WriteHeader(resp.StatusCode)
			fmt.Fprintln(w, resp.Body)
			return
		}
	}

	w.WriteHeader(404)
	fmt.Fprintln(w, "404 - page not found")
}

func LoadConfig(filename string) Config {
	fp, err := os.Open(filename)
	if err != nil {
		log.Fatal(err)
	}
	defer fp.Close()

	d := json.NewDecoder(fp)
	var conf Config
	err = d.Decode(&conf)
	if err != nil {
		log.Fatal(err)
	}

	return conf
}
