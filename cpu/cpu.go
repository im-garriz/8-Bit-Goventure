package cpu

import (
	"fmt"
	"main/cartridge"

	"github.com/golang-collections/collections/stack"
)

type CPU struct {
	Registers         *CPURegisters
	Decoder           cartridge.Decoder
	memory            GBMemory
	InterruptsEnabled bool
	Stack             stack.Stack // Stack.Push(a), a := Stack.Pop()
}

func (cpu *CPU) Init(cartridgeFile string) error {
	cpu.Registers = GetCPURegisters()
	cpu.Registers.Set16bitRegister("PC", 0x150)
	dec, err := cartridge.GetDecoder(cartridgeFile)
	if err != nil {
		return err
	}
	cpu.Decoder = dec
	cpu.memory = initCPUMemory()
	cpu.InterruptsEnabled = false
	cpu.Stack = *stack.New()
	return nil
}

func (cpu *CPU) execute(instruction cartridge.InstructionData) error {
	switch instruction.Mnemonic {
	case "NOP":
		//fmt.Printf("%s Succesfully executed\n", instruction.GetCMD())
		return nil
	default:
		return &InstructionError{fmt.Sprintf("Cannot execute %v", instruction.GetCMD())}
	}
}

func (cpu *CPU) Run() error {

	for {
		address, err := cpu.Registers.Get16bitRegister("PC")
		if err != nil {
			return err
		}

		nextAddress, instruction, err := cpu.Decoder.Decode(address)
		if err != nil {
			return err
		}
		PC, _ := cpu.Registers.Get16bitRegister("PC")
		fmt.Printf("Instruction: %s PC: %x\n", instruction.GetCMD(), PC)

		err = cpu.execute(instruction)
		cpu.Registers.Set16bitRegister("PC", nextAddress)

		// Cycles
		if err != nil {
			return err
		}
	}
}
