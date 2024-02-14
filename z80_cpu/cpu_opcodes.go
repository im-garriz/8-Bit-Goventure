package cpu

// https://rgbds.gbdev.io/docs/v0.6.1/gbz80.7/

/* ADC */

func (cpu *CPU) _ADC_(target *uint8, value uint8) {

	carry, _ := cpu.Registers.GetFlag("C")
	result := uint16(*target) + uint16(value) + uint16(carry)

	cpu.Registers.SetFlag("C", result > 0xFF)
	cpu.Registers.SetFlag("H", ((*target&0x0F)+(value&0x0F)+carry) > 0x0F)
	cpu.Registers.SetFlag("Z", (result&0xFF) == 0)
	cpu.Registers.SetFlag("N", false)

	*target = uint8(result & 0xFF)
}

func (cpu *CPU) ADC_A_r8(reg string) {

	A, _ := cpu.Registers.Get8bitRegister("A")
	value, _ := cpu.Registers.Get8bitRegister(reg)
	cpu._ADC_(&A, value)
}

func (cpu *CPU) ADC_A_HL() {

	hlPointer, _ := cpu.Registers.Get16bitRegister("HL")
	value := cpu.ReadAddr(hlPointer)
	A, _ := cpu.Registers.Get8bitRegister("A")
	cpu._ADC_(&A, value)
}

func (cpu *CPU) ADC_A_n8(value uint8) {

	A, _ := cpu.Registers.Get8bitRegister("A")
	cpu._ADC_(&A, value)
}

/* ADD */

func (cpu *CPU) _ADD_(target *uint8, value uint8) {

	result := uint16(*target) + uint16(value)

	cpu.Registers.SetFlag("C", result > 0xFF)
	cpu.Registers.SetFlag("H", ((*target&0x0F)+(value&0x0F)) > 0x0F)
	cpu.Registers.SetFlag("Z", (result&0xFF) == 0)
	cpu.Registers.SetFlag("N", false)

	*target = uint8(result & 0xFF)
}

func (cpu *CPU) _ADD_16b(target *uint16, value uint16) {

	result := uint32(*target) + uint32(value)

	cpu.Registers.SetFlag("C", result > 0xFFFF)
	cpu.Registers.SetFlag("H", ((*target&0xFFF)+(value&0xFFF)) > 0x0FFF)
	cpu.Registers.SetFlag("N", false)

	*target = uint16(result & 0xFFFF)
}

func (cpu *CPU) ADD_A_r8(reg string) {
	A, _ := cpu.Registers.Get8bitRegister("A")
	value, _ := cpu.Registers.Get8bitRegister(reg)
	cpu._ADD_(&A, value)
}

func (cpu *CPU) ADD_A_HL() {
	A, _ := cpu.Registers.Get8bitRegister("A")
	hlPointer, _ := cpu.Registers.Get16bitRegister("HL")
	value := cpu.ReadAddr(hlPointer)
	cpu._ADD_(&A, value)
}

func (cpu *CPU) ADD_A_n8(value uint8) {
	A, _ := cpu.Registers.Get8bitRegister("A")
	cpu._ADD_(&A, value)
}

func (cpu *CPU) ADD_HL_r16(reg string) {
	HL, _ := cpu.Registers.Get16bitRegister("HL")
	value, _ := cpu.Registers.Get16bitRegister(reg)
	cpu._ADD_16b(&HL, value)
}

func (cpu *CPU) ADD_HL_SP(reg string) {
	HL, _ := cpu.Registers.Get16bitRegister("HL")
	SP, _ := cpu.Registers.Get16bitRegister("SP")
	cpu._ADD_16b(&HL, SP)
}

func (cpu *CPU) ADD_SP_e8(reg string, value int8) {

	/* CHECK */

	SP, _ := cpu.Registers.Get16bitRegister("SP")

	result := int32(SP) + int32(value)

	cpu.Registers.SetFlag("C", result > 0xFF)
	cpu.Registers.SetFlag("H", (SP^uint16(value)^(uint16(result&0xFFFF)))&0x10 == 0x10)
	cpu.Registers.SetFlag("Z", false)
	cpu.Registers.SetFlag("N", false)

	cpu.Registers.Set16bitRegister("SP", uint16(result))
}

