package controllers

import (
	"bytes"
	"encoding/json"
	"io"
	"log"
	"net/http"
	"net/url"
	"time"
)

const DataApiUrl = "https://corona.lmao.ninja/"

type Client struct {
	BaseURL   *url.URL
	UserAgent string

	httpClient *http.Client
}

type All struct {
	Cases     int
	Deaths    int
	Recovered int
	Updated   int64
	UpdatedAt time.Time
	Active    int
}

//func (c *Client) ListUsers() ([]User, error) {
//	rel := &url.URL{Path: "/all"}
//	u := c.BaseURL.ResolveReference(rel)
//	req, err := http.NewRequest("GET", u.String(), nil)
//	if err != nil {
//		return nil, err
//	}
//	req.Header.Set("Accept", "application/json")
//	req.Header.Set("User-Agent", c.UserAgent)
//
//	resp, err := c.httpClient.Do(req)
//	if err != nil {
//		return nil, err
//	}
//	defer resp.Body.Close()
//	var allData []All
//	err = json.NewDecoder(resp.Body).Decode(&allData)
//	return allData, err
//}
func (c *Client) AllCountries() (*All, error) {
	req, err := c.newRequest("GET", "/all", nil)
	if err != nil {
		return nil, err
	}
	var allData All
	_, err = c.do(req, &allData)
	if err == nil {
		allData.UpdatedAt = time.Unix(0, allData.Updated * int64(time.Millisecond))
	}
	return &allData, err
}

func (c *Client) newRequest(method, path string, body interface{}) (*http.Request, error) {
	log.Printf("New request for %s", path)

	rel := &url.URL{Path: path}
	u := c.BaseURL.ResolveReference(rel)
	var buf io.ReadWriter
	if body != nil {
		buf = new(bytes.Buffer)
		err := json.NewEncoder(buf).Encode(body)
		if err != nil {
			return nil, err
		}
	}
	req, err := http.NewRequest(method, u.String(), buf)
	if err != nil {
		return nil, err
	}
	if body != nil {
		req.Header.Set("Content-Type", "application/json")
	}
	req.Header.Set("Accept", "application/json")
	req.Header.Set("User-Agent", c.UserAgent)
	return req, nil
}

func (c *Client) do(req *http.Request, v interface{}) (*http.Response, error) {
	resp, err := c.httpClient.Do(req)
	if err != nil {
		return nil, err
	}
	defer resp.Body.Close()
	//log.Printf("Body %s", resp.Body)
	err = json.NewDecoder(resp.Body).Decode(v)
	return resp, err
}
