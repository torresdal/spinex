package client

import (
  "io/ioutil"
  "encoding/json"
  "github.com/torresdal/spinex/client/types"
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

func f(e types.Execution, confID string) bool {
  return e.PipelineConfigID == confID
}