/* AND */

func (cpu *CPU) _AND_(target *uint8, value uint8) {

	*target &= value

	cpu.Registers.SetFlag("C", false)
	cpu.Registers.SetFlag("H", true)
	cpu.Registers.SetFlag("Z", (*target&0xFF) == 0)
	cpu.Registers.SetFlag("N", false)

}

func (cpu *CPU) AND_A_r8(reg string) {

	A, _ := cpu.Registers.Get8bitRegister("A")
	value, _ := cpu.Registers.Get8bitRegister(reg)
	cpu._AND_(&A, value)
}

func (cpu *CPU) AND_A_HL(reg string) {

	A, _ := cpu.Registers.Get8bitRegister("A")
	hlPointer, _ := cpu.Registers.Get16bitRegister("HL")
	value := cpu.ReadAddr(hlPointer)
	cpu._AND_(&A, value)
}

func (cpu *CPU) AND_A_n8(value uint8) {

	A, _ := cpu.Registers.Get8bitRegister("A")
	cpu._AND_(&A, value)
}

/* BIT */

func (cpu *CPU) _BIT_(bit uint8, value uint8) {

	cpu.Registers.SetFlag("Z", ((value>>bit)&1) == 1)
	cpu.Registers.SetFlag("H", true)
	cpu.Registers.SetFlag("N", false)
}

func (cpu *CPU) BIT_u3_r8(bit uint8, reg string) {
	value, _ := cpu.Registers.Get8bitRegister(reg)
	cpu._BIT_(bit, value)
}

func (cpu *CPU) BIT_u3_HL(bit uint8) {
	hlPointer, _ := cpu.Registers.Get16bitRegister("HL")
	value := cpu.ReadAddr(hlPointer)
	cpu._BIT_(bit, value)
}

/* CALL */

func (cpu *CPU) CALL_n16(addr uint16) {

	PC, _ := cpu.Registers.Get16bitRegister("PC")
	cpu.Stack.Push(PC)
	cpu.Registers.Set16bitRegister("PC", addr)
}

func (cpu *CPU) CALL_cc_n16(addr uint16, cond string) {
	if cpu.CCIsSatisfied(cond) {
		cpu.CALL_n16(addr)
	}
}

/* CFF */

func (cpu *CPU) CFF(addr uint16, cond string) {
	cpu.Registers.SetFlag("N", false)
	cpu.Registers.SetFlag("H", false)
	C, _ := cpu.Registers.GetFlag("C")
	var Cbool bool
	if C == 0 {
		Cbool = false
	} else {
		Cbool = true
	}
	cpu.Registers.SetFlag("C", !Cbool)
}

/* CP */

func (cpu *CPU) _CP_(val uint8) {
	A, _ := cpu.Registers.Get8bitRegister("A")
	res := A - val

	cpu.Registers.SetFlag("Z", res == 0)
	cpu.Registers.SetFlag("N", true)
	cpu.Registers.SetFlag("H", ((A&0xF)-(val&0xF)) < 0)
	cpu.Registers.SetFlag("C", A < val)
}

func (cpu *CPU) CP_A_r8(reg string) {
	val, _ := cpu.Registers.Get8bitRegister(reg)
	cpu._CP_(val)
}

func (cpu *CPU) CP_A_HL() {
	hlPointer, _ := cpu.Registers.Get16bitRegister("HL")
	val := cpu.ReadAddr(hlPointer)
	cpu._CP_(val)
}

func (cpu *CPU) CP_A_n8(val uint8) {
	cpu._CP_(val)
}

/* CPL */

func (cpu *CPU) CPL(val uint8) {
	A, _ := cpu.Registers.Get8bitRegister("A")
	cpu.Registers.Set8bitRegister("A", ^A)
	cpu.Registers.SetFlag("N", true)
	cpu.Registers.SetFlag("H", true)
}

/* DAA */

func (cpu *CPU) DAA(val uint8) {
	A, _ := cpu.Registers.Get8bitRegister("A")

	cpu.Registers.Set8bitRegister("A", ^A)
	cpu.Registers.SetFlag("N", true)
	cpu.Registers.SetFlag("H", true)
}
