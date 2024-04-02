package FRExlending

import (
	"context"
	"errors"
	"sync"
	"time"
)

// List of errors
var (
	ErrOrderNonceTooLow  = errors.New("OrderNonce too low")
	ErrOrderNonceTooHigh = errors.New("OrderNonce too high")
)

// PublicFREXLendingAPI provides the FREX RPC service that can be
// use publicly without security implications.
type PublicFREXLendingAPI struct {
	t        *Lending
	mu       sync.Mutex
	lastUsed map[string]time.Time // keeps track when a filter was polled for the last time.

}

// NewPublicFREXLendingAPI create a new RPC FREX service.
func NewPublicFREXLendingAPI(t *Lending) *PublicFREXLendingAPI {
	api := &PublicFREXLendingAPI{
		t:        t,
		lastUsed: make(map[string]time.Time),
	}
	return api
}

// Version returns the Lending sub-protocol version.
func (api *PublicFREXLendingAPI) Version(ctx context.Context) string {
	return ProtocolVersionStr
}
