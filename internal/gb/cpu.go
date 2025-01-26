package gb

import (
	"fmt"

	"github.com/souzaramon/go-emu-gb/internal/bit"
)

type Instruction struct {
	Type           string
	AddressingMode string
	reg1           string
	cond           string
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

func (c *CPU) CheckCondition() bool {
	flag_z := bit.GetNth(c.f, 7)
	flag_c := bit.GetNth(c.f, 4)

	switch c.CurrentInstruction.cond {
	case "NONE":
		return true
	case "C":
		return flag_c
	case "NC":
		return !flag_c
	case "Z":
		return flag_z
	case "NZ":
		return !flag_z
	}

	return false
}

func (c *CPU) SetZFlag(v bool) {
	c.f = bit.SetNth(c.f, 7, v)
}

func (c *CPU) SetCFlag(v bool) {
	c.f = bit.SetNth(c.f, 4, v)
}

func (c *CPU) SetNFlag(v bool) {
	c.f = bit.SetNth(c.f, 6, v)

}

func (c *CPU) SetHFlag(v bool) {
	c.f = bit.SetNth(c.f, 5, v)
}

func (c *CPU) FetchData() (int, error) {
	switch c.CurrentInstruction.AddressingMode {
	case "IMP":
		return 0, nil

	case "R":
		data, err := c.ReadRegister(c.CurrentInstruction.reg1)

		if err != nil {
			return 0, fmt.Errorf("CPU.ReadRegister failed: %w", err)
		}
		c.data = data
		return 0, nil

	case "R_D8":
		c.data = uint16(c.Bus.Read(c.pc))
		c.pc++
		return 1, nil

	case "D16":
		lo := uint16(c.Bus.Read(c.pc))
		hi := uint16(c.Bus.Read(c.pc + 1))
		c.data = lo | (hi << 8)
		c.pc += 2
		return 2, nil

	default:
		return 0, fmt.Errorf("unknown addressing mode %s for instruction at PC %04X", c.CurrentInstruction.AddressingMode, c.pc)
	}
}

func (c *CPU) ExecInstruction() (int, error) {
	switch c.CurrentInstruction.Type {
	case "NOP":
		return 0, nil

	case "JP":
		// if c.CheckCond() {
		// 	c.pc = c.data
		// 	c.cycles(1)
		// }
		return 0, nil

	case "LD":
		// TODO:
		return 0, nil
	case "XOR":
		// TODO:
		return 0, nil
	case "DI":
		// TODO:
		return 0, nil
	case "DEC":
		// TODO:
		return 0, nil
	default:
		return 0, fmt.Errorf("unknown instruction type '%s' at PC %04X", c.CurrentInstruction.Type, c.pc)
	}
}

func (c *CPU) Tick() (int, error) {
	buf_cycles := 1

	if c.IsHalted {
		return buf_cycles, nil
	}

	c.pc++
	c.CurrentOpCode = c.Bus.Read(c.pc)

	CurrentInstruction, exists := c.Instructions[c.CurrentOpCode]
	c.CurrentInstruction = CurrentInstruction

	if !exists {
		return buf_cycles, fmt.Errorf("unknown instruction (0x%02X) encountered at PC: 0x%04X", c.CurrentOpCode, c.pc)
	}

	cycles_fetch, err := c.FetchData()
	buf_cycles += cycles_fetch

	if err != nil {
		return 0, fmt.Errorf("CPU.FetchData failed: %w", err)
	}

	cycles_exec, err := c.ExecInstruction()
	buf_cycles += cycles_exec

	if err != nil {
		return 0, fmt.Errorf("CPU.ExecInstruction failed: %w", err)
	}

	return buf_cycles, nil
}
