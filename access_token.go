package peanuts

import (
	"net/url"
)

type App struct {
	Id   string `json:"id"`
	Link string `json:"link"`
	Name string `json:"name"`
}

type Token struct {
	App      App      `json:"app"`
	Scopes   []string `json:"scopes"`
	User     User     `json:"user"`
	ClientId string   `json:"client_id"`
}

type AccessTokenResult struct {
	AccessToken string `json:"access_token"`
	Token       Token  `json:"token"`
	UserId      string `json:"user_id"`
	Username    string `json:"username"`
}

func (c *Client) AccessToken(code string, redirectURI string) (result AccessTokenResult, err error) {
	v := url.Values{}
	v.Set("client_id", c.clientId)
	v.Set("client_secret", c.clientSecret)
	v.Set("code", code)
	v.Set("redirect_uri", redirectURI)
	v.Set("grant_type", "authorization_code")
	response_ch := make(chan response)
	c.queryQueue <- query{url: OAUTH_ACCESS_TOKEN_API, form: v, data: &result, method: "POST", response_ch: response_ch}
	return result, (<-response_ch).err
}
