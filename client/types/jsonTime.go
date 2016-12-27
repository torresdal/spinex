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
