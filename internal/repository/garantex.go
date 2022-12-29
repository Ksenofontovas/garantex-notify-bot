package repository

import (
	"encoding/json"
	"fmt"
	"io/ioutil"
	"net/http"
	"tg"
	"time"
)

const (
	URL = "api/v2/depth"
)

type Client struct {
	host       string
	httpClient *http.Client
}

func NewClient(host string, timeout time.Duration) *Client {
	client := &http.Client{
		Timeout: timeout,
	}
	return &Client{
		host:       host,
		httpClient: client,
	}
}

func (c *Client) do(method, endpoint string, params map[string]string) (*http.Response, error) {
	baseURL := fmt.Sprintf("%s/%s", c.host, endpoint)
	req, err := http.NewRequest(method, baseURL, nil)
	if err != nil {
		return nil, err
	}
	req.Header.Add("Content-Type", "application/json")
	q := req.URL.Query()
	for key, val := range params {
		q.Set(key, val)
	}
	req.URL.RawQuery = q.Encode()
	return c.httpClient.Do(req)
}

func (c *Client) GetDepth() (resp tg.DepthResponse, err error) {
	pair := make(map[string]string)
	pair["market"] = "usdtrub"
	res, err := c.do(http.MethodGet, URL, pair)
	if err != nil {
		return
	}

	defer res.Body.Close()

	body, err := ioutil.ReadAll(res.Body)
	if err != nil {
		return resp, err
	}

	if err = json.Unmarshal(body, &resp); err != nil {
		return resp, err
	}
	return
}
