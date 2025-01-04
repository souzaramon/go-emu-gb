package gb

import (
	"fmt"
)

type CPU struct {
	CurrentInstruction Instruction
	CurrentOpCode      uint8
	IsHalted           bool
	cycles             func(int)

	memory_destination    uint16
	destination_is_memory bool
	data                  uint16

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
}

func NewCPU() *CPU {
	return &CPU{
		memory_destination:    0,
		destination_is_memory: false,
		pc:                    0x100,
		a:                     0x01,
		f:                     0x0,
		b:                     0x0,
		c:                     0x0,
		d:                     0x0,
		e:                     0x0,
		h:                     0x0,
		l:                     0x0,
		sp:                    0x0,
		cycles:                func(i int) {},
	}
}

func (c *CPU) ReadRegister(regType int) (uint16, error) {
	switch regType {
	case RT_A:
		return uint16(c.a), nil
	case RT_F:
		return uint16(c.f), nil
	case RT_B:
		return uint16(c.b), nil
	case RT_C:
		return uint16(c.c), nil
	case RT_D:
		return uint16(c.d), nil
	case RT_E:
		return uint16(c.e), nil
	case RT_H:
		return uint16(c.h), nil
	case RT_L:
		return uint16(c.l), nil
	case RT_SP:
		return c.sp, nil
	case RT_PC:
		return c.pc, nil
	case RT_AF:
		return (uint16(c.f) << 8) | uint16(c.a), nil
	case RT_BC:
		return (uint16(c.c) << 8) | uint16(c.b), nil
	case RT_DE:
		return (uint16(c.e) << 8) | uint16(c.d), nil
	case RT_HL:
		return (uint16(c.l) << 8) | uint16(c.h), nil
	default:
		return 0, fmt.Errorf("unknown register type (%d) encountered at PC: 0x%04X", regType, c.pc)
	}
}

func (c *CPU) Tick() error {
	c.cycles(1)

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
		return nil

	case AM_R:
		data, err := c.ReadRegister(c.CurrentInstruction.reg1)

		if err != nil {
			return fmt.Errorf("CPU.ReadRegister failed: %w", err)
		}

		c.data = data
		return nil

	case AM_R_D8:
		c.data = uint16(c.Bus.Read(c.pc))
		c.cycles(1)
		c.pc++
		return nil
	case AM_D16:
		lo := uint16(c.Bus.Read(c.pc))
		c.cycles(1)

		hi := uint16(c.Bus.Read(c.pc + 1))
		c.cycles(1)

		c.data = lo | (hi << 8)
		c.pc += 2
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
