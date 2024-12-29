package main

import (
	"os"

	"github.com/souzaramon/go-emu-gb/internal/emu_gb"
)

func main() {
	e := emu_gb.EmuGB{};
	os.Exit(e.Run());
}