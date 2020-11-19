package giphy

import (
	"net/url"
	"strconv"
	"strings"
)

type SearchQuery struct {
	Q        []string
	Limit    int
	Offset   int
	Rating   string
	Lang     string
	RandomID string
}

// Search returns a search response from the Giphy API
func (c *Client) Search(query SearchQuery) (*Search, error) {
	params := url.Values{}

	if len(query.Q) != 0 {
		params.Add("q", strings.Join(query.Q, " "))
	}

	if query.Limit != 0 {
		params.Add("limit", strconv.Itoa(query.Limit))
	} else {
		params.Add("limit", strconv.Itoa(c.Limit))
	}

	if query.Offset != 0 {
		params.Add("offset", strconv.Itoa(query.Offset))
	}

	if query.Rating != "" {
		params.Add("rating", query.Rating)
	}

	if query.Lang != "" {
		params.Add("lang", query.Lang)
	}

	if query.RandomID != "" {
		params.Add("random_id", query.RandomID)
	}

	u, err := url.Parse("/gifs/search")
	if err != nil {
		return nil, err
	}
	u.RawQuery = params.Encode()

	path := u.String()
	req, err := c.NewRequest(path)
	if err != nil {
		return &Search{}, err
	}

	var search Search
	if _, err = c.Do(req, &search); err != nil {
		return nil, err
	}

	return &search, nil
}
