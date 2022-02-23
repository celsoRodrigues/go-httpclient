package gohttp

import (
	"errors"
	"io"
	"net/http"
)

type HttpClient interface {
	Get(url string, headers http.Header) (*http.Response, error)
	Post(url string, headers http.Header, body io.Reader) (*http.Response, error)
	Put(url string, headers http.Header, body io.Reader) (*http.Response, error)
	Patch(url string, headers http.Header, body io.Reader) (*http.Response, error)
	Delete(url string, headers http.Header) (*http.Response, error)
}

type httpClient struct {
	Header http.Header
}

func (c *httpClient) Get(url string, headers http.Header) (*http.Response, error) {

	return c.do(http.MethodGet, url, headers, nil)

}

func (c *httpClient) Post(url string, headers http.Header, body io.Reader) (*http.Response, error) {
	return c.do(http.MethodPost, url, headers, body)

}

func (c *httpClient) Put(url string, headers http.Header, body io.Reader) (*http.Response, error) {
	return c.do(http.MethodPut, url, headers, body)

}

func (c *httpClient) Patch(url string, headers http.Header, body io.Reader) (*http.Response, error) {
	return c.do(http.MethodPatch, url, headers, body)

}

func (c *httpClient) Delete(url string, headers http.Header) (*http.Response, error) {
	return c.do(http.MethodDelete, url, headers, nil)
}

func (c *httpClient) do(method string, url string, header http.Header, body io.Reader) (*http.Response, error) {

	client := http.Client{}

	req, err := http.NewRequest(method, url, body)
	if err != nil {
		err := errors.New("unable to create request")
		return &http.Response{}, err
	}

	fullHeaders := c.getRequestHeaders(header)

	req.Header = fullHeaders

	return client.Do(req)

}

func (c *httpClient) getRequestHeaders(header http.Header) http.Header {
	fullHeaders := make(http.Header)

	for h, val := range c.Header {
		if len(val) > 0 {
			fullHeaders.Set(h, val[0])
		}
	}

	for h, val := range header {
		if len(val) > 0 {
			fullHeaders.Set(h, val[0])
		}
	}

	return fullHeaders

}

func New() HttpClient {
	c := &httpClient{}

	return c
}
