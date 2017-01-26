package peanuts

import (
	"encoding/json"
	"fmt"
	"io/ioutil"
	"net/http"
	"net/url"
	"strconv"
	"strings"
)

type Client struct {
	clientId            string
	clientSecret        string
	passwordGrantSecret string
	queryQueue          chan query
	Api                 Api
}

func NewClient(clientId string, clientSecret string) *Client {
	queue := make(chan query)
	client := &Client{clientId: clientId, clientSecret: clientSecret, queryQueue: queue}
	client.initialize()
	go client.throttledQuery()
	return client
}

type query struct {
	url         string
	form        url.Values
	data        interface{}
	method      string
	response_ch chan response
}

type response struct {
	data interface{}
	err  error
}

func (c *Client) initialize() {
	c.Api = *&Api{
		accessToken: "",
		HttpClient:  http.DefaultClient,
	}
}

// Generate authorization url
// https://pnut.io/docs/authentication/web-flows
func (c *Client) AuthURL(redirectURI string, scope []string, responseType string) string {
	return AUTHENTICATE_URL + "?client_id=" + c.clientId + "&redirect_uri=" + redirectURI + "&scope=" + strings.Join(scope, "%20") + "&response_type=" + responseType
}

// Set password grant secret
// https://pnut.io/docs/authentication/password-flow
func (c *Client) SetPasswordGrantSecret(passwordGrantSecret string) {
	c.passwordGrantSecret = passwordGrantSecret
}

// Set access token
// https://pnut.io/docs/authentication/web-flows
// https://pnut.io/docs/authentication/password-flow
func (c *Client) SetAccessToken(accessToken string) {
	c.Api.accessToken = accessToken
}

type StreamMeta struct {
	More  bool   `json:"more"`
	MaxId string `json:"max_id"`
	MinId string `json:"min_id"`
}

type Meta struct {
	*StreamMeta
	Code         int    `json:"code"`
	Error        string `json:"error"`
	ErrorMessage string `json:"error_message"`
}

type CommonResponse struct {
	Meta Meta `json:"meta"`
}

func decodeResponse(res *http.Response, data interface{}) error {
	b, err := ioutil.ReadAll(res.Body)
	if err != nil {
		return err
	}
	err = json.Unmarshal(b, data)
	if err != nil {
		return err
	}
	if res.StatusCode < 200 || res.StatusCode >= 300 {
		common := &CommonResponse{}
		err = json.Unmarshal(b, common)
		if err != nil {
			return err
		}
		return fmt.Errorf(strconv.Itoa(res.StatusCode) + ": " + common.Meta.ErrorMessage)
	}

	return nil
}

func (c *Client) execQuery(url string, form url.Values, data interface{}, method string) error {
	req, err := http.NewRequest(
		method,
		url,
		strings.NewReader(form.Encode()),
	)
	if err != nil {
		return err
	}
	req.Header.Set("Content-Type", "application/x-www-form-urlencoded")
	if c.Api.accessToken != "" {
		req.Header.Set("Authorization", "Bearer "+c.Api.accessToken)
	}
	client := &http.Client{}
	resp, err := client.Do(req)
	if err != nil {
		return err
	}
	defer resp.Body.Close()

	return decodeResponse(resp, data)
}

func (c *Client) throttledQuery() {
	for q := range c.queryQueue {
		url := q.url
		form := q.form
		data := q.data
		method := q.method

		response_ch := q.response_ch

		err := c.execQuery(url, form, data, method)

		response_ch <- response{data, err}
	}
}
