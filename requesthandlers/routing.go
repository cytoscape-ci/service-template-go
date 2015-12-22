package handlers


import (
	"net/http"
	"log"
	"github.com/rs/cors"
	"strconv"
)


func StartServer(port int) error {

	mux := http.NewServeMux()

	mux.HandleFunc("/", StatusHandler)
	mux.HandleFunc("/echo", EchoHandler)

	// Add your service endpoints here...


	handler := cors.Default().Handler(mux)

	log.Println("Serving API on port ", port)

	portNumStr := strconv.Itoa(port)

	err := http.ListenAndServe(":" + portNumStr, handler)

	if err != nil {
		log.Fatal("Could not start API server: ", err)
	}
	return err
}
