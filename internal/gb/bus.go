package gb

import (
	"fmt"
	"os"
)

type Bus struct {
	Rom *ROM
	Cpu *CPU
	Ppu *PPU
}

func NewBus() *Bus {
	return &Bus{}
}

func (b *Bus) Read(address uint16) byte {
	switch {
	case address < 0x8000:
		return b.Rom.Read(address)
	default:
		fmt.Println("(Bus) error: Out of bounds read")
		os.Exit(2)
		return 0
	}
}

func (b *Bus) Write(address uint16, value byte) {
	switch {
	case address < 0x8000:
		b.Rom.Write(address, value)
	default:
		fmt.Println("(Bus) error: Out of bounds write")
		os.Exit(2)
	}
}
