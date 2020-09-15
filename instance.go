package goat

import (
	"os"

	"ka/lang/interpreter"

	"ka/lang/types"
)

func getinstance(kast map[string]*types.KaType) *types.Instance {

	if kast == nil {
		return nil
	}

	var ins types.Instance
	dirname, _ := os.Getwd()
	interpreter.FillIns(&ins, kast, dirname, os.Args)

	return &ins
}
