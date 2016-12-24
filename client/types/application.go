package types

import (
  "strings"
)

// Application bla bla
type Application struct {
    Name            string
    Accounts        string
    CloudProviders  CloudProviders
    User            string
    Email           string
    CreateTs        JSONTime
    UpdateTs        JSONTime
}

// CreateApplication bla bla
type CreateApplication struct {
    Name            string `json:"name"`
    Description     string `json:"description"`
    Accounts        string `json:"accounts"`
    CloudProviders  string `json:"cloudProviders"`
    Email           string `json:"email"`
    InstancePort    string `json:"instancePort"`
}

// CloudProviders bla bla
type CloudProviders struct {
  Names string
}

//MarshalJSON will convert time into JavaScript timestamp
func (cp *CloudProviders) MarshalJSON() ([]byte, error) {
  return []byte(cp.Names), nil
}

//UnmarshalJSON takes a JavaScript timestamp and convert to time
func (cp *CloudProviders) UnmarshalJSON(b []byte) (err error) {
  if string(b) == "null" {
		return nil
	}

  cp.Names = strings.Trim(string(b), "\"")
  return
}

// ApplicationResponse bla bla
type ApplicationResponse struct {
  Attributes        Application
}

// DeleteApplication bla bla
type DeleteApplication struct {
  Name            string `json:"name"`
  Accounts        string `json:"accounts"`
  CloudProviders  string `json:"cloudProviders"`
}
