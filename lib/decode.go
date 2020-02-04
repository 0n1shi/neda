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
	isEndOfSub  bool
}

// AccessRange ...
type AccessRange struct {
	Min int
	Max int
}
