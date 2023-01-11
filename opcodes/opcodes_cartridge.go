package opcodes

import (
	"bytes"
	"encoding/binary"
	"fmt"
	"io/ioutil"
)

type GameBoyROM struct {
	ROM    []byte
	Header ROMHeader
}

func (gbr *GameBoyROM) LoadROM(filepath string) error {
	rom, err := ioutil.ReadFile(filepath)
	if err != nil {
		return err
	}
	gbr.ROM = rom

	gbr.Header = ROMHeader{}
	binary.Read(bytes.NewReader(rom[0x100:]), binary.LittleEndian, &gbr.Header)

	fmt.Printf("Title: %s\n", gbr.Header.Title)
	fmt.Printf("Header Checksum: %d\n", gbr.Header.HeaderChecksum)
	fmt.Printf("Global Checksum: %d\n", binary.LittleEndian.Uint16(gbr.Header.GlobalChecksum[:]))

	return nil
}

type ROMHeader struct {
	EntryPoint   [4]byte
	NintendoLogo [48]byte
	Title        [15]byte
	//ManufacturerCode [4]byte
	CGBFlag         byte
	NewLicenseeCode [2]byte
	SGBFlag         byte
	CartridgeType   byte
	ROMSize         byte
	RAMSize         byte
	DestinationCode byte
	OldLicenseeCode byte
	MaskROMVersion  byte
	HeaderChecksum  byte
	GlobalChecksum  [2]byte
}
