package nesrom

// OpecodeType ...
type OpecodeType string

const (
	OpecodeLDA = "LDA"
	OpecodeLDX = "LDX"
	OpecodeLDY = "LDY"
	OpecodeSTA = "STA"
	OpecodeSTX = "STX"
	OpecodeSTY = "STY"
	OpecodeTXA = "TXA"
	OpecodeTYA = "TYA"
	OpecodeTXS = "TXS"
	OpecodeTAY = "TAY"
	OpecodeTAX = "TAX"
	OpecodeTSX = "TSX"
	OpecodePHA = "PHA"
	OpecodePHP = "PHP"
	OpecodePLA = "PLA"
	OpecodePLP = "PLP"
	OpecodeADC = "ADC"
	OpecodeSBC = "SBC"
	OpecodeCPX = "CPX"
	OpecodeCPY = "CPY"
	OpecodeCMP = "CMP"
	OpecodeAND = "AND"
	OpecodeEOR = "EOR"
	OpecodeORA = "ORA"
	OpecodeBIT = "BIT"
	OpecodeASL = "ASL"
	OpecodeLSR = "LSR"
	OpecodeROL = "ROL"
	OpecodeROR = "ROR"
	OpecodeINX = "INX"
	OpecodeINY = "INY"
	OpecodeINC = "INC"
	OpecodeDEX = "DEX"
	OpecodeDEY = "DEY"
	OpecodeDEC = "DEC"
	OpecodeCLC = "CLC"
	OpecodeCLI = "CLI"
	OpecodeCLV = "CLV"
	OpecodeCLD = "CLD"
	OpecodeSEC = "SEC"
	OpecodeSEI = "SEI"
	OpecodeSED = "SED"
	OpecodeJSR = "JSR"
	OpecodeJMP = "JMP"
	OpecodeRTI = "RTI"
	OpecodeRTS = "RTS"
	OpecodeBCC = "BCC"
	OpecodeBCS = "BCS"
	OpecodeBEQ = "BEQ"
	OpecodeBMI = "BMI"
	OpecodeBNE = "BNE"
	OpecodeBPL = "BPL"
	OpecodeBVC = "BVC"
	OpecodeBVS = "BVS"
	OpecodeNOP = "NOP"
	OpecodeBRK = "BRK"
	OpecodeUNK = "unknown" // opecode which does't exist
)

// AddressingType ...
type AddressingType string

const (
	AddressingTypeAccumulator = "Accumulator"
	AddressingTypeImplied     = "Implied"
	AddressingTypeImmediate   = "Immediate"
	AddressingTypeZeroPage    = "ZeroPage"
	AddressingTypeZeroPageX   = "ZeroPageX"
	AddressingTypeZeroPageY   = "ZeroPageY"
	AddressingTypeRelative    = "Relative"
	AddressingTypeAbsolute    = "Absolute"
	AddressingTypeAbsoluteX   = "AbsoluteX"
	AddressingTypeAbsoluteY   = "AbsoluteY"
	AddressingTypeIndirect    = "Indirect"
	AddressingTypeIndirectX   = "IndirectX"
	AddressingTypeIndirectY   = "IndirectY"
)

// Instruction ...
type Instruction struct {
	opecodeType    OpecodeType
	addressingType AddressingType
	bytes          int
	cycle          int
}

