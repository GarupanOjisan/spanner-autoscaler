package mock

import (
	"context"
)

// Scaler is mock Scaler
type Scaler struct {
}

// Scale do nothing
func (s *Scaler) Scale(ctx context.Context, projectID, instanceID string, numNode int32) error {
	return nil
}
