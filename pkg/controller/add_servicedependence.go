package controller

import (
	"github.com/VTN-operator/pkg/controller/servicedependence"
)

func init() {
	// AddToManagerFuncs is a list of functions to create controllers and add them to a manager.
	AddToManagerFuncs = append(AddToManagerFuncs, servicedependence.Add)
}
