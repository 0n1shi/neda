package lib

import (
	"fmt"

	"github.com/k-onishi/nesgo"
)

// DumpHeader ...
func DumpHeader(h nesgo.Header) {
	fmt.Printf("magic number: %s\n", h.MagicNumber)
	fmt.Printf("program bank count: %d (%d bytes)\n", h.ProgramBankCount, h.GetProgramBankSize())
	fmt.Printf("character bank count: %d (%d bytes)\n", h.CharacterBankCount, h.GetCharacterBankSize())
	fmt.Printf("mapper type: %d (%s)\n", h.GetMapper(), h.GetMapperType())
	fmt.Println()
}

// PROMStartAddress ...
const PROMStartAddress = 0xC000

func analyzePBank(rom []byte) []*DecodeInfo {
	var decodeInfoList []*DecodeInfo
	for i := 0; i < len(rom); {
		info := DecodeInfo{}
		info.Address = i + PROMStartAddress
		info.Value = int(rom[i])
		info.Bytes = append(info.Bytes, rom[i]) // save byte code
		instruction, ok := nesgo.InstructionMap[info.Value]
		i++
		if !ok { // might be .db ?
			info.Instruction = nil
			info.Arg = nil
		} else {
			info.Instruction = &instruction

			// fetch argment
			argCount := instruction.Bytes - 1
			args := make([]byte, argCount)
			argValue := 0
			for j := 0; j < argCount; j++ {
				v := rom[i]
				args[j] = v
				info.Bytes = append(info.Bytes, v) // save byte code
				argValue += (int(v) << (8 * j))
				i++
			}
			info.Arg = &argValue
		}
		decodeInfoList = append(decodeInfoList, &info)
	}
	return decodeInfoList
}

func formatAddress(addr int) string {
	return fmt.Sprintf("0x%04X: ", addr)
}

func formatByteCodes(codes []byte) string {
	str := ""
	for _, v := range codes {
		str += fmt.Sprintf("%02X ", v)
	}
	for i := 0; i < 3-len(codes); i++ {
		str += "   "
	}
	str += "  "
	return str
}

func formatOpcode(opcodeType nesgo.OpcodeType) string {
	return fmt.Sprintf("%s ", nesgo.OpcodeMap[opcodeType])
}

func formatOperand(currentAddr int, addrType nesgo.AddressingType, val int) string {
	switch addrType {
	case nesgo.AddressingTypeImmediate:
		return fmt.Sprintf("#$%02X", val)
	case nesgo.AddressingTypeAbsolute:
		return fmt.Sprintf("$%04X", val)
	case nesgo.AddressingTypeZeroPage:
		return fmt.Sprintf("$%02X", val)
	case nesgo.AddressingTypeImplied:
		return ""
	case nesgo.AddressingTypeIndirect:
		return fmt.Sprintf("($%04X)", val)
	case nesgo.AddressingTypeAbsoluteX:
		return fmt.Sprintf("$%04X, X", val)
	case nesgo.AddressingTypeAbsoluteY:
		return fmt.Sprintf("$%04X, Y", val)
	case nesgo.AddressingTypeZeroPageX:
		return fmt.Sprintf("$%02X, X", val)
	case nesgo.AddressingTypeZeroPageY:
		return fmt.Sprintf("$%02X, Y", val)
	case nesgo.AddressingTypeIndirectX:
		return fmt.Sprintf("($%02X, X)", val)
	case nesgo.AddressingTypeIndirectY:
		return fmt.Sprintf("($%02X), Y", val)
	case nesgo.AddressingTypeRelative:
		return fmt.Sprintf("$%04X", currentAddr+val)
	case nesgo.AddressingTypeAccumulator:
		return "A"
	}
	panic("invalid addressing type")
}

// DumpPBank ...
func DumpPBank(rom []byte) {
	infoList := analyzePBank(rom)
	for _, info := range infoList {
		fmt.Printf("%s", formatAddress(info.Address))
		fmt.Printf("%s", formatByteCodes(info.Bytes))
		if info.Instruction != nil {
			fmt.Printf("%s", formatOpcode(info.Instruction.OpcodeType))
			fmt.Printf("%s", formatOperand(info.Address, info.Instruction.AddressingType, *info.Arg))
		}
		fmt.Println()
	}
}
