package arg

import (
	"bytes"
	"os"
	"os/exec"
	"testing"
)

func TestArg_Filename(t *testing.T) {
	const cmd = "go"

	t.Run("arg.Filename():valid", func(t *testing.T) {
		const (
			testProgram = "examples/FileName_valid/main.go"
			testFile    = "/tmp/valid_file"
		)
		t.Cleanup(func() {
			_ = os.Remove(testFile)
		})
		if err := os.WriteFile(testFile, []byte("foo"), 0660); err != nil {
			t.Fatalf("setup error: %v", err)
		}

		args := []string{"run", testProgram}
		actual, err := exec.Command(cmd, args...).Output()
		if err != nil {
			t.Fatalf("Failed to run test program: %v", err)
		}
		expected := []byte{
			27, 91, 48, 109, 118, 97, 108, 117, 101, 58, 47, 116,
			109, 112, 47, 118, 97, 108, 105, 100, 95, 102, 105,
			108, 101, 46, 116, 120, 116}

		if !bytes.Equal(actual, expected) {
			t.Fatalf("Failed.\n"+
				"     got: '%v'\n"+
				"expected: '%v'\n", actual, expected)
		}
	})

	t.Run("arg.Filename():invalid", func(t *testing.T) {
		const (
			testProgram = "examples/FileName_invalid/main.go"
			testFile    = "/tmp/invalid_file"
		)
		t.Cleanup(func() {
			_ = os.Remove(testFile)
		})
		if err := os.WriteFile(testFile, []byte("foo"), 0660); err != nil {
			t.Fatalf("setup error: %v", err)
		}
		_ = os.Remove(testFile) // deleting the file we just created helps ensure it doesn't exist.
		args := []string{"run", testProgram}
		actual, err := exec.Command(cmd, args...).Output()
		if err == nil {
			t.Fatalf("expected error not found")
		}
		expected := []byte{}

		if !bytes.Equal(actual, expected) {
			t.Fatalf("Failed.\n"+
				"     got: '%v' (%s)\n"+
				"expected: '%v'\n", actual, string(actual), expected)
		}
	})

}
