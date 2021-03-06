/*
Copyright © 2021 NAME HERE <EMAIL ADDRESS>

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
	"fmt"

	"github.com/spf13/cobra"
	"github.com/amitkumardube/go-cli-cobra/common"
)


var file_name = common.Get_env("HOME")+"/.cc.json"
var new_project string

// setCmd represents the set command
var setCmd = &cobra.Command{
	Use:   "set",
	Short: "set utility metadata",
	Long: `Sets various utility metdata like project etc`,
	Run: func(cmd *cobra.Command, args []string) {
		common.Write_file(file_name , "cli.project_id",new_project)
		fmt.Println("Project ID "+new_project+" has been successfully set in CLI.")
	},
}

func init() {
	rootCmd.AddCommand(setCmd)


	// Here you will define your flags and configuration settings.

	// Cobra supports Persistent Flags which will work for this command
	// and all subcommands, e.g.:
	// setCmd.PersistentFlags().String("foo", "", "A help for foo")

	// Cobra supports local flags which will only run when this command
	// is called directly, e.g.:
	// setCmd.Flags().BoolP("toggle", "t", false, "Help message for toggle")

	setCmd.Flags().StringVarP(&new_project, "project-id", "p", "" , "Project ID or Number")
	setCmd.MarkFlagRequired("project-id")
}
