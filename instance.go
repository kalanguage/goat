package goat

import (
	"os"

	"omm/lang/interpreter"

	"omm/lang/types"
)

func getinstance(oat map[string]*types.OmmType) *types.Instance {

	if oat == nil {
		return nil
	}

	var ins types.Instance
	dirname, _ := os.Getwd()
	interpreter.FillIns(&ins, oat, dirname, os.Args)

	return &ins
}
