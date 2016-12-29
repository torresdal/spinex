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
