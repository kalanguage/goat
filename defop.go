package goat

import (
	"github.com/tusklang/tusk/lang/interpreter"
	"github.com/tusklang/tusk/lang/types"
	"github.com/tusklang/tusk/native"
)

//DefOp defines a new operator overloader
func DefOp(
	op string,
	type1 string,
	type2 string,
	opfn func(
		val1 types.TuskType,
		val2 types.TuskType,
		tuskpanic func(
			msg string,
			code int,
		) *types.TuskError,
	) (*types.TuskType, *types.TuskError),
) {
	interpreter.Operations[type1+" "+op+" "+type2] = func(val1, val2 types.TuskType, instance *types.Instance, stacktrace []string, line uint64, file string, stacksize uint, namespace string) (*types.TuskType, *types.TuskError, string) {
		var tuskpanicPasser = func(msg string, code int) *types.TuskError {
			return native.TuskPanic(msg, line, file, stacktrace, code)
		}

		ret, e := opfn(val1, val2, tuskpanicPasser)

		return ret, e, ""
	}
}
