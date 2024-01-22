package lib

import (
	"bytes"
	"crypto/tls"
	"encoding/json"
	"errors"
	"fmt"
	log "github.com/sirupsen/logrus"
	"io"
	"io/ioutil"
	"net/http"
	"net/url"
)

type LoginRequest struct {
	UserLogin string `json:"userLogin"`
	Password  string `json:"password"`
}

func (l *LoginRequest) Login(host string) (*ZultysSession, *LoginResponse, error) {
	// Convert credentials to JSON
	jsonCredentials, err := json.Marshal(l)
	if err != nil {
		return nil, nil, errors.New("failed to marshal credentials")
	}

	if len(host) <= 0 {
		return nil, nil, errors.New("host is empty")
	}

	var apiUrl = "https://" + host + "/newapi/users"

	http.DefaultTransport.(*http.Transport).TLSClientConfig = &tls.Config{InsecureSkipVerify: true}

	// URL-encode the JSON string
	encodedData := url.QueryEscape(string(jsonCredentials))

	// Construct form data
	formData := "data=" + encodedData

	// Create a new request
	req, err := http.NewRequest("POST", apiUrl, bytes.NewBufferString(formData))
	if err != nil {
		log.Error(err)
	}
	req.Header.Set("Content-Type", "application/x-www-form-urlencoded")

	// Make the request
	client := &http.Client{}
	resp, err := client.Do(req)
	if err != nil {
		log.Error(err)
	}
	defer func(Body io.ReadCloser) {
		err := Body.Close()
		if err != nil {

		}
	}(resp.Body)

	// Handle the response
	body, err := ioutil.ReadAll(resp.Body)
	if err != nil {
		log.Error(err)
		return nil, nil, err
	}

	response := &LoginResponse{}

	err = json.Unmarshal(body, &response)
	if err != nil {
		log.Error(err)
		return nil, nil, err
	}

	var zSession = &ZultysSession{}
	zSession.Host = host
	zSession.Session = response.Session

	if response.Success == false {
		return nil, response, errors.New("login failed")
	} else if response.ApiAccess.WebAdmin == false {
		return nil, response, errors.New("user does not have WebAdmin access")
	}
	// todo handle normal user accounts for like voicemail administration and such?

	return zSession, response, err
}

type ActivationResponse struct {
	Activation struct {
		CommunicationStatus string `json:"communicationStatus"`
		ConnectionError     string `json:"connectionError"`
		Explanation         string `json:"explanation"`
		SerialNumber        string `json:"serialNumber"`
		Status              string `json:"status"`
		TimeLeftSec         int    `json:"timeLeftSec"`
	} `json:"activation"`
	Command string `json:"command"`
	Index   int    `json:"index"`
	Success bool   `json:"success"`
}

type DevicesResponse struct {
	Command string        `json:"command"`
	Devices []interface{} `json:"devices"`
	Index   int           `json:"index"`
	Success bool          `json:"success"`
}

type UserResponse struct {
	Command  string `json:"command"`
	Index    int    `json:"index"`
	ListInfo struct {
		RecordsAvailable int `json:"recordsAvailable"`
		RecordsLimit     int `json:"recordsLimit"`
		RecordsOffset    int `json:"recordsOffset"`
		RecordsReturned  int `json:"recordsReturned"`
		RecordsTotal     int `json:"recordsTotal"`
	} `json:"listInfo"`
	Success bool          `json:"success"`
	Users   []interface{} `json:"users"`
}

type UserProfileResponse struct {
	Command  string `json:"command"`
	Index    int    `json:"index"`
	Profiles []struct {
		HomeSystem  int    `json:"homeSystem"`
		ProfileId   string `json:"profileId"`
		ProfileName string `json:"profileName"`
	} `json:"profiles"`
	Success bool `json:"success"`
}

type LoginResponse struct {
	ApiAccess struct {
		WebAdmin             bool `json:"WebAdmin"`
		WebCallRecordingView bool `json:"WebCallRecordingView"`
		WebClient            bool `json:"WebClient"`
		WebFiles             bool `json:"WebFiles"`
		WebOutbound          bool `json:"WebOutbound"`
		WebSuperview         bool `json:"WebSuperview"`
		WebUserportal        bool `json:"WebUserportal"`
		WebWallboard         bool `json:"WebWallboard"`
	} `json:"api_access"`
	ApiVersions struct {
		WebAdmin             string `json:"WebAdmin"`
		WebCallRecordingView string `json:"WebCallRecordingView"`
		WebClient            string `json:"WebClient"`
		WebFiles             string `json:"WebFiles"`
		WebOutbound          string `json:"WebOutbound"`
		WebSuperview         string `json:"WebSuperview"`
		WebUserportal        string `json:"WebUserportal"`
		WebWallboard         string `json:"WebWallboard"`
	} `json:"api_versions"`
	Command  string `json:"command"`
	Features struct {
		PRLMode string `json:"PRLMode"`
	} `json:"features"`
	Index         int      `json:"index"`
	ReadCommands  []string `json:"read_commands"`
	Session       string   `json:"session"`
	Success       bool     `json:"success"`
	WriteCommands []string `json:"write_commands"`
}

type ZultysSession struct {
	Session string `json:"session"`
	Host    string `json:"host"`
}

type RequestMethod string

const (
	MethodGet  RequestMethod = "GET"
	MethodPost RequestMethod = "POST"
)

func (z *ZultysSession) SendCommand(method RequestMethod, command string, params map[string]interface{}, postData interface{}) (*http.Response, error) {
	// Construct the base URL
	if z.Host == "" {
		return nil, errors.New("host is empty")
	}

	apiUrl := "https://" + z.Host + "/newapi/?command=" + command + "&session=" + z.Session

	// Prepare query values for GET request
	queryValues := url.Values{}
	for key, value := range params {
		queryValues.Set(key, fmt.Sprintf("%v", value))
	}

	var req *http.Request
	var err error

	if method == MethodGet {
		// Append query string to the URL for GET request
		apiUrl += "&" + queryValues.Encode()
		req, err = http.NewRequest(string(method), apiUrl, nil)
	} else if method == MethodPost {
		// Handle POST request with JSON data in the body
		jsonBody, err := json.Marshal(postData)
		if err != nil {
			return nil, err
		}
		req, err = http.NewRequest(string(method), apiUrl, bytes.NewBuffer(jsonBody))
		req.Header.Set("Content-Type", "application/json")
	} else {
		return nil, errors.New("unsupported method type")
	}

	if err != nil {
		return nil, err
	}

	// Set common headers or query params if needed
	//req.Header.Set("Some-Common-Header", "value")

	// Configure the HTTP client
	http.DefaultTransport.(*http.Transport).TLSClientConfig = &tls.Config{InsecureSkipVerify: true}
	client := &http.Client{}

	// Send the request
	return client.Do(req)
}
