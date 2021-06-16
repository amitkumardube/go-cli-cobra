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
/*
import (
	"fmt"
	"os"
	"github.com/amitkumardube/go-cli-cobra/secret"
	"github.com/amitkumardube/go-cli-cobra/common"

	"github.com/spf13/cobra"
)

// getCmd represents the get command

var project , secret_name , version string
var getCmd = &cobra.Command{
	Use:   "get",
	Short: "Get the value of a secret version",
	Long: `Get the value of specific secret version`,
	Run: func(cmd *cobra.Command, args []string) {
		// logic of this command goes here

		// get the value of secret
		full_secret_name := "projects/"+project+"/secrets/"+secret_name+"/versions/"+version
                secret_value , err := secret.AccessSecretVersion(os.Stdout,full_secret_name)
		// checking for error. If error found then exit with status 1
                if err != nil {
                        cobra.CheckErr(err.Error())
                }
                fmt.Println(secret_value)
	},
}

func init() {
	secretCmd.AddCommand(getCmd)

	// Getting the default compute instance project_id

	project_id , err := common.Get_project_id()
	
	// checking for error. If error found then exit with status 1
	if err != nil {
		cobra.CheckErr(err.Error())
	}

	version = "latest"

	// Here you will define your flags and configuration settings.

	// Cobra supports Persistent Flags which will work for this command
	// and all subcommands, e.g.:
	// getCmd.PersistentFlags().String("foo", "", "A help for foo")

	// Cobra supports local flags which will only run when this command
	// is called directly, e.g.:
	// getCmd.Flags().BoolP("toggle", "t", false, "Help message for toggle")

	// setting up flags for project

	getCmd.Flags().StringVarP(&project, "project", "p", project_id, "Project ID or Name or Number")

	// setting up flags for secret

	getCmd.Flags().StringVarP(&secret_name, "secret", "s", "", "Secret Name")
	getCmd.MarkFlagRequired("secret")
	
	// setting up flags for version
	getCmd.Flags().StringVarP(&version, "version", "v", version , "Secret Version")
}

*/