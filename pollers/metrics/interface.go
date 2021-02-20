package metrics

import (
	"context"
)

// Poller is implements Poller interface
type Poller struct {
}

// Poll returns a recommeneded number of nodes based on metrics below.
// 1. recommended high cpu utilization (https://cloud.google.com/spanner/docs/cpu-utilization#recommended-max)
// 2. recommended limit for storage per node (https://cloud.google.com/spanner/docs/monitoring-cloud#storage)
func (p *Poller) Poll(ctx context.Context, projectID, instanceID string) (int64, error) {
	return 1, nil
}
