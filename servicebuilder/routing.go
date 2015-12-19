package servicebuilder

import (
	"net/http"
	"log"
	req "github.com/keiono/label-generator/requesthandlers"
	"github.com/rs/cors"
	"strconv"
)


func StartServer(portNumber int) (err error) {

	mux := http.NewServeMux()
	mux.HandleFunc("/", req.StatusHandler)
	mux.HandleFunc("/converter/cx2cyjs", req.Cx2CyjsHandler)

	handler := cors.Default().Handler(mux)

	log.Println("Serving API on port ", portNumber)

	portNumStr := strconv.Itoa(portNumber)

	err = http.ListenAndServe(":" + portNumStr, handler)

	if err != nil {
		log.Fatal("Could not start API server: ", err)
	}

	return nil
}
