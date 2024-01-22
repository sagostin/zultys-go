package main

import (
	"encoding/json"
	"fmt"
	log "github.com/sirupsen/logrus"
)

func main() {
	// Create a new LoginRequest
	loginRequest := LoginRequest{
		UserLogin: "Administrator",
		Password:  "",
	}

	zSess, _, err := loginRequest.Login("")

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

	response, err := handleGetUsersResponse(tt)
	if err != nil {
		return
	}

	_, err = json.Marshal(response)
	if err != nil {
		return
	}

	devices, err := zSess.SendCommand(MethodGet, "adm_get_devices", getParams, nil)
	if err != nil {
		log.Fatal(err)
	}

	deviceResp, err := handleDeviceListResponse(devices)
	if err != nil {
		return
	}

	bytesDeviceResp, err := json.Marshal(deviceResp)
	if err != nil {
		return
	}

	fmt.Println("Response: ", string(bytesDeviceResp))
}
