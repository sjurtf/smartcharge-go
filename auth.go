package smartcharge

import (
	"bytes"
	"encoding/json"
	"io/ioutil"
	"net/http"
)

type AuthBody struct {
	NotificationToken string `json:"NotificationToken"`
	AppID             string `json:"appID"`
	Email             string `json:"Email"`
	Password          string `json:"Password"`
	PushState         string `json:"pushState"`
	AppToken          string `json:"appToken"`
}

type AuthResponse struct {
	Result Result `json:"Result"`
}

type Result struct {
	User         User     `json:"user"`
	Customer     Customer `json:"customer"`
	AccessToken  string   `json:"accessToken"`
	RefreshToken string   `json:"rToken"`
}

type User struct {
	Id       int    `json:"PK_UserID"`
	Username string `json:"Username"`
	Email    string `json:"Email"`
}

type Customer struct {
	Id    int    `json:"PK_CustomerID"`
	Email string `json:"Email"`
}

type Authentication struct {
	UserId       int
	CustomerId   int
	AccessToken  string
	RefreshToken string
}

func authenticate(email string, password string) (*Authentication, error) {
	authUrl := DefaultBaseUrl + "v2/Users/Authenticate"

	body := AuthBody{
		AppID:     DefaultAppId,
		Email:     email,
		Password:  password,
		AppToken:  DefaultAppToken,
		PushState: DefaultPushState,
	}

	authBody, err := json.Marshal(body)
	if err != nil {
		return nil, err
	}
	resp, err := http.Post(authUrl, "application/json", bytes.NewBuffer(authBody))
	if err != nil {
		return nil, err
	}

	v := &AuthResponse{}
	b, err := ioutil.ReadAll(resp.Body)
	if err != nil {
		return nil, err
	}
	err = json.Unmarshal(b, v)
	if err != nil {
		return nil, err
	}

	auth := &Authentication{
		UserId:       v.Result.User.Id,
		CustomerId:   v.Result.Customer.Id,
		AccessToken:  v.Result.AccessToken,
		RefreshToken: v.Result.RefreshToken,
	}
	return auth, nil
}
