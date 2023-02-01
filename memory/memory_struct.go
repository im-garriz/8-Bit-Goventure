package memory

type GBMemoryBank struct {
	bios                [256]uint8      // 256-byte BIOS
	rom                 [0x200000]uint8 // 2 MB ROM
	vram                [0x2000]uint8   // 8KB video RAM
	eram                [0x20000]uint8  // 128KB external RAM
	wram                [0x2000]uint8   // 8KB work RAM
	oam                 [0xA0]uint8     // 160-byte sprite attribute table
	hram                [0x7F]uint8     // 127-byte high RAM
	ier                 uint8           // interrupt enable register
	io                  [0x80]uint8     // 128-byte I/O ports
	interruptFlag       uint8           // Interrupt flag register
	interruptEnable     uint8           // Interrupt enable register
	tma                 uint8           // Timer modulo register
	tac                 uint8           // Timer control register
	divider             uint8           // Divider register
	tima                uint8           // Timer counter
	serialData          uint8           // Serial transfer data
	serialControl       uint8           // Serial transfer control
	joyp                uint8           // Joypad register
	wavePatternRAM      [0x10]uint8     // 16-byte wave pattern RAM
	soundControl        [0x10]uint8     // 16 Sound control registers
	soundMixer          uint8           // Sound mixer register
	soundOutputTerminal uint8           // Sound output terminal
	soundOnOff          uint8           // Sound on/off register
	waveRAM             [0x10]uint8     // 16-byte wave RAM
	lcdc                uint8           // LCD control register
	stat                uint8           // LCD status register
	scy                 uint8           // Scroll Y
	scx                 uint8           // Scroll X
	ly                  uint8           // LCDC Y-Coordinate
	lyc                 uint8           // LY compare
	dma                 uint8           // DMA transfer and start address
	bgp                 uint8           // Background palette data
	obp0                uint8           // Object palette 0 data
	obp1                uint8           // Object palette 1 data
	wy                  uint8           // Window Y position
	wx                  uint8           // Window X position
	vBlankInterrupt     uint8           // V-Blank interrupt flag
	lcdcInterrupt       uint8           // LCDC interrupt flag
	timerInterrupt      uint8           // Timer interrupt flag
	serialInterrupt     uint8           // Serial interrupt flag
	joypadInterrupt     uint8           // Joypad interrupt flag
}
