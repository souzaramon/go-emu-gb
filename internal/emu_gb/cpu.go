package emu_gb

import (
	"fmt"
	"os"
)

const (
	RT_NONE int = iota
	RT_A
	RT_F
	RT_B
	RT_C
	RT_D
	RT_E
	RT_H
	RT_L
	RT_AF
	RT_BC
	RT_DE
	RT_HL
	RT_SP
	RT_PC
)

type CPU struct {
	Bus *Bus
	Ppu *PPU

	CurrentInstruction *Instruction
	CurrentOpCode      uint8
	IsHalted           bool

	// NOTE: Registers
	a  uint8
	pc uint16
	// f  uint8
	// b  uint8
	// c  uint8
	// d  uint8
	// e  uint8
	// h  uint8
	// l  uint8
	// sp uint16
}

func NewCPU() *CPU {
	return &CPU{
		pc: 0x100,
		a:  0x01,
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

func (c *CPU) Tick() bool {
	if c.IsHalted {
		return true
	}

	c.pc++
	c.CurrentOpCode = c.Bus.Read(c.pc)
	c.CurrentInstruction = &InstructionSet[0x76]

	typeName := TypeNames[c.CurrentInstruction.Type]
	fmt.Printf("[%04X] %s\n", c.pc, typeName)

	c.FetchData()

	if c.CurrentInstruction == nil {
		fmt.Printf("(CPU) error: Unknown instruction (%02X)\n", c.CurrentInstruction)
		os.Exit(2)
	}

	c.Run()

	return true
}

func (c *CPU) FetchData() {
	// switch c.CurrentInstruction.AddressingMode {
	// default:
	// 	fmt.Println("Error: unknown addressing mode")
	// }
}

func (c *CPU) Run() {
	switch c.CurrentInstruction.Type {
	case IN_HALT:
		c.IsHalted = true
	default:
		fmt.Println("(CPU) error: unknown instruction type")
	}
}
