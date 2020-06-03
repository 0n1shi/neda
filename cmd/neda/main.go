package main

import (
	"bytes"
	"encoding/binary"
	"flag"
	"fmt"
	"os"
	"unsafe"

	"github.com/k-onishi/neda/pkg/neda"
)

const version = "0.7.0"

func main() {
	os.Exit(run())
}

func run() int {
	flag.Usage = func() {
		fmt.Printf("Name: NEDA (NEs rom DisAssembler)\nVersion: %s\nUsage of %s:\n\t%s [Options ...]\nOptions\n", version, os.Args[0], os.Args[0])
		flag.PrintDefaults()
	}
	var (
		ver         = flag.Bool("version", false, "Neda version")
		romFilePath = flag.String("rom", "", "rom file path")
		headerOnly  = flag.Bool("header-only", false, "display header only")
	)
	flag.Parse()
	if *ver {
		fmt.Printf("version: %s\n", version)
		return 0
	}

	if *romFilePath == "" {
		flag.Usage()
		return 1
	}

	// read file from path
	fd, err := os.Open(*romFilePath)
	if err != nil {
		fmt.Println(err.Error())
		return 1
	}
	defer fd.Close()

	// read header
	var header neda.Header
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
	neda.DumpHeader(header)
	if *headerOnly {
		return 0
	}
	neda.DumpPBank(programROM)
	return 0
}
