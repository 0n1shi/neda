package neda

import "fmt"

func formatAddress(addr Address) string {
	return fmt.Sprintf("0x%04X: ", addr)
}

func formatByteCodes(codes []byte) string {
	str := ""
	for _, v := range codes {
		str += fmt.Sprintf("%02X ", v)
	}
	for i := 0; i < 4-len(codes); i++ {
		str += "   "
	}
	str += "  "
	return str
}

func formatOpecode(opcodeType OpecodeType) string {
	return fmt.Sprintf("%s ", OpecodeMap[opcodeType])
}

func formatOperand(currentAddr Address, addrType AddressingType, val int) string {
	switch addrType {
	case AddressingTypeImmediate:
		return fmt.Sprintf("#$%02X", val)
	case AddressingTypeAbsolute:
		return fmt.Sprintf("$%04X", val)
	case AddressingTypeZeroPage:
		return fmt.Sprintf("$%02X", val)
	case AddressingTypeImplied:
		return ""
	case AddressingTypeIndirect:
		return fmt.Sprintf("($%04X)", val)
	case AddressingTypeAbsoluteX:
		return fmt.Sprintf("$%04X, X", val)
	case AddressingTypeAbsoluteY:
		return fmt.Sprintf("$%04X, Y", val)
	case AddressingTypeZeroPageX:
		return fmt.Sprintf("$%02X, X", val)
	case AddressingTypeZeroPageY:
		return fmt.Sprintf("$%02X, Y", val)
	case AddressingTypeIndirectX:
		return fmt.Sprintf("($%02X, X)", val)
	case AddressingTypeIndirectY:
		return fmt.Sprintf("($%02X), Y", val)
	case AddressingTypeRelative:
		return fmt.Sprintf("$%04X", int(currentAddr+2)+int(int8(val)))
	case AddressingTypeAccumulator:
		return "A"
	}
	panic("invalid addressing type")
}
