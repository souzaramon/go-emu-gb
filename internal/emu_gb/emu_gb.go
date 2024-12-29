package emu_gb

import "os"

type EmuGB struct {
	isRunning bool
	tickCount int
}

func (e *EmuGB) Run(rom ROM) {
	cpu := CPU{};
	cpu.Init();

	e.isRunning = true;
	e.tickCount = 0;

	for e.isRunning {
		if !cpu.Tick() {
			os.Exit(2);
		}

		e.tickCount += 1;
	}
}