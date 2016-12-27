package client

import (
  "fmt"
  "os"
  "log"
  "net/http"
  "time"
  "crypto/tls"
  "io/ioutil"
  "encoding/json"
  "text/tabwriter"
  "github.com/torresdal/spinex/client/types"
)

//Client bla bla
type Client struct {
  host string
  x509CertFile string
  x509KeyFile string
}

//NewClient bla bla
func NewClient(host, x509CertFile, x509KeyFile string) *Client {
  return &Client {host: host, x509CertFile: x509CertFile, x509KeyFile: x509KeyFile}
}

//getHTTPClient returns a http.Client with credentials ready for use
func getHTTPClient(client *Client) *http.Client {
  cert1, err := tls.LoadX509KeyPair(client.x509CertFile, client.x509KeyFile)
  if err != nil {
    panic(err)
  }

  // Setup HTTPS client
  tlsConfig := &tls.Config{
    Certificates: []tls.Certificate{cert1},
  }
  tlsConfig.BuildNameToCertificate()

  transport := &http.Transport{TLSClientConfig: tlsConfig}
  return &http.Client{Transport: transport}
}

func checkErr(err error) {
    if err != nil {
        log.Fatal("ERROR:", err)
    }
}

func waitForTask(client *Client, ref string, counter int) string {
  if counter > 10 {
    return "Timed out waiting for task status"
  }

  task := Task(client, ref)

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
    return waitForTask(client, ref, counter+1)
  }

  return fmt.Sprintf("%s%s%s%s", moveCursorUp(1), "\033[K", "Status: ", task.Status)
}

func moveCursorUp(lines int) string {
  return fmt.Sprintf("\033[%dA", lines)
}

// Task will return info and status of a Spinnaker task
func Task(client *Client, ref string) types.TaskResponse {
  httpClient := getHTTPClient(client)
  resp, err := httpClient.Get(client.host + ref)
  defer resp.Body.Close()
  checkErr(err)

  data, err := ioutil.ReadAll(resp.Body)
  checkErr(err)

  var task types.TaskResponse
  err = json.Unmarshal([]byte(data), &task) // here!
  checkErr(err)

  return task
}
