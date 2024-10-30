package model

type SystemInfo struct {
	Hostname string `json:"hostname"`
	Version  string `json:"version"`
	Build    string `json:"build"`
	Commit   string `json:"commit"`
	Env      string `json:"env"`
}
