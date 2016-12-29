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

import(
  "encoding/json"
  "github.com/torresdal/spinex/client/types"
)

// Task will return info and status of a Spinnaker task
func (c Client) Task(ref string) (types.TaskResponse, error) {
  var taskRes types.TaskResponse

  resp, err := c.get(ref, nil)
  defer ensureReaderClosed(resp)

  if err != nil {
    return taskRes, err
  }

  err = json.NewDecoder(resp.body).Decode(&taskRes)
  return taskRes, err
}
