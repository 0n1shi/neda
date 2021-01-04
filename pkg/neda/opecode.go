package neda

type OpcodeType int

const (
	OpcodeLDA = OpcodeType(iota)
	OpcodeLDX
	OpcodeLDY
	OpcodeSTA
	OpcodeSTX
	OpcodeSTY
	OpcodeTXA
	OpcodeTYA
	OpcodeTXS
	OpcodeTAY
	OpcodeTAX
	OpcodeTSX
	OpcodePHA
	OpcodePHP
	OpcodePLA
	OpcodePLP
	OpcodeADC
	OpcodeSBC
	OpcodeCPX
	OpcodeCPY
	OpcodeCMP
	OpcodeAND
	OpcodeEOR
	OpcodeORA
	OpcodeBIT
	OpcodeASL
	OpcodeLSR
	OpcodeROL
	OpcodeROR
	OpcodeINX
	OpcodeINY
	OpcodeINC
	OpcodeDEX
	OpcodeDEY
	OpcodeDEC
	OpcodeCLC
	OpcodeCLI
	OpcodeCLV
	OpcodeCLD
	OpcodeSEC
	OpcodeSEI
	OpcodeSED
	OpcodeJSR
	OpcodeJMP
	OpcodeRTI
	OpcodeRTS
	OpcodeBCC
	OpcodeBCS
	OpcodeBEQ
	OpcodeBMI
	OpcodeBNE
	OpcodeBPL
	OpcodeBVC
	OpcodeBVS
	OpcodeNOP
	OpcodeBRK
)

var OpcodeMap = map[OpcodeType]string{
	OpcodeLDA: "lda",
	OpcodeLDX: "ldx",
	OpcodeLDY: "ldy",
	OpcodeSTA: "sta",
	OpcodeSTX: "stx",
	OpcodeSTY: "sty",
	OpcodeTXA: "txa",
	OpcodeTYA: "tya",
	OpcodeTXS: "txs",
	OpcodeTAY: "tay",
	OpcodeTAX: "tax",
	OpcodeTSX: "tsx",
	OpcodePHA: "pha",
	OpcodePHP: "php",
	OpcodePLA: "pla",
	OpcodePLP: "plp",
	OpcodeADC: "adc",
	OpcodeSBC: "sbc",
	OpcodeCPX: "cpx",
	OpcodeCPY: "cpy",
	OpcodeCMP: "cmp",
	OpcodeAND: "and",
	OpcodeEOR: "eor",
	OpcodeORA: "ora",
	OpcodeBIT: "bit",
	OpcodeASL: "asl",
	OpcodeLSR: "lsr",
	OpcodeROL: "rol",
	OpcodeROR: "ror",
	OpcodeINX: "inx",
	OpcodeINY: "iny",
	OpcodeINC: "inc",
	OpcodeDEX: "dex",
	OpcodeDEY: "dey",
	OpcodeDEC: "dec",
	OpcodeCLC: "clc",
	OpcodeCLI: "cli",
	OpcodeCLV: "clv",
	OpcodeCLD: "cld",
	OpcodeSEC: "sec",
	OpcodeSEI: "sei",
	OpcodeSED: "sed",
	OpcodeJSR: "jsr",
	OpcodeJMP: "jmp",
	OpcodeRTI: "rti",
	OpcodeRTS: "rts",
	OpcodeBCC: "bcc",
	OpcodeBCS: "bcs",
	OpcodeBEQ: "beq",
	OpcodeBMI: "bmi",
	OpcodeBNE: "bne",
	OpcodeBPL: "bpl",
	OpcodeBVC: "bvc",
	OpcodeBVS: "bvs",
	OpcodeNOP: "nop",
	OpcodeBRK: "brk",
}
