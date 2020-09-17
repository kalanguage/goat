package goat

import (
	"errors"
	"fmt"
	"path/filepath"

	"github.com/tusklang/tusk/lang/compiler"
	"github.com/tusklang/tusk/lang/interpreter"
	"github.com/tusklang/tusk/lang/types"
	kastenc "github.com/tusklang/oat/format/encoding"
)

//LoadLibrary decompiles an Oat file and store it in an Tusk instance
func LoadLibrary(file string, params types.CliParams) (*types.Instance, error) {

	params.Name = file

	var oat map[string]*types.TuskType
	var e error

	switch filepath.Ext(file) {
	case ".klr":
		return nil, errors.New("Tuskyl is unimplemented")
	case ".oat":
		oat, e = kastenc.TuskstDecode(file)
	default:
		oat, e = compiler.Compile(params)
	}

	return getinstance(oat), e
}

/*
	DefOp defines a new operation (note that operations are global to all instances in the current go program)
	Note that DefOp does not work with kayl oats, since the operation types are checked at compile time
*/
func DefOp(type1, type2, operation string, cb func(val1, val2 types.TuskType, instance *types.Instance, stacktrace []string, line uint64, file string, stacksize uint) *types.TuskType) {
	interpreter.Operations[fmt.Sprintf("%s %s %s", type1, operation, type2)] = cb
}
