package traceroute

const (
	DEFAULT_PORT        = 33434
	DEFAULT_MAX_HOPS    = 64
	DEFAULT_FIRST_HOP   = 1
	DEFAULT_TIMEOUT_MS  = 500
	DEFAULT_RETRIES     = 3
	DEFAULT_PACKET_SIZE = 52
)

// DefaultOption provides a default option.
var DefaultOption = &Option{
	port:       DEFAULT_PORT,
	maxHops:    DEFAULT_MAX_HOPS,
	firstHop:   DEFAULT_FIRST_HOP,
	timeoutMs:  DEFAULT_TIMEOUT_MS,
	retries:    DEFAULT_RETRIES,
	packetSize: DEFAULT_PACKET_SIZE,
}

// TracrouteOption type
type Option struct {
	port       int
	maxHops    int
	firstHop   int
	timeoutMs  int
	retries    int
	packetSize int
}

func (Option *Option) Port() int {
	if Option.port == 0 {
		Option.port = DEFAULT_PORT
	}
	return Option.port
}

func (Option *Option) SetPort(port int) {
	Option.port = port
}

func (Option *Option) MaxHops() int {
	if Option.maxHops == 0 {
		Option.maxHops = DEFAULT_MAX_HOPS
	}
	return Option.maxHops
}

func (Option *Option) SetMaxHops(maxHops int) {
	Option.maxHops = maxHops
}

func (Option *Option) FirstHop() int {
	if Option.firstHop == 0 {
		Option.firstHop = DEFAULT_FIRST_HOP
	}
	return Option.firstHop
}

func (Option *Option) SetFirstHop(firstHop int) {
	Option.firstHop = firstHop
}

func (Option *Option) TimeoutMs() int {
	if Option.timeoutMs == 0 {
		Option.timeoutMs = DEFAULT_TIMEOUT_MS
	}
	return Option.timeoutMs
}

func (Option *Option) SetTimeoutMs(timeoutMs int) {
	Option.timeoutMs = timeoutMs
}

func (Option *Option) Retries() int {
	if Option.retries == 0 {
		Option.retries = DEFAULT_RETRIES
	}
	return Option.retries
}

func (Option *Option) SetRetries(retries int) {
	Option.retries = retries
}

func (Option *Option) PacketSize() int {
	if Option.packetSize == 0 {
		Option.packetSize = DEFAULT_PACKET_SIZE
	}
	return Option.packetSize
}

func (Option *Option) SetPacketSize(packetSize int) {
	Option.packetSize = packetSize
}
