package goat

import "io/ioutil"
import "lang/compiler"
import "oat/encoding"
import . "lang/types"

//export CompileFile
func CompileFile(filename string) (map[string][]Action, compiler.CompileErr) {
  f, _ := ioutil.ReadFile(filename)
  vars, e := compiler.Compile(string(f), "goat compile", false, true)
  return vars, e
}

//export CompileString
func CompileString(script string) (map[string][]Action, compiler.CompileErr) {
  vars, e := compiler.Compile(script, "goat compile", false, true)
  return vars, e
}

//export GetOat
func GetOat(file string) (map[string][]Action, error) {
  return oatenc.OatDecode(file, 0)
}
