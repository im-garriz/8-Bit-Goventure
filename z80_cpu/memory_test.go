package z80_cpu

import (
	"testing"
)

func TestReadAddr(t *testing.T) {
	cpu := CPU{}

	cpu.memory.BIOS[0] = 0x1F
	cpu.memory.BIOS[0xFF] = 0xFF

	cpu.memory.VRAM[0] = 0x01
	cpu.memory.VRAM[0x1FFF] = 0x02

	cpu.memory.WRAM[0] = 0x03
	cpu.memory.WRAM[0x1FFF] = 0x04
	cpu.memory.WRAM[0xFFF] = 0x44

	cpu.memory.OAM[0] = 0x07
	cpu.memory.OAM[0x9F] = 0x08

	cpu.memory.HRAM[0] = 0x09
	cpu.memory.HRAM[0x7E] = 0x0A

	// BIOS
	if cpu.ReadAddr(0x00) != 0x1F {
		t.Errorf("Expected 0x1F, got %x", cpu.ReadAddr(0x00))
	}
	if cpu.ReadAddr(0xFF) != 0xFF {
		t.Errorf("Expected 0xFF, got %x", cpu.ReadAddr(0xFF))
	}

	// VRAM
	if cpu.ReadAddr(0x8000) != 0x01 {
		t.Errorf("Expected 0x01, got %x", cpu.ReadAddr(0x8000))
	}
	if cpu.ReadAddr(0x9FFF) != 0x02 {
		t.Errorf("Expected 0x02, got %x", cpu.ReadAddr(0x9FFF))
	}

	// WRAM
	if cpu.ReadAddr(0xC000) != 0x03 {
		t.Errorf("Expected 0x03, got %x", cpu.ReadAddr(0xC000))
	}
	if cpu.ReadAddr(0xDFFF) != 0x04 {
		t.Errorf("Expected 0x04, got %x", cpu.ReadAddr(0xDFFF))
	}
	if cpu.ReadAddr(0xE000) != 0x03 {
		t.Errorf("Expected 0x03, got %x", cpu.ReadAddr(0xE000))
	}
	if cpu.ReadAddr(0xEFFF) != 0x44 {
		t.Errorf("Expected 0x44, got %x", cpu.ReadAddr(0xEFFF))
	}

	// OAM
	if cpu.ReadAddr(0xFE00) != 0x07 {
		t.Errorf("Expected 0x07, got %x", cpu.ReadAddr(0xFE00))
	}
	if cpu.ReadAddr(0xFE9F) != 0x08 {
		t.Errorf("Expected 0x08, got %x", cpu.ReadAddr(0xFE9F))
	}

	// HRAM
	if cpu.ReadAddr(0xFF80) != 0x09 {
		t.Errorf("Expected 0x09, got %x", cpu.ReadAddr(0xFF80))
	}
	if cpu.ReadAddr(0xFFFE) != 0x0A {
		t.Errorf("Expected 0x0A, got %x", cpu.ReadAddr(0xFFFE))
	}
}

func TestWriteAddr(t *testing.T) {
	cpu := CPU{}

	cpu.WriteAddr(0x8000, 0x01)
	cpu.WriteAddr(0x9FFF, 0x02)

	cpu.WriteAddr(0xC000, 0x03)
	cpu.WriteAddr(0xDFFF, 0x04)
	cpu.WriteAddr(0xEFFF, 0x44)

	cpu.WriteAddr(0xFE00, 0x07)
	cpu.WriteAddr(0xFE9F, 0x08)

	cpu.WriteAddr(0xFF80, 0x09)
	cpu.WriteAddr(0xFFFE, 0x0A)

	// VRAM
	if cpu.ReadAddr(0x8000) != 0x01 {
		t.Errorf("Expected 0x01, got %x", cpu.ReadAddr(0x8000))
	}
	if cpu.ReadAddr(0x9FFF) != 0x02 {
		t.Errorf("Expected 0x02, got %x", cpu.ReadAddr(0x9FFF))
	}

	// WRAM
	if cpu.ReadAddr(0xC000) != 0x03 {
		t.Errorf("Expected 0x03, got %x", cpu.ReadAddr(0xC000))
	}
	if cpu.ReadAddr(0xDFFF) != 0x04 {
		t.Errorf("Expected 0x04, got %x", cpu.ReadAddr(0xDFFF))
	}
	if cpu.ReadAddr(0xE000) != 0x03 {
		t.Errorf("Expected 0x03, got %x", cpu.ReadAddr(0xE000))
	}
	if cpu.ReadAddr(0xEFFF) != 0x44 {
		t.Errorf("Expected 0x44, got %x", cpu.ReadAddr(0xEFFF))
	}

	// OAM
	if cpu.ReadAddr(0xFE00) != 0x07 {
		t.Errorf("Expected 0x07, got %x", cpu.ReadAddr(0xFE00))
	}
	if cpu.ReadAddr(0xFE9F) != 0x08 {
		t.Errorf("Expected 0x08, got %x", cpu.ReadAddr(0xFE9F))
	}

	// HRAM
	if cpu.ReadAddr(0xFF80) != 0x09 {
		t.Errorf("Expected 0x09, got %x", cpu.ReadAddr(0xFF80))
	}
	if cpu.ReadAddr(0xFFFE) != 0x0A {
		t.Errorf("Expected 0x0A, got %x", cpu.ReadAddr(0xFFFE))
	}

}
