package disassembler

import (
	"encoding/binary"
	"fmt"
	"testing"
)

func TestDissasemblerHeader(t *testing.T) {
	// Defining the columns of the table
	var tests = []struct {
		name  string
		input string
		want  []string
	}{
		// the table itself
		{"Snake Header", "../../etc/snake.gb", []string{"Yvar's GB Snake", "66", "51166"}},
	}
	// The execution loop
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			gbr, err := LoadROM(tt.input)

			if err != nil {
				t.Errorf("Error loading ROM %s", tt.input)
			}

			title := fmt.Sprintf("%s", gbr.Header.Title)
			headerChecksum := fmt.Sprintf("%d", gbr.Header.HeaderChecksum)
			globalChecksum := fmt.Sprintf("%d", binary.LittleEndian.Uint16(gbr.Header.GlobalChecksum[:]))

			if title != tt.want[0] {
				t.Errorf("Title: got %s, want %s", title, tt.want[0])
			}

			if headerChecksum != tt.want[1] {
				t.Errorf("Header checksum: got %s, want %s", headerChecksum, tt.want[1])
			}

			if globalChecksum != tt.want[2] {
				t.Errorf("Global checksum: got %s, want %s", globalChecksum, tt.want[2])
			}
		})
	}
}
