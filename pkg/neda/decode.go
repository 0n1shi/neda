package neda

// Address ...
type Address int

// DecodeInfo ...
type DecodeInfo struct {
	Value       int
	Address     Address
	Instruction *Instruction
	Arg         *int
	Bytes       []byte
	isEndOfSub  bool
}

// AccessRange ...
type AccessRange struct {
	Min       Address
	Max       Address
	IsInvalid bool
}

// AnalysisInfo ...
type AnalysisInfo struct {
	decodeInfoMap   map[Address]*DecodeInfo
	accessRangeList []*AccessRange
}
