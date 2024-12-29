package emu_gb

const clock = 4

type CPU struct {
	Bus *Bus
	Ppu *PPU
}

func NewCPU() *CPU {
	return &CPU{}
}

func (c *CPU) Cycle(n int) {
	for i := 0; i < n; i++ {
		for j := 0; j < clock; j++ {
			c.Ppu.Tick()
		}
	}
}

func (c *CPU) Tick() bool {
	c.Cycle(1)

	return true
}
