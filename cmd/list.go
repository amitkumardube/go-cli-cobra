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
	"os"
	"github.com/spf13/cobra"

	"github.com/amitkumardube/go-cli-cobra/common"
	"github.com/amitkumardube/go-cli-cobra/secret"
)

var project string

// listCmd represents the list command
var listCmd = &cobra.Command{
	Use:   "list",
	Short: "List secrets in project",
	Long:  `This command will list secrets in the project.`,
	Run: func(cmd *cobra.Command, args []string) {

		// logic of this command goes here

		// get the list of secrets
		full_project_name := "projects/" + project
		err := secret.ListSecrets(os.Stdout, full_project_name)
		// checking for error. If error found then exit with status 1
		if err != nil {
			cobra.CheckErr(err.Error())
		}
	},
}

func init() {
	secretCmd.AddCommand(listCmd)

	// Getting the default compute instance project_id
	// if project_id is not set then consider using the default project id where is compute instance lies

	var project_id string
	var err error
	conf := common.Read_file(file_name)
	if conf.Cli.Project_id == "" {
		project_id, err = common.Get_project_id()
		
	}else {
		project_id = conf.Cli.Project_id
	}
	if err != nil {
		cobra.CheckErr(err.Error())
	}

	// Here you will define your flags and configuration settings.

	// Cobra supports Persistent Flags which will work for this command
	// and all subcommands, e.g.:
	// listCmd.PersistentFlags().String("foo", "", "A help for foo")

	// Cobra supports local flags which will only run when this command
	// is called directly, e.g.:
	// listCmd.Flags().BoolP("toggle", "t", false, "Help message for toggle")

	// setting up flags for project
	listCmd.Flags().StringVarP(&project, "project", "p", project_id, "Project ID or Number")

}
