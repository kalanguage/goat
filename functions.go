package goat

import (
	"errors"
	"fmt"

	"github.com/tusklang/tusk/lang/interpreter"

	"github.com/tusklang/tusk/lang/types"
)

//CallOatFunc calls a (global) functions in a given instance and returns the return value
func CallOatFunc(instance *types.Instance, fname string, args ...*types.TuskType) (*types.TuskType, error) {
	fnvar := instance.Fetch("$" + fname)

	if fnvar == nil {
		return nil, fmt.Errorf("Could not find variable %s", fname)
	}

	fn := fnvar.Value

	if (*fn).Type() != "function" {
		return nil, errors.New("Given variable is not a function")
	}

	var argv types.TuskArray

	for _, v := range args {
		argv.PushBack(*v)
	}

	return interpreter.Operations["function <- array"](*fn, argv, instance, []string{"at goat caller "}, 0, "github.com/tusklang/goat", 0), nil
}
