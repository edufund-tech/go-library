package middleman

import (
	"bytes"
	"context"
	"encoding/json"
	"errors"
	"fmt"
	"io/ioutil"
	"log"
	"net/http"
	"strings"
)

//LoginResponse legacy login response
type LoginResponse struct {
	Token   string            `json:"token,omitempty"`
	Status  string            `json:"status,omitempty"`
	Message string            `json:"message,omitempty"`
	Account map[string]string `json:"account,omitempty"`
}

// Verify legacy token
func Verify(ctx context.Context, token string, permissions map[string]interface{}) (verifyResponse map[string]interface{}, err error) {
	client := &http.Client{}
	middlemanURL := ctx.Value("URL_MIDDLEMAN_VERIFY").(string)
	req, err := http.NewRequest("GET", middlemanURL, nil)
	if err != nil {
		log.Printf("[Middleman Verify] error 1 %v", err)
		return nil, err
	}
	req.Header.Add("Authorization", token)
	resp, err := client.Do(req)
	if err != nil {
		log.Printf("[Middleman Verify] error 2 %v", err)
		return nil, err
	}
	defer resp.Body.Close()
	//Read the response body
	if resp.StatusCode != 200 {
		log.Printf("[Middleman Verify] error 3 %v", resp.Status)
		err = errors.New(fmt.Sprintf("Legacy Host Error Code %d (%s)", resp.StatusCode, resp.Status))
		return nil, err
	}
	body, err := ioutil.ReadAll(resp.Body)
	if err != nil {
		log.Printf("[Middleman Verify] error 4 %v", err)
		return nil, err
	}
	json.Unmarshal(body, &verifyResponse)

	return verifyResponse, err
}

// Login legacy token
func Login(ctx context.Context, payload []byte, permissions map[string]interface{}) (loginResponse LoginResponse, err error) {
	responseBody := bytes.NewBuffer(payload)
	//Leverage Go's HTTP Post function to make request
	middlemanURL := ctx.Value("AUTH_MIDDLEMAN_LOGIN").(string)
	resp, err := http.Post(middlemanURL, "application/json", responseBody)

	//Handle Error
	if err != nil {
		log.Printf("[Middleman Login] error 1 %v", err)
		return loginResponse, err
	}
	defer resp.Body.Close()

	if resp.StatusCode != 200 {
		log.Printf("[Middleman Login] error 2 %v", resp.Status)
		err = errors.New(fmt.Sprintf("Legacy Host Error Code %d (%s)", resp.StatusCode, resp.Status))
		return loginResponse, err
	}
	//Read the response body
	body, err := ioutil.ReadAll(resp.Body)
	if err != nil {
		log.Printf("[Middleman Login] error 3 %v", err)
		return loginResponse, err
	}
	json.Unmarshal(body, &loginResponse)

	if strings.ToLower(loginResponse.Status) != "success" {
		err = errors.New(loginResponse.Message)
		log.Printf("[Middleman Login] error 4 %v", err)
	}
	return loginResponse, err
}
