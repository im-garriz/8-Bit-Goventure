// Package disassembler provides functionality for disassembling GameBoy emulator ROMs.
package disassembler

import (
	"encoding/binary"
	"fmt"
)

// Dissasembler represents a GameBoy ROM disassembler.
type Disassembler struct {
	Cartridge    *GameBoyROM   // GameBoy ROM to disassemble.
	Address      uint16        // Current address in the disassembler.
	Instructions *Instructions // Set of GameBoy assembly instructions.
}

// Read reads n (1 or 2) bytes from the specified address and returns the result in little endian.
// It returns the read value and an error, if any.
func (d *Disassembler) Read(address uint16, n uint8) (uint16, error) {

	/* Reads n (1 or 2) bytes from address and
	   returns the result in little endian
	*/

	if n < 1 || n > 2 {
		return 0, fmt.Errorf("n must be 1 or 2, %d received", n)
	}

	// Check if tries to read out of the cartridge
	limit := uint16(n) + address
	if limit > uint16(len(d.Cartridge.ROM)) {
		return 0, fmt.Errorf("%d + %d if out of range", address, n)
	}

	data := d.Cartridge.ROM[address:limit]

	if n == 1 {
		return uint16(data[0]), nil /* Only one byte */
	} else {
		return binary.LittleEndian.Uint16(data), nil /* 2 Bytes: to uint16 little endian */
	}

}

// Decode decodes the instruction at the specified address and returns the updated address, decoded instruction, and an error, if any.
func (d *Disassembler) Decode(address uint16) (uint16, Instruction, error) {
	opcode, err := d.Read(address, 1)
	if err != nil {
		return 0, Instruction{}, err
	}

	address += 1
	var instruction Instruction

	if opcode == 0xCB {
		opcode, err := d.Read(address, 1)
		if err != nil {
			return 0, Instruction{}, err
		}
		address += 1
		instruction = d.Instructions.CBprefixed[fmt.Sprintf("0x%02X", opcode)]
		instruction.CBprefixed = true
	} else {
		instruction = d.Instructions.Unprefixed[fmt.Sprintf("0x%02X", opcode)]
		instruction.CBprefixed = false
	}

	instruction.Opcode = uint8(opcode)

	operandsList := make([]Operand, 0, len(instruction.Operands))

	for _, operand := range instruction.Operands {

		if operand.Bytes > 0 {

			value, err := d.Read(address, operand.Bytes)
			if err != nil {
				return 0, Instruction{}, err
			}

			address += uint16(operand.Bytes)

			operand.HasValue = true
			operand.Value = value
		}

		operandsList = append(operandsList, operand)
	}

	instruction.Operands = operandsList

	return address, instruction, nil
}

// GetDissassembler initializes a Dissasembler for the specified GameBoy ROM file and returns a pointer to it.
// It takes the file path of the GameBoy ROM and returns a Dissasembler and an error, if any.
func GetDissassembler(cartridgeFile string, opcodesJSONFile string) (*Disassembler, error) {
	instructions, err := GetAssemblyInstructions(opcodesJSONFile)
	if err != nil {
		return nil, err
	}
	gameBoyRom, err := LoadROM(cartridgeFile)
	if err != nil {
		return nil, err
	}

	dissasembler := Disassembler{Cartridge: gameBoyRom, Address: 0, Instructions: instructions}

	return &dissasembler, nil
}
