package z80_cpu

import (
	"testing"
)

func TestSet16bitRegister(t *testing.T) {
	registers := GetCPURegisters()
	registers.Set16bitRegister("AF", 0xFF00)
	registers.Set16bitRegister("BC", 0xFF00)
	registers.Set16bitRegister("DE", 0xFF00)
	registers.Set16bitRegister("HL", 0xFF00)
	registers.Set16bitRegister("PC", 0xFF00)
	registers.Set16bitRegister("SP", 0xFF00)

	if registers.values["AF"] != 0xFF00 {
		t.Errorf("Expected 0xFF00, got %X", registers.values["AF"])
	}
	if registers.values["BC"] != 0xFF00 {
		t.Errorf("Expected 0xFF00, got %X", registers.values["BC"])
	}
	if registers.values["DE"] != 0xFF00 {
		t.Errorf("Expected 0xFF00, got %X", registers.values["DE"])
	}
	if registers.values["HL"] != 0xFF00 {
		t.Errorf("Expected 0xFF00, got %X", registers.values["HL"])
	}
	if registers.values["PC"] != 0xFF00 {
		t.Errorf("Expected 0xFF00, got %X", registers.values["PC"])
	}
	if registers.values["SP"] != 0xFF00 {
		t.Errorf("Expected 0xFF00, got %X", registers.values["SP"])
	}
}

func TestSet8bitRegister(t *testing.T) {
	registers := GetCPURegisters()

	registers.Set16bitRegister("AF", 0xFFFF)
	registers.Set16bitRegister("BC", 0xFFFF)
	registers.Set16bitRegister("DE", 0xFFFF)
	registers.Set16bitRegister("HL", 0xFFFF)
	registers.Set16bitRegister("PC", 0xFFFF)
	registers.Set16bitRegister("SP", 0xFFFF)

	registers.Set8bitRegister("A", 0x05)
	registers.Set8bitRegister("B", 0x05)
	registers.Set8bitRegister("D", 0x05)
	registers.Set8bitRegister("H", 0x05)
	registers.Set8bitRegister("C", 0x05)
	registers.Set8bitRegister("E", 0x05)
	registers.Set8bitRegister("L", 0x05)

	if registers.values["AF"] != 0x05FF {
		t.Errorf("[AF] Expected 0x05FF, got %X", registers.values["AF"])
	}
	if registers.values["BC"] != 0x0505 {
		t.Errorf("[BC] Expected 0x0505, got %X", registers.values["BC"])
	}
	if registers.values["DE"] != 0x0505 {
		t.Errorf("[DE] Expected 0x0505, got %X", registers.values["DE"])
	}
	if registers.values["HL"] != 0x0505 {
		t.Errorf("[HL] Expected 0x0505, got %X", registers.values["HL"])
	}
}

func TestSetFlag(t *testing.T) {
	registers := GetCPURegisters()
	registers.SetFlag("c", true)
	registers.SetFlag("h", true)
	registers.SetFlag("n", true)
	registers.SetFlag("z", true)

	if registers.values["AF"] != 0xF0 {
		t.Errorf("Expected 0xF0, got %X", registers.values["AF"])
	}
}

func TestGet16bitRegister(t *testing.T) {
	registers := GetCPURegisters()
	registers.Set16bitRegister("AF", 0xFF00)
	registers.Set16bitRegister("BC", 0xFF00)
	registers.Set16bitRegister("DE", 0xFF00)
	registers.Set16bitRegister("HL", 0xFF00)
	registers.Set16bitRegister("PC", 0xFF00)
	registers.Set16bitRegister("SP", 0xFF00)

	if val, err := registers.Get16bitRegister("AF"); err != nil {
		t.Errorf("Expected 0xFF00, got %X", val)
	}
	if val, err := registers.Get16bitRegister("BC"); err != nil {
		t.Errorf("Expected 0xFF00, got %X", val)
	}
	if val, err := registers.Get16bitRegister("DE"); err != nil {
		t.Errorf("Expected 0xFF00, got %X", val)
	}
	if val, err := registers.Get16bitRegister("HL"); err != nil {
		t.Errorf("Expected 0xFF00, got %X", val)
	}
	if val, err := registers.Get16bitRegister("PC"); err != nil {
		t.Errorf("Expected 0xFF00, got %X", val)
	}
	if val, err := registers.Get16bitRegister("SP"); err != nil {
		t.Errorf("Expected 0xFF00, got %X", val)
	}
}

func TestGet8bitRegister(t *testing.T) {
	registers := GetCPURegisters()
	registers.Set8bitRegister("A", 0xFF)
	registers.Set8bitRegister("B", 0xFF)
	registers.Set8bitRegister("D", 0xFF)
	registers.Set8bitRegister("H", 0xFF)
	registers.Set8bitRegister("C", 0xFF)
	registers.Set8bitRegister("E", 0xFF)
	registers.Set8bitRegister("L", 0xFF)

	if val, err := registers.Get8bitRegister("A"); err != nil {
		t.Errorf("Expected 0xFF, got %X", val)
	}
	if val, err := registers.Get8bitRegister("B"); err != nil {
		t.Errorf("Expected 0xFF, got %X", val)
	}
	if val, err := registers.Get8bitRegister("D"); err != nil {
		t.Errorf("Expected 0xFF, got %X", val)
	}
	if val, err := registers.Get8bitRegister("H"); err != nil {
		t.Errorf("Expected 0xFF, got %X", val)
	}
	if val, err := registers.Get8bitRegister("C"); err != nil {
		t.Errorf("Expected 0xFF, got %X", val)
	}
	if val, err := registers.Get8bitRegister("E"); err != nil {
		t.Errorf("Expected 0xFF, got %X", val)
	}
	if val, err := registers.Get8bitRegister("L"); err != nil {
		t.Errorf("Expected 0xFF, got %X", val)
	}
}

func TestGetFlag(t *testing.T) {
	registers := GetCPURegisters()
	registers.SetFlag("c", true)
	registers.SetFlag("h", true)
	registers.SetFlag("n", true)
	registers.SetFlag("z", true)

	if val, err := registers.GetFlag("c"); err != nil {
		t.Errorf("Expected 1, got %X", val)
	}
	if val, err := registers.GetFlag("h"); err != nil {
		t.Errorf("Expected 1, got %X", val)
	}
	if val, err := registers.GetFlag("n"); err != nil {
		t.Errorf("Expected 1, got %X", val)
	}
	if val, err := registers.GetFlag("z"); err != nil {
		t.Errorf("Expected 1, got %X", val)
	}
}
