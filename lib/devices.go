package lib

import (
	"encoding/json"
	"io"
	"io/ioutil"
	"net/http"
)

type DeviceLine struct {
	DeviceId string `json:"deviceId"`
	SipAuth  string `json:"sipAuth"`
	SipPwd   string `json:"sipPwd"`
}

type DeviceUser struct {
	UserId    string `json:"userId"`
	UserLogin string `json:"userLogin"`
}

type Device struct {
	DeviceId       string       `json:"deviceId"`
	DeviceType     string       `json:"deviceType"`
	DeviceUniqueId string       `json:"deviceUniqueId"`
	IpAddress      string       `json:"ipAddress"`
	IsFree         bool         `json:"isFree"`
	Lines          []DeviceLine `json:"lines"`
	Location       string       `json:"location"`
	MacAddress     string       `json:"macAddress"`
	ProfileName    string       `json:"profileName"`
	Users          []DeviceUser `json:"users"`
}

type DeviceListResponse struct {
	Command string   `json:"command"`
	Devices []Device `json:"devices"`
	Index   int      `json:"index"`
	Success bool     `json:"success"`
}

func HandleDeviceListResponse(resp *http.Response) (*DeviceListResponse, error) {
	defer func(Body io.ReadCloser) {
		err := Body.Close()
		if err != nil {

		}
	}(resp.Body)

	body, err := ioutil.ReadAll(resp.Body)
	if err != nil {
		return nil, err
	}

	var deviceListResponse DeviceListResponse
	err = json.Unmarshal(body, &deviceListResponse)
	if err != nil {
		return nil, err
	}

	return &deviceListResponse, nil
}
