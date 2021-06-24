package gofair

import (
	"encoding/json"
	"errors"
	"io"
	"io/ioutil"
	"log"
	"net/http"
	"strings"
)

func createUrl(endpoint string, method string) string {
	return endpoint + method
}

type Detail struct{}

type ErrorResponse struct {
	//b'{"faultcode":"Client","faultstring":"DSC-0018","detail":{}}'
	FaultCode   string  `json:"faultcode"`
	FaultString string  `json:"faultstring"`
	Detail      *Detail `json:"detail"`
}

func logError(data []byte) error {
	var errorResp ErrorResponse
	if err := json.Unmarshal(data, &errorResp); err != nil {
		return err
	}
	log.Println("Error:", string(data))
	return nil
}

func (c *Client) RequestJson(url string, params interface{}) ([]byte, error) {
	var body io.Reader
	if params != nil {
		bytes, err := json.Marshal(params)
		if err != nil {
			return nil, err
		}
		body = strings.NewReader(string(bytes))
	}

	req, err := http.NewRequest("POST", url, body)
	if err != nil {
		return nil, err
	}

	// set headers
	req.Header.Set("X-Application", c.config.AppKey)
	req.Header.Set("X-Authentication", c.session.SessionToken)
	req.Header.Set("Accept", "application/json")
	req.Header.Set("Content-Type", "application/json")
	req.Header.Set("Connection", "keep-alive")

	client := http.DefaultClient

	resp, err := client.Do(req)

	if err != nil {
		return nil, err
	}

	defer resp.Body.Close()
	data, err := ioutil.ReadAll(resp.Body)
	if err != nil {
		return nil, err
	}

	return data, nil
}

func (c *Client) Request(url string, params interface{}, v interface{}) error {
	//params.Locale = b.Client.config.Locale

	var body io.Reader
	if params != nil {
		bytes, err := json.Marshal(params)
		if err != nil {
			return err
		}
		body = strings.NewReader(string(bytes))
	}

	req, err := http.NewRequest("POST", url, body)
	if err != nil {
		return err
	}

	// set headers
	req.Header.Set("X-Application", c.config.AppKey)
	req.Header.Set("X-Authentication", c.session.SessionToken)
	req.Header.Set("Accept", "application/json")
	req.Header.Set("Content-Type", "application/json")
	req.Header.Set("Connection", "keep-alive")

	client := http.DefaultClient

	resp, err := client.Do(req)

	if err != nil {
		return err
	}

	defer resp.Body.Close()
	data, err := ioutil.ReadAll(resp.Body)
	if err != nil {
		return err
	}

	if resp.StatusCode != 200 {
		logError(data)
		return errors.New(resp.Status)
	} else {
		if err := json.Unmarshal(data, v); err != nil {
			return err
		}
	}

	return nil
}
