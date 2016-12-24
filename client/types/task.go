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
