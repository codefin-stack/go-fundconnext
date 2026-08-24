package fundconnext

import (
	"bytes"
	"encoding/json"
	"errors"
	"fmt"
	"io/ioutil"
	"net/http"
	"time"

	"github.com/dgrijalva/jwt-go"
)

type AuthResponse struct {
	Username    string `json:"username"`
	AccessToken string `json:"access_token"`
	SACode      string `json:"sa_code"`
	Claims      *AuthClaims
}

type AuthClaims struct {
	jwt.StandardClaims
	SellingAgentCode string `json:"sellingAgentCode"`
	SellingAgentId   string `json:"sellingAgentId"`
	Username         string `json:"username"`
}

const defaultAuthTimeout = 10 * time.Second

func Login(env, username, password, proxy string) (*AuthResponse, error) {
	return LoginWithTimeout(env, username, password, proxy, defaultAuthTimeout)
}

func LoginWithTimeout(env, username, password, proxy string, timeout time.Duration) (*AuthResponse, error) {
	if timeout <= 0 {
		timeout = defaultAuthTimeout
	}
	client := http.Client{
		Timeout: timeout,
	}
	if proxy != "" {
		transport, err := setProxy(&HTTPSetting{
			Proxy: proxy,
		})
		if err != nil {
			return nil, err
		}
		client.Transport = transport
	}

	reqBody, err := json.Marshal(map[string]string{
		"username": username,
		"password": password,
	})

	if err != nil {
		return nil, err
	}

	req, err := http.NewRequest("POST", fmt.Sprintf("%s/api/auth", env), bytes.NewBuffer(reqBody))
	if err != nil {
		return nil, MakeInternalError(err.Error())
	}
	req.Header.Add("Content-Type", "application/json")
	resp, err := client.Do(req)
	if err != nil {
		return nil, err
	}

	defer resp.Body.Close()
	body, err := ioutil.ReadAll(resp.Body)
	if err != nil {
		return nil, err
	}
	if resp.StatusCode == 503 {
		return nil, errors.New("service is unavailable")
	}
	if resp.StatusCode != 200 {
		var errMsg FCError
		if err := json.Unmarshal(body, &errMsg); err != nil {
			return nil, err
		}
		return nil, &errMsg
	}

	var authResp AuthResponse
	if err := json.Unmarshal(body, &authResp); err != nil {
		return nil, err
	}
	token, _ := jwt.ParseWithClaims(authResp.AccessToken, &AuthClaims{}, nil)
	if err != nil {
		panic(err)
	}
	c := token.Claims.(*AuthClaims)
	authResp.Claims = c
	return &authResp, nil
}
