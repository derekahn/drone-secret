package plugin

import (
	"context"
	"testing"
)

func TestPlugin(t *testing.T) {
	secret := map[string]string{awsKey: awsVal}

	tests := []struct {
		name   string
		input  Args
		hasErr bool
	}{
		{
			name: "successful execution",
			input: Args{
				Directory: "../test",
				DenyList:  []string{errFile},
				Secrets:   secret,
			},
			hasErr: false,
		},
		{
			name: "errors on file with elevated permissions",
			input: Args{
				Directory: "../test",
				DenyList:  []string{awsFile},
				Secrets:   secret,
			},
			hasErr: true,
		},
		{
			name: "errors on non-existent directory",
			input: Args{
				Directory: "./NOWHERE",
				Secrets:   secret,
			},
			hasErr: true,
		},
	}

	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			err := Exec(context.Background(), tt.input)
			if tt.hasErr && err == nil {
				t.Error("expected an error")
				t.Fail()
			}
			if !tt.hasErr {
				cleanup()
			}
		})
	}
}
