package handlers


import (
	"net/http"
	"encoding/json"
)


type Status struct {
	Name        string `json:"name"`
	Version     string `json:"version"`
	Description string `json:"description"`
	Documents   string `json:"documents"`
}


func StatusHandler(w http.ResponseWriter, r *http.Request) {

	status := Status{
		Name: "Sample CI Service Server",
		Version: "v1",
		Description: "CI Service Template for Go.",
		Documents: "https://github.com/cytoscape-ci/service-template-go",
	}

	method := r.Method

	switch method {

	case GET:
		json.NewEncoder(w).Encode(status)
	default:
		unsupported(w, r, GET)
	}
}
