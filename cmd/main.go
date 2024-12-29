package main

import (
	"fmt"
	"os"

	"github.com/souzaramon/go-emu-gb/internal/emu_gb"
)

func main() {
	if len(os.Args) < 2 {
		fmt.Println("Error: go run ./cmd/main.go <path>");
		os.Exit(1);
	}

	rom := emu_gb.ROM{};
	rom.Load(os.Args[1]);
	fmt.Println(rom);
	
	e := emu_gb.EmuGB{};
	e.Run(rom);
}