package main

import (
	"fmt"
	"os"

	"github.com/souzaramon/go-emu-gb/internal/gb"
)

func main() {
	fmt.Println(gb.ASCII())

	if len(os.Args) < 2 {
		fmt.Println("Error: go run ./cmd/main.go <path>")
		os.Exit(1)
	}

	rom_data, err := os.ReadFile(os.Args[1])

	if err != nil {
		fmt.Println("Error: unable to open file")
		os.Exit(1)
	}

	rom := gb.ROM{}
	rom.Load(rom_data)

	e := gb.NewGB(&rom)
	e.Run()
}
