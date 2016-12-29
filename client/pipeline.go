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
  "encoding/json"
  "github.com/torresdal/spinex/client/types"
  "net/url"
)

//Pipelines returns all Pipelines for a Spinnaker application
func (c *Client) Pipelines(application string) ([]types.Pipeline, error) {
  var pipes []types.Pipeline
  resp, err := c.get("/applications/" + application + "/pipelineConfigs", nil)
  defer ensureReaderClosed(resp)

  if err != nil {
    return pipes, err
  }

  err = json.NewDecoder(resp.body).Decode(&pipes)
  return pipes, err
}

//Pipeline returns a pipeline for a Spinnaker application
func (c *Client) Pipeline(application string, pipeline string) (types.Pipeline, error) {
  var pipe types.Pipeline

  resp, err := c.get("/applications/" + application + "/pipelineConfigs/" +pipeline, nil)
  defer ensureReaderClosed(resp)

  if err != nil {
    return pipe, err
  }

  err = json.NewDecoder(resp.body).Decode(&pipe)
  return pipe, err
}

// StartPipeline will start a new pipeline execution
func (c *Client) StartPipeline(app string, pipeline string, trigger interface{}) (types.TaskRef, error) {
  var taskRef types.TaskRef

  // pipe, err := c.Pipeline(app, pipeline)

  // if err != nil {
  //   return taskRef, err
  // }

  // numOfTriggers := len(pipe.Triggers)
  // var body interface{}
  //
  // if numOfTriggers > 0 {
  //   if numOfTriggers != 1 {
  //     log.Fatal("Spinex currently only support pipelines with max one trigger")
  //   }
  //
  //   switch pipe.Triggers[0].Type {
  //     case "docker":
  //       trigger := getDockerTrigger(pipe.Triggers[0].Values)
  //       if dockerTag != "" {
  //         trigger.Tag = dockerTag
  //       } else if trigger.Tag == "" {
  //         tags := findTags(c, trigger.Account, trigger.Repository)
  //         tag := promptForTag(tags)
  //         trigger.Tag = tag.Tag
  //       }
  //       body = trigger
  //     default:
  //       body = types.PipelineBaseTrigger {
  //         Type          : "manual",
  //         Description   : "Started by Spinex",
  //         User          : "spinex",
  //       }
  //   }
  // }

  resp, err := c.post("/pipelines/" + app + "/" + pipeline, trigger)
  defer ensureReaderClosed(resp)
  if err != nil {
    return taskRef, err
  }

  err = json.NewDecoder(resp.body).Decode(&taskRef)
  return taskRef, err
}

func findTags(c *Client, account string, repo string) ([]types.Tag, error) {
  var tags []types.Tag
  var query url.Values

  query.Add("account", account)
  query.Add("count", "20")
  query.Add("provider", "dockerRegistry")
  query.Add("q", repo + ":")

  resp, err := c.get("/images/find", query)
  defer ensureReaderClosed(resp)
  if err != nil {
    return tags, err
  }

  err = json.NewDecoder(resp.body).Decode(&tags)
  return tags, err
}

// func promptForTag(tags []types.Tag) types.Tag {
//   var t string
//   fmt.Println()
//   fmt.Println("Available tags:")
//
//   for i, tag := range tags {
//     fmt.Printf("  %d) %s\n", i+1, tag.Tag)
//   }
//   fmt.Println()
//
//   i := -1
//
//   fmt.Print("Tag number: ")
//   fmt.Scanln(&t)
//
// 	n, err := strconv.Atoi(t)
//   checkErr(err)
//
// 	if n > 0 && n <= len(tags) {
// 		i = n - 1
// 	} else {
//     log.Fatal("Number for tag does not exist")
//   }
//
//   return tags[i]
// }

func getDockerTrigger(values map[string]interface{}) types.PipelineDockerTrigger {
  var tag string
  if values["tag"] != nil {
    tag = values["tag"].(string)
  }

  trigger := types.PipelineDockerTrigger {
    Type          : "manual",
    Description   : "Started by Spinex",
    User          : "spinex",

    Account       : values["account"].(string),
    Enabled       : values["enabled"].(bool),
    Organization  : values["organization"].(string),
    Registry      : values["registry"].(string),
    Repository    : values["repository"].(string),
    Tag           : tag,
  }

  return trigger
}
