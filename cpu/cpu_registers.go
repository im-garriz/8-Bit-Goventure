package cpu

import (
	"errors"
	"fmt"
)

type CPURegisters struct {
	highRegisters map[string]string
	lowRegisters  map[string]string
	registers16b  []string
	flags         map[string]uint8

	values map[string]uint16
}

func GetCPURegisters() *CPURegisters {
	return &CPURegisters{
		highRegisters: map[string]string{"F": "AF", "C": "BC", "E": "DE", "L": "HL"},
		lowRegisters:  map[string]string{"A": "AF", "B": "BC", "D": "DE", "H": "HL"},
		registers16b:  []string{"AF", "BC", "DE", "HL", "PC", "SP"},
		flags:         map[string]uint8{"c": 4, "h": 5, "n": 6, "z": 7},
		values:        map[string]uint16{"AF": 0x0, "BC": 0x0, "DE": 0x0, "HL": 0x0, "PC": 0x0, "SP": 0x0},
	}
}

func (c *CPURegisters) Get8bitRegister(register string) (uint8, error) {

	if c.registerIsHigh(register) {
		return getMostSignificantByte(c.values[register]), nil
	} else if c.registerIsLow(register) {
		return getLeastSignificantByte(c.values[register]), nil
	} else if c.registerIsFlag(register) {
		return c.flags[register], nil
	} else if c.registerIs16Bit(register) {
		return 0, errors.New(fmt.Sprintf("%s register is 16 bit, not 8", register))
	}

	return 0, errors.New(fmt.Sprintf("Register %s does not exist", register))
}

func (c *CPURegisters) Set8bitRegister(register string, val uint8) error {

	if c.registerIsHigh(register) {
		v := c.values[c.highRegisters[register]]
		setMostSignificantByte(&v, val)
		return nil
	} else if c.registerIsLow(register) {
		v := c.values[c.lowRegisters[register]]
		setMostSignificantByte(&v, val)
		return nil
	} else if c.registerIsFlag(register) {

	} else if c.registerIs16Bit(register) {

	}

	return errors.New(fmt.Sprintf("Register %s does not exist", register))

}

func (c *CPURegisters) registerIsHigh(reg string) bool {
	_, exists := c.highRegisters[reg]
	return exists
}

func (c *CPURegisters) registerIsLow(reg string) bool {
	_, exists := c.lowRegisters[reg]
	return exists
}

func (c *CPURegisters) registerIsFlag(reg string) bool {
	_, exists := c.flags[reg]
	return exists
}

func (c *CPURegisters) registerIs16Bit(reg string) bool {
	for _, register := range c.registers16b {
		if register == reg {
			return true
		}
	}
	return false
}
