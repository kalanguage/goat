package goat

import (
	"errors"
	"path/filepath"

	oatenc "github.com/tusklang/oat/format/encoding"
	"github.com/tusklang/tusk/lang/compiler"
	"github.com/tusklang/tusk/lang/types"
)

//LoadLibrary decompiles an Oat file and store it in an Tusk instance
func LoadLibrary(file string, params types.CliParams) (*types.Instance, error) {

	params.Name = file

	var oat map[string]*types.TuskType
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
