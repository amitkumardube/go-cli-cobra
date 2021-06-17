package common


type conf struct {
	Cli cli `json:"cli,omitempty"`
}

type cli struct {
	Project_id string`json:"project_id,omitempty"`
}
