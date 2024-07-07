package arg

import (
	"bytes"
	"errors"
	"fmt"
	"os/exec"
	"testing"
)

func TestArg_Cidr(t *testing.T) {
	const cmd = "go"
	const testProgram = "examples/Cidr/main.go"

	type Data struct {
		err error
		out []byte
	}

	addresses := map[string]Data{
		"192.168.1.0/24": {
			err: nil,
			out: []byte{
				27, 91, 48, 109, 49, 57, 50, 46, 49, 54, 56, 46, 49, 46, 48, 47, 50, 52,
			},
		},
		"192.168.1.0": {
			err: nil,
			out: []byte{
				27, 91, 48, 109, 105, 110, 118, 97, 108, 105, 100, 32,
				67, 73, 68, 82, 32, 97, 100, 100, 114, 101, 115, 115,
				58, 32, 49, 57, 50, 46, 49, 54, 56, 46, 49, 46, 48, 10,
			},
		},
		"2001:db8::/32": {
			err: nil,
			out: []byte{
				27, 91, 48, 109, 50, 48, 48, 49, 58, 100, 98, 56, 58, 58, 47, 51, 50,
			},
		},
		"256.256.256.256/32": {
			err: fmt.Errorf("exit status 1"),
			out: []byte{
				27, 91, 48, 109, 105, 110, 118, 97, 108, 105, 100, 32, 67, 73, 68,
				82, 32, 97, 100, 100, 114, 101, 115, 115, 58, 32, 50, 53, 54, 46,
				50, 53, 54, 46, 50, 53, 54, 46, 50, 53, 54, 47, 51, 50, 10,
			},
		},
		"": {
			err: nil,
			out: []byte{
				27, 91, 48, 109, 105, 110, 118, 97, 108, 105, 100, 32, 67,
				73, 68, 82, 32, 97, 100, 100, 114, 101, 115, 115, 58, 32, 10,
			},
		},
		"bad": {
			err: fmt.Errorf("exit status 1"),
			out: []byte{
				27, 91, 48, 109, 105, 110, 118, 97, 108, 105, 100, 32, 67, 73, 68,
				82, 32, 97, 100, 100, 114, 101, 115, 115, 58, 32, 98, 97, 100, 10,
			},
		},
	}

	for address, expected := range addresses {
		t.Run(fmt.Sprintf("arg.Cidr() %s [%v]", address, expected.err), func(t *testing.T) {

			args := []string{"run", testProgram, "-cidr", address}
			actual, err := exec.Command(cmd, args...).Output()
			if err != nil && errors.Is(err, expected.err) {
				t.Fatalf("Failed to run test program: %v", err)
			}
			if !bytes.Equal(actual, expected.out) {
				t.Fatalf("Failed (mismatch)\n"+
					"     got: '%s' (%v)\n"+
					"expected: '%s' (%v)\n",
					string(actual), actual, string(expected.out), expected.out)
			}
		})

	}
}
