package arg

import (
	"bytes"
	"os/exec"
	"testing"
)

func TestArg_Debug(t *testing.T) {
	const cmd = "go"
	const testProgram = "examples/debug.go"

	t.Run("arg.Debug()", func(t *testing.T) {
		args := []string{"run", testProgram}
		actual, err := exec.Command(cmd, args...).Output()
		if err != nil {
			t.Fatalf("Failed to run test program (no debug): %v", err)
		}
		expected := []byte{27, 91, 48, 109, 100, 101, 98, 117, 103, 58, 102, 97, 108, 115, 101, 10}
		if !bytes.Equal(actual, expected) {
			t.Fatalf("Failed.\n"+
				"     got: '%v'\n"+
				"expected: '%v'\n", actual, expected)
		}
	})

	t.Run("arg.Debug()", func(t *testing.T) {
		args := []string{"run", testProgram, "-debug"}
		actual, err := exec.Command(cmd, args...).Output()
		if err != nil {
			t.Fatalf("Failed to run test program (has debug): %v", err)
		}
		expected := []byte{27, 91, 48, 109, 100, 101, 98, 117, 103, 58, 116, 114, 117, 101, 10}
		if !bytes.Equal(actual, expected) {
			t.Fatalf("Failed.\n"+
				"     got: '%v'\n"+
				"expected: '%v'\n", actual, expected)
		}
	})
}
