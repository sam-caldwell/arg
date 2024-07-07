package arg

import (
	"bytes"
	"os/exec"
	"testing"
)

func TestArg_Int(t *testing.T) {
	const cmd = "go"
	const testProgram = "examples/Int/main.go"

	t.Run("arg.Int() false", func(t *testing.T) {
		args := []string{"run", testProgram}
		actual, err := exec.Command(cmd, args...).Output()
		if err != nil {
			t.Fatalf("Failed to run test program: %v", err)
		}
		expected := []byte{
			27, 91, 48, 109, 118, 97, 108, 117, 101, 58, 48, 10,
		}
		if !bytes.Equal(actual, expected) {
			t.Fatalf("Failed (mismatch)\n"+
				"     got: '%s' (%v)\n"+
				"expected: '%s' (%v)\n",
				string(actual), actual, string(expected), expected)
		}
	})

}
