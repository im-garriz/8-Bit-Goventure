package cpu

import "fmt"

/*

https://gbdev.io/pandocs/Memory_Map.html

0000 - 3FFF: 16KB ROM bank 00, From cartridge, usually a fixed bank (C)
	* 0000 - 00FF (255 bytes) Restart and Interrupt Vectors (BIOS)
	* 0100 - 014F (80 bytes) Cartridge Header Area
	* 0150 - 3FFF (~16KB) Cartridge ROM - Bank 0 (fixed)
4000 - 7FFF: 16KB ROM Bank 01~NN, From cartridge, switchable bank via MBC (if any)
8000 - 9FFF: 8KB Video RAM (VRAM), Only bank 0 in Non-CGB mode, Switchable bank 0/1 in CGB mode
	* 8000 - 97FF (6KB) Character RAM (VRAM)
	* 9800 - 9BFF (1KB) BG Map Data 1 (VRAM)
	* 9C00 - 9FFF (1KB) BG Map Data 2 (VRAM)
A000 - BFFF: 8KB External Cartridge RAM, In cartridge, switchable bank if any (C)
C000 - CFFF: 4KB Work RAM (WRAM) bank 0
D000 - DFFF: 4KB Work RAM (WRAM) bank 1~N, Only bank 1 in Non-CGB mode, Switchable bank 1~7 in CGB mode
E000 - FDFF: Mirror of C000~DDFF (ECHO RAM), Typically not used
FE00 - FE9F: Sprite attribute table (OAM) Object Attribute Memory
FEA0 - FEFF: Not Usable
FF00 - FF7F: I/O Registers
FF80 - FFFE: High RAM (HRAM)
FFFF - FFFF: Interrupts Enable Register (IE)

*/

func initCPUMemory() GBMemory {
	// TODO
	return GBMemory{}
}

type GBMemory struct {
	BIOS [0x100]uint8 // 256-byte BIOS
	// Cartridge 32KB
	VRAM [0x2000]uint8 // 8KB video RAM 8000 - 9FFF
	// Cartridge 8KB
	WRAM [0x2000]uint8 // 8KB work RAM
	OAM  [0xA0]uint8   // 160-byte sprite attribute table
	IO   [0x80]uint8   // 128-byte I/O ports
	HRAM [0x7F]uint8   // 127-byte high RAM
	IER  uint8         // interrupt enable register
}

func (cpu *CPU) ReadAddr(addr uint16) uint8 {
	switch {
	case addr <= 0xFF: // TODO boot_room_active()
		return cpu.memory.BIOS[addr]
	case addr <= 0x7FFF:
		return cpu.Disassembler.Cartridge.ROM[addr]
	case addr <= 0x9FFF:
		return cpu.memory.VRAM[addr-0x8000]
	case addr <= 0xBFFF:
		return cpu.Disassembler.Cartridge.ROM[addr] // TODO check length of cartridge
	case addr <= 0xDFFF:
		return cpu.memory.WRAM[addr-0xC000]
	case addr <= 0xFDFF:
		return cpu.memory.WRAM[addr-0xC000-0x2000]
	case addr <= 0xFE9F:
		return cpu.memory.OAM[addr-0xFE00]
	case addr <= 0xFEFF:
		fmt.Printf("Attemting to read from not usable memory address 0x%x\n", addr)
		return 0xFF
	case addr <= 0xFF7F:
		// TODO
		//return cpu.ReadIO(addr)
		return 0x00
	case addr <= 0xFFFE:
		return cpu.memory.HRAM[addr-0xFF80]
	case addr == 0xFFFF:
		return cpu.memory.IER
	}

	fmt.Printf("Attemting to read from not unmapped memory address 0x%x\n", addr)
	return 0xFF
}

func (cpu *CPU) WriteAddr(addr uint16, val uint8) {
	switch {
	case addr <= 0x7FFF:
		cpu.Disassembler.Cartridge.ROM[addr] = val
		return
	case addr <= 0x9FFF:
		cpu.memory.VRAM[addr-0x8000] = val
		return
	case addr <= 0xBFFF:
		cpu.Disassembler.Cartridge.ROM[addr] = val // TODO check length of cartridge
		return
	case addr <= 0xDFFF:
		cpu.memory.WRAM[addr-0xC000] = val
		return
	case addr <= 0xFDFF:
		cpu.memory.WRAM[addr-0xC000-0x2000] = val
		return
	case addr <= 0xFE9F:
		cpu.memory.OAM[addr-0xFE00] = val
		return
	case addr <= 0xFEFF:
		fmt.Printf("Attemting to write into not usable memory address 0x%x\n", addr)
		return
	case addr <= 0xFF7F:
		// TODO
		//return cpu.WriteIO(addr, val)
		return
	case addr <= 0xFFFE:
		cpu.memory.HRAM[addr-0xFF80] = val
		return
	case addr == 0xFFFF:
		cpu.memory.IER = val
		return
	}

	fmt.Printf("Attemting to write into unmapped memory address 0x%x\n", addr)
	return
}
