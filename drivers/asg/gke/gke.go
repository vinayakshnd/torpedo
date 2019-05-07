package gke

import (
	"github.com/portworx/torpedo/drivers/asg"
	"github.com/sirupsen/logrus"
)

const (
	// DriverName is the name of the GKE ASG driver
	DriverName = "gke"
)

type gke struct {
}

func (g *gke) Init() error {

	logrus.Info("Initializing GKE client.")
	return nil
}

func init() {
	g := &gke{}
	asg.Register(DriverName, g)
}
