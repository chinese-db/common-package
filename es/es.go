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

func (e *EsClient) EsInit(Host string) {
	cfg := elasticsearch.Config{
		Addresses: []string{
			Host,
		},
		// ...
	}
	ES, err = elasticsearch.NewClient(cfg)
}
