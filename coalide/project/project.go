package project

import (
	"github.com/nathan-osman/coalide/coalide/builder"
	"github.com/nathan-osman/coalide/coalide/docker"
	"github.com/satori/go.uuid"
	"sort"
)

// A single project consisting of a collection of source files
type Project struct {
	Version  int
	UUID     string
	Name     string
	Type     string
	Packages []string
	Options  map[string]string
	builder  builder.Builder
}

// Create a new project with the provided information
func NewProject(name string, buildType string, packages []string, options map[string]string) *Project {
	return &Project{
		Version:  1,
		UUID:     uuid.NewV4().String(),
		Name:     name,
		Type:     buildType,
		Packages: packages,
		Options:  options,
	}
}

// Attempt to build the project
func (p *Project) Build(docker *docker.Docker) error {

	// Make sure the container exists and is up-to-date
	if err := p.prepareContainer(docker); err != nil {
		return err
	}

	// Make sure that the correct builder exists
	if err := p.prepareBuilder(); err != nil {
		return err
	}

	// Build the project
	if err := p.builder.Build(docker); err != nil {
		return err
	}

	return nil
}

// Attempt to run the project's executable
func (p *Project) Run(docker *docker.Docker) error {

	// Ensure that the project is built
	if err := p.Build(docker); err != nil {
		return err
	}

	// Run the executable
	if err := p.builder.Run(docker, p.Options); err != nil {
		return err
	}

	return nil
}

// Ensure that the container exists and is up-to-date
func (p *Project) prepareContainer(docker *docker.Docker) error {

	// Check to see if the container exists
	exists, err := docker.ContainerExists(p.UUID)
	if err != nil {
		return err
	}

	// If it exists, check to make sure the packages are up to date
	if exists {
		packages, err := docker.ContainerPackages(p.UUID)
		if err != nil {
			return err
		}

		// If the packages aren't up to date, destroy the container
		if !p.packagesEqual(packages, p.Packages) {
			exists = false
			if err := docker.RemoveContainer(p.UUID); err != nil {
				return err
			}
		}
	}

	// If it does NOT exist, then it must be created
	if !exists {
		if err := docker.CreateContainer(p.UUID, p.Packages); err != nil {
			return err
		}
	}

	return nil
}

// Ensure that the correct builder exists for the project
func (p *Project) prepareBuilder() error {

	// If a builder exists, check to see if it is correct
	if p.builder == nil || p.Type != p.builder.Type() {

		// Create the correct type of builder
		if b, err := builder.New(p.Type, p.Options); err != nil {
			return err
		} else {
			p.builder = b
		}
	}

	return nil
}

// Compare two lists of packages
func (p *Project) packagesEqual(a, b []string) bool {

	// Make an early exit if the slice length is different
	if len(a) != len(b) {
		return false
	}

	// Sort the slices
	aSorted := make([]string, len(a))
	bSorted := make([]string, len(b))

	copy(aSorted, a)
	copy(bSorted, b)

	sort.Strings(aSorted)
	sort.Strings(bSorted)

	// Compare each of the items
	for i, _ := range aSorted {
		if a[i] != b[i] {
			return false
		}
	}

	return true
}
