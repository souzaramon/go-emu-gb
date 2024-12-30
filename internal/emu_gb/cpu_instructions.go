package emu_gb

const (
	IN_HALT int = iota
)

var InstructionSet = [0x100]Instruction{
	0x76: {Type: IN_HALT},
}

var TypeNames = map[int]string{
	IN_HALT: "HALT",
}

type Instruction struct {
	Type           int
	AddressingMode int
}
