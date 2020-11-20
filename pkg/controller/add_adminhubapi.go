package controller

import (
	"github.com/ibm/ibm-adminhub-api-operator/pkg/controller/adminhubapi"
)

func init() {
	// AddToManagerFuncs is a list of functions to create controllers and add them to a manager.
	AddToManagerFuncs = append(AddToManagerFuncs, adminhubapi.Add)
}
