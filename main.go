package main

import (
	"encoding/json"
	"fmt"
	log "github.com/sirupsen/logrus"
	zultys "zultys-go/lib"
)

func main() {
	// Create a new LoginRequest
	loginRequest := zultys.LoginRequest{
		UserLogin: "Administrator",
		Password:  "",
	}

	zSess, _, err := loginRequest.Login("108.165.150.21")

	if err != nil {
		log.Fatal(err)
	}

	// For a GET request
	getParams := map[string]interface{}{
		"page":  "1",
		"limit": "100",
	}
	// page=1&limit=100

	tt, err := zSess.SendCommand(zultys.MethodGet, "adm_get_users", getParams, nil)
	if err != nil {
		log.Fatal(err)
	}

	response, err := zultys.HandleGetUsersResponse(tt)
	if err != nil {
		return
	}

	_, err = json.Marshal(response)
	if err != nil {
		return
	}

	devices, err := zSess.SendCommand(zultys.MethodGet, "adm_get_devices", getParams, nil)
	if err != nil {
		log.Fatal(err)
	}

	deviceResp, err := zultys.HandleDeviceListResponse(devices)
	if err != nil {
		return
	}

	bytesDeviceResp, err := json.Marshal(deviceResp)
	if err != nil {
		return
	}

	fmt.Println("Response: ", string(bytesDeviceResp))

	licenses, err := zSess.SendCommand(zultys.MethodGet, "adm_get_licenses", getParams, nil)
	if err != nil {
		log.Fatal(err)
	}

	licensesResp, err := zultys.HandleLicenseResponse(licenses)
	if err != nil {
		return
	}

	bytesLicensesResp, err := json.Marshal(licensesResp)
	if err != nil {
		return
	}

	fmt.Println("Response: ", string(bytesLicensesResp))
}
