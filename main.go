package main

import (
	"fmt"
	"log"
	"net/http"
	"os"
	"strconv"

	"github.com/EmiSan1998/gpsTiming-backend/datatypes"
	"github.com/EmiSan1998/gpsTiming-backend/globals"
	"github.com/EmiSan1998/gpsTiming-backend/storage"
)

func main() {
	port, err := parseSettingsFromEnvironment()
	if err != nil {
		log.Fatal("Failed to parse port number, it should be an integer")
		os.Exit(1)
	}

	globals.RegisteredServices = globals.Services{
		StorageService: storage.InMemoryStorage{
			Routes: make(map[string]datatypes.Route),
		},
	}

	r := getRouter()

	fmt.Println("gpsTiming-backend")
	fmt.Println("version " + globals.ProgramVersion + " 2020")
	fmt.Println("system online on port " + strconv.Itoa(port))

	addr := ":" + strconv.Itoa(port)
	err = http.ListenAndServe(addr, r)
	if err != nil {
		log.Fatal("ListenAndServe", err)
	}
}
