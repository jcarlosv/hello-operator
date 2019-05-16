package controller

import (
	"github.com/jcarlosv/hello-operator/pkg/controller/carlos"
)

func init() {
	// AddToManagerFuncs is a list of functions to create controllers and add them to a manager.
	AddToManagerFuncs = append(AddToManagerFuncs, carlos.Add)
}
