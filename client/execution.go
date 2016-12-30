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
  "strings"
  "net/url"
  "strconv"
  "github.com/torresdal/spinex/client/types"
  "encoding/json"
)

//Executions returns all executions for an application
func (c *Client) Executions(application string, limit int, statuses string) (types.ExecutionSlice, error) {
  var executions types.ExecutionSlice
  var query = url.Values{}

  if limit > 0 {
    query.Add("limit", strconv.Itoa(limit))
  }

  if statuses != "" {
    query.Add("statuses", strings.ToUpper(statuses))
  }

  resp, err := c.get("/applications/" + application + "/pipelines", query)
  defer ensureReaderClosed(resp)
  if err != nil {
    return executions, err
  }

  err = json.NewDecoder(resp.body).Decode(&executions)
  return executions, err
}

// CancelExecution will cancel a running pipeline execution
func (c *Client) CancelExecution(id string, reason string) error {
  var queries = url.Values{}
  if reason != "" {
    queries.Add("reason", reason)
  }

  resp, err := c.put("/pipelines/" + id + "/cancel", nil, queries)
  ensureReaderClosed(resp)
  return err
}

// DeleteExecution will cancel a running pipeline execution
func (c *Client) DeleteExecution(id string) error {
  resp, err := c.delete("/pipelines/" + id)
  ensureReaderClosed(resp)
  return err
}

// ExecutionInfo gets detailed information about a pipeline execution
func (c *Client) ExecutionInfo(id string) (types.Execution, error) {
  var ex types.Execution
  resp, err := c.get("/pipelines/" + id, nil)
  defer ensureReaderClosed(resp)

  if err != nil {
    return ex, err
  }

  err = json.NewDecoder(resp.body).Decode(&ex)
  return ex, err
}
