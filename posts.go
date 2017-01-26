package peanuts

import (
	"net/url"
)

type PostResult struct {
	*CommonResponse
	Data Post `json:"data"`
}

// Post post
// https://pnut.io/docs/resources/posts/lifecycle#post-posts
func (c *Client) Post(v url.Values) (result PostResult, err error) {
	response_ch := make(chan response)
	c.queryQueue <- query{url: POST_API, form: v, data: &result, method: "POST", response_ch: response_ch}
	return result, (<-response_ch).err
}
