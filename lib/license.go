package lib

import (
	"encoding/json"
	log "github.com/sirupsen/logrus"
	"io"
	"io/ioutil"
	"net/http"
)

type Capacity struct {
	ID    string `json:"id"`
	Value int    `json:"value"`
}

type Function struct {
	ID    string `json:"id"`
	Value bool   `json:"value"`
}

type License struct {
	Capacity              []Capacity `json:"capacity"`
	DaysLeftInLicense     int        `json:"daysLeftInLicense"`
	Function              []Function `json:"function"`
	LicenseExpirationType string     `json:"licenseExpirationType"`
	LicenseId             int        `json:"licenseId"`
	StartedDate           int64      `json:"startedDate"`
	Text                  string     `json:"text"`
	UpdateAllowed         int64      `json:"updateAllowed"`
}

type InstalledLicenses struct {
	LatestLicenseId int       `json:"latestLicenseId"`
	LicenseList     []License `json:"licenseList"`
}

type Emergency struct {
	DaysInEmergencyLicense     int  `json:"daysInEmergencyLicense"`
	DaysLeftInEmergencyLicense int  `json:"daysLeftInEmergencyLicense"`
	EmergencyAvailable         bool `json:"emergencyAvailable"`
	EmergencyMode              bool `json:"emergencyMode"`
}

type OptionDescription struct {
	ID       string `json:"id"`
	Name     string `json:"name"`
	Priority int    `json:"priority"`
	Max      int    `json:"max,omitempty"`
}

type MxLicenses struct {
	Emergency          Emergency         `json:"emergency"`
	InstalledLicenses  InstalledLicenses `json:"installedLicenses"`
	OptionDescriptions struct {
		Capacity []OptionDescription `json:"capacity"`
		Function []OptionDescription `json:"function"`
	} `json:"optionDescriptions"`
}

type LicenseResponse struct {
	Command    string     `json:"command"`
	Index      int        `json:"index"`
	MxLicenses MxLicenses `json:"mxLicenses"`
	Success    bool       `json:"success"`
}

func HandleLicenseResponse(resp *http.Response) (*LicenseResponse, error) {
	defer func(Body io.ReadCloser) {
		err := Body.Close()
		if err != nil {
			log.Fatal(err)
		}
	}(resp.Body)

	body, err := ioutil.ReadAll(resp.Body)
	if err != nil {
		return nil, err
	}

	var licenseResponse LicenseResponse
	err = json.Unmarshal(body, &licenseResponse)
	if err != nil {
		return nil, err
	}

	return &licenseResponse, nil
}
