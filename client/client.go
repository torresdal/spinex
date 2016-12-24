package client

import (
  "fmt"
  "os"
  "net/http"
  "bytes"
  "strings"
  "time"
  "crypto/tls"
  "io/ioutil"
  "encoding/json"
  "text/tabwriter"
  "github.com/torresdal/spinex/client/types"
)

//Client bla bla
type Client struct {
  host string
  x509CertFile string
  x509KeyFile string
}

//NewClient bla bla
func NewClient(host, x509CertFile, x509KeyFile string) *Client {
  return &Client {host: host, x509CertFile: x509CertFile, x509KeyFile: x509KeyFile}
}

//getHTTPClient returns a http.Client with credentials ready for use
func getHTTPClient(client *Client) *http.Client {
  cert1, err := tls.LoadX509KeyPair(client.x509CertFile, client.x509KeyFile)
  if err != nil {
    panic(err)
  }

  // Setup HTTPS client
  tlsConfig := &tls.Config{
    Certificates: []tls.Certificate{cert1},
  }
  tlsConfig.BuildNameToCertificate()

  transport := &http.Transport{TLSClientConfig: tlsConfig}
  return &http.Client{Transport: transport}
}

//Applications returns all registered Spinnaker applications
func Applications(client *Client) {
  httpClient := getHTTPClient(client)
  resp, err := httpClient.Get(client.host + "/applications")
  defer resp.Body.Close()

  if err != nil {
    panic(err)
  }

  data, err := ioutil.ReadAll(resp.Body)

  if err != nil {
    panic(err)
  }

  var jsonData[] types.Application
  err = json.Unmarshal([]byte(data), &jsonData) // here!

  if err != nil {
    panic(err)
  }

  const dateFormat = "2006-01-02 15:04:05 MST"

  w := new(tabwriter.Writer)
  // w.Init(output, minwidth, tabwidth, padding, padchar, flags)
  w.Init(os.Stdout, 5, 8, 4, '\t', 0)

  fmt.Println("")
  fmt.Fprintln(w, "Application\t Created\t Updated\t CloudProviders\t Accounts")
  fmt.Fprintln(w, "-----------\t -------\t -------\t --------------\t --------")

  for _, app := range jsonData {
    created := "-"
    if !app.CreateTs.IsZero() {
      created = app.CreateTs.Format(dateFormat)
    }

    updated := "-"
    if !app.UpdateTs.IsZero() {
      updated = app.UpdateTs.Format(dateFormat)
    }

    fmt.Fprintln(w, app.Name, "\t", created, "\t", updated, "\t", app.CloudProviders.Names, "\t", app.Accounts)
  }
  w.Flush()
}

// Application bla bla
func Application(client *Client, name string) types.Application {
  httpClient := getHTTPClient(client)
  resp, err := httpClient.Get(client.host + "/applications/" + name)
  defer resp.Body.Close()

  if err != nil {
    panic(err)
  }

  data, err := ioutil.ReadAll(resp.Body)

  if err != nil {
    panic(err)
  }

  var app types.ApplicationResponse
  err = json.Unmarshal([]byte(data), &app) // here!

  if err != nil {
    panic(err)
  }

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
  if err != nil {
    panic(err)
  }

  httpClient := getHTTPClient(client)
  resp, err := httpClient.Post(client.host + "/applications/" + name + "/tasks", "application/json", bytes.NewBuffer(out))
  defer resp.Body.Close()

  if err != nil {
    panic(err)
  }

  data, err := ioutil.ReadAll(resp.Body)

  if err != nil {
    panic(err)
  }

  var taskRef types.TaskRef
  err = json.Unmarshal([]byte(data), &taskRef) // here!

  if err != nil {
    panic(err)
  }

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
  if err != nil {
    panic(err)
  }

  httpClient := getHTTPClient(client)
  resp, err := httpClient.Post(client.host + "/applications/" + name + "/tasks", "application/json", bytes.NewBuffer(out))
  defer resp.Body.Close()

  if err != nil {
    panic(err)
  }

  data, err := ioutil.ReadAll(resp.Body)

  if err != nil {
    panic(err)
  }

  var taskRef types.TaskRef
  err = json.Unmarshal([]byte(data), &taskRef) // here!

  if err != nil {
    panic(err)
  }

  status := waitForTask(client, taskRef.Ref, 0)
  fmt.Println(status)
}

func waitForTask(client *Client, ref string, counter int) string {
  if counter > 10 {
    return "Timed out waiting for task status"
  }

  task := Task(client, ref)
  // var status string

  var mes string
  if counter > 0 {
    mes += moveCursorUp(len(task.Steps)+4)
  }

  mes += "\nSteps:\n"
  for _, step := range task.Steps {
    // if step.Name == "stageStart" || step.Name == "stageEnd" {
    //   continue
    // }
    mes += fmt.Sprintf("%s\t%s\t%s\n", "\033[K", step.Name, step.Status)
  }
  mes += "\nStatus: In Progress"

  w := new(tabwriter.Writer)
  w.Init(os.Stdout, 5, 8, 4, '\t', 0)

  fmt.Fprintln(w, mes)
  w.Flush()

  if task.Status == "RUNNING" {
    time.Sleep(time.Millisecond * 100)
    return waitForTask(client, ref, counter+1)
  }

  return fmt.Sprintf("%s%s%s%s", moveCursorUp(1), "\033[K", "Status: ", task.Status)
}

func moveCursorUp(lines int) string {
  return fmt.Sprintf("\033[%dA", lines)
}

//Pipelines returns all Pipelines for a Spinnaker application
func Pipelines(client *Client, application string) {
  httpClient := getHTTPClient(client)
  resp, err := httpClient.Get(client.host + "/applications/" + application + "/pipelineConfigs" )
  defer resp.Body.Close()

  if err != nil {
    panic(err)
  }

  data, err := ioutil.ReadAll(resp.Body)

  if err != nil {
    panic(err)
  }

  var jsonData[] types.Pipeline
  err = json.Unmarshal([]byte(data), &jsonData) // here!

  if err != nil {
    panic(err)
  }

  const dateFormat = "2006-01-02 15:04:05 MST"

  w := new(tabwriter.Writer)
  // w.Init(output, minwidth, tabwidth, padding, padchar, flags)
  w.Init(os.Stdout, 5, 8, 4, '\t', 0)

  fmt.Println("")
  fmt.Fprintln(w, "Pipelines\t Id\t Updated")
  fmt.Fprintln(w, "-----------\t -------\t -------")

  for _, pipe := range jsonData {
    updated := "-"
    if !pipe.UpdateTs.IsZero() {
      updated = pipe.UpdateTs.Format(dateFormat)
    }

    fmt.Fprintln(w, pipe.Name, "\t", pipe.ID, "\t", updated)
  }
  w.Flush()
}

// Task will return info and status of a Spinnaker task
func Task(client *Client, ref string) types.TaskResponse {
  httpClient := getHTTPClient(client)
  resp, err := httpClient.Get(client.host + ref)
  defer resp.Body.Close()

  if err != nil {
    panic(err)
  }

  data, err := ioutil.ReadAll(resp.Body)

  if err != nil {
    panic(err)
  }

  var task types.TaskResponse
  err = json.Unmarshal([]byte(data), &task) // here!

  if err != nil {
    panic(err)
  }

  return task
}
