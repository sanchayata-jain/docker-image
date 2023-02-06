package server

import (
	"fmt"
	"net/http"

	"github.com/sanchayata-jain/docker-image/internal/config"
)

func New(conf *config.Config) error {
	http.HandleFunc("/", handler)
	return http.ListenAndServe(fmt.Sprintf(":%d", conf.Port), nil)
}

func handler(w http.ResponseWriter, r *http.Request) {
	w.Header().Set("Content-Type", "text/html")
	fmt.Fprint(w, "<h1> Sanchayata's Server </h1>")
}
