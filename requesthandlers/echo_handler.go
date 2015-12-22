package handlers


import (
	"net/http"
	"io/ioutil"
)


func EchoHandler(w http.ResponseWriter, r *http.Request) {

	method := r.Method

	switch method {

	case POST:
		// Simply return user input.
		echo(w, r)
	default:
		unsupported(w, r, POST)
	}
}


func echo(w http.ResponseWriter, r *http.Request) {
	body, err := ioutil.ReadAll(r.Body)

	if err != nil {
		msg := getErrorMsg(400, "Failed to read input.", err)
		http.Error(w, msg, 400)
		return
	} else {
		w.Write(body)
	}
}
