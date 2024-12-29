package emu_gb

import (
	"fmt"

	rl "github.com/gen2brain/raylib-go/raylib"
)

const (
	width  = 160 * 2
	height = 144 * 2
)

type EmuGB struct {
	isRunning bool

	Cpu *CPU
	Ppu *PPU
	Bus *Bus
}

func NewEmuGB(rom ROM) EmuGB {
	cpu := NewCPU()
	ppu := NewPPU()
	bus := NewBus()

	bus.Rom = &rom
	bus.Cpu = cpu
	bus.Ppu = ppu

	cpu.Bus = bus
	cpu.Ppu = ppu

	ppu.Bus = bus

	return EmuGB{
		Cpu: cpu,
		Ppu: ppu,
		Bus: bus,
	}
}

func (e *EmuGB) Run() {

	e.isRunning = true

	go func() {
		for e.isRunning {
			e.Cpu.Tick()
			rl.WaitTime(0.01)
		}
	}()

	rl.InitWindow(width, height, "go-emu-gb")
	rl.SetTargetFPS(30)

	previous_frame := 0

	for !rl.WindowShouldClose() {
		current_frame := e.Ppu.CurrentFrame

		if previous_frame != current_frame {
			rl.BeginDrawing()
			rl.ClearBackground(rl.White)
			rl.DrawText(fmt.Sprintf("%d", current_frame), 10, 20, 20, rl.Black)
			rl.EndDrawing()
		}

		previous_frame = current_frame
	}

	rl.CloseWindow()
}
