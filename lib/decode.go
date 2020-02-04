package lib

import (
	"github.com/k-onishi/nesgo"
)

// DecodeInfo ...
type DecodeInfo struct {
	Value       int
	Address     int
	Instruction *nesgo.Instruction
	Arg         *int
	Bytes       []byte
}

// AccessRange ...
type AccessRange struct {
	Min int
	Max int
}
