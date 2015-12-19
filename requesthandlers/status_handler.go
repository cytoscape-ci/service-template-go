package handlers

import (
	"net/http"
	"encoding/json"
)


type Status struct {
	Name string `json:"name"`
	Version string `json:"version"`
	Description string `json:"description"`
	Documents string `json:"documents"`
}


func StatusHandler(w http.ResponseWriter, r *http.Request) {
	serviceStatus := Status{
		Name:"CX Converter",
		Version:"v1",
		Description:"Converts CX stream into other formats.",
		Documents: "https://github.com/cytoscape-ci/service-go",
	}

	if r.Method == "GET" {
		json.NewEncoder(w).Encode(serviceStatus)
	} else {
		http.Error(w, "Request method must be GET.", 405)
	}
}
