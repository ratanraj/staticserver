package staticserver


import (
	"net/http"
	"fmt"
)

type Response struct {
	StatusCode int `json:"status_code"`
	Body string `json:"body"`
}

type Config struct {
	Port int `json:"port"`
	Routes map[string]Response `json:"routes"`
}

func (c Config) ServeHTTP(w http.ResponseWriter, req *http.Request) {
	for path, resp := range c.Routes {
		if path == req.URL.Path {
			w.WriteHeader(resp.StatusCode)
			fmt.Fprintln(w, resp.Body)
			return
		}
	}
	
	w.WriteHeader(404)
	fmt.Fprintln(w, "404 - page not found")
}
