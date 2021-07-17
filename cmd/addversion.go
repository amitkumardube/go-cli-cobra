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
	"github.com/amitkumardube/go-cli-cobra/common"
	"github.com/amitkumardube/go-cli-cobra/secret"
	"os"
	"github.com/spf13/cobra"
)

var secret_key, project_num string

// addversionCmd represents the addversion command
var addversionCmd = &cobra.Command{
	Use:   "add-version",
	Short: "Add secret version in project",
	Long: `This command will add secret version to an existing secret in a GCP project.`,
	Run: func(cmd *cobra.Command, args []string) {
		full_secret_key_path := "projects/" + project_num + "/secrets/"+secret_key
		err := secret.AddSecretVersion(os.Stdout , full_secret_key_path)
		if err != nil {
			cobra.CheckErr(err.Error())
		}
	},
	SuggestFor: []string{"secret-value", "new-value"},
}

func init() {
	secretCmd.AddCommand(addversionCmd)

	// Getting the default compute instance project_id

	project_id, err := common.Get_project_id()	

	// checking for error. If error found then exit with status 1
	if err != nil {
		cobra.CheckErr(err.Error())
	}

	// Here you will define your flags and configuration settings.

	// Cobra supports Persistent Flags which will work for this command
	// and all subcommands, e.g.:
	// addversionCmd.PersistentFlags().String("foo", "", "A help for foo")

	// Cobra supports local flags which will only run when this command
	// is called directly, e.g.:
	// addversionCmd.Flags().BoolP("toggle", "t", false, "Help message for toggle")
	addversionCmd.Flags().StringVarP(&project_num, "project", "p", project_id, "Project ID or Number")
	addversionCmd.Flags().StringVarP(&secret_key, "secret-key", "s", "", "Secret Key")
	addversionCmd.MarkFlagRequired("secret-key")
}
