package peanuts

import (
	"net/http"
)

const (
	AUTHENTICATE_URL       = "https://pnut.io/oauth/authenticate"
	API_BASE_URL           = "https://api.pnut.io/v0/"
	OAUTH_ACCESS_TOKEN_API = API_BASE_URL + "oauth/access_token"
	POST_API               = API_BASE_URL + "posts"
)

type Api struct {
	accessToken          string
	ReturnRateLimitError bool
	HttpClient           *http.Client
}
