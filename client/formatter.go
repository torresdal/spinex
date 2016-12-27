package client

import (
  "os"
  "fmt"
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
