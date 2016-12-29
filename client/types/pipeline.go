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

package types

import (
  // "fmt"
  "log"
  // "io"
  // "strings"
  "encoding/json"
)

// Pipeline bla bla
type Pipeline struct {
  ID              string
  Name            string
  UpdateTs        JSONTime
  Triggers        []PipelineTrigger

  // Account         string
  // Enabled         bool
  // Organization    string
  // Registry        string
  // Repository      string
  // Tag             string
  // Type            string
  // Description     string
  // User            string
}

// PipelineTrigger bla bla
type PipelineTrigger struct {
  Type        string
  Values      map[string]interface{}
}

//UnmarshalJSON takes a JavaScript timestamp and convert to time
func (ct *PipelineTrigger) UnmarshalJSON(b []byte) (err error) {
  // s := string(b)

  var values interface{}
  mErr := json.Unmarshal(b, &values)

  if mErr != nil {
    log.Fatal(err)
  }

  v := values.(map[string]interface{})

  ct.Type = v["type"].(string)
  ct.Values = v
  return
}

// func triggerType(b []byte) (t *AnonymousPipelineTrigger, err error) {
//   s := string(b)
//   dec := json.NewDecoder(strings.NewReader(s))
//
// 	if err := dec.Decode(&t); err == io.EOF {
//     return t, nil
// 	} else if err != nil {
// 		return t, err
// 	}
//   return
// }

// PipelineBaseTrigger bla bla
type PipelineBaseTrigger struct {
  Type                  string `json:"type"`
  Description           string `json:"description"`
  User                  string `json:"user"`
}

// PipelineDockerTrigger bla bla
type PipelineDockerTrigger struct {
  Type                  string `json:"type"`
  Description           string `json:"description"`
  User                  string `json:"user"`

  Account               string `json:"account"`
  Enabled               bool `json:"enabled"`
  Organization          string `json:"organization"`
  Registry              string `json:"registry"`
  Repository            string `json:"repository"`
  Tag                   string `json:"tag"`
}

// Tag describes docker tags
type Tag struct {
  Account       string
  Registry      string
  Repository    string
  Tag           string
}
