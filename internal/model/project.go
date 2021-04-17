package model

type Project struct {
	Name       string       `json:"project"`
	Components []Components `json:"components"`
}
type Components struct {
	Type       string `json:"type"`
	GroupID    string `json:"groupID"`
	ArtifactID string `json:"artifactID"`
	Version    string `json:"version"`
	SvmID      int    `json:"svmID"`
	SourcesURL string `json:"sources"`
}
