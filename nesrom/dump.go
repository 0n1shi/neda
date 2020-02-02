package nesrom

import "fmt"

// DumpHeader ...
func DumpHeader(h Header) {
	fmt.Printf("magic number: %s\n", h.MagicNumber)
	fmt.Printf("program bank count: %d (%d bytes)\n", h.ProgramBankCount, h.GetProgramBankSize())
	fmt.Printf("character bank count: %d (%d bytes)\n", h.CharacterBankCount, h.GetCharacterBankSize())
	fmt.Printf("mapper type: %d (%s)\n", h.GetMapper(), h.GetMapperType())
	fmt.Println()
}

func parseOperand(currentAddr int, addrType AddressingType, val int) string {
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
		return fmt.Sprintf("$%04X", currentAddr+val)
	case AddressingTypeAccumulator:
		return "A"
	}
	panic("invalid addressing type")
}

// DumpProgramBank ...
func DumpProgramBank(rom []byte) {
	skipCounter := 0
	for i := 0; i < len(rom); {
		addr := i
		fmt.Printf("0x%04X\t", addr)

		opecode := int(rom[i])
		instruction, ok := InstructionMap[opecode]
		i++
		if !ok {
			skipCounter++
			fmt.Printf("$%02X\n", opecode)
			continue
			// panic("invalid opecode")
		} else {
			if skipCounter > 0 {
				//fmt.Printf("\n(... skip %d bytes)\n\n", skipCounter)
				skipCounter = 0
			}
		}

		argCount := instruction.bytes - 1
		args := make([]byte, argCount)
		argValue := 0
		for j := 0; j < argCount; j++ {
			v := rom[i]
			args[j] = v
			argValue += (int(v) << (8 * j))
			i++
		}

		fmt.Printf("%02X ", opecode)
		for j := 0; j < argCount; j++ {
			fmt.Printf("%02X ", args[j])
		}
		tabFormatStr := "\t"
		if argCount < 2 {
			tabFormatStr = "\t\t"
		}
		fmt.Printf(tabFormatStr)
		fmt.Printf("%s ", instruction.opcodeType)
		fmt.Println(parseOperand(addr, instruction.addressingType, argValue))
	}
}
