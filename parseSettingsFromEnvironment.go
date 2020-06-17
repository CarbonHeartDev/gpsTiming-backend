package main

import (
	"errors"
	"os"
	"strconv"
)

func parseSettingsFromEnvironment() (port int, err error) {
	if len(os.Getenv("GPSTIMING_BACKEND_PORT")) == 0 {
		return 8080, nil
	} else if port, error := strconv.Atoi(os.Getenv("GPSTIMING_BACKEND_PORT")); error == nil {
		return port, nil
	} else {
		return 0, errors.New("GPSTIMING_BACKEND_PORT should be an integer number")
	}
}
