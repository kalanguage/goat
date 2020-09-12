package goat

import (
	"errors"
	"fmt"
	"path/filepath"

	oatenc "oat/format/encoding"
	"omm/lang/compiler"
	"omm/lang/interpreter"
	"omm/lang/types"
)

//LoadLibrary decompiles an Oat file and store it in an Omm instance
func LoadLibrary(file string, params types.CliParams) (*types.Instance, error) {

	params.Name = file

	var oat map[string]*types.OmmType
	var e error

	switch filepath.Ext(file) {
	case ".klr":
		return nil, errors.New("Kayl is unimplemented")
	case ".oat":
		oat, e = oatenc.OatDecode(file)
	default:
		oat, e = compiler.Compile(params)
	}

	return getinstance(oat), e
}

/*
	DefOp defines a new operation (note that operations are global to all instances in the current go program)
	Note that DefOp does not work with kayl oats, since the operation types are checked at compile time
*/
func DefOp(type1, type2, operation string, cb func(val1, val2 types.OmmType, instance *types.Instance, stacktrace []string, line uint64, file string, stacksize uint) *types.OmmType) {
	interpreter.Operations[fmt.Sprintf("%s %s %s", type1, operation, type2)] = cb
}
