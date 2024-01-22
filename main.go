package main

import (
	"encoding/json"
	"fmt"
	log "github.com/sirupsen/logrus"
)

func main() {
	// Create a new LoginRequest
	loginRequest := LoginRequest{
		UserLogin: "",
		Password:  "",
	}

	zSess, _, err := loginRequest.Login("108.165.150.21")

	// For a GET request
	getParams := map[string]interface{}{
		"page":  "1",
		"limit": "100",
	}
	// page=1&limit=100

	tt, err := zSess.SendCommand(MethodGet, "adm_get_users", getParams, nil)
	if err != nil {
		log.Fatal(err)
	}

	response, err := getUsersResponse(tt)
	if err != nil {
		return
	}

	bytes, err := json.Marshal(response)
	if err != nil {
		return
	}

	fmt.Println("Response: ", string(bytes))
}
