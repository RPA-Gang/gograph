package gograph

import "net/http"

type IGraphClient interface {
	Client() *http.Client
	SetClient(*http.Client)
	BaseUrl() string
	SetBaseUrl(string)
}

type GraphClient struct {
	client  *http.Client
	baseUrl string
}

func (c *GraphClient) Client() *http.Client {
	return c.client
}
func (c *GraphClient) SetClient(client *http.Client) {
	c.client = client
}
func (c *GraphClient) BaseUrl() string {
	return c.baseUrl
}
func (c *GraphClient) SetBaseUrl(baseUrl string) {
	c.baseUrl = baseUrl
}

func NewGraphClient(client *http.Client, baseUrl string) IGraphClient {
	return &GraphClient{
		client:  client,
		baseUrl: baseUrl,
	}
}
