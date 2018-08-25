package go_hetzner_robot

import (
	"context"
	"net/http"
)

const serverMarketProductBasePath = "order/server_market/product"

// OrderService is an interface for interfacing with the order
// endpoints of the Hetzner Robot API
// See: https://robot.your-server.de/doc/webservice/en.html#get-order-server-product
type OrderService interface {
	ListServerMarketProducts(context.Context, *ProductSearchRequest) ([]Product, *Response, error)
}

// OrderServiceOp handles communication with the order related methods of the
// Hetzner Robot API.
type OrderServiceOp struct {
	client *Client
}

var _ OrderService = &OrderServiceOp{}

// Product represents a Hetzner Robot Product
type Product struct {
	ID             int      `json:"id"`
	Name           string   `json:"name"`
	Description    []string `json:"description"`
	Traffic        string   `json:"traffic"`
	Dist           []string `json:"dist"`
	Arch           []int    `json:"arch"`
	Lang           []string `json:"lang"`
	CPU            string   `json:"cpu"`
	CPUBenchmark   int      `json:"cpu_benchmark"`
	MemorySize     int      `json:"memory_size"`
	HddSize        int      `json:"hdd_size"`
	HddText        string   `json:"hdd_text"`
	HddCount       int      `json:"hdd_count"`
	Datacenter     string   `json:"datacenter"`
	NetworkSpeed   string   `json:"network_speed"`
	Price          string   `json:"price"`
	PriceSetup     string   `json:"price_setup"`
	PriceVat       string   `json:"price_vat"`
	PriceSetupVat  string   `json:"price_setup_vat"`
	FixedPrice     bool     `json:"fixed_price"`
	NextReduce     int      `json:"next_reduce"`
	NextReduceDate string   `json:"next_reduce_date"`
}

type ProductSearchRequest struct {
	// CPU model name
	CPU string `url:"cpu,omitempty"`
	// Minimum CPU benchmark value
	MinCPUBenchmark string `url:"min_cpu_benchmark,omitempty"`
	// Maximum CPU benchmark value
	MaxCPUBenchmark string `url:"max_cpu_benchmark,omitempty"`
	// Minimum memory size in GB
	MinMemorySize string `url:"min_memory_size,omitempty"`
	// Maximum memory size in GB
	MaxMemorySize string `url:"max_memory_size,omitempty"`
	// Minimum drive size in GB
	MinHDDSize string `url:"min_hdd_size,omitempty"`
	// Maximum drive size in GB
	MaxHDDSize string `url:"max_hdd_size,omitempty"`
	// Minimum drive count
	MinHDDCount string `url:"min_hdd_count,omitempty"`
	// Maximum drive count
	MaxHDDCount string `url:"max_hdd_count,omitempty"`
	// Full text search
	Search string `url:"search,omitempty"`
	// Minimum monthly price
	MinPrice string `url:"min_price,omitempty"`
	// Maximum monthly price
	MaxPrice string `url:"max_price,omitempty"`
}

type productsRoot struct {
	Product Product `json:"product"`
}

//ListServerMarketProducts returns a product overview of currently offered server market products
func (s *OrderServiceOp) ListServerMarketProducts(ctx context.Context, opt *ProductSearchRequest) ([]Product, *Response, error) {
	path := serverMarketProductBasePath
	path, err := addOptions(path, opt)
	if err != nil {
		return nil, nil, err
	}

	req, err := s.client.NewRequest(ctx, http.MethodGet, path, nil)
	if err != nil {
		return nil, nil, err
	}

	var results []Product
	var rootItems []productsRoot
	resp, err := s.client.Do(ctx, req, &rootItems)
	if err != nil {
		return nil, resp, err
	}

	for _, item := range rootItems {
		results = append(results, item.Product)
	}

	return results, resp, err
}
