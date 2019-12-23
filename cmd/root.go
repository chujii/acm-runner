/*
Copyright © 2019 NAME HERE <EMAIL ADDRESS>

Licensed under the Apache License, Version 2.0 (the "License");
you may not use this file except in compliance with the License.
You may obtain a copy of the License at

    http://www.apache.org/licenses/LICENSE-2.0

Unless required by applicable law or agreed to in writing, software
distributed under the License is distributed on an "AS IS" BASIS,
WITHOUT WARRANTIES OR CONDITIONS OF ANY KIND, either express or implied.
See the License for the specific language governing permissions and
limitations under the License.
*/
package cmd

import (
  "acm-runner/handler"
  "fmt"
  "github.com/spf13/cobra"
  "io/ioutil"
  "log"
  "os"
)

var cfgFile string

// rootCmd represents the base command when called without any subcommands
var rootCmd = &cobra.Command{
  Use:   "acm-runner",
  Short: "acm runner client",
  Long: `Acm runner is a client for aliyun configuration management, listen config change`,
  Run: runCmd,
}

// Execute adds all child commands to the root command and sets flags appropriately.
// This is called by main.main(). It only needs to happen once to the rootCmd.
func Execute() {
  if err := rootCmd.Execute(); err != nil {
    fmt.Println(err)
    os.Exit(1)
  }
}

func init() {
  cobra.OnInitialize(func() {
      handler.InitConfig(cfgFile)
  })
  rootCmd.Flags().StringVarP(&cfgFile,"config", "c", "./acm-runner.yaml.example", "config file (default is ./acm-runner.yaml.example)")
}

func runCmd(cmd *cobra.Command, args []string)  {
  nc, err := handler.NewNocos(handler.Conf["namespace"].(map[string]interface{}))
  if err != nil {
   fmt.Println(err)
   os.Exit(1)
  }

  nc.ListenConfig(handler.Conf["list"].([]interface{}), func(data string, filename string) {
    fmt.Println(data)
    fmt.Println(filename)
    if err := ioutil.WriteFile(filename, []byte(data), 0666); err != nil {
      log.Fatalln(err)
    }
  })

  <- make(<-chan bool)
}

