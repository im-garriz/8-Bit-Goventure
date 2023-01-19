package cpu

import (
	"fmt"
	"main/opcodes"
)

type InstructionError struct {
	msg string
}

func (e *InstructionError) Error() string {
	return e.msg
}

type CPU struct {
	Registers CPURegisters
	Decoder   opcodes.Decoder
}

func (cpu *CPU) execute(instruction opcodes.InstructionData) error {
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
		err = cpu.execute(instruction)
		cpu.Registers.Set16bitRegister("PC", nextAddress)
		// Cycles
		if err != nil {
			return err
		}
	}
}
