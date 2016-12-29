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

type fn func(e Execution, name string) bool

// Filter will filter
func (s ExecutionList) Filter(f fn, name string) ExecutionList {
  b := s[:0]
  for _, x := range s {
      if f(x, name) {
          b = append(b, x)
      }
  }
  return b
}

// ExecutionList is a list of Execution objects
type ExecutionList []Execution

// First finds the first occourance of Execution with the given criteria
func (s ExecutionList) First(name string) (*Execution, error) {
  lName := strings.ToLower(name)
  for _, e := range s {
    if(strings.ToLower(e.Name) == lName) {
      return &e, nil
    }
  }
  return nil, errors.New("Could not find Executon with name " + name)
}

// Len bla bla
func (s ExecutionList) Len() int      { return len(s) }
// Swap bla bla
func (s ExecutionList) Swap(i, j int) { s[i], s[j] = s[j], s[i] }

// ByNameAsc sor by Name
type ByNameAsc struct{ ExecutionList }

// ByNameDesc sor by Name
type ByNameDesc struct{ ExecutionList }

// ByStartTimeAsc sor by StartTime
type ByStartTimeAsc struct{ ExecutionList }

// ByStartTimeDesc sor by StartTime
type ByStartTimeDesc struct{ ExecutionList }

// ByEndTimeDesc sor by StartTime
type ByEndTimeDesc struct{ ExecutionList }

// ByEndTimeAsc sor by StartTime
type ByEndTimeAsc struct{ ExecutionList }

// ByStatus sor by Status
type ByStatus struct{ ExecutionList }

// Less bla bla
func (s ByNameAsc) Less(i, j int) bool { return s.ExecutionList[i].Name < s.ExecutionList[j].Name }
// Less bla bla
func (s ByNameDesc) Less(i, j int) bool { return s.ExecutionList[i].Name > s.ExecutionList[j].Name }
// Less bla bla
func (s ByStatus) Less(i, j int) bool { return s.ExecutionList[i].Status < s.ExecutionList[j].Status }
// Less bla bla
func (s ByStartTimeAsc) Less(i, j int) bool { return s.ExecutionList[i].StartTime.Before(s.ExecutionList[j].StartTime.Time) }
// Less bla bla
func (s ByStartTimeDesc) Less(i, j int) bool { return s.ExecutionList[i].StartTime.After(s.ExecutionList[j].StartTime.Time) }
// Less bla bla
func (s ByEndTimeAsc) Less(i, j int) bool { return s.ExecutionList[i].EndTime.Before(s.ExecutionList[j].EndTime.Time) }
// Less bla bla
func (s ByEndTimeDesc) Less(i, j int) bool { return s.ExecutionList[i].EndTime.After(s.ExecutionList[j].EndTime.Time) }

// // ByStartTimeDesc sort by StartTime descending
// type ByStartTimeDesc ExecutionList
//
// func (a ByStartTimeDesc) Len() int           { return len(a) }
// func (a ByStartTimeDesc) Less(i, j int) bool { return a[i].StartTime.After(a[j].StartTime.Time) }
// func (a ByStartTimeDesc) Swap(i, j int)      { a[i], a[j] = a[j], a[i] }
//
// // ByStartTimeAsc sort by StartTime ascending
// type ByStartTimeAsc []Execution
//
// func (a ByStartTimeAsc) Len() int           { return len(a) }
// func (a ByStartTimeAsc) Less(i, j int) bool { return a[i].StartTime.Before(a[j].StartTime.Time) }
// func (a ByStartTimeAsc) Swap(i, j int)      { a[i], a[j] = a[j], a[i] }
//
// // ByEndTimeDesc sort by EndTime
// type ByEndTimeDesc []Execution
//
// func (a ByEndTimeDesc) Len() int           { return len(a) }
// func (a ByEndTimeDesc) Less(i, j int) bool { return a[i].EndTime.After(a[j].EndTime.Time) }
// func (a ByEndTimeDesc) Swap(i, j int)      { a[i], a[j] = a[j], a[i] }
//
// // ByEndTimeAsc sort by EndTime ascending
// type ByEndTimeAsc []Execution
//
// func (a ByEndTimeAsc) Len() int           { return len(a) }
// func (a ByEndTimeAsc) Less(i, j int) bool { return a[i].EndTime.Before(a[j].EndTime.Time) }
// func (a ByEndTimeAsc) Swap(i, j int)      { a[i], a[j] = a[j], a[i] }
//
// // ByName sort by Name
// type ByName []Execution
//
// func (a ByName) Len() int           { return len(a) }
// func (a ByName) Less(i, j int) bool { return a[i].Name < a[j].Name }
// func (a ByName) Swap(i, j int)      { a[i], a[j] = a[j], a[i] }
//
// // ByStatus sort by Status
// type ByStatus []Execution
//
// func (a ByStatus) Len() int           { return len(a) }
// func (a ByStatus) Less(i, j int) bool { return a[i].Status < a[j].Status }
// func (a ByStatus) Swap(i, j int)      { a[i], a[j] = a[j], a[i] }
