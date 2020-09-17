package goat

import (
	"os"

	"tusk/lang/interpreter"

	"tusk/lang/types"
)

func getinstance(oat map[string]*types.TuskType) *types.Instance {

	if oat == nil {
		return nil
	}

	var ins types.Instance
	dirname, _ := os.Getwd()
	interpreter.FillIns(&ins, oat, dirname, os.Args)

	return &ins
}
