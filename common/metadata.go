package common

import (
	"cloud.google.com/go/compute/metadata"
)

// this function returns the default project Id for this compute engine instance

func Get_project_id()(string , error) {

        var project_id string
        var err error
	var file_name = Get_env("HOME")+"/.cc.json"

	conf := Read_file(file_name)
        if conf.Cli.Project_id == "" {
		// using the metadata library to get the current project ID if config file is empty
                project_id, err = metadata.ProjectID()

        }else {
		// else get the project ID from the file
                project_id = conf.Cli.Project_id
        }
        /*if err != nil {
                cobra.CheckErr(err.Error())
        }*/

	// using the metadata library to get the current project ID

	return project_id  , err
}
