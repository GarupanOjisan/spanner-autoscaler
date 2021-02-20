package voters

import (
	"context"
)

// Voter is interface of class which returns a recommended number of spanner node
type Voter interface {
	Poll(ctx context.Context, projectID, instanceID string) (int64, error)
}
