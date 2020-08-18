package goat

import "lang/interpreter"
import "os"
import . "lang/types"

//export NewInstance
func NewInstance(oat map[string][]Action) *Instance {
  var ins Instance
  __dirname, _ := os.Getwd()
  interpreter.FillIns(&ins, oat, __dirname, os.Args)
  return &ins
}
