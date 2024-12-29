package emu_gb

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

func (b *Bus) Read(address uint16) uint8 {
	switch {
	case address < 0x8000:
		return b.Rom.Read(address)
	default:
		fmt.Println("Error: Out of bounds read")
		os.Exit(2)
		return 0
	}
}

func (b *Bus) Write(address uint16, value uint8) {
	switch {
	case address < 0x8000:
		b.Rom.Write(address, value)
	default:
		fmt.Println("Error: Out of bounds write")
		os.Exit(2)
	}
}
