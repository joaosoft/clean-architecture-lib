package rest

import (
	"bytes"
	"encoding/json"
	"fmt"
	"io"
	"net/http"
	"reflect"
	"time"
)

type Server struct {
	Endpoint      string `yaml:"endpoint"`
	Timeout       int    `yaml:"timeout"`
	Authorization string `yaml:"authorization"`
}

type Adapter struct {
	Server *Server
}

func NewRestAdapter(server *Server) *Adapter {
	return &Adapter{
		Server: server,
	}
}

func (adapter *Adapter) getFullURL(path string) string {
	return fmt.Sprintf("%s/%s", adapter.Server.Endpoint, path)
}

func (adapter *Adapter) parseResponse(response *http.Response, result interface{}) error {
	defer func(Body io.ReadCloser) {
		_ = Body.Close()
	}(response.Body)

	body, err := io.ReadAll(response.Body)
	if err != nil {
		return err
	}

	if response.StatusCode >= 400 {
		// Handle error response based on specific REST server error format
		// You can customize this according to your REST server's response format
		var errResponse struct {
			Message string `json:"message"`
			Error   string `json:"error"`
			// Other error response fields
		}
		if err = json.Unmarshal(body, &errResponse); err != nil {
			return fmt.Errorf("API error: %s", body)
		}

		if errResponse.Message != "" {
			return fmt.Errorf("API error: %s", errResponse.Message)
		} else {
			return fmt.Errorf("API error: %s", errResponse.Error)
		}
	}

	if result != nil {
		if err = json.Unmarshal(body, result); err != nil {
			return err
		}
	}

	return nil
}

func (adapter *Adapter) call(method, path string, queryParams map[string]string, requestBody interface{}, result interface{}) (*http.Response, error) {
	if result != nil &&
		reflect.ValueOf(result).Kind() != reflect.Ptr {
		return nil, fmt.Errorf("error: result can't be nil and must be a pointer")
	}

	fullURL := adapter.getFullURL(path)

	client := &http.Client{
		Timeout: time.Duration(adapter.Server.Timeout) * time.Second,
	}

	// Marshal the request body to JSON
	bodyBuffer := bytes.NewBuffer(nil)

	if requestBody != nil {
		jsonBody, err := json.Marshal(requestBody)
		if err != nil {
			return nil, err
		}

		bodyBuffer = bytes.NewBuffer(jsonBody)
	}

	request, err := http.NewRequest(method, fullURL, bodyBuffer)
	if err != nil {
		return nil, err
	}

	// Set query parameters if any
	q := request.URL.Query()
	for key, value := range queryParams {
		q.Add(key, value)
	}
	request.URL.RawQuery = q.Encode()
	request.Header.Set("Content-Type", "application/json")

	// Set authorization header if provided
	if adapter.Server.Authorization != "" {
		request.Header.Set("Authorization", adapter.Server.Authorization)
	}

	response, err := client.Do(request)
	if err != nil {
		return nil, err
	}

	if err = adapter.parseResponse(response, result); err != nil {
		return nil, err
	}

	return response, nil
}

func (adapter *Adapter) Get(path string, queryParams map[string]string, result interface{}) error {
	_, err := adapter.call(http.MethodGet, path, queryParams, nil, result)
	if err != nil {
		return err
	}

	return nil
}

func (adapter *Adapter) Post(path string, queryParams map[string]string, requestBody interface{}, result interface{}) error {
	_, err := adapter.call(http.MethodPost, path, queryParams, requestBody, result)
	if err != nil {
		return err
	}

	return nil
}

func (adapter *Adapter) Put(path string, queryParams map[string]string, requestBody interface{}, result interface{}) error {
	_, err := adapter.call(http.MethodPut, path, queryParams, requestBody, result)
	if err != nil {
		return err
	}

	return nil
}

func (adapter *Adapter) Delete(path string, queryParams map[string]string, result interface{}) error {
	_, err := adapter.call(http.MethodDelete, path, queryParams, nil, result)
	if err != nil {
		return err
	}

	return nil
}
