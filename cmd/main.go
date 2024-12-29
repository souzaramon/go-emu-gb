package main

import (
	"fmt"
	"os"

	"github.com/souzaramon/go-emu-gb/internal/emu_gb"
)

func main() {
	if len(os.Args) < 2 {
		fmt.Println("Error: go run ./cmd/main.go <path>")
		os.Exit(1)
	}

	rom_data, err := os.ReadFile(os.Args[1])

	if err != nil {
		fmt.Println("Error: unable to open file")
		os.Exit(1)
	}

	rom := emu_gb.ROM{}
	rom.Load(rom_data)

	e := emu_gb.CreateEmuGB(rom)
	e.Run()
}
