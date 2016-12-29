// Copyright © 2016 Jon Arild Torresdal <jon@torresdal.net>
//
// Licensed under the Apache License, Version 2.0 (the "License");
// you may not use this file except in compliance with the License.
// You may obtain a copy of the License at
//
//     http://www.apache.org/licenses/LICENSE-2.0
//
// Unless required by applicable law or agreed to in writing, software
// distributed under the License is distributed on an "AS IS" BASIS,
// WITHOUT WARRANTIES OR CONDITIONS OF ANY KIND, either express or implied.
// See the License for the specific language governing permissions and
// limitations under the License.

package client

import (
  "fmt"
  "io"
  "io/ioutil"
  "bytes"
  "net/http"
  "net/url"
  "encoding/json"
  "time"
  "crypto/tls"
)

//Client bla bla
type Client struct {
  host string
  httpClient *http.Client
}

// NewConfigClient bla bla
func NewConfigClient(conf *Config) (*Client, error) {
  cert1, err := tls.LoadX509KeyPair(conf.X509CertFile, conf.X509KeyFile)
  if err != nil {
    panic(err)
  }

  tlsConfig := &tls.Config{
    Certificates: []tls.Certificate{cert1},
  }
  tlsConfig.BuildNameToCertificate()

  transport := &http.Transport{TLSClientConfig: tlsConfig}
  httpClient := &http.Client{ Transport: transport, Timeout: time.Second * 10 }
  return NewClient(conf.Host, httpClient)
}

//NewClient bla bla
func NewClient(host string, httpClient *http.Client) (*Client, error) {
  return &Client {host: host, httpClient: httpClient}, nil
}

func (c Client) get(path string, query url.Values) (serverResponse, error) {
  u := c.getAPIPath(path, query)
  resp, err := c.httpClient.Get(u)
  if err != nil {
    return serverResponse{}, nil
  }

  return c.convertResponse(resp)
}

func (c Client) post(path string, data interface{}) (serverResponse, error) {
  body, err := encodeData(data)
  if err != nil {
    return serverResponse{}, err
  }

  u := c.getAPIPath(path, nil)
  resp, err := c.httpClient.Post(u, "application/json", body)
  if err != nil {
    return serverResponse{}, nil
  }

  return c.convertResponse(resp)
}

func (c Client) put(path string, body []byte, query url.Values) (serverResponse, error) {
  u := c.getAPIPath(path, query)
  request, err := http.NewRequest("PUT", u, bytes.NewBuffer(body))
  if err != nil {
    return serverResponse{}, nil
  }

  resp, err := c.httpClient.Do(request)
  if err != nil {
    return serverResponse{}, nil
  }

  return c.convertResponse(resp)
}

func (c Client) delete(path string) (serverResponse, error) {
  u := c.getAPIPath(path, nil)
  request, err := http.NewRequest("DELETE", u, nil)
  if err != nil {
    return serverResponse{}, err
  }

  resp, err := c.httpClient.Do(request)
  if err != nil {
    return serverResponse{}, err
  }

  return c.convertResponse(resp)
}

func (c Client) convertResponse(resp *http.Response) (serverResponse, error) {
  serverResp := serverResponse{statusCode: -1}

  if resp != nil {
    serverResp.statusCode = resp.StatusCode
  }

  serverResp.body = resp.Body
  serverResp.header = resp.Header
  return serverResp, nil
}

// serverResponse is a wrapper for http API responses.
type serverResponse struct {
	body       io.ReadCloser
	header     http.Header
	statusCode int
}

func (c Client) getAPIPath(path string, query url.Values) string {
	u := &url.URL{
		Path: fmt.Sprintf("%s%s", c.host, path),
	}

	if len(query) > 0 {
		u.RawQuery = query.Encode()
	}
	return u.String()
}

func encodeData(data interface{}) (*bytes.Buffer, error) {
	params := bytes.NewBuffer(nil)
	if data != nil {
		if err := json.NewEncoder(params).Encode(data); err != nil {
			return nil, err
		}
	}
	return params, nil
}

func ensureReaderClosed(response serverResponse) {
	if body := response.body; body != nil {
		// Drain up to 512 bytes and close the body to let the Transport reuse the connection
		io.CopyN(ioutil.Discard, body, 512)
		response.body.Close()
	}
}
