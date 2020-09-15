package goat

import (
	"errors"
	"fmt"
	"path/filepath"

	"ka/lang/compiler"
	"ka/lang/interpreter"
	"ka/lang/types"
	kastenc "kast/format/encoding"
)

//LoadLibrary decompiles an Oat file and store it in an Ka instance
func LoadLibrary(file string, params types.CliParams) (*types.Instance, error) {

	params.Name = file

	var kast map[string]*types.KaType
	var e error

	switch filepath.Ext(file) {
	case ".klr":
		return nil, errors.New("Kayl is unimplemented")
	case ".kast":
		kast, e = kastenc.KastDecode(file)
	default:
		kast, e = compiler.Compile(params)
	}

	return getinstance(kast), e
}

/*
	DefOp defines a new operation (note that operations are global to all instances in the current go program)
	Note that DefOp does not work with kayl oats, since the operation types are checked at compile time
*/
func DefOp(type1, type2, operation string, cb func(val1, val2 types.KaType, instance *types.Instance, stacktrace []string, line uint64, file string, stacksize uint) *types.KaType) {
	interpreter.Operations[fmt.Sprintf("%s %s %s", type1, operation, type2)] = cb
}
