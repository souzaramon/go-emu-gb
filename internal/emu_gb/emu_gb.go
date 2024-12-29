package emu_gb

type EmuGB struct {
	isRunning bool
	tickCount int
}

func (e *EmuGB) Run() int {
	cpu := CPU{};
	cpu.Init();

	e.isRunning = true;
	e.tickCount = 0;

	for e.isRunning {
		if !cpu.Tick() {
			return 1;
		}

		e.tickCount += 1;
	}

	return 0;
}