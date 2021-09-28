package plugin

import (
	"testing"
)

func TestGetFiles(t *testing.T) {
	tests := []struct {
		name   string
		input  string
		expect int
		hasErr bool
	}{
		{"recursively reads a directory", "../test", 2, false},
		{"error on unknown path", "./NOWHERE", 0, true},
	}

	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			files, err := getFiles(tt.input)
			if tt.hasErr && err == nil {
				t.Error("expected an error")
				t.Fail()
			}

			got := len(files)
			if tt.expect != got {
				t.Errorf("expected '%d' files, instead got '%d'", tt.expect, got)
				t.Fail()
			}
		})
	}
}
