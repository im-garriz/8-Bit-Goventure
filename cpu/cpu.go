package cpu

import (
	"fmt"
	"main/cpu/disassembler"

	"github.com/golang-collections/collections/stack"
)

type CPU struct {
	Registers         *CPURegisters
	Disassembler      *disassembler.Disassembler
	memory            GBMemory
	InterruptsEnabled bool
	Stack             stack.Stack // Stack.Push(a), a := Stack.Pop()
}

func (cpu *CPU) Init(cartridgeFile string) error {
	cpu.Registers = GetCPURegisters()
	cpu.Registers.Set16bitRegister("PC", 0x150)
	dec, err := disassembler.GetDissassembler(cartridgeFile)
	if err != nil {
		return err
	}
	cpu.Disassembler = dec
	cpu.memory = initCPUMemory()
	cpu.InterruptsEnabled = false
	cpu.Stack = *stack.New()
	return nil
}

func (cpu *CPU) execute(instruction disassembler.Instruction) error {
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

		nextAddress, instruction, err := cpu.Disassembler.Decode(address)
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
