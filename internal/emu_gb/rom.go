package emu_gb

import (
	"fmt"
	"os"
)

type Header struct {
	Title []byte
}

type ROM struct {
	Header Header
}

func (r *ROM) Load(path string) {
	rom_data, err := os.ReadFile(os.Args[1]);

	if err != nil {
		fmt.Println("Error: unable to open file")
		os.Exit(1)

	}

	r.Header.Title = rom_data[0x0134:0x0143]
}

func (r ROM) String() string {
	return fmt.Sprintf(
		`
ROM: %s
------------------------
		`,
		string(r.Header.Title),
	);
}
