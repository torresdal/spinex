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

package cmd

import (
  "os"
  "fmt"
  "log"
  "time"
  "text/tabwriter"
  "github.com/torresdal/spinex/client"
  "github.com/torresdal/spinex/client/types"
  tw "github.com/olekukonko/tablewriter"
)

func getTable(headers []string) *tw.Table {
  table := tw.NewWriter(os.Stdout)
  if(headers != nil) {
    table.SetHeader(headers)
  }
  table.SetHeaderAlignment(tw.ALIGN_LEFT)
  table.SetAlignment(tw.ALIGN_LEFT)
  table.SetBorder(false)
  table.SetCenterSeparator(" ")
  table.SetColumnSeparator(" ")

  return table
}

// FormatApplicationList will format a list of applications
func FormatApplicationList(apps []types.Application) {
  const dateFormat = "2006-01-02 15:04:05 MST"

  table := getTable([]string {"Application", "Created", "Updated", "Owner", "Accounts"})

  for _, app := range apps {
    created := "-"
    if !app.CreateTs.IsZero() {
      created = app.CreateTs.Format(dateFormat)
    }

    updated := "-"
    if !app.UpdateTs.IsZero() {
      updated = app.UpdateTs.Format(dateFormat)
    }

    table.Append([]string {app.Name, created, updated, app.Email, app.Accounts})
  }
  table.Render()
}

// FormatPipelineList ..
func FormatPipelineList(pipelines []types.Pipeline) {
  const dateFormat = "2006-01-02 15:04:05 MST"

  table := getTable([]string {"Pipelines", "Id", "Updated"})

  for _, pipe := range pipelines {
    updated := "-"
    if !pipe.UpdateTs.IsZero() {
      updated = pipe.UpdateTs.Format(dateFormat)
    }
    table.Append([]string {pipe.Name, pipe.ID, updated})
  }
  table.Render()
}

// FormatExecutionList formats output of a list of pipeline executions
func FormatExecutionList(executions []types.Execution) {
  const dateFormat = "2006-01-02 15:04:05 MST"

  table := getTable([]string {"Name", "Id", "Start", "End", "Duration", "Status"})

  for _, exec := range executions {
    started := "-"
    if !exec.StartTime.IsZero() {
      started = exec.StartTime.Format(dateFormat)
    }

    ended := "-"
    if !exec.EndTime.IsZero() {
      ended = exec.EndTime.Format(dateFormat)
    }

    duration := "-"
    if !exec.EndTime.IsZero() && !exec.StartTime.IsZero() {
      duration = exec.EndTime.Sub(exec.StartTime.Time).String()
    }
    table.Append([]string {exec.Name, exec.ID, started, ended, duration, exec.Status})
  }
  table.Render()
}

// FormatExecutionInfo formats information about pipeline execution
func FormatExecutionInfo(exec types.Execution) {
  const dateFormat = "2006-01-02 15:04:05 MST"

  table := getTable(nil)

  started := "-"
  if !exec.StartTime.IsZero() {
    started = exec.StartTime.Format(dateFormat)
  }

  ended := "-"
  if !exec.EndTime.IsZero() {
    ended = exec.EndTime.Format(dateFormat)
  }

  duration := "-"
  if !exec.EndTime.IsZero() && !exec.StartTime.IsZero() {
    duration = exec.EndTime.Sub(exec.StartTime.Time).String()
  }
  table.Append([]string {"Name", exec.Name})
  table.Append([]string {"ID", exec.ID})
  table.Append([]string {"Application", exec.Application})
  table.Append([]string {"Status", exec.Status})
  table.Append([]string {"Start", started})
  table.Append([]string {"End", ended})
  table.Append([]string {"Duration", duration})
  table.Render()

  fmt.Println("")
  table = getTable(nil)
  for _, stage := range exec.Stages {
    if stage.SyntheticStageOwner == "" {
      duration := stage.EndTime.Sub(stage.StartTime.Time).String()
      table.Append([]string {stage.Name})
      table.Append([]string {"", "Id", stage.ID})
      table.Append([]string {"", "Status", stage.Status})
      table.Append([]string {"", "Start", stage.StartTime.String()})
      table.Append([]string {"", "End", stage.EndTime.String()})
      table.Append([]string {"", "Duration", duration})
      table.Append([]string {"", "Type", stage.Type})
    }
  }
  table.Render()
}

func waitForTask(cli *client.Client, ref string, counter int) string {
  if counter > 10 {
    return "Timed out waiting for task status"
  }

  task, err := cli.Task(ref)

  if err != nil {
    log.Fatal(err)
  }

  var mes string
  if counter > 0 {
    mes += moveCursorUp(len(task.Steps)+4)
  }

  mes += "\nSteps:\n"
  for _, step := range task.Steps {
    mes += fmt.Sprintf("%s\t%s\t%s\n", "\033[K", step.Name, step.Status)
  }
  mes += "\nStatus: In Progress"

  w := new(tabwriter.Writer)
  w.Init(os.Stdout, 5, 8, 4, '\t', 0)

  fmt.Fprintln(w, mes)
  w.Flush()

  if task.Status == "RUNNING" {
    time.Sleep(time.Millisecond * 100)
    return waitForTask(cli, ref, counter+1)
  }

  return fmt.Sprintf("%s%s%s%s", moveCursorUp(1), "\033[K", "Status: ", task.Status)
}

func moveCursorUp(lines int) string {
  return fmt.Sprintf("\033[%dA", lines)
}
