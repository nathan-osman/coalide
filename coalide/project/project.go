package project

import (
	"github.com/nathan-osman/coalide/coalide/docker"
	"regexp"
	"sort"
	"strings"
)

// A single project consisting of a collection of source files
type Project struct {
	Name     string            `json:"name"`
	Type     string            `json:"type"`
	Packages []string          `json:"packages"`
	Options  map[string]string `json:"options"`
}

// Matches all non-safe characters in a project name
var unsafeChars = regexp.MustCompile("\\W")

// Create a new project with the provided information
func NewProject(name string, buildType string, packages []string, options map[string]string) *Project {
	return &Project{
		Name:     name,
		Type:     buildType,
		Packages: packages,
		Options:  options,
	}
}

// Attempt to build the project
func (p *Project) Build(docker *docker.Docker) error {

	// Make sure the container exists and is up-to-date
	if err := p.checkContainer(docker); err != nil {
		return err
	}

	// TODO: Invoke the appropriate builder in the container

	return nil
}

func (p *Project) Run(docker *docker.Docker) error {

	// TODO: Invoke the executable from within the container

	return nil
}

// Ensure that the container exists and is up-to-date
func (p *Project) checkContainer(docker *docker.Docker) error {

	// Check to see if the container exists
	exists, err := docker.ContainerExists(p.containerName())
	if err != nil {
		return err
	}

	// If it exists, check to make sure the packages are up to date
	if exists {
		packages, err := docker.ContainerPackages(p.containerName())
		if err != nil {
			return err
		}

		// If the packages aren't up to date, destroy the container
		if p.packagesEqual(packages, p.Packages) {
			exists = false
			if err := docker.RemoveContainer(p.containerName()); err != nil {
				return err
			}
		}
	}

	// If it does NOT exist, then it must be created
	if !exists {
		if err := docker.CreateContainer(p.Name, p.Packages); err != nil {
			return err
		}
	}

	return nil
}

// Generate a "safe" representation of the project name
func (p *Project) safeName() string {
	return strings.ToLower(unsafeChars.ReplaceAllString(p.Name, "-"))
}

// Generate a name for the container that will run the executables
func (p *Project) containerName() string {
	return "coalide-" + p.safeName()
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
