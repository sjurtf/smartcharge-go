package smartcharge

import (
	"bytes"
	"encoding/json"
	"errors"
	"fmt"
	"io"
	"io/ioutil"
	"net/http"
	"net/url"
	"strings"
)

const (
	LibraryVersion = "0.1.0"
	DefaultBaseUrl = "https://api.smartcharge.io/"
	userAgent      = "smartcharge-go/" + LibraryVersion

	DefaultAppId     = "40a38b82-d7a5-4c0f-bd25-9736bc20a4d2"
	DefaultAppToken  = ""
	DefaultPushState = "unknown"
)

type Client struct {
	httpClient *http.Client

	BaseURL *url.URL

	UserAgent string

	Auth *Authentication

	ChargePoint *ChargePointService
	Session     *SessionService
	Invoice     *InvoiceService
}

func NewClient(httpClient *http.Client) *Client {
	if httpClient == nil {
		httpClient = http.DefaultClient
	}

	baseUrl, _ := url.Parse(DefaultBaseUrl)

	c := &Client{
		httpClient: httpClient,
		BaseURL:    baseUrl,
		UserAgent:  userAgent,
	}

	c.ChargePoint = &ChargePointService{client: c}
	c.Session = &SessionService{client: c}
	c.Invoice = &InvoiceService{client: c}

	return c
}

func (c *Client) SetAuthentication(auth *Authentication) {
	c.Auth = auth
}

func (c *Client) NewRequest(method string, urlStr string, body interface{}) (*http.Request, error) {
	rel, err := url.Parse(urlStr)
	if err != nil {
		return nil, err
	}

	u := c.BaseURL.ResolveReference(rel)

	bodyReader, ok := body.(io.Reader)
	if !ok && body != nil {
		buf := &bytes.Buffer{}
		err := json.NewEncoder(buf).Encode(body)
		if err != nil {
			return nil, err
		}
		bodyReader = buf
	}

	req, err := http.NewRequest(method, u.String(), bodyReader)
	if err != nil {
		return nil, err
	}

	if c.UserAgent != "" {
		req.Header.Add("User-Agent", c.UserAgent)
	}

	if c.Auth != nil {
		req.Header.Add("Authorization", "Bearer "+c.Auth.AccessToken)
	}

	return req, nil
}

func (c *Client) Do(req *http.Request, v interface{}) (*http.Response, error) {
	resp, err := c.httpClient.Do(req)
	if err != nil {
		return nil, err
	}

	if v != nil {
		defer resp.Body.Close()
	}

	err = CheckResponse(resp)
	if err != nil {
		if v == nil {
			_ = resp.Body.Close()
		}
		return resp, err
	}

	if v != nil {
		if w, ok := v.(io.Writer); ok {
			_, _ = io.Copy(w, resp.Body)
		} else {
			err = json.NewDecoder(resp.Body).Decode(v)
		}
	}
	return resp, err
}

// TODO handle 401
func CheckResponse(r *http.Response) error {
	c := r.StatusCode
	if 200 <= c && c <= 299 {
		return nil
	}

	errBody := ""
	if data, err := ioutil.ReadAll(r.Body); err == nil {
		errBody = strings.TrimSpace(string(data))
	}

	errMsg := fmt.Sprintf("HTTP code %v: %q: ", c, r.Status)
	if errBody == "" {
		errMsg += "no response body"
	} else {
		errMsg += fmt.Sprintf("response body: %q", errBody)
	}

	return errors.New(errMsg)
}
