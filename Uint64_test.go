package arg

import (
	"bytes"
	"os/exec"
	"testing"
)

func TestArg_Uint64(t *testing.T) {
	const cmd = "go"
	const testProgram = "examples/Uint64/main.go"

	t.Run("arg.Uint64() false", func(t *testing.T) {
		args := []string{"run", testProgram}
		actual, err := exec.Command(cmd, args...).Output()
		if err != nil {
			t.Fatalf("Failed to run test program: %v", err)
		}
		expected := []byte{27, 91, 48, 109, 118, 97, 108, 117, 101, 58, 53, 10}
		if !bytes.Equal(actual, expected) {
			t.Fatalf("Failed\n"+
				"     got: '%v'\n"+
				"expected: '%v'\n", actual, expected)
		}
	})

}
