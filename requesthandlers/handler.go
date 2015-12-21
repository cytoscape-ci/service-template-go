package handlers


import (
	"net/http"
	"encoding/json"
	"errors"
	"log"
)

const (
	GET = "GET"
	POST = "POST"
)

type ErrorMessage struct {
	Code    int `json:"code"`
	Message string `json:"message"`
	Error   string `json:"error,omitempty"`
}

func unsupported(w http.ResponseWriter, r *http.Request) {
	code := 405
	res := getErrorMsg(code, "You need to POST your data to use this service.",
		errors.New("Invalid HTTP method used: " + r.Method))

	log.Println("Unsupported method call from: ", r.RemoteAddr)

	http.Error(w, res, code)
}


func getErrorMsg(code int, msg string, err error) (jsonMsg string) {
	message := ErrorMessage{
		Code: code,
		Message: msg,
	}

	if err != nil {
		message.Error = err.Error()
	}
	result, err := json.Marshal(message)

	if err != nil {
		// TODO: What should I return?
	}

	return string(result)
}

