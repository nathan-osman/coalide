package builders

import (
	"errors"
)

// Generic interface for project builders
type Builder interface {
	Build() (bool, string)
}

// Recognized builders
const (
	Makefile = "makefile"
	CMake    = "cmake"
	Qmake    = "qmake"
)

// Create a builder for the specified type
func CreateBuilder(buildType string) (*Builder, error) {
	switch buildType {
	case Makefile, CMake, Qmake:
		return nil, errors.New("Not yet implemented")
	default:
		return nil, errors.New("Invalid builder specified")
	}
}
