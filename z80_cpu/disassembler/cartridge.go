// Package disassembler provides functionality for disassembling GameBoy emulator ROMs.
package disassembler

import (
	"bytes"
	"encoding/binary"
	"os"
)

// ROMHeader represents the header information of a GameBoy ROM.
type ROMHeader struct {
	EntryPoint   [4]byte  // The address of the entry point.
	NintendoLogo [48]byte // Nintendo logo data.
	Title        [15]byte // Title of the game.
	//ManufacturerCode [4]byte
	CGBFlag         byte    // Flag indicating CGB features.
	NewLicenseeCode [2]byte // New licensee code.
	SGBFlag         byte    // Flag indicating SGB features.
	CartridgeType   byte    // Type of cartridge.
	ROMSize         byte    // Size of the ROM.
	RAMSize         byte    // Size of the RAM.
	DestinationCode byte    // Destination code.
	OldLicenseeCode byte    // Old licensee code.
	MaskROMVersion  byte    // Version of the mask ROM.
	HeaderChecksum  byte    // Checksum for the header.
	GlobalChecksum  [2]byte // Global checksum.
}

// GameBoyROM represents a GameBoy ROM with its header information.
type GameBoyROM struct {
	ROM    []byte    // Actual ROM data.
	Header ROMHeader // Header information.
}

// LoadROM reads a GameBoy ROM file, extracts header information, and returns a GameBoyROM structure.
// It takes the file path as an argument and returns a pointer to a GameBoyROM and an error, if any.
func LoadROM(filepath string) (*GameBoyROM, error) {

	gbr := GameBoyROM{}

	// Read the entire ROM file.
	rom, err := os.ReadFile(filepath)
	if err != nil {
		return nil, err
	}
	gbr.ROM = rom

	// Extract header information from the ROM.
	gbr.Header = ROMHeader{}
	binary.Read(bytes.NewReader(rom[0x100:]), binary.LittleEndian, &gbr.Header)

	// // Print some information for testing purposes.
	// fmt.Printf("Title: %s\n", gbr.Header.Title)
	// fmt.Printf("Header Checksum: %d\n", gbr.Header.HeaderChecksum)
	// fmt.Printf("Global Checksum: %d\n", binary.LittleEndian.Uint16(gbr.Header.GlobalChecksum[:]))

	return &gbr, nil
}
