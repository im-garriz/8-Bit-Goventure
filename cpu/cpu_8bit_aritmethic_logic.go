package cpu

func (cpu *CPU) ADC(target *uint8, value uint8) {

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
	cpu.ADC(&A, value)
}

func (cpu *CPU) ADC_A_HL() {

	hlPointer, _ := cpu.Registers.Get16bitRegister("HL")
	value := cpu.Memory.Read(hlPointer)
	A, _ := cpu.Registers.Get8bitRegister("A")
	cpu.ADC(&A, value)
}

func (cpu *CPU) ADC_A_n8(value uint8) {

	A, _ := cpu.Registers.Get8bitRegister("A")
	cpu.ADC(&A, value)
}

func (cpu *CPU) ADD(target *uint8, value uint8) {

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
	cpu.ADD(&A, value)
}

func (cpu *CPU) ADD_A_HL() {
	A, _ := cpu.Registers.Get8bitRegister("A")
	hlPointer, _ := cpu.Registers.Get16bitRegister("HL")
	value := cpu.Memory.Read(hlPointer)
	cpu.ADD(&A, value)
}

func (cpu *CPU) ADD_A_n8(value uint8) {
	A, _ := cpu.Registers.Get8bitRegister("A")
	cpu.ADD(&A, value)
}

func (cpu *CPU) AND(target *uint8, value uint8) {

	*target &= value

	cpu.Registers.SetFlag("C", false)
	cpu.Registers.SetFlag("H", true)
	cpu.Registers.SetFlag("Z", (*target&0xFF) == 0)
	cpu.Registers.SetFlag("N", false)

}

func (cpu *CPU) AND_A_r8(reg string) {

	A, _ := cpu.Registers.Get8bitRegister("A")
	value, _ := cpu.Registers.Get8bitRegister(reg)
	cpu.AND(&A, value)
}

func (cpu *CPU) AND_A_HL(reg string) {

	A, _ := cpu.Registers.Get8bitRegister("A")
	hlPointer, _ := cpu.Registers.Get16bitRegister("HL")
	value := cpu.Memory.Read(hlPointer)
	cpu.AND(&A, value)
}

func (cpu *CPU) AND_A_n8(value uint8) {

	A, _ := cpu.Registers.Get8bitRegister("A")
	cpu.AND(&A, value)
}
