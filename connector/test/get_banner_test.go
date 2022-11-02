package test

import (
	"net/http"
	"net/url"
	"testing"

	"github.com/stretchr/testify/assert"
	"gitlab.com/distributed_lab/kit/kv"
	"gitlab.com/tokend/connectors/signed"
	connector2 "gitlab.com/tokend/nft-books/blob-svc/connector"
	"gitlab.com/tokend/nft-books/blob-svc/internal/config"
)

func TestGetBanner(t *testing.T) {
	cfg := config.New(kv.NewViperFile("test_config.yaml"))

	mocked := url.URL{}
	connector := connector2.NewConnector(signed.NewClient(http.DefaultClient, &mocked), *cfg.AWSConfig(), "mock")

	_, err := connector.GetReadableLink("wrong key")
	assert.NotNil(t, err, "should be err")

	link, err := connector.GetReadableLink("2.png")
	assert.Nil(t, err, "should find the link")

	t.Logf("Link: %s", link.Data.Attributes.Url)
}
