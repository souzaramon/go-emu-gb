package emu_gb

type Bus struct {}

func (b *Bus) Read(address uint16) uint8 {
	return 0
}

func (b *Bus) Write(address uint16, value uint8) {}
