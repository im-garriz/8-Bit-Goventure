package system

 import (
	"main/system/z80_cpu"
	"main/system/cartridge"
 )

type GameBoySystem struct {
	Cartridge *cartridge.GameBoyROM
	CPU  *z80_cpu.CPU
}

func (system *GameBoySystem) LoadCartridge(cartridgeFile string) error {
	rom, err := cartridge.LoadROM(cartridgeFile)
	if err != nil {
		return err
	}

	system.Cartridge = rom
	
	return err
}

func (system *GameBoySystem) StartCPU() error {

	cpu, err := z80_cpu.GetCPU(system.Cartridge)
	if err != nil {
		return err
	}

	system.CPU = cpu
	return err
}

func LoadSystem() GameBoySystem {

	system := GameBoySystem{}

	return system
}