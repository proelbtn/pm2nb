package netbox

import (
	"context"
	"crypto/tls"
	"encoding/json"
	"errors"
	"fmt"
	"io"
	"log"
	"net/http"
	"net/url"
	"path"
	"strings"
)

type Credential struct {
	Token string
}

type Client struct {
	baseUrl     *url.URL
	httpClient  *http.Client
	logger      *log.Logger
	cred *Credential
}

func NewClient(baseUrl, token string, insecure bool) (*Client, error) {
	u, err := url.Parse(baseUrl)
	if err != nil {
		return nil, err
	}

	transport := &http.Transport{}
	if u.Scheme == "https" && insecure {
		transport.TLSClientConfig = &tls.Config{
			InsecureSkipVerify: insecure,
		}
	}

	return &Client{
		baseUrl: u,
		httpClient: &http.Client{
			Transport: transport,
		},
		cred: &Credential{
			Token: token,
		},
	}, nil
}

func (c *Client) newRequest(ctx context.Context, method, endpoint string, body io.Reader) (*http.Request, error) {
	u := *c.baseUrl
	u.Path = path.Join(c.baseUrl.Path, endpoint) + "/"

	req, err := http.NewRequest(method, u.String(), body)
	if err != nil {
		return nil, err
	}

	if c.cred != nil {
		req.Header.Add("Authorization", fmt.Sprintf("Token %s", c.cred.Token))
	}
	req.WithContext(ctx)
	return req, nil
}

func (c *Client) decodeBody(res *http.Response, data interface{}) error {
	defer res.Body.Close()

	if !(200 <= res.StatusCode && res.StatusCode < 300) {
		content := c.dumpBody(res)
		return errors.New(fmt.Sprintf(
			"failed: %d, %s",
			res.StatusCode, content,
		))
	}

	err := json.NewDecoder(res.Body).Decode(&data)
	if err != nil {
		return err
	}

	return err
}

func (c *Client) dumpBody(res *http.Response) string {
	buf := make([]byte, 32768)
	res.Body.Read(buf)
	return string(buf)
}

func (c *Client) DoGet(ctx context.Context, endpoint string, body url.Values, data interface{}) error {
	req, err := c.newRequest(ctx, http.MethodGet, endpoint, nil)
	if err != nil {
		return err
	}

	req.URL.RawQuery = body.Encode()

	res, err := c.httpClient.Do(req)
	if err != nil {
		return nil
	}

	err = c.decodeBody(res, &data)
	if err != nil {
		return err
	}

	return nil
}

func (c *Client) DoPost(ctx context.Context, endpoint string, body interface{}, data interface{}) error {
	content, err := json.Marshal(body)
	if err != nil {
		return err
	}

	req, err := c.newRequest(ctx, http.MethodPost, endpoint, strings.NewReader(string(content)))
	if err != nil {
		return err
	}

	req.Header.Add("Content-Type", "application/json")

	res, err := c.httpClient.Do(req)
	if err != nil {
		return nil
	}

	err = c.decodeBody(res, &data)
	if err != nil {
		return err
	}

	return nil
}

func (c *Client) DoDelete(ctx context.Context, endpoint string) error {
	req, err := c.newRequest(ctx, http.MethodDelete, endpoint, nil)
	if err != nil {
		return err
	}

	res, err := c.httpClient.Do(req)
	if err != nil {
		return nil
	}

	if !(200 <= res.StatusCode && res.StatusCode < 300) {
		content := c.dumpBody(res)
		return errors.New(fmt.Sprintf(
			"failed: %d, %s",
			res.StatusCode, content,
		))
	}

	return nil
}

func (c *Client) DoPatch(ctx context.Context, endpoint string, body interface{}, data interface{}) error {
	content, err := json.Marshal(body)
	if err != nil {
		return err
	}

	req, err := c.newRequest(ctx, http.MethodPatch, endpoint, strings.NewReader(string(content)))
	if err != nil {
		return err
	}

	req.Header.Add("Content-Type", "application/json")

	res, err := c.httpClient.Do(req)
	if err != nil {
		return nil
	}

	err = c.decodeBody(res, &data)
	if err != nil {
		return err
	}

	return nil
}
