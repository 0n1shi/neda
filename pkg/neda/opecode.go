package neda

type OpecodeType int

const (
	OpecodeLDA = OpecodeType(iota)
	OpecodeLDX
	OpecodeLDY
	OpecodeSTA
	OpecodeSTX
	OpecodeSTY
	OpecodeTXA
	OpecodeTYA
	OpecodeTXS
	OpecodeTAY
	OpecodeTAX
	OpecodeTSX
	OpecodePHA
	OpecodePHP
	OpecodePLA
	OpecodePLP
	OpecodeADC
	OpecodeSBC
	OpecodeCPX
	OpecodeCPY
	OpecodeCMP
	OpecodeAND
	OpecodeEOR
	OpecodeORA
	OpecodeBIT
	OpecodeASL
	OpecodeLSR
	OpecodeROL
	OpecodeROR
	OpecodeINX
	OpecodeINY
	OpecodeINC
	OpecodeDEX
	OpecodeDEY
	OpecodeDEC
	OpecodeCLC
	OpecodeCLI
	OpecodeCLV
	OpecodeCLD
	OpecodeSEC
	OpecodeSEI
	OpecodeSED
	OpecodeJSR
	OpecodeJMP
	OpecodeRTI
	OpecodeRTS
	OpecodeBCC
	OpecodeBCS
	OpecodeBEQ
	OpecodeBMI
	OpecodeBNE
	OpecodeBPL
	OpecodeBVC
	OpecodeBVS
	OpecodeNOP
	OpecodeBRK
)

var OpecodeMap = map[OpecodeType]string{
	OpecodeLDA: "lda",
	OpecodeLDX: "ldx",
	OpecodeLDY: "ldy",
	OpecodeSTA: "sta",
	OpecodeSTX: "stx",
	OpecodeSTY: "sty",
	OpecodeTXA: "txa",
	OpecodeTYA: "tya",
	OpecodeTXS: "txs",
	OpecodeTAY: "tay",
	OpecodeTAX: "tax",
	OpecodeTSX: "tsx",
	OpecodePHA: "pha",
	OpecodePHP: "php",
	OpecodePLA: "pla",
	OpecodePLP: "plp",
	OpecodeADC: "adc",
	OpecodeSBC: "sbc",
	OpecodeCPX: "cpx",
	OpecodeCPY: "cpy",
	OpecodeCMP: "cmp",
	OpecodeAND: "and",
	OpecodeEOR: "eor",
	OpecodeORA: "ora",
	OpecodeBIT: "bit",
	OpecodeASL: "asl",
	OpecodeLSR: "lsr",
	OpecodeROL: "rol",
	OpecodeROR: "ror",
	OpecodeINX: "inx",
	OpecodeINY: "iny",
	OpecodeINC: "inc",
	OpecodeDEX: "dex",
	OpecodeDEY: "dey",
	OpecodeDEC: "dec",
	OpecodeCLC: "clc",
	OpecodeCLI: "cli",
	OpecodeCLV: "clv",
	OpecodeCLD: "cld",
	OpecodeSEC: "sec",
	OpecodeSEI: "sei",
	OpecodeSED: "sed",
	OpecodeJSR: "jsr",
	OpecodeJMP: "jmp",
	OpecodeRTI: "rti",
	OpecodeRTS: "rts",
	OpecodeBCC: "bcc",
	OpecodeBCS: "bcs",
	OpecodeBEQ: "beq",
	OpecodeBMI: "bmi",
	OpecodeBNE: "bne",
	OpecodeBPL: "bpl",
	OpecodeBVC: "bvc",
	OpecodeBVS: "bvs",
	OpecodeNOP: "nop",
	OpecodeBRK: "brk",
}
