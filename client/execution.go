package client

import (
  "fmt"
  "strings"
  "io/ioutil"
  "net/url"
  "strconv"
  "sort"
  "github.com/torresdal/spinex/client/types"
  "encoding/json"
  "log"
  "net/http"
)
//Executions returns all executions for an application
func Executions(client *Client, application string, limit int, statuses string, sortBy string, desc bool, name string) {
  var qLimit string
  var qStatuses string
  var q string

  fmt.Println("Name: ", name)
  if limit > 0 {
    qLimit += "limit=" + strconv.Itoa(limit)
  }

  if statuses != "" {
    qStatuses += "&statuses=" + strings.ToUpper(statuses)
  }

  if qLimit != "" && qStatuses != "" {
    q = "?" + qLimit + "&" + qStatuses
  } else if qLimit != "" {
    q = "?" + qLimit
  } else if qStatuses != "" {
    q = "?" + qStatuses
  }

  url := client.host + "/applications/" + application + "/pipelines" + q

  httpClient := getHTTPClient(client)
  resp, err := httpClient.Get(url)
  defer resp.Body.Close()
  checkErr(err)

  data, err := ioutil.ReadAll(resp.Body)
  checkErr(err)

  var jsonData types.ExecutionList
  err = json.Unmarshal([]byte(data), &jsonData)
  checkErr(err)

  if name != "" {
    first, err := types.ExecutionList.First(jsonData, name)
    if err != nil {
      log.Fatal(err)
    }
    jsonData = types.ExecutionList.Filter(jsonData, f, first.PipelineConfigID)
  }

  var sortMsg string
  switch strings.ToLower(sortBy) {
    case "name":
      if desc {
        sort.Sort(types.ByNameDesc{ExecutionList: jsonData})
      } else {
        sort.Sort(types.ByNameAsc{ExecutionList: jsonData})
      }
      sortMsg = "Sorted by NAME"
    case "end":
      sort.Sort(types.ByEndTimeDesc{ExecutionList: jsonData})
      sortMsg = "Sorted by END desc"
    case "status":
      sort.Sort(types.ByStatus{ExecutionList: jsonData})
      sortMsg = "Sorted by STATUS"
    default:
      sort.Sort(types.ByStartTimeDesc{ExecutionList: jsonData})
      sortMsg = "Sorted by START desc"
  }

  FormatExecutionList(jsonData)
  fmt.Println()
  fmt.Print(sortMsg)
  fmt.Println()
}

// CancelExecution will cancel a running pipeline execution
func CancelExecution(client *Client, id string, reason string) {
  httpClient := getHTTPClient(client)
  rURL := client.host + "/pipelines/" + id + "/cancel"

  if reason != "" {
    fReason := url.QueryEscape(reason)
    rURL += "?reason=" + fReason
  }
  request, err := http.NewRequest("PUT", rURL, nil)
  checkErr(err)

  resp, err := httpClient.Do(request)
  defer resp.Body.Close()
  checkErr(err)

  data, err := ioutil.ReadAll(resp.Body)
  checkErr(err)

  fmt.Println(data)
}

// DeleteExecution will cancel a running pipeline execution
func DeleteExecution(client *Client, id string) {
  httpClient := getHTTPClient(client)
  rURL := client.host + "/pipelines/" + id

  request, err := http.NewRequest("DELETE", rURL, nil)
  checkErr(err)

  resp, err := httpClient.Do(request)
  defer resp.Body.Close()
  checkErr(err)

  data, err := ioutil.ReadAll(resp.Body)
  checkErr(err)

  fmt.Println(data)
}

// ExecutionInfo gets detailed information about a pipeline execution
func ExecutionInfo(client *Client, id string) {
  rURL := client.host + "/pipelines/" + id

  httpClient := getHTTPClient(client)
  resp, err := httpClient.Get(rURL)
  defer resp.Body.Close()
  checkErr(err)

  data, err := ioutil.ReadAll(resp.Body)
  checkErr(err)

  var jsonData types.Execution
  err = json.Unmarshal([]byte(data), &jsonData)
  checkErr(err)

  FormatExecutionInfo(jsonData)
}