package client

import (
	"bytes"
	"encoding/json"
	"fmt"
	"io"
	"net/http"

	"github.com/popescu-af/saas-y/pkg/connection"

	"github.com/popescu-af/optiopay/services/main-svc/pkg/exports"
)

// MainSvcClient is the structure that encompasses a main-svc client.
type MainSvcClient struct {
	connectionManager *connection.FullDuplexManager
	remoteAddress     string
}

// NewMainSvcClient creates a new instance of main-svc client.
func NewMainSvcClient(remoteAddress string) *MainSvcClient {
	return &MainSvcClient{
		connectionManager: connection.NewFullDuplexManager(),
		remoteAddress:     remoteAddress,
	}
}

// AddEmployee is the client function for POST '/add'.
func (c *MainSvcClient) AddEmployee(input *exports.AddInfo) error {
	var body io.Reader

	b, err := json.Marshal(input)
	if err != nil {
		return err
	}

	body = bytes.NewBuffer(b)

	url := "http://" + c.remoteAddress + fmt.Sprintf("/add")

	request, err := http.NewRequest("POST", url, body)

	response, err := http.DefaultClient.Do(request)
	if err != nil {
		return err
	}
	defer response.Body.Close()

	if response.StatusCode != http.StatusOK {
		return fmt.Errorf("POST %s failed with status code %d", url, response.StatusCode)
	}

	return nil
}

// RemoveEmployee is the client function for POST '/remove'.
func (c *MainSvcClient) RemoveEmployee(input *exports.RemoveInfo) error {
	var body io.Reader

	b, err := json.Marshal(input)
	if err != nil {
		return err
	}

	body = bytes.NewBuffer(b)

	url := "http://" + c.remoteAddress + fmt.Sprintf("/remove")

	request, err := http.NewRequest("POST", url, body)

	response, err := http.DefaultClient.Do(request)
	if err != nil {
		return err
	}
	defer response.Body.Close()

	if response.StatusCode != http.StatusOK {
		return fmt.Errorf("POST %s failed with status code %d", url, response.StatusCode)
	}

	return nil
}

// Manager is the client function for GET '/manager'.
func (c *MainSvcClient) Manager(firstEmployee string, secondEmployee string) (*exports.ManagerInfo, error) {
	var body io.Reader

	url := "http://" + c.remoteAddress + fmt.Sprintf("/manager")

	request, err := http.NewRequest("GET", url, body)
	request.Header.Set("first_employee", fmt.Sprintf("%s", firstEmployee))
	request.Header.Set("second_employee", fmt.Sprintf("%s", secondEmployee))

	response, err := http.DefaultClient.Do(request)
	if err != nil {
		return nil, err
	}
	defer response.Body.Close()

	if response.StatusCode != http.StatusOK {
		return nil, fmt.Errorf("GET %s failed with status code %d", url, response.StatusCode)
	}

	result := new(exports.ManagerInfo)
	if err := json.NewDecoder(response.Body).Decode(result); err != nil {
		return nil, err
	}

	return result, nil
}

// Hierarchy is the client function for GET '/hierarchy'.
func (c *MainSvcClient) Hierarchy() (*exports.HierarchyInfo, error) {
	var body io.Reader

	url := "http://" + c.remoteAddress + fmt.Sprintf("/hierarchy")

	request, err := http.NewRequest("GET", url, body)

	response, err := http.DefaultClient.Do(request)
	if err != nil {
		return nil, err
	}
	defer response.Body.Close()

	if response.StatusCode != http.StatusOK {
		return nil, fmt.Errorf("GET %s failed with status code %d", url, response.StatusCode)
	}

	result := new(exports.HierarchyInfo)
	if err := json.NewDecoder(response.Body).Decode(result); err != nil {
		return nil, err
	}

	return result, nil
}
