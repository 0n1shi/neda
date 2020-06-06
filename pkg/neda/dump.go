package neda

import (
	"fmt"
)

// DumpHeader ...
func DumpHeader(h Header) {
	fmt.Printf("magic number: %s\n", h.MagicNumber)
	fmt.Printf("program bank count: %d (%d bytes)\n", h.ProgramBankCount, h.GetProgramBankSize())
	fmt.Printf("character bank count: %d (%d bytes)\n", h.CharacterBankCount, h.GetCharacterBankSize())
	fmt.Printf("mapper type: %d (%s)\n", h.GetMapper(), h.GetMapperType())
	fmt.Println()
}

// PROMStartAddress ...
const PROMStartAddress = 0x8000

func isEndOfSubroutinue(opcode OpcodeType) bool {
	if opcode == OpcodeJMP || opcode == OpcodeRTS || opcode == OpcodeRTI || opcode == OpcodeJSR || opcode == OpcodeBRK {
		return true
	}
	return false
}

func analyzePBank(rom []byte) *AnalysisInfo {
	decodeInfoMap := make(map[Address]*DecodeInfo)
	var accessRangeList []*AccessRange
	accessRange := &AccessRange{}
	for i := 0; i < len(rom); {
		info := DecodeInfo{}
		info.Address = Address(i + PROMStartAddress)
		info.Value = int(rom[i])
		info.Bytes = append(info.Bytes, rom[i]) // save byte code
		instruction, ok := InstructionMap[info.Value]
		i++
		if !ok { // might be .db ?
			info.Instruction = nil
			info.Arg = nil
			info.isEndOfSub = false
			accessRange.IsInvalid = true
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
			if isEndOfSubroutinue(instruction.OpcodeType) {
				info.isEndOfSub = true
				accessRange.Max = info.Address
				accessRangeList = append(accessRangeList, accessRange)
				accessRange = &AccessRange{
					Min: info.Address + Address(len(info.Bytes)),
					Max: 0,
				}
			}
		}
		decodeInfoMap[info.Address] = &info
	}
	return &AnalysisInfo{
		decodeInfoMap:   decodeInfoMap,
		accessRangeList: accessRangeList,
	}
}

func collectInvalidChunk(accessInfo *AnalysisInfo) *AnalysisInfo {
	accessRangeList := accessInfo.accessRangeList
	decodeInfoMap := accessInfo.decodeInfoMap
	for _, accessRange := range accessRangeList {
		// if it does not have .db pseudo opcode
		if !accessRange.IsInvalid {
			continue
		}
		// parse invalid opcode into byte codes
		var byteList []byte
		for addr := accessRange.Min; addr <= accessRange.Max; addr++ {
			decodeInfo, ok := decodeInfoMap[addr]
			delete(decodeInfoMap, addr)
			if !ok {
				continue
			}
			byteList = append(byteList, decodeInfo.Bytes...)
		}
		// put them back to the map
		addrBase := accessRange.Min
		var rowData []byte
		var lastDecodeInfo *DecodeInfo
		var addr Address
		var lastAddr Address
		for i, b := range byteList {
			if len(rowData) == 0 {
				addr = addrBase + Address(i)
			}
			rowData = append(rowData, b)
			if len(rowData) > 3 {
				decodeInfo := &DecodeInfo{
					Bytes:   rowData,
					Address: addr,
				}
				lastAddr = addr
				decodeInfoMap[addr] = decodeInfo
				lastDecodeInfo = decodeInfo
				rowData = []byte{}
			}
		}
		if len(rowData) > 0 {
			addr := lastAddr + 4
			decodeInfo := &DecodeInfo{
				Bytes:   rowData,
				Address: addr,
			}
			decodeInfoMap[accessRange.Max] = decodeInfo
			lastDecodeInfo = decodeInfo
		}
		lastDecodeInfo.isEndOfSub = true
	}
	return &AnalysisInfo{
		accessRangeList: accessRangeList,
		decodeInfoMap:   decodeInfoMap,
	}
}

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

func formatOpcode(opcodeType OpcodeType) string {
	return fmt.Sprintf("%s ", OpcodeMap[opcodeType])
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

// DumpPBank ...
func DumpPBank(rom []byte, stupid bool) {
	analysisInfo := analyzePBank(rom)
	if !stupid {
		analysisInfo = collectInvalidChunk(analysisInfo)
	}
	for i := 0x8000; i < 0xFFFF; i++ {
		info, ok := analysisInfo.decodeInfoMap[Address(i)]
		if !ok {
			continue
		}
		fmt.Printf("%s", formatAddress(info.Address))
		fmt.Printf("%s", formatByteCodes(info.Bytes))
		if info.Instruction != nil {
			fmt.Printf("%s", formatOpcode(info.Instruction.OpcodeType))
			fmt.Printf("%s", formatOperand(info.Address, info.Instruction.AddressingType, *info.Arg))
		} else {
			varStr := ".db "
			for _, b := range info.Bytes {
				varStr += fmt.Sprintf("$%02X ", b)
			}
			fmt.Print(varStr)
		}
		fmt.Println()
		if info.isEndOfSub {
			fmt.Println()
		}
	}
}
