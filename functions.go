package goat

import (
	"errors"
	"fmt"

	"omm/lang/interpreter"

	"omm/lang/types"
)

//CallOatFunc calls a (global) functions in a given instance and returns the return value
func CallOatFunc(instance *types.Instance, fname string, args ...*types.OmmType) (*types.OmmType, error) {
	fnvar := instance.Fetch("$" + fname)

	if fnvar == nil {
		return nil, fmt.Errorf("Could not find variable %s", fname)
	}

	fn := fnvar.Value

	if (*fn).Type() != "function" {
		return nil, errors.New("Given variable is not a function")
	}

	var argv types.OmmArray

	for _, v := range args {
		argv.PushBack(*v)
	}

	return interpreter.Operations["function <- array"](*fn, argv, instance, []string{"at goat caller "}, 0, "goat", 0), nil
}
