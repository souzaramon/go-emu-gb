package gb

type PPU struct {
	Bus          *Bus
	CurrentFrame int
	VideoBuffer  uint32
}

func NewPPU() *PPU {
	return &PPU{
		CurrentFrame: 0,
	}
}

func (p *PPU) Tick() {
	p.CurrentFrame = p.CurrentFrame + 1
}
