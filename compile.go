package goat

import (
	oatenc "github.com/omm-lang/omm/oat/encoding"

	. "github.com/omm-lang/omm/lang/types"
)

//LoadLibrary decompiles an Oat file and store it in a map[string][]types.Action.
func LoadLibrary(file string) (map[string][]Action, error) {
	return oatenc.OatDecode(file, 0)
}
