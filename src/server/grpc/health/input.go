package health

import "time"

type Initiator interface {
	Init() error
}

// Accumulator is Input depended
type Accumulator interface {
	// AddStatus is the same server status, but different status for function
	AddStatus(server string, val map[string]interface{}, time ...time.Time)

	// AddError report a error
	AddError(server string, err error)
}

// Input is the data enter which health checking
type Input interface {
	Initiator
	Gather(acc Accumulator) error
}

// Output is the data exit which health checking
type Output interface {
	Initiator
	Report([]byte) error
}
