package plugin

import (
	"io/ioutil"
	"testing"
)

const (
	testPath = "../test/"

	awsFile = "secret.yaml"
	errFile = "admin.yaml"

	awsVal = "AKIAS4ZG5BB8LHHG23O1"
	awsKey = "${AWS_ACCESS_KEY_ID}"

	original = `---
apiVersion: v1
kind: Secret
metadata:
  name: test-secret
type: Opaque
data:
  AWS_REGION: us-east-1
  AWS_ACCESS_KEY_ID: ${AWS_ACCESS_KEY_ID}
  AWS_SECRET_NAME: test/secret
`
	snapshot = `---
apiVersion: v1
kind: Secret
metadata:
  name: test-secret
type: Opaque
data:
  AWS_REGION: us-east-1
  AWS_ACCESS_KEY_ID: QUtJQVM0Wkc1QkI4TEhIRzIzTzE=
  AWS_SECRET_NAME: test/secret
`
)

func TestGetFiles(t *testing.T) {
	t.Parallel()

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

func TestFilter(t *testing.T) {
	t.Parallel()

	denyList := []string{"bar", "baz"}

	tests := []struct {
		name   string
		input  []string
		expect int
	}{
		{"filters unwanted: bar", []string{"alpha", "bar", "bravo"}, 2},
		{"filters unwanted: bar, baz", []string{"foo", "bar", "baz"}, 1},
	}

	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {

			list := filter(tt.input, isAllowed(denyList))
			got := len(list)

			if tt.expect != got {
				t.Errorf("expected '%d', instead got '%d'", tt.expect, got)
				t.Fail()
			}
		})
	}
}

func TestFindAndReplace(t *testing.T) {
	tests := []struct {
		name    string
		file    string
		find    string
		replace string
		hasErr  bool
	}{
		{"error ioutil.WriteFile", testPath + errFile, "NOTHING", "something", true},
		{"success", testPath + awsFile, awsKey, awsVal, false},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			err := findAndReplace([]string{tt.file}, tt.find, tt.replace)
			if tt.hasErr && err == nil {
				t.Error("expected an error")
				t.Fail()
			}

			if !tt.hasErr {
				defer cleanup()
				file, _ := ioutil.ReadFile(tt.file)
				if string(file) != snapshot {
					t.Error("expected to interpolate " + awsFile)
					t.Fail()
				}
			}
		})
	}
}

// cleanup reverts the test file back
func cleanup() error {
	return ioutil.WriteFile(testPath+awsFile, []byte(original), 0644)
}
