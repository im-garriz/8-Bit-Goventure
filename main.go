package main

import (
	"fmt"
	"main/cpu"
)

func main() {

	cpu := cpu.CPU{}
	cpu.Init("cartridge/etc/snake.gb")
	err := cpu.Run()
	if err != nil {
		fmt.Println(err)
	}
}

/*func main() {
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

	for i := 0; i < 16; i++ {
		address, instruction, err := decoder.Decode(addr)
		if err != nil {
			fmt.Println(err)
			return
		}

		printableInst := instruction.GetCMD()
		fmt.Printf("%04X %s\n", addr, printableInst)
		addr = address
	}
}*/
