package arg

import (
	"bytes"
	"os/exec"
	"testing"
)

func TestArg_Choices(t *testing.T) {
	const cmd = "go"
	const testProgram = "examples/Choices/main.go"

	t.Run("arg.Choices() false", func(t *testing.T) {
		args := []string{"run", testProgram}
		actual, err := exec.Command(cmd, args...).Output()
		if err != nil {
			t.Fatalf("Failed to run test program: %v", err)
		}
		expected := []byte{27, 91, 48, 109, 118, 97, 108, 117, 101, 58, 97}
		if !bytes.Equal(actual, expected) {
			t.Fatalf("exit code should return no output.\n"+
				"     got: '%v'\n"+
				"expected: '%v'\n", actual, expected)
		}
	})

	t.Run("arg.Choices() false", func(t *testing.T) {
		args := []string{"run", testProgram, "-value", "b"}
		actual, err := exec.Command(cmd, args...).Output()
		if err != nil {
			t.Fatalf("Failed to run test program: %v", err)
		}
		expected := []byte{27, 91, 48, 109, 118, 97, 108, 117, 101, 58, 98}
		if !bytes.Equal(actual, expected) {
			t.Fatalf("exit code should return no output.\n"+
				"     got: '%v'\n"+
				"expected: '%v'\n", actual, expected)
		}
	})

	t.Run("arg.Choices() false", func(t *testing.T) {
		args := []string{"run", testProgram, "-value", "c"}
		actual, err := exec.Command(cmd, args...).Output()
		if err != nil {
			t.Fatalf("Failed to run test program: %v", err)
		}
		expected := []byte{27, 91, 48, 109, 118, 97, 108, 117, 101, 58, 99}
		if !bytes.Equal(actual, expected) {
			t.Fatalf("exit code should return no output.\n"+
				"     got: '%v'\n"+
				"expected: '%v'\n", actual, expected)
		}
	})

	t.Run("arg.Choices() false", func(t *testing.T) {
		args := []string{"run", testProgram, "-value", "d"}
		actual, err := exec.Command(cmd, args...).Output()
		if err != nil {
			t.Fatalf("Failed to run test program: %v", err)
		}
		expected := []byte{27, 91, 48, 109, 118, 97, 108, 117, 101, 58, 100}
		if !bytes.Equal(actual, expected) {
			t.Fatalf("Failed.\n"+
				"     got: '%v'\n"+
				"expected: '%v'\n", actual, expected)
		}
	})

	t.Run("arg.Choices() false", func(t *testing.T) {
		args := []string{"run", testProgram, "-value", "e"}
		actual, err := exec.Command(cmd, args...).Output()
		if err == nil {
			t.Fatalf("expected error")
		}
		expected := []byte{
			27, 91, 48, 109, 105, 110, 118, 97, 108,
			105, 100, 32, 105, 110, 112, 117, 116, 10}
		if !bytes.Equal(actual, expected) {
			t.Fatalf("Failed.\n"+
				"     got: '%v'\n"+
				"expected: '%v'\n", actual, expected)
		}
	})
}
