package posnode

import (
	"github.com/Fantom-foundation/go-lachesis/src/hash"
)

// Consensus is a consensus interface.
type Consensus interface {
	// PushEvent takes event for processing.
	PushEvent(hash.Event)
}
