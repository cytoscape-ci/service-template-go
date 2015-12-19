package servicebuilder

import (
	"net/http"
	"log"
	"github.com/rs/cors"
	"strconv"
	req "github.com/cytoscape-ci/service-cx/requesthandlers"
)


func StartServer(portNumber int) (err error) {

	mux := http.NewServeMux()
	mux.HandleFunc("/", req.StatusHandler)
	mux.HandleFunc("/converter/cx2cyjs", req.Cx2CyjsHandler)
	mux.HandleFunc("/converter/cx2sif", req.Cx2SifHandler)

	handler := cors.Default().Handler(mux)

	log.Println("Serving API on port ", portNumber)

	portNumStr := strconv.Itoa(portNumber)

	err = http.ListenAndServe(":" + portNumStr, handler)

	if err != nil {
		log.Fatal("Could not start API server: ", err)
	}

	return nil
}
