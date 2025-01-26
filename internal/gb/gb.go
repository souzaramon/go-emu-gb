package gb

import (
	"fmt"
	"os"
	"time"

	rl "github.com/gen2brain/raylib-go/raylib"
)

const (
	width  = 160 * 2
	height = 144 * 2
)

type GB struct {
	isRunning bool
	ticks     int

	Cpu *CPU
	Ppu *PPU
	Bus *Bus
}

func NewGB(rom *ROM) *GB {
	cpu := NewCPU()
	ppu := NewPPU()
	bus := NewBus()

	bus.Rom = rom
	bus.Cpu = cpu
	bus.Ppu = ppu
	cpu.Bus = bus
	ppu.Bus = bus

	GB := &GB{
		ticks:     0,
		isRunning: true,
		Cpu:       cpu,
		Ppu:       ppu,
		Bus:       bus,
	}

	return GB
}

func (e *GB) Run() {
	rl.SetTraceLogLevel(rl.LogError)
	rl.InitWindow(width, height, "go-emu-gb")
	rl.SetTargetFPS(30)

	for e.isRunning && !rl.WindowShouldClose() {
		err, cycles := e.Cpu.Tick()

		if err != nil {
			fmt.Println(err)
			time.Sleep(time.Second * 5)
			os.Exit(2)
		}

		// NOTE: GB cycles
		for i := 0; i < cycles; i++ {
			for j := 0; j < 4; j++ {
				e.ticks++
			}
		}

		rl.BeginDrawing()
		rl.ClearBackground(rl.White)
		rl.DrawText(fmt.Sprintf("%d", e.ticks), 10, 20, 20, rl.Black)
		rl.EndDrawing()
		rl.WaitTime(0.1)
	}

	rl.CloseWindow()
}
