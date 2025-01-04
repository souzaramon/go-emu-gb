package gb

type Instruction struct {
	Type           int
	AddressingMode int
	reg1           int
	// reg2           int
	// cond           int
}

var Instructions = map[byte]Instruction{
	0x00: {Type: IN_NOP, AddressingMode: AM_IMP},
	0x05: {Type: IN_DEC, AddressingMode: AM_R, reg1: RT_B},
	0x0E: {Type: IN_LD, AddressingMode: AM_R_D8, reg1: RT_C},
	0xAF: {Type: IN_XOR, AddressingMode: AM_R, reg1: RT_A},
	0xC3: {Type: IN_JP, AddressingMode: AM_D16},
	0xF3: {Type: IN_DI},
}

var TypeNames = map[int]string{
	IN_NONE: "<NONE>",
	IN_NOP:  "NOP",
	IN_LD:   "LD",
	IN_INC:  "INC",
	IN_DEC:  "DEC",
	IN_RLCA: "RLCA",
	IN_ADD:  "ADD",
	IN_RRCA: "RRCA",
	IN_STOP: "STOP",
	IN_RLA:  "RLA",
	IN_JR:   "JR",
	IN_RRA:  "RRA",
	IN_DAA:  "DAA",
	IN_CPL:  "CPL",
	IN_SCF:  "SCF",
	IN_CCF:  "CCF",
	IN_HALT: "HALT",
	IN_ADC:  "ADC",
	IN_SUB:  "SUB",
	IN_SBC:  "SBC",
	IN_AND:  "AND",
	IN_XOR:  "XOR",
	IN_OR:   "OR",
	IN_CP:   "CP",
	IN_POP:  "POP",
	IN_JP:   "JP",
	IN_PUSH: "PUSH",
	IN_RET:  "RET",
	IN_CB:   "CB",
	IN_CALL: "CALL",
	IN_RETI: "RETI",
	IN_LDH:  "LDH",
	IN_JPHL: "JPHL",
	IN_DI:   "DI",
	IN_EI:   "EI",
	IN_RST:  "RST",
	IN_ERR:  "IN_ERR",
	IN_RLC:  "IN_RLC",
	IN_RRC:  "IN_RRC",
	IN_RL:   "IN_RL",
	IN_RR:   "IN_RR",
	IN_SLA:  "IN_SLA",
	IN_SRA:  "IN_SRA",
	IN_SWAP: "IN_SWAP",
	IN_SRL:  "IN_SRL",
	IN_BIT:  "IN_BIT",
	IN_RES:  "IN_RES",
	IN_SET:  "IN_SET",
}

const (
	IN_NONE int = iota
	IN_NOP
	IN_LD
	IN_INC
	IN_DEC
	IN_RLCA
	IN_ADD
	IN_RRCA
	IN_STOP
	IN_RLA
	IN_JR
	IN_RRA
	IN_DAA
	IN_CPL
	IN_SCF
	IN_CCF
	IN_HALT
	IN_ADC
	IN_SUB
	IN_SBC
	IN_AND
	IN_XOR
	IN_OR
	IN_CP
	IN_POP
	IN_JP
	IN_PUSH
	IN_RET
	IN_CB
	IN_CALL
	IN_RETI
	IN_LDH
	IN_JPHL
	IN_DI
	IN_EI
	IN_RST
	IN_ERR
	IN_RLC
	IN_RRC
	IN_RL
	IN_RR
	IN_SLA
	IN_SRA
	IN_SWAP
	IN_SRL
	IN_BIT
	IN_RES
	IN_SET
)

const (
	AM_IMP int = iota
	AM_R_D16
	AM_R_R
	AM_MR_R
	AM_R
	AM_R_D8
	AM_R_MR
	AM_R_HLI
	AM_R_HLD
	AM_HLI_R
	AM_HLD_R
	AM_R_A8
	AM_A8_R
	AM_HL_SPR
	AM_D16
	AM_D8
	AM_D16_R
	AM_MR_D8
	AM_MR
	AM_A16_R
	AM_R_A16
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

const (
	CT_NONE int = iota
	CT_NZ
	CT_Z
	CT_NC
	CT_C
)
