package opcodes

import (
	"bytes"
	"encoding/binary"
	"errors"
	"fmt"
)

type InstructionData struct {
	Mnemonic  string
	Bytes     uint8 /* Length */
	Cycles    []uint8
	Operands  []Operand
	Immediate bool
	Flags     map[string]string
}

type Operand struct {
	Name      string
	Bytes     uint8 /* Length */
	Immediate bool
	HasValue  bool   /* Flag that indicates if operator has a value */
	Value     uint16 /* Value of the operand */
}

type Instructions struct {
	Unprefixed map[string]InstructionData
	CBprefixed map[string]InstructionData
}

func (o *Operand) GetOpCMD() string {

	/* Builds the Operator cmd into a string */

	var buffer bytes.Buffer
	var val string

	if o.HasValue {
		if o.Bytes != 0 {
			val = fmt.Sprintf("0x%x", o.Value) /* Hexadecimal */
		} else {
			val = fmt.Sprintf("0x%d", o.Value) /* Decimal */
		}
	} else {
		val = o.Name
	}

	if o.Immediate {
		fmt.Fprintf(&buffer, "%s", val)
	} else {
		fmt.Fprintf(&buffer, "(%s)", val)
	}

	return buffer.String()
}

func (i *InstructionData) GetCMD() string {
	operands := ""
	for i, operand := range i.Operands {

		if i == 0 {
			operands += operand.GetOpCMD()
		} else {
			operands += ", " + operand.GetOpCMD()
		}
	}

	var buffer bytes.Buffer
	fmt.Fprintf(&buffer, "%-8s %s", i.Mnemonic, operands)

	return buffer.String()
}

type Decoder struct {
	Cartridge    GameBoyROM
	Address      uint16
	Instructions Instructions
}

func (d *Decoder) Read(address uint16, n uint8) (uint16, error) {

	/* Reads n (1 or 2) bytes from address and
	   returns the result in little endian
	*/

	if n < 1 || n > 2 {
		return 0, errors.New(fmt.Sprintf("N must be 1 or 2, %d received", n))
	}

	// Check if tries to read out of the cartridge
	limit := uint16(n) + address
	if limit > uint16(len(d.Cartridge.ROM)) {
		return 0, errors.New(fmt.Sprintf("%d + %d if out of range", address, n))
	}

	data := d.Cartridge.ROM[address:limit]

	if n == 1 {
		return uint16(data[0]), nil /* Only one byte */
	} else {
		return binary.LittleEndian.Uint16(data), nil /* 2 Bytes: to uint16 little endian */
	}

}

func (d *Decoder) Decode(address uint16) (uint16, InstructionData, error) {
	opcode, err := d.Read(address, 1)
	if err != nil {
		return 0, InstructionData{}, err
	}

	address += 1
	var instruction InstructionData

	if opcode == 0xCB {
		opcode, err := d.Read(address, 1)
		if err != nil {
			return 0, InstructionData{}, err
		}
		address += 1
		instruction = d.Instructions.Unprefixed[fmt.Sprintf("0x%02X", opcode)]
	} else {
		instruction = d.Instructions.Unprefixed[fmt.Sprintf("0x%02X", opcode)]
	}

	operandsList := make([]Operand, 0, len(instruction.Operands))

	for _, operand := range instruction.Operands {

		if operand.Bytes > 0 {

			value, err := d.Read(address, operand.Bytes)
			if err != nil {
				return 0, InstructionData{}, err
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
