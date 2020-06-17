package handlers

import (
	"encoding/json"
	"io/ioutil"
	"log"
	"net/http"

	"github.com/EmiSan1998/gpsTiming-backend/globals"

	"github.com/EmiSan1998/gpsTiming-backend/datatypes"
	"github.com/julienschmidt/httprouter"
)

// GetRoute ...
func GetRoute(w http.ResponseWriter, r *http.Request, p httprouter.Params) {
	k := p.ByName("key")

	if route, ok := globals.RegisteredServices.StorageService.GetRoute(k); ok {
		routeJSON, err := json.Marshal(route)
		if err != nil {
			log.Fatal("Serialization error")
		}
		w.Header().Set("Content-Type", "application/json")
		log.Println("Route " + route.Name + " loaded " + "from key " + k)
		w.Write(routeJSON)
	} else {
		log.Println("Route not exists")
		w.Write([]byte("Route not exists"))
	}
}

// PostRoute ...
func PostRoute(w http.ResponseWriter, r *http.Request, p httprouter.Params) {
	var route datatypes.Route

	jsn, err := ioutil.ReadAll(r.Body)
	if err != nil {
		log.Fatal("Error reading the body", err)
	}

	err = json.Unmarshal(jsn, &route)
	if err != nil {
		log.Fatal("Decoding error: ", err)
	}

	newKey := globals.RegisteredServices.StorageService.CreateRoute(route)

	log.Println("Route " + route.Name + " saved " + "on key " + newKey.String())
	var response struct {
		ID   string
		Name string
	}
	response.ID = newKey.String()
	response.Name = route.Name

	responseBytes, err := json.Marshal(response)
	if err != nil {
		log.Fatal("Encoding error: ", err)
	}

	w.Write(responseBytes)
}
