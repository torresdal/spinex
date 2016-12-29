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

package client

// Config contains all configuration settings for Spinex
type Config struct {
  X509CertFile    string
  X509KeyFile     string
  Host            string
}

// NewConfig will create a new config based on spinex.yml config file
func NewConfig(conf map[string]string) *Config {
  return &Config { Host: conf["host"], X509CertFile: conf["x509certfile"], X509KeyFile: conf["x509keyfile"]}
}
