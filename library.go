package goat

import (
	"errors"
	"fmt"

	oatenc "github.com/omm-lang/omm/oat/encoding"

	"github.com/omm-lang/omm/lang/interpreter"
	"github.com/omm-lang/omm/lang/types"
)

//LoadLibrary decompiles an Oat file and store it in a map[string][]types.Action.
func LoadLibrary(file string) (map[string][]types.Action, error) {
	return oatenc.OatDecode(file, 0)
}

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

//DefOp defines a new operation (note that operations are global to all instances in the current go program)
func DefOp(type1, type2, operation string, cb func(val1, val2 types.OmmType, instance *types.Instance, stacktrace []string, line uint64, file string, stacksize uint) *types.OmmType) {
	interpreter.Operations[fmt.Sprintf("%s %s %s", type1, operation, type2)] = cb
}
