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

// Task bla bla
type Task struct {
//{"job":[{"type":"deleteApplication","account":"miles-kube-1","application":{"name":"temp2","accounts":"miles-kube-1","cloudProviders":"kubernetes,aws"},"user":"jon.arild.torresdal@miles.no"}],"application":"temp2","description":"Deleting Application: temp2"}
  Job           []Job `json:"job"`
  Application   string `json:"application"`
  Description   string `json:"description"`
}

// Job bla bla
type Job struct {
  Type          string `json:"type"`
  Account       string `json:"account"`
  Application   interface{} `json:"application"`
  User          string `json:"user"`
}

// TaskResponse contains information about a Task
type TaskResponse struct {
  Application     string
  BuildTime       JSONTime
  StartTime       JSONTime
  EndTime         JSONTime
  ID              string
  Name            string
  Status          string
  Steps           []TaskStep
}

// TaskStep contains the details of each step in a Task
type TaskStep struct {
  ID              string
  Name            string
  StartTime       JSONTime
  EndTime         JSONTime
  Status          string
}

// TaskRef has the reference URL for getting detailed Task info
type TaskRef struct {
  Ref   string
}
