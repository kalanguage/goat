package goat

import (
	"os"

	"github.com/omm-lang/omm/lang/interpreter"

	. "github.com/omm-lang/omm/lang/types"
)

//export NewInstance
func NewInstance(oat map[string][]Action) *Instance {
	var ins Instance
	__dirname, _ := os.Getwd()
	interpreter.FillIns(&ins, oat, __dirname, os.Args)
	return &ins
}
