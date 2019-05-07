package asg

import (
	"fmt"

	"github.com/portworx/torpedo/pkg/errors"
)

var (
	asgDrivers = make(map[string]Driver)
)

type asg struct {
}

// Driver provides the asg driver interface
type Driver interface {
	// Init initializes the asg driver under the given scheduler
	Init() error
}

// Register registers the given asg driver
func Register(name string, d Driver) error {
	if _, ok := asgDrivers[name]; !ok {
		asgDrivers[name] = d
	} else {
		return fmt.Errorf("asg driver: %s is already registered", name)
	}

	return nil
}

// Get returns a registered asg driver
func Get(name string) (Driver, error) {
	if d, ok := asgDrivers[name]; ok {
		return d, nil
	}
	return nil, &errors.ErrNotFound{
		ID:   name,
		Type: "Node Driver",
	}
}
