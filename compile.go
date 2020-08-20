package goat

import (
	"io/ioutil"

	"github.com/omm-lang/omm/lang/compiler"
	oatenc "github.com/omm-lang/omm/oat/encoding"

	. "github.com/omm-lang/omm/lang/types"
)

//Compile a Omm file
func CompileFile(filename string) (map[string][]Action, compiler.CompileErr) {
	f, _ := ioutil.ReadFile(filename)
	vars, e := compiler.Compile(string(f), "goat compile", false, true)
	return vars, e
}

//Compile an Omm string
func CompileString(script string) (map[string][]Action, compiler.CompileErr) {
	vars, e := compiler.Compile(script, "goat compile", false, true)
	return vars, e
}

//Decompile an Oat file and store it in a map[string][]types.Action.
func GetOat(file string) (map[string][]Action, error) {
	return oatenc.OatDecode(file, 0)
}
