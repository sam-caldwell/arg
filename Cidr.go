package arg

// Cidr - Represents an IPv4 CIDR address
//
//	CIDR (Classless Inter-Domain Routing) is a method for allocating IP addresses and IP
//	routing that replaces the older system based on classful networking. It uses variable-length
//	subnet masking (VLSM) to create more flexible and efficient IP address allocation, represented
//	by an IP address followed by a slash and a prefix length (e.g., 192.168.1.0/24).
type Cidr struct {
	value *string
}
