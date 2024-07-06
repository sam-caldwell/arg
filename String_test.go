package arg

import (
	"bytes"
	"os/exec"
	"testing"
)

func TestArg_String(t *testing.T) {
	const cmd = "go"
	const testProgram = "examples/String.go"

	t.Run("arg.String() false", func(t *testing.T) {
		args := []string{"run", testProgram}
		actual, err := exec.Command(cmd, args...).Output()
		if err != nil {
			t.Fatalf("Failed to run test program: %v", err)
		}
		expected := append([]byte{27, 91, 48, 109}, []byte("value:5")...)
		if !bytes.Equal(actual, expected) {
			t.Fatalf("Failed\n"+
				"     got: '%v'\n"+
				"expected: '%v'\n", actual, expected)
		}
	})
}
