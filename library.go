package goat

import (
	"errors"
	"fmt"
	"path/filepath"

	oatenc "github.com/omm-lang/oat/format/encoding"
	"github.com/omm-lang/omm/lang/compiler"
	"github.com/omm-lang/omm/lang/interpreter"
	"github.com/omm-lang/omm/lang/types"
)

//LoadLibrary decompiles an Oat file and store it in an Omm instance
func LoadLibrary(file string, params types.CliParams) (*types.Instance, error) {

	params.Name = file

	var oat map[string]*types.OmmType
	var e error

	switch filepath.Ext(file) {
	case ".omm":
		oat, e = compiler.Compile(params)
	case ".klr":
		return nil, errors.New("Kayl is unimplemented")
	case ".oat":
		fallthrough
	default:
		oat, e = oatenc.OatDecode(file)
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
