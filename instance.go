package goat

import (
	"os"

	"github.com/omm-lang/omm/lang/interpreter"

	"github.com/omm-lang/omm/lang/types"
)

//NewInstance creates a new instance of oat given a map[string][]types.Action
func NewInstance(oat map[string][]types.Action) *types.Instance {
	var ins types.Instance
	dirname, _ := os.Getwd()
	interpreter.FillIns(&ins, oat, dirname, os.Args)
	return &ins
}
