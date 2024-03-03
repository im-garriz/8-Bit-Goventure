// Package disassembler provides functionality for disassembling GameBoy emulator ROMs.
package disassembler

import (
	"bytes"
	"encoding/json"
	"fmt"
	"os"
)

type Instruction struct {
	Opcode     uint8
	CBprefixed bool
	Mnemonic   string
	Bytes      uint8 /* Length */
	Cycles     []uint8
	Operands   []Operand
	Immediate  bool
	Flags      map[string]string
}

type Operand struct {
	Name      string
	Bytes     uint8 /* Length */
	Immediate bool
	HasValue  bool   /* Flag that indicates if operator has a value */
	Value     uint16 /* Value of the operand */
}

// Instructions holds unprefixed and CB-prefixed GameBoy assembly instructions.
type Instructions struct {
	Unprefixed map[string]Instruction // Unprefixed instructions mapped by opcode.
	CBprefixed map[string]Instruction // CB-prefixed instructions mapped by opcode.
}

// GetOpCMD builds the operator command into a string.
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

func (i *Instruction) GetCMD() string {
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

func GetAssemblyInstructions() (*Instructions, error) {

	byteValue, err := os.ReadFile("etc/opcodes.json")
	if err != nil {
		return nil, err
	}

	var instructions Instructions
	json.Unmarshal(byteValue, &instructions)

	if err != nil {
		return nil, err
	}

	return &instructions, nil
}
