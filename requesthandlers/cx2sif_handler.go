package handlers

import (
	"net/http"
	"github.com/cytoscape-ci/cxtool/converter"
)

func Cx2SifHandler(w http.ResponseWriter, r *http.Request) {
	method := r.Method
	switch method {
	case POST:
		cx2sif(w, r)
	default:
		unsupported(w, r)
	}
}

func cx2sif(w http.ResponseWriter, r *http.Request) {
	conv := converter.Cx2Sif{}
	conv.Convert(r.Body, w)
}


