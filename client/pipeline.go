package client

import (
  "io/ioutil"
  "bytes"
  "fmt"
  "strconv"
  "log"
  "encoding/json"
  "github.com/torresdal/spinex/client/types"
  "net/url"
)

//Pipelines returns all Pipelines for a Spinnaker application
func Pipelines(client *Client, application string) {
  httpClient := getHTTPClient(client)
  resp, err := httpClient.Get(client.host + "/applications/" + application + "/pipelineConfigs" )
  defer resp.Body.Close()
  checkErr(err)

  data, err := ioutil.ReadAll(resp.Body)
  checkErr(err)

  var jsonData[] types.Pipeline
  err = json.Unmarshal([]byte(data), &jsonData) // here!
  checkErr(err)

  FormatPipelineList(jsonData)
}

//Pipeline returns a pipeline for a Spinnaker application
func Pipeline(client *Client, application string, pipeline string) types.Pipeline {
  httpClient := getHTTPClient(client)
  resp, err := httpClient.Get(client.host + "/applications/" + application + "/pipelineConfigs/" +pipeline)
  defer resp.Body.Close()
  checkErr(err)

  data, err := ioutil.ReadAll(resp.Body)
  checkErr(err)

  var jsonData types.Pipeline
  err = json.Unmarshal([]byte(data), &jsonData) // here!
  checkErr(err)

  return jsonData
}

// StartPipeline will start a new pipeline execution
func StartPipeline(client *Client, app string, pipeline string, dockerTag string) {
  pipe := Pipeline(client, app, pipeline)

  numOfTriggers := len(pipe.Triggers)
  var body interface{}

  if numOfTriggers > 0 {
    if numOfTriggers != 1 {
      log.Fatal("Spinex currently only support pipelines with max one trigger")
    }

    switch pipe.Triggers[0].Type {
      case "docker":
        trigger := getDockerTrigger(pipe.Triggers[0].Values)
        if dockerTag != "" {
          trigger.Tag = dockerTag
        } else if trigger.Tag == "" {
          tags := findTags(client, trigger.Account, trigger.Repository)
          tag := promptForTag(tags)
          trigger.Tag = tag.Tag
        }
        body = trigger
      default:
        body = types.PipelineBaseTrigger {
          Type          : "manual",
          Description   : "Started by Spinex",
          User          : "spinex",
        }
    }
  }

  bodyJSON, err := json.Marshal(body)
  checkErr(err)

  fmt.Println(string(bodyJSON))

  httpClient := getHTTPClient(client)
  resp, err := httpClient.Post(client.host + "/pipelines/" + app + "/" + pipeline, "application/json", bytes.NewBuffer(bodyJSON))
  defer resp.Body.Close()
  checkErr(err)

  data, err := ioutil.ReadAll(resp.Body)
  checkErr(err)

  var pipeRef types.TaskRef
  err = json.Unmarshal([]byte(data), &pipeRef) // here!
  checkErr(err)

  fmt.Println(string(data))
}

func findTags(client *Client, account string, repo string) []types.Tag {
//https://deploy.milescloud.io:8084/images/find?
  httpClient := getHTTPClient(client)
  qStr := fmt.Sprintf("?account=%s&count=20&provider=dockerRegistry&q=%s", url.QueryEscape(account), url.QueryEscape(repo + ":"))
  resp, err := httpClient.Get(client.host + "/images/find" + qStr)
  defer resp.Body.Close()
  checkErr(err)

  data, err := ioutil.ReadAll(resp.Body)
  checkErr(err)

  var tags []types.Tag
  err = json.Unmarshal([]byte(data), &tags) // here!
  checkErr(err)

  return tags
}

func promptForTag(tags []types.Tag) types.Tag {
  var t string
  fmt.Println()
  fmt.Println("Available tags:")

  for i, tag := range tags {
    fmt.Printf("  %d) %s\n", i+1, tag.Tag)
  }
  fmt.Println()

  i := -1

  fmt.Print("Tag number: ")
  fmt.Scanln(&t)

	n, err := strconv.Atoi(t)
  checkErr(err)

	if n > 0 && n <= len(tags) {
		i = n - 1
	} else {
    log.Fatal("Number for tag does not exist")
  }

  return tags[i]
}

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
