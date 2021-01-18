package goat

import (
	"fmt"

	"github.com/tusklang/tusk/lang/interpreter"

	"github.com/tusklang/tusk/lang/types"
)

//CallOatFunc calls a (global) functions in a given instance and returns the return value
func CallOatFunc(instance *types.Instance, fname string, args ...*types.TuskType) (*types.TuskType, *types.TuskError) {
	fnvar := instance.Fetch(fname)

	if fnvar == nil {
		var tmp types.TuskError
		tmp.Err = fmt.Sprintf("Could not find variable %s", fname)
		return nil, &tmp
	}

	fn := fnvar.Value

	if (*fn).Type() != "function" {
		var tmp types.TuskError
		tmp.Err = fmt.Sprintf("Could not find variable %s", fname)
		return nil, &tmp
	}

	var argv types.TuskArray

	for _, v := range args {
		argv.PushBack(*v)
	}

	a, b, _ := interpreter.Operations["function <- array"](*fn, argv, instance, []string{"at goat caller "}, 0, "github.com/tusklang/goat", 0, "global")

	return a, b
}
