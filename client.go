package poloniex

import (
	"crypto/hmac"
	"crypto/sha512"
	"encoding/base64"
	"hash"
	"net/http"
	"net/url"
)

const (
	publicURL    = "https://poloniex.com/public"
	tradeURL     = "https://poloniex.com/tradingApi"
	websocketURL = "wss://api.poloniex.com"
)

// Client is used to expose the various methods used on the exchange.
type Client struct {
	key        string
	secret     []byte
	httpClient *http.Client
	hash       *hash.Hash
}

// NewClient returns a Client object.
func NewClient(key string, secret []byte) *Client {
	h := hmac.New(sha512.New, secret)
	return &Client{key, secret, new(http.Client), &h}
}

func (c *Client) req(method string, val url.Values, params interface{}) (interface{}, error) {
	return nil, nil
}

func (c *Client) sign(post string) (sig string) {
	(*c.hash).Write([]byte(post))
	sig = base64.StdEncoding.EncodeToString((*c.hash).Sum(nil))
	(*c.hash).Reset()
	return
}
