package goat

import (
	"os"

	"github.com/omm-lang/omm/lang/interpreter"

	. "github.com/omm-lang/omm/lang/types"
)

//Create a new instance of oat given a map[string][]types.Action
func NewInstance(oat map[string][]Action) *Instance {
	var ins Instance
	__dirname, _ := os.Getwd()
	interpreter.FillIns(&ins, oat, __dirname, os.Args)
	return &ins
}
