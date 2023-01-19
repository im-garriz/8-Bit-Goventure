package main

import (
	"fmt"
	"main/cpu"
	"main/opcodes"
)

func main() {

	decoder, err := getDecoder()
	if err != nil {
		fmt.Println(err)
		return
	}
	registers := *cpu.GetCPURegisters()
	registers.Set16bitRegister("PC", 0x150)
	cpu := cpu.CPU{Registers: registers, Decoder: decoder}
	err = cpu.Run()
	if err != nil {
		fmt.Println(err)
	}
}

func getDecoder() (opcodes.Decoder, error) {
	instructions, err := opcodes.ReadOpcodes(false)
	if err != nil {
		fmt.Println(err)
		return opcodes.Decoder{}, err
	}
	cartridge := opcodes.GameBoyROM{}
	err = cartridge.LoadROM("opcodes/etc/snake.gb")
	if err != nil {
		fmt.Println(err)
		return opcodes.Decoder{}, err
	}

	decoder := opcodes.Decoder{Cartridge: cartridge, Address: 0, Instructions: instructions}

	return decoder, nil
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
