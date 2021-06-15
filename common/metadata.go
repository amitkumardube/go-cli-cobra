package common

import (
	"cloud.google.com/go/compute/metadata"
)

// this function returns the default project Id for this compute engine instance

func Get_project_id()(string , error) {

	// using the metadata library to get the current project ID

	var project_id, err = metadata.ProjectID()

	return project_id  , err
}
