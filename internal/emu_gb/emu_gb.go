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

	return emuGB
}

func (e *EmuGB) Run() {
	e.tickCounter = 0
	e.isRunning = true

	// rl.InitWindow(width, height, "go-emu-gb")
	// rl.SetTargetFPS(30)

	for e.isRunning {
		if err := e.Cpu.Tick(); err != nil {
			fmt.Println(err)
			os.Exit(2)
		}

		e.tickCounter++
		// rl.WaitTime(0.01)
	}

	// previous_frame := 0

	// for !rl.WindowShouldClose() {
	// 	current_frame := e.Ppu.CurrentFrame

	// 	if previous_frame != current_frame {
	// 		rl.BeginDrawing()
	// 		rl.ClearBackground(rl.White)
	// 		rl.DrawText(fmt.Sprintf("%d", current_frame), 10, 20, 20, rl.Black)
	// 		rl.EndDrawing()
	// 	}

	// 	previous_frame = current_frame
	// }

	// rl.CloseWindow()
}
