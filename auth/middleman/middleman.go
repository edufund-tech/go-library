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

//LegacyResponse legacy response
type LegacyResponse struct {
	Token   string            `json:"token,omitempty"`
	Status  string            `json:"status,omitempty"`
	Message string            `json:"message,omitempty"`
	Account map[string]string `json:"account,omitempty"`
}

// Verify legacy token
func Verify(ctx context.Context, token string, permissions map[string]interface{}) (legacyResponse map[string]interface{}, err error) {
	client := &http.Client{}

	req, err := http.NewRequest("GET", ctx.Value("URL_MIDDLEMAN_VERIFY").(string), nil)
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
	json.Unmarshal(body, &legacyResponse)

	return legacyResponse, err
}

// Login legacy token
func Login(ctx context.Context, payload []byte, permissions map[string]interface{}) (legacyResponse LegacyResponse, err error) {
	responseBody := bytes.NewBuffer(payload)
	//Leverage Go's HTTP Post function to make request
	resp, err := http.Post(ctx.Value("AUTH_MIDDLEMAN_LOGIN").(string), "application/json", responseBody)

	//Handle Error
	if err != nil {
		log.Printf("[Middleman Login] error 1 %v", err)
		return legacyResponse, err
	}
	defer resp.Body.Close()

	if resp.StatusCode != 200 {
		log.Printf("[Middleman Login] error 2 %v", resp.Status)
		err = errors.New(fmt.Sprintf("Legacy Host Error Code %d (%s)", resp.StatusCode, resp.Status))
		return legacyResponse, err
	}
	//Read the response body
	body, err := ioutil.ReadAll(resp.Body)
	if err != nil {
		log.Printf("[Middleman Login] error 3 %v", err)
		return legacyResponse, err
	}
	json.Unmarshal(body, &legacyResponse)

	if strings.ToLower(legacyResponse.Status) != "success" {
		err = errors.New(legacyResponse.Message)
		log.Printf("[Middleman Login] error 4 %v", err)
	}
	return legacyResponse, err
}
