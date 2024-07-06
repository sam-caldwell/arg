package arg

import (
	"bytes"
	"os/exec"
	"testing"
)

func TestArg_Bool(t *testing.T) {
	const cmd = "go"
	const testProgram = "examples/bool.go"

	t.Run("arg.Bool() false", func(t *testing.T) {
		args := []string{"run", testProgram}
		actual, err := exec.Command(cmd, args...).Output()
		if err != nil {
			t.Fatalf("Failed to run test program: %v", err)
		}
		expected := append([]byte{27, 91, 48, 109}, []byte("value:false")...)
		if !bytes.Equal(actual, expected) {
			t.Fatalf("Failed\n"+
				"     got: '%v'\n"+
				"expected: '%v'\n", actual, expected)
		}
	})
	t.Run("arg.Bool() true", func(t *testing.T) {
		args := []string{"run", testProgram, "-value"}
		actual, err := exec.Command(cmd, args...).Output()
		if err != nil {
			t.Fatalf("Failed to run test program: %v", err)
		}
		expected := append([]byte{27, 91, 48, 109}, []byte("value:true")...)
		if !bytes.Equal(actual, expected) {
			t.Fatalf("Failed\n"+
				"     got: '%v'\n"+
				"expected: '%v'\n", actual, expected)
		}
	})
}
