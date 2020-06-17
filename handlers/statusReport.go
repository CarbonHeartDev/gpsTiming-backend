package handlers

import (
	"encoding/json"
	"log"
	"net/http"

	"github.com/EmiSan1998/gpsTiming-backend/globals"
	"github.com/julienschmidt/httprouter"
)

// GetStatusReport ...
func GetStatusReport(w http.ResponseWriter, r *http.Request, p httprouter.Params) {
	var response struct {
		BackendVersion string
	}
	response.BackendVersion = globals.ProgramVersion

	log.Println("Status report requested")

	responseBytes, err := json.Marshal(response)
	if err != nil {
		log.Fatal("Encoding error: ", err)
	}

	w.Write(responseBytes)
}
