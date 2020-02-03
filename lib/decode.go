package lib

import (
	"github.com/k-onishi/nesgo"
)

// DecodeInfo ...
type DecodeInfo struct {
	Address     int
	Instruction nesgo.Instruction
	Arg         int
}

// AccessRange ...
type AccessRange struct {
	Min int
	Max int
}
