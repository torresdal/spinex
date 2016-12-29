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
  "encoding/json"
  "github.com/torresdal/spinex/client/types"
)

//Applications returns all registered Spinnaker applications
func (c Client) Applications() ([]types.Application, error) {
  resp, err := c.get("/applications", nil)
  defer ensureReaderClosed(resp)
  if err != nil {
    return nil, err
  }

  var apps []types.Application
  err = json.NewDecoder(resp.body).Decode(&apps)
  return apps, err
}

// Application returns application with a given name
func (c Client) Application(name string) (types.Application, error) {
  var app types.ApplicationResponse

  resp, err := c.get("/applications/" + name, nil)
  defer ensureReaderClosed(resp)
  if err != nil {
    return app.Attributes, err
  }

  err = json.NewDecoder(resp.body).Decode(&app)

  if err != nil {
    return app.Attributes, err
  }
  return app.Attributes, err
}

// CreateApplication will create a new application in Spinnaker
func (c Client) CreateApplication(name string, email string, accounts string, cloudProviders string, instancePort string, description string) (types.TaskRef, error) {
  var taskRef types.TaskRef
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

  resp, err := c.post("/applications/" + name + "/tasks", task)
  defer ensureReaderClosed(resp)
  if err != nil {
    return taskRef, err
  }

  err = json.NewDecoder(resp.body).Decode(&taskRef)
  return taskRef, err
}

// DeleteApplication will delete a Spinnaker application with a given name
func (c Client) DeleteApplication(name string) (types.TaskRef, error) {
  var taskRef types.TaskRef

  app, err := c.Application(name)
  if err != nil {
    return taskRef, err
  }

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

  resp, err := c.post("/applications/" + name + "/tasks", task)
  defer ensureReaderClosed(resp)
  if err != nil {
    return taskRef, err
  }

  err = json.NewDecoder(resp.body).Decode(&taskRef)
  return taskRef, err
}
