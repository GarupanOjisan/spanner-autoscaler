package pollers

import (
	"context"
)

// Poller is interface of class which returns a recommended number of spanner node
type Poller interface {
	Poll(ctx context.Context, projectID, instanceID string) (int64, error)
}
