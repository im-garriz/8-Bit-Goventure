package z80_cpu

import (
	"testing"
)

type TestParams struct {
	name     string
	target   uint8
	value    uint8
	carry    bool
	res      uint8
	tg_flags map[string]uint8
}

func checkOP(cpu *CPU, tt *TestParams, t *testing.T) {

	if tt.target != tt.res {
		t.Errorf("RES: got %d, want %d", tt.target, tt.res)
	}

	z, _ := cpu.Registers.GetFlag("Z")
	n, _ := cpu.Registers.GetFlag("N")
	h, _ := cpu.Registers.GetFlag("H")
	c, _ := cpu.Registers.GetFlag("C")
	generated_flags := map[string]uint8{"z": z, "n": n, "h": h, "c": c}

	for _, flag := range []string{"z", "n", "h", "c"} {
		target_flag := tt.tg_flags[flag]
		generated_flag := generated_flags[flag]

		if target_flag != generated_flag {
			t.Errorf("%s: got %d, want %d", flag, generated_flag, target_flag)
		}
	}
}

func TestALL(t *testing.T) {
	// Defining the columns of the table

	cpu, err := GetCPU("../etc/snake.gb", "../etc/opcodes.json")
	if err != nil {
		t.Fatalf("Error in GetCPU:\n[E]: %s\n", err)
	}

	var tests = []TestParams{
		// the table itself
		{"ADC: 50+50+0=100", 50, 50, false, 100, map[string]uint8{"z": 0, "n": 0, "h": 0, "c": 0}},
		{"ADC: 50+50+1=101", 50, 50, true, 101, map[string]uint8{"z": 0, "n": 0, "h": 0, "c": 0}},
		{"ADC: 255+1+0=0 (overflow Z)", 255, 1, false, 0, map[string]uint8{"z": 1, "n": 0, "h": 1, "c": 1}},
		{"ADC: 255+1+1=0 (overflow)", 255, 1, true, 1, map[string]uint8{"z": 0, "n": 0, "h": 1, "c": 1}},
	}
	// The execution loop
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			cpu.Registers.SetFlag("C", tt.carry)
			cpu._ADC_(&tt.target, tt.value)

			checkOP(&cpu, &tt, t)
		})
	}

	tests = []TestParams{
		// the table itself
		{"ADD: 50+50=100", 50, 50, false, 100, map[string]uint8{"z": 0, "n": 0, "h": 0, "c": 0}},
		{"ADD: 25+150=175", 25, 150, false, 175, map[string]uint8{"z": 0, "n": 0, "h": 0, "c": 0}},
		{"ADD: 50+206=0", 50, 206, false, 0, map[string]uint8{"z": 1, "n": 0, "h": 1, "c": 1}},
		{"ADD: 4+1=5", 4, 1, false, 5, map[string]uint8{"z": 0, "n": 0, "h": 0, "c": 0}},
	}

	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			cpu._ADD_(&tt.target, tt.value)

			checkOP(&cpu, &tt, t)
		})
	}

	tests = []TestParams{
		// the table itself
		{"AND: 50&50=50", 50, 50, false, 50, map[string]uint8{"z": 0, "n": 0, "h": 1, "c": 0}},
		{"AND: 15&240=0", 15, 240, false, 0, map[string]uint8{"z": 1, "n": 0, "h": 1, "c": 0}},
	}

	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			cpu._AND_(&tt.target, tt.value)

			checkOP(&cpu, &tt, t)
		})
	}

	tests = []TestParams{
		// the table itself
		{"CP: 100-50", 100, 50, false, 100, map[string]uint8{"z": 0, "n": 1, "h": 0, "c": 0}},
		{"CP: 25-240=0", 25, 240, false, 25, map[string]uint8{"z": 0, "n": 1, "h": 0, "c": 1}},
		{"CP: 240-50=0", 240, 50, false, 240, map[string]uint8{"z": 0, "n": 1, "h": 1, "c": 0}},
		{"CP: 240-240=0", 240, 240, false, 240, map[string]uint8{"z": 1, "n": 1, "h": 0, "c": 0}},
	}

	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			cpu._CP_(&tt.target, tt.value)

			checkOP(&cpu, &tt, t)
		})
	}

	tests = []TestParams{
		// the table itself
		{"DEC: 100->99", 100, 0, false, 99, map[string]uint8{"z": 0, "n": 1, "h": 0, "c": 0}},
		{"DEC: 1->0", 1, 0, true, 0, map[string]uint8{"z": 1, "n": 1, "h": 1, "c": 1}},
		{"DEC: 0->255", 0, 0, false, 255, map[string]uint8{"z": 0, "n": 1, "h": 0, "c": 0}},
		{"DEC: 241->240", 241, 0, true, 240, map[string]uint8{"z": 0, "n": 1, "h": 1, "c": 1}},
	}

	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			cpu.Registers.SetFlag("C", tt.carry)
			cpu._DEC_(&tt.target)

			checkOP(&cpu, &tt, t)
		})
	}

	tests = []TestParams{
		// the table itself
		{"INC: 100->101", 100, 0, false, 101, map[string]uint8{"z": 0, "n": 0, "h": 0, "c": 0}},
		{"INC: 63->64", 63, 0, true, 64, map[string]uint8{"z": 0, "n": 0, "h": 1, "c": 1}},
		{"INC: 255->0", 255, 0, true, 0, map[string]uint8{"z": 1, "n": 0, "h": 1, "c": 1}},
	}

	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			cpu.Registers.SetFlag("C", tt.carry)
			cpu._INC_(&tt.target)

			checkOP(&cpu, &tt, t)
		})
	}

	tests = []TestParams{
		// the table itself
		{"OR: 100,50->118", 100, 50, false, 118, map[string]uint8{"z": 0, "n": 0, "h": 0, "c": 0}},
		{"OR: 0,0->0", 0, 0, true, 0, map[string]uint8{"z": 1, "n": 0, "h": 0, "c": 0}},
	}

	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			cpu.Registers.SetFlag("C", tt.carry)
			cpu._OR_(&tt.target, tt.value)

			checkOP(&cpu, &tt, t)
		})
	}

	tests = []TestParams{
		// the table itself
		{"SBC: 100,50,0->50", 100, 50, false, 50, map[string]uint8{"z": 0, "n": 1, "h": 0, "c": 0}},
		{"SBC: 100,50,1->50", 100, 50, true, 49, map[string]uint8{"z": 0, "n": 1, "h": 0, "c": 0}},
		{"SBC: 50,100,0->206", 50, 100, false, 206, map[string]uint8{"z": 0, "n": 1, "h": 1, "c": 1}},
		{"SBC: 100,100,0->0", 100, 100, false, 0, map[string]uint8{"z": 1, "n": 1, "h": 0, "c": 0}},
	}

	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			cpu.Registers.SetFlag("C", tt.carry)
			cpu._SBC_(&tt.target, tt.value)

			checkOP(&cpu, &tt, t)
		})
	}

	tests = []TestParams{
		// the table itself
		{"SUB: 100,50,0->50", 100, 50, false, 50, map[string]uint8{"z": 0, "n": 1, "h": 0, "c": 0}},
		{"SUB: 50,100,0->206", 50, 100, false, 206, map[string]uint8{"z": 0, "n": 1, "h": 1, "c": 1}},
		{"SUB: 100,100,0->0", 100, 100, false, 0, map[string]uint8{"z": 1, "n": 1, "h": 0, "c": 0}},
	}

	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			cpu.Registers.SetFlag("C", tt.carry)
			cpu._SUB_(&tt.target, tt.value)

			checkOP(&cpu, &tt, t)
		})
	}

	tests = []TestParams{
		// the table itself
		{"XOR: 255,50->205", 255, 50, false, 205, map[string]uint8{"z": 0, "n": 0, "h": 0, "c": 0}},
		{"XOR: 255,255->0", 255, 255, false, 0, map[string]uint8{"z": 1, "n": 0, "h": 0, "c": 0}},
	}

	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			cpu.Registers.SetFlag("C", tt.carry)
			cpu._XOR_(&tt.target, tt.value)

			checkOP(&cpu, &tt, t)
		})
	}
}
