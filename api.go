package peanuts

import (
	"net/http"
)

const (
	AUTHENTICATE_URL       = "https://pnut.io/oauth/authenticate"
	API_BASE_URL           = "https://api.pnut.io/v0/"
	OAUTH_ACCESS_TOKEN_API = API_BASE_URL + "oauth/access_token"
	POST_API               = API_BASE_URL + "posts"
	STREAM_BASE_URL        = POST_API + "/" + "streams"
	STREAM_ME_API          = STREAM_BASE_URL + "/me"
	STREAM_UNIFIED_API     = STREAM_BASE_URL + "/unified"
	STREAM_GLOBAL_API      = STREAM_BASE_URL + "/global"
	STREAM_TAG_BASE_URL    = POST_API + "/" + "tag"
	USER_API               = API_BASE_URL + "users"
	USER_ME_API            = USER_API + "/me"
)

type Api struct {
	accessToken string
	HttpClient  *http.Client
}
