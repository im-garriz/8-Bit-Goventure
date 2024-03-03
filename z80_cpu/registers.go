package z80_cpu

import (
	"errors"
	"fmt"
	"strings"
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
		lowRegisters:  map[string]string{"F": "AF", "C": "BC", "E": "DE", "L": "HL"},
		highRegisters: map[string]string{"A": "AF", "B": "BC", "D": "DE", "H": "HL"},
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
		return 0, errors.New(fmt.Sprintf("%s is a FLAG, not a register", register))
	} else if c.registerIs16Bit(register) {
		return 0, errors.New(fmt.Sprintf("%s register is 16 bit, not 8", register))
	}

	return 0, errors.New(fmt.Sprintf("Register %s does not exist", register))
}

func (c *CPURegisters) Set8bitRegister(register string, val uint8) error {

	if c.registerIsHigh(register) {
		v := c.values[c.highRegisters[register]]
		setMostSignificantByte(&v, val)
		c.values[c.highRegisters[register]] = v
		return nil
	} else if c.registerIsLow(register) {
		v := c.values[c.lowRegisters[register]]
		setLeastSignificantByte(&v, val)
		c.values[c.lowRegisters[register]] = v
		return nil
	} else if c.registerIsFlag(register) {
		return errors.New(fmt.Sprintf("%s is a FLAG, not a register", register))
	} else if c.registerIs16Bit(register) {
		return errors.New(fmt.Sprintf("%s register is 16 bit, not 8", register))
	}

	return errors.New(fmt.Sprintf("Register %s does not exist", register))

}

func (c *CPURegisters) GetFlag(flag string) (uint8, error) {

	flag = strings.ToLower(flag)

	if c.registerIsFlag(flag) {
		flagBit := c.flags[flag]
		/* c.values["AF"] >> flagBit (4): shifts to right until flagBit (4) bit is the last
		then does logical AND (&) with it and returns */
		return (uint8(c.values["AF"]) >> flagBit) & 1, nil
	} else if c.registerIs16Bit(flag) || c.registerIsHigh(flag) || c.registerIsLow(flag) {
		return 0, errors.New(fmt.Sprintf("%s is a register, not flag", flag))
	}

	return 0, errors.New(fmt.Sprintf("Flag %s does not exist", flag))
}

func (c *CPURegisters) SetFlag(flag string, value bool) error {

	flag = strings.ToLower(flag)

	if c.registerIsFlag(flag) {

		flagBit := c.flags[flag]
		if value == false {
			/* 1 << flagBit (5) = 00100000 (5th pos)
			^(1 << flagBit (5)) = ^(00100000) = 11011111 (bitwise negation)
			By doing bitwise AND (&) 5th bit is set to 0 and the rest remain as they were */
			c.values["AF"] = c.values["AF"] & ^(1 << flagBit)
		} else {
			/* 1 << flagBit (5) = 00100000 (5th pos)
			By doing bitwise OR (|) 5th bit is set to 1 and the rest remain as they were */
			c.values["AF"] = c.values["AF"] | (1 << flagBit)
		}
		return nil

	} else if c.registerIs16Bit(flag) || c.registerIsHigh(flag) || c.registerIsLow(flag) {
		return errors.New(fmt.Sprintf("%s is a register, not flag", flag))
	}
	return errors.New(fmt.Sprintf("Flag %s does not exist", flag))
}

func (c *CPURegisters) Get16bitRegister(register string) (uint16, error) {

	if c.registerIs16Bit(register) {
		return c.values[register], nil
	} else if c.registerIsHigh(register) || c.registerIsLow(register) {
		return 0, errors.New(fmt.Sprintf("%s is a 8 bit register, not 16", register))
	} else if c.registerIsFlag(register) {
		return 0, errors.New(fmt.Sprintf("%s is a FLAG, not a register", register))
	}

	return 0, errors.New(fmt.Sprintf("Register %s does not exist", register))
}

func (c *CPURegisters) Set16bitRegister(register string, val uint16) error {

	if c.registerIs16Bit(register) {
		c.values[register] = val
		return nil
	} else if c.registerIsHigh(register) || c.registerIsLow(register) {
		return errors.New(fmt.Sprintf("%s is a 8 bit register, not 16", register))
	} else if c.registerIsFlag(register) {
		return errors.New(fmt.Sprintf("%s is a FLAG, not a register", register))
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
