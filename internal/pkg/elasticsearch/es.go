package elasticsearch

import (
	"go-grpc-skeleton/config"

	"github.com/elastic/go-elasticsearch/v7"
)

func NewElasticsearchClient(cfg config.ElasticsearchConfig) (*elasticsearch.Client, error) {
	es, err := elasticsearch.NewClient(elasticsearch.Config{
		Addresses: []string{
			cfg.Addr,
		},
	})
	if err != nil {
		return nil, err
	}
	return es, nil
}
