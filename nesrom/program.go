package nesrom

// OpcodeType ...
type OpcodeType string

const (
	// official opcodes below
	OpcodeLDA = "lda"
	OpcodeLDX = "ldx"
	OpcodeLDY = "ldy"
	OpcodeSTA = "sta"
	OpcodeSTX = "stx"
	OpcodeSTY = "sty"
	OpcodeTXA = "txa"
	OpcodeTYA = "tya"
	OpcodeTXS = "txs"
	OpcodeTAY = "tay"
	OpcodeTAX = "tax"
	OpcodeTSX = "tsx"
	OpcodePHA = "pha"
	OpcodePHP = "php"
	OpcodePLA = "pla"
	OpcodePLP = "plp"
	OpcodeADC = "adc"
	OpcodeSBC = "sbc"
	OpcodeCPX = "cpx"
	OpcodeCPY = "cpy"
	OpcodeCMP = "cmp"
	OpcodeAND = "and"
	OpcodeEOR = "eor"
	OpcodeORA = "ora"
	OpcodeBIT = "bit"
	OpcodeASL = "asl"
	OpcodeLSR = "lsr"
	OpcodeROL = "rol"
	OpcodeROR = "ror"
	OpcodeINX = "inx"
	OpcodeINY = "iny"
	OpcodeINC = "inc"
	OpcodeDEX = "dex"
	OpcodeDEY = "dey"
	OpcodeDEC = "dec"
	OpcodeCLC = "clc"
	OpcodeCLI = "cli"
	OpcodeCLV = "clv"
	OpcodeCLD = "cld"
	OpcodeSEC = "sec"
	OpcodeSEI = "sei"
	OpcodeSED = "sed"
	OpcodeJSR = "jsr"
	OpcodeJMP = "jmp"
	OpcodeRTI = "rti"
	OpcodeRTS = "rts"
	OpcodeBCC = "bcc"
	OpcodeBCS = "bcs"
	OpcodeBEQ = "beq"
	OpcodeBMI = "bmi"
	OpcodeBNE = "bne"
	OpcodeBPL = "bpl"
	OpcodeBVC = "bvc"
	OpcodeBVS = "bvs"
	OpcodeNOP = "nop"
	OpcodeBRK = "brk"
	// unofficial opcodes below
	// OpcodeLAX = "lax"
	// OpcodeASO = "aso"
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
	opcodeType     OpcodeType
	addressingType AddressingType
	bytes          int
	cycle          int
}

