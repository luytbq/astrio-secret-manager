package common

import (
	"bytes"
	"encoding/json"
	"errors"
	"io"
	"log"
	"net/http"

	"github.com/luytbq/astrio-secret-manager/config"
)

func RequestAAS(method, uri string, body any, header *http.Header) (statusCode int, resBytes []byte, err error) {
	statusCode, resBytes, err = httpRequest(method, config.App.AAS_URL+uri, body, header)

	if err != nil {
		statusCode = http.StatusUnauthorized
		err = errors.New("unauthorized")
		return
	}

	if statusCode == http.StatusBadRequest {
		statusCode = http.StatusUnauthorized
		err = errors.New("unauthorized")
		return
	}

	if statusCode == http.StatusInternalServerError {
		statusCode = http.StatusUnauthorized
		err = errors.New("unauthorized")
		return
	}

	if statusCode != http.StatusOK {
		statusCode = http.StatusUnauthorized
		err = errors.New("unauthorized")
		return
	}

	return
}

func httpRequest(method, url string, body any, header *http.Header) (statusCode int, resBytes []byte, err error) {
	var reqBody io.Reader
	if body != nil {
		jsonData, err := json.Marshal(body)

		if err != nil {
			log.Printf("Error creating body: %v", err)
			return 500, nil, err
		}

		reqBody = bytes.NewBuffer(jsonData)
	}

	req, err := http.NewRequest(method, url, reqBody)

	if err != nil {
		log.Printf("Error creating request: %v", err)
		return
	}

	if header != nil {
		req.Header = *header
	}

	client := http.Client{}

	res, err := client.Do(req)
	statusCode = res.StatusCode

	if err != nil {
		log.Printf("Error making request: %v", err)
		return
	}

	defer res.Body.Close()

	resBytes, err = io.ReadAll(res.Body)

	if err != nil {
		log.Printf("Error reading response body: %v", err)
		return
	}

	return statusCode, resBytes, nil
}
