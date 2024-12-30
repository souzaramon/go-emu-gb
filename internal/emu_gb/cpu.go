package emu_gb

import (
	"fmt"
)

type CPU struct {
	CurrentInstruction Instruction
	CurrentOpCode      uint8
	IsHalted           bool

	pc uint16
	a  uint8
	f  uint8
	b  uint8
	c  uint8
	d  uint8
	e  uint8
	h  uint8
	l  uint8
	sp uint16

	Bus *Bus
	Ppu *PPU
}

func NewCPU() CPU {
	return CPU{
		pc: 0x100,
		a:  0x01,
		f:  0x0,
		b:  0x0,
		c:  0x0,
		d:  0x0,
		e:  0x0,
		h:  0x0,
		l:  0x0,
		sp: 0x0,
	}
}

// func (c *CPU) Cycle(n int) {
// 	for i := 0; i < n; i++ {
// 		for j := 0; j < 4; j++ {
// 			c.Ppu.Tick()
// 		}
// 	}
// }

func (c *CPU) ReadReg(regType int) {}

func (c *CPU) Tick() error {
	if c.IsHalted {
		return nil
	}

	c.pc++
	c.CurrentOpCode = c.Bus.Read(c.pc)

	CurrentInstruction, exists := Instructions[c.CurrentOpCode]
	c.CurrentInstruction = CurrentInstruction

	if !exists {
		return fmt.Errorf("unknown instruction (0x%02X) encountered at PC: 0x%04X", c.CurrentOpCode, c.pc)
	}

	if err := c.FetchData(); err != nil {
		return fmt.Errorf("CPU.FetchData failed: %w", err)
	}

	if err := c.Run(); err != nil {
		return fmt.Errorf("CPU.Run failed: %w", err)
	}

	return nil
}

func (c *CPU) FetchData() error {
	switch c.CurrentInstruction.AddressingMode {
	case AM_IMP:
		// TODO:
		return nil
	case AM_R:
		// TODO:
		return nil
	case AM_R_D8:
		// TODO:
		return nil
	case AM_D16:
		// TODO:
		return nil
	default:
		return fmt.Errorf("unknown addressing mode %d for instruction at PC %04X", c.CurrentInstruction.AddressingMode, c.pc)
	}
}

func (c *CPU) Run() error {
	switch c.CurrentInstruction.Type {
	case IN_NOP:
		// TODO:
		return nil
	case IN_JP:
		// TODO:
		return nil
	case IN_NONE:
		// TODO:
		return nil
	case IN_LD:
		// TODO:
		return nil
	case IN_XOR:
		// TODO:
		return nil
	case IN_DI:
		// TODO:
		return nil
	case IN_DEC:
		// TODO:
		return nil
	default:
		typeName := TypeNames[c.CurrentInstruction.Type]
		return fmt.Errorf("unknown instruction type '%s' at PC %04X", typeName, c.pc)
	}
}
