package traceroute

const (
	DEFAULT_PORT        = 33434
	DEFAULT_MAX_HOPS    = 64
	DEFAULT_FIRST_HOP   = 1
	DEFAULT_TIMEOUT_MS  = 500
	DEFAULT_NQUERIES    = 1
	DEFAULT_PACKET_SIZE = 0
)

// DefaultOption provides a default opt.
var DefaultOption = &Option{
	port:         DEFAULT_PORT,
	maxHops:      DEFAULT_MAX_HOPS,
	firstHop:     DEFAULT_FIRST_HOP,
	timeoutMs:    DEFAULT_TIMEOUT_MS,
	nqueries:     DEFAULT_NQUERIES,
	packetSize:   DEFAULT_PACKET_SIZE,
	privileged:   true,
	fixedDstPort: false,
}

// TracrouteOption type
type Option struct {
	port         int
	maxHops      int
	firstHop     int
	timeoutMs    int
	nqueries     int
	packetSize   int
	resolveHost  bool
	privileged   bool
	fixedDstPort bool
}

func (opt *Option) Port() int {
	if opt.port == 0 {
		opt.port = DEFAULT_PORT
	}
	return opt.port
}

func (opt *Option) SetPort(port int) {
	opt.port = port
}

func (opt *Option) MaxHops() int {
	if opt.maxHops == 0 {
		opt.maxHops = DEFAULT_MAX_HOPS
	}
	return opt.maxHops
}

func (opt *Option) SetMaxHops(maxHops int) {
	opt.maxHops = maxHops
}

func (opt *Option) FirstHop() int {
	if opt.firstHop == 0 {
		opt.firstHop = DEFAULT_FIRST_HOP
	}
	return opt.firstHop
}

func (opt *Option) SetFirstHop(firstHop int) {
	opt.firstHop = firstHop
}

func (opt *Option) TimeoutMs() int {
	if opt.timeoutMs == 0 {
		opt.timeoutMs = DEFAULT_TIMEOUT_MS
	}
	return opt.timeoutMs
}

func (opt *Option) SetTimeoutMs(timeoutMs int) {
	opt.timeoutMs = timeoutMs
}

func (opt *Option) NRequeries() int {
	if opt.nqueries == 0 {
		opt.nqueries = DEFAULT_NQUERIES
	}
	return opt.nqueries
}

func (opt *Option) SetNRequeries(nqueries int) {
	opt.nqueries = nqueries
}

func (opt *Option) PacketSize() int {
	if opt.packetSize == 0 {
		opt.packetSize = DEFAULT_PACKET_SIZE
	}
	return opt.packetSize
}

func (opt *Option) SetPacketSize(packetSize int) {
	opt.packetSize = packetSize
}

func (opt *Option) ResolveHost() bool {
	return opt.resolveHost
}

func (opt *Option) EnableResolveHost() {
	opt.resolveHost = true
}

func (opt *Option) DisableResolveHost() {
	opt.resolveHost = false
}

func (opt *Option) Privileged() bool {
	return opt.privileged
}

func (opt *Option) EnablePrivileged() {
	opt.privileged = true
}

func (opt *Option) DisablePrivileged() {
	opt.privileged = false
}

func (opt *Option) FixedDstPort() bool {
	return opt.fixedDstPort
}

func (opt *Option) EnableFixedDstPort() {
	opt.fixedDstPort = true
}

func (opt *Option) DisableFixedDstPort() {
	opt.fixedDstPort = false
}
