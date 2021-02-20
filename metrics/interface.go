package metrics

import "context"

// Collector is interface of a spanner metrics collector.
type Collector interface {
	// GetCPUHighPriorityTotal returns CPU utilization (High priority total)
	GetCPUHighPriorityTotal(ctx context.Context, projectID, instanceID string) (float64, error)
	// GetCPU24HourSmoothedAggregate returns CPU utilization (24-hour smoothed aggregate)
	GetCPU24HourSmoothedAggregate(ctx context.Context, projectID, instanceID string) (float64, error)
	// GetStorageUtilizationPerNode returns usage of Storage
	GetStorageUtilizationPerNode(ctx context.Context, projectID, instanceID string) (float64, error)
}
