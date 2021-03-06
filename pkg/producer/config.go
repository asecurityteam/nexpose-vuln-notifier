package producer

import (
	"context"
	"net/url"
)

// ProducerConfig holds configuration required to send Nexpose assets
// to a queue via an HTTP Producer
type ProducerConfig struct {
	Endpoint string `description:"The scheme and URL of an HTTP producer."`
}

// Name is used by the settings library and will add a "HTTPPRODUCER"
// prefix to ProducerConfig environment variables
func (c *ProducerConfig) Name() string {
	return "HTTPProducer"
}

// ProducerConfigComponent satisfies the settings library Component
// API, and may be used by the settings.NewComponent function.
type ProducerConfigComponent struct{}

// Settings can be used to populate default values if there are any
func (*ProducerConfigComponent) Settings() *ProducerConfig { return &ProducerConfig{} }

// New constructs a AssetProducer from a config.
func (*ProducerConfigComponent) New(_ context.Context, c *ProducerConfig) (*AssetProducer, error) {
	endpoint, err := url.Parse(c.Endpoint)
	if err != nil {
		return nil, err
	}

	return &AssetProducer{Endpoint: endpoint}, nil
}
