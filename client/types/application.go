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

  cp.Names = string(b)
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
