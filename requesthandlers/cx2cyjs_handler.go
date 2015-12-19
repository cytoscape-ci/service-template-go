package handlers

import (
	"net/http"
	"encoding/json"
	"log"
	"errors"
	"github.com/cytoscape-ci/cxtool/converter"
)


const (
	GET = "GET"
	POST = "POST"

	ids = "ids"
	idTypes = "idTypes"
)

type Message struct {
	Code int `json:"code"`
	Message string `json:"message"`
	Error string `json:"error,omitempty"`
}

func Cx2CyjsHandler(w http.ResponseWriter, r *http.Request) {

	method := r.Method
	switch method {
	case POST:
		post(w, r)
	default:
		unsupported(w, r)
	}
}


func post(w http.ResponseWriter, r *http.Request) {
	conv := converter.Cx2Cyjs{}
	conv.Convert(r.Body, w)
}

func unsupported(w http.ResponseWriter, r *http.Request) {
	code := 405
	res := getErrorMsg(code, "You need to POST your data to use this service.",
		errors.New("Invalid HTTP method used: " + r.Method))
	log.Println("Unsupported method call from: ", r.RemoteAddr)
	http.Error(w, res, code)
}


func getErrorMsg(code int, msg string, err error) (jsonMsg string) {

	message := Message{
		Code: code,
		Message: msg,
	}

	if err != nil {
		message.Error = err.Error()
	}

	result, err := json.Marshal(message)

	if err !=nil {
		// TODO: What should I return?
	}

	return string(result)
}
