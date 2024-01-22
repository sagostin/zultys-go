package lib

import (
	"encoding/json"
	"io"
	"io/ioutil"
	"net/http"
)

type UserProfile struct {
	ProfileId   string `json:"profileId"`
	ProfileName string `json:"profileName"`
}

type User struct {
	AdminProfile       UserProfile `json:"adminProfile"`
	CallRecProfile     UserProfile `json:"callRecProfile"`
	CallerId           string      `json:"callerId"`
	CellPhone          string      `json:"cellPhone"`
	ChangePwdOnLogin   bool        `json:"changePwdOnLogin"`
	DefaultRole        UserProfile `json:"defaultRole"`
	DevicePassword     string      `json:"devicePassword"`
	Devices            string      `json:"devices"`
	Did                string      `json:"did"`
	Email              string      `json:"email"`
	Email2             string      `json:"email2"`
	Extension          string      `json:"extension"`
	FaxDID             string      `json:"faxDID"`
	FaxNumber          string      `json:"faxNumber"`
	FirstName          string      `json:"firstName"`
	LastName           string      `json:"lastName"`
	LdapAuthentication bool        `json:"ldapAuthentication"`
	ModifiedTS         int64       `json:"modifiedTS"`
	MsExchangeId       string      `json:"msExchangeId"`
	MxnNodeId          int         `json:"mxnNodeId"`
	MxnNodeIdCurrent   int         `json:"mxnNodeIdCurrent"`
	PagingProfile      UserProfile `json:"pagingProfile"`
	Password           string      `json:"password"`
	Pin                string      `json:"pin"`
	PrlMode            string      `json:"prlMode"`
	Pseudonym          string      `json:"pseudonym"`
	UniqueId           string      `json:"uniqueId"`
	UserBlocked        bool        `json:"userBlocked"`
	UserId             string      `json:"userId"`
	UserLogin          string      `json:"userLogin"`
	UserProfile        UserProfile `json:"userProfile"`
}

type UserListResponse struct {
	Command  string `json:"command"`
	Index    int    `json:"index"`
	ListInfo struct {
		RecordsAvailable int `json:"recordsAvailable"`
		RecordsLimit     int `json:"recordsLimit"`
		RecordsOffset    int `json:"recordsOffset"`
		RecordsReturned  int `json:"recordsReturned"`
		RecordsTotal     int `json:"recordsTotal"`
	} `json:"listInfo"`
	Success bool   `json:"success"`
	Users   []User `json:"users"`
}

func HandleGetUsersResponse(resp *http.Response) (*UserListResponse, error) {
	defer func(Body io.ReadCloser) {
		err := Body.Close()
		if err != nil {

		}
	}(resp.Body)

	body, err := ioutil.ReadAll(resp.Body)
	if err != nil {
		return nil, err
	}

	var userListResponse UserListResponse
	err = json.Unmarshal(body, &userListResponse)
	if err != nil {
		return nil, err
	}

	return &userListResponse, nil
}
