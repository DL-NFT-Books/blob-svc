package connector

import (
	"encoding/json"
	"net/url"

	"gitlab.com/distributed_lab/json-api-connector/base"
	"gitlab.com/distributed_lab/json-api-connector/client"
	"gitlab.com/tokend/nft-books/blob-svc/internal/config"
)

const (
	DocumentEndpoint = "/integrations/documents"
)

type Connector struct {
	base      *base.Connector
	client    client.Client
	token     string
	awsParams *config.AWSConfig
}

func NewConnector(client client.Client, awsConfig config.AWSConfig, token string) *Connector {
	return &Connector{
		client:    client,
		base:      base.NewConnector(client),
		token:     token,
		awsParams: &awsConfig,
	}
}

func (c *Connector) Get(endpoint *url.URL, dst interface{}) (err error) {
	response, err := c.base.Get(endpoint)
	if err != nil {
		return err
	}

	if response == nil || dst == nil {
		return nil
	}

	return json.Unmarshal(response, dst)
}
