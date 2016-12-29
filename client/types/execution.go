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
  "errors"
  "strings"
)
// Execution represents a pipeline execution
type Execution struct {
  Application           string
  BuildTime             JSONTime
  Canceled              bool
  CanceledBy            string
  CancellationReason    string
  EndTime               JSONTime
  ID                    string
  KeepWaitingPipelines  bool
  LimitConcurrent       bool
  Name                  string
  Parallel              bool
  Paused                bool
  PipelineConfigID      string
  StartTime             JSONTime
  Status                string
  Stages                []ExecutionStage
}

// ExecutionStage has details of a stage during pipeline execution
type ExecutionStage struct {
  EndTime               JSONTime
  ID                    string
  Immutable             bool
  InitializationStage   bool
  // LastModified          JSONTime
  Name                  string
  ParentStageID         string
  RefID                 string
  RequisiteStageRefIds  []string
  scheduledTime         JSONTime
  StartTime             JSONTime
  Status                string
  SyntheticStageOwner   string
  Type                  string
}

// ExecutionRequest describes a request for a new pipeline execution
type ExecutionRequest struct {
  Type          string // manual, docker...
  User          string //jon@torresdal.net
}

// DockerTriggerExecutionRequest describes a request for a new pipeline execution using a Docker trigger
type DockerTriggerExecutionRequest struct {
  ExecutionRequest

  Account               string
  Enabled               bool
  Organization          string
  Registry              string
  Repository            string
  Tag                   string
}

// Filter will filter
func (s ExecutionSlice) Filter(name string) (ExecutionSlice, error) {
  first, err := s.First(name);
  if err != nil {
    return s, err
  }

  b := s[:0]
  for _, x := range s {
      if first.PipelineConfigID == x.PipelineConfigID {
          b = append(b, x)
      }
  }
  return b, err
}

// ExecutionSlice is a list of Execution objects
type ExecutionSlice []Execution

// First finds the first occourance of Execution with the given criteria
func (s ExecutionSlice) First(name string) (*Execution, error) {
  lName := strings.ToLower(name)
  for _, e := range s {
    if(strings.ToLower(e.Name) == lName) {
      return &e, nil
    }
  }
  return nil, errors.New("Could not find Executon with name " + name)
}

// Len bla bla
func (s ExecutionSlice) Len() int      { return len(s) }
// Swap bla bla
func (s ExecutionSlice) Swap(i, j int) { s[i], s[j] = s[j], s[i] }

// ByNameAsc sor by Name
type ByNameAsc struct{ ExecutionSlice }

// ByNameDesc sor by Name
type ByNameDesc struct{ ExecutionSlice }

// ByStartTimeAsc sor by StartTime
type ByStartTimeAsc struct{ ExecutionSlice }

// ByStartTimeDesc sor by StartTime
type ByStartTimeDesc struct{ ExecutionSlice }

// ByEndTimeDesc sor by StartTime
type ByEndTimeDesc struct{ ExecutionSlice }

// ByEndTimeAsc sor by StartTime
type ByEndTimeAsc struct{ ExecutionSlice }

// ByStatus sor by Status
type ByStatus struct{ ExecutionSlice }

// Less bla bla
func (s ByNameAsc) Less(i, j int) bool { return s.ExecutionSlice[i].Name < s.ExecutionSlice[j].Name }
// Less bla bla
func (s ByNameDesc) Less(i, j int) bool { return s.ExecutionSlice[i].Name > s.ExecutionSlice[j].Name }
// Less bla bla
func (s ByStatus) Less(i, j int) bool { return s.ExecutionSlice[i].Status < s.ExecutionSlice[j].Status }
// Less bla bla
func (s ByStartTimeAsc) Less(i, j int) bool { return s.ExecutionSlice[i].StartTime.Before(s.ExecutionSlice[j].StartTime.Time) }
// Less bla bla
func (s ByStartTimeDesc) Less(i, j int) bool { return s.ExecutionSlice[i].StartTime.After(s.ExecutionSlice[j].StartTime.Time) }
// Less bla bla
func (s ByEndTimeAsc) Less(i, j int) bool { return s.ExecutionSlice[i].EndTime.Before(s.ExecutionSlice[j].EndTime.Time) }
// Less bla bla
func (s ByEndTimeDesc) Less(i, j int) bool { return s.ExecutionSlice[i].EndTime.After(s.ExecutionSlice[j].EndTime.Time) }
