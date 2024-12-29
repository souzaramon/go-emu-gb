package emu_gb

import (
	"fmt"
	"os"
	"strings"
)

type Header struct {
	Checksum byte
	Title 	 []byte
	Logo  	 []byte
}

type ROM struct {
	FilePath string
	Header Header
}

func (r *ROM) Load(path string) {
	rom_data, err := os.ReadFile(os.Args[1]);

	if err != nil {
		fmt.Println("Error: unable to open file");
		os.Exit(1);
	}

	checksum := rom_data[0x014D];
	checksum_acc := 0;
	for address := 0x0134; address <= 0x014C; address++ {
		checksum_acc = checksum_acc - int(rom_data[address]) - 1;
	}

	if int(checksum) != checksum_acc & 0xFF {
		fmt.Println("Error: checksum missmatch");
		os.Exit(1);
	}

	r.FilePath = os.Args[1];
	r.Header.Checksum = checksum;
	r.Header.Title = rom_data[0x0134:0x0143];
	r.Header.Logo = rom_data[0x0104:0x0133];
}

func (r ROM) String() string {
	return fmt.Sprintf(
		`
%s
%s
Title    : %s
Checksum : %2.2X
		`,
		r.FilePath,
		strings.Repeat("-", len(r.FilePath)),
		string(r.Header.Title),
		r.Header.Checksum,
	);
}
