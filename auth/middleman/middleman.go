package middleman

import (
	"bytes"
	"context"
	"crypto/tls"
	"crypto/x509"
	"encoding/json"
	"errors"
	"fmt"
	"io/ioutil"
	"log"
	"net/http"
	"os"
	"strings"
)

// User domain model
type EncryptedVerifyResponse struct {
	Account string `json:"Account"`
}

type VerifyResponse struct {
	Account Account `json:"Account"`
}

type Account struct {
	ID       string `json:"AccountID"`
	Code     string `json:"BorrowerID"`
	Borrower struct {
		BorrowerID string `json:"BorrowerID"`
	} `json:"msborrower"`
	Email string `json:"email,omitempty"`
}

//LoginResponse legacy login response
type LoginResponse struct {
	Token   string            `json:"token,omitempty"`
	Status  string            `json:"status,omitempty"`
	Message string            `json:"message,omitempty"`
	Account map[string]string `json:"account,omitempty"`
}

func validateCert() (client *http.Client) {
	if _, err := os.Stat("ssl-cert/cert.crt"); os.IsNotExist(err) {
		tr := &http.Transport{
			TLSClientConfig: &tls.Config{InsecureSkipVerify: true},
		}
		client := &http.Client{Transport: tr}
		return client
	}
	caCert, err := ioutil.ReadFile("ssl-cert/cert.crt")
	if err != nil {
		log.Fatal(err)
	}
	caCertPool := x509.NewCertPool()
	caCertPool.AppendCertsFromPEM(caCert)

	client = &http.Client{
		Transport: &http.Transport{
			TLSClientConfig: &tls.Config{
				RootCAs: caCertPool,
			},
		},
	}
	return client
}

// Verify legacy token
func Verify(ctx context.Context, token string, permissions map[string]interface{}) (verifyResponse VerifyResponse, err error) {

	client := validateCert()

	middlemanURL := ctx.Value("URL_MIDDLEMAN_VERIFY").(string)
	req, err := http.NewRequest("GET", middlemanURL, nil)
	if err != nil {
		log.Printf("[Middleman Verify] error 1 %v", err)
		return verifyResponse, err
	}
	req.Header.Add("Authorization", token)
	resp, err := client.Do(req)
	if err != nil {
		log.Printf("[Middleman Verify] error 2 %v", err)
		return verifyResponse, err
	}
	defer resp.Body.Close()
	//Read the response body
	if resp.StatusCode != 200 {
		log.Printf("[Middleman Verify] error 3 %v", resp.Status)
		err = errors.New(fmt.Sprintf("Legacy Host Error Code %d (%s)", resp.StatusCode, resp.Status))
		return verifyResponse, err
	}

	var encrypted EncryptedVerifyResponse
	body, err := ioutil.ReadAll(resp.Body)
	if err != nil {
		log.Printf("[Middleman Verify] error 4 %v", err)
		return verifyResponse, err
	}
	json.Unmarshal(body, &encrypted)

	r := strings.NewReader(fmt.Sprintf(`{"message": "%s"}`, encrypted.Account))
	client = &http.Client{}

	decryptURL := ctx.Value("URL_MIDDLEMAN_DECRYPT").(string)
	if decryptURL == "" {
		decryptURL = "https://api.edufund.co.id/api/general/decrypt-account"
	}

	req, err = http.NewRequest("POST", decryptURL, r)
	if err != nil {
		log.Printf("[Middleman Verify] error 5 %v", err)
		return verifyResponse, err
	}
	req.Header.Add("Authorization", token)
	req.Header.Add("Content-Type", "application/json")
	resp, err = client.Do(req)
	if err != nil {
		log.Printf("[Middleman Verify] error 6 %v", err)
		return verifyResponse, err
	}
	defer resp.Body.Close()

	body, err = ioutil.ReadAll(resp.Body)
	if err != nil {
		log.Printf("[Middleman Verify] error 7 %v", err)
		return verifyResponse, err
	}

	json.Unmarshal(body, &verifyResponse)
	verifyResponse.Account.Code = verifyResponse.Account.Borrower.BorrowerID

	return verifyResponse, err
}

// Login legacy token
func Login(ctx context.Context, payload []byte, permissions map[string]interface{}) (loginResponse LoginResponse, err error) {
	client := validateCert()
	responseBody := bytes.NewBuffer(payload)
	//Leverage Go's HTTP Post function to make request
	middlemanURL := ctx.Value("URL_MIDDLEMAN_LOGIN").(string)
	resp, err := client.Post(middlemanURL, "application/json", responseBody)

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
