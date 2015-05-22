package project

import (
	"encoding/json"
	"github.com/nathan-osman/coalide/coalide/builders"
)

// A single project consisting of a collection of source files
type Project struct {
	Path    string
	Builder *builders.Builder
	Name    string
}

// Create a new project with the provided information
func NewProject(path string, buildType string, name string) (*Project, error) {
	builder, err := builders.CreateBuilder(buildType)
	if err != nil {
		return nil, err
	}
	return &Project{
		Path:    path,
		Builder: builder,
		Name:    name,
	}, nil
}

// Attempt to load a project from JSON data
func NewProjectFromJSON(path string, jsonData []byte) (*Project, error) {

	// A separate struct is used for unmarshalling
	type ProjectData struct {
		Name      string `json:"name"`
		BuildType string `json:"build_type"`
	}
	projectData := &ProjectData{}

	// Attempt to unmarshall the JSON data into the struct
	if err := json.Unmarshal(jsonData, projectData); err != nil {
		return nil, err
	}

	// Use the NewProject constructor to initialize the project
	return NewProject(path, projectData.BuildType, projectData.Name)
}
