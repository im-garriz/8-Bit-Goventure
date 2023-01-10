package opcodes

import (
	"encoding/json"
	"fmt"
	"io/ioutil"
)

func ReadOpcodes(printOpcodes bool) (Instructions, error) {

	byteValue, err := ioutil.ReadFile("opcodes/etc/opcodes.json")
	if err != nil {
		return Instructions{}, err
	}

	var instructions Instructions
	json.Unmarshal(byteValue, &instructions)

	if err != nil {
		return Instructions{}, err
	}

	if printOpcodes {
		printInstructions(instructions)
	}

	return instructions, nil
}

func printInstructions(instructions Instructions) {
	fmt.Println("Unprefixed instructions:")
	for opcode, data := range instructions.Unprefixed {
		fmt.Println("Opcode:", opcode)
		fmt.Println("Mnemonic:", data.Mnemonic)
		fmt.Println("Bytes:", data.Bytes)
		fmt.Println("Cycles:", data.Cycles)
		fmt.Println("Operands:")

		for _, operand := range data.Operands {
			fmt.Printf("- Name: %s, Bytes: %d, Immediate: %v\n", operand.Name, operand.Bytes, operand.Immediate)
		}
		fmt.Println("Immediate:", data.Immediate)
		fmt.Println("Flags:", data.Flags)
		fmt.Println()

		//break
	}

	fmt.Println("CB prefixed instructions:")
	for opcode, data := range instructions.CBprefixed {
		fmt.Println("Opcode:", opcode)
		fmt.Println("Mnemonic:", data.Mnemonic)
		fmt.Println("Bytes:", data.Bytes)
		fmt.Println("Cycles:", data.Cycles)
		fmt.Println("Operands:")

		for _, operand := range data.Operands {
			fmt.Printf("- Name: %s, Bytes: %d, Immediate: %v\n", operand.Name, operand.Bytes, operand.Immediate)
		}
		fmt.Println("Immediate:", data.Immediate)
		fmt.Println("Flags:", data.Flags)
		fmt.Println()

		//break
	}

}
