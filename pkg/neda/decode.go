package neda

type Address int

type DecodeInfo struct {
	Value       int
	Address     Address
	Instruction *Instruction
	Arg         *int
	Bytes       []byte
	isEndOfSub  bool
}

type AccessRange struct {
	Min       Address
	Max       Address
	IsInvalid bool
}

type AnalysisInfo struct {
	decodeInfoMap   map[Address]*DecodeInfo
	accessRangeList []*AccessRange
}
