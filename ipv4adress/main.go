package main

import (
	"fmt"
)

func main() {
	ipNum := uint32(2149583361) // Example number
	ipAddress := Int32ToIp(ipNum)
	fmt.Println("The IPv4 address is:", ipAddress)
}

// Int32ToIP converts an unsigned 32-bit integer to its IPv4 address representation.
func Int32ToIp(n uint32) string {
	// Extract each octet by shifting where necessary and mask.
	octet1 := (n >> 24) & 0xFF
	octet2 := (n >> 16) & 0xFF
	octet3 := (n >> 8) & 0xFF
	octet4 := n & 0xFF

	// Format the IPv4 address
	ipAddress := fmt.Sprintf("%d.%d.%d.%d", octet1, octet2, octet3, octet4)
	return ipAddress
}