package builder

import (
	"errors"
	"github.com/nathan-osman/coalide/coalide/docker"
)

// Builder for a particular type of project
type Builder interface {
	Type() string
	Build(docker *docker.Docker) error
	Run(docker *docker.Docker, options map[string]string) error
}

// Recognized build types
const (
	Makefile = "makefile"
	CMake    = "cmake"
	Qmake    = "qmake"
)

// Create a builder for the specified type
func New(buildType string, options map[string]string) (Builder, error) {
	switch buildType {
	default:
		return nil, errors.New("Invalid builder specified")
	}
}
