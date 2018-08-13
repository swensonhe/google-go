package google

import (
	"encoding/json"
	"fmt"
	"io/ioutil"
	"net/http"
	"net/url"
)

const (
	tokenInfoEndpoint = "https://www.googleapis.com/oauth2/v3/tokeninfo"

	queryIdToken = "id_token"
)

type Client struct {
	*http.Client
}

func NewClient() *Client {
	return &Client{
		Client: &http.Client{},
	}
}

// GetTokenInfo returns the token info as a user.
func (c *Client) GetTokenInfo(token string) (*User, error) {
	u, err := url.Parse(tokenInfoEndpoint)
	if err != nil {
		return nil, ErrInternal
	}

	q := u.Query()
	q.Set(queryIdToken, token)
	u.RawQuery = q.Encode()

	req, err := http.NewRequest(http.MethodGet, u.String(), nil)
	if err != nil {
		fmt.Println(err)
		return nil, ErrInternal
	}

	resp, err := c.Client.Do(req)
	if err != nil {
		fmt.Println(err)
		return nil, ErrInvalidCredentials
	}
	defer resp.Body.Close()

	if resp.StatusCode >= 300 || resp.StatusCode < 200 {
		return nil, ErrInvalidCredentials
	}

	user, err := parseUser(resp)
	if err != nil {
		return nil, ErrInternal
	}

	return user, nil
}

func parseUser(resp *http.Response) (*User, error) {
	body, err := ioutil.ReadAll(resp.Body)
	if err != nil {
		return nil, err
	}

	var user User
	err = json.Unmarshal(body, &user)
	if err != nil {
		return nil, err
	}

	return &user, nil
}
