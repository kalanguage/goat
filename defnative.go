package goat

import (
	"github.com/tusklang/tusk/lang/types"
	"github.com/tusklang/tusk/native"
)

//DefNative defines a new Tusk native function
func DefNative(
	name string,
	sigs [][]string,
	f func(
		args []*types.TuskType,
		tuskpanic Panic,
	) (*types.TuskType, *types.TuskError),
) {

	var gf types.TuskType = native.TuskGoFunc{
		Function: func(args []*types.TuskType, stacktrace []string, line uint64, file string, instance *types.Instance) (*types.TuskType, *types.TuskError) {
			var tuskpanicPasser Panic = func(msg string, code int) *types.TuskError {
				return native.TuskPanic(msg, line, file, stacktrace, code)
			}

			return f(args, tuskpanicPasser)
		},
		Signatures: sigs,
	}

	native.Native[name] = &gf
}
