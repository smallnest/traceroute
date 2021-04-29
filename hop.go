package traceroute

import (
	"net"
	"time"
)

// Hop type
type Hop struct {
	Success     bool
	Address     net.Addr
	Host        string
	N           int
	ElapsedTime time.Duration
	TTL         int
}

func (hop *Hop) AddressString() string {
	return hop.Address.String()
}

func (hop *Hop) String() string {
	hostOrAddr := hop.AddressString()
	if hop.Host != "" {
		hostOrAddr = hop.Host
	}
	return hostOrAddr
}
