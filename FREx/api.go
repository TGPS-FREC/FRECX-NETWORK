package FREx

import (
	"context"
	"errors"
	"sync"
	"time"
)

const (
	LimitThresholdOrderNonceInQueue = 100
)

// List of errors
var (
	ErrNoTopics          = errors.New("missing topic(s)")
	ErrOrderNonceTooLow  = errors.New("OrderNonce too low")
	ErrOrderNonceTooHigh = errors.New("OrderNonce too high")
)

// PublicFREXAPI provides the FREX RPC service that can be
// use publicly without security implications.
type PublicFREXAPI struct {
	t        *FREX
	mu       sync.Mutex
	lastUsed map[string]time.Time // keeps track when a filter was polled for the last time.

}

// NewPublicFREXAPI create a new RPC FREX service.
func NewPublicFREXAPI(t *FREX) *PublicFREXAPI {
	api := &PublicFREXAPI{
		t:        t,
		lastUsed: make(map[string]time.Time),
	}
	return api
}

// Version returns the FREX sub-protocol version.
func (api *PublicFREXAPI) Version(ctx context.Context) string {
	return ProtocolVersionStr
}
