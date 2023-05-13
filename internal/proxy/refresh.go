package proxy

import (
	"context"

	"github.com/centrifugal/centrifugo/v5/internal/proxyproto"
)

// RefreshProxy allows to send refresh requests.
type RefreshProxy interface {
	ProxyRefresh(context.Context, *proxyproto.RefreshRequest) (*proxyproto.RefreshResponse, error)
	// Protocol for metrics and logging.
	Protocol() string
	// UseBase64 for bytes in requests from Centrifugo to application backend.
	UseBase64() bool
	// IncludeMeta ...
	IncludeMeta() bool
}
