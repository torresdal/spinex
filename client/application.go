package client

import (
  "fmt"
  "bytes"
  "strings"
  "io/ioutil"
  "encoding/json"
  "github.com/torresdal/spinex/client/types"
)

//Applications returns all registered Spinnaker applications
func Applications(client *Client) {
  httpClient := getHTTPClient(client)
  resp, err := httpClient.Get(client.host + "/applications")
  defer resp.Body.Close()
  checkErr(err)

  data, err := ioutil.ReadAll(resp.Body)
  checkErr(err)

  var jsonData[] types.Application
  err = json.Unmarshal([]byte(data), &jsonData) // here!
  checkErr(err)

  FormatApplicationList(jsonData)
}

// Application bla bla
func Application(client *Client, name string) types.Application {
  httpClient := getHTTPClient(client)
  resp, err := httpClient.Get(client.host + "/applications/" + name)
  defer resp.Body.Close()
  checkErr(err)

  data, err := ioutil.ReadAll(resp.Body)
  checkErr(err)

  var app types.ApplicationResponse
  err = json.Unmarshal([]byte(data), &app) // here!
  checkErr(err)

  return app.Attributes
}

// CreateApplication bla bla
func CreateApplication(client *Client, name string, email string, accounts string, cloudProviders string, instancePort string, description string) {
  var jobs []types.Job

  a := strings.Split(accounts, ",")

  for _, account := range a {
    jobs = append(jobs, types.Job {
      Type: "createApplication",
      Account: account,
      User: "",
      Application: types.CreateApplication {
        Name: name,
        Description: description,
        Accounts: accounts,
        CloudProviders: cloudProviders,
        Email: email,
        InstancePort: instancePort,
      },
    })
  }

  task := types.Task {
    Application: name,
    Description: "Create Application: " + name,
    Job : jobs,
  }

  out, err := json.Marshal(task)
  checkErr(err)

  httpClient := getHTTPClient(client)
  resp, err := httpClient.Post(client.host + "/applications/" + name + "/tasks", "application/json", bytes.NewBuffer(out))
  defer resp.Body.Close()
  checkErr(err)

  data, err := ioutil.ReadAll(resp.Body)
  checkErr(err)

  var taskRef types.TaskRef
  err = json.Unmarshal([]byte(data), &taskRef) // here!
  checkErr(err)

  status := waitForTask(client, taskRef.Ref, 0)
  fmt.Println(status)
}

// DeleteApplication will delete a Spinnaker application
func DeleteApplication(client *Client, name string) {
  app := Application(client, name)
  accounts := strings.Split(app.Accounts, ",")

  var jobs []types.Job

  for _, account := range accounts {
    jobs = append(jobs, types.Job {
      Type: "deleteApplication",
      Account: account,
      User: "",
      Application: types.DeleteApplication {
        Name: app.Name,
        Accounts: app.Accounts,
        CloudProviders: app.CloudProviders.Names,
      },
    })
  }

  task := types.Task {
    Application: app.Name,
    Description: "Deleting Application: " + app.Name,
    Job : jobs,
  }

  out, err := json.Marshal(task)
  checkErr(err)

  httpClient := getHTTPClient(client)
  resp, err := httpClient.Post(client.host + "/applications/" + name + "/tasks", "application/json", bytes.NewBuffer(out))
  defer resp.Body.Close()
  checkErr(err)

  data, err := ioutil.ReadAll(resp.Body)
  checkErr(err)

  var taskRef types.TaskRef
  err = json.Unmarshal([]byte(data), &taskRef) // here!
  checkErr(err)

  status := waitForTask(client, taskRef.Ref, 0)
  fmt.Println(status)
}
