package es

import (
	"github.com/elastic/go-elasticsearch/v7"
)

var (
	ES  *elasticsearch.Client
	err error
)

type EsClient struct {
	Host string
}

func EsNewClient(host string) *EsClient {
	return &EsClient{Host: host}
}

func (e *EsClient) EsInit() *elasticsearch.Client {
	cfg := elasticsearch.Config{
		Addresses: []string{
			e.Host,
		},
		// ...
	}
	ES, err = elasticsearch.NewClient(cfg)
	return ES
}
