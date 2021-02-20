package scaler

import "context"

// Scaler is interface of class which scale in/out spanner instance
type Scaler interface {
	// Scale scale in/out spanner instance
	Scale(ctx context.Context, projectID, instanceID string, numNode int32) error
}
