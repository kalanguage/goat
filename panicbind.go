package goat

import "github.com/tusklang/tusk/lang/types"

//Panic represents a panic function used in defining usable go code
type Panic func(msg string, code int) *types.TuskError
