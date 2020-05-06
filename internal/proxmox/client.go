package proxmox

import (
	"context"
	"crypto/tls"
	"encoding/json"
	"errors"
	"fmt"
	"log"
	"net/http"
	"net/url"
	"path"
	"strings"
)

type Credential struct {
	Username string
	CSRFToken string
	Ticket string
}

type Client struct {
	baseUrl     *url.URL
	httpClient  *http.Client
	logger      *log.Logger
	cred *Credential
}

type Ticket struct {
	Data struct {
		Username string `json:"username"`
		CSRFPreventionToken string `json:"CSRFPreventionToken"`
		ClusterName string `json:"clustername"`
		Ticket string `json:"ticket"`
	} `json:"data"`
}

func NewClient(baseUrl string, insecure bool) (*Client, error) {
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
		cred: nil,
	}, nil
}

func (c *Client) newRequest(ctx context.Context, method, endpoint string, body url.Values) (*http.Request, error) {
	u := *c.baseUrl
	u.Path = path.Join(c.baseUrl.Path, endpoint)

	bodyContent := ""
	if method == http.MethodPost && body != nil {
		bodyContent = body.Encode()
	}

	req, err := http.NewRequest(method, u.String(), strings.NewReader(bodyContent))
	if err != nil {
		return nil, err
	}

	if method == http.MethodGet && body != nil {
		req.URL.RawQuery = body.Encode()
	}

	if c.cred != nil {
		req.Header.Add("CSRFPreventionToken", c.cred.CSRFToken)
		req.AddCookie(&http.Cookie{
			Name: "PVEAuthCookie",
			Value: c.cred.Ticket,
		})
	}
	req.WithContext(ctx)
	return req, nil
}

func (c *Client) decodeBody(res *http.Response, data interface{}) error {
	defer res.Body.Close()

	if res.StatusCode != 200 {
		buf := make([]byte, 4096)
		res.Body.Read(buf)
		return errors.New(fmt.Sprintf(
			"failed: %d, %s",
			res.StatusCode, string(buf),
		))
	}

	err := json.NewDecoder(res.Body).Decode(&data)
	if err != nil {
		return err
	}

	return err
}

func (c *Client) dumpBody(res *http.Response) {
	buf := make([]byte, 32768)
	res.Body.Read(buf)
	fmt.Println(string(buf))
}

func (c *Client) Login(ctx context.Context, username, password string) error {
	endpoint := "/api2/json/access/ticket"

	body := url.Values{}
	body.Add("username", username)
	body.Add("password", password)

	req, err := c.newRequest(ctx, http.MethodPost, endpoint, body)
	if err != nil {
		return err
	}

	res, err := c.httpClient.Do(req)
	if err != nil {
		return nil
	}

	data := Ticket{}
	err = c.decodeBody(res, &data)
	if err != nil {
		return err
	}

	c.cred = &Credential{
		Username: data.Data.Username,
		CSRFToken: data.Data.CSRFPreventionToken,
		Ticket: data.Data.Ticket,
	}

	return err
}
