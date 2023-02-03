package cpu

import "strings"

type InstructionError struct {
	msg string
}

func (e *InstructionError) Error() string {
	return e.msg
}

func getMostSignificantByte(value uint16) byte {
	return uint8(value >> 8)
}

func getLeastSignificantByte(value uint16) byte {
	return uint8(value & 0xFF)
}

func setMostSignificantByte(value *uint16, newByte byte) {
	*value = (*value & 0x00FF) | (uint16(newByte) << 8)
}

func setLeastSignificantByte(value *uint16, newByte byte) {
	*value = (*value & 0xFF00) | uint16(newByte)
}

func (cpu *CPU) CCIsSatisfied(cond string) bool {

	if strings.HasSuffix(cond, "Z") {

		Z, _ := cpu.Registers.GetFlag("Z")

		if strings.HasPrefix(cond, "N") {
			if Z == 0 {
				return true
			} else {
				return false
			}
		} else {
			if Z == 1 {
				return true
			} else {
				return false
			}
		}

	} else {
		C, _ := cpu.Registers.GetFlag("C")

		if strings.HasPrefix(cond, "N") {
			if C == 0 {
				return true
			} else {
				return false
			}
		} else {
			if C == 1 {
				return true
			} else {
				return false
			}
		}
	}
}
