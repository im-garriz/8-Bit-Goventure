package main

import (
	"fmt"
	"main/opcodes"
)

func main() {
	instructions, err := opcodes.ReadOpcodes(false)
	if err != nil {
		fmt.Println(err)
		return
	}
	cartridge := opcodes.GameBoyROM{}
	err = cartridge.LoadROM("opcodes/etc/snake.gb")
	if err != nil {
		fmt.Println(err)
		return
	}

	decoder := opcodes.Decoder{Cartridge: cartridge, Address: 0, Instructions: instructions}

	var addr uint16
	addr = 0x150

	for i := 0; i < 15; i++ {
		address, instruction, err := decoder.Decode(addr)
		if err != nil {
			fmt.Println(err)
			return
		}
		printableInst := instruction.GetCMD()
		fmt.Printf("%04X %s\n", address, printableInst)
		addr = address
	}

}
