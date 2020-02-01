package main

import (
	"bytes"
	"encoding/binary"
	"flag"
	"fmt"
	"os"
	"unsafe"

	"github.com/k-onishi/nes-rom-analyzer/nesrom"
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

	var header nesrom.Header
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

	header.Dump()

	return 0
}
