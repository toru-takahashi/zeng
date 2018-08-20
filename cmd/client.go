package cmd

import (
	"context"
	"encoding/json"
	"fmt"
	"io"
	"net/http"
	"net/url"
	"path"

	"github.com/pkg/errors"
	"github.com/spf13/viper"
)

// Client is a object for Zendesk Auth
type Client struct {
	EndpointURL *url.URL
	HTTPClient  *http.Client
	UserAgent   string
	SubDomain   string
	Email       string
	Password    string
	Locale      string
	APIKEY      string
}

func newClient(endpointURL string, httpClient *http.Client, userAgent string, config *Configuration) (*Client, error) {
	parsedURL, err := url.ParseRequestURI(endpointURL)
	if err != nil {
		return nil, errors.Wrapf(err, "failed to parse url: %s", endpointURL)
	}
	client := &Client{}
	if len(config.password) > 0 {
		client = &Client{
			EndpointURL: parsedURL,
			HTTPClient:  httpClient,
			UserAgent:   userAgent,
			SubDomain:   config.subdomain,
			Email:       config.email,
			Password:    config.password,
			Locale:      config.locale,
		}
	} else if len(config.apikey) > 0 {
		client = &Client{
			EndpointURL: parsedURL,
			HTTPClient:  httpClient,
			UserAgent:   userAgent,
			SubDomain:   config.subdomain,
			Email:       config.email,
			APIKEY:      config.apikey,
			Locale:      config.locale,
		}
	}
	return client, nil
}

func (client *Client) newRequest(ctx context.Context, method string, subPath string, body io.Reader) (*http.Request, error) {
	endpointURL := *client.EndpointURL
	endpointURL.Path = path.Join(client.EndpointURL.Path, subPath)

	req, err := http.NewRequest(method, endpointURL.String(), body)

	if err != nil {
		return nil, err
	}

	req = req.WithContext(ctx)

	req.Header.Add("Content-Type", "application/json; charset=utf-8")
	req.Header.Add("User-Agent", client.UserAgent)
	if len(client.Password) > 0 {
		req.SetBasicAuth(client.Email, client.Password)
	} else if len(client.APIKEY) > 0 {
		req.SetBasicAuth(client.Email+"/token", client.APIKEY)
	}
	return req, nil
}

func decodeBody(resp *http.Response, out interface{}) error {
	defer resp.Body.Close()
	// To check a response
	// body, _ := ioutil.ReadAll(resp.Body)
	// fmt.Println(string(body))
	//return nil
	decoder := json.NewDecoder(resp.Body)
	return decoder.Decode(out)
}

func (client *Client) previewPath(subPath string) string {
	return fmt.Sprintf("https://%s.zendesk.com/%s", viper.GetString("subdomain"), subPath)
}
