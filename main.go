package main

import (
	"bytes"
	"encoding/binary"
	"flag"
	"fmt"
	"os"
	"unsafe"

	"github.com/k-onishi/neda/lib"
	"github.com/k-onishi/nesgo"
)

func main() {
	os.Exit(run())
}

func run() int {
	filePath := flag.String("rom", "", "rom file path")
	flag.Parse()
	if *filePath == "" {
		fmt.Println(flag.ErrHelp.Error())
		return 1
	}
	fd, err := os.Open(*filePath)
	if err != nil {
		fmt.Println(err.Error())
		return 1
	}
	defer fd.Close()

	// read header
	var header nesgo.Header
	rawHeader := make([]byte, unsafe.Sizeof(header))
	_, err = fd.Read(rawHeader)
	if err != nil {
		fmt.Println(err.Error())
		return 1
	}
	if err = binary.Read(bytes.NewReader(rawHeader), binary.LittleEndian, &header); err != nil {
		fmt.Println(err.Error())
		return 1
	}
	if !header.IsValid(header.MagicNumber) {
		fmt.Println("invalid ROM")
		return 1
	}

	// read program ROM
	programROM := make([]byte, header.GetProgramBankSize())
	_, err = fd.Read(programROM)
	if err != nil {
		fmt.Println(err.Error())
		return 1
	}

	// read character ROM
	characterROM := make([]byte, header.GetCharacterBankSize())
	_, err = fd.Read(characterROM)
	if err != nil {
		fmt.Println(err.Error())
		return 1
	}

	// show them
	lib.DumpHeader(header)
	lib.DumpPBank(programROM)
	return 0
}
