package z80_cpu

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

/* CP */

func (cpu *CPU) _CP_(target *uint8, value uint8) {

	res := *target - value

	cpu.Registers.SetFlag("Z", res == 0)
	cpu.Registers.SetFlag("N", true)
	cpu.Registers.SetFlag("H", *target&0xF < value&0xF)
	cpu.Registers.SetFlag("C", *target < value)
}

func (cpu *CPU) CP_A_r8(reg string) {
	A, _ := cpu.Registers.Get8bitRegister("A")
	val, _ := cpu.Registers.Get8bitRegister(reg)

	cpu._CP_(&A, val)
}

func (cpu *CPU) CP_A_HL() {
	A, _ := cpu.Registers.Get8bitRegister("A")
	hlPointer, _ := cpu.Registers.Get16bitRegister("HL")
	val := cpu.ReadAddr(hlPointer)
	cpu._CP_(&A, val)
}

func (cpu *CPU) CP_A_n8(val uint8) {
	A, _ := cpu.Registers.Get8bitRegister("A")
	cpu._CP_(&A, val)
}

/* DEC */

func (cpu *CPU) _DEC_(target *uint8) {

	*target--

	cpu.Registers.SetFlag("Z", *target == 0)
	cpu.Registers.SetFlag("N", true)
	cpu.Registers.SetFlag("H", 0x01 > *target&0x0F)
}

func (cpu *CPU) DEC_r8(reg string) {
	val, _ := cpu.Registers.Get8bitRegister(reg)
	cpu._DEC_(&val)
}

func (cpu *CPU) DEC_HL() {
	hlPointer, _ := cpu.Registers.Get16bitRegister("HL")
	val := cpu.ReadAddr(hlPointer)
	cpu._DEC_(&val)
}

/* INC */

func (cpu *CPU) _INC_(target *uint8) {

	cpu.Registers.SetFlag("H", (*target&0x0F+0x01) > 0x0F)

	*target++

	cpu.Registers.SetFlag("Z", *target == 0)
	cpu.Registers.SetFlag("N", false)

}

func (cpu *CPU) INC_r8(reg string) {
	val, _ := cpu.Registers.Get8bitRegister(reg)
	cpu._INC_(&val)
}

func (cpu *CPU) INC_HL() {
	hlPointer, _ := cpu.Registers.Get16bitRegister("HL")
	val := cpu.ReadAddr(hlPointer)
	cpu._INC_(&val)
}

/* OR */

func (cpu *CPU) _OR_(target *uint8, value uint8) {

	*target |= value

	cpu.Registers.SetFlag("C", false)
	cpu.Registers.SetFlag("H", false)
	cpu.Registers.SetFlag("Z", (*target&0xFF) == 0)
	cpu.Registers.SetFlag("N", false)

}

func (cpu *CPU) OR_A_r8(reg string) {

	A, _ := cpu.Registers.Get8bitRegister("A")
	value, _ := cpu.Registers.Get8bitRegister(reg)
	cpu._OR_(&A, value)
}

func (cpu *CPU) OR_A_HL(reg string) {

	A, _ := cpu.Registers.Get8bitRegister("A")
	hlPointer, _ := cpu.Registers.Get16bitRegister("HL")
	value := cpu.ReadAddr(hlPointer)
	cpu._OR_(&A, value)
}

func (cpu *CPU) OR_A_n8(value uint8) {

	A, _ := cpu.Registers.Get8bitRegister("A")
	cpu._OR_(&A, value)
}

/* SBC */

func (cpu *CPU) _SBC_(target *uint8, value uint8) {

	carry, _ := cpu.Registers.GetFlag("C")

	result := int16(*target) - int16(value) - int16(carry)

	cpu.Registers.SetFlag("C", (value+carry) > *target)
	cpu.Registers.SetFlag("H", (value&0x0F+carry) > *target&0x0F)
	cpu.Registers.SetFlag("Z", (result&0xFF) == 0)
	cpu.Registers.SetFlag("N", true)

	*target = uint8(result & 0xFF)
}

func (cpu *CPU) SBC_A_r8(reg string) {

	A, _ := cpu.Registers.Get8bitRegister("A")
	value, _ := cpu.Registers.Get8bitRegister(reg)
	cpu._SBC_(&A, value)
}

func (cpu *CPU) SBC_A_HL() {

	hlPointer, _ := cpu.Registers.Get16bitRegister("HL")
	value := cpu.ReadAddr(hlPointer)
	A, _ := cpu.Registers.Get8bitRegister("A")
	cpu._SBC_(&A, value)
}

func (cpu *CPU) SBC_A_n8(value uint8) {

	A, _ := cpu.Registers.Get8bitRegister("A")
	cpu._SBC_(&A, value)
}

/* SUB */

func (cpu *CPU) _SUB_(target *uint8, value uint8) {

	result := int16(*target) - int16(value)

	cpu.Registers.SetFlag("C", value > *target)
	cpu.Registers.SetFlag("H", value&0x0F > *target&0x0F)
	cpu.Registers.SetFlag("Z", (result&0xFF) == 0)
	cpu.Registers.SetFlag("N", true)

	*target = uint8(result & 0xFF)
}

func (cpu *CPU) SUB_A_r8(reg string) {

	A, _ := cpu.Registers.Get8bitRegister("A")
	value, _ := cpu.Registers.Get8bitRegister(reg)
	cpu._SUB_(&A, value)
}

func (cpu *CPU) SUB_A_HL() {

	hlPointer, _ := cpu.Registers.Get16bitRegister("HL")
	value := cpu.ReadAddr(hlPointer)
	A, _ := cpu.Registers.Get8bitRegister("A")
	cpu._SUB_(&A, value)
}

func (cpu *CPU) SUB_A_n8(value uint8) {

	A, _ := cpu.Registers.Get8bitRegister("A")
	cpu._SUB_(&A, value)
}

/* XOR */

func (cpu *CPU) _XOR_(target *uint8, value uint8) {

	*target ^= value

	cpu.Registers.SetFlag("C", false)
	cpu.Registers.SetFlag("H", false)
	cpu.Registers.SetFlag("Z", (*target&0xFF) == 0)
	cpu.Registers.SetFlag("N", false)

}

func (cpu *CPU) XOR_A_r8(reg string) {

	A, _ := cpu.Registers.Get8bitRegister("A")
	value, _ := cpu.Registers.Get8bitRegister(reg)
	cpu._XOR_(&A, value)
}

func (cpu *CPU) XOR_A_HL(reg string) {

	A, _ := cpu.Registers.Get8bitRegister("A")
	hlPointer, _ := cpu.Registers.Get16bitRegister("HL")
	value := cpu.ReadAddr(hlPointer)
	cpu._XOR_(&A, value)
}

func (cpu *CPU) XOR_A_n8(value uint8) {

	A, _ := cpu.Registers.Get8bitRegister("A")
	cpu._XOR_(&A, value)
}
