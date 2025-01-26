package gb

import (
	"fmt"
)

type Instruction struct {
	Type           string
	AddressingMode string
	reg1           string
	// reg2           string
	// cond           int
}

type CPU struct {
	Instructions       map[byte]Instruction
	CurrentInstruction Instruction
	CurrentOpCode      byte

	IsHalted              bool
	memory_destination    uint16
	destination_is_memory bool
	data                  uint16

	pc uint16
	a  byte
	f  byte
	b  byte
	c  byte
	d  byte
	e  byte
	h  byte
	l  byte
	sp uint16

	// cycles func(int)
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
		// cycles:                func(i int) {},
		Instructions: map[byte]Instruction{
			0x00: {Type: "NOP", AddressingMode: "IMP"},
			0x05: {Type: "DEC", AddressingMode: "R", reg1: "B"},
			0x0E: {Type: "LD", AddressingMode: "R_D8", reg1: "C"},
			0xAF: {Type: "XOR", AddressingMode: "R", reg1: "A"},
			0xC3: {Type: "JP", AddressingMode: "D16"},
			0xF3: {Type: "DI"},
		},
	}
}

func (c *CPU) ReadRegister(name string) (uint16, error) {
	switch name {
	case "A":
		return uint16(c.a), nil
	case "F":
		return uint16(c.f), nil
	case "B":
		return uint16(c.b), nil
	case "C":
		return uint16(c.c), nil
	case "D":
		return uint16(c.d), nil
	case "E":
		return uint16(c.e), nil
	case "H":
		return uint16(c.h), nil
	case "L":
		return uint16(c.l), nil
	case "SP":
		return c.sp, nil
	case "PC":
		return c.pc, nil
	case "AF":
		return (uint16(c.f) << 8) | uint16(c.a), nil
	case "BC":
		return (uint16(c.c) << 8) | uint16(c.b), nil
	case "DE":
		return (uint16(c.e) << 8) | uint16(c.d), nil
	case "HL":
		return (uint16(c.l) << 8) | uint16(c.h), nil
	default:
		return 0, fmt.Errorf("unknown register type (%s) encountered at PC: 0x%04X", name, c.pc)
	}
}

// func (c *CPU) CheckCond() bool {
// 	flag_z := bit.GetNth(c.f, 7)
// 	flag_c := bit.GetNth(c.f, 4)

// 	switch c.CurrentInstruction.cond {
// 	case "NONE":
// 		return true
// 	case "C":
// 		return flag_c
// 	case "NC":
// 		return !flag_c
// 	case "Z":
// 		return flag_z
// 	case "NZ":
// 		return !flag_z
// 	}

// 	return false
// }

func (c *CPU) FetchData() (error, int) {
	switch c.CurrentInstruction.AddressingMode {
	case "IMP":
		return nil, 0

	case "R":
		data, err := c.ReadRegister(c.CurrentInstruction.reg1)

		if err != nil {
			return fmt.Errorf("CPU.ReadRegister failed: %w", err), 0
		}
		c.data = data
		return nil, 0

	case "R_D8":
		c.data = uint16(c.Bus.Read(c.pc))
		c.pc++
		return nil, 1

	case "D16":
		lo := uint16(c.Bus.Read(c.pc))
		hi := uint16(c.Bus.Read(c.pc + 1))
		c.data = lo | (hi << 8)
		c.pc += 2
		return nil, 2

	default:
		return fmt.Errorf("unknown addressing mode %s for instruction at PC %04X", c.CurrentInstruction.AddressingMode, c.pc), 0
	}
}

func (c *CPU) ExecInstruction() (error, int) {
	switch c.CurrentInstruction.Type {
	case "NOP":
		return nil, 0

	case "JP":
		// if c.CheckCond() {
		// 	c.pc = c.data
		// 	c.cycles(1)
		// }
		return nil, 0

	case "LD":
		// TODO:
		return nil, 0
	case "XOR":
		// TODO:
		return nil, 0
	case "DI":
		// TODO:
		return nil, 0
	case "DEC":
		// TODO:
		return nil, 0
	default:
		return fmt.Errorf("unknown instruction type '%s' at PC %04X", c.CurrentInstruction.Type, c.pc), 0
	}
}

func (c *CPU) Tick() (error, int) {
	buf_cycles := 1

	if c.IsHalted {
		return nil, buf_cycles
	}

	c.pc++
	c.CurrentOpCode = c.Bus.Read(c.pc)

	CurrentInstruction, exists := c.Instructions[c.CurrentOpCode]
	c.CurrentInstruction = CurrentInstruction

	if !exists {
		return fmt.Errorf("unknown instruction (0x%02X) encountered at PC: 0x%04X", c.CurrentOpCode, c.pc), buf_cycles
	}

	err, cycles_fetch := c.FetchData()
	buf_cycles += cycles_fetch

	if err != nil {
		return fmt.Errorf("CPU.FetchData failed: %w", err), 0
	}

	err, cycles_exec := c.ExecInstruction()
	buf_cycles += cycles_exec

	if err != nil {
		return fmt.Errorf("CPU.ExecInstruction failed: %w", err), 0
	}

	return nil, buf_cycles
}
