package internal

import (
	"bytes"
	"encoding/json"
	"fmt"
	"io/ioutil"
	"net/http"
	"time"
)

type houndClient struct {
	token   string
	baseURL string
}

type createEndpointRequest struct {
	Endpoint string `json:"endpoint"`
	Protocol string `json:"protocol"`
}

type updateEndpointRequest struct {
	ID       string `json:"id"`
	Endpoint string `json:"endpoint"`
	Protocol string `json:"protocol"`
}

type endpointResponse struct {
	Endpoint string `json:"endpoint"`
	Protocol string `json:"protocol"`
	ID       string `json:"id"`
}

type listEndpointsResponse struct {
	Endpoints []endpointResponse `json:"endpoints"`
}

func (h houndClient) listEndpoints() (listEndpointsResponse, error) {
	req, err := http.NewRequest("GET", h.baseURL, nil)
	if err != nil {
		return listEndpointsResponse{}, err
	}
	req.Header.Add("Accept", "application/json")
	req.Header.Add("Authorization", fmt.Sprintf("Bearer %s", h.token))

	client := &http.Client{Timeout: time.Second * 10}
	resp, err := client.Do(req)
	if err != nil {
		return listEndpointsResponse{}, err
	}

	if resp.StatusCode != 200 {
		return listEndpointsResponse{}, fmt.Errorf("request failed: %d", resp.StatusCode)
	}

	body, err := ioutil.ReadAll(resp.Body)
	if err != nil {
		return listEndpointsResponse{}, err
	}
	if err = resp.Body.Close(); err != nil {
		return listEndpointsResponse{}, err
	}

	var respBody listEndpointsResponse
	if err = json.Unmarshal([]byte(body), &respBody); err != nil {
		return listEndpointsResponse{}, err
	}
	return respBody, nil
}

func (h houndClient) createEndpoint(endpoint, protocol string) (endpointResponse, error) {
	reqBody, err := json.Marshal(createEndpointRequest{Endpoint: endpoint, Protocol: protocol})
	if err != nil {
		return endpointResponse{}, err
	}

	req, err := http.NewRequest("POST", h.baseURL, bytes.NewBuffer(reqBody))
	if err != nil {
		return endpointResponse{}, err
	}
	req.Header.Add("Content-Type", "application/json")
	req.Header.Add("Authorization", fmt.Sprintf("Bearer %s", h.token))

	client := &http.Client{Timeout: time.Second * 10}
	resp, err := client.Do(req)
	if err != nil {
		return endpointResponse{}, err
	}

	if resp.StatusCode != 200 {
		return endpointResponse{}, fmt.Errorf("request failed: %d", resp.StatusCode)
	}

	body, err := ioutil.ReadAll(resp.Body)
	if err != nil {
		return endpointResponse{}, err
	}
	if err = resp.Body.Close(); err != nil {
		return endpointResponse{}, err
	}

	var createEndpointResponse endpointResponse
	if err = json.Unmarshal([]byte(body), &createEndpointResponse); err != nil {
		return endpointResponse{}, err
	}
	return createEndpointResponse, nil
}

func (h houndClient) getEndpoint(id string) (endpointResponse, error) {
	req, err := http.NewRequest("GET", fmt.Sprintf("%s/%s", h.baseURL, id), nil)
	if err != nil {
		return endpointResponse{}, err
	}
	req.Header.Add("Accept", "application/json")
	req.Header.Add("Authorization", fmt.Sprintf("Bearer %s", h.token))

	client := &http.Client{Timeout: time.Second * 10}
	resp, err := client.Do(req)
	if err != nil {
		return endpointResponse{}, err
	}

	if resp.StatusCode != 200 {
		return endpointResponse{}, fmt.Errorf("request failed: %d", resp.StatusCode)
	}

	body, err := ioutil.ReadAll(resp.Body)
	if err != nil {
		return endpointResponse{}, err
	}
	if err = resp.Body.Close(); err != nil {
		return endpointResponse{}, err
	}

	var getEndpointResponse endpointResponse
	if err = json.Unmarshal([]byte(body), &getEndpointResponse); err != nil {
		return endpointResponse{}, err
	}
	return getEndpointResponse, nil
}

func (h houndClient) updateEndpoint(id, endpoint, protocol string) (endpointResponse, error) {
	reqBody, err := json.Marshal(updateEndpointRequest{Endpoint: endpoint, Protocol: protocol})
	if err != nil {
		return endpointResponse{}, err
	}

	req, err := http.NewRequest("POST", fmt.Sprintf("%s/%s", h.baseURL, id), bytes.NewBuffer(reqBody))
	if err != nil {
		return endpointResponse{}, err
	}
	req.Header.Add("Accept", "application/json")
	req.Header.Add("Authorization", fmt.Sprintf("Bearer %s", h.token))
	req.Header.Add("Content-Type", "application/json")

	client := &http.Client{Timeout: time.Second * 10}
	resp, err := client.Do(req)
	if err != nil {
		return endpointResponse{}, err
	}

	if resp.StatusCode != 200 {
		return endpointResponse{}, fmt.Errorf("request failed: %d", resp.StatusCode)
	}

	body, err := ioutil.ReadAll(resp.Body)
	if err != nil {
		return endpointResponse{}, err
	}
	if err = resp.Body.Close(); err != nil {
		return endpointResponse{}, err
	}

	var updateEndpointResponse endpointResponse
	if err = json.Unmarshal([]byte(body), &updateEndpointResponse); err != nil {
		return endpointResponse{}, err
	}
	return updateEndpointResponse, nil
}

func (h houndClient) deleteEndpoint(id string) error {
	req, err := http.NewRequest("DELETE", fmt.Sprintf("%s/%s", h.baseURL, id), nil)
	if err != nil {
		return err
	}
	req.Header.Add("Accept", "application/json")
	req.Header.Add("Authorization", fmt.Sprintf("Bearer %s", h.token))

	client := &http.Client{Timeout: time.Second * 10}
	resp, err := client.Do(req)
	if err != nil {
		return err
	}

	if resp.StatusCode != 200 {
		return err
	}
	return nil
}
