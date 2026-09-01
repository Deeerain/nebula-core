package core

import (
	"encoding/json"
	"fmt"
	"log"
	"net/http"
)

type TelemtService interface {
	GetConnections() (map[string]any, error)
}

type telemtService struct {
	baseURL string
	client  *http.Client
}

func NewTelemtService(host string, port int) TelemtService {
	return &telemtService{
		baseURL: fmt.Sprintf("http://%s:%v", host, port),
		client:  http.DefaultClient,
	}
}

func (s *telemtService) GetConnections() (map[string]any, error) {
	req, err := http.NewRequest("GET", fmt.Sprintf("%s/v1/users", s.baseURL), nil)
	if err != nil {
		log.Fatalln("failed to create request: ", err.Error())
	}
	req.Header.Add("Content-type", "application/json")

	response, err := s.client.Do(req)
	if err != nil {
		return nil, fmt.Errorf("response error: %w", err)
	}

	var data map[string]any
	if err = json.NewDecoder(response.Body).Decode(&data); err != nil {
		return nil, fmt.Errorf("failed to decode body: %w", err)
	}

	switch response.StatusCode {
	case 200:
		return data, nil
	default:
		return nil, fmt.Errorf("telemt error: %v", data["error"])
	}
}
