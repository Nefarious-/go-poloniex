package poloniex

import (
	"crypto/hmac"
	"crypto/sha512"
	"encoding/hex"
	"encoding/json"
	"fmt"
	"hash"
	"io/ioutil"
	"net/http"
	"net/url"
	"strconv"
	"strings"
	"time"
)

const (
	publicURL    = "https://poloniex.com/public"
	tradeURL     = "https://poloniex.com/tradingApi"
	websocketURL = "wss://api.poloniex.com"
)

// Client exposes the various methods used on the exchange.
type Client struct {
	key, secret string
	httpClient  *http.Client
	hash        *hash.Hash
}

// NewClient returns a Client object.
func NewClient(key, secret string) *Client {
	h := hmac.New(sha512.New, []byte(secret))
	return &Client{key, secret, new(http.Client), &h}
}

func (c *Client) req(method, command string, val url.Values, obj interface{}) error {
	var addr string
	switch method {
	case "GET":
		addr = fmt.Sprintf("%s?command=%s&%s", publicURL, command, val.Encode())
	case "POST":
		addr = tradeURL
	}
	r, err := http.NewRequest(method, addr, nil)
	if err != nil {
		return err
	}
	if method == "POST" {
		val.Add("nonce", strconv.Itoa(int(time.Now().UnixNano())))
		val.Add("command", command)
		data := val.Encode()
		r.Body = ioutil.NopCloser(strings.NewReader(data))
		r.Header.Set("Key", c.key)
		r.Header.Set("Sign", c.sign(data))
		r.Header.Set("Content-Type", "application/x-www-form-urlencoded")
		r.ContentLength = int64(len(data))
	}
	resp, err := c.httpClient.Do(r)
	if err != nil {
		return err
	}
	defer resp.Body.Close()
	return json.NewDecoder(resp.Body).Decode(obj)
}

func (c *Client) sign(post string) (sig string) {
	(*c.hash).Write([]byte(post))
	sig = hex.EncodeToString((*c.hash).Sum(nil))
	(*c.hash).Reset()
	return
}