// InstructionMap ...
var InstructionMap = map[int]Instruction{
	0x00: Instruction{
		opcodeType:     OpcodeBRK,
		addressingType: AddressingTypeImplied,
		bytes:          1,
		cycle:          7,
	},
	0x01: Instruction{
		opcodeType:     OpcodeORA,
		addressingType: AddressingTypeIndirectX,
		bytes:          2,
		cycle:          6,
	},
	0x05: Instruction{
		opcodeType:     OpcodeORA,
		addressingType: AddressingTypeZeroPage,
		bytes:          2,
		cycle:          3,
	},
	0x06: Instruction{
		opcodeType:     OpcodeASL,
		addressingType: AddressingTypeZeroPage,
		bytes:          2,
		cycle:          5,
	},
	0x08: Instruction{
		opcodeType:     OpcodePHP,
		addressingType: AddressingTypeImplied,
		bytes:          1,
		cycle:          3,
	},
	0x09: Instruction{
		opcodeType:     OpcodeORA,
		addressingType: AddressingTypeImmediate,
		bytes:          2,
		cycle:          2,
	},
	0x0A: Instruction{
		opcodeType:     OpcodeASL,
		addressingType: AddressingTypeAccumulator,
		bytes:          1,
		cycle:          2,
	},
	0x0D: Instruction{
		opcodeType:     OpcodeORA,
		addressingType: AddressingTypeAbsolute,
		bytes:          3,
		cycle:          4,
	},
	0x0E: Instruction{
		opcodeType:     OpcodeASL,
		addressingType: AddressingTypeAbsolute,
		bytes:          3,
		cycle:          6,
	},
	0x10: Instruction{
		opcodeType:     OpcodeBPL,
		addressingType: AddressingTypeRelative,
		bytes:          2,
		cycle:          2,
	},
	0x11: Instruction{
		opcodeType:     OpcodeORA,
		addressingType: AddressingTypeIndirectY,
		bytes:          2,
		cycle:          5,
	},
	0x15: Instruction{
		opcodeType:     OpcodeORA,
		addressingType: AddressingTypeZeroPageX,
		bytes:          2,
		cycle:          4,
	},
	0x16: Instruction{
		opcodeType:     OpcodeASL,
		addressingType: AddressingTypeZeroPageX,
		bytes:          2,
		cycle:          6,
	},
	0x18: Instruction{
		opcodeType:     OpcodeCLC,
		addressingType: AddressingTypeImplied,
		bytes:          1,
		cycle:          2,
	},
	0x19: Instruction{
		opcodeType:     OpcodeORA,
		addressingType: AddressingTypeAbsoluteY,
		bytes:          3,
		cycle:          4,
	},
	0x1D: Instruction{
		opcodeType:     OpcodeORA,
		addressingType: AddressingTypeAbsoluteX,
		bytes:          3,
		cycle:          4,
	},
	0x1E: Instruction{
		opcodeType:     OpcodeASL,
		addressingType: AddressingTypeAbsoluteX,
		bytes:          3,
		cycle:          7,
	},
	0x20: Instruction{
		opcodeType:     OpcodeJSR,
		addressingType: AddressingTypeAbsolute,
		bytes:          3,
		cycle:          6,
	},
	0x21: Instruction{
		opcodeType:     OpcodeAND,
		addressingType: AddressingTypeIndirectX,
		bytes:          2,
		cycle:          6,
	},
	0x24: Instruction{
		opcodeType:     OpcodeBIT,
		addressingType: AddressingTypeZeroPage,
		bytes:          2,
		cycle:          3,
	},
	0x25: Instruction{
		opcodeType:     OpcodeAND,
		addressingType: AddressingTypeZeroPage,
		bytes:          2,
		cycle:          3,
	},
	0x26: Instruction{
		opcodeType:     OpcodeROL,
		addressingType: AddressingTypeZeroPage,
		bytes:          2,
		cycle:          5,
	},
	0x28: Instruction{
		opcodeType:     OpcodePLP,
		addressingType: AddressingTypeImplied,
		bytes:          1,
		cycle:          4,
	},
	0x29: Instruction{
		opcodeType:     OpcodeAND,
		addressingType: AddressingTypeImmediate,
		bytes:          2,
		cycle:          2,
	},
	0x2A: Instruction{
		opcodeType:     OpcodeROL,
		addressingType: AddressingTypeAccumulator,
		bytes:          1,
		cycle:          2,
	},
	0x2C: Instruction{
		opcodeType:     OpcodeBIT,
		addressingType: AddressingTypeAbsolute,
		bytes:          3,
		cycle:          4,
	},
	0x2D: Instruction{
		opcodeType:     OpcodeAND,
		addressingType: AddressingTypeAbsolute,
		bytes:          3,
		cycle:          4,
	},
	0x2E: Instruction{
		opcodeType:     OpcodeROL,
		addressingType: AddressingTypeAbsolute,
		bytes:          3,
		cycle:          6,
	},
	0x30: Instruction{
		opcodeType:     OpcodeBMI,
		addressingType: AddressingTypeRelative,
		bytes:          2,
		cycle:          2,
	},
	0x31: Instruction{
		opcodeType:     OpcodeAND,
		addressingType: AddressingTypeIndirectY,
		bytes:          2,
		cycle:          5,
	},
	0x35: Instruction{
		opcodeType:     OpcodeAND,
		addressingType: AddressingTypeZeroPageX,
		bytes:          2,
		cycle:          4,
	},
	0x36: Instruction{
		opcodeType:     OpcodeROL,
		addressingType: AddressingTypeZeroPageX,
		bytes:          2,
		cycle:          6,
	},
	0x38: Instruction{
		opcodeType:     OpcodeSEC,
		addressingType: AddressingTypeImplied,
		bytes:          1,
		cycle:          2,
	},
	0x39: Instruction{
		opcodeType:     OpcodeAND,
		addressingType: AddressingTypeAbsoluteY,
		bytes:          3,
		cycle:          4,
	},
	0x3D: Instruction{
		opcodeType:     OpcodeAND,
		addressingType: AddressingTypeAbsoluteX,
		bytes:          3,
		cycle:          4,
	},
	0x3E: Instruction{
		opcodeType:     OpcodeROL,
		addressingType: AddressingTypeAbsoluteX,
		bytes:          3,
		cycle:          7,
	},
	0x40: Instruction{
		opcodeType:     OpcodeRTI,
		addressingType: AddressingTypeImplied,
		bytes:          1,
		cycle:          6,
	},
	0x41: Instruction{
		opcodeType:     OpcodeEOR,
		addressingType: AddressingTypeIndirectX,
		bytes:          2,
		cycle:          6,
	},
	// 0x44: Instruction{
	// 	opcodeType:     OpcodeNOP,
	// 	addressingType: AddressingTypeImplied,
	// 	bytes:          1,
	// 	cycle:          2,
	// },
	0x45: Instruction{
		opcodeType:     OpcodeEOR,
		addressingType: AddressingTypeZeroPage,
		bytes:          1,
		cycle:          3,
	},
	0x46: Instruction{
		opcodeType:     OpcodeLSR,
		addressingType: AddressingTypeZeroPage,
		bytes:          2,
		cycle:          5,
	},
	0x48: Instruction{
		opcodeType:     OpcodePHA,
		addressingType: AddressingTypeImplied,
		bytes:          1,
		cycle:          3,
	},
	0x49: Instruction{
		opcodeType:     OpcodeEOR,
		addressingType: AddressingTypeImmediate,
		bytes:          2,
		cycle:          3,
	},
	0x4A: Instruction{
		opcodeType:     OpcodeLSR,
		addressingType: AddressingTypeAccumulator,
		bytes:          1,
		cycle:          2,
	},
	0x4C: Instruction{
		opcodeType:     OpcodeJMP,
		addressingType: AddressingTypeAbsolute,
		bytes:          3,
		cycle:          3,
	},
	0x4D: Instruction{
		opcodeType:     OpcodeEOR,
		addressingType: AddressingTypeAbsolute,
		bytes:          3,
		cycle:          4,
	},
	0x4E: Instruction{
		opcodeType:     OpcodeLSR,
		addressingType: AddressingTypeAbsolute,
		bytes:          3,
		cycle:          6,
	},
	0x50: Instruction{
		opcodeType:     OpcodeBVC,
		addressingType: AddressingTypeRelative,
		bytes:          2,
		cycle:          2,
	},
	0x51: Instruction{
		opcodeType:     OpcodeEOR,
		addressingType: AddressingTypeIndirectY,
		bytes:          2,
		cycle:          5,
	},
	// 0x54: Instruction{
	// 	opcodeType:     OpcodeNOP,
	// 	addressingType: AddressingTypeImplied,
	// 	bytes:          1,
	// 	cycle:          2,
	// },
	0x55: Instruction{
		opcodeType:     OpcodeEOR,
		addressingType: AddressingTypeZeroPageX,
		bytes:          2,
		cycle:          4,
	},
	0x56: Instruction{
		opcodeType:     OpcodeLSR,
		addressingType: AddressingTypeZeroPageX,
		bytes:          2,
		cycle:          6,
	},
	0x58: Instruction{
		opcodeType:     OpcodeCLI,
		addressingType: AddressingTypeImplied,
		bytes:          1,
		cycle:          2,
	},
	0x59: Instruction{
		opcodeType:     OpcodeEOR,
		addressingType: AddressingTypeAbsoluteY,
		bytes:          3,
		cycle:          4,
	},
	0x5D: Instruction{
		opcodeType:     OpcodeEOR,
		addressingType: AddressingTypeAbsoluteX,
		bytes:          3,
		cycle:          4,
	},
	0x5E: Instruction{
		opcodeType:     OpcodeLSR,
		addressingType: AddressingTypeAbsoluteX,
		bytes:          3,
		cycle:          7,
	},
	0x60: Instruction{
		opcodeType:     OpcodeRTS,
		addressingType: AddressingTypeImplied,
		bytes:          1,
		cycle:          6,
	},
	0x61: Instruction{
		opcodeType:     OpcodeADC,
		addressingType: AddressingTypeIndirectX,
		bytes:          2,
		cycle:          6,
	},
	0x65: Instruction{
		opcodeType:     OpcodeADC,
		addressingType: AddressingTypeZeroPage,
		bytes:          2,
		cycle:          3,
	},
	0x66: Instruction{
		opcodeType:     OpcodeROR,
		addressingType: AddressingTypeZeroPage,
		bytes:          2,
		cycle:          5,
	},
	0x68: Instruction{
		opcodeType:     OpcodePLA,
		addressingType: AddressingTypeImplied,
		bytes:          1,
		cycle:          4,
	},
	0x69: Instruction{
		opcodeType:     OpcodeADC,
		addressingType: AddressingTypeImmediate,
		bytes:          2,
		cycle:          2,
	},
	0x6A: Instruction{
		opcodeType:     OpcodeROR,
		addressingType: AddressingTypeAccumulator,
		bytes:          1,
		cycle:          2,
	},
	0x6C: Instruction{
		opcodeType:     OpcodeJMP,
		addressingType: AddressingTypeIndirect,
		bytes:          3,
		cycle:          5,
	},
	0x6D: Instruction{
		opcodeType:     OpcodeADC,
		addressingType: AddressingTypeAbsolute,
		bytes:          3,
		cycle:          4,
	},
	0x6E: Instruction{
		opcodeType:     OpcodeROR,
		addressingType: AddressingTypeAbsolute,
		bytes:          3,
		cycle:          6,
	},
	0x70: Instruction{
		opcodeType:     OpcodeBVS,
		addressingType: AddressingTypeRelative,
		bytes:          2,
		cycle:          2,
	},
	0x71: Instruction{
		opcodeType:     OpcodeADC,
		addressingType: AddressingTypeAbsoluteY,
		bytes:          2,
		cycle:          5,
	},
	0x75: Instruction{
		opcodeType:     OpcodeADC,
		addressingType: AddressingTypeZeroPageX,
		bytes:          2,
		cycle:          4,
	},
	0x76: Instruction{
		opcodeType:     OpcodeROR,
		addressingType: AddressingTypeZeroPageX,
		bytes:          2,
		cycle:          6,
	},
	0x78: Instruction{
		opcodeType:     OpcodeSEI,
		addressingType: AddressingTypeImplied,
		bytes:          1,
		cycle:          2,
	},
	0x79: Instruction{
		opcodeType:     OpcodeADC,
		addressingType: AddressingTypeAbsoluteY,
		bytes:          3,
		cycle:          4,
	},
	// 0x7C: Instruction{
	// 	opcodeType:     OpcodeNOP,
	// 	addressingType: AddressingTypeImplied,
	// 	bytes:          1,
	// 	cycle:          2,
	// },
	0x7D: Instruction{
		opcodeType:     OpcodeADC,
		addressingType: AddressingTypeAbsoluteX,
		bytes:          3,
		cycle:          4,
	},
	0x7E: Instruction{
		opcodeType:     OpcodeROR,
		addressingType: AddressingTypeAbsoluteX,
		bytes:          3,
		cycle:          7,
	},
	0x81: Instruction{
		opcodeType:     OpcodeSTA,
		addressingType: AddressingTypeIndirectX,
		bytes:          2,
		cycle:          6,
	},
	0x84: Instruction{
		opcodeType:     OpcodeSTY,
		addressingType: AddressingTypeZeroPage,
		bytes:          2,
		cycle:          3,
	},
	0x85: Instruction{
		opcodeType:     OpcodeSTA,
		addressingType: AddressingTypeZeroPage,
		bytes:          2,
		cycle:          3,
	},
	0x86: Instruction{
		opcodeType:     OpcodeSTX,
		addressingType: AddressingTypeZeroPage,
		bytes:          2,
		cycle:          3,
	},
	// 0x87: Instruction{
	// 	opcodeType:     OpcodeLAX,
	// 	addressingType: AddressingTypeZeroPageY,
	// 	bytes:          3,
	// 	cycle:          4,
	// },
	0x88: Instruction{
		opcodeType:     OpcodeDEY,
		addressingType: AddressingTypeImplied,
		bytes:          1,
		cycle:          2,
	},
	0x8A: Instruction{
		opcodeType:     OpcodeTXA,
		addressingType: AddressingTypeImplied,
		bytes:          1,
		cycle:          2,
	},
	0x8C: Instruction{
		opcodeType:     OpcodeSTY,
		addressingType: AddressingTypeAbsolute,
		bytes:          3,
		cycle:          4,
	},
	0x8D: Instruction{
		opcodeType:     OpcodeSTA,
		addressingType: AddressingTypeAbsolute,
		bytes:          3,
		cycle:          4,
	},
	0x8E: Instruction{
		opcodeType:     OpcodeSTX,
		addressingType: AddressingTypeAbsolute,
		bytes:          3,
		cycle:          4,
	},
	0x90: Instruction{
		opcodeType:     OpcodeBCC,
		addressingType: AddressingTypeRelative,
		bytes:          2,
		cycle:          2,
	},
	0x91: Instruction{
		opcodeType:     OpcodeSTA,
		addressingType: AddressingTypeIndirectY,
		bytes:          2,
		cycle:          6,
	},
	0x94: Instruction{
		opcodeType:     OpcodeSTY,
		addressingType: AddressingTypeZeroPageX,
		bytes:          2,
		cycle:          4,
	},
	0x95: Instruction{
		opcodeType:     OpcodeSTA,
		addressingType: AddressingTypeZeroPageX,
		bytes:          2,
		cycle:          4,
	},
	0x96: Instruction{
		opcodeType:     OpcodeSTX,
		addressingType: AddressingTypeZeroPageY,
		bytes:          2,
		cycle:          4,
	},
	0x98: Instruction{
		opcodeType:     OpcodeTYA,
		addressingType: AddressingTypeImplied,
		bytes:          1,
		cycle:          2,
	},
	0x99: Instruction{
		opcodeType:     OpcodeSTA,
		addressingType: AddressingTypeAbsoluteY,
		bytes:          3,
		cycle:          5,
	},
	0x9A: Instruction{
		opcodeType:     OpcodeTXS,
		addressingType: AddressingTypeImplied,
		bytes:          1,
		cycle:          2,
	},
	0x9D: Instruction{
		opcodeType:     OpcodeSTA,
		addressingType: AddressingTypeAbsoluteX,
		bytes:          3,
		cycle:          5,
	},
	0xA0: Instruction{
		opcodeType:     OpcodeLDY,
		addressingType: AddressingTypeImmediate,
		bytes:          2,
		cycle:          2,
	},
	0xA1: Instruction{
		opcodeType:     OpcodeLDA,
		addressingType: AddressingTypeIndirectX,
		bytes:          2,
		cycle:          6,
	},
	0xA2: Instruction{
		opcodeType:     OpcodeLDX,
		addressingType: AddressingTypeImmediate,
		bytes:          2,
		cycle:          2,
	},
	// 0xA3: Instruction{
	// 	opcodeType:     OpcodeLAX,
	// 	addressingType: AddressingTypeIndirectX,
	// 	bytes:          2,
	// 	cycle:          6,
	// },
	0xA4: Instruction{
		opcodeType:     OpcodeLDY,
		addressingType: AddressingTypeZeroPage,
		bytes:          2,
		cycle:          3,
	},
	0xA5: Instruction{
		opcodeType:     OpcodeLDA,
		addressingType: AddressingTypeZeroPage,
		bytes:          2,
		cycle:          3,
	},
	0xA6: Instruction{
		opcodeType:     OpcodeLDX,
		addressingType: AddressingTypeZeroPage,
		bytes:          2,
		cycle:          3,
	},
	// 0xA7: Instruction{
	// 	opcodeType:     OpcodeLAX,
	// 	addressingType: AddressingTypeZeroPage,
	// 	bytes:          2,
	// 	cycle:          3,
	// },
	0xA8: Instruction{
		opcodeType:     OpcodeTAY,
		addressingType: AddressingTypeImplied,
		bytes:          1,
		cycle:          2,
	},
	0xA9: Instruction{
		opcodeType:     OpcodeLDA,
		addressingType: AddressingTypeImmediate,
		bytes:          2,
		cycle:          2,
	},
	0xAA: Instruction{
		opcodeType:     OpcodeTAX,
		addressingType: AddressingTypeImplied,
		bytes:          1,
		cycle:          2,
	},
	0xAC: Instruction{
		opcodeType:     OpcodeLDY,
		addressingType: AddressingTypeAbsolute,
		bytes:          3,
		cycle:          4,
	},
	0xAD: Instruction{
		opcodeType:     OpcodeLDA,
		addressingType: AddressingTypeAbsolute,
		bytes:          3,
		cycle:          4,
	},
	0xAE: Instruction{
		opcodeType:     OpcodeLDX,
		addressingType: AddressingTypeAbsolute,
		bytes:          3,
		cycle:          4,
	},
	// 0xAF: Instruction{
	// 	opcodeType:     OpcodeLAX,
	// 	addressingType: AddressingTypeAbsolute,
	// 	bytes:          3,
	// 	cycle:          4,
	// },
	0xB0: Instruction{
		opcodeType:     OpcodeBCS,
		addressingType: AddressingTypeRelative,
		bytes:          2,
		cycle:          2,
	},
	0xB1: Instruction{
		opcodeType:     OpcodeLDA,
		addressingType: AddressingTypeIndirectY,
		bytes:          2,
		cycle:          5,
	},
	// 0xB3: Instruction{
	// 	opcodeType:     OpcodeLAX,
	// 	addressingType: AddressingTypeIndirectY,
	// 	bytes:          2,
	// 	cycle:          5,
	// },
	0xB4: Instruction{
		opcodeType:     OpcodeLDY,
		addressingType: AddressingTypeZeroPageX,
		bytes:          2,
		cycle:          4,
	},
	0xB5: Instruction{
		opcodeType:     OpcodeLDA,
		addressingType: AddressingTypeZeroPageX,
		bytes:          2,
		cycle:          4,
	},
	0xB6: Instruction{
		opcodeType:     OpcodeLDX,
		addressingType: AddressingTypeZeroPageY,
		bytes:          2,
		cycle:          4,
	},
	0xB8: Instruction{
		opcodeType:     OpcodeCLV,
		addressingType: AddressingTypeImplied,
		bytes:          1,
		cycle:          2,
	},
	0xB9: Instruction{
		opcodeType:     OpcodeLDA,
		addressingType: AddressingTypeAbsoluteY,
		bytes:          3,
		cycle:          4,
	},
	0xBA: Instruction{
		opcodeType:     OpcodeTSX,
		addressingType: AddressingTypeImplied,
		bytes:          1,
		cycle:          2,
	},
	0xBC: Instruction{
		opcodeType:     OpcodeLDY,
		addressingType: AddressingTypeAbsolute,
		bytes:          3,
		cycle:          4,
	},
	0xBD: Instruction{
		opcodeType:     OpcodeLDA,
		addressingType: AddressingTypeAbsoluteX,
		bytes:          3,
		cycle:          4,
	},
	0xBE: Instruction{
		opcodeType:     OpcodeLDX,
		addressingType: AddressingTypeAbsoluteY,
		bytes:          3,
		cycle:          4,
	},
	// 0xBF: Instruction{
	// 	opcodeType:     OpcodeLAX,
	// 	addressingType: AddressingTypeAbsoluteY,
	// 	bytes:          3,
	// 	cycle:          4,
	// },
	0xC0: Instruction{
		opcodeType:     OpcodeCPY,
		addressingType: AddressingTypeImmediate,
		bytes:          2,
		cycle:          2,
	},
	0xC1: Instruction{
		opcodeType:     OpcodeCMP,
		addressingType: AddressingTypeIndirectX,
		bytes:          2,
		cycle:          6,
	},
	0xC4: Instruction{
		opcodeType:     OpcodeCPY,
		addressingType: AddressingTypeZeroPage,
		bytes:          2,
		cycle:          3,
	},
	0xC5: Instruction{
		opcodeType:     OpcodeCMP,
		addressingType: AddressingTypeZeroPage,
		bytes:          2,
		cycle:          3,
	},
	0xC6: Instruction{
		opcodeType:     OpcodeDEC,
		addressingType: AddressingTypeZeroPage,
		bytes:          2,
		cycle:          5,
	},
	0xC8: Instruction{
		opcodeType:     OpcodeINY,
		addressingType: AddressingTypeImplied,
		bytes:          1,
		cycle:          2,
	},
	0xC9: Instruction{
		opcodeType:     OpcodeCMP,
		addressingType: AddressingTypeImmediate,
		bytes:          2,
		cycle:          2,
	},
	0xCA: Instruction{
		opcodeType:     OpcodeDEX,
		addressingType: AddressingTypeImplied,
		bytes:          1,
		cycle:          2,
	},
	0xCC: Instruction{
		opcodeType:     OpcodeCPY,
		addressingType: AddressingTypeAbsolute,
		bytes:          3,
		cycle:          4,
	},
	0xCD: Instruction{
		opcodeType:     OpcodeCMP,
		addressingType: AddressingTypeAbsolute,
		bytes:          3,
		cycle:          4,
	},
	0xCE: Instruction{
		opcodeType:     OpcodeDEC,
		addressingType: AddressingTypeAbsolute,
		bytes:          3,
		cycle:          6,
	},
	0xD0: Instruction{
		opcodeType:     OpcodeBNE,
		addressingType: AddressingTypeRelative,
		bytes:          2,
		cycle:          2,
	},
	0xD1: Instruction{
		opcodeType:     OpcodeCMP,
		addressingType: AddressingTypeIndirectY,
		bytes:          2,
		cycle:          5,
	},
	0xD5: Instruction{
		opcodeType:     OpcodeCMP,
		addressingType: AddressingTypeZeroPageX,
		bytes:          2,
		cycle:          4,
	},
	0xD6: Instruction{
		opcodeType:     OpcodeDEC,
		addressingType: AddressingTypeZeroPageX,
		bytes:          2,
		cycle:          6,
	},
	0xD8: Instruction{
		opcodeType:     OpcodeCLD,
		addressingType: AddressingTypeImplied,
		bytes:          1,
		cycle:          2,
	},
	0xD9: Instruction{
		opcodeType:     OpcodeCMP,
		addressingType: AddressingTypeAbsoluteY,
		bytes:          3,
		cycle:          4,
	},
	0xDE: Instruction{
		opcodeType:     OpcodeDEC,
		addressingType: AddressingTypeAbsoluteX,
		bytes:          3,
		cycle:          7,
	},
	0xE0: Instruction{
		opcodeType:     OpcodeCPX,
		addressingType: AddressingTypeImmediate,
		bytes:          2,
		cycle:          2,
	},
	0xE1: Instruction{
		opcodeType:     OpcodeSBC,
		addressingType: AddressingTypeIndirectX,
		bytes:          2,
		cycle:          6,
	},
	0xE4: Instruction{
		opcodeType:     OpcodeCPX,
		addressingType: AddressingTypeZeroPage,
		bytes:          2,
		cycle:          3,
	},
	0xE5: Instruction{
		opcodeType:     OpcodeSBC,
		addressingType: AddressingTypeZeroPage,
		bytes:          2,
		cycle:          3,
	},
	0xE6: Instruction{
		opcodeType:     OpcodeINC,
		addressingType: AddressingTypeZeroPage,
		bytes:          2,
		cycle:          5,
	},
	0xE8: Instruction{
		opcodeType:     OpcodeINX,
		addressingType: AddressingTypeImplied,
		bytes:          1,
		cycle:          2,
	},
	0xE9: Instruction{
		opcodeType:     OpcodeSBC,
		addressingType: AddressingTypeImmediate,
		bytes:          2,
		cycle:          2,
	},
	0xEA: Instruction{
		opcodeType:     OpcodeNOP,
		addressingType: AddressingTypeImplied,
		bytes:          1,
		cycle:          2,
	},
	0xEC: Instruction{
		opcodeType:     OpcodeCPX,
		addressingType: AddressingTypeAbsolute,
		bytes:          3,
		cycle:          4,
	},
	0xED: Instruction{
		opcodeType:     OpcodeSBC,
		addressingType: AddressingTypeAbsolute,
		bytes:          3,
		cycle:          4,
	},
	0xEE: Instruction{
		opcodeType:     OpcodeINC,
		addressingType: AddressingTypeAbsolute,
		bytes:          3,
		cycle:          6,
	},
	0xF0: Instruction{
		opcodeType:     OpcodeBEQ,
		addressingType: AddressingTypeRelative,
		bytes:          2,
		cycle:          2,
	},
	0xF1: Instruction{
		opcodeType:     OpcodeSBC,
		addressingType: AddressingTypeIndirectY,
		bytes:          2,
		cycle:          5,
	},
	0xF5: Instruction{
		opcodeType:     OpcodeSBC,
		addressingType: AddressingTypeZeroPageX,
		bytes:          2,
		cycle:          4,
	},
	0xF6: Instruction{
		opcodeType:     OpcodeINC,
		addressingType: AddressingTypeZeroPageX,
		bytes:          2,
		cycle:          6,
	},
	0xF8: Instruction{
		opcodeType:     OpcodeSED,
		addressingType: AddressingTypeImplied,
		bytes:          1,
		cycle:          2,
	},
	0xF9: Instruction{
		opcodeType:     OpcodeSBC,
		addressingType: AddressingTypeAbsoluteY,
		bytes:          3,
		cycle:          4,
	},
	0xFD: Instruction{
		opcodeType:     OpcodeSBC,
		addressingType: AddressingTypeAbsoluteX,
		bytes:          3,
		cycle:          4,
	},
	0xFE: Instruction{
		opcodeType:     OpcodeINC,
		addressingType: AddressingTypeAbsoluteX,
		bytes:          3,
		cycle:          7,
	},
}
