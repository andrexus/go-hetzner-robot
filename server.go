package go_hetzner_robot

import (
	"context"
	"net/http"
)

const serverBasePath = "server"

// ServerService is an interface for interfacing with the order
// endpoints of the Hetzner Robot API
// See: https://robot.your-server.de/doc/webservice/en.html#get-order-server-product
type ServerService interface {
	ListServers(context.Context) ([]Server, *Response, error)
}

// ServerServiceOp handles communication with the server related methods of the
// Hetzner Robot API.
type ServerServiceOp struct {
	client *Client
}

var _ ServerService = &ServerServiceOp{}

//Server represents a Hetzner ordered Server
type Server struct {
	ServerIP     string `json:"server_ip"`
	ServerNumber int    `json:"server_number"`
	ServerName   string `json:"server_name"`
	Product      string `json:"product"`
	Dc           string `json:"dc"`
	Traffic      string `json:"traffic"`
	Flatrate     bool   `json:"flatrate"`
	Status       string `json:"status"`
	Throttled    bool   `json:"throttled"`
	Cancelled    bool   `json:"cancelled"`
	PaidUntil    string `json:"paid_until"`
}

type serversRoot struct {
	Server Server `json:"server"`
}

//ListServers queries data of all servers
func (s *ServerServiceOp) ListServers(ctx context.Context) ([]Server, *Response, error) {
	path := serverBasePath

	req, err := s.client.NewRequest(ctx, http.MethodGet, path, nil)
	if err != nil {
		return nil, nil, err
	}

	var results []Server
	var rootItems []serversRoot
	resp, err := s.client.Do(ctx, req, &rootItems)
	if err != nil {
		return nil, resp, err
	}

	for _, item := range rootItems {
		results = append(results, item.Server)
	}

	return results, resp, err
}
