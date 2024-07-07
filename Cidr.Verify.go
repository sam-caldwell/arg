package arg

import "net"

// Verify - Ensure the given string is a valid CIDR.
func (b *Cidr) Verify() (err error) {

	_, _, err = net.ParseCIDR(*b.value)

	return err

}