// InstructionMap ...
var InstructionMap = map[int]Instruction{
	0x00: Instruction{
		opecodeType:    OpecodeBRK,
		addressingType: AddressingTypeImplied,
		bytes:          1,
		cycle:          7,
	},
	0x01: {
		opecodeType:    OpecodeORA,
		addressingType: AddressingTypeIndirectX,
		bytes:          2,
		cycle:          6,
	},
	0x05: {
		opecodeType:    OpecodeORA,
		addressingType: AddressingTypeZeroPage,
		bytes:          2,
		cycle:          3,
	},
	0x06: {
		opecodeType:    OpecodeASL,
		addressingType: AddressingTypeZeroPage,
		bytes:          2,
		cycle:          5,
	},
	0x08: {
		opecodeType:    OpecodePHP,
		addressingType: AddressingTypeImplied,
		bytes:          1,
		cycle:          3,
	},
	0x09: {
		opecodeType:    OpecodeORA,
		addressingType: AddressingTypeImmediate,
		bytes:          2,
		cycle:          2,
	},
	0x0a: {
		opecodeType:    OpecodeASL,
		addressingType: AddressingTypeAccumulator,
		bytes:          1,
		cycle:          2,
	},
	0x0d: {
		opecodeType:    OpecodeORA,
		addressingType: AddressingTypeAbsolute,
		bytes:          3,
		cycle:          4,
	},
	0x0e: {
		opecodeType:    OpecodeASL,
		addressingType: AddressingTypeAbsolute,
		bytes:          3,
		cycle:          6,
	},
	0x10: {
		opecodeType:    OpecodeBPL,
		addressingType: AddressingTypeRelative,
		bytes:          2,
		cycle:          2,
	},
	0x11: {
		opecodeType:    OpecodeORA,
		addressingType: AddressingTypeIndirectY,
		bytes:          2,
		cycle:          5,
	},
	0x15: {
		opecodeType:    OpecodeORA,
		addressingType: AddressingTypeZeroPageX,
		bytes:          2,
		cycle:          4,
	},
	0x16: {
		opecodeType:    OpecodeASL,
		addressingType: AddressingTypeZeroPageX,
		bytes:          2,
		cycle:          6,
	},
	0x18: {
		opecodeType:    OpecodeCLC,
		addressingType: AddressingTypeImplied,
		bytes:          1,
		cycle:          2,
	},
	0x19: {
		opecodeType:    OpecodeORA,
		addressingType: AddressingTypeAbsoluteY,
		bytes:          3,
		cycle:          4,
	},
	0x1d: {
		opecodeType:    OpecodeORA,
		addressingType: AddressingTypeAbsoluteX,
		bytes:          3,
		cycle:          4,
	},
	0x1e: {
		opecodeType:    OpecodeASL,
		addressingType: AddressingTypeAbsoluteX,
		bytes:          3,
		cycle:          7,
	},
	0x20: {
		opecodeType:    OpecodeJSR,
		addressingType: AddressingTypeAbsolute,
		bytes:          3,
		cycle:          6,
	},
	0x21: {
		opecodeType:    OpecodeAND,
		addressingType: AddressingTypeIndirectX,
		bytes:          2,
		cycle:          6,
	},
	0x24: {
		opecodeType:    OpecodeBIT,
		addressingType: AddressingTypeZeroPage,
		bytes:          2,
		cycle:          3,
	},
	0x25: {
		opecodeType:    OpecodeAND,
		addressingType: AddressingTypeZeroPage,
		bytes:          2,
		cycle:          3,
	},
	0x26: {
		opecodeType:    OpecodeROL,
		addressingType: AddressingTypeZeroPage,
		bytes:          2,
		cycle:          5,
	},
	0x28: {
		opecodeType:    OpecodePLP,
		addressingType: AddressingTypeImplied,
		bytes:          1,
		cycle:          4,
	},
	0x29: {
		opecodeType:    OpecodeAND,
		addressingType: AddressingTypeImmediate,
		bytes:          2,
		cycle:          2,
	},
	0x2a: {
		opecodeType:    OpecodeROL,
		addressingType: AddressingTypeAccumulator,
		bytes:          1,
		cycle:          2,
	},
	0x2c: {
		opecodeType:    OpecodeBIT,
		addressingType: AddressingTypeAbsolute,
		bytes:          3,
		cycle:          4,
	},
	0x2d: {
		opecodeType:    OpecodeAND,
		addressingType: AddressingTypeAbsolute,
		bytes:          3,
		cycle:          4,
	},
	0x2e: {
		opecodeType:    OpecodeROL,
		addressingType: AddressingTypeAbsolute,
		bytes:          3,
		cycle:          6,
	},
	0x30: {
		opecodeType:    OpecodeBMI,
		addressingType: AddressingTypeRelative,
		bytes:          2,
		cycle:          2,
	},
	0x31: {
		opecodeType:    OpecodeAND,
		addressingType: AddressingTypeIndirectY,
		bytes:          2,
		cycle:          5,
	},
	0x35: {
		opecodeType:    OpecodeAND,
		addressingType: AddressingTypeZeroPageX,
		bytes:          2,
		cycle:          4,
	},
	0x36: {
		opecodeType:    OpecodeROL,
		addressingType: AddressingTypeZeroPageX,
		bytes:          2,
		cycle:          6,
	},
	0x38: {
		opecodeType:    OpecodeSEC,
		addressingType: AddressingTypeImplied,
		bytes:          1,
		cycle:          2,
	},
	0x39: {
		opecodeType:    OpecodeAND,
		addressingType: AddressingTypeAbsoluteY,
		bytes:          3,
		cycle:          4,
	},
	0x3d: {
		opecodeType:    OpecodeAND,
		addressingType: AddressingTypeAbsoluteX,
		bytes:          3,
		cycle:          4,
	},
	0x3e: {
		opecodeType:    OpecodeROL,
		addressingType: AddressingTypeAbsoluteX,
		bytes:          3,
		cycle:          7,
	},
	0x40: {
		opecodeType:    OpecodeRTI,
		addressingType: AddressingTypeImplied,
		bytes:          1,
		cycle:          6,
	},
	0x41: {
		opecodeType:    OpecodeEOR,
		addressingType: AddressingTypeIndirectX,
		bytes:          2,
		cycle:          6,
	},
	0x45: {
		opecodeType:    OpecodeEOR,
		addressingType: AddressingTypeZeroPage,
		bytes:          2,
		cycle:          3,
	},
	0x46: {
		opecodeType:    OpecodeLSR,
		addressingType: AddressingTypeZeroPage,
		bytes:          2,
		cycle:          5,
	},
	0x48: {
		opecodeType:    OpecodePHA,
		addressingType: AddressingTypeImplied,
		bytes:          1,
		cycle:          3,
	},
	0x49: {
		opecodeType:    OpecodeEOR,
		addressingType: AddressingTypeImmediate,
		bytes:          2,
		cycle:          3,
	},
	0x4a: {
		opecodeType:    OpecodeLSR,
		addressingType: AddressingTypeAccumulator,
		bytes:          1,
		cycle:          2,
	},
	0x4c: {
		opecodeType:    OpecodeJMP,
		addressingType: AddressingTypeAbsolute,
		bytes:          3,
		cycle:          3,
	},
	0x4d: {
		opecodeType:    OpecodeEOR,
		addressingType: AddressingTypeAbsolute,
		bytes:          3,
		cycle:          4,
	},
	0x4e: {
		opecodeType:    OpecodeLSR,
		addressingType: AddressingTypeAbsolute,
		bytes:          3,
		cycle:          6,
	},
	0x50: {
		opecodeType:    OpecodeBVC,
		addressingType: AddressingTypeRelative,
		bytes:          2,
		cycle:          2,
	},
	0x51: {
		opecodeType:    OpecodeEOR,
		addressingType: AddressingTypeIndirectY,
		bytes:          2,
		cycle:          5,
	},
	0x55: {
		opecodeType:    OpecodeEOR,
		addressingType: AddressingTypeZeroPageX,
		bytes:          2,
		cycle:          4,
	},
	0x56: {
		opecodeType:    OpecodeLSR,
		addressingType: AddressingTypeZeroPageX,
		bytes:          2,
		cycle:          6,
	},
	0x58: {
		opecodeType:    OpecodeCLI,
		addressingType: AddressingTypeImplied,
		bytes:          1,
		cycle:          2,
	},
	0x59: {
		opecodeType:    OpecodeEOR,
		addressingType: AddressingTypeAbsoluteY,
		bytes:          3,
		cycle:          4,
	},
	0x5d: {
		opecodeType:    OpecodeEOR,
		addressingType: AddressingTypeAbsoluteX,
		bytes:          3,
		cycle:          4,
	},
	0x5e: {
		opecodeType:    OpecodeLSR,
		addressingType: AddressingTypeAbsoluteX,
		bytes:          3,
		cycle:          7,
	},
	0x60: {
		opecodeType:    OpecodeRTS,
		addressingType: AddressingTypeImplied,
		bytes:          1,
		cycle:          6,
	},
	0x61: {
		opecodeType:    OpecodeADC,
		addressingType: AddressingTypeIndirectX,
		bytes:          2,
		cycle:          6,
	},
	0x65: {
		opecodeType:    OpecodeADC,
		addressingType: AddressingTypeZeroPage,
		bytes:          2,
		cycle:          3,
	},
	0x66: {
		opecodeType:    OpecodeROR,
		addressingType: AddressingTypeZeroPage,
		bytes:          2,
		cycle:          5,
	},
	0x68: {
		opecodeType:    OpecodePLA,
		addressingType: AddressingTypeImplied,
		bytes:          1,
		cycle:          4,
	},
	0x69: {
		opecodeType:    OpecodeADC,
		addressingType: AddressingTypeImmediate,
		bytes:          2,
		cycle:          2,
	},
	0x6a: {
		opecodeType:    OpecodeROR,
		addressingType: AddressingTypeAccumulator,
		bytes:          1,
		cycle:          2,
	},
	0x6c: {
		opecodeType:    OpecodeJMP,
		addressingType: AddressingTypeIndirect,
		bytes:          3,
		cycle:          5,
	},
	0x6d: {
		opecodeType:    OpecodeADC,
		addressingType: AddressingTypeAbsolute,
		bytes:          3,
		cycle:          4,
	},
	0x6e: {
		opecodeType:    OpecodeROR,
		addressingType: AddressingTypeAbsolute,
		bytes:          3,
		cycle:          6,
	},
	0x70: {
		opecodeType:    OpecodeBVS,
		addressingType: AddressingTypeRelative,
		bytes:          2,
		cycle:          2,
	},
	0x71: {
		opecodeType:    OpecodeADC,
		addressingType: AddressingTypeAbsoluteY,
		bytes:          2,
		cycle:          5,
	},
	0x75: {
		opecodeType:    OpecodeADC,
		addressingType: AddressingTypeZeroPageX,
		bytes:          2,
		cycle:          4,
	},
	0x76: {
		opecodeType:    OpecodeROR,
		addressingType: AddressingTypeZeroPageX,
		bytes:          2,
		cycle:          6,
	},
	0x78: {
		opecodeType:    OpecodeSEI,
		addressingType: AddressingTypeImplied,
		bytes:          1,
		cycle:          2,
	},
	0x79: {
		opecodeType:    OpecodeADC,
		addressingType: AddressingTypeAbsoluteY,
		bytes:          3,
		cycle:          4,
	},
	0x7d: {
		opecodeType:    OpecodeADC,
		addressingType: AddressingTypeAbsoluteX,
		bytes:          3,
		cycle:          4,
	},
	0x7e: {
		opecodeType:    OpecodeROR,
		addressingType: AddressingTypeAbsoluteX,
		bytes:          3,
		cycle:          7,
	},
	0x81: {
		opecodeType:    OpecodeSTA,
		addressingType: AddressingTypeIndirectX,
		bytes:          2,
		cycle:          6,
	},
	0x84: {
		opecodeType:    OpecodeSTY,
		addressingType: AddressingTypeZeroPage,
		bytes:          2,
		cycle:          3,
	},
	0x85: {
		opecodeType:    OpecodeSTA,
		addressingType: AddressingTypeZeroPage,
		bytes:          2,
		cycle:          3,
	},
	0x86: {
		opecodeType:    OpecodeSTX,
		addressingType: AddressingTypeZeroPage,
		bytes:          2,
		cycle:          3,
	},
	0x88: {
		opecodeType:    OpecodeDEY,
		addressingType: AddressingTypeImplied,
		bytes:          1,
		cycle:          2,
	},
	0x8a: {
		opecodeType:    OpecodeTXA,
		addressingType: AddressingTypeImplied,
		bytes:          1,
		cycle:          2,
	},
	0x8c: {
		opecodeType:    OpecodeSTY,
		addressingType: AddressingTypeAbsolute,
		bytes:          3,
		cycle:          4,
	},
	0x8d: {
		opecodeType:    OpecodeSTA,
		addressingType: AddressingTypeAbsolute,
		bytes:          3,
		cycle:          4,
	},
	0x8e: {
		opecodeType:    OpecodeSTX,
		addressingType: AddressingTypeAbsolute,
		bytes:          3,
		cycle:          4,
	},
	0x90: {
		opecodeType:    OpecodeBCC,
		addressingType: AddressingTypeRelative,
		bytes:          2,
		cycle:          2,
	},
	0x91: {
		opecodeType:    OpecodeSTA,
		addressingType: AddressingTypeIndirectY,
		bytes:          2,
		cycle:          6,
	},
	0x94: {
		opecodeType:    OpecodeSTY,
		addressingType: AddressingTypeZeroPageX,
		bytes:          2,
		cycle:          4,
	},
	0x95: {
		opecodeType:    OpecodeSTA,
		addressingType: AddressingTypeZeroPageX,
		bytes:          2,
		cycle:          4,
	},
	0x96: {
		opecodeType:    OpecodeSTX,
		addressingType: AddressingTypeZeroPageY,
		bytes:          2,
		cycle:          4,
	},
	0x98: {
		opecodeType:    OpecodeTYA,
		addressingType: AddressingTypeImplied,
		bytes:          1,
		cycle:          2,
	},
	0x99: {
		opecodeType:    OpecodeSTA,
		addressingType: AddressingTypeAbsoluteY,
		bytes:          3,
		cycle:          5,
	},
	0x9a: {
		opecodeType:    OpecodeTXS,
		addressingType: AddressingTypeImplied,
		bytes:          1,
		cycle:          2,
	},
	0x9d: {
		opecodeType:    OpecodeSTA,
		addressingType: AddressingTypeAbsoluteX,
		bytes:          3,
		cycle:          5,
	},
	0xa0: {
		opecodeType:    OpecodeLDY,
		addressingType: AddressingTypeImmediate,
		bytes:          2,
		cycle:          2,
	},
	0xa1: {
		opecodeType:    OpecodeLDA,
		addressingType: AddressingTypeIndirectX,
		bytes:          2,
		cycle:          6,
	},
	0xa2: {
		opecodeType:    OpecodeLDX,
		addressingType: AddressingTypeImmediate,
		bytes:          2,
		cycle:          2,
	},
	0xa4: {
		opecodeType:    OpecodeLDY,
		addressingType: AddressingTypeZeroPage,
		bytes:          2,
		cycle:          3,
	},
	0xa5: {
		opecodeType:    OpecodeLDA,
		addressingType: AddressingTypeZeroPage,
		bytes:          2,
		cycle:          3,
	},
	0xa6: {
		opecodeType:    OpecodeLDX,
		addressingType: AddressingTypeZeroPage,
		bytes:          2,
		cycle:          3,
	},
	0xa8: {
		opecodeType:    OpecodeTAY,
		addressingType: AddressingTypeImplied,
		bytes:          1,
		cycle:          2,
	},
	0xa9: {
		opecodeType:    OpecodeLDA,
		addressingType: AddressingTypeImmediate,
		bytes:          2,
		cycle:          2,
	},
	0xaa: {
		opecodeType:    OpecodeTAX,
		addressingType: AddressingTypeImplied,
		bytes:          1,
		cycle:          2,
	},
	0xac: {
		opecodeType:    OpecodeLDY,
		addressingType: AddressingTypeAbsolute,
		bytes:          3,
		cycle:          4,
	},
	0xad: {
		opecodeType:    OpecodeLDA,
		addressingType: AddressingTypeAbsolute,
		bytes:          3,
		cycle:          4,
	},
	0xae: {
		opecodeType:    OpecodeLDX,
		addressingType: AddressingTypeAbsolute,
		bytes:          3,
		cycle:          4,
	},
	0xb0: {
		opecodeType:    OpecodeBCS,
		addressingType: AddressingTypeRelative,
		bytes:          2,
		cycle:          2,
	},
	0xb1: {
		opecodeType:    OpecodeLDA,
		addressingType: AddressingTypeIndirectY,
		bytes:          2,
		cycle:          5,
	},
	0xb4: {
		opecodeType:    OpecodeLDY,
		addressingType: AddressingTypeZeroPageX,
		bytes:          2,
		cycle:          4,
	},
	0xb5: {
		opecodeType:    OpecodeLDA,
		addressingType: AddressingTypeZeroPageX,
		bytes:          2,
		cycle:          4,
	},
	0xb6: {
		opecodeType:    OpecodeLDX,
		addressingType: AddressingTypeZeroPageY,
		bytes:          2,
		cycle:          4,
	},
	0xb8: {
		opecodeType:    OpecodeCLV,
		addressingType: AddressingTypeImplied,
		bytes:          1,
		cycle:          2,
	},
	0xb9: {
		opecodeType:    OpecodeLDA,
		addressingType: AddressingTypeAbsoluteY,
		bytes:          3,
		cycle:          4,
	},
	0xba: {
		opecodeType:    OpecodeTSX,
		addressingType: AddressingTypeImplied,
		bytes:          1,
		cycle:          2,
	},
	0xbc: {
		opecodeType:    OpecodeLDY,
		addressingType: AddressingTypeAbsolute,
		bytes:          3,
		cycle:          4,
	},
	0xbd: {
		opecodeType:    OpecodeLDA,
		addressingType: AddressingTypeAbsoluteX,
		bytes:          3,
		cycle:          4,
	},
	0xbe: {
		opecodeType:    OpecodeLDX,
		addressingType: AddressingTypeAbsoluteY,
		bytes:          3,
		cycle:          4,
	},
	0xc0: {
		opecodeType:    OpecodeCPY,
		addressingType: AddressingTypeImmediate,
		bytes:          2,
		cycle:          2,
	},
	0xc1: {
		opecodeType:    OpecodeCMP,
		addressingType: AddressingTypeIndirectX,
		bytes:          2,
		cycle:          6,
	},
	0xc4: {
		opecodeType:    OpecodeCPY,
		addressingType: AddressingTypeZeroPage,
		bytes:          2,
		cycle:          3,
	},
	0xc5: {
		opecodeType:    OpecodeCMP,
		addressingType: AddressingTypeZeroPage,
		bytes:          2,
		cycle:          3,
	},
	0xc6: {
		opecodeType:    OpecodeDEC,
		addressingType: AddressingTypeZeroPage,
		bytes:          2,
		cycle:          5,
	},
	0xc8: {
		opecodeType:    OpecodeINY,
		addressingType: AddressingTypeImplied,
		bytes:          1,
		cycle:          2,
	},
	0xc9: {
		opecodeType:    OpecodeCMP,
		addressingType: AddressingTypeImmediate,
		bytes:          2,
		cycle:          2,
	},
	0xca: {
		opecodeType:    OpecodeDEX,
		addressingType: AddressingTypeImplied,
		bytes:          1,
		cycle:          2,
	},
	0xcc: {
		opecodeType:    OpecodeCPY,
		addressingType: AddressingTypeAbsolute,
		bytes:          3,
		cycle:          4,
	},
	0xcd: {
		opecodeType:    OpecodeCMP,
		addressingType: AddressingTypeAbsolute,
		bytes:          3,
		cycle:          4,
	},
	0xce: {
		opecodeType:    OpecodeDEC,
		addressingType: AddressingTypeAbsolute,
		bytes:          3,
		cycle:          6,
	},
	0xd0: {
		opecodeType:    OpecodeBNE,
		addressingType: AddressingTypeRelative,
		bytes:          2,
		cycle:          2,
	},
	0xd1: {
		opecodeType:    OpecodeCMP,
		addressingType: AddressingTypeIndirectY,
		bytes:          2,
		cycle:          5,
	},
	0xd5: {
		opecodeType:    OpecodeCMP,
		addressingType: AddressingTypeZeroPageX,
		bytes:          2,
		cycle:          4,
	},
	0xd6: {
		opecodeType:    OpecodeDEC,
		addressingType: AddressingTypeZeroPageX,
		bytes:          2,
		cycle:          6,
	},
	0xd8: {
		opecodeType:    OpecodeCLD,
		addressingType: AddressingTypeImplied,
		bytes:          1,
		cycle:          2,
	},
	0xd9: {
		opecodeType:    OpecodeCMP,
		addressingType: AddressingTypeAbsoluteY,
		bytes:          3,
		cycle:          4,
	},
	0xde: {
		opecodeType:    OpecodeDEC,
		addressingType: AddressingTypeAbsoluteX,
		bytes:          3,
		cycle:          7,
	},
	0xe0: {
		opecodeType:    OpecodeCPX,
		addressingType: AddressingTypeImmediate,
		bytes:          2,
		cycle:          2,
	},
	0xe1: {
		opecodeType:    OpecodeSBC,
		addressingType: AddressingTypeIndirectX,
		bytes:          2,
		cycle:          6,
	},
	0xe4: {
		opecodeType:    OpecodeCPX,
		addressingType: AddressingTypeZeroPage,
		bytes:          2,
		cycle:          3,
	},
	0xe5: {
		opecodeType:    OpecodeSBC,
		addressingType: AddressingTypeZeroPage,
		bytes:          2,
		cycle:          3,
	},
	0xe6: {
		opecodeType:    OpecodeINC,
		addressingType: AddressingTypeZeroPage,
		bytes:          2,
		cycle:          5,
	},
	0xe8: {
		opecodeType:    OpecodeINX,
		addressingType: AddressingTypeImplied,
		bytes:          1,
		cycle:          2,
	},
	0xe9: {
		opecodeType:    OpecodeSBC,
		addressingType: AddressingTypeImmediate,
		bytes:          2,
		cycle:          2,
	},
	0xea: {
		opecodeType:    OpecodeNOP,
		addressingType: AddressingTypeImplied,
		bytes:          1,
		cycle:          2,
	},
	0xec: {
		opecodeType:    OpecodeCPX,
		addressingType: AddressingTypeAbsolute,
		bytes:          3,
		cycle:          4,
	},
	0xed: {
		opecodeType:    OpecodeSBC,
		addressingType: AddressingTypeAbsolute,
		bytes:          3,
		cycle:          4,
	},
	0xee: {
		opecodeType:    OpecodeINC,
		addressingType: AddressingTypeAbsolute,
		bytes:          3,
		cycle:          6,
	},
	0xf0: {
		opecodeType:    OpecodeBEQ,
		addressingType: AddressingTypeRelative,
		bytes:          2,
		cycle:          2,
	},
	0xf1: {
		opecodeType:    OpecodeSBC,
		addressingType: AddressingTypeIndirectY,
		bytes:          2,
		cycle:          5,
	},
	0xf5: {
		opecodeType:    OpecodeSBC,
		addressingType: AddressingTypeZeroPageX,
		bytes:          2,
		cycle:          4,
	},
	0xf6: {
		opecodeType:    OpecodeINC,
		addressingType: AddressingTypeZeroPageX,
		bytes:          2,
		cycle:          6,
	},
	0xf8: {
		opecodeType:    OpecodeSED,
		addressingType: AddressingTypeImplied,
		bytes:          1,
		cycle:          2,
	},
	0xf9: {
		opecodeType:    OpecodeSBC,
		addressingType: AddressingTypeAbsoluteY,
		bytes:          3,
		cycle:          4,
	},
	0xfd: {
		opecodeType:    OpecodeSBC,
		addressingType: AddressingTypeAbsoluteX,
		bytes:          3,
		cycle:          4,
	},
	0xfe: {
		opecodeType:    OpecodeINC,
		addressingType: AddressingTypeAbsoluteX,
		bytes:          3,
		cycle:          7,
	},
}
