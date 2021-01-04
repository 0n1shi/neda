package neda

type AddressingType int

const (
	AddressingTypeAccumulator = AddressingType(iota)
	AddressingTypeImplied
	AddressingTypeImmediate
	AddressingTypeZeroPage
	AddressingTypeZeroPageX
	AddressingTypeZeroPageY
	AddressingTypeRelative
	AddressingTypeAbsolute
	AddressingTypeAbsoluteX
	AddressingTypeAbsoluteY
	AddressingTypeIndirect
	AddressingTypeIndirectX
	AddressingTypeIndirectY
)
