package emu_gb

import (
	"fmt"
	"os"
	// rl "github.com/gen2brain/raylib-go/raylib"
)

// const (
// 	width  = 160 * 2
// 	height = 144 * 2
// )

type EmuGB struct {
	isRunning   bool
	tickCounter int

	Cpu *CPU
	Ppu *PPU
	Bus *Bus
}

func NewEmuGB(rom *ROM) *EmuGB {
	cpu := NewCPU()
	ppu := NewPPU()
	bus := NewBus()

	bus.Rom = rom
	bus.Cpu = cpu
	bus.Ppu = ppu
	cpu.Bus = bus
	ppu.Bus = bus

	emuGB := &EmuGB{
		tickCounter: 0,
		isRunning:   true,
		Cpu:         cpu,
		Ppu:         ppu,
		Bus:         bus,
	}

	cpu.cycles = func(i int) {
		emuGB.tickCounter++
	}

	return emuGB
}

func (e *EmuGB) Run() {
	// rl.InitWindow(width, height, "go-emu-gb")
	// rl.SetTargetFPS(30)

	// for e.isRunning && !rl.WindowShouldClose() {
	for e.isRunning {

		if err := e.Cpu.Tick(); err != nil {
			fmt.Println(err)
			os.Exit(2)
		}

		// rl.BeginDrawing()
		// rl.ClearBackground(rl.White)
		// rl.DrawText(fmt.Sprintf("%d", e.tickCounter), 10, 20, 20, rl.Black)
		// rl.EndDrawing()
		// rl.WaitTime(0.1)
	}

	// rl.CloseWindow()
}
