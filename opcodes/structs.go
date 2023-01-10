package opcodes

import (
	"bytes"
	"encoding/binary"
	"errors"
	"fmt"
	"io/ioutil"
)

type InstructionData struct {
	Mnemonic  string
	Bytes     uint8
	Cycles    []uint8
	Operands  []Operand
	Immediate bool
	Flags     map[string]string
}

type Operand struct {
	Name      string
	Bytes     uint8
	Immediate bool
	Value     uint16 /* Value of the operand*/
}

type Instructions struct {
	Unprefixed map[string]InstructionData
	CBprefixed map[string]InstructionData
}

func (o *Operand) GetOpCMD() string {

	var buffer bytes.Buffer
	var val string

	if o.Value != 0 {
		if o.Bytes != 0 {
			val = fmt.Sprintf("%x", o.Value) /* Hexadecimal */
		} else {
			val = fmt.Sprintf("%d", o.Value) /* Decimal */
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

type GameBoyROM struct {
	ROM []byte
}

func (gbr *GameBoyROM) LoadROM(filepath string) error {
	rom, err := ioutil.ReadFile(filepath)
	if err != nil {
		return err
	}
	gbr.ROM = rom
	return nil
}

type Decoder struct {
	Cartridge    GameBoyROM
	Address      uint16
	Instructions Instructions
}

func (d *Decoder) Read(address uint16, count uint16) (uint16, error) {

	limit := count + address
	if limit > uint16(len(d.Cartridge.ROM)) {
		// error
		return 0, errors.New(fmt.Sprintf("%d + %d if out of range", address, count))
	}

	data := d.Cartridge.ROM[address : address+count]
	return uint16(binary.LittleEndian.Uint32(data)), nil
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
		instruction = d.Instructions.CBprefixed[fmt.Sprintf("%x", opcode)]

	} else {
		instruction = d.Instructions.Unprefixed[fmt.Sprintf("%x", opcode)]
	}

	operandsList := make([]Operand, 0, len(instruction.Operands))

	for i, operand := range instruction.Operands {
		if operand.Bytes > 0 {

			value, err := d.Read(address, uint16(operand.Bytes))
			if err != nil {
				return 0, InstructionData{}, err
			}

			address += uint16(operand.Bytes)
			operand.Value = value

		}

		operandsList[i] = operand
	}

	instruction.Operands = operandsList

	return address, instruction, nil
}
