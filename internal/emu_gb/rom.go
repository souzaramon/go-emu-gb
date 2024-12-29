package emu_gb

import (
	"fmt"
	"os"
)

type Header struct {
	Checksum byte
	Title    []byte
	Logo     []byte
}

type ROM struct {
	Header Header
}

func (r *ROM) Load(rom_data []byte) {
	checksum := rom_data[0x014D]
	checksum_acc := 0
	for address := 0x0134; address <= 0x014C; address++ {
		checksum_acc = checksum_acc - int(rom_data[address]) - 1
	}

	if int(checksum) != checksum_acc&0xFF {
		fmt.Println("Error: checksum missmatch")
		os.Exit(1)
	}

	r.Header.Checksum = checksum
	r.Header.Title = rom_data[0x0134:0x0143]
	r.Header.Logo = rom_data[0x0104:0x0133]
}

func (r ROM) String() string {
	return fmt.Sprintf(
		`
Title    : %s
Checksum : %2.2X
		`,
		string(r.Header.Title),
		r.Header.Checksum,
	)
}
