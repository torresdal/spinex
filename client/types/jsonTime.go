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
  "fmt"
  "time"
  "strings"
  "strconv"
)

//JSONTime is a type of for handling convertions of UNIX time
type JSONTime struct {
    time.Time
}

//MarshalJSON will convert time into JavaScript timestamp
func (ct *JSONTime) MarshalJSON() ([]byte, error) {
  if ct.Time.UnixNano() == nilTime {
    return []byte("null"), nil
  }
  return []byte(fmt.Sprintf("\"%s\"", ct.Time.Format(ctLayout))), nil
}

var nilTime = (time.Time{}).UnixNano()

const ctLayout = "2006/01/02|15:04:05"

//UnmarshalJSON takes a JavaScript timestamp and convert to time
func (ct *JSONTime) UnmarshalJSON(b []byte) (err error) {
  s := strings.Trim(string(b), "\"")
  if s == "null" {
    ct.Time = time.Date(1, time.January, 1, 0, 0, 0, 0, time.UTC)
    return
  }

  i, err := strconv.ParseInt(s, 10, 64)
  if err != nil {
      panic(err)
  }

  i = i/1000

  ct.Time = time.Unix(i,0)
  return
}
